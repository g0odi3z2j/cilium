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

// InvocationResult is a nested struct in ecs response
type InvocationResult struct {
	InvocationStatus   string `json:"InvocationStatus" xml:"InvocationStatus"`
	Repeats            int    `json:"Repeats" xml:"Repeats"`
	CommandId          string `json:"CommandId" xml:"CommandId"`
	InstanceId         string `json:"InstanceId" xml:"InstanceId"`
	Output             string `json:"Output" xml:"Output"`
	Dropped            int    `json:"Dropped" xml:"Dropped"`
	StopTime           string `json:"StopTime" xml:"StopTime"`
	ExitCode           int64  `json:"ExitCode" xml:"ExitCode"`
	StartTime          string `json:"StartTime" xml:"StartTime"`
	ErrorInfo          string `json:"ErrorInfo" xml:"ErrorInfo"`
	ErrorCode          string `json:"ErrorCode" xml:"ErrorCode"`
	FinishedTime       string `json:"FinishedTime" xml:"FinishedTime"`
	InvokeId           string `json:"InvokeId" xml:"InvokeId"`
	InvokeRecordStatus string `json:"InvokeRecordStatus" xml:"InvokeRecordStatus"`
	Username           string `json:"Username" xml:"Username"`
}
