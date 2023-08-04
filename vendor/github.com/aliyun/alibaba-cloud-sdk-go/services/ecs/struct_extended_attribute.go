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

// ExtendedAttribute is a nested struct in ecs response
type ExtendedAttribute struct {
	HostId             string                                       `json:"HostId" xml:"HostId"`
	Device             string                                       `json:"Device" xml:"Device"`
	Code               string                                       `json:"Code" xml:"Code"`
	OnlineRepairPolicy string                                       `json:"OnlineRepairPolicy" xml:"OnlineRepairPolicy"`
	DiskId             string                                       `json:"DiskId" xml:"DiskId"`
	HostType           string                                       `json:"HostType" xml:"HostType"`
	PunishType         string                                       `json:"PunishType" xml:"PunishType"`
	PunishUrl          string                                       `json:"PunishUrl" xml:"PunishUrl"`
	PunishDomain       string                                       `json:"PunishDomain" xml:"PunishDomain"`
	Rack               string                                       `json:"Rack" xml:"Rack"`
	CanAccept          string                                       `json:"CanAccept" xml:"CanAccept"`
	ResponseResult     string                                       `json:"ResponseResult" xml:"ResponseResult"`
	MigrationOptions   MigrationOptions                             `json:"MigrationOptions" xml:"MigrationOptions"`
	InactiveDisks      InactiveDisksInDescribeInstanceHistoryEvents `json:"InactiveDisks" xml:"InactiveDisks"`
}
