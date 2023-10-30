// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package statedb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new statedb API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for statedb API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetStatedbDump(params *GetStatedbDumpParams, writer io.Writer, opts ...ClientOption) (*GetStatedbDumpOK, error)

	GetStatedbQueryTable(params *GetStatedbQueryTableParams, writer io.Writer, opts ...ClientOption) (*GetStatedbQueryTableOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetStatedbDump dumps state d b contents
*/
func (a *Client) GetStatedbDump(params *GetStatedbDumpParams, writer io.Writer, opts ...ClientOption) (*GetStatedbDumpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatedbDumpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStatedbDump",
		Method:             "GET",
		PathPattern:        "/statedb/dump",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStatedbDumpReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*GetStatedbDumpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStatedbDump: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetStatedbQueryTable performs a query against a state d b table
*/
func (a *Client) GetStatedbQueryTable(params *GetStatedbQueryTableParams, writer io.Writer, opts ...ClientOption) (*GetStatedbQueryTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatedbQueryTableParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetStatedbQueryTable",
		Method:             "GET",
		PathPattern:        "/statedb/query/{table}",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetStatedbQueryTableReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*GetStatedbQueryTableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStatedbQueryTable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
