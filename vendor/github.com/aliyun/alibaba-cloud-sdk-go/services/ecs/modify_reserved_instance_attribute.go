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

// ModifyReservedInstanceAttribute invokes the ecs.ModifyReservedInstanceAttribute API synchronously
func (client *Client) ModifyReservedInstanceAttribute(request *ModifyReservedInstanceAttributeRequest) (response *ModifyReservedInstanceAttributeResponse, err error) {
	response = CreateModifyReservedInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyReservedInstanceAttributeWithChan invokes the ecs.ModifyReservedInstanceAttribute API asynchronously
func (client *Client) ModifyReservedInstanceAttributeWithChan(request *ModifyReservedInstanceAttributeRequest) (<-chan *ModifyReservedInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyReservedInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyReservedInstanceAttribute(request)
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

// ModifyReservedInstanceAttributeWithCallback invokes the ecs.ModifyReservedInstanceAttribute API asynchronously
func (client *Client) ModifyReservedInstanceAttributeWithCallback(request *ModifyReservedInstanceAttributeRequest, callback func(response *ModifyReservedInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyReservedInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyReservedInstanceAttribute(request)
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

// ModifyReservedInstanceAttributeRequest is the request struct for api ModifyReservedInstanceAttribute
type ModifyReservedInstanceAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ReservedInstanceId   string           `position:"Query" name:"ReservedInstanceId"`
	ReservedInstanceName string           `position:"Query" name:"ReservedInstanceName"`
}

// ModifyReservedInstanceAttributeResponse is the response struct for api ModifyReservedInstanceAttribute
type ModifyReservedInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateModifyReservedInstanceAttributeRequest creates a request to invoke ModifyReservedInstanceAttribute API
func CreateModifyReservedInstanceAttributeRequest() (request *ModifyReservedInstanceAttributeRequest) {
	request = &ModifyReservedInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyReservedInstanceAttribute", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyReservedInstanceAttributeResponse creates a response to parse from ModifyReservedInstanceAttribute response
func CreateModifyReservedInstanceAttributeResponse() (response *ModifyReservedInstanceAttributeResponse) {
	response = &ModifyReservedInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
