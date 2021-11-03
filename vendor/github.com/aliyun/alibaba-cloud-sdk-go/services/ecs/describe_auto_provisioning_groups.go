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

// DescribeAutoProvisioningGroups invokes the ecs.DescribeAutoProvisioningGroups API synchronously
func (client *Client) DescribeAutoProvisioningGroups(request *DescribeAutoProvisioningGroupsRequest) (response *DescribeAutoProvisioningGroupsResponse, err error) {
	response = CreateDescribeAutoProvisioningGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutoProvisioningGroupsWithChan invokes the ecs.DescribeAutoProvisioningGroups API asynchronously
func (client *Client) DescribeAutoProvisioningGroupsWithChan(request *DescribeAutoProvisioningGroupsRequest) (<-chan *DescribeAutoProvisioningGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeAutoProvisioningGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutoProvisioningGroups(request)
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

// DescribeAutoProvisioningGroupsWithCallback invokes the ecs.DescribeAutoProvisioningGroups API asynchronously
func (client *Client) DescribeAutoProvisioningGroupsWithCallback(request *DescribeAutoProvisioningGroupsRequest, callback func(response *DescribeAutoProvisioningGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutoProvisioningGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutoProvisioningGroups(request)
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

// DescribeAutoProvisioningGroupsRequest is the request struct for api DescribeAutoProvisioningGroups
type DescribeAutoProvisioningGroupsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId             requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber                  requests.Integer `position:"Query" name:"PageNumber"`
	PageSize                    requests.Integer `position:"Query" name:"PageSize"`
	AutoProvisioningGroupStatus *[]string        `position:"Query" name:"AutoProvisioningGroupStatus"  type:"Repeated"`
	ResourceOwnerAccount        string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string           `position:"Query" name:"OwnerAccount"`
	OwnerId                     requests.Integer `position:"Query" name:"OwnerId"`
	AutoProvisioningGroupId     *[]string        `position:"Query" name:"AutoProvisioningGroupId"  type:"Repeated"`
	AutoProvisioningGroupName   string           `position:"Query" name:"AutoProvisioningGroupName"`
}

// DescribeAutoProvisioningGroupsResponse is the response struct for api DescribeAutoProvisioningGroups
type DescribeAutoProvisioningGroupsResponse struct {
	*responses.BaseResponse
	PageSize               int                    `json:"PageSize" xml:"PageSize"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	PageNumber             int                    `json:"PageNumber" xml:"PageNumber"`
	TotalCount             int                    `json:"TotalCount" xml:"TotalCount"`
	AutoProvisioningGroups AutoProvisioningGroups `json:"AutoProvisioningGroups" xml:"AutoProvisioningGroups"`
}

// CreateDescribeAutoProvisioningGroupsRequest creates a request to invoke DescribeAutoProvisioningGroups API
func CreateDescribeAutoProvisioningGroupsRequest() (request *DescribeAutoProvisioningGroupsRequest) {
	request = &DescribeAutoProvisioningGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAutoProvisioningGroups", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAutoProvisioningGroupsResponse creates a response to parse from DescribeAutoProvisioningGroups response
func CreateDescribeAutoProvisioningGroupsResponse() (response *DescribeAutoProvisioningGroupsResponse) {
	response = &DescribeAutoProvisioningGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
