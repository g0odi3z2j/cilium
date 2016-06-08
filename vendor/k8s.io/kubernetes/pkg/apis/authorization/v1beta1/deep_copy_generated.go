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

package v1beta1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_v1beta1_LocalSubjectAccessReview,
		DeepCopy_v1beta1_NonResourceAttributes,
		DeepCopy_v1beta1_ResourceAttributes,
		DeepCopy_v1beta1_SelfSubjectAccessReview,
		DeepCopy_v1beta1_SelfSubjectAccessReviewSpec,
		DeepCopy_v1beta1_SubjectAccessReview,
		DeepCopy_v1beta1_SubjectAccessReviewSpec,
		DeepCopy_v1beta1_SubjectAccessReviewStatus,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1beta1_LocalSubjectAccessReview(in LocalSubjectAccessReview, out *LocalSubjectAccessReview, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_SubjectAccessReviewSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_SubjectAccessReviewStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_NonResourceAttributes(in NonResourceAttributes, out *NonResourceAttributes, c *conversion.Cloner) error {
	out.Path = in.Path
	out.Verb = in.Verb
	return nil
}

func DeepCopy_v1beta1_ResourceAttributes(in ResourceAttributes, out *ResourceAttributes, c *conversion.Cloner) error {
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.Subresource = in.Subresource
	out.Name = in.Name
	return nil
}

func DeepCopy_v1beta1_SelfSubjectAccessReview(in SelfSubjectAccessReview, out *SelfSubjectAccessReview, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_SelfSubjectAccessReviewSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_SubjectAccessReviewStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_SelfSubjectAccessReviewSpec(in SelfSubjectAccessReviewSpec, out *SelfSubjectAccessReviewSpec, c *conversion.Cloner) error {
	if in.ResourceAttributes != nil {
		in, out := in.ResourceAttributes, &out.ResourceAttributes
		*out = new(ResourceAttributes)
		if err := DeepCopy_v1beta1_ResourceAttributes(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.ResourceAttributes = nil
	}
	if in.NonResourceAttributes != nil {
		in, out := in.NonResourceAttributes, &out.NonResourceAttributes
		*out = new(NonResourceAttributes)
		if err := DeepCopy_v1beta1_NonResourceAttributes(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.NonResourceAttributes = nil
	}
	return nil
}

func DeepCopy_v1beta1_SubjectAccessReview(in SubjectAccessReview, out *SubjectAccessReview, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_SubjectAccessReviewSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_SubjectAccessReviewStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_SubjectAccessReviewSpec(in SubjectAccessReviewSpec, out *SubjectAccessReviewSpec, c *conversion.Cloner) error {
	if in.ResourceAttributes != nil {
		in, out := in.ResourceAttributes, &out.ResourceAttributes
		*out = new(ResourceAttributes)
		if err := DeepCopy_v1beta1_ResourceAttributes(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.ResourceAttributes = nil
	}
	if in.NonResourceAttributes != nil {
		in, out := in.NonResourceAttributes, &out.NonResourceAttributes
		*out = new(NonResourceAttributes)
		if err := DeepCopy_v1beta1_NonResourceAttributes(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.NonResourceAttributes = nil
	}
	out.User = in.User
	if in.Groups != nil {
		in, out := in.Groups, &out.Groups
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Groups = nil
	}
	if in.Extra != nil {
		in, out := in.Extra, &out.Extra
		*out = make(map[string][]string)
		for key, val := range in {
			if newVal, err := c.DeepCopy(val); err != nil {
				return err
			} else {
				(*out)[key] = newVal.([]string)
			}
		}
	} else {
		out.Extra = nil
	}
	return nil
}

func DeepCopy_v1beta1_SubjectAccessReviewStatus(in SubjectAccessReviewStatus, out *SubjectAccessReviewStatus, c *conversion.Cloner) error {
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	return nil
}
