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

// DescribeBgpGroups invokes the vpc.DescribeBgpGroups API synchronously
func (client *Client) DescribeBgpGroups(request *DescribeBgpGroupsRequest) (response *DescribeBgpGroupsResponse, err error) {
	response = CreateDescribeBgpGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBgpGroupsWithChan invokes the vpc.DescribeBgpGroups API asynchronously
func (client *Client) DescribeBgpGroupsWithChan(request *DescribeBgpGroupsRequest) (<-chan *DescribeBgpGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeBgpGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBgpGroups(request)
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

// DescribeBgpGroupsWithCallback invokes the vpc.DescribeBgpGroups API asynchronously
func (client *Client) DescribeBgpGroupsWithCallback(request *DescribeBgpGroupsRequest, callback func(response *DescribeBgpGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBgpGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeBgpGroups(request)
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

// DescribeBgpGroupsRequest is the request struct for api DescribeBgpGroups
type DescribeBgpGroupsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BgpGroupId           string           `position:"Query" name:"BgpGroupId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	IsDefault            requests.Boolean `position:"Query" name:"IsDefault"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouterId             string           `position:"Query" name:"RouterId"`
}

// DescribeBgpGroupsResponse is the response struct for api DescribeBgpGroups
type DescribeBgpGroupsResponse struct {
	*responses.BaseResponse
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	BgpGroups  BgpGroups `json:"BgpGroups" xml:"BgpGroups"`
}

// CreateDescribeBgpGroupsRequest creates a request to invoke DescribeBgpGroups API
func CreateDescribeBgpGroupsRequest() (request *DescribeBgpGroupsRequest) {
	request = &DescribeBgpGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeBgpGroups", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBgpGroupsResponse creates a response to parse from DescribeBgpGroups response
func CreateDescribeBgpGroupsResponse() (response *DescribeBgpGroupsResponse) {
	response = &DescribeBgpGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
