// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new daemon API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for daemon API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCgroupDumpMetadata(params *GetCgroupDumpMetadataParams, opts ...ClientOption) (*GetCgroupDumpMetadataOK, error)

	GetClusterNodes(params *GetClusterNodesParams, opts ...ClientOption) (*GetClusterNodesOK, error)

	GetConfig(params *GetConfigParams, opts ...ClientOption) (*GetConfigOK, error)

	GetDebuginfo(params *GetDebuginfoParams, opts ...ClientOption) (*GetDebuginfoOK, error)

	GetHealth(params *GetHealthParams, opts ...ClientOption) (*GetHealthOK, error)

	GetHealthz(params *GetHealthzParams, opts ...ClientOption) (*GetHealthzOK, error)

	GetMap(params *GetMapParams, opts ...ClientOption) (*GetMapOK, error)

	GetMapName(params *GetMapNameParams, opts ...ClientOption) (*GetMapNameOK, error)

	GetMapNameEvents(params *GetMapNameEventsParams, writer io.Writer, opts ...ClientOption) (*GetMapNameEventsOK, error)

	GetNodeIds(params *GetNodeIdsParams, opts ...ClientOption) (*GetNodeIdsOK, error)

	PatchConfig(params *PatchConfigParams, opts ...ClientOption) (*PatchConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetCgroupDumpMetadata retrieves cgroup metadata for all pods
*/
func (a *Client) GetCgroupDumpMetadata(params *GetCgroupDumpMetadataParams, opts ...ClientOption) (*GetCgroupDumpMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCgroupDumpMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCgroupDumpMetadata",
		Method:             "GET",
		PathPattern:        "/cgroup-dump-metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCgroupDumpMetadataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCgroupDumpMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCgroupDumpMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetClusterNodes gets nodes information stored in the cilium agent
*/
func (a *Client) GetClusterNodes(params *GetClusterNodesParams, opts ...ClientOption) (*GetClusterNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterNodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetClusterNodes",
		Method:             "GET",
		PathPattern:        "/cluster/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterNodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetClusterNodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetConfig gets configuration of cilium daemon

Returns the configuration of the Cilium daemon.
*/
func (a *Client) GetConfig(params *GetConfigParams, opts ...ClientOption) (*GetConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetConfig",
		Method:             "GET",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDebuginfo retrieves information about the agent and environment for debugging
*/
func (a *Client) GetDebuginfo(params *GetDebuginfoParams, opts ...ClientOption) (*GetDebuginfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDebuginfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDebuginfo",
		Method:             "GET",
		PathPattern:        "/debuginfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDebuginfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDebuginfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDebuginfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHealth gets modules health of cilium daemon

Returns modules health and status information of the Cilium daemon.
*/
func (a *Client) GetHealth(params *GetHealthParams, opts ...ClientOption) (*GetHealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHealth",
		Method:             "GET",
		PathPattern:        "/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHealthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHealthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHealth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetHealthz gets health of cilium daemon

	Returns health and status information of the Cilium daemon and related

components such as the local container runtime, connected datastore,
Kubernetes integration and Hubble.
*/
func (a *Client) GetHealthz(params *GetHealthzParams, opts ...ClientOption) (*GetHealthzOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthzParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHealthz",
		Method:             "GET",
		PathPattern:        "/healthz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHealthzReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHealthzOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHealthz: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMap lists all open maps
*/
func (a *Client) GetMap(params *GetMapParams, opts ...ClientOption) (*GetMapOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMapParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMap",
		Method:             "GET",
		PathPattern:        "/map",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMapReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMapOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMap: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMapName retrieves contents of b p f map
*/
func (a *Client) GetMapName(params *GetMapNameParams, opts ...ClientOption) (*GetMapNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMapNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMapName",
		Method:             "GET",
		PathPattern:        "/map/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMapNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMapNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMapName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMapNameEvents retrieves the recent event logs associated with this endpoint
*/
func (a *Client) GetMapNameEvents(params *GetMapNameEventsParams, writer io.Writer, opts ...ClientOption) (*GetMapNameEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMapNameEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMapNameEvents",
		Method:             "GET",
		PathPattern:        "/map/{name}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMapNameEventsReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMapNameEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMapNameEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetNodeIds lists information about known node i ds

	Retrieves a list of node IDs allocated by the agent and their

associated node IP addresses.
*/
func (a *Client) GetNodeIds(params *GetNodeIdsParams, opts ...ClientOption) (*GetNodeIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeIdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNodeIds",
		Method:             "GET",
		PathPattern:        "/node/ids",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNodeIdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNodeIdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNodeIds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PatchConfig modifies daemon configuration

	Updates the daemon configuration by applying the provided

ConfigurationMap and regenerates & recompiles all required datapath
components.
*/
func (a *Client) PatchConfig(params *PatchConfigParams, opts ...ClientOption) (*PatchConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchConfig",
		Method:             "PATCH",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
