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

// InstanceType is a nested struct in ecs response
type InstanceType struct {
	MemorySize                  float64 `json:"MemorySize" xml:"MemorySize"`
	InstancePpsRx               int64   `json:"InstancePpsRx" xml:"InstancePpsRx"`
	EriQuantity                 int     `json:"EriQuantity" xml:"EriQuantity"`
	EniPrivateIpAddressQuantity int     `json:"EniPrivateIpAddressQuantity" xml:"EniPrivateIpAddressQuantity"`
	CpuCoreCount                int     `json:"CpuCoreCount" xml:"CpuCoreCount"`
	EniTotalQuantity            int     `json:"EniTotalQuantity" xml:"EniTotalQuantity"`
	Cores                       int     `json:"Cores" xml:"Cores"`
	InstanceTypeId              string  `json:"InstanceTypeId" xml:"InstanceTypeId"`
	InstanceBandwidthRx         int     `json:"InstanceBandwidthRx" xml:"InstanceBandwidthRx"`
	InstanceType                string  `json:"InstanceType" xml:"InstanceType"`
	QueuePairNumber             int     `json:"QueuePairNumber" xml:"QueuePairNumber"`
	EniQuantity                 int     `json:"EniQuantity" xml:"EniQuantity"`
	Generation                  string  `json:"Generation" xml:"Generation"`
	SupportIoOptimized          string  `json:"SupportIoOptimized" xml:"SupportIoOptimized"`
	InstanceTypeFamily          string  `json:"InstanceTypeFamily" xml:"InstanceTypeFamily"`
	InitialCredit               int     `json:"InitialCredit" xml:"InitialCredit"`
	InstancePpsTx               int64   `json:"InstancePpsTx" xml:"InstancePpsTx"`
	InstanceFamilyLevel         string  `json:"InstanceFamilyLevel" xml:"InstanceFamilyLevel"`
	LocalStorageAmount          int     `json:"LocalStorageAmount" xml:"LocalStorageAmount"`
	TotalEniQueueQuantity       int     `json:"TotalEniQueueQuantity" xml:"TotalEniQueueQuantity"`
	GPUSpec                     string  `json:"GPUSpec" xml:"GPUSpec"`
	SecondaryEniQueueNumber     int     `json:"SecondaryEniQueueNumber" xml:"SecondaryEniQueueNumber"`
	InstanceBandwidthTx         int     `json:"InstanceBandwidthTx" xml:"InstanceBandwidthTx"`
	MaximumQueueNumberPerEni    int     `json:"MaximumQueueNumberPerEni" xml:"MaximumQueueNumberPerEni"`
	DiskQuantity                int     `json:"DiskQuantity" xml:"DiskQuantity"`
	PrimaryEniQueueNumber       int     `json:"PrimaryEniQueueNumber" xml:"PrimaryEniQueueNumber"`
	Memory                      int     `json:"Memory" xml:"Memory"`
	BaselineCredit              int     `json:"BaselineCredit" xml:"BaselineCredit"`
	EniTrunkSupported           bool    `json:"EniTrunkSupported" xml:"EniTrunkSupported"`
	GPUAmount                   int     `json:"GPUAmount" xml:"GPUAmount"`
	NvmeSupport                 string  `json:"NvmeSupport" xml:"NvmeSupport"`
	EniIpv6AddressQuantity      int     `json:"EniIpv6AddressQuantity" xml:"EniIpv6AddressQuantity"`
	LocalStorageCapacity        int64   `json:"LocalStorageCapacity" xml:"LocalStorageCapacity"`
	LocalStorageCategory        string  `json:"LocalStorageCategory" xml:"LocalStorageCategory"`
}
