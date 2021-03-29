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

package v1beta2

import (
	v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeploymentConditionApplyConfiguration represents an declarative configuration of the DeploymentCondition type for use
// with apply.
type DeploymentConditionApplyConfiguration struct {
	Type               *v1beta2.DeploymentConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus              `json:"status,omitempty"`
	LastUpdateTime     *metav1.Time                     `json:"lastUpdateTime,omitempty"`
	LastTransitionTime *metav1.Time                     `json:"lastTransitionTime,omitempty"`
	Reason             *string                          `json:"reason,omitempty"`
	Message            *string                          `json:"message,omitempty"`
}

// DeploymentConditionApplyConfiguration constructs an declarative configuration of the DeploymentCondition type for use with
// apply.
func DeploymentCondition() *DeploymentConditionApplyConfiguration {
	return &DeploymentConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *DeploymentConditionApplyConfiguration) WithType(value v1beta2.DeploymentConditionType) *DeploymentConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *DeploymentConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *DeploymentConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastUpdateTime sets the LastUpdateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastUpdateTime field is set to the value of the last call.
func (b *DeploymentConditionApplyConfiguration) WithLastUpdateTime(value metav1.Time) *DeploymentConditionApplyConfiguration {
	b.LastUpdateTime = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *DeploymentConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *DeploymentConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *DeploymentConditionApplyConfiguration) WithReason(value string) *DeploymentConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *DeploymentConditionApplyConfiguration) WithMessage(value string) *DeploymentConditionApplyConfiguration {
	b.Message = &value
	return b
}
