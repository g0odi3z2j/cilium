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

// FullNatEntry is a nested struct in vpc response
type FullNatEntry struct {
	NatIp                   string `json:"NatIp" xml:"NatIp"`
	NatIpPort               string `json:"NatIpPort" xml:"NatIpPort"`
	AccessIp                string `json:"AccessIp" xml:"AccessIp"`
	AccessPort              string `json:"AccessPort" xml:"AccessPort"`
	IpProtocol              string `json:"IpProtocol" xml:"IpProtocol"`
	NetworkInterfaceId      string `json:"NetworkInterfaceId" xml:"NetworkInterfaceId"`
	NetworkInterfaceType    string `json:"NetworkInterfaceType" xml:"NetworkInterfaceType"`
	FullNatEntryName        string `json:"FullNatEntryName" xml:"FullNatEntryName"`
	FullNatEntryDescription string `json:"FullNatEntryDescription" xml:"FullNatEntryDescription"`
	CreationTime            string `json:"CreationTime" xml:"CreationTime"`
	FullNatEntryId          string `json:"FullNatEntryId" xml:"FullNatEntryId"`
	FullNatEntryStatus      string `json:"FullNatEntryStatus" xml:"FullNatEntryStatus"`
}
