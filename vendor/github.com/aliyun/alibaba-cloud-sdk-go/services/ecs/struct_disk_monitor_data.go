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

// DiskMonitorData is a nested struct in ecs response
type DiskMonitorData struct {
	BPSRead      int    `json:"BPSRead" xml:"BPSRead"`
	IOPSRead     int    `json:"IOPSRead" xml:"IOPSRead"`
	LatencyRead  int    `json:"LatencyRead" xml:"LatencyRead"`
	BPSTotal     int    `json:"BPSTotal" xml:"BPSTotal"`
	IOPSTotal    int    `json:"IOPSTotal" xml:"IOPSTotal"`
	TimeStamp    string `json:"TimeStamp" xml:"TimeStamp"`
	LatencyWrite int    `json:"LatencyWrite" xml:"LatencyWrite"`
	IOPSWrite    int    `json:"IOPSWrite" xml:"IOPSWrite"`
	DiskId       string `json:"DiskId" xml:"DiskId"`
	BPSWrite     int    `json:"BPSWrite" xml:"BPSWrite"`
}
