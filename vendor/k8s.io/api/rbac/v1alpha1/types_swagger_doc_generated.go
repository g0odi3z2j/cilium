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

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_AggregationRule = map[string]string{
	"":                     "AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole",
	"clusterRoleSelectors": "ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added",
}

func (AggregationRule) SwaggerDoc() map[string]string {
	return map_AggregationRule
}

var map_ClusterRole = map[string]string{
	"":                "ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRole, and will no longer be served in v1.22.",
	"metadata":        "Standard object's metadata.",
	"rules":           "Rules holds all the PolicyRules for this ClusterRole",
	"aggregationRule": "AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.",
}

func (ClusterRole) SwaggerDoc() map[string]string {
	return map_ClusterRole
}

var map_ClusterRoleBinding = map[string]string{
	"":         "ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoleBinding, and will no longer be served in v1.22.",
	"metadata": "Standard object's metadata.",
	"subjects": "Subjects holds references to the objects the role applies to.",
	"roleRef":  "RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.",
}

func (ClusterRoleBinding) SwaggerDoc() map[string]string {
	return map_ClusterRoleBinding
}

var map_ClusterRoleBindingList = map[string]string{
	"":         "ClusterRoleBindingList is a collection of ClusterRoleBindings. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoleBindings, and will no longer be served in v1.22.",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of ClusterRoleBindings",
}

func (ClusterRoleBindingList) SwaggerDoc() map[string]string {
	return map_ClusterRoleBindingList
}

var map_ClusterRoleList = map[string]string{
	"":         "ClusterRoleList is a collection of ClusterRoles. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoles, and will no longer be served in v1.22.",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of ClusterRoles",
}

func (ClusterRoleList) SwaggerDoc() map[string]string {
	return map_ClusterRoleList
}

var map_PolicyRule = map[string]string{
	"":                "PolicyRule holds information that describes a policy rule, but does not contain information about who the rule applies to or which namespace the rule applies to.",
	"verbs":           "Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. '*' represents all verbs.",
	"apiGroups":       "APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.",
	"resources":       "Resources is a list of resources this rule applies to. '*' represents all resources.",
	"resourceNames":   "ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.",
	"nonResourceURLs": "NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as \"pods\" or \"secrets\") or non-resource URL paths (such as \"/api\"),  but not both.",
}

func (PolicyRule) SwaggerDoc() map[string]string {
	return map_PolicyRule
}

var map_Role = map[string]string{
	"":         "Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 Role, and will no longer be served in v1.22.",
	"metadata": "Standard object's metadata.",
	"rules":    "Rules holds all the PolicyRules for this Role",
}

func (Role) SwaggerDoc() map[string]string {
	return map_Role
}

var map_RoleBinding = map[string]string{
	"":         "RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 RoleBinding, and will no longer be served in v1.22.",
	"metadata": "Standard object's metadata.",
	"subjects": "Subjects holds references to the objects the role applies to.",
	"roleRef":  "RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.",
}

func (RoleBinding) SwaggerDoc() map[string]string {
	return map_RoleBinding
}

var map_RoleBindingList = map[string]string{
	"":         "RoleBindingList is a collection of RoleBindings Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 RoleBindingList, and will no longer be served in v1.22.",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of RoleBindings",
}

func (RoleBindingList) SwaggerDoc() map[string]string {
	return map_RoleBindingList
}

var map_RoleList = map[string]string{
	"":         "RoleList is a collection of Roles. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 RoleList, and will no longer be served in v1.22.",
	"metadata": "Standard object's metadata.",
	"items":    "Items is a list of Roles",
}

func (RoleList) SwaggerDoc() map[string]string {
	return map_RoleList
}

var map_RoleRef = map[string]string{
	"":         "RoleRef contains information that points to the role being used",
	"apiGroup": "APIGroup is the group for the resource being referenced",
	"kind":     "Kind is the type of resource being referenced",
	"name":     "Name is the name of resource being referenced",
}

func (RoleRef) SwaggerDoc() map[string]string {
	return map_RoleRef
}

var map_Subject = map[string]string{
	"":           "Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.",
	"kind":       "Kind of object being referenced. Values defined by this API group are \"User\", \"Group\", and \"ServiceAccount\". If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
	"apiVersion": "APIVersion holds the API group and version of the referenced subject. Defaults to \"v1\" for ServiceAccount subjects. Defaults to \"rbac.authorization.k8s.io/v1alpha1\" for User and Group subjects.",
	"name":       "Name of the object being referenced.",
	"namespace":  "Namespace of the referenced object.  If the object kind is non-namespace, such as \"User\" or \"Group\", and this value is not empty the Authorizer should report an error.",
}

func (Subject) SwaggerDoc() map[string]string {
	return map_Subject
}

// AUTO-GENERATED FUNCTIONS END HERE
