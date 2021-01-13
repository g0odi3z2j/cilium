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

// Ipv6Translator is a nested struct in vpc response
type Ipv6Translator struct {
	Ipv6TranslatorId       string                 `json:"Ipv6TranslatorId" xml:"Ipv6TranslatorId"`
	CreateTime             int64                  `json:"CreateTime" xml:"CreateTime"`
	EndTime                int64                  `json:"EndTime" xml:"EndTime"`
	Spec                   string                 `json:"Spec" xml:"Spec"`
	Name                   string                 `json:"Name" xml:"Name"`
	Description            string                 `json:"Description" xml:"Description"`
	Status                 string                 `json:"Status" xml:"Status"`
	BusinessStatus         string                 `json:"BusinessStatus" xml:"BusinessStatus"`
	PayType                string                 `json:"PayType" xml:"PayType"`
	Bandwidth              int                    `json:"Bandwidth" xml:"Bandwidth"`
	AllocateIpv6Addr       string                 `json:"AllocateIpv6Addr" xml:"AllocateIpv6Addr"`
	AllocateIpv4Addr       string                 `json:"AllocateIpv4Addr" xml:"AllocateIpv4Addr"`
	AvailableBandwidth     string                 `json:"AvailableBandwidth" xml:"AvailableBandwidth"`
	RegionId               string                 `json:"RegionId" xml:"RegionId"`
	Ipv6TranslatorEntryIds Ipv6TranslatorEntryIds `json:"Ipv6TranslatorEntryIds" xml:"Ipv6TranslatorEntryIds"`
}
