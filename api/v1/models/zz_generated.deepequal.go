//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepequal-gen. DO NOT EDIT.

package models

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *ControllerStatusConfiguration) DeepEqual(other *ControllerStatusConfiguration) bool {
	if other == nil {
		return false
	}

	if in.ErrorRetry != other.ErrorRetry {
		return false
	}
	if in.ErrorRetryBase != other.ErrorRetryBase {
		return false
	}
	if in.Interval != other.Interval {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointHealth) DeepEqual(other *EndpointHealth) bool {
	if other == nil {
		return false
	}

	if in.Bpf != other.Bpf {
		return false
	}
	if in.Connected != other.Connected {
		return false
	}
	if in.OverallHealth != other.OverallHealth {
		return false
	}
	if in.Policy != other.Policy {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointIdentifiers) DeepEqual(other *EndpointIdentifiers) bool {
	if other == nil {
		return false
	}

	if in.ContainerID != other.ContainerID {
		return false
	}
	if in.ContainerName != other.ContainerName {
		return false
	}
	if in.DockerEndpointID != other.DockerEndpointID {
		return false
	}
	if in.DockerNetworkID != other.DockerNetworkID {
		return false
	}
	if in.K8sNamespace != other.K8sNamespace {
		return false
	}
	if in.K8sPodName != other.K8sPodName {
		return false
	}
	if in.PodName != other.PodName {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EndpointStatusChange) DeepEqual(other *EndpointStatusChange) bool {
	if other == nil {
		return false
	}

	if in.Code != other.Code {
		return false
	}
	if in.Message != other.Message {
		return false
	}
	if in.State != other.State {
		return false
	}
	if in.Timestamp != other.Timestamp {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *NamedPorts) DeepEqual(other *NamedPorts) bool {
	if other == nil {
		return false
	}

	if len(*in) != len(*other) {
		return false
	} else {
		for i, inElement := range *in {
			if !inElement.DeepEqual((*other)[i]) {
				return false
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *Port) DeepEqual(other *Port) bool {
	if other == nil {
		return false
	}

	if in.Name != other.Name {
		return false
	}
	if in.Port != other.Port {
		return false
	}
	if in.Protocol != other.Protocol {
		return false
	}

	return true
}
