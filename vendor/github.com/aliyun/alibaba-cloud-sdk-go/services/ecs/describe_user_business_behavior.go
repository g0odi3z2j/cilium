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

// DescribeUserBusinessBehavior invokes the ecs.DescribeUserBusinessBehavior API synchronously
func (client *Client) DescribeUserBusinessBehavior(request *DescribeUserBusinessBehaviorRequest) (response *DescribeUserBusinessBehaviorResponse, err error) {
	response = CreateDescribeUserBusinessBehaviorResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserBusinessBehaviorWithChan invokes the ecs.DescribeUserBusinessBehavior API asynchronously
func (client *Client) DescribeUserBusinessBehaviorWithChan(request *DescribeUserBusinessBehaviorRequest) (<-chan *DescribeUserBusinessBehaviorResponse, <-chan error) {
	responseChan := make(chan *DescribeUserBusinessBehaviorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserBusinessBehavior(request)
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

// DescribeUserBusinessBehaviorWithCallback invokes the ecs.DescribeUserBusinessBehavior API asynchronously
func (client *Client) DescribeUserBusinessBehaviorWithCallback(request *DescribeUserBusinessBehaviorRequest, callback func(response *DescribeUserBusinessBehaviorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserBusinessBehaviorResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserBusinessBehavior(request)
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

// DescribeUserBusinessBehaviorRequest is the request struct for api DescribeUserBusinessBehavior
type DescribeUserBusinessBehaviorRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	StatusKey            string           `position:"Query" name:"statusKey"`
}

// DescribeUserBusinessBehaviorResponse is the response struct for api DescribeUserBusinessBehavior
type DescribeUserBusinessBehaviorResponse struct {
	*responses.BaseResponse
	StatusValue string `json:"StatusValue" xml:"StatusValue"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeUserBusinessBehaviorRequest creates a request to invoke DescribeUserBusinessBehavior API
func CreateDescribeUserBusinessBehaviorRequest() (request *DescribeUserBusinessBehaviorRequest) {
	request = &DescribeUserBusinessBehaviorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeUserBusinessBehavior", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUserBusinessBehaviorResponse creates a response to parse from DescribeUserBusinessBehavior response
func CreateDescribeUserBusinessBehaviorResponse() (response *DescribeUserBusinessBehaviorResponse) {
	response = &DescribeUserBusinessBehaviorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
