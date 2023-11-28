/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ClusterTrustBundleProjectionApplyConfiguration represents an declarative configuration of the ClusterTrustBundleProjection type for use
// with apply.
type ClusterTrustBundleProjectionApplyConfiguration struct {
	Name          *string                             `json:"name,omitempty"`
	SignerName    *string                             `json:"signerName,omitempty"`
	LabelSelector *v1.LabelSelectorApplyConfiguration `json:"labelSelector,omitempty"`
	Optional      *bool                               `json:"optional,omitempty"`
	Path          *string                             `json:"path,omitempty"`
}

// ClusterTrustBundleProjectionApplyConfiguration constructs an declarative configuration of the ClusterTrustBundleProjection type for use with
// apply.
func ClusterTrustBundleProjection() *ClusterTrustBundleProjectionApplyConfiguration {
	return &ClusterTrustBundleProjectionApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ClusterTrustBundleProjectionApplyConfiguration) WithName(value string) *ClusterTrustBundleProjectionApplyConfiguration {
	b.Name = &value
	return b
}

// WithSignerName sets the SignerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignerName field is set to the value of the last call.
func (b *ClusterTrustBundleProjectionApplyConfiguration) WithSignerName(value string) *ClusterTrustBundleProjectionApplyConfiguration {
	b.SignerName = &value
	return b
}

// WithLabelSelector sets the LabelSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelSelector field is set to the value of the last call.
func (b *ClusterTrustBundleProjectionApplyConfiguration) WithLabelSelector(value *v1.LabelSelectorApplyConfiguration) *ClusterTrustBundleProjectionApplyConfiguration {
	b.LabelSelector = value
	return b
}

// WithOptional sets the Optional field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Optional field is set to the value of the last call.
func (b *ClusterTrustBundleProjectionApplyConfiguration) WithOptional(value bool) *ClusterTrustBundleProjectionApplyConfiguration {
	b.Optional = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *ClusterTrustBundleProjectionApplyConfiguration) WithPath(value string) *ClusterTrustBundleProjectionApplyConfiguration {
	b.Path = &value
	return b
}
