// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	time "time"

	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	versioned "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/cilium/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v2 "github.com/cilium/cilium/pkg/k8s/client/listers/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CiliumClusterwideNetworkPolicyInformer provides access to a shared informer and lister for
// CiliumClusterwideNetworkPolicies.
type CiliumClusterwideNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.CiliumClusterwideNetworkPolicyLister
}

type ciliumClusterwideNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCiliumClusterwideNetworkPolicyInformer constructs a new informer for CiliumClusterwideNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCiliumClusterwideNetworkPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCiliumClusterwideNetworkPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCiliumClusterwideNetworkPolicyInformer constructs a new informer for CiliumClusterwideNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCiliumClusterwideNetworkPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumClusterwideNetworkPolicies().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumClusterwideNetworkPolicies().Watch(options)
			},
		},
		&ciliumiov2.CiliumClusterwideNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *ciliumClusterwideNetworkPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCiliumClusterwideNetworkPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ciliumClusterwideNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ciliumiov2.CiliumClusterwideNetworkPolicy{}, f.defaultInformer)
}

func (f *ciliumClusterwideNetworkPolicyInformer) Lister() v2.CiliumClusterwideNetworkPolicyLister {
	return v2.NewCiliumClusterwideNetworkPolicyLister(f.Informer().GetIndexer())
}
