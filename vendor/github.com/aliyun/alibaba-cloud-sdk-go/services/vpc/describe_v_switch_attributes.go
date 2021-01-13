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

// DescribeVSwitchAttributes invokes the vpc.DescribeVSwitchAttributes API synchronously
func (client *Client) DescribeVSwitchAttributes(request *DescribeVSwitchAttributesRequest) (response *DescribeVSwitchAttributesResponse, err error) {
	response = CreateDescribeVSwitchAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVSwitchAttributesWithChan invokes the vpc.DescribeVSwitchAttributes API asynchronously
func (client *Client) DescribeVSwitchAttributesWithChan(request *DescribeVSwitchAttributesRequest) (<-chan *DescribeVSwitchAttributesResponse, <-chan error) {
	responseChan := make(chan *DescribeVSwitchAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVSwitchAttributes(request)
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

// DescribeVSwitchAttributesWithCallback invokes the vpc.DescribeVSwitchAttributes API asynchronously
func (client *Client) DescribeVSwitchAttributesWithCallback(request *DescribeVSwitchAttributesRequest, callback func(response *DescribeVSwitchAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVSwitchAttributesResponse
		var err error
		defer close(result)
		response, err = client.DescribeVSwitchAttributes(request)
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

// DescribeVSwitchAttributesRequest is the request struct for api DescribeVSwitchAttributes
type DescribeVSwitchAttributesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
}

// DescribeVSwitchAttributesResponse is the response struct for api DescribeVSwitchAttributes
type DescribeVSwitchAttributesResponse struct {
	*responses.BaseResponse
	RequestId               string                                    `json:"RequestId" xml:"RequestId"`
	VSwitchId               string                                    `json:"VSwitchId" xml:"VSwitchId"`
	VpcId                   string                                    `json:"VpcId" xml:"VpcId"`
	Status                  string                                    `json:"Status" xml:"Status"`
	CidrBlock               string                                    `json:"CidrBlock" xml:"CidrBlock"`
	Ipv6CidrBlock           string                                    `json:"Ipv6CidrBlock" xml:"Ipv6CidrBlock"`
	ZoneId                  string                                    `json:"ZoneId" xml:"ZoneId"`
	AvailableIpAddressCount int64                                     `json:"AvailableIpAddressCount" xml:"AvailableIpAddressCount"`
	Description             string                                    `json:"Description" xml:"Description"`
	VSwitchName             string                                    `json:"VSwitchName" xml:"VSwitchName"`
	CreationTime            string                                    `json:"CreationTime" xml:"CreationTime"`
	IsDefault               bool                                      `json:"IsDefault" xml:"IsDefault"`
	ResourceGroupId         string                                    `json:"ResourceGroupId" xml:"ResourceGroupId"`
	NetworkAclId            string                                    `json:"NetworkAclId" xml:"NetworkAclId"`
	OwnerId                 int64                                     `json:"OwnerId" xml:"OwnerId"`
	ShareType               string                                    `json:"ShareType" xml:"ShareType"`
	RouteTable              RouteTable                                `json:"RouteTable" xml:"RouteTable"`
	CloudResources          CloudResourcesInDescribeVSwitchAttributes `json:"CloudResources" xml:"CloudResources"`
}

// CreateDescribeVSwitchAttributesRequest creates a request to invoke DescribeVSwitchAttributes API
func CreateDescribeVSwitchAttributesRequest() (request *DescribeVSwitchAttributesRequest) {
	request = &DescribeVSwitchAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVSwitchAttributes", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVSwitchAttributesResponse creates a response to parse from DescribeVSwitchAttributes response
func CreateDescribeVSwitchAttributesResponse() (response *DescribeVSwitchAttributesResponse) {
	response = &DescribeVSwitchAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
