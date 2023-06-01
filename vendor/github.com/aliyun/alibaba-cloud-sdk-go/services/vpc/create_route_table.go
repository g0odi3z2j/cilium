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

// CreateRouteTable invokes the vpc.CreateRouteTable API synchronously
func (client *Client) CreateRouteTable(request *CreateRouteTableRequest) (response *CreateRouteTableResponse, err error) {
	response = CreateCreateRouteTableResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRouteTableWithChan invokes the vpc.CreateRouteTable API asynchronously
func (client *Client) CreateRouteTableWithChan(request *CreateRouteTableRequest) (<-chan *CreateRouteTableResponse, <-chan error) {
	responseChan := make(chan *CreateRouteTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRouteTable(request)
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

// CreateRouteTableWithCallback invokes the vpc.CreateRouteTable API asynchronously
func (client *Client) CreateRouteTableWithCallback(request *CreateRouteTableRequest, callback func(response *CreateRouteTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRouteTableResponse
		var err error
		defer close(result)
		response, err = client.CreateRouteTable(request)
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

// CreateRouteTableRequest is the request struct for api CreateRouteTable
type CreateRouteTableRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	RouteTableName       string           `position:"Query" name:"RouteTableName"`
	AssociateType        string           `position:"Query" name:"AssociateType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
}

// CreateRouteTableResponse is the response struct for api CreateRouteTable
type CreateRouteTableResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	RouteTableId string `json:"RouteTableId" xml:"RouteTableId"`
}

// CreateCreateRouteTableRequest creates a request to invoke CreateRouteTable API
func CreateCreateRouteTableRequest() (request *CreateRouteTableRequest) {
	request = &CreateRouteTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateRouteTable", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRouteTableResponse creates a response to parse from CreateRouteTable response
func CreateCreateRouteTableResponse() (response *CreateRouteTableResponse) {
	response = &CreateRouteTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
