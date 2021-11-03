package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreatePhysicalConnection invokes the ecs.CreatePhysicalConnection API synchronously
func (client *Client) CreatePhysicalConnection(request *CreatePhysicalConnectionRequest) (response *CreatePhysicalConnectionResponse, err error) {
	response = CreateCreatePhysicalConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePhysicalConnectionWithChan invokes the ecs.CreatePhysicalConnection API asynchronously
func (client *Client) CreatePhysicalConnectionWithChan(request *CreatePhysicalConnectionRequest) (<-chan *CreatePhysicalConnectionResponse, <-chan error) {
	responseChan := make(chan *CreatePhysicalConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePhysicalConnection(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreatePhysicalConnectionWithCallback invokes the ecs.CreatePhysicalConnection API asynchronously
func (client *Client) CreatePhysicalConnectionWithCallback(request *CreatePhysicalConnectionRequest, callback func(response *CreatePhysicalConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePhysicalConnectionResponse
		var err error
		defer close(result)
		response, err = client.CreatePhysicalConnection(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreatePhysicalConnectionRequest is the request struct for api CreatePhysicalConnection
type CreatePhysicalConnectionRequest struct {
	*requests.RpcRequest
	AccessPointId                 string           `position:"Query" name:"AccessPointId"`
	ResourceOwnerId               requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PortType                      string           `position:"Query" name:"PortType"`
	CircuitCode                   string           `position:"Query" name:"CircuitCode"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
	Description                   string           `position:"Query" name:"Description"`
	Type                          string           `position:"Query" name:"Type"`
	UserCidr                      string           `position:"Query" name:"UserCidr"`
	RedundantPhysicalConnectionId string           `position:"Query" name:"RedundantPhysicalConnectionId"`
	PeerLocation                  string           `position:"Query" name:"PeerLocation"`
	Bandwidth                     requests.Integer `position:"Query" name:"bandwidth"`
	ResourceOwnerAccount          string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string           `position:"Query" name:"OwnerAccount"`
	OwnerId                       requests.Integer `position:"Query" name:"OwnerId"`
	LineOperator                  string           `position:"Query" name:"LineOperator"`
	Name                          string           `position:"Query" name:"Name"`
}

// CreatePhysicalConnectionResponse is the response struct for api CreatePhysicalConnection
type CreatePhysicalConnectionResponse struct {
	*responses.BaseResponse
	PhysicalConnectionId string `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
	RequestId            string `json:"RequestId" xml:"RequestId"`
}

// CreateCreatePhysicalConnectionRequest creates a request to invoke CreatePhysicalConnection API
func CreateCreatePhysicalConnectionRequest() (request *CreatePhysicalConnectionRequest) {
	request = &CreatePhysicalConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreatePhysicalConnection", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePhysicalConnectionResponse creates a response to parse from CreatePhysicalConnection response
func CreateCreatePhysicalConnectionResponse() (response *CreatePhysicalConnectionResponse) {
	response = &CreatePhysicalConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
