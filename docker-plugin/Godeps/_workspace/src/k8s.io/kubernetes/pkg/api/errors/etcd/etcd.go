/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package etcd

import (
	"github.com/noironetworks/cilium-net/docker-plugin/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api/errors"
	"github.com/noironetworks/cilium-net/docker-plugin/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/storage"
)

// InterpretListError converts a generic error on a retrieval
// operation into the appropriate API error.
func InterpretListError(err error, qualifiedResource unversioned.GroupResource) error {
	switch {
	case storage.IsNotFound(err):
		return errors.NewNotFound(qualifiedResource, "")
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(qualifiedResource, "list", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}

// InterpretGetError converts a generic error on a retrieval
// operation into the appropriate API error.
func InterpretGetError(err error, qualifiedResource unversioned.GroupResource, name string) error {
	switch {
	case storage.IsNotFound(err):
		return errors.NewNotFound(qualifiedResource, name)
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(qualifiedResource, "get", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}

// InterpretCreateError converts a generic error on a create
// operation into the appropriate API error.
func InterpretCreateError(err error, qualifiedResource unversioned.GroupResource, name string) error {
	switch {
	case storage.IsNodeExist(err):
		return errors.NewAlreadyExists(qualifiedResource, name)
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(qualifiedResource, "create", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}

// InterpretUpdateError converts a generic error on a update
// operation into the appropriate API error.
func InterpretUpdateError(err error, qualifiedResource unversioned.GroupResource, name string) error {
	switch {
	case storage.IsTestFailed(err), storage.IsNodeExist(err):
		return errors.NewConflict(qualifiedResource, name, err)
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(qualifiedResource, "update", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}

// InterpretDeleteError converts a generic error on a delete
// operation into the appropriate API error.
func InterpretDeleteError(err error, qualifiedResource unversioned.GroupResource, name string) error {
	switch {
	case storage.IsNotFound(err):
		return errors.NewNotFound(qualifiedResource, name)
	case storage.IsUnreachable(err):
		return errors.NewServerTimeout(qualifiedResource, "delete", 2) // TODO: make configurable or handled at a higher level
	default:
		return err
	}
}
