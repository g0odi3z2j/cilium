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

// DescribeVpnRouteEntries invokes the vpc.DescribeVpnRouteEntries API synchronously
func (client *Client) DescribeVpnRouteEntries(request *DescribeVpnRouteEntriesRequest) (response *DescribeVpnRouteEntriesResponse, err error) {
	response = CreateDescribeVpnRouteEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVpnRouteEntriesWithChan invokes the vpc.DescribeVpnRouteEntries API asynchronously
func (client *Client) DescribeVpnRouteEntriesWithChan(request *DescribeVpnRouteEntriesRequest) (<-chan *DescribeVpnRouteEntriesResponse, <-chan error) {
	responseChan := make(chan *DescribeVpnRouteEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpnRouteEntries(request)
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

// DescribeVpnRouteEntriesWithCallback invokes the vpc.DescribeVpnRouteEntries API asynchronously
func (client *Client) DescribeVpnRouteEntriesWithCallback(request *DescribeVpnRouteEntriesRequest, callback func(response *DescribeVpnRouteEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpnRouteEntriesResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpnRouteEntries(request)
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

// DescribeVpnRouteEntriesRequest is the request struct for api DescribeVpnRouteEntries
type DescribeVpnRouteEntriesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	RouteEntryType       string           `position:"Query" name:"RouteEntryType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string           `position:"Query" name:"VpnGatewayId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVpnRouteEntriesResponse is the response struct for api DescribeVpnRouteEntries
type DescribeVpnRouteEntriesResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	TotalCount      int             `json:"TotalCount" xml:"TotalCount"`
	PageNumber      int             `json:"PageNumber" xml:"PageNumber"`
	PageSize        int             `json:"PageSize" xml:"PageSize"`
	VpnRouteEntries VpnRouteEntries `json:"VpnRouteEntries" xml:"VpnRouteEntries"`
}

// CreateDescribeVpnRouteEntriesRequest creates a request to invoke DescribeVpnRouteEntries API
func CreateDescribeVpnRouteEntriesRequest() (request *DescribeVpnRouteEntriesRequest) {
	request = &DescribeVpnRouteEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnRouteEntries", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVpnRouteEntriesResponse creates a response to parse from DescribeVpnRouteEntries response
func CreateDescribeVpnRouteEntriesResponse() (response *DescribeVpnRouteEntriesResponse) {
	response = &DescribeVpnRouteEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
