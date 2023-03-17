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

// HaVip is a nested struct in vpc response
type HaVip struct {
	VpcId                  string                 `json:"VpcId" xml:"VpcId"`
	Status                 string                 `json:"Status" xml:"Status"`
	HaVipId                string                 `json:"HaVipId" xml:"HaVipId"`
	AssociatedInstanceType string                 `json:"AssociatedInstanceType" xml:"AssociatedInstanceType"`
	CreateTime             string                 `json:"CreateTime" xml:"CreateTime"`
	ChargeType             string                 `json:"ChargeType" xml:"ChargeType"`
	RegionId               string                 `json:"RegionId" xml:"RegionId"`
	VSwitchId              string                 `json:"VSwitchId" xml:"VSwitchId"`
	IpAddress              string                 `json:"IpAddress" xml:"IpAddress"`
	Description            string                 `json:"Description" xml:"Description"`
	MasterInstanceId       string                 `json:"MasterInstanceId" xml:"MasterInstanceId"`
	Name                   string                 `json:"Name" xml:"Name"`
	ResourceGroupId        string                 `json:"ResourceGroupId" xml:"ResourceGroupId"`
	AssociatedEipAddresses AssociatedEipAddresses `json:"AssociatedEipAddresses" xml:"AssociatedEipAddresses"`
	AssociatedInstances    AssociatedInstances    `json:"AssociatedInstances" xml:"AssociatedInstances"`
	Tags                   TagsInDescribeHaVips   `json:"Tags" xml:"Tags"`
}
