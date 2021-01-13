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

// InstallCloudAssistant invokes the ecs.InstallCloudAssistant API synchronously
func (client *Client) InstallCloudAssistant(request *InstallCloudAssistantRequest) (response *InstallCloudAssistantResponse, err error) {
	response = CreateInstallCloudAssistantResponse()
	err = client.DoAction(request, response)
	return
}

// InstallCloudAssistantWithChan invokes the ecs.InstallCloudAssistant API asynchronously
func (client *Client) InstallCloudAssistantWithChan(request *InstallCloudAssistantRequest) (<-chan *InstallCloudAssistantResponse, <-chan error) {
	responseChan := make(chan *InstallCloudAssistantResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallCloudAssistant(request)
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

// InstallCloudAssistantWithCallback invokes the ecs.InstallCloudAssistant API asynchronously
func (client *Client) InstallCloudAssistantWithCallback(request *InstallCloudAssistantRequest, callback func(response *InstallCloudAssistantResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallCloudAssistantResponse
		var err error
		defer close(result)
		response, err = client.InstallCloudAssistant(request)
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

// InstallCloudAssistantRequest is the request struct for api InstallCloudAssistant
type InstallCloudAssistantRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// InstallCloudAssistantResponse is the response struct for api InstallCloudAssistant
type InstallCloudAssistantResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateInstallCloudAssistantRequest creates a request to invoke InstallCloudAssistant API
func CreateInstallCloudAssistantRequest() (request *InstallCloudAssistantRequest) {
	request = &InstallCloudAssistantRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "InstallCloudAssistant", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInstallCloudAssistantResponse creates a response to parse from InstallCloudAssistant response
func CreateInstallCloudAssistantResponse() (response *InstallCloudAssistantResponse) {
	response = &InstallCloudAssistantResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
