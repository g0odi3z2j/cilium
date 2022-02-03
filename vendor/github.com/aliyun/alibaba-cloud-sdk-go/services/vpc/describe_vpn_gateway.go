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

// DescribeVpnGateway invokes the vpc.DescribeVpnGateway API synchronously
func (client *Client) DescribeVpnGateway(request *DescribeVpnGatewayRequest) (response *DescribeVpnGatewayResponse, err error) {
	response = CreateDescribeVpnGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVpnGatewayWithChan invokes the vpc.DescribeVpnGateway API asynchronously
func (client *Client) DescribeVpnGatewayWithChan(request *DescribeVpnGatewayRequest) (<-chan *DescribeVpnGatewayResponse, <-chan error) {
	responseChan := make(chan *DescribeVpnGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpnGateway(request)
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

// DescribeVpnGatewayWithCallback invokes the vpc.DescribeVpnGateway API asynchronously
func (client *Client) DescribeVpnGatewayWithCallback(request *DescribeVpnGatewayRequest, callback func(response *DescribeVpnGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpnGatewayResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpnGateway(request)
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

// DescribeVpnGatewayRequest is the request struct for api DescribeVpnGateway
type DescribeVpnGatewayRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	IncludeReservationData requests.Boolean `position:"Query" name:"IncludeReservationData"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	VpnGatewayId           string           `position:"Query" name:"VpnGatewayId"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVpnGatewayResponse is the response struct for api DescribeVpnGateway
type DescribeVpnGatewayResponse struct {
	*responses.BaseResponse
	VpnType           string                   `json:"VpnType" xml:"VpnType"`
	Status            string                   `json:"Status" xml:"Status"`
	VpcId             string                   `json:"VpcId" xml:"VpcId"`
	SslMaxConnections int64                    `json:"SslMaxConnections" xml:"SslMaxConnections"`
	Spec              string                   `json:"Spec" xml:"Spec"`
	InternetIp        string                   `json:"InternetIp" xml:"InternetIp"`
	CreateTime        int64                    `json:"CreateTime" xml:"CreateTime"`
	AutoPropagate     bool                     `json:"AutoPropagate" xml:"AutoPropagate"`
	ChargeType        string                   `json:"ChargeType" xml:"ChargeType"`
	VpnGatewayId      string                   `json:"VpnGatewayId" xml:"VpnGatewayId"`
	Tag               string                   `json:"Tag" xml:"Tag"`
	IpsecVpn          string                   `json:"IpsecVpn" xml:"IpsecVpn"`
	EndTime           int64                    `json:"EndTime" xml:"EndTime"`
	VSwitchId         string                   `json:"VSwitchId" xml:"VSwitchId"`
	RequestId         string                   `json:"RequestId" xml:"RequestId"`
	Description       string                   `json:"Description" xml:"Description"`
	EnableBgp         bool                     `json:"EnableBgp" xml:"EnableBgp"`
	BusinessStatus    string                   `json:"BusinessStatus" xml:"BusinessStatus"`
	SslVpn            string                   `json:"SslVpn" xml:"SslVpn"`
	Name              string                   `json:"Name" xml:"Name"`
	ReservationData   ReservationData          `json:"ReservationData" xml:"ReservationData"`
	Tags              TagsInDescribeVpnGateway `json:"Tags" xml:"Tags"`
}

// CreateDescribeVpnGatewayRequest creates a request to invoke DescribeVpnGateway API
func CreateDescribeVpnGatewayRequest() (request *DescribeVpnGatewayRequest) {
	request = &DescribeVpnGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnGateway", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVpnGatewayResponse creates a response to parse from DescribeVpnGateway response
func CreateDescribeVpnGatewayResponse() (response *DescribeVpnGatewayResponse) {
	response = &DescribeVpnGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
