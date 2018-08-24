package k8sTest

import (
	"fmt"
	"strings"

	. "github.com/cilium/cilium/test/ginkgo-ext"
	"github.com/cilium/cilium/test/helpers"

	. "github.com/onsi/gomega"
)

var _ = Describe("K8sUpdates", func() {

	// This test runs 8 steps as following:
	// 1 - delete all pods. Clean cilium, this can be, and should be achieved by
	// `clean-cilium-state: "true"` option that we have in configmap
	// 2 - install cilium `cilium:v1.1.4`
	// 3 - make endpoints talk with each other with policy
	// 4 - upgrade cilium to `k8s1:5000/cilium/cilium-dev:latest`
	// 5 - make endpoints talk with each other with policy
	// 6 - downgrade cilium to `cilium:v1.1.4`
	// 7 - make endpoints talk with each other with policy
	// 8 - delete all pods. Clean cilium, this can be, and should be achieved by
	// `clean-cilium-state: "true"` option that we have in configmap.
	// This makes sure the upgrade tests won't affect any other test
	// 9 - re install cilium:latest image for remaining tests.

	var (
		kubectl *helpers.Kubectl

		microscopeErr    error
		microscopeCancel = func() error { return nil }
		cleanupCallback  = func() { return }
	)

	BeforeAll(func() {
		kubectl = helpers.CreateKubectl(helpers.K8s1VMName(), logger)

		_ = kubectl.Delete(helpers.DNSDeployment())

		// Delete kube-dns because if not will be a restore the old endpoints
		// from master instead of create the new ones.
		_ = kubectl.DeleteResource(
			"deploy", fmt.Sprintf("-n %s kube-dns", helpers.KubeSystemNamespace))

		// Sometimes PolicyGen has a lot of pods running around without delete
		// it. Using this we are sure that we delete before this test start
		kubectl.Exec(fmt.Sprintf(
			"%s delete --all pods,svc,cnp -n %s", helpers.KubectlCmd, helpers.DefaultNamespace))

		ExpectAllPodsTerminated(kubectl)
	})

	JustBeforeEach(func() {
		microscopeErr, microscopeCancel = kubectl.MicroscopeStart()
		Expect(microscopeErr).To(BeNil(), "Microscope cannot be started")
	})

	AfterAll(func() {
		_ = kubectl.Apply(helpers.DNSDeployment())

		// Re-install the cilium developer image after tests are completed
		err := kubectl.CiliumInstall(helpers.CiliumDefaultDSPatch, helpers.CiliumConfigMapPatch)
		Expect(err).To(BeNil(), "Cilium cannot be installed")

		ExpectCiliumReady(kubectl)
		ExpectCiliumReady(kubectl)
	})

	AfterFailed(func() {
		kubectl.CiliumReport(helpers.KubeSystemNamespace, "cilium endpoint list")
	})

	JustAfterEach(func() {
		kubectl.ValidateNoErrorsOnLogs(CurrentGinkgoTestDescription().Duration)
		Expect(microscopeCancel()).To(BeNil(), "cannot stop microscope")
	})

	AfterEach(func() {
		cleanupCallback()
		ExpectAllPodsTerminated(kubectl)
	})

	BeforeEach(func() {
		// Making sure that we deleted the  cilium ds. No assert message
		// because maybe is not present
		_ = kubectl.DeleteResource("ds", fmt.Sprintf("-n %s cilium", helpers.KubeSystemNamespace))

		// Delete kube-dns because if not will be a restore the old endpoints
		// from master instead of create the new ones.
		_ = kubectl.Delete(helpers.DNSDeployment())

		ExpectAllPodsTerminated(kubectl)

		err := kubectl.CiliumInstallVersion(
			helpers.CiliumDefaultDSPatch,
			"cilium-cm-patch-clean-cilium-state.yaml",
			helpers.CiliumStableVersion,
		)
		Expect(err).To(BeNil(), fmt.Sprintf("Cilium %s was not able to be deployed", helpers.CiliumStableVersion))
		ExpectCiliumReady(kubectl)
	})

	It("Tests upgrade and downgrade from a Cilium stable image to master", func() {
		var assertUpgradeSuccessful func()
		assertUpgradeSuccessful, cleanupCallback =
			ValidateCiliumUpgrades(kubectl, helpers.CiliumStableVersion, helpers.CiliumDeveloperImage)
		assertUpgradeSuccessful()
	})
})

