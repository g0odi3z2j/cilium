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

// CreateDefaultVSwitch invokes the vpc.CreateDefaultVSwitch API synchronously
func (client *Client) CreateDefaultVSwitch(request *CreateDefaultVSwitchRequest) (response *CreateDefaultVSwitchResponse, err error) {
	response = CreateCreateDefaultVSwitchResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDefaultVSwitchWithChan invokes the vpc.CreateDefaultVSwitch API asynchronously
func (client *Client) CreateDefaultVSwitchWithChan(request *CreateDefaultVSwitchRequest) (<-chan *CreateDefaultVSwitchResponse, <-chan error) {
	responseChan := make(chan *CreateDefaultVSwitchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDefaultVSwitch(request)
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

// CreateDefaultVSwitchWithCallback invokes the vpc.CreateDefaultVSwitch API asynchronously
func (client *Client) CreateDefaultVSwitchWithCallback(request *CreateDefaultVSwitchRequest, callback func(response *CreateDefaultVSwitchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDefaultVSwitchResponse
		var err error
		defer close(result)
		response, err = client.CreateDefaultVSwitch(request)
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

// CreateDefaultVSwitchRequest is the request struct for api CreateDefaultVSwitch
type CreateDefaultVSwitchRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Ipv6CidrBlock        requests.Integer `position:"Query" name:"Ipv6CidrBlock"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// CreateDefaultVSwitchResponse is the response struct for api CreateDefaultVSwitch
type CreateDefaultVSwitchResponse struct {
	*responses.BaseResponse
	VSwitchId string `json:"VSwitchId" xml:"VSwitchId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDefaultVSwitchRequest creates a request to invoke CreateDefaultVSwitch API
func CreateCreateDefaultVSwitchRequest() (request *CreateDefaultVSwitchRequest) {
	request = &CreateDefaultVSwitchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateDefaultVSwitch", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDefaultVSwitchResponse creates a response to parse from CreateDefaultVSwitch response
func CreateCreateDefaultVSwitchResponse() (response *CreateDefaultVSwitchResponse) {
	response = &CreateDefaultVSwitchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
