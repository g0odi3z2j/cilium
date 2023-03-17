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

// DescribeIpv6GatewayAttribute invokes the vpc.DescribeIpv6GatewayAttribute API synchronously
func (client *Client) DescribeIpv6GatewayAttribute(request *DescribeIpv6GatewayAttributeRequest) (response *DescribeIpv6GatewayAttributeResponse, err error) {
	response = CreateDescribeIpv6GatewayAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIpv6GatewayAttributeWithChan invokes the vpc.DescribeIpv6GatewayAttribute API asynchronously
func (client *Client) DescribeIpv6GatewayAttributeWithChan(request *DescribeIpv6GatewayAttributeRequest) (<-chan *DescribeIpv6GatewayAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeIpv6GatewayAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIpv6GatewayAttribute(request)
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

// DescribeIpv6GatewayAttributeWithCallback invokes the vpc.DescribeIpv6GatewayAttribute API asynchronously
func (client *Client) DescribeIpv6GatewayAttributeWithCallback(request *DescribeIpv6GatewayAttributeRequest, callback func(response *DescribeIpv6GatewayAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIpv6GatewayAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeIpv6GatewayAttribute(request)
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

// DescribeIpv6GatewayAttributeRequest is the request struct for api DescribeIpv6GatewayAttribute
type DescribeIpv6GatewayAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Ipv6GatewayId        string           `position:"Query" name:"Ipv6GatewayId"`
}

// DescribeIpv6GatewayAttributeResponse is the response struct for api DescribeIpv6GatewayAttribute
type DescribeIpv6GatewayAttributeResponse struct {
	*responses.BaseResponse
	VpcId              string                             `json:"VpcId" xml:"VpcId"`
	Status             string                             `json:"Status" xml:"Status"`
	CreationTime       string                             `json:"CreationTime" xml:"CreationTime"`
	Spec               string                             `json:"Spec" xml:"Spec"`
	RegionId           string                             `json:"RegionId" xml:"RegionId"`
	InstanceChargeType string                             `json:"InstanceChargeType" xml:"InstanceChargeType"`
	RequestId          string                             `json:"RequestId" xml:"RequestId"`
	Ipv6GatewayId      string                             `json:"Ipv6GatewayId" xml:"Ipv6GatewayId"`
	Description        string                             `json:"Description" xml:"Description"`
	ExpiredTime        string                             `json:"ExpiredTime" xml:"ExpiredTime"`
	BusinessStatus     string                             `json:"BusinessStatus" xml:"BusinessStatus"`
	Name               string                             `json:"Name" xml:"Name"`
	ResourceGroupId    string                             `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tags               TagsInDescribeIpv6GatewayAttribute `json:"Tags" xml:"Tags"`
}

// CreateDescribeIpv6GatewayAttributeRequest creates a request to invoke DescribeIpv6GatewayAttribute API
func CreateDescribeIpv6GatewayAttributeRequest() (request *DescribeIpv6GatewayAttributeRequest) {
	request = &DescribeIpv6GatewayAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeIpv6GatewayAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeIpv6GatewayAttributeResponse creates a response to parse from DescribeIpv6GatewayAttribute response
func CreateDescribeIpv6GatewayAttributeResponse() (response *DescribeIpv6GatewayAttributeResponse) {
	response = &DescribeIpv6GatewayAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
