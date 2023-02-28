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

// AttachNetworkInterface invokes the ecs.AttachNetworkInterface API synchronously
func (client *Client) AttachNetworkInterface(request *AttachNetworkInterfaceRequest) (response *AttachNetworkInterfaceResponse, err error) {
	response = CreateAttachNetworkInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// AttachNetworkInterfaceWithChan invokes the ecs.AttachNetworkInterface API asynchronously
func (client *Client) AttachNetworkInterfaceWithChan(request *AttachNetworkInterfaceRequest) (<-chan *AttachNetworkInterfaceResponse, <-chan error) {
	responseChan := make(chan *AttachNetworkInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachNetworkInterface(request)
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

// AttachNetworkInterfaceWithCallback invokes the ecs.AttachNetworkInterface API asynchronously
func (client *Client) AttachNetworkInterfaceWithCallback(request *AttachNetworkInterfaceRequest, callback func(response *AttachNetworkInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachNetworkInterfaceResponse
		var err error
		defer close(result)
		response, err = client.AttachNetworkInterface(request)
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

// AttachNetworkInterfaceRequest is the request struct for api AttachNetworkInterface
type AttachNetworkInterfaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TrunkNetworkInstanceId           string           `position:"Query" name:"TrunkNetworkInstanceId"`
	ResourceOwnerAccount             string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                     string           `position:"Query" name:"OwnerAccount"`
	WaitForNetworkConfigurationReady requests.Boolean `position:"Query" name:"WaitForNetworkConfigurationReady"`
	OwnerId                          requests.Integer `position:"Query" name:"OwnerId"`
	NetworkCardIndex                 requests.Integer `position:"Query" name:"NetworkCardIndex"`
	InstanceId                       string           `position:"Query" name:"InstanceId"`
	NetworkInterfaceId               string           `position:"Query" name:"NetworkInterfaceId"`
}

// AttachNetworkInterfaceResponse is the response struct for api AttachNetworkInterface
type AttachNetworkInterfaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachNetworkInterfaceRequest creates a request to invoke AttachNetworkInterface API
func CreateAttachNetworkInterfaceRequest() (request *AttachNetworkInterfaceRequest) {
	request = &AttachNetworkInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AttachNetworkInterface", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachNetworkInterfaceResponse creates a response to parse from AttachNetworkInterface response
func CreateAttachNetworkInterfaceResponse() (response *AttachNetworkInterfaceResponse) {
	response = &AttachNetworkInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
