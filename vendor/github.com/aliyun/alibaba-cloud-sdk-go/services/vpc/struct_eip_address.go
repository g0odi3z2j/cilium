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

// EipAddress is a nested struct in vpc response
type EipAddress struct {
	ReservationActiveTime         string                                        `json:"ReservationActiveTime" xml:"ReservationActiveTime"`
	Status                        string                                        `json:"Status" xml:"Status"`
	ReservationOrderType          string                                        `json:"ReservationOrderType" xml:"ReservationOrderType"`
	AllocationTime                string                                        `json:"AllocationTime" xml:"AllocationTime"`
	Netmode                       string                                        `json:"Netmode" xml:"Netmode"`
	ChargeType                    string                                        `json:"ChargeType" xml:"ChargeType"`
	Descritpion                   string                                        `json:"Descritpion" xml:"Descritpion"`
	Mode                          string                                        `json:"Mode" xml:"Mode"`
	SegmentInstanceId             string                                        `json:"SegmentInstanceId" xml:"SegmentInstanceId"`
	ReservationInternetChargeType string                                        `json:"ReservationInternetChargeType" xml:"ReservationInternetChargeType"`
	BandwidthPackageId            string                                        `json:"BandwidthPackageId" xml:"BandwidthPackageId"`
	IpAddress                     string                                        `json:"IpAddress" xml:"IpAddress"`
	Bandwidth                     string                                        `json:"Bandwidth" xml:"Bandwidth"`
	ReservationBandwidth          string                                        `json:"ReservationBandwidth" xml:"ReservationBandwidth"`
	EipBandwidth                  string                                        `json:"EipBandwidth" xml:"EipBandwidth"`
	Name                          string                                        `json:"Name" xml:"Name"`
	PrivateIpAddress              string                                        `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
	InstanceRegionId              string                                        `json:"InstanceRegionId" xml:"InstanceRegionId"`
	DeletionProtection            bool                                          `json:"DeletionProtection" xml:"DeletionProtection"`
	InstanceId                    string                                        `json:"InstanceId" xml:"InstanceId"`
	SecondLimited                 bool                                          `json:"SecondLimited" xml:"SecondLimited"`
	InstanceType                  string                                        `json:"InstanceType" xml:"InstanceType"`
	HDMonitorStatus               string                                        `json:"HDMonitorStatus" xml:"HDMonitorStatus"`
	RegionId                      string                                        `json:"RegionId" xml:"RegionId"`
	BandwidthPackageBandwidth     string                                        `json:"BandwidthPackageBandwidth" xml:"BandwidthPackageBandwidth"`
	ServiceManaged                int                                           `json:"ServiceManaged" xml:"ServiceManaged"`
	ExpiredTime                   string                                        `json:"ExpiredTime" xml:"ExpiredTime"`
	ResourceGroupId               string                                        `json:"ResourceGroupId" xml:"ResourceGroupId"`
	AllocationId                  string                                        `json:"AllocationId" xml:"AllocationId"`
	InternetChargeType            string                                        `json:"InternetChargeType" xml:"InternetChargeType"`
	BusinessStatus                string                                        `json:"BusinessStatus" xml:"BusinessStatus"`
	BandwidthPackageType          string                                        `json:"BandwidthPackageType" xml:"BandwidthPackageType"`
	HasReservationData            string                                        `json:"HasReservationData" xml:"HasReservationData"`
	ISP                           string                                        `json:"ISP" xml:"ISP"`
	AvailableRegions              AvailableRegions                              `json:"AvailableRegions" xml:"AvailableRegions"`
	SecurityProtectionTypes       SecurityProtectionTypesInDescribeEipAddresses `json:"SecurityProtectionTypes" xml:"SecurityProtectionTypes"`
	OperationLocks                OperationLocksInDescribeEipAddresses          `json:"OperationLocks" xml:"OperationLocks"`
	Tags                          TagsInDescribeEipAddresses                    `json:"Tags" xml:"Tags"`
}
