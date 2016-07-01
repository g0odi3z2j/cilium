package daemon

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/noironetworks/cilium-net/common"
	"github.com/noironetworks/cilium-net/common/types"

	cniTypes "github.com/appc/cni/pkg/types"
	hb "github.com/appc/cni/plugins/ipam/host-local/backend"
	dClient "github.com/docker/engine-api/client"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/op/go-logging"
	k8sClientConfig "k8s.io/kubernetes/pkg/client/restclient"
	k8sClient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/registry/service/ipallocator"
)

var (
	log = logging.MustGetLogger("cilium-net")
)

// Daemon is the cilium daemon that is in charge of perform all necessary plumbing,
// monitoring when a LXC starts.
type Daemon struct {
	ipamConf                  map[types.IPAMType]*types.IPAMConfig
	consul                    *consulAPI.Client
	containers                map[string]*types.Container
	containersMU              sync.Mutex
	endpoints                 map[string]*types.Endpoint
	endpointsDocker           map[string]*types.Endpoint
	endpointsDockerEP         map[string]*types.Endpoint
	endpointsMU               sync.Mutex
	endpointsLearning         map[string]types.LearningLabel
	endpointsLearningMU       sync.Mutex
	endpointsLearningRegister chan types.LearningLabel
	validLabelPrefixesMU      sync.Mutex
	dockerClient              *dClient.Client
	k8sClient                 *k8sClient.Client
	conf                      *Config
	policyTree                types.PolicyTree
	policyTreeMU              sync.Mutex
	cacheIteration            int
	reservedConsumables       []*types.Consumable
	uiTopo                    types.UITopo
	uiListeners               map[*Conn]bool
	uiListenersMU             sync.Mutex
	registerUIListener        chan *Conn
	configMutex               sync.Mutex
}

func createConsulClient(config *consulAPI.Config) (*consulAPI.Client, error) {
	var (
		c   *consulAPI.Client
		err error
	)
	if config != nil {
		c, err = consulAPI.NewClient(config)
	} else {
		c, err = consulAPI.NewClient(consulAPI.DefaultConfig())
	}
	if err != nil {
		return nil, err
	}
	maxRetries := 30
	i := 0
	for {
		leader, err := c.Status().Leader()
		if err != nil || leader == "" {
			log.Info("Waiting for consul client to be ready...")
			time.Sleep(2 * time.Second)
			i++
			if i > maxRetries {
				e := fmt.Errorf("Unable to contact consul")
				log.Error(e)
				return nil, e
			}
		} else {
			log.Info("Consul client ready")
			break
		}
	}
	return c, nil
}

func createDockerClient(endpoint string) (*dClient.Client, error) {
	defaultHeaders := map[string]string{"User-Agent": "cilium"}
	return dClient.NewClient(endpoint, "v1.21", nil, defaultHeaders)
}

func createK8sClient(endpoint string) (*k8sClient.Client, error) {
	config := k8sClientConfig.Config{Host: endpoint}
	k8sClientConfig.SetKubernetesDefaults(&config)
	return k8sClient.New(&config)
}

// NewDaemon creates and returns a new Daemon with the parameters set in c.
func NewDaemon(c *Config) (*Daemon, error) {
	if c == nil {
		return nil, fmt.Errorf("Configuration is nil")
	}
	ones, bits := common.NodeIPv6Mask.Size()
	maskPerIPAMType := ones + 1
	ipamSubnets := net.IPNet{IP: c.NodeAddress, Mask: net.CIDRMask(maskPerIPAMType, bits)}
	cniIPAMSubnet := ipamSubnets
	libnetworkIPAMSubnet := net.IPNet{IP: common.NextNetwork(ipamSubnets), Mask: net.CIDRMask(maskPerIPAMType, bits)}
	nodeRoute := net.IPNet{IP: c.NodeAddress, Mask: common.ContainerIPv6Mask}

	ipamConf := map[types.IPAMType]*types.IPAMConfig{
		types.CNIIPAMType: &types.IPAMConfig{
			IPAMConfig: hb.IPAMConfig{
				Name:    string(types.CNIIPAMType),
				Subnet:  cniTypes.IPNet(cniIPAMSubnet),
				Gateway: c.NodeAddress,
				Routes: []cniTypes.Route{
					cniTypes.Route{
						Dst: nodeRoute,
					},
					cniTypes.Route{
						Dst: net.IPNet{IP: net.IPv6zero, Mask: net.CIDRMask(0, 128)},
						GW:  c.NodeAddress,
					},
				},
			},
			IPAllocator: ipallocator.NewCIDRRange(&cniIPAMSubnet),
		},
		types.LibnetworkIPAMType: &types.IPAMConfig{
			IPAMConfig: hb.IPAMConfig{
				Name:    string(types.LibnetworkIPAMType),
				Subnet:  cniTypes.IPNet(libnetworkIPAMSubnet),
				Gateway: c.NodeAddress,
				Routes: []cniTypes.Route{
					cniTypes.Route{
						Dst: nodeRoute,
					},
					cniTypes.Route{
						Dst: net.IPNet{IP: net.IPv6zero, Mask: net.CIDRMask(0, 128)},
						GW:  c.NodeAddress,
					},
				},
			},
			IPAllocator: ipallocator.NewCIDRRange(&libnetworkIPAMSubnet),
		},
	}

	consul, err := createConsulClient(c.ConsulConfig)
	if err != nil {
		return nil, err
	}

	dockerClient, err := createDockerClient(c.DockerEndpoint)
	if err != nil {
		return nil, err
	}

	k8sClient, err := createK8sClient(c.K8sEndpoint)
	if err != nil {
		return nil, err
	}

	rootNode := types.PolicyTree{
		Root: types.NewPolicyNode(common.GlobalLabelPrefix, nil),
	}

	rootNode.Root.Path()

	d := Daemon{
		conf:                      c,
		ipamConf:                  ipamConf,
		consul:                    consul,
		dockerClient:              dockerClient,
		k8sClient:                 k8sClient,
		containers:                make(map[string]*types.Container),
		endpoints:                 make(map[string]*types.Endpoint),
		endpointsDocker:           make(map[string]*types.Endpoint),
		endpointsDockerEP:         make(map[string]*types.Endpoint),
		endpointsLearning:         make(map[string]types.LearningLabel),
		endpointsLearningRegister: make(chan types.LearningLabel, 1),
		cacheIteration:            1,
		reservedConsumables:       make([]*types.Consumable, 0),
		policyTree:                rootNode,
		uiTopo:                    types.NewUITopo(),
		uiListeners:               make(map[*Conn]bool),
		registerUIListener:        make(chan *Conn, 1),
	}

	if d.conf.IsUIEnabled() {
		d.ListenBuildUIEvents()
	}

	if c.RestoreState {
		if err := d.SyncState(common.CiliumPath, true); err != nil {
			log.Warningf("Error while recovering endpoints: %s\n", err)
		}
	}

	return &d, nil
}

func changedOption(key string, value bool, data interface{}) {
}

func (d *Daemon) Update(opts types.OptionMap) error {
	if err := d.conf.Opts.Validate(opts); err != nil {
		return err
	}

	d.configMutex.Lock()
	defer d.configMutex.Unlock()

	d.conf.Opts.Apply(opts, changedOption, d)

	return nil
}
