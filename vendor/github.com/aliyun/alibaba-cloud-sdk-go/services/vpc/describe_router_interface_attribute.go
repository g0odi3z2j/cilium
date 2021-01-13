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

// DescribeRouterInterfaceAttribute invokes the vpc.DescribeRouterInterfaceAttribute API synchronously
func (client *Client) DescribeRouterInterfaceAttribute(request *DescribeRouterInterfaceAttributeRequest) (response *DescribeRouterInterfaceAttributeResponse, err error) {
	response = CreateDescribeRouterInterfaceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRouterInterfaceAttributeWithChan invokes the vpc.DescribeRouterInterfaceAttribute API asynchronously
func (client *Client) DescribeRouterInterfaceAttributeWithChan(request *DescribeRouterInterfaceAttributeRequest) (<-chan *DescribeRouterInterfaceAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeRouterInterfaceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRouterInterfaceAttribute(request)
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

// DescribeRouterInterfaceAttributeWithCallback invokes the vpc.DescribeRouterInterfaceAttribute API asynchronously
func (client *Client) DescribeRouterInterfaceAttributeWithCallback(request *DescribeRouterInterfaceAttributeRequest, callback func(response *DescribeRouterInterfaceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRouterInterfaceAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeRouterInterfaceAttribute(request)
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

// DescribeRouterInterfaceAttributeRequest is the request struct for api DescribeRouterInterfaceAttribute
type DescribeRouterInterfaceAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeRouterInterfaceAttributeResponse is the response struct for api DescribeRouterInterfaceAttribute
type DescribeRouterInterfaceAttributeResponse struct {
	*responses.BaseResponse
	RequestId                       string `json:"RequestId" xml:"RequestId"`
	Code                            string `json:"Code" xml:"Code"`
	Message                         string `json:"Message" xml:"Message"`
	Success                         bool   `json:"Success" xml:"Success"`
	RouterInterfaceId               string `json:"RouterInterfaceId" xml:"RouterInterfaceId"`
	OppositeRegionId                string `json:"OppositeRegionId" xml:"OppositeRegionId"`
	Role                            string `json:"Role" xml:"Role"`
	Spec                            string `json:"Spec" xml:"Spec"`
	Name                            string `json:"Name" xml:"Name"`
	Description                     string `json:"Description" xml:"Description"`
	RouterId                        string `json:"RouterId" xml:"RouterId"`
	RouterType                      string `json:"RouterType" xml:"RouterType"`
	CreationTime                    string `json:"CreationTime" xml:"CreationTime"`
	GmtModified                     string `json:"GmtModified" xml:"GmtModified"`
	EndTime                         string `json:"EndTime" xml:"EndTime"`
	ChargeType                      string `json:"ChargeType" xml:"ChargeType"`
	Status                          string `json:"Status" xml:"Status"`
	BusinessStatus                  string `json:"BusinessStatus" xml:"BusinessStatus"`
	ConnectedTime                   string `json:"ConnectedTime" xml:"ConnectedTime"`
	OppositeInterfaceId             string `json:"OppositeInterfaceId" xml:"OppositeInterfaceId"`
	OppositeInterfaceSpec           string `json:"OppositeInterfaceSpec" xml:"OppositeInterfaceSpec"`
	OppositeInterfaceStatus         string `json:"OppositeInterfaceStatus" xml:"OppositeInterfaceStatus"`
	OppositeInterfaceBusinessStatus string `json:"OppositeInterfaceBusinessStatus" xml:"OppositeInterfaceBusinessStatus"`
	OppositeRouterId                string `json:"OppositeRouterId" xml:"OppositeRouterId"`
	OppositeRouterType              string `json:"OppositeRouterType" xml:"OppositeRouterType"`
	OppositeInterfaceOwnerId        string `json:"OppositeInterfaceOwnerId" xml:"OppositeInterfaceOwnerId"`
	AccessPointId                   string `json:"AccessPointId" xml:"AccessPointId"`
	OppositeAccessPointId           string `json:"OppositeAccessPointId" xml:"OppositeAccessPointId"`
	HealthCheckSourceIp             string `json:"HealthCheckSourceIp" xml:"HealthCheckSourceIp"`
	HealthCheckTargetIp             string `json:"HealthCheckTargetIp" xml:"HealthCheckTargetIp"`
	OppositeVpcInstanceId           string `json:"OppositeVpcInstanceId" xml:"OppositeVpcInstanceId"`
	Bandwidth                       int    `json:"Bandwidth" xml:"Bandwidth"`
	VpcInstanceId                   string `json:"VpcInstanceId" xml:"VpcInstanceId"`
	OppositeBandwidth               int    `json:"OppositeBandwidth" xml:"OppositeBandwidth"`
	HasReservationData              string `json:"HasReservationData" xml:"HasReservationData"`
	ReservationBandwidth            string `json:"ReservationBandwidth" xml:"ReservationBandwidth"`
	ReservationInternetChargeType   string `json:"ReservationInternetChargeType" xml:"ReservationInternetChargeType"`
	ReservationActiveTime           string `json:"ReservationActiveTime" xml:"ReservationActiveTime"`
	ReservationOrderType            string `json:"ReservationOrderType" xml:"ReservationOrderType"`
	CrossBorder                     bool   `json:"CrossBorder" xml:"CrossBorder"`
	HcThreshold                     int    `json:"HcThreshold" xml:"HcThreshold"`
	HcRate                          int    `json:"HcRate" xml:"HcRate"`
	HealthCheckStatus               string `json:"HealthCheckStatus" xml:"HealthCheckStatus"`
}

// CreateDescribeRouterInterfaceAttributeRequest creates a request to invoke DescribeRouterInterfaceAttribute API
func CreateDescribeRouterInterfaceAttributeRequest() (request *DescribeRouterInterfaceAttributeRequest) {
	request = &DescribeRouterInterfaceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouterInterfaceAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRouterInterfaceAttributeResponse creates a response to parse from DescribeRouterInterfaceAttribute response
func CreateDescribeRouterInterfaceAttributeResponse() (response *DescribeRouterInterfaceAttributeResponse) {
	response = &DescribeRouterInterfaceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
