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

// TunnelOptionsInModifyVpnConnectionAttribute is a nested struct in vpc response
type TunnelOptionsInModifyVpnConnectionAttribute struct {
	CustomerGatewayId   string                                          `json:"CustomerGatewayId" xml:"CustomerGatewayId"`
	EnableDpd           bool                                            `json:"EnableDpd" xml:"EnableDpd"`
	EnableNatTraversal  bool                                            `json:"EnableNatTraversal" xml:"EnableNatTraversal"`
	InternetIp          string                                          `json:"InternetIp" xml:"InternetIp"`
	RemoteCaCertificate string                                          `json:"RemoteCaCertificate" xml:"RemoteCaCertificate"`
	Role                string                                          `json:"Role" xml:"Role"`
	State               string                                          `json:"State" xml:"State"`
	Status              string                                          `json:"Status" xml:"Status"`
	TunnelId            string                                          `json:"TunnelId" xml:"TunnelId"`
	ZoneNo              string                                          `json:"ZoneNo" xml:"ZoneNo"`
	TunnelBgpConfig     TunnelBgpConfigInModifyVpnConnectionAttribute   `json:"TunnelBgpConfig" xml:"TunnelBgpConfig"`
	TunnelIkeConfig     TunnelIkeConfigInModifyVpnConnectionAttribute   `json:"TunnelIkeConfig" xml:"TunnelIkeConfig"`
	TunnelIpsecConfig   TunnelIpsecConfigInModifyVpnConnectionAttribute `json:"TunnelIpsecConfig" xml:"TunnelIpsecConfig"`
}
