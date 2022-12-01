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

// ModifyVpnAttachmentAttribute invokes the vpc.ModifyVpnAttachmentAttribute API synchronously
func (client *Client) ModifyVpnAttachmentAttribute(request *ModifyVpnAttachmentAttributeRequest) (response *ModifyVpnAttachmentAttributeResponse, err error) {
	response = CreateModifyVpnAttachmentAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpnAttachmentAttributeWithChan invokes the vpc.ModifyVpnAttachmentAttribute API asynchronously
func (client *Client) ModifyVpnAttachmentAttributeWithChan(request *ModifyVpnAttachmentAttributeRequest) (<-chan *ModifyVpnAttachmentAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVpnAttachmentAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpnAttachmentAttribute(request)
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

// ModifyVpnAttachmentAttributeWithCallback invokes the vpc.ModifyVpnAttachmentAttribute API asynchronously
func (client *Client) ModifyVpnAttachmentAttributeWithCallback(request *ModifyVpnAttachmentAttributeRequest, callback func(response *ModifyVpnAttachmentAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpnAttachmentAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpnAttachmentAttribute(request)
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

// ModifyVpnAttachmentAttributeRequest is the request struct for api ModifyVpnAttachmentAttribute
type ModifyVpnAttachmentAttributeRequest struct {
	*requests.RpcRequest
	IkeConfig            string           `position:"Query" name:"IkeConfig"`
	AutoConfigRoute      requests.Boolean `position:"Query" name:"AutoConfigRoute"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	IpsecConfig          string           `position:"Query" name:"IpsecConfig"`
	BgpConfig            string           `position:"Query" name:"BgpConfig"`
	NetworkType          string           `position:"Query" name:"NetworkType"`
	HealthCheckConfig    string           `position:"Query" name:"HealthCheckConfig"`
	CustomerGatewayId    string           `position:"Query" name:"CustomerGatewayId"`
	LocalSubnet          string           `position:"Query" name:"LocalSubnet"`
	RemoteCaCert         string           `position:"Query" name:"RemoteCaCert"`
	RemoteSubnet         string           `position:"Query" name:"RemoteSubnet"`
	EffectImmediately    requests.Boolean `position:"Query" name:"EffectImmediately"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EnableDpd            requests.Boolean `position:"Query" name:"EnableDpd"`
	VpnConnectionId      string           `position:"Query" name:"VpnConnectionId"`
	Name                 string           `position:"Query" name:"Name"`
	EnableNatTraversal   requests.Boolean `position:"Query" name:"EnableNatTraversal"`
}

// ModifyVpnAttachmentAttributeResponse is the response struct for api ModifyVpnAttachmentAttribute
type ModifyVpnAttachmentAttributeResponse struct {
	*responses.BaseResponse
	VpnConnectionId    string         `json:"VpnConnectionId" xml:"VpnConnectionId"`
	CustomerGatewayId  string         `json:"CustomerGatewayId" xml:"CustomerGatewayId"`
	VpnGatewayId       string         `json:"VpnGatewayId" xml:"VpnGatewayId"`
	Name               string         `json:"Name" xml:"Name"`
	Description        string         `json:"Description" xml:"Description"`
	LocalSubnet        string         `json:"LocalSubnet" xml:"LocalSubnet"`
	RemoteSubnet       string         `json:"RemoteSubnet" xml:"RemoteSubnet"`
	CreateTime         int64          `json:"CreateTime" xml:"CreateTime"`
	EffectImmediately  bool           `json:"EffectImmediately" xml:"EffectImmediately"`
	Status             string         `json:"Status" xml:"Status"`
	EnableDpd          bool           `json:"EnableDpd" xml:"EnableDpd"`
	EnableNatTraversal bool           `json:"EnableNatTraversal" xml:"EnableNatTraversal"`
	AttachType         string         `json:"AttachType" xml:"AttachType"`
	NetworkType        string         `json:"NetworkType" xml:"NetworkType"`
	AttachInstanceId   string         `json:"AttachInstanceId" xml:"AttachInstanceId"`
	Spec               string         `json:"Spec" xml:"Spec"`
	RequestId          string         `json:"RequestId" xml:"RequestId"`
	IkeConfig          IkeConfig      `json:"IkeConfig" xml:"IkeConfig"`
	IpsecConfig        IpsecConfig    `json:"IpsecConfig" xml:"IpsecConfig"`
	VcoHealthCheck     VcoHealthCheck `json:"VcoHealthCheck" xml:"VcoHealthCheck"`
	VpnBgpConfig       VpnBgpConfig   `json:"VpnBgpConfig" xml:"VpnBgpConfig"`
}

// CreateModifyVpnAttachmentAttributeRequest creates a request to invoke ModifyVpnAttachmentAttribute API
func CreateModifyVpnAttachmentAttributeRequest() (request *ModifyVpnAttachmentAttributeRequest) {
	request = &ModifyVpnAttachmentAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpnAttachmentAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpnAttachmentAttributeResponse creates a response to parse from ModifyVpnAttachmentAttribute response
func CreateModifyVpnAttachmentAttributeResponse() (response *ModifyVpnAttachmentAttributeResponse) {
	response = &ModifyVpnAttachmentAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
