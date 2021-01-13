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

// AllocateEipAddressPro invokes the vpc.AllocateEipAddressPro API synchronously
func (client *Client) AllocateEipAddressPro(request *AllocateEipAddressProRequest) (response *AllocateEipAddressProResponse, err error) {
	response = CreateAllocateEipAddressProResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateEipAddressProWithChan invokes the vpc.AllocateEipAddressPro API asynchronously
func (client *Client) AllocateEipAddressProWithChan(request *AllocateEipAddressProRequest) (<-chan *AllocateEipAddressProResponse, <-chan error) {
	responseChan := make(chan *AllocateEipAddressProResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateEipAddressPro(request)
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

// AllocateEipAddressProWithCallback invokes the vpc.AllocateEipAddressPro API asynchronously
func (client *Client) AllocateEipAddressProWithCallback(request *AllocateEipAddressProRequest, callback func(response *AllocateEipAddressProResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateEipAddressProResponse
		var err error
		defer close(result)
		response, err = client.AllocateEipAddressPro(request)
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

// AllocateEipAddressProRequest is the request struct for api AllocateEipAddressPro
type AllocateEipAddressProRequest struct {
	*requests.RpcRequest
	IpAddress            string           `position:"Query" name:"IpAddress"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ISP                  string           `position:"Query" name:"ISP"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	Netmode              string           `position:"Query" name:"Netmode"`
	InstanceChargeType   string           `position:"Query" name:"InstanceChargeType"`
	Period               requests.Integer `position:"Query" name:"Period"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string           `position:"Query" name:"Bandwidth"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	InternetChargeType   string           `position:"Query" name:"InternetChargeType"`
	PricingCycle         string           `position:"Query" name:"PricingCycle"`
}

// AllocateEipAddressProResponse is the response struct for api AllocateEipAddressPro
type AllocateEipAddressProResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	AllocationId    string `json:"AllocationId" xml:"AllocationId"`
	EipAddress      string `json:"EipAddress" xml:"EipAddress"`
	OrderId         int64  `json:"OrderId" xml:"OrderId"`
	ResourceGroupId string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

// CreateAllocateEipAddressProRequest creates a request to invoke AllocateEipAddressPro API
func CreateAllocateEipAddressProRequest() (request *AllocateEipAddressProRequest) {
	request = &AllocateEipAddressProRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AllocateEipAddressPro", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAllocateEipAddressProResponse creates a response to parse from AllocateEipAddressPro response
func CreateAllocateEipAddressProResponse() (response *AllocateEipAddressProResponse) {
	response = &AllocateEipAddressProResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
