// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package model

import (
	"sort"
	"strconv"
	"strings"
)

// Model holds an abstracted data model representing the translation
// of various types of Kubernetes config to Cilium config.
type Model struct {
	HTTP []HTTPListener `json:"http,omitempty"`
}

// HTTPListener holds configuration for any listener that terminates and proxies HTTP
// including HTTP and HTTPS.
// Each holds the configuration info for one distinct HTTP listener, by
//   - Hostname
//   - TLS
//   - Address
//   - Port
type HTTPListener struct {
	// Name of the HTTPListener
	Name string `json:"name,omitempty"`
	// Sources is a slice of fully qualified resources this HTTPListener is sourced
	// from.
	Sources []FullyQualifiedResource `json:"sources,omitempty"`
	// IPAddress that the listener should listen on.
	// The string must be parseable as an IP address.
	Address string `json:"address,omitempty"`
	// Port on which the service can be expected to be accessed by clients.
	Port uint32 `json:"port,omitempty"`
	// Hostname that the listener should match.
	// Wildcards are supported in prefix or suffix forms, or the special wildcard `*`.
	// An empty list means that the Listener should match all hostnames.
	Hostname string `json:"hostname,omitempty"`
	// TLS Certificate information. If omitted, then the listener is a cleartext HTTP listener.
	TLS []TLSSecret `json:"tls,omitempty"`
	// Routes associated with HTTP traffic to the service.
	// An empty list means that traffic will not be routed.
	Routes []HTTPRoute `json:"routes,omitempty"`
}

// FullyQualifiedResource stores the full details of a Kubernetes resource, including
// the Group, Version, and Kind.
// Namespace must be set to the empty string for cluster-scoped resources.
type FullyQualifiedResource struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Group     string `json:"group,omitempty"`
	Version   string `json:"version,omitempty"`
	Kind      string `json:"kind,omitempty"`
	UID       string `json:"uuid,omitempty"`
}

// TLSSecret holds a reference to a secret containing a TLS keypair.
type TLSSecret struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// HTTPRoute holds all the details needed to route HTTP traffic to a backend.
type HTTPRoute struct {
	Name string `json:"name,omitempty"`
	// PathMatch specifies that the HTTPRoute should match a path.
	PathMatch StringMatch `json:"path_match,omitempty"`
	// HeadersMatch specifies that the HTTPRoute should match a set of headers.
	HeadersMatch []KeyValueMatch `json:"headers_match,omitempty"`
	// QueryParamsMatch specifies that the HTTPRoute should match a set of query parameters.
	QueryParamsMatch []KeyValueMatch `json:"query_params_match,omitempty"`
	Method           *string         `json:"method,omitempty"`
	Backends         []Backend       `json:"backends,omitempty"`
}

// GetMatchKey returns the key to be used for matching the backend.
func (r *HTTPRoute) GetMatchKey() string {
	sb := strings.Builder{}

	if r.Method != nil {
		sb.WriteString("method:")
		sb.WriteString(*r.Method)
		sb.WriteString("|")
	}

	sb.WriteString("path:")
	sb.WriteString(r.PathMatch.String())
	sb.WriteString("|")

	sort.Slice(r.HeadersMatch, func(i, j int) bool {
		return r.HeadersMatch[i].String() < r.HeadersMatch[j].String()
	})
	for _, hm := range r.HeadersMatch {
		sb.WriteString("header:")
		sb.WriteString(hm.String())
		sb.WriteString("|")
	}

	sort.Slice(r.QueryParamsMatch, func(i, j int) bool {
		return r.QueryParamsMatch[i].String() < r.QueryParamsMatch[j].String()
	})
	for _, qm := range r.QueryParamsMatch {
		sb.WriteString("query:")
		sb.WriteString(qm.String())
		sb.WriteString("|")
	}

	return sb.String()
}

// StringMatch describes various types of string matching.
// Only one field may be set.
// If no fields are set, all paths should match (no path match criteria should
// be generated for Envoy.)
type StringMatch struct {
	Prefix string `json:"prefix,omitempty"`
	Exact  string `json:"exact,omitempty"`
	Regex  string `json:"regex,omitempty"`
}

func (sm StringMatch) String() string {
	sb := strings.Builder{}
	if sm.Prefix != "" {
		sb.WriteString("prefix:")
		sb.WriteString(sm.Prefix)
	} else if sm.Exact != "" {
		sb.WriteString("exact:")
		sb.WriteString(sm.Exact)
	} else if sm.Regex != "" {
		sb.WriteString("regex:")
		sb.WriteString(sm.Regex)
	}
	return sb.String()
}

type KeyValueMatch struct {
	Key   string      `json:"key,omitempty"`
	Match StringMatch `json:"match,omitempty"`
}

func (kv KeyValueMatch) String() string {
	sb := strings.Builder{}
	sb.WriteString("kv:")
	sb.WriteString(kv.Key)
	sb.WriteString(":")
	sb.WriteString(kv.Match.String())
	return sb.String()
}

// Backend holds a Kubernetes Service that points to a backend for traffic.
type Backend struct {
	// Name of the Service.
	Name string `json:"name,omitempty"`
	// Namespace of the Service.
	Namespace string `json:"namespace,omitempty"`
	// Port contains the details of the port on the Service to connect to
	// If unset, the same port as the top-level Listener will be used.
	Port *BackendPort `json:"port,omitempty"`

	// Weight specifies the percentage of traffic to send to this backend.
	// This is computed as weight/(sum of all weights in backends) * 100.
	Weight *int32 `json:"weight,omitempty"`
}

// BackendPort holds the details of what port on the Service to connect to.
// Only one of Port or Name can be set.
type BackendPort struct {
	// Port holds the numeric port to connect to.
	Port uint32 `json:"port,omitempty"`
	// Name holds a string which will be used to connect to the port with a
	// matching spec.ports[].name in the target Service.
	Name string `json:"name,omitempty"`
}

// GetPort return the string representation of the port (either the port number or the port name)
func (be *BackendPort) GetPort() string {
	if be.Port != 0 {
		return strconv.Itoa(int(be.Port))
	}
	return be.Name
}
