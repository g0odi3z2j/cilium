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

// DeleteRouteTable invokes the vpc.DeleteRouteTable API synchronously
func (client *Client) DeleteRouteTable(request *DeleteRouteTableRequest) (response *DeleteRouteTableResponse, err error) {
	response = CreateDeleteRouteTableResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRouteTableWithChan invokes the vpc.DeleteRouteTable API asynchronously
func (client *Client) DeleteRouteTableWithChan(request *DeleteRouteTableRequest) (<-chan *DeleteRouteTableResponse, <-chan error) {
	responseChan := make(chan *DeleteRouteTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRouteTable(request)
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

// DeleteRouteTableWithCallback invokes the vpc.DeleteRouteTable API asynchronously
func (client *Client) DeleteRouteTableWithCallback(request *DeleteRouteTableRequest, callback func(response *DeleteRouteTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRouteTableResponse
		var err error
		defer close(result)
		response, err = client.DeleteRouteTable(request)
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

// DeleteRouteTableRequest is the request struct for api DeleteRouteTable
type DeleteRouteTableRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouteTableId         string           `position:"Query" name:"RouteTableId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteRouteTableResponse is the response struct for api DeleteRouteTable
type DeleteRouteTableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRouteTableRequest creates a request to invoke DeleteRouteTable API
func CreateDeleteRouteTableRequest() (request *DeleteRouteTableRequest) {
	request = &DeleteRouteTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteRouteTable", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteRouteTableResponse creates a response to parse from DeleteRouteTable response
func CreateDeleteRouteTableResponse() (response *DeleteRouteTableResponse) {
	response = &DeleteRouteTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
