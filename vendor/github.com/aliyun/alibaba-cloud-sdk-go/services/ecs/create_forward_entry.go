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

// CreateForwardEntry invokes the ecs.CreateForwardEntry API synchronously
func (client *Client) CreateForwardEntry(request *CreateForwardEntryRequest) (response *CreateForwardEntryResponse, err error) {
	response = CreateCreateForwardEntryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateForwardEntryWithChan invokes the ecs.CreateForwardEntry API asynchronously
func (client *Client) CreateForwardEntryWithChan(request *CreateForwardEntryRequest) (<-chan *CreateForwardEntryResponse, <-chan error) {
	responseChan := make(chan *CreateForwardEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateForwardEntry(request)
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

// CreateForwardEntryWithCallback invokes the ecs.CreateForwardEntry API asynchronously
func (client *Client) CreateForwardEntryWithCallback(request *CreateForwardEntryRequest, callback func(response *CreateForwardEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateForwardEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateForwardEntry(request)
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

// CreateForwardEntryRequest is the request struct for api CreateForwardEntry
type CreateForwardEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ForwardTableId       string           `position:"Query" name:"ForwardTableId"`
	InternalIp           string           `position:"Query" name:"InternalIp"`
	ExternalIp           string           `position:"Query" name:"ExternalIp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string           `position:"Query" name:"IpProtocol"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InternalPort         string           `position:"Query" name:"InternalPort"`
	ExternalPort         string           `position:"Query" name:"ExternalPort"`
}

// CreateForwardEntryResponse is the response struct for api CreateForwardEntry
type CreateForwardEntryResponse struct {
	*responses.BaseResponse
	ForwardEntryId string `json:"ForwardEntryId" xml:"ForwardEntryId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateForwardEntryRequest creates a request to invoke CreateForwardEntry API
func CreateCreateForwardEntryRequest() (request *CreateForwardEntryRequest) {
	request = &CreateForwardEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateForwardEntry", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateForwardEntryResponse creates a response to parse from CreateForwardEntry response
func CreateCreateForwardEntryResponse() (response *CreateForwardEntryResponse) {
	response = &CreateForwardEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