// ValidateCiliumUpgrades tests if the oldVersion can be upgrade to the newVersion
// and if the newVersion can be downgraded to the oldVersion.
// It returns two callbacks, the first one is the assertfunction that need to
// run, and the second one are the cleanup actions
func ValidateCiliumUpgrades(kubectl *helpers.Kubectl, oldVersion, newVersion string) (func(), func()) {
	demoPath := helpers.ManifestGet("demo.yaml")
	l7Policy := helpers.ManifestGet("l7-policy.yaml")
	apps := []string{helpers.App1, helpers.App2, helpers.App3}
	app1Service := "app1-service"

	cleanupCallback := func() {
		kubectl.Delete(l7Policy)
		kubectl.Delete(demoPath)

		// make sure that Kubedns is deleted correctly
		_ = kubectl.Delete(helpers.DNSDeployment())

		// make sure we clean everything up before doing any other test
		err := kubectl.CiliumInstall(
			helpers.CiliumDefaultDSPatch,
			"cilium-cm-patch-clean-cilium-state.yaml",
		)
		Expect(err).To(BeNil(), fmt.Sprintf("Cilium %s was not able to be deployed", newVersion))
		ExpectCiliumReady(kubectl)

		_ = kubectl.DeleteResource(
			"ds", fmt.Sprintf("-n %s cilium", helpers.KubeSystemNamespace))
	}

	testfunc := func() {
		validatedImage := func(image string) {
			By("Checking that installed image is %q", image)

			filter := `{.items[*].status.containerStatuses[0].image}`
			data, err := kubectl.GetPods(
				helpers.KubeSystemNamespace, "-l k8s-app=cilium").Filter(filter)
			ExpectWithOffset(1, err).To(BeNil(), "Cannot get cilium pods")

			for _, val := range strings.Split(data.String(), " ") {
				ExpectWithOffset(1, image).To(ContainSubstring(val), "Cilium image didn't update correctly")
			}
		}

		validateEndpointsConnection := func() {
			err := kubectl.CiliumEndpointWaitReady()
			ExpectWithOffset(1, err).To(BeNil(), "Endpoints are not ready after timeout")

			ExpectKubeDNSReady(kubectl)

			err = kubectl.WaitForKubeDNSEntry(app1Service, helpers.DefaultNamespace)
			ExpectWithOffset(1, err).To(BeNil(), "DNS entry is not ready after timeout")

			err = kubectl.CiliumEndpointWaitReady()
			ExpectWithOffset(1, err).To(BeNil(), "Endpoints are not ready after timeout")

			appPods := helpers.GetAppPods(apps, helpers.DefaultNamespace, kubectl, "id")

			err = kubectl.WaitForKubeDNSEntry(app1Service, helpers.DefaultNamespace)
			ExpectWithOffset(1, err).To(BeNil(), "DNS entry is not ready after timeout")

			By("Making L7 requests between endpoints")
			res := kubectl.ExecPodCmd(
				helpers.DefaultNamespace, appPods[helpers.App2],
				helpers.CurlFail("http://%s/public", app1Service))
			ExpectWithOffset(1, res).Should(helpers.CMDSuccess(), "Cannot curl app1-service")

			res = kubectl.ExecPodCmd(
				helpers.DefaultNamespace, appPods[helpers.App2],
				helpers.CurlFail("http://%s/private", app1Service))
			ExpectWithOffset(1, res).ShouldNot(helpers.CMDSuccess(), "Expect a 403 from app1-service")
		}

		By("Installing kube-dns")
		kubectl.Apply(helpers.DNSDeployment()).ExpectSuccess("Kube-dns cannot be installed")

		By("Creating some endpoints and L7 policy")
		kubectl.Apply(demoPath).ExpectSuccess()

		err := kubectl.WaitforPods(helpers.DefaultNamespace, "-l zgroup=testapp", timeout)
		Expect(err).Should(BeNil(), "Test pods are not ready after timeout")

		ExpectKubeDNSReady(kubectl)

		_, err = kubectl.CiliumPolicyAction(
			helpers.KubeSystemNamespace, l7Policy, helpers.KubectlApply, timeout)
		Expect(err).Should(BeNil(), "cannot import l7 policy: %v", l7Policy)

		validateEndpointsConnection()

		By("Updating cilium to master image")

		waitForUpdateImage := func(image string) func() bool {
			return func() bool {
				pods, err := kubectl.GetCiliumPods(helpers.KubeSystemNamespace)
				if err != nil {
					return false
				}

				filter := `{.items[*].status.containerStatuses[0].image}`
				data, err := kubectl.GetPods(
					helpers.KubeSystemNamespace, "-l k8s-app=cilium").Filter(filter)
				if err != nil {
					return false
				}
				number := strings.Count(data.String(), image)
				if number == len(pods) {
					return true
				}
				log.Infof("Only '%v' of '%v' cilium pods updated to the new image",
					number, len(pods))
				return false
			}
		}

		err = kubectl.CiliumInstall(helpers.CiliumDefaultDSPatch, helpers.CiliumConfigMapPatch)
		Expect(err).To(BeNil(), fmt.Sprintf("Cilium %s was not able to be deployed", newVersion))

		err = helpers.WithTimeout(
			waitForUpdateImage(newVersion),
			"Cilium Pods are not updating correctly",
			&helpers.TimeoutConfig{Timeout: timeout})
		Expect(err).To(BeNil(), "Pods are not updating")

		err = kubectl.WaitforPods(
			helpers.KubeSystemNamespace, "-l k8s-app=cilium", timeout)
		Expect(err).Should(BeNil(), "Cilium is not ready after timeout")

		validatedImage(newVersion)

		validateEndpointsConnection()

		By("Downgrading cilium to %s image", oldVersion)

		err = kubectl.CiliumInstallVersion(
			helpers.CiliumDefaultDSPatch,
			helpers.CiliumConfigMapPatch,
			helpers.CiliumStableVersion,
		)
		Expect(err).To(BeNil(), fmt.Sprintf("Cilium %s was not able to be deployed", helpers.CiliumStableVersion))

		err = helpers.WithTimeout(
			waitForUpdateImage(helpers.CiliumStableVersion),
			"Cilium Pods are not updating correctly",
			&helpers.TimeoutConfig{Timeout: timeout})
		Expect(err).To(BeNil(), "Pods are not updating")

		err = kubectl.WaitforPods(
			helpers.KubeSystemNamespace, "-l k8s-app=cilium", timeout)
		Expect(err).Should(BeNil(), "Cilium is not ready after timeout")

		validatedImage(helpers.CiliumStableImageVersion)

		validateEndpointsConnection()

	}
	return testfunc, cleanupCallback
}
