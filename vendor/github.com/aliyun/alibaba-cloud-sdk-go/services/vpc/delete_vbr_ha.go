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

// DeleteVbrHa invokes the vpc.DeleteVbrHa API synchronously
func (client *Client) DeleteVbrHa(request *DeleteVbrHaRequest) (response *DeleteVbrHaResponse, err error) {
	response = CreateDeleteVbrHaResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVbrHaWithChan invokes the vpc.DeleteVbrHa API asynchronously
func (client *Client) DeleteVbrHaWithChan(request *DeleteVbrHaRequest) (<-chan *DeleteVbrHaResponse, <-chan error) {
	responseChan := make(chan *DeleteVbrHaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVbrHa(request)
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

// DeleteVbrHaWithCallback invokes the vpc.DeleteVbrHa API asynchronously
func (client *Client) DeleteVbrHaWithCallback(request *DeleteVbrHaRequest, callback func(response *DeleteVbrHaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVbrHaResponse
		var err error
		defer close(result)
		response, err = client.DeleteVbrHa(request)
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

// DeleteVbrHaRequest is the request struct for api DeleteVbrHa
type DeleteVbrHaRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DeleteVbrHaResponse is the response struct for api DeleteVbrHa
type DeleteVbrHaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVbrHaRequest creates a request to invoke DeleteVbrHa API
func CreateDeleteVbrHaRequest() (request *DeleteVbrHaRequest) {
	request = &DeleteVbrHaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVbrHa", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVbrHaResponse creates a response to parse from DeleteVbrHa response
func CreateDeleteVbrHaResponse() (response *DeleteVbrHaResponse) {
	response = &DeleteVbrHaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
