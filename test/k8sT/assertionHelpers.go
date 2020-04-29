// Copyright 2018-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sTest

import (
	"fmt"
	"time"

	"github.com/cilium/cilium/test/config"
	. "github.com/cilium/cilium/test/ginkgo-ext"
	"github.com/cilium/cilium/test/helpers"
	. "github.com/onsi/gomega"
)

var longTimeout = 10 * time.Minute

// ExpectKubeDNSReady is a wrapper around helpers/WaitKubeDNS. It asserts that
// the error returned by that function is nil.
func ExpectKubeDNSReady(vm *helpers.Kubectl) {
	By("Waiting for kube-dns to be ready")
	err := vm.WaitKubeDNS()
	ExpectWithOffset(1, err).Should(BeNil(), "kube-dns was not able to get into ready state")

	By("Running kube-dns preflight check")
	err = vm.KubeDNSPreFlightCheck()
	ExpectWithOffset(1, err).Should(BeNil(), "kube-dns service not ready")
}

// ExpectCiliumReady is a wrapper around helpers/WaitForPods. It asserts that
// the error returned by that function is nil.
func ExpectCiliumReady(vm *helpers.Kubectl) {
	err := vm.WaitForCiliumReadiness()
	Expect(err).To(BeNil(), "Timeout while waiting for Cilium to become ready")

	err = vm.CiliumPreFlightCheck()
	ExpectWithOffset(1, err).Should(BeNil(), "cilium pre-flight checks failed")
}

// ExpectCiliumOperatorReady is a wrapper around helpers/WaitForPods. It asserts that
// the error returned by that function is nil.
func ExpectCiliumOperatorReady(vm *helpers.Kubectl) {
	By("Waiting for cilium-operator to be ready")
	err := vm.WaitforPods(helpers.CiliumNamespace, "-l name=cilium-operator", longTimeout)
	ExpectWithOffset(1, err).Should(BeNil(), "Cilium operator was not able to get into ready state")
}

// ExpectAllPodsTerminated is a wrapper around helpers/WaitCleanAllTerminatingPods.
// It asserts that the error returned by that function is nil.
func ExpectAllPodsTerminated(vm *helpers.Kubectl) {
	err := vm.WaitCleanAllTerminatingPods(helpers.HelperTimeout)
	ExpectWithOffset(1, err).To(BeNil(), "terminating containers are not deleted after timeout")
}

// ExpectCiliumPreFlightInstallReady is a wrapper around helpers/WaitForNPods.
// It asserts the error returned by that function is nil.
func ExpectCiliumPreFlightInstallReady(vm *helpers.Kubectl) {
	By("Waiting for all cilium pre-flight pods to be ready")

	err := vm.WaitforPods(helpers.CiliumNamespace, "-l k8s-app=cilium-pre-flight-check", longTimeout)
	warningMessage := ""
	if err != nil {
		res := vm.Exec(fmt.Sprintf(
			"%s -n %s get pods -l k8s-app=cilium-pre-flight-check",
			helpers.KubectlCmd, helpers.CiliumNamespace))
		warningMessage = res.Output().String()
	}
	Expect(err).To(BeNil(), "cilium pre-flight check is not ready after timeout, pods status:\n %s", warningMessage)
}

// DeployCiliumAndDNS deploys DNS and cilium into the kubernetes cluster
func DeployCiliumAndDNS(vm *helpers.Kubectl, ciliumFilename string) {
	DeployCiliumOptionsAndDNS(vm, ciliumFilename, map[string]string{"global.debug.verbose": "flow"})
}

func redeployCilium(vm *helpers.Kubectl, ciliumFilename string, options map[string]string) {
	By("Installing Cilium")
	err := vm.CiliumInstall(ciliumFilename, options)
	Expect(err).To(BeNil(), "Cilium cannot be installed")

	err = vm.WaitForCiliumReadiness()
	Expect(err).To(BeNil(), "Timeout while waiting for Cilium to become ready")
}

