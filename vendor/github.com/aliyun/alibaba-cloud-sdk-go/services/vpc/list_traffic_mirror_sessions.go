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

// ListTrafficMirrorSessions invokes the vpc.ListTrafficMirrorSessions API synchronously
func (client *Client) ListTrafficMirrorSessions(request *ListTrafficMirrorSessionsRequest) (response *ListTrafficMirrorSessionsResponse, err error) {
	response = CreateListTrafficMirrorSessionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTrafficMirrorSessionsWithChan invokes the vpc.ListTrafficMirrorSessions API asynchronously
func (client *Client) ListTrafficMirrorSessionsWithChan(request *ListTrafficMirrorSessionsRequest) (<-chan *ListTrafficMirrorSessionsResponse, <-chan error) {
	responseChan := make(chan *ListTrafficMirrorSessionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTrafficMirrorSessions(request)
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

// ListTrafficMirrorSessionsWithCallback invokes the vpc.ListTrafficMirrorSessions API asynchronously
func (client *Client) ListTrafficMirrorSessionsWithCallback(request *ListTrafficMirrorSessionsRequest, callback func(response *ListTrafficMirrorSessionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTrafficMirrorSessionsResponse
		var err error
		defer close(result)
		response, err = client.ListTrafficMirrorSessions(request)
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

// ListTrafficMirrorSessionsRequest is the request struct for api ListTrafficMirrorSessions
type ListTrafficMirrorSessionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId          requests.Integer                 `position:"Query" name:"ResourceOwnerId"`
	TrafficMirrorSourceId    string                           `position:"Query" name:"TrafficMirrorSourceId"`
	Enabled                  requests.Boolean                 `position:"Query" name:"Enabled"`
	ResourceGroupId          string                           `position:"Query" name:"ResourceGroupId"`
	TrafficMirrorSessionName string                           `position:"Query" name:"TrafficMirrorSessionName"`
	NextToken                string                           `position:"Query" name:"NextToken"`
	TrafficMirrorSessionIds  *[]string                        `position:"Query" name:"TrafficMirrorSessionIds"  type:"Repeated"`
	ResourceOwnerAccount     string                           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string                           `position:"Query" name:"OwnerAccount"`
	Priority                 requests.Integer                 `position:"Query" name:"Priority"`
	OwnerId                  requests.Integer                 `position:"Query" name:"OwnerId"`
	TrafficMirrorTargetId    string                           `position:"Query" name:"TrafficMirrorTargetId"`
	TrafficMirrorFilterId    string                           `position:"Query" name:"TrafficMirrorFilterId"`
	Tags                     *[]ListTrafficMirrorSessionsTags `position:"Query" name:"Tags"  type:"Repeated"`
	MaxResults               requests.Integer                 `position:"Query" name:"MaxResults"`
	VirtualNetworkId         requests.Integer                 `position:"Query" name:"VirtualNetworkId"`
}

// ListTrafficMirrorSessionsTags is a repeated param struct in ListTrafficMirrorSessionsRequest
type ListTrafficMirrorSessionsTags struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// ListTrafficMirrorSessionsResponse is the response struct for api ListTrafficMirrorSessions
type ListTrafficMirrorSessionsResponse struct {
	*responses.BaseResponse
	NextToken             string                 `json:"NextToken" xml:"NextToken"`
	RequestId             string                 `json:"RequestId" xml:"RequestId"`
	TotalCount            string                 `json:"TotalCount" xml:"TotalCount"`
	TrafficMirrorSessions []TrafficMirrorSession `json:"TrafficMirrorSessions" xml:"TrafficMirrorSessions"`
}

// CreateListTrafficMirrorSessionsRequest creates a request to invoke ListTrafficMirrorSessions API
func CreateListTrafficMirrorSessionsRequest() (request *ListTrafficMirrorSessionsRequest) {
	request = &ListTrafficMirrorSessionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ListTrafficMirrorSessions", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTrafficMirrorSessionsResponse creates a response to parse from ListTrafficMirrorSessions response
func CreateListTrafficMirrorSessionsResponse() (response *ListTrafficMirrorSessionsResponse) {
	response = &ListTrafficMirrorSessionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
