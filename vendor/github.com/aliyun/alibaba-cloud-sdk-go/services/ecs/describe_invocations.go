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

// DescribeInvocations invokes the ecs.DescribeInvocations API synchronously
func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (response *DescribeInvocationsResponse, err error) {
	response = CreateDescribeInvocationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInvocationsWithChan invokes the ecs.DescribeInvocations API asynchronously
func (client *Client) DescribeInvocationsWithChan(request *DescribeInvocationsRequest) (<-chan *DescribeInvocationsResponse, <-chan error) {
	responseChan := make(chan *DescribeInvocationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInvocations(request)
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

// DescribeInvocationsWithCallback invokes the ecs.DescribeInvocations API asynchronously
func (client *Client) DescribeInvocationsWithCallback(request *DescribeInvocationsRequest, callback func(response *DescribeInvocationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInvocationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeInvocations(request)
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

// DescribeInvocationsRequest is the request struct for api DescribeInvocations
type DescribeInvocationsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer          `position:"Query" name:"ResourceOwnerId"`
	InvokeStatus         string                    `position:"Query" name:"InvokeStatus"`
	IncludeOutput        requests.Boolean          `position:"Query" name:"IncludeOutput"`
	CommandId            string                    `position:"Query" name:"CommandId"`
	PageNumber           requests.Integer          `position:"Query" name:"PageNumber"`
	ResourceGroupId      string                    `position:"Query" name:"ResourceGroupId"`
	ContentEncoding      string                    `position:"Query" name:"ContentEncoding"`
	RepeatMode           string                    `position:"Query" name:"RepeatMode"`
	PageSize             requests.Integer          `position:"Query" name:"PageSize"`
	Tag                  *[]DescribeInvocationsTag `position:"Query" name:"Tag"  type:"Repeated"`
	InvokeId             string                    `position:"Query" name:"InvokeId"`
	Timed                requests.Boolean          `position:"Query" name:"Timed"`
	CommandName          string                    `position:"Query" name:"CommandName"`
	ResourceOwnerAccount string                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                    `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer          `position:"Query" name:"OwnerId"`
	CommandType          string                    `position:"Query" name:"CommandType"`
	InstanceId           string                    `position:"Query" name:"InstanceId"`
}

// DescribeInvocationsTag is a repeated param struct in DescribeInvocationsRequest
type DescribeInvocationsTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeInvocationsResponse is the response struct for api DescribeInvocations
type DescribeInvocationsResponse struct {
	*responses.BaseResponse
	PageSize    int64                            `json:"PageSize" xml:"PageSize"`
	RequestId   string                           `json:"RequestId" xml:"RequestId"`
	PageNumber  int64                            `json:"PageNumber" xml:"PageNumber"`
	TotalCount  int64                            `json:"TotalCount" xml:"TotalCount"`
	Invocations InvocationsInDescribeInvocations `json:"Invocations" xml:"Invocations"`
}

// CreateDescribeInvocationsRequest creates a request to invoke DescribeInvocations API
func CreateDescribeInvocationsRequest() (request *DescribeInvocationsRequest) {
	request = &DescribeInvocationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInvocations", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInvocationsResponse creates a response to parse from DescribeInvocations response
func CreateDescribeInvocationsResponse() (response *DescribeInvocationsResponse) {
	response = &DescribeInvocationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
