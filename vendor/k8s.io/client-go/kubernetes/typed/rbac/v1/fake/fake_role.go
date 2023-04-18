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

	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rbacv1 "k8s.io/client-go/applyconfigurations/rbac/v1"
	testing "k8s.io/client-go/testing"
)

// FakeRoles implements RoleInterface
type FakeRoles struct {
	Fake *FakeRbacV1
	ns   string
}

var rolesResource = v1.SchemeGroupVersion.WithResource("roles")

var rolesKind = v1.SchemeGroupVersion.WithKind("Role")

// Get takes name of the role, and returns the corresponding role object, and an error if there is any.
func (c *FakeRoles) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rolesResource, c.ns, name), &v1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Role), err
}

// List takes label and field selectors, and returns the list of Roles that match those selectors.
func (c *FakeRoles) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rolesResource, rolesKind, c.ns, opts), &v1.RoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RoleList{ListMeta: obj.(*v1.RoleList).ListMeta}
	for _, item := range obj.(*v1.RoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roles.
func (c *FakeRoles) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rolesResource, c.ns, opts))

}

// Create takes the representation of a role and creates it.  Returns the server's representation of the role, and an error, if there is any.
func (c *FakeRoles) Create(ctx context.Context, role *v1.Role, opts metav1.CreateOptions) (result *v1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rolesResource, c.ns, role), &v1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Role), err
}

// Update takes the representation of a role and updates it. Returns the server's representation of the role, and an error, if there is any.
func (c *FakeRoles) Update(ctx context.Context, role *v1.Role, opts metav1.UpdateOptions) (result *v1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rolesResource, c.ns, role), &v1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Role), err
}

// Delete takes name of the role and deletes it. Returns an error if one occurs.
func (c *FakeRoles) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rolesResource, c.ns, name, opts), &v1.Role{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoles) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rolesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RoleList{})
	return err
}

// Patch applies the patch and returns the patched role.
func (c *FakeRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rolesResource, c.ns, name, pt, data, subresources...), &v1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Role), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied role.
func (c *FakeRoles) Apply(ctx context.Context, role *rbacv1.RoleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Role, err error) {
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
		Invokes(testing.NewPatchSubresourceAction(rolesResource, c.ns, *name, types.ApplyPatchType, data), &v1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Role), err
}
