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

// CreateIpv4Gateway invokes the vpc.CreateIpv4Gateway API synchronously
func (client *Client) CreateIpv4Gateway(request *CreateIpv4GatewayRequest) (response *CreateIpv4GatewayResponse, err error) {
	response = CreateCreateIpv4GatewayResponse()
	err = client.DoAction(request, response)
	return
}

// CreateIpv4GatewayWithChan invokes the vpc.CreateIpv4Gateway API asynchronously
func (client *Client) CreateIpv4GatewayWithChan(request *CreateIpv4GatewayRequest) (<-chan *CreateIpv4GatewayResponse, <-chan error) {
	responseChan := make(chan *CreateIpv4GatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateIpv4Gateway(request)
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

// CreateIpv4GatewayWithCallback invokes the vpc.CreateIpv4Gateway API asynchronously
func (client *Client) CreateIpv4GatewayWithCallback(request *CreateIpv4GatewayRequest, callback func(response *CreateIpv4GatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateIpv4GatewayResponse
		var err error
		defer close(result)
		response, err = client.CreateIpv4Gateway(request)
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

// CreateIpv4GatewayRequest is the request struct for api CreateIpv4Gateway
type CreateIpv4GatewayRequest struct {
	*requests.RpcRequest
	Ipv4GatewayDescription string                  `position:"Query" name:"Ipv4GatewayDescription"`
	ResourceOwnerId        requests.Integer        `position:"Query" name:"ResourceOwnerId"`
	ClientToken            string                  `position:"Query" name:"ClientToken"`
	Ipv4GatewayName        string                  `position:"Query" name:"Ipv4GatewayName"`
	ResourceGroupId        string                  `position:"Query" name:"ResourceGroupId"`
	Tag                    *[]CreateIpv4GatewayTag `position:"Query" name:"Tag"  type:"Repeated"`
	DryRun                 requests.Boolean        `position:"Query" name:"DryRun"`
	ResourceOwnerAccount   string                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string                  `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer        `position:"Query" name:"OwnerId"`
	VpcId                  string                  `position:"Query" name:"VpcId"`
}

// CreateIpv4GatewayTag is a repeated param struct in CreateIpv4GatewayRequest
type CreateIpv4GatewayTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateIpv4GatewayResponse is the response struct for api CreateIpv4Gateway
type CreateIpv4GatewayResponse struct {
	*responses.BaseResponse
	Ipv4GatewayId   string `json:"Ipv4GatewayId" xml:"Ipv4GatewayId"`
	ResourceGroupId string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateIpv4GatewayRequest creates a request to invoke CreateIpv4Gateway API
func CreateCreateIpv4GatewayRequest() (request *CreateIpv4GatewayRequest) {
	request = &CreateIpv4GatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateIpv4Gateway", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateIpv4GatewayResponse creates a response to parse from CreateIpv4Gateway response
func CreateCreateIpv4GatewayResponse() (response *CreateIpv4GatewayResponse) {
	response = &CreateIpv4GatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
