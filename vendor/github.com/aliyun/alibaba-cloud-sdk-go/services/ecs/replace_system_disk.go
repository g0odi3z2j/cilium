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

// ReplaceSystemDisk invokes the ecs.ReplaceSystemDisk API synchronously
func (client *Client) ReplaceSystemDisk(request *ReplaceSystemDiskRequest) (response *ReplaceSystemDiskResponse, err error) {
	response = CreateReplaceSystemDiskResponse()
	err = client.DoAction(request, response)
	return
}

// ReplaceSystemDiskWithChan invokes the ecs.ReplaceSystemDisk API asynchronously
func (client *Client) ReplaceSystemDiskWithChan(request *ReplaceSystemDiskRequest) (<-chan *ReplaceSystemDiskResponse, <-chan error) {
	responseChan := make(chan *ReplaceSystemDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReplaceSystemDisk(request)
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

// ReplaceSystemDiskWithCallback invokes the ecs.ReplaceSystemDisk API asynchronously
func (client *Client) ReplaceSystemDiskWithCallback(request *ReplaceSystemDiskRequest, callback func(response *ReplaceSystemDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReplaceSystemDiskResponse
		var err error
		defer close(result)
		response, err = client.ReplaceSystemDisk(request)
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

// ReplaceSystemDiskRequest is the request struct for api ReplaceSystemDisk
type ReplaceSystemDiskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId             requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ImageId                     string           `position:"Query" name:"ImageId"`
	ClientToken                 string           `position:"Query" name:"ClientToken"`
	SecurityEnhancementStrategy string           `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string           `position:"Query" name:"KeyPairName"`
	Platform                    string           `position:"Query" name:"Platform"`
	Password                    string           `position:"Query" name:"Password"`
	PasswordInherit             requests.Boolean `position:"Query" name:"PasswordInherit"`
	DiskId                      string           `position:"Query" name:"DiskId"`
	Architecture                string           `position:"Query" name:"Architecture"`
	ResourceOwnerAccount        string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string           `position:"Query" name:"OwnerAccount"`
	OwnerId                     requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId                  string           `position:"Query" name:"InstanceId"`
	SystemDiskSize              requests.Integer `position:"Query" name:"SystemDisk.Size"`
	UseAdditionalService        requests.Boolean `position:"Query" name:"UseAdditionalService"`
}

// ReplaceSystemDiskResponse is the response struct for api ReplaceSystemDisk
type ReplaceSystemDiskResponse struct {
	*responses.BaseResponse
	DiskId    string `json:"DiskId" xml:"DiskId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReplaceSystemDiskRequest creates a request to invoke ReplaceSystemDisk API
func CreateReplaceSystemDiskRequest() (request *ReplaceSystemDiskRequest) {
	request = &ReplaceSystemDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ReplaceSystemDisk", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReplaceSystemDiskResponse creates a response to parse from ReplaceSystemDisk response
func CreateReplaceSystemDiskResponse() (response *ReplaceSystemDiskResponse) {
	response = &ReplaceSystemDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
