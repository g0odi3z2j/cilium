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

// Command is a nested struct in ecs response
type Command struct {
	CreationTime         string                 `json:"CreationTime" xml:"CreationTime"`
	Type                 string                 `json:"Type" xml:"Type"`
	Timeout              int64                  `json:"Timeout" xml:"Timeout"`
	InvokeTimes          int                    `json:"InvokeTimes" xml:"InvokeTimes"`
	CommandId            string                 `json:"CommandId" xml:"CommandId"`
	WorkingDir           string                 `json:"WorkingDir" xml:"WorkingDir"`
	Description          string                 `json:"Description" xml:"Description"`
	Version              int                    `json:"Version" xml:"Version"`
	Provider             string                 `json:"Provider" xml:"Provider"`
	CommandContent       string                 `json:"CommandContent" xml:"CommandContent"`
	Category             string                 `json:"Category" xml:"Category"`
	Latest               bool                   `json:"Latest" xml:"Latest"`
	Name                 string                 `json:"Name" xml:"Name"`
	EnableParameter      bool                   `json:"EnableParameter" xml:"EnableParameter"`
	ParameterNames       ParameterNames         `json:"ParameterNames" xml:"ParameterNames"`
	ParameterDefinitions ParameterDefinitions   `json:"ParameterDefinitions" xml:"ParameterDefinitions"`
	Tags                 TagsInDescribeCommands `json:"Tags" xml:"Tags"`
}
