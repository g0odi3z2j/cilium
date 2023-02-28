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

// CreateDisk invokes the ecs.CreateDisk API synchronously
func (client *Client) CreateDisk(request *CreateDiskRequest) (response *CreateDiskResponse, err error) {
	response = CreateCreateDiskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDiskWithChan invokes the ecs.CreateDisk API asynchronously
func (client *Client) CreateDiskWithChan(request *CreateDiskRequest) (<-chan *CreateDiskResponse, <-chan error) {
	responseChan := make(chan *CreateDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDisk(request)
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

// CreateDiskWithCallback invokes the ecs.CreateDisk API asynchronously
func (client *Client) CreateDiskWithCallback(request *CreateDiskRequest, callback func(response *CreateDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDiskResponse
		var err error
		defer close(result)
		response, err = client.CreateDisk(request)
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

// CreateDiskRequest is the request struct for api CreateDisk
type CreateDiskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer       `position:"Query" name:"ResourceOwnerId"`
	EncryptAlgorithm          string                 `position:"Query" name:"EncryptAlgorithm"`
	DiskName                  string                 `position:"Query" name:"DiskName"`
	ResourceGroupId           string                 `position:"Query" name:"ResourceGroupId"`
	StorageSetPartitionNumber requests.Integer       `position:"Query" name:"StorageSetPartitionNumber"`
	Tag                       *[]CreateDiskTag       `position:"Query" name:"Tag"  type:"Repeated"`
	OwnerId                   requests.Integer       `position:"Query" name:"OwnerId"`
	ProvisionedIops           requests.Integer       `position:"Query" name:"ProvisionedIops"`
	InstanceId                string                 `position:"Query" name:"InstanceId"`
	Size                      requests.Integer       `position:"Query" name:"Size"`
	ZoneId                    string                 `position:"Query" name:"ZoneId"`
	StorageClusterId          string                 `position:"Query" name:"StorageClusterId"`
	SnapshotId                string                 `position:"Query" name:"SnapshotId"`
	ClientToken               string                 `position:"Query" name:"ClientToken"`
	SystemTag                 *[]CreateDiskSystemTag `position:"Query" name:"SystemTag"  type:"Repeated"`
	Description               string                 `position:"Query" name:"Description"`
	DiskCategory              string                 `position:"Query" name:"DiskCategory"`
	MultiAttach               string                 `position:"Query" name:"MultiAttach"`
	AdvancedFeatures          string                 `position:"Query" name:"AdvancedFeatures"`
	Arn                       *[]CreateDiskArn       `position:"Query" name:"Arn"  type:"Repeated"`
	ResourceOwnerAccount      string                 `position:"Query" name:"ResourceOwnerAccount"`
	PerformanceLevel          string                 `position:"Query" name:"PerformanceLevel"`
	OwnerAccount              string                 `position:"Query" name:"OwnerAccount"`
	BurstingEnabled           requests.Boolean       `position:"Query" name:"BurstingEnabled"`
	StorageSetId              string                 `position:"Query" name:"StorageSetId"`
	Encrypted                 requests.Boolean       `position:"Query" name:"Encrypted"`
	KMSKeyId                  string                 `position:"Query" name:"KMSKeyId"`
}

// CreateDiskTag is a repeated param struct in CreateDiskRequest
type CreateDiskTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateDiskSystemTag is a repeated param struct in CreateDiskRequest
type CreateDiskSystemTag struct {
	Scope string `name:"Scope"`
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateDiskArn is a repeated param struct in CreateDiskRequest
type CreateDiskArn struct {
	Rolearn       string `name:"Rolearn"`
	RoleType      string `name:"RoleType"`
	AssumeRoleFor string `name:"AssumeRoleFor"`
}

// CreateDiskResponse is the response struct for api CreateDisk
type CreateDiskResponse struct {
	*responses.BaseResponse
	DiskId    string `json:"DiskId" xml:"DiskId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateDiskRequest creates a request to invoke CreateDisk API
func CreateCreateDiskRequest() (request *CreateDiskRequest) {
	request = &CreateDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateDisk", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDiskResponse creates a response to parse from CreateDisk response
func CreateCreateDiskResponse() (response *CreateDiskResponse) {
	response = &CreateDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
