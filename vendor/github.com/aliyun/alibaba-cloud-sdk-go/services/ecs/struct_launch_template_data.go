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

// LaunchTemplateData is a nested struct in ecs response
type LaunchTemplateData struct {
	ImageId                      string                                            `json:"ImageId" xml:"ImageId"`
	ImageOwnerAlias              string                                            `json:"ImageOwnerAlias" xml:"ImageOwnerAlias"`
	PasswordInherit              bool                                              `json:"PasswordInherit" xml:"PasswordInherit"`
	InstanceType                 string                                            `json:"InstanceType" xml:"InstanceType"`
	SecurityGroupId              string                                            `json:"SecurityGroupId" xml:"SecurityGroupId"`
	VpcId                        string                                            `json:"VpcId" xml:"VpcId"`
	VSwitchId                    string                                            `json:"VSwitchId" xml:"VSwitchId"`
	InstanceName                 string                                            `json:"InstanceName" xml:"InstanceName"`
	Description                  string                                            `json:"Description" xml:"Description"`
	InternetMaxBandwidthIn       int                                               `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut      int                                               `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	HostName                     string                                            `json:"HostName" xml:"HostName"`
	ZoneId                       string                                            `json:"ZoneId" xml:"ZoneId"`
	SystemDiskSize               int                                               `json:"SystemDisk.Size" xml:"SystemDisk.Size"`
	SystemDiskCategory           string                                            `json:"SystemDisk.Category" xml:"SystemDisk.Category"`
	SystemDiskDiskName           string                                            `json:"SystemDisk.DiskName" xml:"SystemDisk.DiskName"`
	SystemDiskDescription        string                                            `json:"SystemDisk.Description" xml:"SystemDisk.Description"`
	SystemDiskIops               int                                               `json:"SystemDisk.Iops" xml:"SystemDisk.Iops"`
	SystemDiskPerformanceLevel   string                                            `json:"SystemDisk.PerformanceLevel" xml:"SystemDisk.PerformanceLevel"`
	SystemDiskDeleteWithInstance bool                                              `json:"SystemDisk.DeleteWithInstance" xml:"SystemDisk.DeleteWithInstance"`
	IoOptimized                  string                                            `json:"IoOptimized" xml:"IoOptimized"`
	InstanceChargeType           string                                            `json:"InstanceChargeType" xml:"InstanceChargeType"`
	Period                       int                                               `json:"Period" xml:"Period"`
	InternetChargeType           string                                            `json:"InternetChargeType" xml:"InternetChargeType"`
	EnableVmOsConfig             bool                                              `json:"EnableVmOsConfig" xml:"EnableVmOsConfig"`
	NetworkType                  string                                            `json:"NetworkType" xml:"NetworkType"`
	UserData                     string                                            `json:"UserData" xml:"UserData"`
	KeyPairName                  string                                            `json:"KeyPairName" xml:"KeyPairName"`
	RamRoleName                  string                                            `json:"RamRoleName" xml:"RamRoleName"`
	AutoReleaseTime              string                                            `json:"AutoReleaseTime" xml:"AutoReleaseTime"`
	SpotStrategy                 string                                            `json:"SpotStrategy" xml:"SpotStrategy"`
	SpotPriceLimit               float64                                           `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
	SpotDuration                 int                                               `json:"SpotDuration" xml:"SpotDuration"`
	ResourceGroupId              string                                            `json:"ResourceGroupId" xml:"ResourceGroupId"`
	SecurityEnhancementStrategy  string                                            `json:"SecurityEnhancementStrategy" xml:"SecurityEnhancementStrategy"`
	PrivateIpAddress             string                                            `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
	DeploymentSetId              string                                            `json:"DeploymentSetId" xml:"DeploymentSetId"`
	SecurityGroupIds             SecurityGroupIdsInDescribeLaunchTemplateVersions  `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	DataDisks                    DataDisks                                         `json:"DataDisks" xml:"DataDisks"`
	NetworkInterfaces            NetworkInterfacesInDescribeLaunchTemplateVersions `json:"NetworkInterfaces" xml:"NetworkInterfaces"`
	Tags                         TagsInDescribeLaunchTemplateVersions              `json:"Tags" xml:"Tags"`
}
