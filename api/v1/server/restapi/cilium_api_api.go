// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cilium/cilium/api/v1/server/restapi/daemon"
	"github.com/cilium/cilium/api/v1/server/restapi/endpoint"
	"github.com/cilium/cilium/api/v1/server/restapi/ipam"
	"github.com/cilium/cilium/api/v1/server/restapi/metrics"
	"github.com/cilium/cilium/api/v1/server/restapi/policy"
	"github.com/cilium/cilium/api/v1/server/restapi/prefilter"
	"github.com/cilium/cilium/api/v1/server/restapi/service"
)

// NewCiliumAPIAPI creates a new CiliumAPI instance
func NewCiliumAPIAPI(spec *loads.Document) *CiliumAPIAPI {
	return &CiliumAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		EndpointDeleteEndpointIDHandler: endpoint.DeleteEndpointIDHandlerFunc(func(params endpoint.DeleteEndpointIDParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.DeleteEndpointID has not yet been implemented")
		}),
		PolicyDeleteFqdnCacheHandler: policy.DeleteFqdnCacheHandlerFunc(func(params policy.DeleteFqdnCacheParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.DeleteFqdnCache has not yet been implemented")
		}),
		IpamDeleteIpamIPHandler: ipam.DeleteIpamIPHandlerFunc(func(params ipam.DeleteIpamIPParams) middleware.Responder {
			return middleware.NotImplemented("operation ipam.DeleteIpamIP has not yet been implemented")
		}),
		PolicyDeletePolicyHandler: policy.DeletePolicyHandlerFunc(func(params policy.DeletePolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.DeletePolicy has not yet been implemented")
		}),
		PrefilterDeletePrefilterHandler: prefilter.DeletePrefilterHandlerFunc(func(params prefilter.DeletePrefilterParams) middleware.Responder {
			return middleware.NotImplemented("operation prefilter.DeletePrefilter has not yet been implemented")
		}),
		ServiceDeleteServiceIDHandler: service.DeleteServiceIDHandlerFunc(func(params service.DeleteServiceIDParams) middleware.Responder {
			return middleware.NotImplemented("operation service.DeleteServiceID has not yet been implemented")
		}),
		DaemonGetClusterNodesHandler: daemon.GetClusterNodesHandlerFunc(func(params daemon.GetClusterNodesParams) middleware.Responder {
			return middleware.NotImplemented("operation daemon.GetClusterNodes has not yet been implemented")
		}),
		DaemonGetConfigHandler: daemon.GetConfigHandlerFunc(func(params daemon.GetConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation daemon.GetConfig has not yet been implemented")
		}),
		DaemonGetDebuginfoHandler: daemon.GetDebuginfoHandlerFunc(func(params daemon.GetDebuginfoParams) middleware.Responder {
			return middleware.NotImplemented("operation daemon.GetDebuginfo has not yet been implemented")
		}),
		EndpointGetEndpointHandler: endpoint.GetEndpointHandlerFunc(func(params endpoint.GetEndpointParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.GetEndpoint has not yet been implemented")
		}),
		EndpointGetEndpointIDHandler: endpoint.GetEndpointIDHandlerFunc(func(params endpoint.GetEndpointIDParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.GetEndpointID has not yet been implemented")
		}),
		EndpointGetEndpointIDConfigHandler: endpoint.GetEndpointIDConfigHandlerFunc(func(params endpoint.GetEndpointIDConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.GetEndpointIDConfig has not yet been implemented")
		}),
		EndpointGetEndpointIDHealthzHandler: endpoint.GetEndpointIDHealthzHandlerFunc(func(params endpoint.GetEndpointIDHealthzParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.GetEndpointIDHealthz has not yet been implemented")
		}),
		EndpointGetEndpointIDLabelsHandler: endpoint.GetEndpointIDLabelsHandlerFunc(func(params endpoint.GetEndpointIDLabelsParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.GetEndpointIDLabels has not yet been implemented")
		}),
		EndpointGetEndpointIDLogHandler: endpoint.GetEndpointIDLogHandlerFunc(func(params endpoint.GetEndpointIDLogParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.GetEndpointIDLog has not yet been implemented")
		}),
		PolicyGetFqdnCacheHandler: policy.GetFqdnCacheHandlerFunc(func(params policy.GetFqdnCacheParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetFqdnCache has not yet been implemented")
		}),
		PolicyGetFqdnCacheIDHandler: policy.GetFqdnCacheIDHandlerFunc(func(params policy.GetFqdnCacheIDParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetFqdnCacheID has not yet been implemented")
		}),
		PolicyGetFqdnNamesHandler: policy.GetFqdnNamesHandlerFunc(func(params policy.GetFqdnNamesParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetFqdnNames has not yet been implemented")
		}),
		DaemonGetHealthzHandler: daemon.GetHealthzHandlerFunc(func(params daemon.GetHealthzParams) middleware.Responder {
			return middleware.NotImplemented("operation daemon.GetHealthz has not yet been implemented")
		}),
		PolicyGetIPHandler: policy.GetIPHandlerFunc(func(params policy.GetIPParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetIP has not yet been implemented")
		}),
		PolicyGetIdentityHandler: policy.GetIdentityHandlerFunc(func(params policy.GetIdentityParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetIdentity has not yet been implemented")
		}),
		PolicyGetIdentityEndpointsHandler: policy.GetIdentityEndpointsHandlerFunc(func(params policy.GetIdentityEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetIdentityEndpoints has not yet been implemented")
		}),
		PolicyGetIdentityIDHandler: policy.GetIdentityIDHandlerFunc(func(params policy.GetIdentityIDParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetIdentityID has not yet been implemented")
		}),
		DaemonGetMapHandler: daemon.GetMapHandlerFunc(func(params daemon.GetMapParams) middleware.Responder {
			return middleware.NotImplemented("operation daemon.GetMap has not yet been implemented")
		}),
		DaemonGetMapNameHandler: daemon.GetMapNameHandlerFunc(func(params daemon.GetMapNameParams) middleware.Responder {
			return middleware.NotImplemented("operation daemon.GetMapName has not yet been implemented")
		}),
		MetricsGetMetricsHandler: metrics.GetMetricsHandlerFunc(func(params metrics.GetMetricsParams) middleware.Responder {
			return middleware.NotImplemented("operation metrics.GetMetrics has not yet been implemented")
		}),
		PolicyGetPolicyHandler: policy.GetPolicyHandlerFunc(func(params policy.GetPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetPolicy has not yet been implemented")
		}),
		PolicyGetPolicyResolveHandler: policy.GetPolicyResolveHandlerFunc(func(params policy.GetPolicyResolveParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetPolicyResolve has not yet been implemented")
		}),
		PolicyGetPolicySelectorsHandler: policy.GetPolicySelectorsHandlerFunc(func(params policy.GetPolicySelectorsParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetPolicySelectors has not yet been implemented")
		}),
		PrefilterGetPrefilterHandler: prefilter.GetPrefilterHandlerFunc(func(params prefilter.GetPrefilterParams) middleware.Responder {
			return middleware.NotImplemented("operation prefilter.GetPrefilter has not yet been implemented")
		}),
		ServiceGetServiceHandler: service.GetServiceHandlerFunc(func(params service.GetServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation service.GetService has not yet been implemented")
		}),
		ServiceGetServiceIDHandler: service.GetServiceIDHandlerFunc(func(params service.GetServiceIDParams) middleware.Responder {
			return middleware.NotImplemented("operation service.GetServiceID has not yet been implemented")
		}),
		DaemonPatchConfigHandler: daemon.PatchConfigHandlerFunc(func(params daemon.PatchConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation daemon.PatchConfig has not yet been implemented")
		}),
		EndpointPatchEndpointIDHandler: endpoint.PatchEndpointIDHandlerFunc(func(params endpoint.PatchEndpointIDParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.PatchEndpointID has not yet been implemented")
		}),
		EndpointPatchEndpointIDConfigHandler: endpoint.PatchEndpointIDConfigHandlerFunc(func(params endpoint.PatchEndpointIDConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.PatchEndpointIDConfig has not yet been implemented")
		}),
		EndpointPatchEndpointIDLabelsHandler: endpoint.PatchEndpointIDLabelsHandlerFunc(func(params endpoint.PatchEndpointIDLabelsParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.PatchEndpointIDLabels has not yet been implemented")
		}),
		PrefilterPatchPrefilterHandler: prefilter.PatchPrefilterHandlerFunc(func(params prefilter.PatchPrefilterParams) middleware.Responder {
			return middleware.NotImplemented("operation prefilter.PatchPrefilter has not yet been implemented")
		}),
		IpamPostIpamHandler: ipam.PostIpamHandlerFunc(func(params ipam.PostIpamParams) middleware.Responder {
			return middleware.NotImplemented("operation ipam.PostIpam has not yet been implemented")
		}),
		IpamPostIpamIPHandler: ipam.PostIpamIPHandlerFunc(func(params ipam.PostIpamIPParams) middleware.Responder {
			return middleware.NotImplemented("operation ipam.PostIpamIP has not yet been implemented")
		}),
		EndpointPutEndpointIDHandler: endpoint.PutEndpointIDHandlerFunc(func(params endpoint.PutEndpointIDParams) middleware.Responder {
			return middleware.NotImplemented("operation endpoint.PutEndpointID has not yet been implemented")
		}),
		PolicyPutPolicyHandler: policy.PutPolicyHandlerFunc(func(params policy.PutPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.PutPolicy has not yet been implemented")
		}),
		ServicePutServiceIDHandler: service.PutServiceIDHandlerFunc(func(params service.PutServiceIDParams) middleware.Responder {
			return middleware.NotImplemented("operation service.PutServiceID has not yet been implemented")
		}),
	}
}

