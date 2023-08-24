// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2alpha1

import (
	v2alpha1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumBGPNodeConfigLister helps list CiliumBGPNodeConfigs.
// All objects returned here must be treated as read-only.
type CiliumBGPNodeConfigLister interface {
	// List lists all CiliumBGPNodeConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.CiliumBGPNodeConfig, err error)
	// Get retrieves the CiliumBGPNodeConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2alpha1.CiliumBGPNodeConfig, error)
	CiliumBGPNodeConfigListerExpansion
}

// ciliumBGPNodeConfigLister implements the CiliumBGPNodeConfigLister interface.
type ciliumBGPNodeConfigLister struct {
	indexer cache.Indexer
}

// NewCiliumBGPNodeConfigLister returns a new CiliumBGPNodeConfigLister.
func NewCiliumBGPNodeConfigLister(indexer cache.Indexer) CiliumBGPNodeConfigLister {
	return &ciliumBGPNodeConfigLister{indexer: indexer}
}

// List lists all CiliumBGPNodeConfigs in the indexer.
func (s *ciliumBGPNodeConfigLister) List(selector labels.Selector) (ret []*v2alpha1.CiliumBGPNodeConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.CiliumBGPNodeConfig))
	})
	return ret, err
}

// Get retrieves the CiliumBGPNodeConfig from the index for a given name.
func (s *ciliumBGPNodeConfigLister) Get(name string) (*v2alpha1.CiliumBGPNodeConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2alpha1.Resource("ciliumbgpnodeconfig"), name)
	}
	return obj.(*v2alpha1.CiliumBGPNodeConfig), nil
}
