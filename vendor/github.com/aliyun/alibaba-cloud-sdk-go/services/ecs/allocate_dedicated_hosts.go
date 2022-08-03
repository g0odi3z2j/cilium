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

// AllocateDedicatedHosts invokes the ecs.AllocateDedicatedHosts API synchronously
func (client *Client) AllocateDedicatedHosts(request *AllocateDedicatedHostsRequest) (response *AllocateDedicatedHostsResponse, err error) {
	response = CreateAllocateDedicatedHostsResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateDedicatedHostsWithChan invokes the ecs.AllocateDedicatedHosts API asynchronously
func (client *Client) AllocateDedicatedHostsWithChan(request *AllocateDedicatedHostsRequest) (<-chan *AllocateDedicatedHostsResponse, <-chan error) {
	responseChan := make(chan *AllocateDedicatedHostsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateDedicatedHosts(request)
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

// AllocateDedicatedHostsWithCallback invokes the ecs.AllocateDedicatedHosts API asynchronously
func (client *Client) AllocateDedicatedHostsWithCallback(request *AllocateDedicatedHostsRequest, callback func(response *AllocateDedicatedHostsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateDedicatedHostsResponse
		var err error
		defer close(result)
		response, err = client.AllocateDedicatedHosts(request)
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

// AllocateDedicatedHostsRequest is the request struct for api AllocateDedicatedHosts
type AllocateDedicatedHostsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                requests.Integer             `position:"Query" name:"ResourceOwnerId"`
	ClientToken                    string                       `position:"Query" name:"ClientToken"`
	Description                    string                       `position:"Query" name:"Description"`
	CpuOverCommitRatio             requests.Float               `position:"Query" name:"CpuOverCommitRatio"`
	ResourceGroupId                string                       `position:"Query" name:"ResourceGroupId"`
	MinQuantity                    requests.Integer             `position:"Query" name:"MinQuantity"`
	ActionOnMaintenance            string                       `position:"Query" name:"ActionOnMaintenance"`
	DedicatedHostClusterId         string                       `position:"Query" name:"DedicatedHostClusterId"`
	Tag                            *[]AllocateDedicatedHostsTag `position:"Query" name:"Tag"  type:"Repeated"`
	DedicatedHostType              string                       `position:"Query" name:"DedicatedHostType"`
	AutoRenewPeriod                requests.Integer             `position:"Query" name:"AutoRenewPeriod"`
	Period                         requests.Integer             `position:"Query" name:"Period"`
	Quantity                       requests.Integer             `position:"Query" name:"Quantity"`
	DedicatedHostName              string                       `position:"Query" name:"DedicatedHostName"`
	ResourceOwnerAccount           string                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                   string                       `position:"Query" name:"OwnerAccount"`
	AutoReleaseTime                string                       `position:"Query" name:"AutoReleaseTime"`
	OwnerId                        requests.Integer             `position:"Query" name:"OwnerId"`
	SchedulerOptionsFenceId        string                       `position:"Query" name:"SchedulerOptions.FenceId"`
	PeriodUnit                     string                       `position:"Query" name:"PeriodUnit"`
	AutoRenew                      requests.Boolean             `position:"Query" name:"AutoRenew"`
	NetworkAttributesSlbUdpTimeout requests.Integer             `position:"Query" name:"NetworkAttributes.SlbUdpTimeout"`
	ZoneId                         string                       `position:"Query" name:"ZoneId"`
	AutoPlacement                  string                       `position:"Query" name:"AutoPlacement"`
	ChargeType                     string                       `position:"Query" name:"ChargeType"`
	NetworkAttributesUdpTimeout    requests.Integer             `position:"Query" name:"NetworkAttributes.UdpTimeout"`
}

// AllocateDedicatedHostsTag is a repeated param struct in AllocateDedicatedHostsRequest
type AllocateDedicatedHostsTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// AllocateDedicatedHostsResponse is the response struct for api AllocateDedicatedHosts
type AllocateDedicatedHostsResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	DedicatedHostIdSets DedicatedHostIdSets `json:"DedicatedHostIdSets" xml:"DedicatedHostIdSets"`
}

// CreateAllocateDedicatedHostsRequest creates a request to invoke AllocateDedicatedHosts API
func CreateAllocateDedicatedHostsRequest() (request *AllocateDedicatedHostsRequest) {
	request = &AllocateDedicatedHostsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AllocateDedicatedHosts", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAllocateDedicatedHostsResponse creates a response to parse from AllocateDedicatedHosts response
func CreateAllocateDedicatedHostsResponse() (response *AllocateDedicatedHostsResponse) {
	response = &AllocateDedicatedHostsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
