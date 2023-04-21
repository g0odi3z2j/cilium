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

package v1alpha1

import (
	v1 "k8s.io/api/admissionregistration/v1"
	admissionregistrationv1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1"
)

// RuleWithOperationsApplyConfiguration represents an declarative configuration of the RuleWithOperations type for use
// with apply.
type RuleWithOperationsApplyConfiguration struct {
	Operations                                     []v1.OperationType `json:"operations,omitempty"`
	admissionregistrationv1.RuleApplyConfiguration `json:",inline"`
}

// RuleWithOperationsApplyConfiguration constructs an declarative configuration of the RuleWithOperations type for use with
// apply.
func RuleWithOperations() *RuleWithOperationsApplyConfiguration {
	return &RuleWithOperationsApplyConfiguration{}
}

// WithOperations adds the given value to the Operations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Operations field.
func (b *RuleWithOperationsApplyConfiguration) WithOperations(values ...v1.OperationType) *RuleWithOperationsApplyConfiguration {
	for i := range values {
		b.Operations = append(b.Operations, values[i])
	}
	return b
}

// WithAPIGroups adds the given value to the APIGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIGroups field.
func (b *RuleWithOperationsApplyConfiguration) WithAPIGroups(values ...string) *RuleWithOperationsApplyConfiguration {
	for i := range values {
		b.APIGroups = append(b.APIGroups, values[i])
	}
	return b
}

// WithAPIVersions adds the given value to the APIVersions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIVersions field.
func (b *RuleWithOperationsApplyConfiguration) WithAPIVersions(values ...string) *RuleWithOperationsApplyConfiguration {
	for i := range values {
		b.APIVersions = append(b.APIVersions, values[i])
	}
	return b
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *RuleWithOperationsApplyConfiguration) WithResources(values ...string) *RuleWithOperationsApplyConfiguration {
	for i := range values {
		b.Resources = append(b.Resources, values[i])
	}
	return b
}

// WithScope sets the Scope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scope field is set to the value of the last call.
func (b *RuleWithOperationsApplyConfiguration) WithScope(value v1.ScopeType) *RuleWithOperationsApplyConfiguration {
	b.Scope = &value
	return b
}
