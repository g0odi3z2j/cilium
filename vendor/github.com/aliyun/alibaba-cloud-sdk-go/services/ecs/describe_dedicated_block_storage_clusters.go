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

// DescribeDedicatedBlockStorageClusters invokes the ecs.DescribeDedicatedBlockStorageClusters API synchronously
func (client *Client) DescribeDedicatedBlockStorageClusters(request *DescribeDedicatedBlockStorageClustersRequest) (response *DescribeDedicatedBlockStorageClustersResponse, err error) {
	response = CreateDescribeDedicatedBlockStorageClustersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDedicatedBlockStorageClustersWithChan invokes the ecs.DescribeDedicatedBlockStorageClusters API asynchronously
func (client *Client) DescribeDedicatedBlockStorageClustersWithChan(request *DescribeDedicatedBlockStorageClustersRequest) (<-chan *DescribeDedicatedBlockStorageClustersResponse, <-chan error) {
	responseChan := make(chan *DescribeDedicatedBlockStorageClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDedicatedBlockStorageClusters(request)
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

// DescribeDedicatedBlockStorageClustersWithCallback invokes the ecs.DescribeDedicatedBlockStorageClusters API asynchronously
func (client *Client) DescribeDedicatedBlockStorageClustersWithCallback(request *DescribeDedicatedBlockStorageClustersRequest, callback func(response *DescribeDedicatedBlockStorageClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDedicatedBlockStorageClustersResponse
		var err error
		defer close(result)
		response, err = client.DescribeDedicatedBlockStorageClusters(request)
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

// DescribeDedicatedBlockStorageClustersRequest is the request struct for api DescribeDedicatedBlockStorageClusters
type DescribeDedicatedBlockStorageClustersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NextToken                      string           `position:"Query" name:"NextToken"`
	DedicatedBlockStorageClusterId *[]string        `position:"Query" name:"DedicatedBlockStorageClusterId"  type:"Repeated"`
	ResourceOwnerAccount           string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                   string           `position:"Query" name:"OwnerAccount"`
	OwnerId                        requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId                         string           `position:"Query" name:"ZoneId"`
	MaxResults                     requests.Integer `position:"Query" name:"MaxResults"`
	Category                       string           `position:"Query" name:"Category"`
	Status                         *[]string        `position:"Query" name:"Status"  type:"Repeated"`
}

// DescribeDedicatedBlockStorageClustersResponse is the response struct for api DescribeDedicatedBlockStorageClusters
type DescribeDedicatedBlockStorageClustersResponse struct {
	*responses.BaseResponse
	RequestId                     string                        `json:"RequestId" xml:"RequestId"`
	NextToken                     string                        `json:"NextToken" xml:"NextToken"`
	DedicatedBlockStorageClusters DedicatedBlockStorageClusters `json:"DedicatedBlockStorageClusters" xml:"DedicatedBlockStorageClusters"`
}

// CreateDescribeDedicatedBlockStorageClustersRequest creates a request to invoke DescribeDedicatedBlockStorageClusters API
func CreateDescribeDedicatedBlockStorageClustersRequest() (request *DescribeDedicatedBlockStorageClustersRequest) {
	request = &DescribeDedicatedBlockStorageClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDedicatedBlockStorageClusters", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDedicatedBlockStorageClustersResponse creates a response to parse from DescribeDedicatedBlockStorageClusters response
func CreateDescribeDedicatedBlockStorageClustersResponse() (response *DescribeDedicatedBlockStorageClustersResponse) {
	response = &DescribeDedicatedBlockStorageClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
