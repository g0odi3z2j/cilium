// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package runtime

import (
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func DeepCopy_runtime_RawExtension(in RawExtension, out *RawExtension, c *conversion.Cloner) error {
	if in.Raw != nil {
		in, out := in.Raw, &out.Raw
		*out = make([]byte, len(in))
		copy(*out, in)
	} else {
		out.Raw = nil
	}
	if in.Object == nil {
		out.Object = nil
	} else if newVal, err := c.DeepCopy(in.Object); err != nil {
		return err
	} else {
		out.Object = newVal.(Object)
	}
	return nil
}

func DeepCopy_runtime_Scheme(in Scheme, out *Scheme, c *conversion.Cloner) error {
	if in.gvkToType != nil {
		in, out := in.gvkToType, &out.gvkToType
		*out = make(map[unversioned.GroupVersionKind]reflect.Type)
		for range in {
			// FIXME: Copying unassignable keys unsupported unversioned.GroupVersionKind
		}
	} else {
		out.gvkToType = nil
	}
	if in.typeToGVK != nil {
		in, out := in.typeToGVK, &out.typeToGVK
		*out = make(map[reflect.Type][]unversioned.GroupVersionKind)
		for range in {
			// FIXME: Copying unassignable keys unsupported reflect.Type
		}
	} else {
		out.typeToGVK = nil
	}
	if in.unversionedTypes != nil {
		in, out := in.unversionedTypes, &out.unversionedTypes
		*out = make(map[reflect.Type]unversioned.GroupVersionKind)
		for range in {
			// FIXME: Copying unassignable keys unsupported reflect.Type
		}
	} else {
		out.unversionedTypes = nil
	}
	if in.unversionedKinds != nil {
		in, out := in.unversionedKinds, &out.unversionedKinds
		*out = make(map[string]reflect.Type)
		for key, val := range in {
			if newVal, err := c.DeepCopy(val); err != nil {
				return err
			} else {
				(*out)[key] = newVal.(reflect.Type)
			}
		}
	} else {
		out.unversionedKinds = nil
	}
	if in.fieldLabelConversionFuncs != nil {
		in, out := in.fieldLabelConversionFuncs, &out.fieldLabelConversionFuncs
		*out = make(map[string]map[string]FieldLabelConversionFunc)
		for key, val := range in {
			if newVal, err := c.DeepCopy(val); err != nil {
				return err
			} else {
				(*out)[key] = newVal.(map[string]FieldLabelConversionFunc)
			}
		}
	} else {
		out.fieldLabelConversionFuncs = nil
	}
	if in.converter != nil {
		in, out := in.converter, &out.converter
		*out = new(conversion.Converter)
		if err := conversion.DeepCopy_conversion_Converter(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.converter = nil
	}
	if in.cloner != nil {
		in, out := in.cloner, &out.cloner
		*out = new(conversion.Cloner)
		if err := conversion.DeepCopy_conversion_Cloner(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.cloner = nil
	}
	return nil
}

func DeepCopy_runtime_SerializerInfo(in SerializerInfo, out *SerializerInfo, c *conversion.Cloner) error {
	if in.Serializer == nil {
		out.Serializer = nil
	} else if newVal, err := c.DeepCopy(in.Serializer); err != nil {
		return err
	} else {
		out.Serializer = newVal.(Serializer)
	}
	out.EncodesAsText = in.EncodesAsText
	out.MediaType = in.MediaType
	return nil
}

func DeepCopy_runtime_StreamSerializerInfo(in StreamSerializerInfo, out *StreamSerializerInfo, c *conversion.Cloner) error {
	if err := DeepCopy_runtime_SerializerInfo(in.SerializerInfo, &out.SerializerInfo, c); err != nil {
		return err
	}
	if in.Framer == nil {
		out.Framer = nil
	} else if newVal, err := c.DeepCopy(in.Framer); err != nil {
		return err
	} else {
		out.Framer = newVal.(Framer)
	}
	if err := DeepCopy_runtime_SerializerInfo(in.Embedded, &out.Embedded, c); err != nil {
		return err
	}
	return nil
}
