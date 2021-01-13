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

// DetachDisk invokes the ecs.DetachDisk API synchronously
func (client *Client) DetachDisk(request *DetachDiskRequest) (response *DetachDiskResponse, err error) {
	response = CreateDetachDiskResponse()
	err = client.DoAction(request, response)
	return
}

// DetachDiskWithChan invokes the ecs.DetachDisk API asynchronously
func (client *Client) DetachDiskWithChan(request *DetachDiskRequest) (<-chan *DetachDiskResponse, <-chan error) {
	responseChan := make(chan *DetachDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachDisk(request)
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

// DetachDiskWithCallback invokes the ecs.DetachDisk API asynchronously
func (client *Client) DetachDiskWithCallback(request *DetachDiskRequest, callback func(response *DetachDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachDiskResponse
		var err error
		defer close(result)
		response, err = client.DetachDisk(request)
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

// DetachDiskRequest is the request struct for api DetachDisk
type DetachDiskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DiskId               string           `position:"Query" name:"DiskId"`
	DeleteWithInstance   requests.Boolean `position:"Query" name:"DeleteWithInstance"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DetachDiskResponse is the response struct for api DetachDisk
type DetachDiskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachDiskRequest creates a request to invoke DetachDisk API
func CreateDetachDiskRequest() (request *DetachDiskRequest) {
	request = &DetachDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DetachDisk", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachDiskResponse creates a response to parse from DetachDisk response
func CreateDetachDiskResponse() (response *DetachDiskResponse) {
	response = &DetachDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
