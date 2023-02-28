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

// SendFile invokes the ecs.SendFile API synchronously
func (client *Client) SendFile(request *SendFileRequest) (response *SendFileResponse, err error) {
	response = CreateSendFileResponse()
	err = client.DoAction(request, response)
	return
}

// SendFileWithChan invokes the ecs.SendFile API asynchronously
func (client *Client) SendFileWithChan(request *SendFileRequest) (<-chan *SendFileResponse, <-chan error) {
	responseChan := make(chan *SendFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendFile(request)
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

// SendFileWithCallback invokes the ecs.SendFile API asynchronously
func (client *Client) SendFileWithCallback(request *SendFileRequest, callback func(response *SendFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendFileResponse
		var err error
		defer close(result)
		response, err = client.SendFile(request)
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

// SendFileRequest is the request struct for api SendFile
type SendFileRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	Timeout              requests.Integer `position:"Query" name:"Timeout"`
	Content              string           `position:"Query" name:"Content"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	FileOwner            string           `position:"Query" name:"FileOwner"`
	Tag                  *[]SendFileTag   `position:"Query" name:"Tag"  type:"Repeated"`
	Overwrite            requests.Boolean `position:"Query" name:"Overwrite"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	FileMode             string           `position:"Query" name:"FileMode"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ContentType          string           `position:"Query" name:"ContentType"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
	Name                 string           `position:"Query" name:"Name"`
	FileGroup            string           `position:"Query" name:"FileGroup"`
	TargetDir            string           `position:"Query" name:"TargetDir"`
}

// SendFileTag is a repeated param struct in SendFileRequest
type SendFileTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// SendFileResponse is the response struct for api SendFile
type SendFileResponse struct {
	*responses.BaseResponse
	InvokeId  string `json:"InvokeId" xml:"InvokeId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSendFileRequest creates a request to invoke SendFile API
func CreateSendFileRequest() (request *SendFileRequest) {
	request = &SendFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "SendFile", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSendFileResponse creates a response to parse from SendFile response
func CreateSendFileResponse() (response *SendFileResponse) {
	response = &SendFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
