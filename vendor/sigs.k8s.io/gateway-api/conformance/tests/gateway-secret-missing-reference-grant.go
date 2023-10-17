/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	v1 "sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/gateway-api/conformance/utils/kubernetes"
	"sigs.k8s.io/gateway-api/conformance/utils/suite"
)

func init() {
	ConformanceTests = append(ConformanceTests, GatewaySecretMissingReferenceGrant)
}

var GatewaySecretMissingReferenceGrant = suite.ConformanceTest{
	ShortName:   "GatewaySecretMissingReferenceGrant",
	Description: "A Gateway in the gateway-conformance-infra namespace should fail to become programmed if the Gateway has a certificateRef for a Secret in the gateway-conformance-web-backend namespace and a ReferenceGrant granting permission to the Secret does not exist",
	Features: []suite.SupportedFeature{
		suite.SupportGateway,
		suite.SupportReferenceGrant,
	},
	Manifests: []string{"tests/gateway-secret-missing-reference-grant.yaml"},
	Test: func(t *testing.T, s *suite.ConformanceTestSuite) {
		gwNN := types.NamespacedName{Name: "gateway-secret-missing-reference-grant", Namespace: "gateway-conformance-infra"}

		t.Run("Gateway listener should have a false ResolvedRefs condition with reason RefNotPermitted", func(t *testing.T) {
			listeners := []v1.ListenerStatus{{
				Name: v1.SectionName("https"),
				SupportedKinds: []v1.RouteGroupKind{{
					Group: (*v1.Group)(&v1.GroupVersion.Group),
					Kind:  v1.Kind("HTTPRoute"),
				}},
				Conditions: []metav1.Condition{{
					Type:   string(v1.ListenerConditionResolvedRefs),
					Status: metav1.ConditionFalse,
					Reason: string(v1.ListenerReasonRefNotPermitted),
				}},
				AttachedRoutes: 0,
			}}

			kubernetes.GatewayStatusMustHaveListeners(t, s.Client, s.TimeoutConfig, gwNN, listeners)
		})
	},
}
