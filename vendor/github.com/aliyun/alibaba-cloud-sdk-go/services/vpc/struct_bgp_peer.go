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

// BgpPeer is a nested struct in vpc response
type BgpPeer struct {
	Status        string `json:"Status" xml:"Status"`
	PeerIpAddress string `json:"PeerIpAddress" xml:"PeerIpAddress"`
	RouterId      string `json:"RouterId" xml:"RouterId"`
	BgpGroupId    string `json:"BgpGroupId" xml:"BgpGroupId"`
	BgpStatus     string `json:"BgpStatus" xml:"BgpStatus"`
	BfdMultiHop   int    `json:"BfdMultiHop" xml:"BfdMultiHop"`
	PeerAsn       string `json:"PeerAsn" xml:"PeerAsn"`
	LocalAsn      string `json:"LocalAsn" xml:"LocalAsn"`
	RegionId      string `json:"RegionId" xml:"RegionId"`
	BgpPeerId     string `json:"BgpPeerId" xml:"BgpPeerId"`
	EnableBfd     bool   `json:"EnableBfd" xml:"EnableBfd"`
	Hold          string `json:"Hold" xml:"Hold"`
	IpVersion     string `json:"IpVersion" xml:"IpVersion"`
	Keepalive     string `json:"Keepalive" xml:"Keepalive"`
	Description   string `json:"Description" xml:"Description"`
	RouteLimit    string `json:"RouteLimit" xml:"RouteLimit"`
	IsFake        string `json:"IsFake" xml:"IsFake"`
	AuthKey       string `json:"AuthKey" xml:"AuthKey"`
	Name          string `json:"Name" xml:"Name"`
}
