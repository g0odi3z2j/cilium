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

// RouterTableListType is a nested struct in vpc response
type RouterTableListType struct {
	VpcId           string                             `json:"VpcId" xml:"VpcId"`
	CreationTime    string                             `json:"CreationTime" xml:"CreationTime"`
	Status          string                             `json:"Status" xml:"Status"`
	RouterId        string                             `json:"RouterId" xml:"RouterId"`
	AssociateType   string                             `json:"AssociateType" xml:"AssociateType"`
	RouteTableId    string                             `json:"RouteTableId" xml:"RouteTableId"`
	OwnerId         int64                              `json:"OwnerId" xml:"OwnerId"`
	Description     string                             `json:"Description" xml:"Description"`
	RouteTableType  string                             `json:"RouteTableType" xml:"RouteTableType"`
	ResourceGroupId string                             `json:"ResourceGroupId" xml:"ResourceGroupId"`
	RouterType      string                             `json:"RouterType" xml:"RouterType"`
	RouteTableName  string                             `json:"RouteTableName" xml:"RouteTableName"`
	VSwitchIds      VSwitchIdsInDescribeRouteTableList `json:"VSwitchIds" xml:"VSwitchIds"`
	GatewayIds      GatewayIds                         `json:"GatewayIds" xml:"GatewayIds"`
	Tags            TagsInDescribeRouteTableList       `json:"Tags" xml:"Tags"`
}
