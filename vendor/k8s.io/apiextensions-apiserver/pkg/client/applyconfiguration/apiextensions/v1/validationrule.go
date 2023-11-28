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
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// ValidationRuleApplyConfiguration represents an declarative configuration of the ValidationRule type for use
// with apply.
type ValidationRuleApplyConfiguration struct {
	Rule              *string                   `json:"rule,omitempty"`
	Message           *string                   `json:"message,omitempty"`
	MessageExpression *string                   `json:"messageExpression,omitempty"`
	Reason            *v1.FieldValueErrorReason `json:"reason,omitempty"`
	FieldPath         *string                   `json:"fieldPath,omitempty"`
	OptionalOldSelf   *bool                     `json:"optionalOldSelf,omitempty"`
}

// ValidationRuleApplyConfiguration constructs an declarative configuration of the ValidationRule type for use with
// apply.
func ValidationRule() *ValidationRuleApplyConfiguration {
	return &ValidationRuleApplyConfiguration{}
}

// WithRule sets the Rule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Rule field is set to the value of the last call.
func (b *ValidationRuleApplyConfiguration) WithRule(value string) *ValidationRuleApplyConfiguration {
	b.Rule = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ValidationRuleApplyConfiguration) WithMessage(value string) *ValidationRuleApplyConfiguration {
	b.Message = &value
	return b
}

// WithMessageExpression sets the MessageExpression field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MessageExpression field is set to the value of the last call.
func (b *ValidationRuleApplyConfiguration) WithMessageExpression(value string) *ValidationRuleApplyConfiguration {
	b.MessageExpression = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ValidationRuleApplyConfiguration) WithReason(value v1.FieldValueErrorReason) *ValidationRuleApplyConfiguration {
	b.Reason = &value
	return b
}

// WithFieldPath sets the FieldPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FieldPath field is set to the value of the last call.
func (b *ValidationRuleApplyConfiguration) WithFieldPath(value string) *ValidationRuleApplyConfiguration {
	b.FieldPath = &value
	return b
}

// WithOptionalOldSelf sets the OptionalOldSelf field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OptionalOldSelf field is set to the value of the last call.
func (b *ValidationRuleApplyConfiguration) WithOptionalOldSelf(value bool) *ValidationRuleApplyConfiguration {
	b.OptionalOldSelf = &value
	return b
}
