// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"

	"github.com/cilium/cilium/operator/api"
	operatorMetrics "github.com/cilium/cilium/operator/metrics"
	operatorOption "github.com/cilium/cilium/operator/option"
	ces "github.com/cilium/cilium/operator/pkg/ciliumendpointslice"
	"github.com/cilium/cilium/operator/pkg/ingress"
	operatorWatchers "github.com/cilium/cilium/operator/watchers"

	"github.com/cilium/cilium/pkg/components"
	"github.com/cilium/cilium/pkg/defaults"
	"github.com/cilium/cilium/pkg/gops"
	"github.com/cilium/cilium/pkg/hive"
	"github.com/cilium/cilium/pkg/hive/cell"
	"github.com/cilium/cilium/pkg/ipam/allocator"
	ipamOption "github.com/cilium/cilium/pkg/ipam/option"
	"github.com/cilium/cilium/pkg/k8s"
	"github.com/cilium/cilium/pkg/k8s/apis/cilium.io/client"
	k8sClient "github.com/cilium/cilium/pkg/k8s/client"
	k8sversion "github.com/cilium/cilium/pkg/k8s/version"
	"github.com/cilium/cilium/pkg/kvstore"
	"github.com/cilium/cilium/pkg/logging"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/metrics"
	"github.com/cilium/cilium/pkg/option"
	"github.com/cilium/cilium/pkg/pprof"
	"github.com/cilium/cilium/pkg/rand"
	"github.com/cilium/cilium/pkg/rate"
	"github.com/cilium/cilium/pkg/version"
)

