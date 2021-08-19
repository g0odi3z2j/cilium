//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by main. DO NOT EDIT.

package labels

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Label) DeepEqual(other *Label) bool {
	if other == nil {
		return false
	}

	if in.Key != other.Key {
		return false
	}
	if in.Value != other.Value {
		return false
	}
	if in.Source != other.Source {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *LabelArray) DeepEqual(other *LabelArray) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for i, inElement := range *in {
			if !inElement.DeepEqual(&(*other)[i]) {
				return false
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *LabelArrayList) DeepEqual(other *LabelArrayList) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for i, inElement := range *in {
			if !inElement.DeepEqual(&(*other)[i]) {
				return false
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Labels) DeepEqual(other *Labels) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for key, inValue := range *in {
			if otherValue, present := (*other)[key]; !present {
				return false
			} else {
				if !inValue.DeepEqual(&otherValue) {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *OpLabels) DeepEqual(other *OpLabels) bool {
	if other == nil {
		return false
	}

	if ((in.Custom != nil) && (other.Custom != nil)) || ((in.Custom == nil) != (other.Custom == nil)) {
		in, other := &in.Custom, &other.Custom
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.OrchestrationIdentity != nil) && (other.OrchestrationIdentity != nil)) || ((in.OrchestrationIdentity == nil) != (other.OrchestrationIdentity == nil)) {
		in, other := &in.OrchestrationIdentity, &other.OrchestrationIdentity
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.Disabled != nil) && (other.Disabled != nil)) || ((in.Disabled == nil) != (other.Disabled == nil)) {
		in, other := &in.Disabled, &other.Disabled
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	if ((in.OrchestrationInfo != nil) && (other.OrchestrationInfo != nil)) || ((in.OrchestrationInfo == nil) != (other.OrchestrationInfo == nil)) {
		in, other := &in.OrchestrationInfo, &other.OrchestrationInfo
		if other == nil || !in.DeepEqual(other) {
			return false
		}
	}

	return true
}
