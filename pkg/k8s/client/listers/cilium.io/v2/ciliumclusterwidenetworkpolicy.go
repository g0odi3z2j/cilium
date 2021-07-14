// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumClusterwideNetworkPolicyLister helps list CiliumClusterwideNetworkPolicies.
// All objects returned here must be treated as read-only.
type CiliumClusterwideNetworkPolicyLister interface {
	// List lists all CiliumClusterwideNetworkPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.CiliumClusterwideNetworkPolicy, err error)
	// Get retrieves the CiliumClusterwideNetworkPolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2.CiliumClusterwideNetworkPolicy, error)
	CiliumClusterwideNetworkPolicyListerExpansion
}

// ciliumClusterwideNetworkPolicyLister implements the CiliumClusterwideNetworkPolicyLister interface.
type ciliumClusterwideNetworkPolicyLister struct {
	indexer cache.Indexer
}

// NewCiliumClusterwideNetworkPolicyLister returns a new CiliumClusterwideNetworkPolicyLister.
func NewCiliumClusterwideNetworkPolicyLister(indexer cache.Indexer) CiliumClusterwideNetworkPolicyLister {
	return &ciliumClusterwideNetworkPolicyLister{indexer: indexer}
}

// List lists all CiliumClusterwideNetworkPolicies in the indexer.
func (s *ciliumClusterwideNetworkPolicyLister) List(selector labels.Selector) (ret []*v2.CiliumClusterwideNetworkPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumClusterwideNetworkPolicy))
	})
	return ret, err
}

// Get retrieves the CiliumClusterwideNetworkPolicy from the index for a given name.
func (s *ciliumClusterwideNetworkPolicyLister) Get(name string) (*v2.CiliumClusterwideNetworkPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("ciliumclusterwidenetworkpolicy"), name)
	}
	return obj.(*v2.CiliumClusterwideNetworkPolicy), nil
}
