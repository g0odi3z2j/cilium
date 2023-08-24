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

// CiliumBGPClusterConfigLister helps list CiliumBGPClusterConfigs.
// All objects returned here must be treated as read-only.
type CiliumBGPClusterConfigLister interface {
	// List lists all CiliumBGPClusterConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.CiliumBGPClusterConfig, err error)
	// Get retrieves the CiliumBGPClusterConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2alpha1.CiliumBGPClusterConfig, error)
	CiliumBGPClusterConfigListerExpansion
}

// ciliumBGPClusterConfigLister implements the CiliumBGPClusterConfigLister interface.
type ciliumBGPClusterConfigLister struct {
	indexer cache.Indexer
}

// NewCiliumBGPClusterConfigLister returns a new CiliumBGPClusterConfigLister.
func NewCiliumBGPClusterConfigLister(indexer cache.Indexer) CiliumBGPClusterConfigLister {
	return &ciliumBGPClusterConfigLister{indexer: indexer}
}

// List lists all CiliumBGPClusterConfigs in the indexer.
func (s *ciliumBGPClusterConfigLister) List(selector labels.Selector) (ret []*v2alpha1.CiliumBGPClusterConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.CiliumBGPClusterConfig))
	})
	return ret, err
}

// Get retrieves the CiliumBGPClusterConfig from the index for a given name.
func (s *ciliumBGPClusterConfigLister) Get(name string) (*v2alpha1.CiliumBGPClusterConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2alpha1.Resource("ciliumbgpclusterconfig"), name)
	}
	return obj.(*v2alpha1.CiliumBGPClusterConfig), nil
}
