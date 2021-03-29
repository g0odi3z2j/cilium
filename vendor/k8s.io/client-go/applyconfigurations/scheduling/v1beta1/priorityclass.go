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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/scheduling/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "k8s.io/client-go/applyconfigurations/internal"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// PriorityClassApplyConfiguration represents an declarative configuration of the PriorityClass type for use
// with apply.
type PriorityClassApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Value                            *int32                   `json:"value,omitempty"`
	GlobalDefault                    *bool                    `json:"globalDefault,omitempty"`
	Description                      *string                  `json:"description,omitempty"`
	PreemptionPolicy                 *corev1.PreemptionPolicy `json:"preemptionPolicy,omitempty"`
}

// PriorityClass constructs an declarative configuration of the PriorityClass type for use with
// apply.
func PriorityClass(name string) *PriorityClassApplyConfiguration {
	b := &PriorityClassApplyConfiguration{}
	b.WithName(name)
	b.WithKind("PriorityClass")
	b.WithAPIVersion("scheduling.k8s.io/v1beta1")
	return b
}

// ExtractPriorityClass extracts the applied configuration owned by fieldManager from
// priorityClass. If no managedFields are found in priorityClass for fieldManager, a
// PriorityClassApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. Is is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// priorityClass must be a unmodified PriorityClass API object that was retrieved from the Kubernetes API.
// ExtractPriorityClass provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractPriorityClass(priorityClass *v1beta1.PriorityClass, fieldManager string) (*PriorityClassApplyConfiguration, error) {
	b := &PriorityClassApplyConfiguration{}
	err := managedfields.ExtractInto(priorityClass, internal.Parser().Type("io.k8s.api.scheduling.v1beta1.PriorityClass"), fieldManager, b)
	if err != nil {
		return nil, err
	}
	b.WithName(priorityClass.Name)

	b.WithKind("PriorityClass")
	b.WithAPIVersion("scheduling.k8s.io/v1beta1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithKind(value string) *PriorityClassApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithAPIVersion(value string) *PriorityClassApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithName(value string) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithGenerateName(value string) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithNamespace(value string) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithSelfLink sets the SelfLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelfLink field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithSelfLink(value string) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.SelfLink = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithUID(value types.UID) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithResourceVersion(value string) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithGeneration(value int64) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithCreationTimestamp(value metav1.Time) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *PriorityClassApplyConfiguration) WithLabels(entries map[string]string) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *PriorityClassApplyConfiguration) WithAnnotations(entries map[string]string) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *PriorityClassApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *PriorityClassApplyConfiguration) WithFinalizers(values ...string) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithClusterName(value string) *PriorityClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ClusterName = &value
	return b
}

func (b *PriorityClassApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithValue(value int32) *PriorityClassApplyConfiguration {
	b.Value = &value
	return b
}

// WithGlobalDefault sets the GlobalDefault field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GlobalDefault field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithGlobalDefault(value bool) *PriorityClassApplyConfiguration {
	b.GlobalDefault = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithDescription(value string) *PriorityClassApplyConfiguration {
	b.Description = &value
	return b
}

// WithPreemptionPolicy sets the PreemptionPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreemptionPolicy field is set to the value of the last call.
func (b *PriorityClassApplyConfiguration) WithPreemptionPolicy(value corev1.PreemptionPolicy) *PriorityClassApplyConfiguration {
	b.PreemptionPolicy = &value
	return b
}
