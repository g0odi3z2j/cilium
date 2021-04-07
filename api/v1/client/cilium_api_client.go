// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/client/daemon"
	"github.com/cilium/cilium/api/v1/client/endpoint"
	"github.com/cilium/cilium/api/v1/client/ipam"
	"github.com/cilium/cilium/api/v1/client/metrics"
	"github.com/cilium/cilium/api/v1/client/policy"
	"github.com/cilium/cilium/api/v1/client/prefilter"
	"github.com/cilium/cilium/api/v1/client/recorder"
	"github.com/cilium/cilium/api/v1/client/service"
)

// Default cilium API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new cilium API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CiliumAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cilium API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *CiliumAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cilium API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CiliumAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(CiliumAPI)
	cli.Transport = transport
	cli.Daemon = daemon.New(transport, formats)
	cli.Endpoint = endpoint.New(transport, formats)
	cli.Ipam = ipam.New(transport, formats)
	cli.Metrics = metrics.New(transport, formats)
	cli.Policy = policy.New(transport, formats)
	cli.Prefilter = prefilter.New(transport, formats)
	cli.Recorder = recorder.New(transport, formats)
	cli.Service = service.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// CiliumAPI is a client for cilium API
type CiliumAPI struct {
	Daemon daemon.ClientService

	Endpoint endpoint.ClientService

	Ipam ipam.ClientService

	Metrics metrics.ClientService

	Policy policy.ClientService

	Prefilter prefilter.ClientService

	Recorder recorder.ClientService

	Service service.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CiliumAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Daemon.SetTransport(transport)
	c.Endpoint.SetTransport(transport)
	c.Ipam.SetTransport(transport)
	c.Metrics.SetTransport(transport)
	c.Policy.SetTransport(transport)
	c.Prefilter.SetTransport(transport)
	c.Recorder.SetTransport(transport)
	c.Service.SetTransport(transport)
}
