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

// AttachClassicLinkVpc invokes the ecs.AttachClassicLinkVpc API synchronously
func (client *Client) AttachClassicLinkVpc(request *AttachClassicLinkVpcRequest) (response *AttachClassicLinkVpcResponse, err error) {
	response = CreateAttachClassicLinkVpcResponse()
	err = client.DoAction(request, response)
	return
}

// AttachClassicLinkVpcWithChan invokes the ecs.AttachClassicLinkVpc API asynchronously
func (client *Client) AttachClassicLinkVpcWithChan(request *AttachClassicLinkVpcRequest) (<-chan *AttachClassicLinkVpcResponse, <-chan error) {
	responseChan := make(chan *AttachClassicLinkVpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachClassicLinkVpc(request)
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

// AttachClassicLinkVpcWithCallback invokes the ecs.AttachClassicLinkVpc API asynchronously
func (client *Client) AttachClassicLinkVpcWithCallback(request *AttachClassicLinkVpcRequest, callback func(response *AttachClassicLinkVpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachClassicLinkVpcResponse
		var err error
		defer close(result)
		response, err = client.AttachClassicLinkVpc(request)
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

// AttachClassicLinkVpcRequest is the request struct for api AttachClassicLinkVpc
type AttachClassicLinkVpcRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	VpcId                string           `position:"Query" name:"VpcId"`
}

// AttachClassicLinkVpcResponse is the response struct for api AttachClassicLinkVpc
type AttachClassicLinkVpcResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachClassicLinkVpcRequest creates a request to invoke AttachClassicLinkVpc API
func CreateAttachClassicLinkVpcRequest() (request *AttachClassicLinkVpcRequest) {
	request = &AttachClassicLinkVpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AttachClassicLinkVpc", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachClassicLinkVpcResponse creates a response to parse from AttachClassicLinkVpc response
func CreateAttachClassicLinkVpcResponse() (response *AttachClassicLinkVpcResponse) {
	response = &AttachClassicLinkVpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
