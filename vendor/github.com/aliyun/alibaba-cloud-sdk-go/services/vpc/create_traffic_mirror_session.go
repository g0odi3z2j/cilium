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

// CreateTrafficMirrorSession invokes the vpc.CreateTrafficMirrorSession API synchronously
func (client *Client) CreateTrafficMirrorSession(request *CreateTrafficMirrorSessionRequest) (response *CreateTrafficMirrorSessionResponse, err error) {
	response = CreateCreateTrafficMirrorSessionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTrafficMirrorSessionWithChan invokes the vpc.CreateTrafficMirrorSession API asynchronously
func (client *Client) CreateTrafficMirrorSessionWithChan(request *CreateTrafficMirrorSessionRequest) (<-chan *CreateTrafficMirrorSessionResponse, <-chan error) {
	responseChan := make(chan *CreateTrafficMirrorSessionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTrafficMirrorSession(request)
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

// CreateTrafficMirrorSessionWithCallback invokes the vpc.CreateTrafficMirrorSession API asynchronously
func (client *Client) CreateTrafficMirrorSessionWithCallback(request *CreateTrafficMirrorSessionRequest, callback func(response *CreateTrafficMirrorSessionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTrafficMirrorSessionResponse
		var err error
		defer close(result)
		response, err = client.CreateTrafficMirrorSession(request)
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

// CreateTrafficMirrorSessionRequest is the request struct for api CreateTrafficMirrorSession
type CreateTrafficMirrorSessionRequest struct {
	*requests.RpcRequest
	TrafficMirrorTargetType         string           `position:"Query" name:"TrafficMirrorTargetType"`
	ResourceOwnerId                 requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                     string           `position:"Query" name:"ClientToken"`
	Enabled                         requests.Boolean `position:"Query" name:"Enabled"`
	ResourceGroupId                 string           `position:"Query" name:"ResourceGroupId"`
	TrafficMirrorSessionName        string           `position:"Query" name:"TrafficMirrorSessionName"`
	TrafficMirrorSessionDescription string           `position:"Query" name:"TrafficMirrorSessionDescription"`
	TrafficMirrorSourceIds          *[]string        `position:"Query" name:"TrafficMirrorSourceIds"  type:"Repeated"`
	DryRun                          requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount            string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                    string           `position:"Query" name:"OwnerAccount"`
	Priority                        requests.Integer `position:"Query" name:"Priority"`
	OwnerId                         requests.Integer `position:"Query" name:"OwnerId"`
	TrafficMirrorTargetId           string           `position:"Query" name:"TrafficMirrorTargetId"`
	TrafficMirrorFilterId           string           `position:"Query" name:"TrafficMirrorFilterId"`
	PacketLength                    requests.Integer `position:"Query" name:"PacketLength"`
	VirtualNetworkId                requests.Integer `position:"Query" name:"VirtualNetworkId"`
}

// CreateTrafficMirrorSessionResponse is the response struct for api CreateTrafficMirrorSession
type CreateTrafficMirrorSessionResponse struct {
	*responses.BaseResponse
	TrafficMirrorSessionId string `json:"TrafficMirrorSessionId" xml:"TrafficMirrorSessionId"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
	ResourceGroupId        string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

// CreateCreateTrafficMirrorSessionRequest creates a request to invoke CreateTrafficMirrorSession API
func CreateCreateTrafficMirrorSessionRequest() (request *CreateTrafficMirrorSessionRequest) {
	request = &CreateTrafficMirrorSessionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateTrafficMirrorSession", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTrafficMirrorSessionResponse creates a response to parse from CreateTrafficMirrorSession response
func CreateCreateTrafficMirrorSessionResponse() (response *CreateTrafficMirrorSessionResponse) {
	response = &CreateTrafficMirrorSessionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
