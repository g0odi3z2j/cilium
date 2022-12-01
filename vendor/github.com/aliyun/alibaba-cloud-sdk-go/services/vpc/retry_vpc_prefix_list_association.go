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

// RetryVpcPrefixListAssociation invokes the vpc.RetryVpcPrefixListAssociation API synchronously
func (client *Client) RetryVpcPrefixListAssociation(request *RetryVpcPrefixListAssociationRequest) (response *RetryVpcPrefixListAssociationResponse, err error) {
	response = CreateRetryVpcPrefixListAssociationResponse()
	err = client.DoAction(request, response)
	return
}

// RetryVpcPrefixListAssociationWithChan invokes the vpc.RetryVpcPrefixListAssociation API asynchronously
func (client *Client) RetryVpcPrefixListAssociationWithChan(request *RetryVpcPrefixListAssociationRequest) (<-chan *RetryVpcPrefixListAssociationResponse, <-chan error) {
	responseChan := make(chan *RetryVpcPrefixListAssociationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RetryVpcPrefixListAssociation(request)
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

// RetryVpcPrefixListAssociationWithCallback invokes the vpc.RetryVpcPrefixListAssociation API asynchronously
func (client *Client) RetryVpcPrefixListAssociationWithCallback(request *RetryVpcPrefixListAssociationRequest, callback func(response *RetryVpcPrefixListAssociationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RetryVpcPrefixListAssociationResponse
		var err error
		defer close(result)
		response, err = client.RetryVpcPrefixListAssociation(request)
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

// RetryVpcPrefixListAssociationRequest is the request struct for api RetryVpcPrefixListAssociation
type RetryVpcPrefixListAssociationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	PrefixListId         string           `position:"Query" name:"PrefixListId"`
	ResourceId           string           `position:"Query" name:"ResourceId"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
}

// RetryVpcPrefixListAssociationResponse is the response struct for api RetryVpcPrefixListAssociation
type RetryVpcPrefixListAssociationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRetryVpcPrefixListAssociationRequest creates a request to invoke RetryVpcPrefixListAssociation API
func CreateRetryVpcPrefixListAssociationRequest() (request *RetryVpcPrefixListAssociationRequest) {
	request = &RetryVpcPrefixListAssociationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "RetryVpcPrefixListAssociation", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRetryVpcPrefixListAssociationResponse creates a response to parse from RetryVpcPrefixListAssociation response
func CreateRetryVpcPrefixListAssociationResponse() (response *RetryVpcPrefixListAssociationResponse) {
	response = &RetryVpcPrefixListAssociationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
