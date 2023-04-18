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

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-codegen.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_ClusterCIDR = map[string]string{
	"":         "ClusterCIDR represents a single configuration for per-Node Pod CIDR allocations when the MultiCIDRRangeAllocator is enabled (see the config for kube-controller-manager).  A cluster may have any number of ClusterCIDR resources, all of which will be considered when allocating a CIDR for a Node.  A ClusterCIDR is eligible to be used for a given Node when the node selector matches the node in question and has free CIDRs to allocate.  In case of multiple matching ClusterCIDR resources, the allocator will attempt to break ties using internal heuristics, but any ClusterCIDR whose node selector matches the Node may be used.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec is the desired state of the ClusterCIDR. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (ClusterCIDR) SwaggerDoc() map[string]string {
	return map_ClusterCIDR
}

var map_ClusterCIDRList = map[string]string{
	"":         "ClusterCIDRList contains a list of ClusterCIDR.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of ClusterCIDRs.",
}

func (ClusterCIDRList) SwaggerDoc() map[string]string {
	return map_ClusterCIDRList
}

var map_ClusterCIDRSpec = map[string]string{
	"":                "ClusterCIDRSpec defines the desired state of ClusterCIDR.",
	"nodeSelector":    "nodeSelector defines which nodes the config is applicable to. An empty or nil nodeSelector selects all nodes. This field is immutable.",
	"perNodeHostBits": "perNodeHostBits defines the number of host bits to be configured per node. A subnet mask determines how much of the address is used for network bits and host bits. For example an IPv4 address of 192.168.0.0/24, splits the address into 24 bits for the network portion and 8 bits for the host portion. To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6). Minimum value is 4 (16 IPs). This field is immutable.",
	"ipv4":            "ipv4 defines an IPv4 IP block in CIDR notation(e.g. \"10.0.0.0/8\"). At least one of ipv4 and ipv6 must be specified. This field is immutable.",
	"ipv6":            "ipv6 defines an IPv6 IP block in CIDR notation(e.g. \"2001:db8::/64\"). At least one of ipv4 and ipv6 must be specified. This field is immutable.",
}

func (ClusterCIDRSpec) SwaggerDoc() map[string]string {
	return map_ClusterCIDRSpec
}

var map_IPAddress = map[string]string{
	"":         "IPAddress represents a single IP of a single IP Family. The object is designed to be used by APIs that operate on IP addresses. The object is used by the Service core API for allocation of IP addresses. An IP address can be represented in different formats, to guarantee the uniqueness of the IP, the name of the object is the IP address in canonical format, four decimal digits separated by dots suppressing leading zeros for IPv4 and the representation defined by RFC 5952 for IPv6. Valid: 192.168.1.5 or 2001:db8::1 or 2001:db8:aaaa:bbbb:cccc:dddd:eeee:1 Invalid: 10.01.2.3 or 2001:db8:0:0:0::1",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec is the desired state of the IPAddress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
}

func (IPAddress) SwaggerDoc() map[string]string {
	return map_IPAddress
}

var map_IPAddressList = map[string]string{
	"":         "IPAddressList contains a list of IPAddress.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of IPAddresses.",
}

func (IPAddressList) SwaggerDoc() map[string]string {
	return map_IPAddressList
}

var map_IPAddressSpec = map[string]string{
	"":          "IPAddressSpec describe the attributes in an IP Address.",
	"parentRef": "ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.",
}

func (IPAddressSpec) SwaggerDoc() map[string]string {
	return map_IPAddressSpec
}

var map_ParentReference = map[string]string{
	"":          "ParentReference describes a reference to a parent object.",
	"group":     "Group is the group of the object being referenced.",
	"resource":  "Resource is the resource of the object being referenced.",
	"namespace": "Namespace is the namespace of the object being referenced.",
	"name":      "Name is the name of the object being referenced.",
	"uid":       "UID is the uid of the object being referenced.",
}

func (ParentReference) SwaggerDoc() map[string]string {
	return map_ParentReference
}

// AUTO-GENERATED FUNCTIONS END HERE
