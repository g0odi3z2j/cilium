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

// CreatePublicIpAddressPool invokes the vpc.CreatePublicIpAddressPool API synchronously
func (client *Client) CreatePublicIpAddressPool(request *CreatePublicIpAddressPoolRequest) (response *CreatePublicIpAddressPoolResponse, err error) {
	response = CreateCreatePublicIpAddressPoolResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePublicIpAddressPoolWithChan invokes the vpc.CreatePublicIpAddressPool API asynchronously
func (client *Client) CreatePublicIpAddressPoolWithChan(request *CreatePublicIpAddressPoolRequest) (<-chan *CreatePublicIpAddressPoolResponse, <-chan error) {
	responseChan := make(chan *CreatePublicIpAddressPoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePublicIpAddressPool(request)
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

// CreatePublicIpAddressPoolWithCallback invokes the vpc.CreatePublicIpAddressPool API asynchronously
func (client *Client) CreatePublicIpAddressPoolWithCallback(request *CreatePublicIpAddressPoolRequest, callback func(response *CreatePublicIpAddressPoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePublicIpAddressPoolResponse
		var err error
		defer close(result)
		response, err = client.CreatePublicIpAddressPool(request)
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

// CreatePublicIpAddressPoolRequest is the request struct for api CreatePublicIpAddressPool
type CreatePublicIpAddressPoolRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string                          `position:"Query" name:"ClientToken"`
	Isp                  string                          `position:"Query" name:"Isp"`
	Description          string                          `position:"Query" name:"Description"`
	ResourceGroupId      string                          `position:"Query" name:"ResourceGroupId"`
	Tag                  *[]CreatePublicIpAddressPoolTag `position:"Query" name:"Tag"  type:"Repeated"`
	DryRun               requests.Boolean                `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string                          `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                          `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                `position:"Query" name:"OwnerId"`
	Name                 string                          `position:"Query" name:"Name"`
}

// CreatePublicIpAddressPoolTag is a repeated param struct in CreatePublicIpAddressPoolRequest
type CreatePublicIpAddressPoolTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreatePublicIpAddressPoolResponse is the response struct for api CreatePublicIpAddressPool
type CreatePublicIpAddressPoolResponse struct {
	*responses.BaseResponse
	PulbicIpAddressPoolId string `json:"PulbicIpAddressPoolId" xml:"PulbicIpAddressPoolId"`
	RequestId             string `json:"RequestId" xml:"RequestId"`
	ResourceGroupId       string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

// CreateCreatePublicIpAddressPoolRequest creates a request to invoke CreatePublicIpAddressPool API
func CreateCreatePublicIpAddressPoolRequest() (request *CreatePublicIpAddressPoolRequest) {
	request = &CreatePublicIpAddressPoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreatePublicIpAddressPool", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePublicIpAddressPoolResponse creates a response to parse from CreatePublicIpAddressPool response
func CreateCreatePublicIpAddressPoolResponse() (response *CreatePublicIpAddressPoolResponse) {
	response = &CreatePublicIpAddressPoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
