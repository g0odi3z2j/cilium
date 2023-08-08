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

package v1

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
var map_Endpoint = map[string]string{
	"":                   "Endpoint represents a single logical \"backend\" implementing a service.",
	"addresses":          "addresses of this endpoint. The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100. These are all assumed to be fungible and clients may choose to only use the first element. Refer to: https://issue.k8s.io/106267",
	"conditions":         "conditions contains information about the current status of the endpoint.",
	"hostname":           "hostname of this endpoint. This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must be lowercase and pass DNS Label (RFC 1123) validation.",
	"targetRef":          "targetRef is a reference to a Kubernetes object that represents this endpoint.",
	"deprecatedTopology": "deprecatedTopology contains topology information part of the v1beta1 API. This field is deprecated, and will be removed when the v1beta1 API is removed (no sooner than kubernetes v1.24).  While this field can hold values, it is not writable through the v1 API, and any attempts to write to it will be silently ignored. Topology information can be found in the zone and nodeName fields instead.",
	"nodeName":           "nodeName represents the name of the Node hosting this endpoint. This can be used to determine endpoints local to a Node.",
	"zone":               "zone is the name of the Zone this endpoint exists in.",
	"hints":              "hints contains information associated with how an endpoint should be consumed.",
}

func (Endpoint) SwaggerDoc() map[string]string {
	return map_Endpoint
}

var map_EndpointConditions = map[string]string{
	"":            "EndpointConditions represents the current condition of an endpoint.",
	"ready":       "ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready. For compatibility reasons, ready should never be \"true\" for terminating endpoints, except when the normal readiness behavior is being explicitly overridden, for example when the associated Service has set the publishNotReadyAddresses flag.",
	"serving":     "serving is identical to ready except that it is set regardless of the terminating state of endpoints. This condition should be set to true for a ready endpoint that is terminating. If nil, consumers should defer to the ready condition.",
	"terminating": "terminating indicates that this endpoint is terminating. A nil value indicates an unknown state. Consumers should interpret this unknown state to mean that the endpoint is not terminating.",
}

func (EndpointConditions) SwaggerDoc() map[string]string {
	return map_EndpointConditions
}

var map_EndpointHints = map[string]string{
	"":         "EndpointHints provides hints describing how an endpoint should be consumed.",
	"forZones": "forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.",
}

func (EndpointHints) SwaggerDoc() map[string]string {
	return map_EndpointHints
}

var map_EndpointPort = map[string]string{
	"":            "EndpointPort represents a Port used by an EndpointSlice",
	"name":        "name represents the name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.",
	"protocol":    "protocol represents the IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
	"port":        "port represents the port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.",
	"appProtocol": "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:\n\n* Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).\n\n* Kubernetes-defined prefixed names:\n  * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540\n  * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455\n  * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455\n\n* Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
}

func (EndpointPort) SwaggerDoc() map[string]string {
	return map_EndpointPort
}

var map_EndpointSlice = map[string]string{
	"":            "EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.",
	"metadata":    "Standard object's metadata.",
	"addressType": "addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.",
	"endpoints":   "endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.",
	"ports":       "ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates \"all ports\". Each slice may include a maximum of 100 ports.",
}

func (EndpointSlice) SwaggerDoc() map[string]string {
	return map_EndpointSlice
}

var map_EndpointSliceList = map[string]string{
	"":         "EndpointSliceList represents a list of endpoint slices",
	"metadata": "Standard list metadata.",
	"items":    "items is the list of endpoint slices",
}

func (EndpointSliceList) SwaggerDoc() map[string]string {
	return map_EndpointSliceList
}

var map_ForZone = map[string]string{
	"":     "ForZone provides information about which zones should consume this endpoint.",
	"name": "name represents the name of the zone.",
}

func (ForZone) SwaggerDoc() map[string]string {
	return map_ForZone
}

// AUTO-GENERATED FUNCTIONS END HERE
