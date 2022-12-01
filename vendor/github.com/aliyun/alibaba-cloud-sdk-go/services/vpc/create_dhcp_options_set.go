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

// CreateDhcpOptionsSet invokes the vpc.CreateDhcpOptionsSet API synchronously
func (client *Client) CreateDhcpOptionsSet(request *CreateDhcpOptionsSetRequest) (response *CreateDhcpOptionsSetResponse, err error) {
	response = CreateCreateDhcpOptionsSetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDhcpOptionsSetWithChan invokes the vpc.CreateDhcpOptionsSet API asynchronously
func (client *Client) CreateDhcpOptionsSetWithChan(request *CreateDhcpOptionsSetRequest) (<-chan *CreateDhcpOptionsSetResponse, <-chan error) {
	responseChan := make(chan *CreateDhcpOptionsSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDhcpOptionsSet(request)
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

// CreateDhcpOptionsSetWithCallback invokes the vpc.CreateDhcpOptionsSet API asynchronously
func (client *Client) CreateDhcpOptionsSetWithCallback(request *CreateDhcpOptionsSetRequest, callback func(response *CreateDhcpOptionsSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDhcpOptionsSetResponse
		var err error
		defer close(result)
		response, err = client.CreateDhcpOptionsSet(request)
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

// CreateDhcpOptionsSetRequest is the request struct for api CreateDhcpOptionsSet
type CreateDhcpOptionsSetRequest struct {
	*requests.RpcRequest
	BootFileName              string           `position:"Query" name:"BootFileName"`
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken               string           `position:"Query" name:"ClientToken"`
	TFTPServerName            string           `position:"Query" name:"TFTPServerName"`
	LeaseTime                 string           `position:"Query" name:"LeaseTime"`
	DomainNameServers         string           `position:"Query" name:"DomainNameServers"`
	DhcpOptionsSetDescription string           `position:"Query" name:"DhcpOptionsSetDescription"`
	DryRun                    requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	DomainName                string           `position:"Query" name:"DomainName"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	DhcpOptionsSetName        string           `position:"Query" name:"DhcpOptionsSetName"`
	Ipv6LeaseTime             string           `position:"Query" name:"Ipv6LeaseTime"`
}

// CreateDhcpOptionsSetResponse is the response struct for api CreateDhcpOptionsSet
type CreateDhcpOptionsSetResponse struct {
	*responses.BaseResponse
	DhcpOptionsSetId string `json:"DhcpOptionsSetId" xml:"DhcpOptionsSetId"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDhcpOptionsSetRequest creates a request to invoke CreateDhcpOptionsSet API
func CreateCreateDhcpOptionsSetRequest() (request *CreateDhcpOptionsSetRequest) {
	request = &CreateDhcpOptionsSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateDhcpOptionsSet", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDhcpOptionsSetResponse creates a response to parse from CreateDhcpOptionsSet response
func CreateCreateDhcpOptionsSetResponse() (response *CreateDhcpOptionsSetResponse) {
	response = &CreateDhcpOptionsSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
