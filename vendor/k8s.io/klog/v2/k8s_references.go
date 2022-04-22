/*
Copyright 2021 The Kubernetes Authors.

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

package klog

import (
	"fmt"
	"reflect"

	"github.com/go-logr/logr"
)

// ObjectRef references a kubernetes object
type ObjectRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

func (ref ObjectRef) String() string {
	if ref.Namespace != "" {
		return fmt.Sprintf("%s/%s", ref.Namespace, ref.Name)
	}
	return ref.Name
}

// MarshalLog ensures that loggers with support for structured output will log
// as a struct by removing the String method via a custom type.
func (ref ObjectRef) MarshalLog() interface{} {
	type or ObjectRef
	return or(ref)
}

var _ logr.Marshaler = ObjectRef{}

// KMetadata is a subset of the kubernetes k8s.io/apimachinery/pkg/apis/meta/v1.Object interface
// this interface may expand in the future, but will always be a subset of the
// kubernetes k8s.io/apimachinery/pkg/apis/meta/v1.Object interface
type KMetadata interface {
	GetName() string
	GetNamespace() string
}

// KObj returns ObjectRef from ObjectMeta
func KObj(obj KMetadata) ObjectRef {
	if obj == nil {
		return ObjectRef{}
	}
	if val := reflect.ValueOf(obj); val.Kind() == reflect.Ptr && val.IsNil() {
		return ObjectRef{}
	}

	return ObjectRef{
		Name:      obj.GetName(),
		Namespace: obj.GetNamespace(),
	}
}

// KRef returns ObjectRef from name and namespace
func KRef(namespace, name string) ObjectRef {
	return ObjectRef{
		Name:      name,
		Namespace: namespace,
	}
}

// KObjs returns slice of ObjectRef from an slice of ObjectMeta
func KObjs(arg interface{}) []ObjectRef {
	s := reflect.ValueOf(arg)
	if s.Kind() != reflect.Slice {
		return nil
	}
	objectRefs := make([]ObjectRef, 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		if v, ok := s.Index(i).Interface().(KMetadata); ok {
			objectRefs = append(objectRefs, KObj(v))
		} else {
			return nil
		}
	}
	return objectRefs
}
