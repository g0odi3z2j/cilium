// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition,
		Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition,
		Convert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition,
		Convert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition,
		Convert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList,
		Convert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList,
		Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames,
		Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames,
		Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec,
		Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec,
		Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus,
		Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus,
	)
}

func autoConvert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(in *CustomResourceDefinition, out *apiextensions.CustomResourceDefinition, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(in *CustomResourceDefinition, out *apiextensions.CustomResourceDefinition, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(in *apiextensions.CustomResourceDefinition, out *CustomResourceDefinition, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(in *apiextensions.CustomResourceDefinition, out *CustomResourceDefinition, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition(in *CustomResourceDefinitionCondition, out *apiextensions.CustomResourceDefinitionCondition, s conversion.Scope) error {
	out.Type = apiextensions.CustomResourceDefinitionConditionType(in.Type)
	out.Status = apiextensions.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition(in *CustomResourceDefinitionCondition, out *apiextensions.CustomResourceDefinitionCondition, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition(in *apiextensions.CustomResourceDefinitionCondition, out *CustomResourceDefinitionCondition, s conversion.Scope) error {
	out.Type = CustomResourceDefinitionConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition(in *apiextensions.CustomResourceDefinitionCondition, out *CustomResourceDefinitionCondition, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList(in *CustomResourceDefinitionList, out *apiextensions.CustomResourceDefinitionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]apiextensions.CustomResourceDefinition)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList(in *CustomResourceDefinitionList, out *apiextensions.CustomResourceDefinitionList, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList(in *apiextensions.CustomResourceDefinitionList, out *CustomResourceDefinitionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]CustomResourceDefinition, 0)
	} else {
		out.Items = *(*[]CustomResourceDefinition)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList(in *apiextensions.CustomResourceDefinitionList, out *CustomResourceDefinitionList, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(in *CustomResourceDefinitionNames, out *apiextensions.CustomResourceDefinitionNames, s conversion.Scope) error {
	out.Plural = in.Plural
	out.Singular = in.Singular
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	out.Kind = in.Kind
	out.ListKind = in.ListKind
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(in *CustomResourceDefinitionNames, out *apiextensions.CustomResourceDefinitionNames, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(in *apiextensions.CustomResourceDefinitionNames, out *CustomResourceDefinitionNames, s conversion.Scope) error {
	out.Plural = in.Plural
	out.Singular = in.Singular
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	out.Kind = in.Kind
	out.ListKind = in.ListKind
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(in *apiextensions.CustomResourceDefinitionNames, out *CustomResourceDefinitionNames, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(in *CustomResourceDefinitionSpec, out *apiextensions.CustomResourceDefinitionSpec, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	if err := Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(&in.Names, &out.Names, s); err != nil {
		return err
	}
	out.Scope = apiextensions.ResourceScope(in.Scope)
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(in *CustomResourceDefinitionSpec, out *apiextensions.CustomResourceDefinitionSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(in *apiextensions.CustomResourceDefinitionSpec, out *CustomResourceDefinitionSpec, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	if err := Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(&in.Names, &out.Names, s); err != nil {
		return err
	}
	out.Scope = ResourceScope(in.Scope)
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(in *apiextensions.CustomResourceDefinitionSpec, out *CustomResourceDefinitionSpec, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(in *CustomResourceDefinitionStatus, out *apiextensions.CustomResourceDefinitionStatus, s conversion.Scope) error {
	out.Conditions = *(*[]apiextensions.CustomResourceDefinitionCondition)(unsafe.Pointer(&in.Conditions))
	if err := Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(&in.AcceptedNames, &out.AcceptedNames, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(in *CustomResourceDefinitionStatus, out *apiextensions.CustomResourceDefinitionStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(in *apiextensions.CustomResourceDefinitionStatus, out *CustomResourceDefinitionStatus, s conversion.Scope) error {
	if in.Conditions == nil {
		out.Conditions = make([]CustomResourceDefinitionCondition, 0)
	} else {
		out.Conditions = *(*[]CustomResourceDefinitionCondition)(unsafe.Pointer(&in.Conditions))
	}
	if err := Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(&in.AcceptedNames, &out.AcceptedNames, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(in *apiextensions.CustomResourceDefinitionStatus, out *CustomResourceDefinitionStatus, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(in, out, s)
}
