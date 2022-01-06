// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2022 Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumExternalWorkloadLister helps list CiliumExternalWorkloads.
// All objects returned here must be treated as read-only.
type CiliumExternalWorkloadLister interface {
	// List lists all CiliumExternalWorkloads in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.CiliumExternalWorkload, err error)
	// Get retrieves the CiliumExternalWorkload from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2.CiliumExternalWorkload, error)
	CiliumExternalWorkloadListerExpansion
}

// ciliumExternalWorkloadLister implements the CiliumExternalWorkloadLister interface.
type ciliumExternalWorkloadLister struct {
	indexer cache.Indexer
}

// NewCiliumExternalWorkloadLister returns a new CiliumExternalWorkloadLister.
func NewCiliumExternalWorkloadLister(indexer cache.Indexer) CiliumExternalWorkloadLister {
	return &ciliumExternalWorkloadLister{indexer: indexer}
}

// List lists all CiliumExternalWorkloads in the indexer.
func (s *ciliumExternalWorkloadLister) List(selector labels.Selector) (ret []*v2.CiliumExternalWorkload, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumExternalWorkload))
	})
	return ret, err
}

// Get retrieves the CiliumExternalWorkload from the index for a given name.
func (s *ciliumExternalWorkloadLister) Get(name string) (*v2.CiliumExternalWorkload, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("ciliumexternalworkload"), name)
	}
	return obj.(*v2.CiliumExternalWorkload), nil
}
