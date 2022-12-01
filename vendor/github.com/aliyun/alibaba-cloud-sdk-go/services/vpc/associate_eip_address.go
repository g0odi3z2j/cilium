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

// AssociateEipAddress invokes the vpc.AssociateEipAddress API synchronously
func (client *Client) AssociateEipAddress(request *AssociateEipAddressRequest) (response *AssociateEipAddressResponse, err error) {
	response = CreateAssociateEipAddressResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateEipAddressWithChan invokes the vpc.AssociateEipAddress API asynchronously
func (client *Client) AssociateEipAddressWithChan(request *AssociateEipAddressRequest) (<-chan *AssociateEipAddressResponse, <-chan error) {
	responseChan := make(chan *AssociateEipAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateEipAddress(request)
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

// AssociateEipAddressWithCallback invokes the vpc.AssociateEipAddress API asynchronously
func (client *Client) AssociateEipAddressWithCallback(request *AssociateEipAddressRequest, callback func(response *AssociateEipAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateEipAddressResponse
		var err error
		defer close(result)
		response, err = client.AssociateEipAddress(request)
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

// AssociateEipAddressRequest is the request struct for api AssociateEipAddress
type AssociateEipAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	AllocationId         string           `position:"Query" name:"AllocationId"`
	Mode                 string           `position:"Query" name:"Mode"`
	InstanceRegionId     string           `position:"Query" name:"InstanceRegionId"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PrivateIpAddress     string           `position:"Query" name:"PrivateIpAddress"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	VpcId                string           `position:"Query" name:"VpcId"`
}

// AssociateEipAddressResponse is the response struct for api AssociateEipAddress
type AssociateEipAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateEipAddressRequest creates a request to invoke AssociateEipAddress API
func CreateAssociateEipAddressRequest() (request *AssociateEipAddressRequest) {
	request = &AssociateEipAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AssociateEipAddress", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateEipAddressResponse creates a response to parse from AssociateEipAddress response
func CreateAssociateEipAddressResponse() (response *AssociateEipAddressResponse) {
	response = &AssociateEipAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
