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

// FlowLog is a nested struct in vpc response
type FlowLog struct {
	Status                     string                 `json:"Status" xml:"Status"`
	CreationTime               string                 `json:"CreationTime" xml:"CreationTime"`
	FlowLogName                string                 `json:"FlowLogName" xml:"FlowLogName"`
	TrafficType                string                 `json:"TrafficType" xml:"TrafficType"`
	ResourceType               string                 `json:"ResourceType" xml:"ResourceType"`
	Description                string                 `json:"Description" xml:"Description"`
	ProjectName                string                 `json:"ProjectName" xml:"ProjectName"`
	LogStoreName               string                 `json:"LogStoreName" xml:"LogStoreName"`
	ResourceId                 string                 `json:"ResourceId" xml:"ResourceId"`
	RegionId                   string                 `json:"RegionId" xml:"RegionId"`
	FlowLogId                  string                 `json:"FlowLogId" xml:"FlowLogId"`
	BusinessStatus             string                 `json:"BusinessStatus" xml:"BusinessStatus"`
	AggregationInterval        int                    `json:"AggregationInterval" xml:"AggregationInterval"`
	ServiceType                string                 `json:"ServiceType" xml:"ServiceType"`
	ResourceGroupId            string                 `json:"ResourceGroupId" xml:"ResourceGroupId"`
	FlowLogDeliverStatus       string                 `json:"FlowLogDeliverStatus" xml:"FlowLogDeliverStatus"`
	FlowLogDeliverErrorMessage string                 `json:"FlowLogDeliverErrorMessage" xml:"FlowLogDeliverErrorMessage"`
	TrafficPath                TrafficPath            `json:"TrafficPath" xml:"TrafficPath"`
	Tags                       TagsInDescribeFlowLogs `json:"Tags" xml:"Tags"`
}
