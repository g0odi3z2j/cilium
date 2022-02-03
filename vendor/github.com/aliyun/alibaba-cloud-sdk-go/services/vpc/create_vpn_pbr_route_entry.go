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

// CreateVpnPbrRouteEntry invokes the vpc.CreateVpnPbrRouteEntry API synchronously
func (client *Client) CreateVpnPbrRouteEntry(request *CreateVpnPbrRouteEntryRequest) (response *CreateVpnPbrRouteEntryResponse, err error) {
	response = CreateCreateVpnPbrRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVpnPbrRouteEntryWithChan invokes the vpc.CreateVpnPbrRouteEntry API asynchronously
func (client *Client) CreateVpnPbrRouteEntryWithChan(request *CreateVpnPbrRouteEntryRequest) (<-chan *CreateVpnPbrRouteEntryResponse, <-chan error) {
	responseChan := make(chan *CreateVpnPbrRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVpnPbrRouteEntry(request)
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

// CreateVpnPbrRouteEntryWithCallback invokes the vpc.CreateVpnPbrRouteEntry API asynchronously
func (client *Client) CreateVpnPbrRouteEntryWithCallback(request *CreateVpnPbrRouteEntryRequest, callback func(response *CreateVpnPbrRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVpnPbrRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateVpnPbrRouteEntry(request)
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

// CreateVpnPbrRouteEntryRequest is the request struct for api CreateVpnPbrRouteEntry
type CreateVpnPbrRouteEntryRequest struct {
	*requests.RpcRequest
	RouteSource          string           `position:"Query" name:"RouteSource"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	PublishVpc           requests.Boolean `position:"Query" name:"PublishVpc"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Weight               requests.Integer `position:"Query" name:"Weight"`
	VpnGatewayId         string           `position:"Query" name:"VpnGatewayId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouteDest            string           `position:"Query" name:"RouteDest"`
	NextHop              string           `position:"Query" name:"NextHop"`
	OverlayMode          string           `position:"Query" name:"OverlayMode"`
}

// CreateVpnPbrRouteEntryResponse is the response struct for api CreateVpnPbrRouteEntry
type CreateVpnPbrRouteEntryResponse struct {
	*responses.BaseResponse
	NextHop       string `json:"NextHop" xml:"NextHop"`
	Weight        int    `json:"Weight" xml:"Weight"`
	RouteDest     string `json:"RouteDest" xml:"RouteDest"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Description   string `json:"Description" xml:"Description"`
	State         string `json:"State" xml:"State"`
	CreateTime    int64  `json:"CreateTime" xml:"CreateTime"`
	OverlayMode   string `json:"OverlayMode" xml:"OverlayMode"`
	RouteSource   string `json:"RouteSource" xml:"RouteSource"`
	VpnInstanceId string `json:"VpnInstanceId" xml:"VpnInstanceId"`
}

// CreateCreateVpnPbrRouteEntryRequest creates a request to invoke CreateVpnPbrRouteEntry API
func CreateCreateVpnPbrRouteEntryRequest() (request *CreateVpnPbrRouteEntryRequest) {
	request = &CreateVpnPbrRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateVpnPbrRouteEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVpnPbrRouteEntryResponse creates a response to parse from CreateVpnPbrRouteEntry response
func CreateCreateVpnPbrRouteEntryResponse() (response *CreateVpnPbrRouteEntryResponse) {
	response = &CreateVpnPbrRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
