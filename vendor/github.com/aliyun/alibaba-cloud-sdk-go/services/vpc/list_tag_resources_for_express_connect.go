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

// ListTagResourcesForExpressConnect invokes the vpc.ListTagResourcesForExpressConnect API synchronously
func (client *Client) ListTagResourcesForExpressConnect(request *ListTagResourcesForExpressConnectRequest) (response *ListTagResourcesForExpressConnectResponse, err error) {
	response = CreateListTagResourcesForExpressConnectResponse()
	err = client.DoAction(request, response)
	return
}

// ListTagResourcesForExpressConnectWithChan invokes the vpc.ListTagResourcesForExpressConnect API asynchronously
func (client *Client) ListTagResourcesForExpressConnectWithChan(request *ListTagResourcesForExpressConnectRequest) (<-chan *ListTagResourcesForExpressConnectResponse, <-chan error) {
	responseChan := make(chan *ListTagResourcesForExpressConnectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTagResourcesForExpressConnect(request)
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

// ListTagResourcesForExpressConnectWithCallback invokes the vpc.ListTagResourcesForExpressConnect API asynchronously
func (client *Client) ListTagResourcesForExpressConnectWithCallback(request *ListTagResourcesForExpressConnectRequest, callback func(response *ListTagResourcesForExpressConnectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTagResourcesForExpressConnectResponse
		var err error
		defer close(result)
		response, err = client.ListTagResourcesForExpressConnect(request)
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

// ListTagResourcesForExpressConnectRequest is the request struct for api ListTagResourcesForExpressConnect
type ListTagResourcesForExpressConnectRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                        `position:"Query" name:"ResourceOwnerId"`
	NextToken            string                                  `position:"Query" name:"NextToken"`
	Tag                  *[]ListTagResourcesForExpressConnectTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceId           *[]string                               `position:"Query" name:"ResourceId"  type:"Repeated"`
	ResourceOwnerAccount string                                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                  `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                        `position:"Query" name:"OwnerId"`
	ResourceType         string                                  `position:"Query" name:"ResourceType"`
	MaxResults           requests.Integer                        `position:"Query" name:"MaxResults"`
}

// ListTagResourcesForExpressConnectTag is a repeated param struct in ListTagResourcesForExpressConnectRequest
type ListTagResourcesForExpressConnectTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListTagResourcesForExpressConnectResponse is the response struct for api ListTagResourcesForExpressConnect
type ListTagResourcesForExpressConnectResponse struct {
	*responses.BaseResponse
	NextToken    string                                          `json:"NextToken" xml:"NextToken"`
	RequestId    string                                          `json:"RequestId" xml:"RequestId"`
	TagResources TagResourcesInListTagResourcesForExpressConnect `json:"TagResources" xml:"TagResources"`
}

// CreateListTagResourcesForExpressConnectRequest creates a request to invoke ListTagResourcesForExpressConnect API
func CreateListTagResourcesForExpressConnectRequest() (request *ListTagResourcesForExpressConnectRequest) {
	request = &ListTagResourcesForExpressConnectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ListTagResourcesForExpressConnect", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTagResourcesForExpressConnectResponse creates a response to parse from ListTagResourcesForExpressConnect response
func CreateListTagResourcesForExpressConnectResponse() (response *ListTagResourcesForExpressConnectResponse) {
	response = &ListTagResourcesForExpressConnectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
