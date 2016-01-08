/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package io

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/noironetworks/cilium-net/docker-plugin/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api"
	"github.com/noironetworks/cilium-net/docker-plugin/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api/latest"
)

// LoadPodFromFile will read, decode, and return a Pod from a file.
func LoadPodFromFile(filePath string) (*api.Pod, error) {
	if filePath == "" {
		return nil, fmt.Errorf("file path not specified")
	}
	podDef, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file path %s: %+v", filePath, err)
	}
	if len(podDef) == 0 {
		return nil, fmt.Errorf("file was empty: %s", filePath)
	}
	pod := &api.Pod{}

	if err := latest.GroupOrDie(api.GroupName).Codec.DecodeInto(podDef, pod); err != nil {
		return nil, fmt.Errorf("failed decoding file: %v", err)
	}
	return pod, nil
}

// SavePodToFile will encode and save a pod to a given path & permissions
func SavePodToFile(pod *api.Pod, filePath string, perm os.FileMode) error {
	if filePath == "" {
		return fmt.Errorf("file path not specified")
	}
	data, err := latest.GroupOrDie(api.GroupName).Codec.Encode(pod)
	if err != nil {
		return fmt.Errorf("failed encoding pod: %v", err)
	}
	return ioutil.WriteFile(filePath, data, perm)
}
