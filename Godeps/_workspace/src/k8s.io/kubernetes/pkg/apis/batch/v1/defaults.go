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

import (
	"github.com/noironetworks/cilium-net/Godeps/_workspace/src/k8s.io/kubernetes/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) {
	scheme.AddDefaultingFuncs(
		func(obj *Job) {
			// For a non-parallel job, you can leave both `.spec.completions` and
			// `.spec.parallelism` unset.  When both are unset, both are defaulted to 1.
			if obj.Spec.Completions == nil && obj.Spec.Parallelism == nil {
				obj.Spec.Completions = new(int32)
				*obj.Spec.Completions = 1
				obj.Spec.Parallelism = new(int32)
				*obj.Spec.Parallelism = 1
			}
			if obj.Spec.Parallelism == nil {
				obj.Spec.Parallelism = new(int32)
				*obj.Spec.Parallelism = 1
			}
		},
	)
}
