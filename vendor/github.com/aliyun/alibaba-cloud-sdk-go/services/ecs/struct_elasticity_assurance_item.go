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

// ElasticityAssuranceItem is a nested struct in ecs response
type ElasticityAssuranceItem struct {
	PrivatePoolOptionsName          string                                           `json:"PrivatePoolOptionsName" xml:"PrivatePoolOptionsName"`
	PrivatePoolOptionsMatchCriteria string                                           `json:"PrivatePoolOptionsMatchCriteria" xml:"PrivatePoolOptionsMatchCriteria"`
	LatestStartTime                 string                                           `json:"LatestStartTime" xml:"LatestStartTime"`
	UsedAssuranceTimes              int                                              `json:"UsedAssuranceTimes" xml:"UsedAssuranceTimes"`
	RegionId                        string                                           `json:"RegionId" xml:"RegionId"`
	StartTime                       string                                           `json:"StartTime" xml:"StartTime"`
	EndTime                         string                                           `json:"EndTime" xml:"EndTime"`
	ResourceGroupId                 string                                           `json:"ResourceGroupId" xml:"ResourceGroupId"`
	TotalAssuranceTimes             string                                           `json:"TotalAssuranceTimes" xml:"TotalAssuranceTimes"`
	Status                          string                                           `json:"Status" xml:"Status"`
	Description                     string                                           `json:"Description" xml:"Description"`
	PrivatePoolOptionsId            string                                           `json:"PrivatePoolOptionsId" xml:"PrivatePoolOptionsId"`
	AllocatedResources              AllocatedResourcesInDescribeElasticityAssurances `json:"AllocatedResources" xml:"AllocatedResources"`
	Tags                            TagsInDescribeElasticityAssurances               `json:"Tags" xml:"Tags"`
}