/*CiliumAPIAPI Cilium */
type CiliumAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// EndpointDeleteEndpointIDHandler sets the operation handler for the delete endpoint ID operation
	EndpointDeleteEndpointIDHandler endpoint.DeleteEndpointIDHandler
	// PolicyDeleteFqdnCacheHandler sets the operation handler for the delete fqdn cache operation
	PolicyDeleteFqdnCacheHandler policy.DeleteFqdnCacheHandler
	// IpamDeleteIpamIPHandler sets the operation handler for the delete ipam IP operation
	IpamDeleteIpamIPHandler ipam.DeleteIpamIPHandler
	// PolicyDeletePolicyHandler sets the operation handler for the delete policy operation
	PolicyDeletePolicyHandler policy.DeletePolicyHandler
	// PrefilterDeletePrefilterHandler sets the operation handler for the delete prefilter operation
	PrefilterDeletePrefilterHandler prefilter.DeletePrefilterHandler
	// ServiceDeleteServiceIDHandler sets the operation handler for the delete service ID operation
	ServiceDeleteServiceIDHandler service.DeleteServiceIDHandler
	// DaemonGetClusterNodesHandler sets the operation handler for the get cluster nodes operation
	DaemonGetClusterNodesHandler daemon.GetClusterNodesHandler
	// DaemonGetConfigHandler sets the operation handler for the get config operation
	DaemonGetConfigHandler daemon.GetConfigHandler
	// DaemonGetDebuginfoHandler sets the operation handler for the get debuginfo operation
	DaemonGetDebuginfoHandler daemon.GetDebuginfoHandler
	// EndpointGetEndpointHandler sets the operation handler for the get endpoint operation
	EndpointGetEndpointHandler endpoint.GetEndpointHandler
	// EndpointGetEndpointIDHandler sets the operation handler for the get endpoint ID operation
	EndpointGetEndpointIDHandler endpoint.GetEndpointIDHandler
	// EndpointGetEndpointIDConfigHandler sets the operation handler for the get endpoint ID config operation
	EndpointGetEndpointIDConfigHandler endpoint.GetEndpointIDConfigHandler
	// EndpointGetEndpointIDHealthzHandler sets the operation handler for the get endpoint ID healthz operation
	EndpointGetEndpointIDHealthzHandler endpoint.GetEndpointIDHealthzHandler
	// EndpointGetEndpointIDLabelsHandler sets the operation handler for the get endpoint ID labels operation
	EndpointGetEndpointIDLabelsHandler endpoint.GetEndpointIDLabelsHandler
	// EndpointGetEndpointIDLogHandler sets the operation handler for the get endpoint ID log operation
	EndpointGetEndpointIDLogHandler endpoint.GetEndpointIDLogHandler
	// PolicyGetFqdnCacheHandler sets the operation handler for the get fqdn cache operation
	PolicyGetFqdnCacheHandler policy.GetFqdnCacheHandler
	// PolicyGetFqdnCacheIDHandler sets the operation handler for the get fqdn cache ID operation
	PolicyGetFqdnCacheIDHandler policy.GetFqdnCacheIDHandler
	// PolicyGetFqdnNamesHandler sets the operation handler for the get fqdn names operation
	PolicyGetFqdnNamesHandler policy.GetFqdnNamesHandler
	// DaemonGetHealthzHandler sets the operation handler for the get healthz operation
	DaemonGetHealthzHandler daemon.GetHealthzHandler
	// PolicyGetIPHandler sets the operation handler for the get IP operation
	PolicyGetIPHandler policy.GetIPHandler
	// PolicyGetIdentityHandler sets the operation handler for the get identity operation
	PolicyGetIdentityHandler policy.GetIdentityHandler
	// PolicyGetIdentityEndpointsHandler sets the operation handler for the get identity endpoints operation
	PolicyGetIdentityEndpointsHandler policy.GetIdentityEndpointsHandler
	// PolicyGetIdentityIDHandler sets the operation handler for the get identity ID operation
	PolicyGetIdentityIDHandler policy.GetIdentityIDHandler
	// DaemonGetMapHandler sets the operation handler for the get map operation
	DaemonGetMapHandler daemon.GetMapHandler
	// DaemonGetMapNameHandler sets the operation handler for the get map name operation
	DaemonGetMapNameHandler daemon.GetMapNameHandler
	// MetricsGetMetricsHandler sets the operation handler for the get metrics operation
	MetricsGetMetricsHandler metrics.GetMetricsHandler
	// PolicyGetPolicyHandler sets the operation handler for the get policy operation
	PolicyGetPolicyHandler policy.GetPolicyHandler
	// PolicyGetPolicyResolveHandler sets the operation handler for the get policy resolve operation
	PolicyGetPolicyResolveHandler policy.GetPolicyResolveHandler
	// PolicyGetPolicySelectorsHandler sets the operation handler for the get policy selectors operation
	PolicyGetPolicySelectorsHandler policy.GetPolicySelectorsHandler
	// PrefilterGetPrefilterHandler sets the operation handler for the get prefilter operation
	PrefilterGetPrefilterHandler prefilter.GetPrefilterHandler
	// ServiceGetServiceHandler sets the operation handler for the get service operation
	ServiceGetServiceHandler service.GetServiceHandler
	// ServiceGetServiceIDHandler sets the operation handler for the get service ID operation
	ServiceGetServiceIDHandler service.GetServiceIDHandler
	// DaemonPatchConfigHandler sets the operation handler for the patch config operation
	DaemonPatchConfigHandler daemon.PatchConfigHandler
	// EndpointPatchEndpointIDHandler sets the operation handler for the patch endpoint ID operation
	EndpointPatchEndpointIDHandler endpoint.PatchEndpointIDHandler
	// EndpointPatchEndpointIDConfigHandler sets the operation handler for the patch endpoint ID config operation
	EndpointPatchEndpointIDConfigHandler endpoint.PatchEndpointIDConfigHandler
	// EndpointPatchEndpointIDLabelsHandler sets the operation handler for the patch endpoint ID labels operation
	EndpointPatchEndpointIDLabelsHandler endpoint.PatchEndpointIDLabelsHandler
	// PrefilterPatchPrefilterHandler sets the operation handler for the patch prefilter operation
	PrefilterPatchPrefilterHandler prefilter.PatchPrefilterHandler
	// IpamPostIpamHandler sets the operation handler for the post ipam operation
	IpamPostIpamHandler ipam.PostIpamHandler
	// IpamPostIpamIPHandler sets the operation handler for the post ipam IP operation
	IpamPostIpamIPHandler ipam.PostIpamIPHandler
	// EndpointPutEndpointIDHandler sets the operation handler for the put endpoint ID operation
	EndpointPutEndpointIDHandler endpoint.PutEndpointIDHandler
	// PolicyPutPolicyHandler sets the operation handler for the put policy operation
	PolicyPutPolicyHandler policy.PutPolicyHandler
	// ServicePutServiceIDHandler sets the operation handler for the put service ID operation
	ServicePutServiceIDHandler service.PutServiceIDHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *CiliumAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *CiliumAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *CiliumAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CiliumAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CiliumAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CiliumAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CiliumAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CiliumAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CiliumAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CiliumAPIAPI
