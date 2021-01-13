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

// PurchaseStorageCapacityUnit invokes the ecs.PurchaseStorageCapacityUnit API synchronously
func (client *Client) PurchaseStorageCapacityUnit(request *PurchaseStorageCapacityUnitRequest) (response *PurchaseStorageCapacityUnitResponse, err error) {
	response = CreatePurchaseStorageCapacityUnitResponse()
	err = client.DoAction(request, response)
	return
}

// PurchaseStorageCapacityUnitWithChan invokes the ecs.PurchaseStorageCapacityUnit API asynchronously
func (client *Client) PurchaseStorageCapacityUnitWithChan(request *PurchaseStorageCapacityUnitRequest) (<-chan *PurchaseStorageCapacityUnitResponse, <-chan error) {
	responseChan := make(chan *PurchaseStorageCapacityUnitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PurchaseStorageCapacityUnit(request)
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

// PurchaseStorageCapacityUnitWithCallback invokes the ecs.PurchaseStorageCapacityUnit API asynchronously
func (client *Client) PurchaseStorageCapacityUnitWithCallback(request *PurchaseStorageCapacityUnitRequest, callback func(response *PurchaseStorageCapacityUnitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PurchaseStorageCapacityUnitResponse
		var err error
		defer close(result)
		response, err = client.PurchaseStorageCapacityUnit(request)
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

// PurchaseStorageCapacityUnitRequest is the request struct for api PurchaseStorageCapacityUnit
type PurchaseStorageCapacityUnitRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	StartTime            string           `position:"Query" name:"StartTime"`
	Capacity             requests.Integer `position:"Query" name:"Capacity"`
	Period               requests.Integer `position:"Query" name:"Period"`
	Amount               requests.Integer `position:"Query" name:"Amount"`
	FromApp              string           `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PeriodUnit           string           `position:"Query" name:"PeriodUnit"`
	Name                 string           `position:"Query" name:"Name"`
}

// PurchaseStorageCapacityUnitResponse is the response struct for api PurchaseStorageCapacityUnit
type PurchaseStorageCapacityUnitResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	OrderId                string                 `json:"OrderId" xml:"OrderId"`
	StorageCapacityUnitIds StorageCapacityUnitIds `json:"StorageCapacityUnitIds" xml:"StorageCapacityUnitIds"`
}

// CreatePurchaseStorageCapacityUnitRequest creates a request to invoke PurchaseStorageCapacityUnit API
func CreatePurchaseStorageCapacityUnitRequest() (request *PurchaseStorageCapacityUnitRequest) {
	request = &PurchaseStorageCapacityUnitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "PurchaseStorageCapacityUnit", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePurchaseStorageCapacityUnitResponse creates a response to parse from PurchaseStorageCapacityUnit response
func CreatePurchaseStorageCapacityUnitResponse() (response *PurchaseStorageCapacityUnitResponse) {
	response = &PurchaseStorageCapacityUnitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
