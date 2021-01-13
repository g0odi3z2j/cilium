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

// Vpc is a nested struct in vpc response
type Vpc struct {
	VpcId                  string                            `json:"VpcId" xml:"VpcId"`
	RegionId               string                            `json:"RegionId" xml:"RegionId"`
	Status                 string                            `json:"Status" xml:"Status"`
	VpcName                string                            `json:"VpcName" xml:"VpcName"`
	CreationTime           string                            `json:"CreationTime" xml:"CreationTime"`
	CidrBlock              string                            `json:"CidrBlock" xml:"CidrBlock"`
	Ipv6CidrBlock          string                            `json:"Ipv6CidrBlock" xml:"Ipv6CidrBlock"`
	VRouterId              string                            `json:"VRouterId" xml:"VRouterId"`
	Description            string                            `json:"Description" xml:"Description"`
	IsDefault              bool                              `json:"IsDefault" xml:"IsDefault"`
	NetworkAclNum          string                            `json:"NetworkAclNum" xml:"NetworkAclNum"`
	ResourceGroupId        string                            `json:"ResourceGroupId" xml:"ResourceGroupId"`
	CenStatus              string                            `json:"CenStatus" xml:"CenStatus"`
	OwnerId                int64                             `json:"OwnerId" xml:"OwnerId"`
	SupportAdvancedFeature bool                              `json:"SupportAdvancedFeature" xml:"SupportAdvancedFeature"`
	AdvancedResource       bool                              `json:"AdvancedResource" xml:"AdvancedResource"`
	DhcpOptionsSetId       string                            `json:"DhcpOptionsSetId" xml:"DhcpOptionsSetId"`
	DhcpOptionsSetStatus   string                            `json:"DhcpOptionsSetStatus" xml:"DhcpOptionsSetStatus"`
	VSwitchIds             VSwitchIdsInDescribeVpcs          `json:"VSwitchIds" xml:"VSwitchIds"`
	UserCidrs              UserCidrsInDescribeVpcs           `json:"UserCidrs" xml:"UserCidrs"`
	NatGatewayIds          NatGatewayIds                     `json:"NatGatewayIds" xml:"NatGatewayIds"`
	RouterTableIds         RouterTableIds                    `json:"RouterTableIds" xml:"RouterTableIds"`
	SecondaryCidrBlocks    SecondaryCidrBlocksInDescribeVpcs `json:"SecondaryCidrBlocks" xml:"SecondaryCidrBlocks"`
	Tags                   TagsInDescribeVpcs                `json:"Tags" xml:"Tags"`
}
