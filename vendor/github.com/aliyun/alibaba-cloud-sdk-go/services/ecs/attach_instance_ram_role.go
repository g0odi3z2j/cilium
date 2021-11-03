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

// AttachInstanceRamRole invokes the ecs.AttachInstanceRamRole API synchronously
func (client *Client) AttachInstanceRamRole(request *AttachInstanceRamRoleRequest) (response *AttachInstanceRamRoleResponse, err error) {
	response = CreateAttachInstanceRamRoleResponse()
	err = client.DoAction(request, response)
	return
}

// AttachInstanceRamRoleWithChan invokes the ecs.AttachInstanceRamRole API asynchronously
func (client *Client) AttachInstanceRamRoleWithChan(request *AttachInstanceRamRoleRequest) (<-chan *AttachInstanceRamRoleResponse, <-chan error) {
	responseChan := make(chan *AttachInstanceRamRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachInstanceRamRole(request)
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

// AttachInstanceRamRoleWithCallback invokes the ecs.AttachInstanceRamRole API asynchronously
func (client *Client) AttachInstanceRamRoleWithCallback(request *AttachInstanceRamRoleRequest, callback func(response *AttachInstanceRamRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachInstanceRamRoleResponse
		var err error
		defer close(result)
		response, err = client.AttachInstanceRamRole(request)
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

// AttachInstanceRamRoleRequest is the request struct for api AttachInstanceRamRole
type AttachInstanceRamRoleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Policy               string           `position:"Query" name:"Policy"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RamRoleName          string           `position:"Query" name:"RamRoleName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
}

// AttachInstanceRamRoleResponse is the response struct for api AttachInstanceRamRole
type AttachInstanceRamRoleResponse struct {
	*responses.BaseResponse
	RamRoleName                  string                       `json:"RamRoleName" xml:"RamRoleName"`
	RequestId                    string                       `json:"RequestId" xml:"RequestId"`
	TotalCount                   int                          `json:"TotalCount" xml:"TotalCount"`
	FailCount                    int                          `json:"FailCount" xml:"FailCount"`
	AttachInstanceRamRoleResults AttachInstanceRamRoleResults `json:"AttachInstanceRamRoleResults" xml:"AttachInstanceRamRoleResults"`
}

// CreateAttachInstanceRamRoleRequest creates a request to invoke AttachInstanceRamRole API
func CreateAttachInstanceRamRoleRequest() (request *AttachInstanceRamRoleRequest) {
	request = &AttachInstanceRamRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AttachInstanceRamRole", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachInstanceRamRoleResponse creates a response to parse from AttachInstanceRamRole response
func CreateAttachInstanceRamRoleResponse() (response *AttachInstanceRamRoleResponse) {
	response = &AttachInstanceRamRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
