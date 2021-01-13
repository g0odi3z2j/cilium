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

// CopySnapshot invokes the ecs.CopySnapshot API synchronously
func (client *Client) CopySnapshot(request *CopySnapshotRequest) (response *CopySnapshotResponse, err error) {
	response = CreateCopySnapshotResponse()
	err = client.DoAction(request, response)
	return
}

// CopySnapshotWithChan invokes the ecs.CopySnapshot API asynchronously
func (client *Client) CopySnapshotWithChan(request *CopySnapshotRequest) (<-chan *CopySnapshotResponse, <-chan error) {
	responseChan := make(chan *CopySnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopySnapshot(request)
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

// CopySnapshotWithCallback invokes the ecs.CopySnapshot API asynchronously
func (client *Client) CopySnapshotWithCallback(request *CopySnapshotRequest, callback func(response *CopySnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopySnapshotResponse
		var err error
		defer close(result)
		response, err = client.CopySnapshot(request)
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

// CopySnapshotRequest is the request struct for api CopySnapshot
type CopySnapshotRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                requests.Integer   `position:"Query" name:"ResourceOwnerId"`
	SnapshotId                     string             `position:"Query" name:"SnapshotId"`
	DestinationRegionId            string             `position:"Query" name:"DestinationRegionId"`
	ResourceGroupId                string             `position:"Query" name:"ResourceGroupId"`
	Tag                            *[]CopySnapshotTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount           string             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId                        requests.Integer   `position:"Query" name:"OwnerId"`
	DestinationSnapshotName        string             `position:"Query" name:"DestinationSnapshotName"`
	DestinationSnapshotDescription string             `position:"Query" name:"DestinationSnapshotDescription"`
	RetentionDays                  requests.Integer   `position:"Query" name:"RetentionDays"`
}

// CopySnapshotTag is a repeated param struct in CopySnapshotRequest
type CopySnapshotTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CopySnapshotResponse is the response struct for api CopySnapshot
type CopySnapshotResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	SnapshotId string `json:"SnapshotId" xml:"SnapshotId"`
}

// CreateCopySnapshotRequest creates a request to invoke CopySnapshot API
func CreateCopySnapshotRequest() (request *CopySnapshotRequest) {
	request = &CopySnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CopySnapshot", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCopySnapshotResponse creates a response to parse from CopySnapshot response
func CreateCopySnapshotResponse() (response *CopySnapshotResponse) {
	response = &CopySnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
