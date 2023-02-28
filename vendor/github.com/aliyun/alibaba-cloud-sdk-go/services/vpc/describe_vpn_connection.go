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

// DescribeVpnConnection invokes the vpc.DescribeVpnConnection API synchronously
func (client *Client) DescribeVpnConnection(request *DescribeVpnConnectionRequest) (response *DescribeVpnConnectionResponse, err error) {
	response = CreateDescribeVpnConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVpnConnectionWithChan invokes the vpc.DescribeVpnConnection API asynchronously
func (client *Client) DescribeVpnConnectionWithChan(request *DescribeVpnConnectionRequest) (<-chan *DescribeVpnConnectionResponse, <-chan error) {
	responseChan := make(chan *DescribeVpnConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpnConnection(request)
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

// DescribeVpnConnectionWithCallback invokes the vpc.DescribeVpnConnection API asynchronously
func (client *Client) DescribeVpnConnectionWithCallback(request *DescribeVpnConnectionRequest, callback func(response *DescribeVpnConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpnConnectionResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpnConnection(request)
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

// DescribeVpnConnectionRequest is the request struct for api DescribeVpnConnection
type DescribeVpnConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VpnConnectionId      string           `position:"Query" name:"VpnConnectionId"`
}

// DescribeVpnConnectionResponse is the response struct for api DescribeVpnConnection
type DescribeVpnConnectionResponse struct {
	*responses.BaseResponse
	Status                 string                      `json:"Status" xml:"Status"`
	RemoteCaCertificate    string                      `json:"RemoteCaCertificate" xml:"RemoteCaCertificate"`
	EnableNatTraversal     bool                        `json:"EnableNatTraversal" xml:"EnableNatTraversal"`
	CreateTime             int64                       `json:"CreateTime" xml:"CreateTime"`
	EffectImmediately      bool                        `json:"EffectImmediately" xml:"EffectImmediately"`
	VpnGatewayId           string                      `json:"VpnGatewayId" xml:"VpnGatewayId"`
	LocalSubnet            string                      `json:"LocalSubnet" xml:"LocalSubnet"`
	RequestId              string                      `json:"RequestId" xml:"RequestId"`
	VpnConnectionId        string                      `json:"VpnConnectionId" xml:"VpnConnectionId"`
	RemoteSubnet           string                      `json:"RemoteSubnet" xml:"RemoteSubnet"`
	CustomerGatewayId      string                      `json:"CustomerGatewayId" xml:"CustomerGatewayId"`
	Name                   string                      `json:"Name" xml:"Name"`
	EnableDpd              bool                        `json:"EnableDpd" xml:"EnableDpd"`
	AttachType             string                      `json:"AttachType" xml:"AttachType"`
	NetworkType            string                      `json:"NetworkType" xml:"NetworkType"`
	AttachInstanceId       string                      `json:"AttachInstanceId" xml:"AttachInstanceId"`
	Spec                   string                      `json:"Spec" xml:"Spec"`
	State                  string                      `json:"State" xml:"State"`
	ZoneNo                 string                      `json:"ZoneNo" xml:"ZoneNo"`
	InternetIp             string                      `json:"InternetIp" xml:"InternetIp"`
	TransitRouterId        string                      `json:"TransitRouterId" xml:"TransitRouterId"`
	TransitRouterName      string                      `json:"TransitRouterName" xml:"TransitRouterName"`
	CrossAccountAuthorized bool                        `json:"CrossAccountAuthorized" xml:"CrossAccountAuthorized"`
	IkeConfig              IkeConfig                   `json:"IkeConfig" xml:"IkeConfig"`
	IpsecConfig            IpsecConfig                 `json:"IpsecConfig" xml:"IpsecConfig"`
	VcoHealthCheck         VcoHealthCheck              `json:"VcoHealthCheck" xml:"VcoHealthCheck"`
	VpnBgpConfig           VpnBgpConfig                `json:"VpnBgpConfig" xml:"VpnBgpConfig"`
	Tags                   TagsInDescribeVpnConnection `json:"Tags" xml:"Tags"`
}

// CreateDescribeVpnConnectionRequest creates a request to invoke DescribeVpnConnection API
func CreateDescribeVpnConnectionRequest() (request *DescribeVpnConnectionRequest) {
	request = &DescribeVpnConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnConnection", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVpnConnectionResponse creates a response to parse from DescribeVpnConnection response
func CreateDescribeVpnConnectionResponse() (response *DescribeVpnConnectionResponse) {
	response = &DescribeVpnConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
