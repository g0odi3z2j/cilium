package vpc

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

// DescribeInstanceAutoRenewAttribute invokes the vpc.DescribeInstanceAutoRenewAttribute API synchronously
func (client *Client) DescribeInstanceAutoRenewAttribute(request *DescribeInstanceAutoRenewAttributeRequest) (response *DescribeInstanceAutoRenewAttributeResponse, err error) {
	response = CreateDescribeInstanceAutoRenewAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceAutoRenewAttributeWithChan invokes the vpc.DescribeInstanceAutoRenewAttribute API asynchronously
func (client *Client) DescribeInstanceAutoRenewAttributeWithChan(request *DescribeInstanceAutoRenewAttributeRequest) (<-chan *DescribeInstanceAutoRenewAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceAutoRenewAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceAutoRenewAttribute(request)
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

// DescribeInstanceAutoRenewAttributeWithCallback invokes the vpc.DescribeInstanceAutoRenewAttribute API asynchronously
func (client *Client) DescribeInstanceAutoRenewAttributeWithCallback(request *DescribeInstanceAutoRenewAttributeRequest, callback func(response *DescribeInstanceAutoRenewAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceAutoRenewAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceAutoRenewAttribute(request)
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

// DescribeInstanceAutoRenewAttributeRequest is the request struct for api DescribeInstanceAutoRenewAttribute
type DescribeInstanceAutoRenewAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	RenewalStatus        string           `position:"Query" name:"RenewalStatus"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeInstanceAutoRenewAttributeResponse is the response struct for api DescribeInstanceAutoRenewAttribute
type DescribeInstanceAutoRenewAttributeResponse struct {
	*responses.BaseResponse
	PageNumber              string                  `json:"PageNumber" xml:"PageNumber"`
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	PageSize                string                  `json:"PageSize" xml:"PageSize"`
	TotalCount              string                  `json:"TotalCount" xml:"TotalCount"`
	InstanceRenewAttributes InstanceRenewAttributes `json:"InstanceRenewAttributes" xml:"InstanceRenewAttributes"`
}

// CreateDescribeInstanceAutoRenewAttributeRequest creates a request to invoke DescribeInstanceAutoRenewAttribute API
func CreateDescribeInstanceAutoRenewAttributeRequest() (request *DescribeInstanceAutoRenewAttributeRequest) {
	request = &DescribeInstanceAutoRenewAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeInstanceAutoRenewAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceAutoRenewAttributeResponse creates a response to parse from DescribeInstanceAutoRenewAttribute response
func CreateDescribeInstanceAutoRenewAttributeResponse() (response *DescribeInstanceAutoRenewAttributeResponse) {
	response = &DescribeInstanceAutoRenewAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
