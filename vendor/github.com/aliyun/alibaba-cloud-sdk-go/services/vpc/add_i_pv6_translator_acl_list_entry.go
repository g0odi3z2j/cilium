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

// AddIPv6TranslatorAclListEntry invokes the vpc.AddIPv6TranslatorAclListEntry API synchronously
func (client *Client) AddIPv6TranslatorAclListEntry(request *AddIPv6TranslatorAclListEntryRequest) (response *AddIPv6TranslatorAclListEntryResponse, err error) {
	response = CreateAddIPv6TranslatorAclListEntryResponse()
	err = client.DoAction(request, response)
	return
}

// AddIPv6TranslatorAclListEntryWithChan invokes the vpc.AddIPv6TranslatorAclListEntry API asynchronously
func (client *Client) AddIPv6TranslatorAclListEntryWithChan(request *AddIPv6TranslatorAclListEntryRequest) (<-chan *AddIPv6TranslatorAclListEntryResponse, <-chan error) {
	responseChan := make(chan *AddIPv6TranslatorAclListEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddIPv6TranslatorAclListEntry(request)
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

// AddIPv6TranslatorAclListEntryWithCallback invokes the vpc.AddIPv6TranslatorAclListEntry API asynchronously
func (client *Client) AddIPv6TranslatorAclListEntryWithCallback(request *AddIPv6TranslatorAclListEntryRequest, callback func(response *AddIPv6TranslatorAclListEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddIPv6TranslatorAclListEntryResponse
		var err error
		defer close(result)
		response, err = client.AddIPv6TranslatorAclListEntry(request)
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

// AddIPv6TranslatorAclListEntryRequest is the request struct for api AddIPv6TranslatorAclListEntry
type AddIPv6TranslatorAclListEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AclId                string           `position:"Query" name:"AclId"`
	AclEntryIp           string           `position:"Query" name:"AclEntryIp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AclEntryComment      string           `position:"Query" name:"AclEntryComment"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// AddIPv6TranslatorAclListEntryResponse is the response struct for api AddIPv6TranslatorAclListEntry
type AddIPv6TranslatorAclListEntryResponse struct {
	*responses.BaseResponse
	AclEntryId string `json:"AclEntryId" xml:"AclEntryId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateAddIPv6TranslatorAclListEntryRequest creates a request to invoke AddIPv6TranslatorAclListEntry API
func CreateAddIPv6TranslatorAclListEntryRequest() (request *AddIPv6TranslatorAclListEntryRequest) {
	request = &AddIPv6TranslatorAclListEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AddIPv6TranslatorAclListEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddIPv6TranslatorAclListEntryResponse creates a response to parse from AddIPv6TranslatorAclListEntry response
func CreateAddIPv6TranslatorAclListEntryResponse() (response *AddIPv6TranslatorAclListEntryResponse) {
	response = &AddIPv6TranslatorAclListEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
