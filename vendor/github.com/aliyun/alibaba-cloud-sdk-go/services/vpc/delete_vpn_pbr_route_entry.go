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

// DeleteVpnPbrRouteEntry invokes the vpc.DeleteVpnPbrRouteEntry API synchronously
func (client *Client) DeleteVpnPbrRouteEntry(request *DeleteVpnPbrRouteEntryRequest) (response *DeleteVpnPbrRouteEntryResponse, err error) {
	response = CreateDeleteVpnPbrRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVpnPbrRouteEntryWithChan invokes the vpc.DeleteVpnPbrRouteEntry API asynchronously
func (client *Client) DeleteVpnPbrRouteEntryWithChan(request *DeleteVpnPbrRouteEntryRequest) (<-chan *DeleteVpnPbrRouteEntryResponse, <-chan error) {
	responseChan := make(chan *DeleteVpnPbrRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVpnPbrRouteEntry(request)
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

// DeleteVpnPbrRouteEntryWithCallback invokes the vpc.DeleteVpnPbrRouteEntry API asynchronously
func (client *Client) DeleteVpnPbrRouteEntryWithCallback(request *DeleteVpnPbrRouteEntryRequest, callback func(response *DeleteVpnPbrRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVpnPbrRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.DeleteVpnPbrRouteEntry(request)
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

// DeleteVpnPbrRouteEntryRequest is the request struct for api DeleteVpnPbrRouteEntry
type DeleteVpnPbrRouteEntryRequest struct {
	*requests.RpcRequest
	RouteSource          string           `position:"Query" name:"RouteSource"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Weight               requests.Integer `position:"Query" name:"Weight"`
	VpnGatewayId         string           `position:"Query" name:"VpnGatewayId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouteDest            string           `position:"Query" name:"RouteDest"`
	NextHop              string           `position:"Query" name:"NextHop"`
	OverlayMode          string           `position:"Query" name:"OverlayMode"`
}

// DeleteVpnPbrRouteEntryResponse is the response struct for api DeleteVpnPbrRouteEntry
type DeleteVpnPbrRouteEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVpnPbrRouteEntryRequest creates a request to invoke DeleteVpnPbrRouteEntry API
func CreateDeleteVpnPbrRouteEntryRequest() (request *DeleteVpnPbrRouteEntryRequest) {
	request = &DeleteVpnPbrRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVpnPbrRouteEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVpnPbrRouteEntryResponse creates a response to parse from DeleteVpnPbrRouteEntry response
func CreateDeleteVpnPbrRouteEntryResponse() (response *DeleteVpnPbrRouteEntryResponse) {
	response = &DeleteVpnPbrRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