// RedeployCilium reinstantiates the Cilium DS and ensures it is running.
//
// This helper is only appropriate for reconfiguring Cilium in the middle of
// an existing testsuite that calls DeployCiliumAndDNS(...).
func RedeployCilium(vm *helpers.Kubectl, ciliumFilename string, options map[string]string) {
	redeployCilium(vm, ciliumFilename, options)
	ExpectCiliumReady(vm)
	ExpectCiliumOperatorReady(vm)
}

// DeployCiliumOptionsAndDNS deploys DNS and cilium with options into the kubernetes cluster
func DeployCiliumOptionsAndDNS(vm *helpers.Kubectl, ciliumFilename string, options map[string]string) {
	redeployCilium(vm, ciliumFilename, options)

	By("Installing DNS Deployment")
	switch helpers.GetCurrentIntegration() {
	case helpers.CIIntegrationMicrok8s:
		By(fmt.Sprintf("%s (hint: %s)",
			"Assuming that microk8s already has DNS deployed...",
			"Use 'microk8s.enable dns' to create deployment"))
	case helpers.CIIntegrationGKE:
		By("Restarting all kube-system pods")
		if res := vm.DeleteResource("pod", fmt.Sprintf("-n %s --all", helpers.KubeSystemNamespace)); !res.WasSuccessful() {
			log.Warningf("Unable to delete kube-system pods: %s", res.OutputPrettyPrint())
		}
	default:
		vm.ApplyDefault(helpers.DNSDeployment(vm.BasePath()))
		By("Restarting DNS Pods")
		if res := vm.DeleteResource("pod", fmt.Sprintf("-n %s -l k8s-app=kube-dns", helpers.KubeSystemNamespace)); !res.WasSuccessful() {
			log.Warningf("Unable to delete DNS pods: %s", res.OutputPrettyPrint())
		}
	}

	switch helpers.GetCurrentIntegration() {
	case helpers.CIIntegrationFlannel:
		By("Installing Flannel")
		vm.ApplyDefault(vm.GetFilePath("../examples/kubernetes/addons/flannel/flannel.yaml"))
	default:
	}

	ExpectCiliumReady(vm)
	ExpectCiliumOperatorReady(vm)
	ExpectKubeDNSReady(vm)

	switch helpers.GetCurrentIntegration() {
	case helpers.CIIntegrationGKE:
		err := vm.WaitforPods(helpers.KubeSystemNamespace, "", longTimeout)
		ExpectWithOffset(1, err).Should(BeNil(), "kube-system pods were not able to get into ready state after restart")
	}
}

// SkipIfBenchmark will skip the test if benchmark is not specified
func SkipIfBenchmark() {
	if !config.CiliumTestConfig.Benchmarks {
		Skip("Benchmarks are skipped, specify -cilium.Benchmarks")
	}
}

// SkipIfIntegration will skip a test if it's running with any of the specified
// integration.
func SkipIfIntegration(integration string) {
	if helpers.IsIntegration(integration) {
		Skip(fmt.Sprintf(
			"This feature is not supported in Cilium %q mode. Skipping test.",
			integration))
	}
}

// SkipItIfNoKubeProxy will skip It if kube-proxy is disabled (= NodePort BPF is
// enabled)
func SkipItIfNoKubeProxy() {
	if !helpers.RunsWithKubeProxy() {
		Skip("kube-proxy is disabled (NodePort BPF is enabled). Skipping test.")
	}
}

func deleteETCDOperator(kubectl *helpers.Kubectl) {
	// Do not assert on success in AfterEach intentionally to avoid
	// incomplete teardown.
	_ = kubectl.DeleteResource("deploy", fmt.Sprintf("-n %s -l io.cilium/app=etcd-operator", helpers.CiliumNamespace))
	_ = kubectl.DeleteResource("pod", fmt.Sprintf("-n %s -l io.cilium/app=etcd-operator", helpers.CiliumNamespace))
	_ = kubectl.WaitCleanAllTerminatingPods(helpers.HelperTimeout)
}
