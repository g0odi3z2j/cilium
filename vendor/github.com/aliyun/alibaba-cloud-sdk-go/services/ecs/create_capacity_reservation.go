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

// CreateCapacityReservation invokes the ecs.CreateCapacityReservation API synchronously
func (client *Client) CreateCapacityReservation(request *CreateCapacityReservationRequest) (response *CreateCapacityReservationResponse, err error) {
	response = CreateCreateCapacityReservationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCapacityReservationWithChan invokes the ecs.CreateCapacityReservation API asynchronously
func (client *Client) CreateCapacityReservationWithChan(request *CreateCapacityReservationRequest) (<-chan *CreateCapacityReservationResponse, <-chan error) {
	responseChan := make(chan *CreateCapacityReservationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCapacityReservation(request)
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

// CreateCapacityReservationWithCallback invokes the ecs.CreateCapacityReservation API asynchronously
func (client *Client) CreateCapacityReservationWithCallback(request *CreateCapacityReservationRequest, callback func(response *CreateCapacityReservationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCapacityReservationResponse
		var err error
		defer close(result)
		response, err = client.CreateCapacityReservation(request)
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

// CreateCapacityReservationRequest is the request struct for api CreateCapacityReservation
type CreateCapacityReservationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                 requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                     string           `position:"Query" name:"ClientToken"`
	Description                     string           `position:"Query" name:"Description"`
	StartTime                       string           `position:"Query" name:"StartTime"`
	Platform                        string           `position:"Query" name:"Platform"`
	PrivatePoolOptionsMatchCriteria string           `position:"Query" name:"PrivatePoolOptions.MatchCriteria"`
	InstanceType                    string           `position:"Query" name:"InstanceType"`
	InstanceChargeType              string           `position:"Query" name:"InstanceChargeType"`
	EfficientStatus                 requests.Integer `position:"Query" name:"EfficientStatus"`
	Period                          requests.Integer `position:"Query" name:"Period"`
	EndTimeType                     string           `position:"Query" name:"EndTimeType"`
	ResourceOwnerAccount            string           `position:"Query" name:"ResourceOwnerAccount"`
	PrivatePoolOptionsName          string           `position:"Query" name:"PrivatePoolOptions.Name"`
	OwnerAccount                    string           `position:"Query" name:"OwnerAccount"`
	EndTime                         string           `position:"Query" name:"EndTime"`
	OwnerId                         requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType                    string           `position:"Query" name:"ResourceType"`
	PeriodUnit                      string           `position:"Query" name:"PeriodUnit"`
	TimeSlot                        string           `position:"Query" name:"TimeSlot"`
	ZoneId                          *[]string        `position:"Query" name:"ZoneId"  type:"Repeated"`
	ChargeType                      string           `position:"Query" name:"ChargeType"`
	PackageType                     string           `position:"Query" name:"PackageType"`
	InstanceAmount                  requests.Integer `position:"Query" name:"InstanceAmount"`
}

// CreateCapacityReservationResponse is the response struct for api CreateCapacityReservation
type CreateCapacityReservationResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	PrivatePoolOptionsId string `json:"PrivatePoolOptionsId" xml:"PrivatePoolOptionsId"`
}

// CreateCreateCapacityReservationRequest creates a request to invoke CreateCapacityReservation API
func CreateCreateCapacityReservationRequest() (request *CreateCapacityReservationRequest) {
	request = &CreateCapacityReservationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateCapacityReservation", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCapacityReservationResponse creates a response to parse from CreateCapacityReservation response
func CreateCreateCapacityReservationResponse() (response *CreateCapacityReservationResponse) {
	response = &CreateCapacityReservationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