var (
	binaryName = filepath.Base(os.Args[0])

	log = logging.DefaultLogger.WithField(logfields.LogSubsys, binaryName)

	rootCmd = &cobra.Command{
		Use:   binaryName,
		Short: "Run " + binaryName,
	}

	shutdownSignal = make(chan struct{})

	leaderElectionResourceLockName = "cilium-operator-resource-lock"

	// Use a Go context so we can tell the leaderelection code when we
	// want to step down
	leaderElectionCtx, leaderElectionCtxCancel = context.WithCancel(context.Background())

	// identityRateLimiter is a rate limiter to rate limit the number of
	// identities being GCed by the operator. See the documentation of
	// rate.Limiter to understand its difference than 'x/time/rate.Limiter'.
	//
	// With our rate.Limiter implementation Cilium will be able to handle bursts
	// of identities being garbage collected with the help of the functionality
	// provided by the 'policy-trigger-interval' in the cilium-agent. With the
	// policy-trigger even if we receive N identity changes over the interval
	// set, Cilium will only need to process all of them at once instead of
	// processing each one individually.
	identityRateLimiter *rate.Limiter

	operatorAddr string

	// IsLeader is an atomic boolean value that is true when the Operator is
	// elected leader. Otherwise, it is false.
	IsLeader atomic.Value

	// OperatorCell are the operator specific cells without infrastructure cells.
	// Used also in tests.
	OperatorCell = cell.Module(
		"operator",

		cell.Invoke(
			registerOperatorHooks,
		),

		// These cells are started only after the operator is elected leader.
		WithLeaderLifecycle(
			legacyCell,
		),
	)

	operatorHive *hive.Hive = newOperatorHive()

	Vp *viper.Viper = operatorHive.Viper()
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func registerOperatorHooks(lc hive.Lifecycle, llc *LeaderLifecycle, clientset k8sClient.Clientset, shutdowner hive.Shutdowner) {

	lc.Append(hive.Hook{
		OnStart: func(context.Context) error {
			go runOperator(llc, clientset, shutdowner)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := llc.Stop(ctx); err != nil {
				return err
			}
			doCleanup()
			return nil
		},
	})
}

func newOperatorHive() *hive.Hive {
	h := hive.New(
		gops.Cell(defaults.GopsPortOperator),
		k8sClient.Cell,
		OperatorCell,
	)
	h.RegisterFlags(rootCmd.Flags())

	// Enable fallback to direct API probing to check for support of Leases in
	// case Discovery API fails.
	h.Viper().Set(option.K8sEnableAPIDiscovery, true)

	return h
}

func init() {
	rootCmd.AddCommand(MetricsCmd)

	rootCmd.Run = func(cobraCmd *cobra.Command, args []string) {
		cmdRefDir := operatorHive.Viper().GetString(option.CMDRef)
		if cmdRefDir != "" {
			genMarkdown(cobraCmd, cmdRefDir)
			os.Exit(0)
		}

		initEnv()

		if err := operatorHive.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func initEnv() {
	vp := operatorHive.Viper()
	// Prepopulate option.Config with options from CLI.
	option.Config.Populate(vp)
	operatorOption.Config.Populate(vp)
	operatorAddr = vp.GetString(operatorOption.OperatorAPIServeAddr)

	// add hooks after setting up metrics in the option.Confog
	logging.DefaultLogger.Hooks.Add(metrics.NewLoggingHook(components.CiliumOperatortName))

	// Logging should always be bootstrapped first. Do not add any code above this!
	if err := logging.SetupLogging(option.Config.LogDriver, logging.LogOptions(option.Config.LogOpt), binaryName, option.Config.Debug); err != nil {
		log.Fatal(err)
	}

	option.LogRegisteredOptions(vp, log)
}

func doCleanup() {
	IsLeader.Store(false)
	close(shutdownSignal)

	// Cancelling this conext here makes sure that if the operator hold the
	// leader lease, it will be released.
	leaderElectionCtxCancel()
}

func getAPIServerAddr() []string {
	if operatorOption.Config.OperatorAPIServeAddr == "" {
		return []string{"127.0.0.1:0", "[::1]:0"}
	}
	return []string{operatorOption.Config.OperatorAPIServeAddr}
}

// checkStatus checks the connection status to the kvstore and
// k8s apiserver and returns an error if any of them is unhealthy
func checkStatus(clientset k8sClient.Clientset) error {
	if kvstoreEnabled() {
		// We check if we are the leader here because only the leader has
		// access to the kvstore client. Otherwise, the kvstore client check
		// will block. It is safe for a non-leader to skip this check, as the
		// it is the leader's responsibility to report the status of the
		// kvstore client.
		if leader, ok := IsLeader.Load().(bool); ok && leader {
			if client := kvstore.Client(); client == nil {
				return fmt.Errorf("kvstore client not configured")
			} else if _, err := client.Status(); err != nil {
				return err
			}
		}
	}

	if _, err := clientset.Discovery().ServerVersion(); err != nil {
		return err
	}

	return nil
}

// runOperator implements the logic of leader election for cilium-operator using
// built-in leader election capbility in kubernetes.
// See: https://github.com/kubernetes/client-go/blob/master/examples/leader-election/main.go
func runOperator(lc *LeaderLifecycle, clientset k8sClient.Clientset, shutdowner hive.Shutdowner) {
	log.Infof("Cilium Operator %s", version.Version)

	allSystemsGo := make(chan struct{})
	IsLeader.Store(false)

	// Configure API server for the operator.
	srv, err := api.NewServer(shutdownSignal, allSystemsGo, getAPIServerAddr()...)
	if err != nil {
		log.WithError(err).Fatalf("Unable to create operator apiserver")
	}
	close(allSystemsGo)

	if operatorOption.Config.EnableK8s {
		k8s.SetClients(clientset, clientset.Slim(), clientset, clientset)

		go func() {
			err = srv.WithStatusCheckFunc(func() error { return checkStatus(clientset) }).StartServer()
			if err != nil {
				log.WithError(err).Fatalf("Unable to start operator apiserver")
			}
		}()
	}

	if operatorOption.Config.EnableMetrics {
		operatorMetrics.Register()
	}

	if operatorOption.Config.PProf {
		pprof.Enable(operatorOption.Config.PProfPort)
	}

	if k8s.IsEnabled() {
		capabilities := k8sversion.Capabilities()
		if !capabilities.MinimalVersionMet {
			log.Fatalf("Minimal kubernetes version not met: %s < %s",
				k8sversion.Version(), k8sversion.MinimalVersionConstraint)
		}
	}

	// Register the CRDs after validating that we are running on a supported
	// version of K8s.
	if k8s.IsEnabled() && !operatorOption.Config.SkipCRDCreation {
		if err := client.RegisterCRDs(); err != nil {
			log.WithError(err).Fatal("Unable to register CRDs")
		}
	} else {
		log.Info("Skipping creation of CRDs")
	}

	// We only support Operator in HA mode for Kubernetes Versions having support for
	// LeasesResourceLock.
	// See docs on capabilities.LeasesResourceLock for more context.
	if !k8sversion.Capabilities().LeasesResourceLock {
		log.Info("Support for coordination.k8s.io/v1 not present, fallback to non HA mode")

		if err := lc.Start(leaderElectionCtx); err != nil {
			log.WithError(err).Fatal("Failed to start leading")
		}
		return
	}

	// Get hostname for identity name of the lease lock holder.
	// We identify the leader of the operator cluster using hostname.
	operatorID, err := os.Hostname()
	if err != nil {
		log.WithError(err).Fatal("Failed to get hostname when generating lease lock identity")
	}
	operatorID = rand.RandomStringWithPrefix(operatorID+"-", 10)

	ns := option.Config.K8sNamespace
	// If due to any reason the CILIUM_K8S_NAMESPACE is not set we assume the operator
	// to be in default namespace.
	if ns == "" {
		ns = metav1.NamespaceDefault
	}

	leResourceLock := &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      leaderElectionResourceLockName,
			Namespace: ns,
		},
		Client: clientset.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			// Identity name of the lock holder
			Identity: operatorID,
		},
	}

	// Start the leader election for running cilium-operators
	log.Info("Waiting for leader election")
	leaderelection.RunOrDie(leaderElectionCtx, leaderelection.LeaderElectionConfig{
		Name: leaderElectionResourceLockName,

		Lock:            leResourceLock,
		ReleaseOnCancel: true,

		LeaseDuration: operatorOption.Config.LeaderElectionLeaseDuration,
		RenewDeadline: operatorOption.Config.LeaderElectionRenewDeadline,
		RetryPeriod:   operatorOption.Config.LeaderElectionRetryPeriod,

		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				if err := lc.Start(ctx); err != nil {
					log.WithError(err).Error("Failed to start when elected leader, shutting down")
					shutdowner.Shutdown(hive.ShutdownWithError(err))
				}
			},
			OnStoppedLeading: func() {
				log.WithField("operator-id", operatorID).Info("Leader election lost")
				// Cleanup everything here, and exit.
				shutdowner.Shutdown(hive.ShutdownWithError(errors.New("Leader election lost")))
			},
			OnNewLeader: func(identity string) {
				if identity == operatorID {
					log.Info("Leading the operator HA deployment")
				} else {
					log.WithFields(logrus.Fields{
						"newLeader":  identity,
						"operatorID": operatorID,
					}).Info("Leader re-election complete")
				}
			},
		},
	})
}