func (o *CiliumAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.EndpointDeleteEndpointIDHandler == nil {
		unregistered = append(unregistered, "endpoint.DeleteEndpointIDHandler")
	}
	if o.PolicyDeleteFqdnCacheHandler == nil {
		unregistered = append(unregistered, "policy.DeleteFqdnCacheHandler")
	}
	if o.IpamDeleteIpamIPHandler == nil {
		unregistered = append(unregistered, "ipam.DeleteIpamIPHandler")
	}
	if o.PolicyDeletePolicyHandler == nil {
		unregistered = append(unregistered, "policy.DeletePolicyHandler")
	}
	if o.PrefilterDeletePrefilterHandler == nil {
		unregistered = append(unregistered, "prefilter.DeletePrefilterHandler")
	}
	if o.ServiceDeleteServiceIDHandler == nil {
		unregistered = append(unregistered, "service.DeleteServiceIDHandler")
	}
	if o.DaemonGetClusterNodesHandler == nil {
		unregistered = append(unregistered, "daemon.GetClusterNodesHandler")
	}
	if o.DaemonGetConfigHandler == nil {
		unregistered = append(unregistered, "daemon.GetConfigHandler")
	}
	if o.DaemonGetDebuginfoHandler == nil {
		unregistered = append(unregistered, "daemon.GetDebuginfoHandler")
	}
	if o.EndpointGetEndpointHandler == nil {
		unregistered = append(unregistered, "endpoint.GetEndpointHandler")
	}
	if o.EndpointGetEndpointIDHandler == nil {
		unregistered = append(unregistered, "endpoint.GetEndpointIDHandler")
	}
	if o.EndpointGetEndpointIDConfigHandler == nil {
		unregistered = append(unregistered, "endpoint.GetEndpointIDConfigHandler")
	}
	if o.EndpointGetEndpointIDHealthzHandler == nil {
		unregistered = append(unregistered, "endpoint.GetEndpointIDHealthzHandler")
	}
	if o.EndpointGetEndpointIDLabelsHandler == nil {
		unregistered = append(unregistered, "endpoint.GetEndpointIDLabelsHandler")
	}
	if o.EndpointGetEndpointIDLogHandler == nil {
		unregistered = append(unregistered, "endpoint.GetEndpointIDLogHandler")
	}
	if o.PolicyGetFqdnCacheHandler == nil {
		unregistered = append(unregistered, "policy.GetFqdnCacheHandler")
	}
	if o.PolicyGetFqdnCacheIDHandler == nil {
		unregistered = append(unregistered, "policy.GetFqdnCacheIDHandler")
	}
	if o.PolicyGetFqdnNamesHandler == nil {
		unregistered = append(unregistered, "policy.GetFqdnNamesHandler")
	}
	if o.DaemonGetHealthzHandler == nil {
		unregistered = append(unregistered, "daemon.GetHealthzHandler")
	}
	if o.PolicyGetIPHandler == nil {
		unregistered = append(unregistered, "policy.GetIPHandler")
	}
	if o.PolicyGetIdentityHandler == nil {
		unregistered = append(unregistered, "policy.GetIdentityHandler")
	}
	if o.PolicyGetIdentityEndpointsHandler == nil {
		unregistered = append(unregistered, "policy.GetIdentityEndpointsHandler")
	}
	if o.PolicyGetIdentityIDHandler == nil {
		unregistered = append(unregistered, "policy.GetIdentityIDHandler")
	}
	if o.DaemonGetMapHandler == nil {
		unregistered = append(unregistered, "daemon.GetMapHandler")
	}
	if o.DaemonGetMapNameHandler == nil {
		unregistered = append(unregistered, "daemon.GetMapNameHandler")
	}
	if o.MetricsGetMetricsHandler == nil {
		unregistered = append(unregistered, "metrics.GetMetricsHandler")
	}
	if o.PolicyGetPolicyHandler == nil {
		unregistered = append(unregistered, "policy.GetPolicyHandler")
	}
	if o.PolicyGetPolicyResolveHandler == nil {
		unregistered = append(unregistered, "policy.GetPolicyResolveHandler")
	}
	if o.PolicyGetPolicySelectorsHandler == nil {
		unregistered = append(unregistered, "policy.GetPolicySelectorsHandler")
	}
	if o.PrefilterGetPrefilterHandler == nil {
		unregistered = append(unregistered, "prefilter.GetPrefilterHandler")
	}
	if o.ServiceGetServiceHandler == nil {
		unregistered = append(unregistered, "service.GetServiceHandler")
	}
	if o.ServiceGetServiceIDHandler == nil {
		unregistered = append(unregistered, "service.GetServiceIDHandler")
	}
	if o.DaemonPatchConfigHandler == nil {
		unregistered = append(unregistered, "daemon.PatchConfigHandler")
	}
	if o.EndpointPatchEndpointIDHandler == nil {
		unregistered = append(unregistered, "endpoint.PatchEndpointIDHandler")
	}
	if o.EndpointPatchEndpointIDConfigHandler == nil {
		unregistered = append(unregistered, "endpoint.PatchEndpointIDConfigHandler")
	}
	if o.EndpointPatchEndpointIDLabelsHandler == nil {
		unregistered = append(unregistered, "endpoint.PatchEndpointIDLabelsHandler")
	}
	if o.PrefilterPatchPrefilterHandler == nil {
		unregistered = append(unregistered, "prefilter.PatchPrefilterHandler")
	}
	if o.IpamPostIpamHandler == nil {
		unregistered = append(unregistered, "ipam.PostIpamHandler")
	}
	if o.IpamPostIpamIPHandler == nil {
		unregistered = append(unregistered, "ipam.PostIpamIPHandler")
	}
	if o.EndpointPutEndpointIDHandler == nil {
		unregistered = append(unregistered, "endpoint.PutEndpointIDHandler")
	}
	if o.PolicyPutPolicyHandler == nil {
		unregistered = append(unregistered, "policy.PutPolicyHandler")
	}
	if o.ServicePutServiceIDHandler == nil {
		unregistered = append(unregistered, "service.PutServiceIDHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CiliumAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CiliumAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *CiliumAPIAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *CiliumAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *CiliumAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CiliumAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the cilium API API
func (o *CiliumAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CiliumAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/endpoint/{id}"] = endpoint.NewDeleteEndpointID(o.context, o.EndpointDeleteEndpointIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/fqdn/cache"] = policy.NewDeleteFqdnCache(o.context, o.PolicyDeleteFqdnCacheHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/ipam/{ip}"] = ipam.NewDeleteIpamIP(o.context, o.IpamDeleteIpamIPHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/policy"] = policy.NewDeletePolicy(o.context, o.PolicyDeletePolicyHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/prefilter"] = prefilter.NewDeletePrefilter(o.context, o.PrefilterDeletePrefilterHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/service/{id}"] = service.NewDeleteServiceID(o.context, o.ServiceDeleteServiceIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/cluster/nodes"] = daemon.NewGetClusterNodes(o.context, o.DaemonGetClusterNodesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/config"] = daemon.NewGetConfig(o.context, o.DaemonGetConfigHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/debuginfo"] = daemon.NewGetDebuginfo(o.context, o.DaemonGetDebuginfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/endpoint"] = endpoint.NewGetEndpoint(o.context, o.EndpointGetEndpointHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/endpoint/{id}"] = endpoint.NewGetEndpointID(o.context, o.EndpointGetEndpointIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/endpoint/{id}/config"] = endpoint.NewGetEndpointIDConfig(o.context, o.EndpointGetEndpointIDConfigHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/endpoint/{id}/healthz"] = endpoint.NewGetEndpointIDHealthz(o.context, o.EndpointGetEndpointIDHealthzHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/endpoint/{id}/labels"] = endpoint.NewGetEndpointIDLabels(o.context, o.EndpointGetEndpointIDLabelsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/endpoint/{id}/log"] = endpoint.NewGetEndpointIDLog(o.context, o.EndpointGetEndpointIDLogHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/fqdn/cache"] = policy.NewGetFqdnCache(o.context, o.PolicyGetFqdnCacheHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/fqdn/cache/{id}"] = policy.NewGetFqdnCacheID(o.context, o.PolicyGetFqdnCacheIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/fqdn/names"] = policy.NewGetFqdnNames(o.context, o.PolicyGetFqdnNamesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/healthz"] = daemon.NewGetHealthz(o.context, o.DaemonGetHealthzHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ip"] = policy.NewGetIP(o.context, o.PolicyGetIPHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/identity"] = policy.NewGetIdentity(o.context, o.PolicyGetIdentityHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/identity/endpoints"] = policy.NewGetIdentityEndpoints(o.context, o.PolicyGetIdentityEndpointsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/identity/{id}"] = policy.NewGetIdentityID(o.context, o.PolicyGetIdentityIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/map"] = daemon.NewGetMap(o.context, o.DaemonGetMapHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/map/{name}"] = daemon.NewGetMapName(o.context, o.DaemonGetMapNameHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/metrics"] = metrics.NewGetMetrics(o.context, o.MetricsGetMetricsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/policy"] = policy.NewGetPolicy(o.context, o.PolicyGetPolicyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/policy/resolve"] = policy.NewGetPolicyResolve(o.context, o.PolicyGetPolicyResolveHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/policy/selectors"] = policy.NewGetPolicySelectors(o.context, o.PolicyGetPolicySelectorsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/prefilter"] = prefilter.NewGetPrefilter(o.context, o.PrefilterGetPrefilterHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/service"] = service.NewGetService(o.context, o.ServiceGetServiceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/service/{id}"] = service.NewGetServiceID(o.context, o.ServiceGetServiceIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/config"] = daemon.NewPatchConfig(o.context, o.DaemonPatchConfigHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/endpoint/{id}"] = endpoint.NewPatchEndpointID(o.context, o.EndpointPatchEndpointIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/endpoint/{id}/config"] = endpoint.NewPatchEndpointIDConfig(o.context, o.EndpointPatchEndpointIDConfigHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/endpoint/{id}/labels"] = endpoint.NewPatchEndpointIDLabels(o.context, o.EndpointPatchEndpointIDLabelsHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/prefilter"] = prefilter.NewPatchPrefilter(o.context, o.PrefilterPatchPrefilterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/ipam"] = ipam.NewPostIpam(o.context, o.IpamPostIpamHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/ipam/{ip}"] = ipam.NewPostIpamIP(o.context, o.IpamPostIpamIPHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/endpoint/{id}"] = endpoint.NewPutEndpointID(o.context, o.EndpointPutEndpointIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/policy"] = policy.NewPutPolicy(o.context, o.PolicyPutPolicyHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/service/{id}"] = service.NewPutServiceID(o.context, o.ServicePutServiceIDHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CiliumAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CiliumAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CiliumAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CiliumAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *CiliumAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
