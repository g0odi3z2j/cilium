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

// AssociateEipAddressBatch invokes the vpc.AssociateEipAddressBatch API synchronously
func (client *Client) AssociateEipAddressBatch(request *AssociateEipAddressBatchRequest) (response *AssociateEipAddressBatchResponse, err error) {
	response = CreateAssociateEipAddressBatchResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateEipAddressBatchWithChan invokes the vpc.AssociateEipAddressBatch API asynchronously
func (client *Client) AssociateEipAddressBatchWithChan(request *AssociateEipAddressBatchRequest) (<-chan *AssociateEipAddressBatchResponse, <-chan error) {
	responseChan := make(chan *AssociateEipAddressBatchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateEipAddressBatch(request)
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

// AssociateEipAddressBatchWithCallback invokes the vpc.AssociateEipAddressBatch API asynchronously
func (client *Client) AssociateEipAddressBatchWithCallback(request *AssociateEipAddressBatchRequest, callback func(response *AssociateEipAddressBatchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateEipAddressBatchResponse
		var err error
		defer close(result)
		response, err = client.AssociateEipAddressBatch(request)
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

// AssociateEipAddressBatchRequest is the request struct for api AssociateEipAddressBatch
type AssociateEipAddressBatchRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	BindedInstanceType   string           `position:"Query" name:"BindedInstanceType"`
	BindedInstanceId     string           `position:"Query" name:"BindedInstanceId"`
	Mode                 string           `position:"Query" name:"Mode"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds          *[]string        `position:"Query" name:"InstanceIds"  type:"Repeated"`
}

// AssociateEipAddressBatchResponse is the response struct for api AssociateEipAddressBatch
type AssociateEipAddressBatchResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateEipAddressBatchRequest creates a request to invoke AssociateEipAddressBatch API
func CreateAssociateEipAddressBatchRequest() (request *AssociateEipAddressBatchRequest) {
	request = &AssociateEipAddressBatchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AssociateEipAddressBatch", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateEipAddressBatchResponse creates a response to parse from AssociateEipAddressBatch response
func CreateAssociateEipAddressBatchResponse() (response *AssociateEipAddressBatchResponse) {
	response = &AssociateEipAddressBatchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
