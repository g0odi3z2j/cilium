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

// ReleaseEipAddress invokes the vpc.ReleaseEipAddress API synchronously
func (client *Client) ReleaseEipAddress(request *ReleaseEipAddressRequest) (response *ReleaseEipAddressResponse, err error) {
	response = CreateReleaseEipAddressResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseEipAddressWithChan invokes the vpc.ReleaseEipAddress API asynchronously
func (client *Client) ReleaseEipAddressWithChan(request *ReleaseEipAddressRequest) (<-chan *ReleaseEipAddressResponse, <-chan error) {
	responseChan := make(chan *ReleaseEipAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseEipAddress(request)
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

// ReleaseEipAddressWithCallback invokes the vpc.ReleaseEipAddress API asynchronously
func (client *Client) ReleaseEipAddressWithCallback(request *ReleaseEipAddressRequest, callback func(response *ReleaseEipAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseEipAddressResponse
		var err error
		defer close(result)
		response, err = client.ReleaseEipAddress(request)
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

// ReleaseEipAddressRequest is the request struct for api ReleaseEipAddress
type ReleaseEipAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AllocationId         string           `position:"Query" name:"AllocationId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ReleaseEipAddressResponse is the response struct for api ReleaseEipAddress
type ReleaseEipAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseEipAddressRequest creates a request to invoke ReleaseEipAddress API
func CreateReleaseEipAddressRequest() (request *ReleaseEipAddressRequest) {
	request = &ReleaseEipAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ReleaseEipAddress", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleaseEipAddressResponse creates a response to parse from ReleaseEipAddress response
func CreateReleaseEipAddressResponse() (response *ReleaseEipAddressResponse) {
	response = &ReleaseEipAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
