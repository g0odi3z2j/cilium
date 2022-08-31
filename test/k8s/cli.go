// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package k8sTest

import (
	"fmt"
	"strings"

	. "github.com/onsi/gomega"

	. "github.com/cilium/cilium/test/ginkgo-ext"
	"github.com/cilium/cilium/test/helpers"
)

var _ = Describe("K8sCLI", func() {
	SkipContextIf(func() bool {
		return helpers.DoesNotRunOnGKE() && helpers.DoesNotRunOnEKS()
	}, "CLI", func() {
		var kubectl *helpers.Kubectl
		var ciliumFilename string

		BeforeAll(func() {
			kubectl = helpers.CreateKubectl(helpers.K8s1VMName(), logger)

			ciliumFilename = helpers.TimestampFilename("cilium.yaml")
			DeployCiliumAndDNS(kubectl, ciliumFilename)
			ExpectCiliumReady(kubectl)
		})

		AfterAll(func() {
			UninstallCiliumFromManifest(kubectl, ciliumFilename)
		})

		JustAfterEach(func() {
			kubectl.ValidateNoErrorsInLogs(CurrentGinkgoTestDescription().Duration)
		})

		Context("Identity CLI testing", func() {
			const (
				manifestYAML = "test-cli.yaml"
				fooID        = "foo"
				fooNode      = "k8s1"
				// These labels are automatically added to all pods in the default namespace.
				defaultLabels = "k8s:io.cilium.k8s.namespace.labels.kubernetes.io/metadata.name=default " +
					"k8s:io.cilium.k8s.policy.cluster=default " +
					"k8s:io.cilium.k8s.policy.serviceaccount=default " +
					"k8s:io.kubernetes.pod.namespace=default"
			)

			var (
				cliManifest string
				ciliumPod   string
				err         error
				identity    int64
			)

			BeforeAll(func() {
				cliManifest = helpers.ManifestGet(kubectl.BasePath(), manifestYAML)
				res := kubectl.ApplyDefault(cliManifest)
				res.ExpectSuccess("Unable to apply %s", cliManifest)
				err = kubectl.WaitforPods(helpers.DefaultNamespace, "-l id", helpers.HelperTimeout)
				Expect(err).Should(BeNil(), "The pods were not ready after timeout")

				ciliumPod, err = kubectl.GetCiliumPodOnNode(fooNode)
				Expect(err).Should(BeNil())

				err := kubectl.WaitForCEPIdentity(helpers.DefaultNamespace, fooID)
				Expect(err).Should(BeNil())

				ep, err := kubectl.GetCiliumEndpoint(helpers.DefaultNamespace, fooID)
				Expect(err).Should(BeNil(), fmt.Sprintf("Unable to get CEP for pod %s", fooID))
				identity = ep.Identity.ID
			})

			AfterAll(func() {
				_ = kubectl.Delete(cliManifest)
				ExpectAllPodsTerminated(kubectl)
			})

			It("Test identity list", func() {
				By("Testing 'cilium identity list' for an endpoint's identity")
				cmd := fmt.Sprintf("cilium identity list k8s:id=%s %s", fooID, defaultLabels)
				res := kubectl.ExecPodCmd(helpers.CiliumNamespace, ciliumPod, cmd)
				res.ExpectSuccess(fmt.Sprintf("Unable to get identity list output for label k8s:id=%s %s", fooID, defaultLabels))

				resSingleOut := res.SingleOut()
				containsIdentity := strings.Contains(resSingleOut, fmt.Sprintf("%d", identity))
				Expect(containsIdentity).To(BeTrue(), "Identity %d of endpoint %s not in 'cilium identity list' output", identity, resSingleOut)
			})

			It("Test cilium bpf metrics list", func() {
				demoManifest := helpers.ManifestGet(kubectl.BasePath(), "demo-named-port.yaml")
				app1Service := "app1-service"
				l3L4DenyPolicy := helpers.ManifestGet(kubectl.BasePath(), "l3-l4-policy-deny.yaml")

				namespaceForTest := helpers.GenerateNamespaceForTest("")
				kubectl.NamespaceDelete(namespaceForTest)
				kubectl.NamespaceCreate(namespaceForTest).ExpectSuccess("could not create namespace")
				kubectl.Apply(helpers.ApplyOptions{FilePath: demoManifest, Namespace: namespaceForTest}).ExpectSuccess("could not create resource")

				err := kubectl.WaitforPods(namespaceForTest, "-l zgroup=testapp", helpers.HelperTimeout)
				Expect(err).To(BeNil(),
					"testapp pods are not ready after timeout in namespace %q", namespaceForTest)

				_, err = kubectl.CiliumPolicyAction(
					namespaceForTest, l3L4DenyPolicy, helpers.KubectlApply, helpers.HelperTimeout)
				Expect(err).Should(BeNil(), "Cannot apply L3 Deny Policy")

				ciliumPodK8s1, err := kubectl.GetCiliumPodOnNode(helpers.K8s1)
				ExpectWithOffset(2, err).Should(BeNil(), "Cannot get cilium pod on k8s1")
				ciliumPodK8s2, err := kubectl.GetCiliumPodOnNode(helpers.K8s2)
				ExpectWithOffset(2, err).Should(BeNil(), "Cannot get cilium pod on k8s2")

				countBeforeK8s1, _ := helpers.GetBPFPacketsCount(kubectl, ciliumPodK8s1, "Policy denied by denylist", "ingress")
				countBeforeK8s2, _ := helpers.GetBPFPacketsCount(kubectl, ciliumPodK8s2, "Policy denied by denylist", "ingress")

				appPods := helpers.GetAppPods([]string{helpers.App2}, namespaceForTest, kubectl, "id")

				clusterIP, _, err := kubectl.GetServiceHostPort(namespaceForTest, app1Service)
				Expect(err).To(BeNil(), "Cannot get service in %q namespace", namespaceForTest)

				res := kubectl.ExecPodCmd(
					namespaceForTest, appPods[helpers.App2],
					helpers.CurlFail("http://%s/public", clusterIP))
				res.ExpectFail("Unexpected connection from %q to 'http://%s/public'",
					appPods[helpers.App2], clusterIP)

				countAfterK8s1, _ := helpers.GetBPFPacketsCount(kubectl, ciliumPodK8s1, "Policy denied by denylist", "ingress")
				countAfterK8s2, _ := helpers.GetBPFPacketsCount(kubectl, ciliumPodK8s2, "Policy denied by denylist", "ingress")

				Expect((countAfterK8s1 + countAfterK8s2) - (countBeforeK8s1 + countBeforeK8s2)).To(Equal(3))

				_, err = kubectl.CiliumPolicyAction(
					namespaceForTest, l3L4DenyPolicy, helpers.KubectlDelete, helpers.HelperTimeout)
				Expect(err).Should(BeNil(), "Cannot delete L3 Policy")

				kubectl.NamespaceDelete(namespaceForTest)
			})
		})
	})
})
