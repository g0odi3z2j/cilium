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

// DescribeNetworkInterfaceAttribute invokes the ecs.DescribeNetworkInterfaceAttribute API synchronously
func (client *Client) DescribeNetworkInterfaceAttribute(request *DescribeNetworkInterfaceAttributeRequest) (response *DescribeNetworkInterfaceAttributeResponse, err error) {
	response = CreateDescribeNetworkInterfaceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkInterfaceAttributeWithChan invokes the ecs.DescribeNetworkInterfaceAttribute API asynchronously
func (client *Client) DescribeNetworkInterfaceAttributeWithChan(request *DescribeNetworkInterfaceAttributeRequest) (<-chan *DescribeNetworkInterfaceAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkInterfaceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkInterfaceAttribute(request)
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

// DescribeNetworkInterfaceAttributeWithCallback invokes the ecs.DescribeNetworkInterfaceAttribute API asynchronously
func (client *Client) DescribeNetworkInterfaceAttributeWithCallback(request *DescribeNetworkInterfaceAttributeRequest, callback func(response *DescribeNetworkInterfaceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkInterfaceAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkInterfaceAttribute(request)
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

// DescribeNetworkInterfaceAttributeRequest is the request struct for api DescribeNetworkInterfaceAttribute
type DescribeNetworkInterfaceAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                        `position:"Query" name:"ResourceOwnerId"`
	Tag                  *[]DescribeNetworkInterfaceAttributeTag `position:"Query" name:"Tag"  type:"Repeated"`
	Attribute            string                                  `position:"Query" name:"Attribute"`
	ResourceOwnerAccount string                                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                  `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                        `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string                                  `position:"Query" name:"NetworkInterfaceId"`
}

// DescribeNetworkInterfaceAttributeTag is a repeated param struct in DescribeNetworkInterfaceAttributeRequest
type DescribeNetworkInterfaceAttributeTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeNetworkInterfaceAttributeResponse is the response struct for api DescribeNetworkInterfaceAttribute
type DescribeNetworkInterfaceAttributeResponse struct {
	*responses.BaseResponse
	RequestId                   string                                              `json:"RequestId" xml:"RequestId"`
	NetworkInterfaceId          string                                              `json:"NetworkInterfaceId" xml:"NetworkInterfaceId"`
	Status                      string                                              `json:"Status" xml:"Status"`
	Type                        string                                              `json:"Type" xml:"Type"`
	VpcId                       string                                              `json:"VpcId" xml:"VpcId"`
	VSwitchId                   string                                              `json:"VSwitchId" xml:"VSwitchId"`
	ZoneId                      string                                              `json:"ZoneId" xml:"ZoneId"`
	PrivateIpAddress            string                                              `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
	MacAddress                  string                                              `json:"MacAddress" xml:"MacAddress"`
	NetworkInterfaceName        string                                              `json:"NetworkInterfaceName" xml:"NetworkInterfaceName"`
	Description                 string                                              `json:"Description" xml:"Description"`
	InstanceId                  string                                              `json:"InstanceId" xml:"InstanceId"`
	CreationTime                string                                              `json:"CreationTime" xml:"CreationTime"`
	ResourceGroupId             string                                              `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ServiceID                   int64                                               `json:"ServiceID" xml:"ServiceID"`
	ServiceManaged              bool                                                `json:"ServiceManaged" xml:"ServiceManaged"`
	QueueNumber                 int                                                 `json:"QueueNumber" xml:"QueueNumber"`
	OwnerId                     string                                              `json:"OwnerId" xml:"OwnerId"`
	NetworkInterfaceTrafficMode string                                              `json:"NetworkInterfaceTrafficMode" xml:"NetworkInterfaceTrafficMode"`
	QueuePairNumber             int                                                 `json:"QueuePairNumber" xml:"QueuePairNumber"`
	SecurityGroupIds            SecurityGroupIdsInDescribeNetworkInterfaceAttribute `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	AssociatedPublicIp          AssociatedPublicIp                                  `json:"AssociatedPublicIp" xml:"AssociatedPublicIp"`
	Attachment                  Attachment                                          `json:"Attachment" xml:"Attachment"`
	PrivateIpSets               PrivateIpSetsInDescribeNetworkInterfaceAttribute    `json:"PrivateIpSets" xml:"PrivateIpSets"`
	Ipv6Sets                    Ipv6SetsInDescribeNetworkInterfaceAttribute         `json:"Ipv6Sets" xml:"Ipv6Sets"`
	Tags                        TagsInDescribeNetworkInterfaceAttribute             `json:"Tags" xml:"Tags"`
}

// CreateDescribeNetworkInterfaceAttributeRequest creates a request to invoke DescribeNetworkInterfaceAttribute API
func CreateDescribeNetworkInterfaceAttributeRequest() (request *DescribeNetworkInterfaceAttributeRequest) {
	request = &DescribeNetworkInterfaceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNetworkInterfaceAttribute", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkInterfaceAttributeResponse creates a response to parse from DescribeNetworkInterfaceAttribute response
func CreateDescribeNetworkInterfaceAttributeResponse() (response *DescribeNetworkInterfaceAttributeResponse) {
	response = &DescribeNetworkInterfaceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
