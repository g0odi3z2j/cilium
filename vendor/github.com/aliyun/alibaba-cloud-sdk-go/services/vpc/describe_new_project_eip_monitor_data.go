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

// DescribeNewProjectEipMonitorData invokes the vpc.DescribeNewProjectEipMonitorData API synchronously
func (client *Client) DescribeNewProjectEipMonitorData(request *DescribeNewProjectEipMonitorDataRequest) (response *DescribeNewProjectEipMonitorDataResponse, err error) {
	response = CreateDescribeNewProjectEipMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNewProjectEipMonitorDataWithChan invokes the vpc.DescribeNewProjectEipMonitorData API asynchronously
func (client *Client) DescribeNewProjectEipMonitorDataWithChan(request *DescribeNewProjectEipMonitorDataRequest) (<-chan *DescribeNewProjectEipMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeNewProjectEipMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNewProjectEipMonitorData(request)
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

// DescribeNewProjectEipMonitorDataWithCallback invokes the vpc.DescribeNewProjectEipMonitorData API asynchronously
func (client *Client) DescribeNewProjectEipMonitorDataWithCallback(request *DescribeNewProjectEipMonitorDataRequest, callback func(response *DescribeNewProjectEipMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNewProjectEipMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeNewProjectEipMonitorData(request)
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

// DescribeNewProjectEipMonitorDataRequest is the request struct for api DescribeNewProjectEipMonitorData
type DescribeNewProjectEipMonitorDataRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AllocationId         string           `position:"Query" name:"AllocationId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	Period               requests.Integer `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeNewProjectEipMonitorDataResponse is the response struct for api DescribeNewProjectEipMonitorData
type DescribeNewProjectEipMonitorDataResponse struct {
	*responses.BaseResponse
	RequestId       string                                            `json:"RequestId" xml:"RequestId"`
	EipMonitorDatas EipMonitorDatasInDescribeNewProjectEipMonitorData `json:"EipMonitorDatas" xml:"EipMonitorDatas"`
}

// CreateDescribeNewProjectEipMonitorDataRequest creates a request to invoke DescribeNewProjectEipMonitorData API
func CreateDescribeNewProjectEipMonitorDataRequest() (request *DescribeNewProjectEipMonitorDataRequest) {
	request = &DescribeNewProjectEipMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeNewProjectEipMonitorData", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNewProjectEipMonitorDataResponse creates a response to parse from DescribeNewProjectEipMonitorData response
func CreateDescribeNewProjectEipMonitorDataResponse() (response *DescribeNewProjectEipMonitorDataResponse) {
	response = &DescribeNewProjectEipMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
