package ecs

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

// DeregisterManagedInstance invokes the ecs.DeregisterManagedInstance API synchronously
func (client *Client) DeregisterManagedInstance(request *DeregisterManagedInstanceRequest) (response *DeregisterManagedInstanceResponse, err error) {
	response = CreateDeregisterManagedInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DeregisterManagedInstanceWithChan invokes the ecs.DeregisterManagedInstance API asynchronously
func (client *Client) DeregisterManagedInstanceWithChan(request *DeregisterManagedInstanceRequest) (<-chan *DeregisterManagedInstanceResponse, <-chan error) {
	responseChan := make(chan *DeregisterManagedInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeregisterManagedInstance(request)
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

// DeregisterManagedInstanceWithCallback invokes the ecs.DeregisterManagedInstance API asynchronously
func (client *Client) DeregisterManagedInstanceWithCallback(request *DeregisterManagedInstanceRequest, callback func(response *DeregisterManagedInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeregisterManagedInstanceResponse
		var err error
		defer close(result)
		response, err = client.DeregisterManagedInstance(request)
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

// DeregisterManagedInstanceRequest is the request struct for api DeregisterManagedInstance
type DeregisterManagedInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DeregisterManagedInstanceResponse is the response struct for api DeregisterManagedInstance
type DeregisterManagedInstanceResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Instance  Instance `json:"Instance" xml:"Instance"`
}

// CreateDeregisterManagedInstanceRequest creates a request to invoke DeregisterManagedInstance API
func CreateDeregisterManagedInstanceRequest() (request *DeregisterManagedInstanceRequest) {
	request = &DeregisterManagedInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeregisterManagedInstance", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeregisterManagedInstanceResponse creates a response to parse from DeregisterManagedInstance response
func CreateDeregisterManagedInstanceResponse() (response *DeregisterManagedInstanceResponse) {
	response = &DeregisterManagedInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
