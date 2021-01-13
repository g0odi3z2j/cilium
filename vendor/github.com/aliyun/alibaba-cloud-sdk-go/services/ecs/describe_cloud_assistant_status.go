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

// DescribeCloudAssistantStatus invokes the ecs.DescribeCloudAssistantStatus API synchronously
func (client *Client) DescribeCloudAssistantStatus(request *DescribeCloudAssistantStatusRequest) (response *DescribeCloudAssistantStatusResponse, err error) {
	response = CreateDescribeCloudAssistantStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudAssistantStatusWithChan invokes the ecs.DescribeCloudAssistantStatus API asynchronously
func (client *Client) DescribeCloudAssistantStatusWithChan(request *DescribeCloudAssistantStatusRequest) (<-chan *DescribeCloudAssistantStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudAssistantStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudAssistantStatus(request)
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

// DescribeCloudAssistantStatusWithCallback invokes the ecs.DescribeCloudAssistantStatus API asynchronously
func (client *Client) DescribeCloudAssistantStatusWithCallback(request *DescribeCloudAssistantStatusRequest, callback func(response *DescribeCloudAssistantStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudAssistantStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudAssistantStatus(request)
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

// DescribeCloudAssistantStatusRequest is the request struct for api DescribeCloudAssistantStatus
type DescribeCloudAssistantStatusRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// DescribeCloudAssistantStatusResponse is the response struct for api DescribeCloudAssistantStatus
type DescribeCloudAssistantStatusResponse struct {
	*responses.BaseResponse
	RequestId                       string                          `json:"RequestId" xml:"RequestId"`
	InstanceCloudAssistantStatusSet InstanceCloudAssistantStatusSet `json:"InstanceCloudAssistantStatusSet" xml:"InstanceCloudAssistantStatusSet"`
}

// CreateDescribeCloudAssistantStatusRequest creates a request to invoke DescribeCloudAssistantStatus API
func CreateDescribeCloudAssistantStatusRequest() (request *DescribeCloudAssistantStatusRequest) {
	request = &DescribeCloudAssistantStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeCloudAssistantStatus", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudAssistantStatusResponse creates a response to parse from DescribeCloudAssistantStatus response
func CreateDescribeCloudAssistantStatusResponse() (response *DescribeCloudAssistantStatusResponse) {
	response = &DescribeCloudAssistantStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
