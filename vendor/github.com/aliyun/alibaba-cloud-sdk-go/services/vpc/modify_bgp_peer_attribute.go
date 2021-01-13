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

// ModifyBgpPeerAttribute invokes the vpc.ModifyBgpPeerAttribute API synchronously
func (client *Client) ModifyBgpPeerAttribute(request *ModifyBgpPeerAttributeRequest) (response *ModifyBgpPeerAttributeResponse, err error) {
	response = CreateModifyBgpPeerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBgpPeerAttributeWithChan invokes the vpc.ModifyBgpPeerAttribute API asynchronously
func (client *Client) ModifyBgpPeerAttributeWithChan(request *ModifyBgpPeerAttributeRequest) (<-chan *ModifyBgpPeerAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyBgpPeerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBgpPeerAttribute(request)
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

// ModifyBgpPeerAttributeWithCallback invokes the vpc.ModifyBgpPeerAttribute API asynchronously
func (client *Client) ModifyBgpPeerAttributeWithCallback(request *ModifyBgpPeerAttributeRequest, callback func(response *ModifyBgpPeerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBgpPeerAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyBgpPeerAttribute(request)
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

// ModifyBgpPeerAttributeRequest is the request struct for api ModifyBgpPeerAttribute
type ModifyBgpPeerAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	BgpGroupId           string           `position:"Query" name:"BgpGroupId"`
	PeerIpAddress        string           `position:"Query" name:"PeerIpAddress"`
	BfdMultiHop          requests.Integer `position:"Query" name:"BfdMultiHop"`
	EnableBfd            requests.Boolean `position:"Query" name:"EnableBfd"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	BgpPeerId            string           `position:"Query" name:"BgpPeerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyBgpPeerAttributeResponse is the response struct for api ModifyBgpPeerAttribute
type ModifyBgpPeerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyBgpPeerAttributeRequest creates a request to invoke ModifyBgpPeerAttribute API
func CreateModifyBgpPeerAttributeRequest() (request *ModifyBgpPeerAttributeRequest) {
	request = &ModifyBgpPeerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyBgpPeerAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyBgpPeerAttributeResponse creates a response to parse from ModifyBgpPeerAttribute response
func CreateModifyBgpPeerAttributeResponse() (response *ModifyBgpPeerAttributeResponse) {
	response = &ModifyBgpPeerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
