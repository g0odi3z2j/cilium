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

// ModifyVpnGatewayAttribute invokes the vpc.ModifyVpnGatewayAttribute API synchronously
func (client *Client) ModifyVpnGatewayAttribute(request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
	response = CreateModifyVpnGatewayAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpnGatewayAttributeWithChan invokes the vpc.ModifyVpnGatewayAttribute API asynchronously
func (client *Client) ModifyVpnGatewayAttributeWithChan(request *ModifyVpnGatewayAttributeRequest) (<-chan *ModifyVpnGatewayAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVpnGatewayAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpnGatewayAttribute(request)
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

// ModifyVpnGatewayAttributeWithCallback invokes the vpc.ModifyVpnGatewayAttribute API asynchronously
func (client *Client) ModifyVpnGatewayAttributeWithCallback(request *ModifyVpnGatewayAttributeRequest, callback func(response *ModifyVpnGatewayAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpnGatewayAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpnGatewayAttribute(request)
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

// ModifyVpnGatewayAttributeRequest is the request struct for api ModifyVpnGatewayAttribute
type ModifyVpnGatewayAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AutoPropagate        requests.Boolean `position:"Query" name:"AutoPropagate"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string           `position:"Query" name:"VpnGatewayId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyVpnGatewayAttributeResponse is the response struct for api ModifyVpnGatewayAttribute
type ModifyVpnGatewayAttributeResponse struct {
	*responses.BaseResponse
	VpcId          string `json:"VpcId" xml:"VpcId"`
	Status         string `json:"Status" xml:"Status"`
	InternetIp     string `json:"InternetIp" xml:"InternetIp"`
	Spec           string `json:"Spec" xml:"Spec"`
	CreateTime     int64  `json:"CreateTime" xml:"CreateTime"`
	AutoPropagate  bool   `json:"AutoPropagate" xml:"AutoPropagate"`
	VpnGatewayId   string `json:"VpnGatewayId" xml:"VpnGatewayId"`
	IntranetIp     string `json:"IntranetIp" xml:"IntranetIp"`
	EndTime        int64  `json:"EndTime" xml:"EndTime"`
	VSwitchId      string `json:"VSwitchId" xml:"VSwitchId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Description    string `json:"Description" xml:"Description"`
	EnableBgp      bool   `json:"EnableBgp" xml:"EnableBgp"`
	BusinessStatus string `json:"BusinessStatus" xml:"BusinessStatus"`
	Name           string `json:"Name" xml:"Name"`
}

// CreateModifyVpnGatewayAttributeRequest creates a request to invoke ModifyVpnGatewayAttribute API
func CreateModifyVpnGatewayAttributeRequest() (request *ModifyVpnGatewayAttributeRequest) {
	request = &ModifyVpnGatewayAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpnGatewayAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpnGatewayAttributeResponse creates a response to parse from ModifyVpnGatewayAttribute response
func CreateModifyVpnGatewayAttributeResponse() (response *ModifyVpnGatewayAttributeResponse) {
	response = &ModifyVpnGatewayAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
