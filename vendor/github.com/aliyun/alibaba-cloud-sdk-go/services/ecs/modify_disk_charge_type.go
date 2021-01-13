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

// ModifyDiskChargeType invokes the ecs.ModifyDiskChargeType API synchronously
func (client *Client) ModifyDiskChargeType(request *ModifyDiskChargeTypeRequest) (response *ModifyDiskChargeTypeResponse, err error) {
	response = CreateModifyDiskChargeTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDiskChargeTypeWithChan invokes the ecs.ModifyDiskChargeType API asynchronously
func (client *Client) ModifyDiskChargeTypeWithChan(request *ModifyDiskChargeTypeRequest) (<-chan *ModifyDiskChargeTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyDiskChargeTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDiskChargeType(request)
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

// ModifyDiskChargeTypeWithCallback invokes the ecs.ModifyDiskChargeType API asynchronously
func (client *Client) ModifyDiskChargeTypeWithCallback(request *ModifyDiskChargeTypeRequest, callback func(response *ModifyDiskChargeTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDiskChargeTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDiskChargeType(request)
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

// ModifyDiskChargeTypeRequest is the request struct for api ModifyDiskChargeType
type ModifyDiskChargeTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DiskChargeType       string           `position:"Query" name:"DiskChargeType"`
	DiskIds              string           `position:"Query" name:"DiskIds"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// ModifyDiskChargeTypeResponse is the response struct for api ModifyDiskChargeType
type ModifyDiskChargeTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateModifyDiskChargeTypeRequest creates a request to invoke ModifyDiskChargeType API
func CreateModifyDiskChargeTypeRequest() (request *ModifyDiskChargeTypeRequest) {
	request = &ModifyDiskChargeTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyDiskChargeType", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDiskChargeTypeResponse creates a response to parse from ModifyDiskChargeType response
func CreateModifyDiskChargeTypeResponse() (response *ModifyDiskChargeTypeResponse) {
	response = &ModifyDiskChargeTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
