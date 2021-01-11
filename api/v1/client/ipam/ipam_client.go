// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ipam API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ipam API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteIpamIP(params *DeleteIpamIPParams) (*DeleteIpamIPOK, error)

	PostIpam(params *PostIpamParams) (*PostIpamCreated, error)

	PostIpamIP(params *PostIpamIPParams) (*PostIpamIPOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteIpamIP releases an allocated IP address
*/
func (a *Client) DeleteIpamIP(params *DeleteIpamIPParams) (*DeleteIpamIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIpamIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteIpamIP",
		Method:             "DELETE",
		PathPattern:        "/ipam/{ip}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteIpamIPReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteIpamIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteIpamIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostIpam allocates an IP address
*/
func (a *Client) PostIpam(params *PostIpamParams) (*PostIpamCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIpamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIpam",
		Method:             "POST",
		PathPattern:        "/ipam",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostIpamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIpamCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostIpam: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostIpamIP allocates an IP address
*/
func (a *Client) PostIpamIP(params *PostIpamIPParams) (*PostIpamIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIpamIPParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIpamIP",
		Method:             "POST",
		PathPattern:        "/ipam/{ip}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostIpamIPReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIpamIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostIpamIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
