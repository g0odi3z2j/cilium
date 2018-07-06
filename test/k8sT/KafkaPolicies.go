// Copyright 2018 Authors of Cilium
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

	"github.com/cilium/cilium/api/v1/models"
	. "github.com/cilium/cilium/test/ginkgo-ext"
	"github.com/cilium/cilium/test/helpers"

	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("K8sValidatedKafkaPolicyTest", func() {

	var (
		kubectl             *helpers.Kubectl
		microscopeErr       error
		microscopeCancel    = func() error { return nil }
		logger              = log.WithFields(logrus.Fields{"testName": "K8sValidatedKafkaPolicyTest"})
		l7Policy            = helpers.ManifestGet("kafka-sw-security-policy.yaml")
		demoPath            = helpers.ManifestGet("kafka-sw-app.yaml")
		kafkaApp            = "kafka"
		zookApp             = "zook"
		backupApp           = "empire-backup"
		empireHqApp         = "empire-hq"
		outpostApp          = "empire-outpost"
		apps                = []string{kafkaApp, zookApp, backupApp, empireHqApp, outpostApp}
		appPods             = map[string]string{}
		topicEmpireAnnounce = "empire-announce"
		topicDeathstarPlans = "deathstar-plans"
		topicTest           = "test-topic"

		prodHqAnnounce    = `-c "echo 'Happy 40th Birthday to General Tagge' | ./kafka-produce.sh --topic empire-announce"`
		conOutpostAnnoune = `-c "./kafka-consume.sh --topic empire-announce --from-beginning --max-messages 1"`
		prodHqDeathStar   = `-c "echo 'deathstar reactor design v3' | ./kafka-produce.sh --topic deathstar-plans"`
		conOutDeathStar   = `-c "./kafka-consume.sh --topic deathstar-plans --from-beginning --max-messages 1"`
		prodBackAnnounce  = `-c "echo 'Happy 40th Birthday to General Tagge' | ./kafka-produce.sh --topic empire-announce"`
		prodOutAnnounce   = `-c "echo 'Vader Booed at Empire Karaoke Party' | ./kafka-produce.sh --topic empire-announce"`
	)

	AfterFailed(func() {
		kubectl.CiliumReport(helpers.KubeSystemNamespace,
			"cilium service list",
			"cilium endpoint list")
	})

	Context("Kafka Policy Tests", func() {
		createTopicCmd := func(topic string) string {
			return fmt.Sprintf("/opt/kafka/bin/kafka-topics.sh --create --zookeeper zook:2181 "+
				"--replication-factor 1 --partitions 1 --topic %s", topic)
		}

		createTopic := func(topic string, pod string) error {
			return kubectl.ExecKafkaPodCmd(helpers.DefaultNamespace, pod, createTopicCmd(topic))
		}

		// WaitKafkaBroker waits for the broker to be ready, by executing
		// a command repeatedly until it succeeds, or a timeout occurs
		waitForKafkaBroker := func(pod string, cmd string) error {
			body := func() bool {
				err := kubectl.ExecKafkaPodCmd(helpers.DefaultNamespace, pod, cmd)
				if err != nil {
					return false
				}
				return true
			}
			err := helpers.WithTimeout(body, "Kafka Broker not ready", &helpers.TimeoutConfig{Timeout: helpers.HelperTimeout})
			return err
		}

		JustBeforeEach(func() {
			microscopeErr, microscopeCancel = kubectl.MicroscopeStart()
			Expect(microscopeErr).To(BeNil(), "Microscope cannot be started")
		})

		JustAfterEach(func() {
			kubectl.ValidateNoErrorsOnLogs(CurrentGinkgoTestDescription().Duration)
			Expect(microscopeCancel()).To(BeNil(), "cannot stop microscope")
		})

		AfterEach(func() {
			// On aftereach don't make assertions to delete all.
			_ = kubectl.Delete(demoPath)
			_ = kubectl.Delete(l7Policy)

			ExpectAllPodsTerminated(kubectl)
		})

		BeforeAll(func() {
			logger = log.WithFields(logrus.Fields{"testName": "K8sValidatedKafkaPolicyTest"})
			kubectl = helpers.CreateKubectl(helpers.K8s1VMName(), logger)

			err := kubectl.CiliumInstall(helpers.CiliumDSPath)
			Expect(err).To(BeNil(), "Cilium cannot be installed")
			ExpectCiliumReady(kubectl)
			ExpectKubeDNSReady(kubectl)

			kubectl.Apply(demoPath)
			err = kubectl.WaitforPods(helpers.DefaultNamespace, "-l zgroup=kafkaTestApp", helpers.HelperTimeout)
			Expect(err).Should(BeNil(), "Kafka Pods are not ready after timeout")

			err = kubectl.WaitForKubeDNSEntry("kafka-service." + helpers.DefaultNamespace)
			Expect(err).To(BeNil(), "DNS entry of kafka-service is not ready after timeout")
			err = kubectl.WaitForKubeDNSEntry("zook." + helpers.DefaultNamespace)
			Expect(err).To(BeNil(), "DNS entry of zook is not ready after timeout")

			appPods = helpers.GetAppPods(apps, helpers.DefaultNamespace, kubectl, "app")

			By("Wait for Kafka broker to be up")
			err = waitForKafkaBroker(appPods[empireHqApp], createTopicCmd(topicTest))
			Expect(err).To(BeNil(), "Timeout: Kafka cluster failed to come up correctly")
		})

		It("KafkaPolicies", func() {
			By("Creating new kafka topic %s", topicEmpireAnnounce)
			err := createTopic(topicEmpireAnnounce, appPods[empireHqApp])
			Expect(err).Should(BeNil(), "Failed to create topic empire-announce")

			By("Creating new kafka topic %s", topicDeathstarPlans)
			err = createTopic(topicDeathstarPlans, appPods[empireHqApp])
			Expect(err).Should(BeNil(), "Failed to create topic deathstar-plans")

			By("Testing basic Kafka Produce and Consume")
			// We need to produce first, since consumer script waits for
			// some messages to be already there by the producer.

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[empireHqApp], fmt.Sprintf(prodHqAnnounce))
			Expect(err).Should(BeNil(), "Failed to produce to empire-hq on topic empire-announce")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[outpostApp], fmt.Sprintf(conOutpostAnnoune))
			Expect(err).Should(BeNil(), "Failed to consume from outpost on topic empire-announce")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[empireHqApp], fmt.Sprintf(prodHqDeathStar))
			Expect(err).Should(BeNil(), "Failed to produce to empire-hq on topic deathstar-plans")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[outpostApp], fmt.Sprintf(conOutDeathStar))
			Expect(err).Should(BeNil(), "Failed to consume from outpost on topic deathstar-plans")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[backupApp], fmt.Sprintf(prodBackAnnounce))
			Expect(err).Should(BeNil(), "Failed to produce to backup on topic empire-announce")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[outpostApp], fmt.Sprintf(prodOutAnnounce))
			Expect(err).Should(BeNil(), "Failed to produce to outpost on topic empire-announce")

			By("Apply L7 kafka policy and wait")

			_, err = kubectl.CiliumPolicyAction(
				helpers.KubeSystemNamespace, l7Policy,
				helpers.KubectlApply, helpers.HelperTimeout)
			Expect(err).To(BeNil(), "L7 policy cannot be imported correctly")

			ExpectCEPUpdates(kubectl)

			By("validate that the pods have the correct policy")

			desiredPolicyStatus := map[string]models.EndpointPolicyEnabled{
				backupApp:   models.EndpointPolicyEnabledNone,
				empireHqApp: models.EndpointPolicyEnabledNone,
				kafkaApp:    models.EndpointPolicyEnabledIngress,
				outpostApp:  models.EndpointPolicyEnabledNone,
				zookApp:     models.EndpointPolicyEnabledNone,
			}

			for app, policy := range desiredPolicyStatus {
				cep := kubectl.CepGet(helpers.DefaultNamespace, appPods[app])
				Expect(cep).ToNot(BeNil(), "cannot get cep for app %q and pod %s", app, appPods[app])
				Expect(cep.Status.Policy.Spec.PolicyEnabled).To(Equal(policy), "Policy for %q mismatch", app)
			}

			By("Testing Kafka L7 policy enforcement status")
			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[empireHqApp], fmt.Sprintf(prodHqAnnounce))
			Expect(err).Should(BeNil(), "Failed to produce to empire-hq on topic empire-announce")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[outpostApp], fmt.Sprintf(conOutpostAnnoune))
			Expect(err).Should(BeNil(), "Failed to consume from outpost on topic empire-announce")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[empireHqApp], fmt.Sprintf(prodHqDeathStar))
			Expect(err).Should(BeNil(), "Failed to produce from empire-hq on topic deathstar-plans")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[outpostApp], fmt.Sprintf(conOutpostAnnoune))
			Expect(err).Should(BeNil(), "Failed to consume from outpost on topic empire-announce")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[backupApp], fmt.Sprintf(prodBackAnnounce))
			Expect(err).Should(HaveOccurred(), " Produce to backup on topic empire-announce should have been denied")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[outpostApp], fmt.Sprintf(conOutDeathStar))
			Expect(err).Should(HaveOccurred(), " Consume from outpost on topic deathstar-plans should have been denied")

			err = kubectl.ExecKafkaPodCmd(
				helpers.DefaultNamespace, appPods[outpostApp], fmt.Sprintf(prodOutAnnounce))
			Expect(err).Should(HaveOccurred(), "Produce to outpost on topic empire-announce should have been denied")
		})
	})
})
