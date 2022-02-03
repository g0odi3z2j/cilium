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

// AssociateRouteTableWithGateway invokes the vpc.AssociateRouteTableWithGateway API synchronously
func (client *Client) AssociateRouteTableWithGateway(request *AssociateRouteTableWithGatewayRequest) (response *AssociateRouteTableWithGatewayResponse, err error) {
	response = CreateAssociateRouteTableWithGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateRouteTableWithGatewayWithChan invokes the vpc.AssociateRouteTableWithGateway API asynchronously
func (client *Client) AssociateRouteTableWithGatewayWithChan(request *AssociateRouteTableWithGatewayRequest) (<-chan *AssociateRouteTableWithGatewayResponse, <-chan error) {
	responseChan := make(chan *AssociateRouteTableWithGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateRouteTableWithGateway(request)
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

// AssociateRouteTableWithGatewayWithCallback invokes the vpc.AssociateRouteTableWithGateway API asynchronously
func (client *Client) AssociateRouteTableWithGatewayWithCallback(request *AssociateRouteTableWithGatewayRequest, callback func(response *AssociateRouteTableWithGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateRouteTableWithGatewayResponse
		var err error
		defer close(result)
		response, err = client.AssociateRouteTableWithGateway(request)
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

// AssociateRouteTableWithGatewayRequest is the request struct for api AssociateRouteTableWithGateway
type AssociateRouteTableWithGatewayRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	GatewayId            string           `position:"Query" name:"GatewayId"`
	RouteTableId         string           `position:"Query" name:"RouteTableId"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// AssociateRouteTableWithGatewayResponse is the response struct for api AssociateRouteTableWithGateway
type AssociateRouteTableWithGatewayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateRouteTableWithGatewayRequest creates a request to invoke AssociateRouteTableWithGateway API
func CreateAssociateRouteTableWithGatewayRequest() (request *AssociateRouteTableWithGatewayRequest) {
	request = &AssociateRouteTableWithGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AssociateRouteTableWithGateway", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateRouteTableWithGatewayResponse creates a response to parse from AssociateRouteTableWithGateway response
func CreateAssociateRouteTableWithGatewayResponse() (response *AssociateRouteTableWithGatewayResponse) {
	response = &AssociateRouteTableWithGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
