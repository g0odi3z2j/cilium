package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new endpoint API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoint API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteEndpointID deletes endpoint

Deletes the endpoint specified by the ID. Deletion is imminent and
atomic, if the deletion request is valid and the endpoint exists,
deletion will occur even if errors are encounted in the process. If
errors have been encountered, the code 202 will be returned, otherwise
200 on success.

All resources associated with the endpoint will be freed and the
workload represented by the endpoint will be disconnected.It will no
longer be able to initiate or receive communications of any sort.

*/
func (a *Client) DeleteEndpointID(params *DeleteEndpointIDParams) (*DeleteEndpointIDOK, *DeleteEndpointIDErrors, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEndpointIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteEndpointID",
		Method:             "DELETE",
		PathPattern:        "/endpoint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteEndpointIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteEndpointIDOK:
		return value, nil, nil
	case *DeleteEndpointIDErrors:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetEndpoint gets list of all endpoints

Returns an array of all local endpoints.

*/
func (a *Client) GetEndpoint(params *GetEndpointParams) (*GetEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpoint",
		Method:             "GET",
		PathPattern:        "/endpoint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEndpointOK), nil

}

/*
GetEndpointID gets endpoint by endpoint ID

Returns endpoint information

*/
func (a *Client) GetEndpointID(params *GetEndpointIDParams) (*GetEndpointIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpointID",
		Method:             "GET",
		PathPattern:        "/endpoint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEndpointIDOK), nil

}

/*
GetEndpointIDConfig retrieves endpoint configuration

Retrieves the configuration of the specified endpoint.

*/
func (a *Client) GetEndpointIDConfig(params *GetEndpointIDConfigParams) (*GetEndpointIDConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointIDConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpointIDConfig",
		Method:             "GET",
		PathPattern:        "/endpoint/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointIDConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEndpointIDConfigOK), nil

}

/*
GetEndpointIDLabels retrieves the list of labels associated with an endpoint
*/
func (a *Client) GetEndpointIDLabels(params *GetEndpointIDLabelsParams) (*GetEndpointIDLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointIDLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEndpointIDLabels",
		Method:             "GET",
		PathPattern:        "/endpoint/{id}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEndpointIDLabelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEndpointIDLabelsOK), nil

}

/*
PatchEndpointID modifies existing endpoint

Applies the endpoint change request to an existing endpoint

*/
func (a *Client) PatchEndpointID(params *PatchEndpointIDParams) (*PatchEndpointIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEndpointIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEndpointID",
		Method:             "PATCH",
		PathPattern:        "/endpoint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchEndpointIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchEndpointIDOK), nil

}

/*
PatchEndpointIDConfig modifies mutable endpoint configuration

Update the configuration of an existing endpoint and regenerates &
recompiles the corresponding programs automatically.

*/
func (a *Client) PatchEndpointIDConfig(params *PatchEndpointIDConfigParams) (*PatchEndpointIDConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEndpointIDConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEndpointIDConfig",
		Method:             "PATCH",
		PathPattern:        "/endpoint/{id}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchEndpointIDConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchEndpointIDConfigOK), nil

}

/*
PutEndpointID creates endpoint

Updates an existing endpoint

*/
func (a *Client) PutEndpointID(params *PutEndpointIDParams) (*PutEndpointIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEndpointIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEndpointID",
		Method:             "PUT",
		PathPattern:        "/endpoint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutEndpointIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutEndpointIDCreated), nil

}

/*
PutEndpointIDLabels modifies label configuration of endpoint

Updates the list of labels associated with an endpoint by applying
a label modificator structure to the label configuration of an
endpoint.

The label configuration mutation is only executed as a whole, i.e.
if any of the labels to be deleted are not either on the list of
orchestration system labels, custom labels, or already disabled,
then the request will fail. Labels to be added which already exist
on either the orchestration list or custom list will be ignored.

*/
func (a *Client) PutEndpointIDLabels(params *PutEndpointIDLabelsParams) (*PutEndpointIDLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEndpointIDLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEndpointIDLabels",
		Method:             "PUT",
		PathPattern:        "/endpoint/{id}/labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutEndpointIDLabelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutEndpointIDLabelsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
