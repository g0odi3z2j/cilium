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

// VpnConnection is a nested struct in vpc response
type VpnConnection struct {
	Status                     string                                             `json:"Status" xml:"Status"`
	EnableNatTraversal         bool                                               `json:"EnableNatTraversal" xml:"EnableNatTraversal"`
	RemoteCaCertificate        string                                             `json:"RemoteCaCertificate" xml:"RemoteCaCertificate"`
	CreateTime                 int64                                              `json:"CreateTime" xml:"CreateTime"`
	EffectImmediately          bool                                               `json:"EffectImmediately" xml:"EffectImmediately"`
	VpnGatewayId               string                                             `json:"VpnGatewayId" xml:"VpnGatewayId"`
	LocalSubnet                string                                             `json:"LocalSubnet" xml:"LocalSubnet"`
	VpnConnectionId            string                                             `json:"VpnConnectionId" xml:"VpnConnectionId"`
	RemoteSubnet               string                                             `json:"RemoteSubnet" xml:"RemoteSubnet"`
	CustomerGatewayId          string                                             `json:"CustomerGatewayId" xml:"CustomerGatewayId"`
	Name                       string                                             `json:"Name" xml:"Name"`
	EnableDpd                  bool                                               `json:"EnableDpd" xml:"EnableDpd"`
	AttachType                 string                                             `json:"AttachType" xml:"AttachType"`
	NetworkType                string                                             `json:"NetworkType" xml:"NetworkType"`
	AttachInstanceId           string                                             `json:"AttachInstanceId" xml:"AttachInstanceId"`
	Spec                       string                                             `json:"Spec" xml:"Spec"`
	State                      string                                             `json:"State" xml:"State"`
	TransitRouterId            string                                             `json:"TransitRouterId" xml:"TransitRouterId"`
	TransitRouterName          string                                             `json:"TransitRouterName" xml:"TransitRouterName"`
	CrossAccountAuthorized     bool                                               `json:"CrossAccountAuthorized" xml:"CrossAccountAuthorized"`
	InternetIp                 string                                             `json:"InternetIp" xml:"InternetIp"`
	IkeConfig                  IkeConfig                                          `json:"IkeConfig" xml:"IkeConfig"`
	IpsecConfig                IpsecConfig                                        `json:"IpsecConfig" xml:"IpsecConfig"`
	VcoHealthCheck             VcoHealthCheck                                     `json:"VcoHealthCheck" xml:"VcoHealthCheck"`
	VpnBgpConfig               VpnBgpConfig                                       `json:"VpnBgpConfig" xml:"VpnBgpConfig"`
	Tag                        TagInDescribeVpnConnections                        `json:"Tag" xml:"Tag"`
	TunnelOptionsSpecification TunnelOptionsSpecificationInDescribeVpnConnections `json:"TunnelOptionsSpecification" xml:"TunnelOptionsSpecification"`
}
