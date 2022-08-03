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

// CreateCommand invokes the ecs.CreateCommand API synchronously
func (client *Client) CreateCommand(request *CreateCommandRequest) (response *CreateCommandResponse, err error) {
	response = CreateCreateCommandResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCommandWithChan invokes the ecs.CreateCommand API asynchronously
func (client *Client) CreateCommandWithChan(request *CreateCommandRequest) (<-chan *CreateCommandResponse, <-chan error) {
	responseChan := make(chan *CreateCommandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCommand(request)
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

// CreateCommandWithCallback invokes the ecs.CreateCommand API asynchronously
func (client *Client) CreateCommandWithCallback(request *CreateCommandRequest, callback func(response *CreateCommandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCommandResponse
		var err error
		defer close(result)
		response, err = client.CreateCommand(request)
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

// CreateCommandRequest is the request struct for api CreateCommand
type CreateCommandRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	WorkingDir           string           `position:"Query" name:"WorkingDir"`
	Description          string           `position:"Query" name:"Description"`
	Type                 string           `position:"Query" name:"Type"`
	CommandContent       string           `position:"Query" name:"CommandContent"`
	Timeout              requests.Integer `position:"Query" name:"Timeout"`
	ContentEncoding      string           `position:"Query" name:"ContentEncoding"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
	EnableParameter      requests.Boolean `position:"Query" name:"EnableParameter"`
}

// CreateCommandResponse is the response struct for api CreateCommand
type CreateCommandResponse struct {
	*responses.BaseResponse
	CommandId string `json:"CommandId" xml:"CommandId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateCommandRequest creates a request to invoke CreateCommand API
func CreateCreateCommandRequest() (request *CreateCommandRequest) {
	request = &CreateCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateCommand", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCommandResponse creates a response to parse from CreateCommand response
func CreateCreateCommandResponse() (response *CreateCommandResponse) {
	response = &CreateCommandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
