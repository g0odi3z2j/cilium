/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1beta1 "k8s.io/api/rbac/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rbacv1beta1 "k8s.io/client-go/applyconfigurations/rbac/v1beta1"
	testing "k8s.io/client-go/testing"
)

// FakeRoles implements RoleInterface
type FakeRoles struct {
	Fake *FakeRbacV1beta1
	ns   string
}

var rolesResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1beta1", Resource: "roles"}

var rolesKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1beta1", Kind: "Role"}

// Get takes name of the role, and returns the corresponding role object, and an error if there is any.
func (c *FakeRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rolesResource, c.ns, name), &v1beta1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Role), err
}

// List takes label and field selectors, and returns the list of Roles that match those selectors.
func (c *FakeRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.RoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rolesResource, rolesKind, c.ns, opts), &v1beta1.RoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.RoleList{ListMeta: obj.(*v1beta1.RoleList).ListMeta}
	for _, item := range obj.(*v1beta1.RoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roles.
func (c *FakeRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rolesResource, c.ns, opts))

}

// Create takes the representation of a role and creates it.  Returns the server's representation of the role, and an error, if there is any.
func (c *FakeRoles) Create(ctx context.Context, role *v1beta1.Role, opts v1.CreateOptions) (result *v1beta1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rolesResource, c.ns, role), &v1beta1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Role), err
}

// Update takes the representation of a role and updates it. Returns the server's representation of the role, and an error, if there is any.
func (c *FakeRoles) Update(ctx context.Context, role *v1beta1.Role, opts v1.UpdateOptions) (result *v1beta1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rolesResource, c.ns, role), &v1beta1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Role), err
}

// Delete takes name of the role and deletes it. Returns an error if one occurs.
func (c *FakeRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rolesResource, c.ns, name), &v1beta1.Role{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rolesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.RoleList{})
	return err
}

// Patch applies the patch and returns the patched role.
func (c *FakeRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rolesResource, c.ns, name, pt, data, subresources...), &v1beta1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Role), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied role.
func (c *FakeRoles) Apply(ctx context.Context, role *rbacv1beta1.RoleApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Role, err error) {
	if role == nil {
		return nil, fmt.Errorf("role provided to Apply must not be nil")
	}
	data, err := json.Marshal(role)
	if err != nil {
		return nil, err
	}
	name := role.Name
	if name == nil {
		return nil, fmt.Errorf("role.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rolesResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Role), err
}
