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

// DeleteVpnRouteEntry invokes the vpc.DeleteVpnRouteEntry API synchronously
func (client *Client) DeleteVpnRouteEntry(request *DeleteVpnRouteEntryRequest) (response *DeleteVpnRouteEntryResponse, err error) {
	response = CreateDeleteVpnRouteEntryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVpnRouteEntryWithChan invokes the vpc.DeleteVpnRouteEntry API asynchronously
func (client *Client) DeleteVpnRouteEntryWithChan(request *DeleteVpnRouteEntryRequest) (<-chan *DeleteVpnRouteEntryResponse, <-chan error) {
	responseChan := make(chan *DeleteVpnRouteEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVpnRouteEntry(request)
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

// DeleteVpnRouteEntryWithCallback invokes the vpc.DeleteVpnRouteEntry API asynchronously
func (client *Client) DeleteVpnRouteEntryWithCallback(request *DeleteVpnRouteEntryRequest, callback func(response *DeleteVpnRouteEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVpnRouteEntryResponse
		var err error
		defer close(result)
		response, err = client.DeleteVpnRouteEntry(request)
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

// DeleteVpnRouteEntryRequest is the request struct for api DeleteVpnRouteEntry
type DeleteVpnRouteEntryRequest struct {
	*requests.RpcRequest
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

// DeleteVpnRouteEntryResponse is the response struct for api DeleteVpnRouteEntry
type DeleteVpnRouteEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVpnRouteEntryRequest creates a request to invoke DeleteVpnRouteEntry API
func CreateDeleteVpnRouteEntryRequest() (request *DeleteVpnRouteEntryRequest) {
	request = &DeleteVpnRouteEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVpnRouteEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVpnRouteEntryResponse creates a response to parse from DeleteVpnRouteEntry response
func CreateDeleteVpnRouteEntryResponse() (response *DeleteVpnRouteEntryResponse) {
	response = &DeleteVpnRouteEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
