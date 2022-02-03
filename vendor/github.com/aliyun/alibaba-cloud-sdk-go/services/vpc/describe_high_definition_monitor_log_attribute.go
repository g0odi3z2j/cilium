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

// DescribeHighDefinitionMonitorLogAttribute invokes the vpc.DescribeHighDefinitionMonitorLogAttribute API synchronously
func (client *Client) DescribeHighDefinitionMonitorLogAttribute(request *DescribeHighDefinitionMonitorLogAttributeRequest) (response *DescribeHighDefinitionMonitorLogAttributeResponse, err error) {
	response = CreateDescribeHighDefinitionMonitorLogAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHighDefinitionMonitorLogAttributeWithChan invokes the vpc.DescribeHighDefinitionMonitorLogAttribute API asynchronously
func (client *Client) DescribeHighDefinitionMonitorLogAttributeWithChan(request *DescribeHighDefinitionMonitorLogAttributeRequest) (<-chan *DescribeHighDefinitionMonitorLogAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeHighDefinitionMonitorLogAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHighDefinitionMonitorLogAttribute(request)
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

// DescribeHighDefinitionMonitorLogAttributeWithCallback invokes the vpc.DescribeHighDefinitionMonitorLogAttribute API asynchronously
func (client *Client) DescribeHighDefinitionMonitorLogAttributeWithCallback(request *DescribeHighDefinitionMonitorLogAttributeRequest, callback func(response *DescribeHighDefinitionMonitorLogAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHighDefinitionMonitorLogAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeHighDefinitionMonitorLogAttribute(request)
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

// DescribeHighDefinitionMonitorLogAttributeRequest is the request struct for api DescribeHighDefinitionMonitorLogAttribute
type DescribeHighDefinitionMonitorLogAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeHighDefinitionMonitorLogAttributeResponse is the response struct for api DescribeHighDefinitionMonitorLogAttribute
type DescribeHighDefinitionMonitorLogAttributeResponse struct {
	*responses.BaseResponse
	LogProject   string `json:"LogProject" xml:"LogProject"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	LogStore     string `json:"LogStore" xml:"LogStore"`
	Success      string `json:"Success" xml:"Success"`
	InstanceId   string `json:"InstanceId" xml:"InstanceId"`
	InstanceType string `json:"InstanceType" xml:"InstanceType"`
}

// CreateDescribeHighDefinitionMonitorLogAttributeRequest creates a request to invoke DescribeHighDefinitionMonitorLogAttribute API
func CreateDescribeHighDefinitionMonitorLogAttributeRequest() (request *DescribeHighDefinitionMonitorLogAttributeRequest) {
	request = &DescribeHighDefinitionMonitorLogAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeHighDefinitionMonitorLogAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHighDefinitionMonitorLogAttributeResponse creates a response to parse from DescribeHighDefinitionMonitorLogAttribute response
func CreateDescribeHighDefinitionMonitorLogAttributeResponse() (response *DescribeHighDefinitionMonitorLogAttributeResponse) {
	response = &DescribeHighDefinitionMonitorLogAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
