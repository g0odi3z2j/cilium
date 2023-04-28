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

// ModifyLaunchTemplateDefaultVersion invokes the ecs.ModifyLaunchTemplateDefaultVersion API synchronously
func (client *Client) ModifyLaunchTemplateDefaultVersion(request *ModifyLaunchTemplateDefaultVersionRequest) (response *ModifyLaunchTemplateDefaultVersionResponse, err error) {
	response = CreateModifyLaunchTemplateDefaultVersionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLaunchTemplateDefaultVersionWithChan invokes the ecs.ModifyLaunchTemplateDefaultVersion API asynchronously
func (client *Client) ModifyLaunchTemplateDefaultVersionWithChan(request *ModifyLaunchTemplateDefaultVersionRequest) (<-chan *ModifyLaunchTemplateDefaultVersionResponse, <-chan error) {
	responseChan := make(chan *ModifyLaunchTemplateDefaultVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLaunchTemplateDefaultVersion(request)
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

// ModifyLaunchTemplateDefaultVersionWithCallback invokes the ecs.ModifyLaunchTemplateDefaultVersion API asynchronously
func (client *Client) ModifyLaunchTemplateDefaultVersionWithCallback(request *ModifyLaunchTemplateDefaultVersionRequest, callback func(response *ModifyLaunchTemplateDefaultVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLaunchTemplateDefaultVersionResponse
		var err error
		defer close(result)
		response, err = client.ModifyLaunchTemplateDefaultVersion(request)
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

// ModifyLaunchTemplateDefaultVersionRequest is the request struct for api ModifyLaunchTemplateDefaultVersion
type ModifyLaunchTemplateDefaultVersionRequest struct {
	*requests.RpcRequest
	LaunchTemplateName   string           `position:"Query" name:"LaunchTemplateName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LaunchTemplateId     string           `position:"Query" name:"LaunchTemplateId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DefaultVersionNumber requests.Integer `position:"Query" name:"DefaultVersionNumber"`
}

// ModifyLaunchTemplateDefaultVersionResponse is the response struct for api ModifyLaunchTemplateDefaultVersion
type ModifyLaunchTemplateDefaultVersionResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	LaunchTemplateId string `json:"LaunchTemplateId" xml:"LaunchTemplateId"`
}

// CreateModifyLaunchTemplateDefaultVersionRequest creates a request to invoke ModifyLaunchTemplateDefaultVersion API
func CreateModifyLaunchTemplateDefaultVersionRequest() (request *ModifyLaunchTemplateDefaultVersionRequest) {
	request = &ModifyLaunchTemplateDefaultVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyLaunchTemplateDefaultVersion", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLaunchTemplateDefaultVersionResponse creates a response to parse from ModifyLaunchTemplateDefaultVersion response
func CreateModifyLaunchTemplateDefaultVersionResponse() (response *ModifyLaunchTemplateDefaultVersionResponse) {
	response = &ModifyLaunchTemplateDefaultVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
