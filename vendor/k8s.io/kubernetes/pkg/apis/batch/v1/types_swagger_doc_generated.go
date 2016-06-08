/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Job = map[string]string{
	"":         "Job represents the configuration of a single job.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Spec is a structure defining the expected behavior of a job. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status":   "Status is a structure describing current status of a job. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (Job) SwaggerDoc() map[string]string {
	return map_Job
}

var map_JobCondition = map[string]string{
	"":                   "JobCondition describes current state of a job.",
	"type":               "Type of job condition, Complete or Failed.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastProbeTime":      "Last time the condition was checked.",
	"lastTransitionTime": "Last time the condition transit from one status to another.",
	"reason":             "(brief) reason for the condition's last transition.",
	"message":            "Human readable message indicating details about last transition.",
}

func (JobCondition) SwaggerDoc() map[string]string {
	return map_JobCondition
}

var map_JobList = map[string]string{
	"":         "JobList is a collection of jobs.",
	"metadata": "Standard list metadata More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of Job.",
}

func (JobList) SwaggerDoc() map[string]string {
	return map_JobList
}

var map_JobSpec = map[string]string{
	"":                      "JobSpec describes how the job execution will look like.",
	"parallelism":           "Parallelism specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: http://releases.k8s.io/HEAD/docs/user-guide/jobs.md",
	"completions":           "Completions specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: http://releases.k8s.io/HEAD/docs/user-guide/jobs.md",
	"activeDeadlineSeconds": "Optional duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer",
	"selector":              "Selector is a label query over pods that should match the pod count. Normally, the system sets this field for you. More info: http://releases.k8s.io/HEAD/docs/user-guide/labels.md#label-selectors",
	"manualSelector":        "ManualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: http://releases.k8s.io/HEAD/docs/design/selector-generation.md",
	"template":              "Template is the object that describes the pod that will be created when executing a job. More info: http://releases.k8s.io/HEAD/docs/user-guide/jobs.md",
}

func (JobSpec) SwaggerDoc() map[string]string {
	return map_JobSpec
}

var map_JobStatus = map[string]string{
	"":               "JobStatus represents the current state of a Job.",
	"conditions":     "Conditions represent the latest available observations of an object's current state. More info: http://releases.k8s.io/HEAD/docs/user-guide/jobs.md",
	"startTime":      "StartTime represents time when the job was acknowledged by the Job Manager. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
	"completionTime": "CompletionTime represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
	"active":         "Active is the number of actively running pods.",
	"succeeded":      "Succeeded is the number of pods which reached Phase Succeeded.",
	"failed":         "Failed is the number of pods which reached Phase Failed.",
}

func (JobStatus) SwaggerDoc() map[string]string {
	return map_JobStatus
}

var map_LabelSelector = map[string]string{
	"":                 "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
	"matchLabels":      "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
	"matchExpressions": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
}

func (LabelSelector) SwaggerDoc() map[string]string {
	return map_LabelSelector
}

var map_LabelSelectorRequirement = map[string]string{
	"":         "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
	"key":      "key is the label key that the selector applies to.",
	"operator": "operator represents a key's relationship to a set of values. Valid operators ard In, NotIn, Exists and DoesNotExist.",
	"values":   "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
}

func (LabelSelectorRequirement) SwaggerDoc() map[string]string {
	return map_LabelSelectorRequirement
}

// AUTO-GENERATED FUNCTIONS END HERE
