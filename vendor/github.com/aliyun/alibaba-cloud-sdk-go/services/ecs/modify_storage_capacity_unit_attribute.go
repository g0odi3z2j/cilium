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

// ModifyStorageCapacityUnitAttribute invokes the ecs.ModifyStorageCapacityUnitAttribute API synchronously
func (client *Client) ModifyStorageCapacityUnitAttribute(request *ModifyStorageCapacityUnitAttributeRequest) (response *ModifyStorageCapacityUnitAttributeResponse, err error) {
	response = CreateModifyStorageCapacityUnitAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyStorageCapacityUnitAttributeWithChan invokes the ecs.ModifyStorageCapacityUnitAttribute API asynchronously
func (client *Client) ModifyStorageCapacityUnitAttributeWithChan(request *ModifyStorageCapacityUnitAttributeRequest) (<-chan *ModifyStorageCapacityUnitAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyStorageCapacityUnitAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyStorageCapacityUnitAttribute(request)
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

// ModifyStorageCapacityUnitAttributeWithCallback invokes the ecs.ModifyStorageCapacityUnitAttribute API asynchronously
func (client *Client) ModifyStorageCapacityUnitAttributeWithCallback(request *ModifyStorageCapacityUnitAttributeRequest, callback func(response *ModifyStorageCapacityUnitAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyStorageCapacityUnitAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyStorageCapacityUnitAttribute(request)
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

// ModifyStorageCapacityUnitAttributeRequest is the request struct for api ModifyStorageCapacityUnitAttribute
type ModifyStorageCapacityUnitAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description           string           `position:"Query" name:"Description"`
	StorageCapacityUnitId string           `position:"Query" name:"StorageCapacityUnitId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	Name                  string           `position:"Query" name:"Name"`
}

// ModifyStorageCapacityUnitAttributeResponse is the response struct for api ModifyStorageCapacityUnitAttribute
type ModifyStorageCapacityUnitAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyStorageCapacityUnitAttributeRequest creates a request to invoke ModifyStorageCapacityUnitAttribute API
func CreateModifyStorageCapacityUnitAttributeRequest() (request *ModifyStorageCapacityUnitAttributeRequest) {
	request = &ModifyStorageCapacityUnitAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyStorageCapacityUnitAttribute", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyStorageCapacityUnitAttributeResponse creates a response to parse from ModifyStorageCapacityUnitAttribute response
func CreateModifyStorageCapacityUnitAttributeResponse() (response *ModifyStorageCapacityUnitAttributeResponse) {
	response = &ModifyStorageCapacityUnitAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
