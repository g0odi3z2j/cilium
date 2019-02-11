// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/client/daemon"
	"github.com/cilium/cilium/api/v1/client/endpoint"
	"github.com/cilium/cilium/api/v1/client/ipam"
	"github.com/cilium/cilium/api/v1/client/metrics"
	"github.com/cilium/cilium/api/v1/client/policy"
	"github.com/cilium/cilium/api/v1/client/prefilter"
	"github.com/cilium/cilium/api/v1/client/service"
)

// Default cilium HTTP client.
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

// NewHTTPClient creates a new cilium HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cilium {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cilium HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Cilium {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cilium client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Cilium {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Cilium)
	cli.Transport = transport

	cli.Daemon = daemon.New(transport, formats)

	cli.Endpoint = endpoint.New(transport, formats)

	cli.IPAM = ipam.New(transport, formats)

	cli.Metrics = metrics.New(transport, formats)

	cli.Policy = policy.New(transport, formats)

	cli.Prefilter = prefilter.New(transport, formats)

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

// Cilium is a client for cilium
type Cilium struct {
	Daemon *daemon.Client

	Endpoint *endpoint.Client

	IPAM *ipam.Client

	Metrics *metrics.Client

	Policy *policy.Client

	Prefilter *prefilter.Client

	Service *service.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cilium) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Daemon.SetTransport(transport)

	c.Endpoint.SetTransport(transport)

	c.IPAM.SetTransport(transport)

	c.Metrics.SetTransport(transport)

	c.Policy.SetTransport(transport)

	c.Prefilter.SetTransport(transport)

	c.Service.SetTransport(transport)

}