func kvstoreEnabled() bool {
	if option.Config.KVStore == "" {
		return false
	}

	return option.Config.IdentityAllocationMode == option.IdentityAllocationModeKVstore ||
		operatorOption.Config.SyncK8sServices ||
		operatorOption.Config.SyncK8sNodes
}

var legacyCell = cell.Invoke(registerLegacyOnLeader)

func registerLegacyOnLeader(lc hive.Lifecycle, clientSet k8sClient.Clientset) {
	ctx, cancel := context.WithCancel(context.Background())
	legacy := &legacyOnLeader{
		ctx:       ctx,
		cancel:    cancel,
		clientset: clientSet,
	}
	lc.Append(hive.Hook{
		OnStart: legacy.onStart,
		OnStop:  legacy.onStop,
	})
}

type legacyOnLeader struct {
	ctx       context.Context
	cancel    context.CancelFunc
	clientset k8sClient.Clientset
}

func (legacy *legacyOnLeader) onStop(_ context.Context) error {
	legacy.cancel()
	return nil
}

// OnOperatorStartLeading is the function called once the operator starts leading
// in HA mode.
func (legacy *legacyOnLeader) onStart(_ context.Context) error {
	IsLeader.Store(true)

	// If CiliumEndpointSlice feature is enabled, create CESController, start CEP watcher and run controller.
	if k8s.IsEnabled() && !option.Config.DisableCiliumEndpointCRD && option.Config.EnableCiliumEndpointSlice {
		log.Info("Create and run CES controller, start CEP watcher")
		// Initialize  the CES controller
		cesController := ces.NewCESController(legacy.clientset,
			operatorOption.Config.CESMaxCEPsInCES,
			operatorOption.Config.CESSlicingMode,
			float64(legacy.clientset.Config().K8sClientQPS),
			legacy.clientset.Config().K8sClientBurst)
		stopCh := make(chan struct{})
		// Start CEP watcher
		operatorWatchers.CiliumEndpointsSliceInit(legacy.clientset, cesController)
		// Start the CES controller, after current CEPs are synced locally in cache.
		go cesController.Run(operatorWatchers.CiliumEndpointStore, stopCh)
	}

	// Restart kube-dns as soon as possible since it helps etcd-operator to be
	// properly setup. If kube-dns is not managed by Cilium it can prevent
	// etcd from reaching out kube-dns in EKS.
	// If this logic is modified, make sure the operator's clusterrole logic for
	// pods/delete is also up-to-date.
	if !k8s.IsEnabled() {
		log.Infof("KubeDNS unmanaged pods controller disabled due to kubernetes support not enabled")
	} else if option.Config.DisableCiliumEndpointCRD {
		log.Infof("KubeDNS unmanaged pods controller disabled as %q option is set to 'disabled' in Cilium ConfigMap", option.DisableCiliumEndpointCRDName)
	} else if operatorOption.Config.UnmanagedPodWatcherInterval != 0 {
		go enableUnmanagedKubeDNSController(legacy.clientset)
	}

	var (
		nodeManager allocator.NodeEventHandler
		err         error
		withKVStore bool
	)

	log.WithField(logfields.Mode, option.Config.IPAM).Info("Initializing IPAM")

	switch ipamMode := option.Config.IPAM; ipamMode {
	case ipamOption.IPAMAzure, ipamOption.IPAMENI, ipamOption.IPAMClusterPool, ipamOption.IPAMClusterPoolV2, ipamOption.IPAMAlibabaCloud:
		alloc, providerBuiltin := allocatorProviders[ipamMode]
		if !providerBuiltin {
			log.Fatalf("%s allocator is not supported by this version of %s", ipamMode, binaryName)
		}

		if err := alloc.Init(legacy.ctx); err != nil {
			log.WithError(err).Fatalf("Unable to init %s allocator", ipamMode)
		}

		nm, err := alloc.Start(legacy.ctx, &ciliumNodeUpdateImplementation{legacy.clientset})
		if err != nil {
			log.WithError(err).Fatalf("Unable to start %s allocator", ipamMode)
		}

		nodeManager = nm
	}

	if operatorOption.Config.BGPAnnounceLBIP {
		log.Info("Starting LB IP allocator")
		operatorWatchers.StartLBIPAllocator(legacy.ctx, option.Config, legacy.clientset)
	}

	if kvstoreEnabled() {
		var goopts *kvstore.ExtraOptions
		scopedLog := log.WithFields(logrus.Fields{
			"kvstore": option.Config.KVStore,
			"address": option.Config.KVStoreOpt[fmt.Sprintf("%s.address", option.Config.KVStore)],
		})

		if k8s.IsEnabled() && operatorOption.Config.SyncK8sServices {
			operatorWatchers.StartSynchronizingServices(legacy.clientset, true, option.Config)
			// If K8s is enabled we can do the service translation automagically by
			// looking at services from k8s and retrieve the service IP from that.
			// This makes cilium to not depend on kube dns to interact with etcd
			if k8s.IsEnabled() {
				svcURL, isETCDOperator := kvstore.IsEtcdOperator(option.Config.KVStore, option.Config.KVStoreOpt, option.Config.K8sNamespace)
				if isETCDOperator {
					scopedLog.Infof("%s running with service synchronization: automatic etcd service translation enabled", binaryName)

					svcGetter := k8s.ServiceIPGetter(&operatorWatchers.K8sSvcCache)

					name, namespace, err := kvstore.SplitK8sServiceURL(svcURL)
					if err != nil {
						// If we couldn't derive the name/namespace for the given
						// svcURL log the error so the user can see it.
						// k8s.CreateCustomDialer won't be able to derive
						// the name/namespace as well so it does not matter that
						// we wait for all services to be synchronized with k8s.
						scopedLog.WithError(err).WithFields(logrus.Fields{
							"url": svcURL,
						}).Error("Unable to derive service name from given url")
					} else {
						scopedLog.WithFields(logrus.Fields{
							logfields.ServiceName:      name,
							logfields.ServiceNamespace: namespace,
						}).Info("Retrieving service spec from k8s to perform automatic etcd service translation")
						k8sSvc, err := legacy.clientset.CoreV1().Services(namespace).Get(legacy.ctx, name, metav1.GetOptions{})
						switch {
						case err == nil:
							// Create another service cache that contains the
							// k8s service for etcd. As soon the k8s caches are
							// synced, this hijack will stop happening.
							sc := k8s.NewServiceCache(nil)
							slimSvcObj := k8s.ConvertToK8sService(k8sSvc)
							slimSvc := k8s.ObjToV1Services(slimSvcObj)
							if slimSvc == nil {
								// This will never happen but still log it
								scopedLog.Warnf("BUG: invalid k8s service: %s", slimSvcObj)
							}
							sc.UpdateService(slimSvc, nil)
							svcGetter = operatorWatchers.NewServiceGetter(&sc)
						case k8sErrors.IsNotFound(err):
							scopedLog.Error("Service not found in k8s")
						default:
							scopedLog.Warning("Unable to get service spec from k8s, this might cause network disruptions with etcd")
						}
					}

					log := log.WithField(logfields.LogSubsys, "etcd")
					goopts = &kvstore.ExtraOptions{
						DialOption: []grpc.DialOption{
							grpc.WithContextDialer(k8s.CreateCustomDialer(svcGetter, log)),
						},
					}
				}
			}
		} else {
			scopedLog.Infof("%s running without service synchronization: automatic etcd service translation disabled", binaryName)
		}
		scopedLog.Info("Connecting to kvstore")
		if err := kvstore.Setup(legacy.ctx, option.Config.KVStore, option.Config.KVStoreOpt, goopts); err != nil {
			scopedLog.WithError(err).Fatal("Unable to setup kvstore")
		}

		if k8s.IsEnabled() && operatorOption.Config.SyncK8sNodes {
			withKVStore = true
		}

		startKvstoreWatchdog()
	}

	if k8s.IsEnabled() &&
		(operatorOption.Config.RemoveCiliumNodeTaints || operatorOption.Config.SetCiliumIsUpCondition) {
		stopCh := make(chan struct{})

		log.WithFields(logrus.Fields{
			logfields.K8sNamespace:       operatorOption.Config.CiliumK8sNamespace,
			"label-selector":             operatorOption.Config.CiliumPodLabels,
			"remove-cilium-node-taints":  operatorOption.Config.RemoveCiliumNodeTaints,
			"set-cilium-is-up-condition": operatorOption.Config.SetCiliumIsUpCondition,
		}).Info("Removing Cilium Node Taints or Setting Cilium Is Up Condition for Kubernetes Nodes")

		operatorWatchers.HandleNodeTolerationAndTaints(legacy.clientset, stopCh)
	}

	if k8s.IsEnabled() {
		if err := startSynchronizingCiliumNodes(legacy.ctx, legacy.clientset, nodeManager, withKVStore); err != nil {
			log.WithError(err).Fatal("Unable to setup node watcher")
		}

		if operatorOption.Config.CNPNodeStatusGCInterval != 0 {
			RunCNPNodeStatusGC(legacy.clientset, ciliumNodeStore)
		}

		if operatorOption.Config.NodesGCInterval != 0 {
			operatorWatchers.RunCiliumNodeGC(legacy.ctx, legacy.clientset, ciliumNodeStore, operatorOption.Config.NodesGCInterval)
		}
	}

	if option.Config.IPAM == ipamOption.IPAMClusterPool || option.Config.IPAM == ipamOption.IPAMClusterPoolV2 {
		// We will use CiliumNodes as the source of truth for the podCIDRs.
		// Once the CiliumNodes are synchronized with the operator we will
		// be able to watch for K8s Node events which they will be used
		// to create the remaining CiliumNodes.
		<-k8sCiliumNodesCacheSynced

		// We don't want CiliumNodes that don't have podCIDRs to be
		// allocated with a podCIDR already being used by another node.
		// For this reason we will call Resync after all CiliumNodes are
		// synced with the operator to signal the node manager, since it
		// knows all podCIDRs that are currently set in the cluster, that
		// it can allocate podCIDRs for the nodes that don't have a podCIDR
		// set.
		nodeManager.Resync(legacy.ctx, time.Time{})
	}

	if operatorOption.Config.IdentityGCInterval != 0 {
		identityRateLimiter = rate.NewLimiter(
			operatorOption.Config.IdentityGCRateInterval,
			operatorOption.Config.IdentityGCRateLimit,
		)
	}

	switch option.Config.IdentityAllocationMode {
	case option.IdentityAllocationModeCRD:
		if !k8s.IsEnabled() {
			log.Fatal("CRD Identity allocation mode requires k8s to be configured.")
		}

		startManagingK8sIdentities(legacy.clientset)

		if operatorOption.Config.IdentityGCInterval != 0 {
			go startCRDIdentityGC(legacy.clientset)
		}
	case option.IdentityAllocationModeKVstore:
		if operatorOption.Config.IdentityGCInterval != 0 {
			startKvstoreIdentityGC()
		}
	}

	if k8s.IsEnabled() {
		if operatorOption.Config.EndpointGCInterval != 0 {
			enableCiliumEndpointSyncGC(legacy.clientset, false)
		} else {
			// Even if the EndpointGC is disabled we still want it to run at least
			// once. This is to prevent leftover CEPs from populating ipcache with
			// stale entries.
			enableCiliumEndpointSyncGC(legacy.clientset, true)
		}

		err = enableCNPWatcher(legacy.clientset)
		if err != nil {
			log.WithError(err).WithField(logfields.LogSubsys, "CNPWatcher").Fatal(
				"Cannot connect to Kubernetes apiserver ")
		}

		err = enableCCNPWatcher(legacy.clientset)
		if err != nil {
			log.WithError(err).WithField(logfields.LogSubsys, "CCNPWatcher").Fatal(
				"Cannot connect to Kubernetes apiserver ")
		}
	}

	if operatorOption.Config.EnableIngressController {
		ingressController, err := ingress.NewController(
			legacy.clientset,
			ingress.WithHTTPSEnforced(operatorOption.Config.EnforceIngressHTTPS),
			ingress.WithSecretsSyncEnabled(operatorOption.Config.EnableIngressSecretsSync),
			ingress.WithSecretsNamespace(operatorOption.Config.IngressSecretsNamespace),
			ingress.WithLBAnnotationPrefixes(operatorOption.Config.IngressLBAnnotationPrefixes),
			ingress.WithCiliumNamespace(operatorOption.Config.CiliumK8sNamespace),
			ingress.WithSharedLBServiceName(operatorOption.Config.IngressSharedLBServiceName),
			ingress.WithDefaultLoadbalancerMode(operatorOption.Config.IngressDefaultLoadbalancerMode),
		)
		if err != nil {
			log.WithError(err).WithField(logfields.LogSubsys, ingress.Subsys).Fatal(
				"Failed to start ingress controller")
		}
		go ingressController.Run()
	}

	log.Info("Initialization complete")
	return nil
}

// ResetCiliumNodesCacheSyncedStatus resets the current status of
// cache synchronization in Cilium nodes as "not synced".
// Should be used in control-plane testing only to reset the operator status
// before executing the next test case.
func ResetCiliumNodesCacheSyncedStatus() {
	k8sCiliumNodesCacheSynced = make(chan struct{})
}
