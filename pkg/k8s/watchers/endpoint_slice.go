// Copyright 2019-2020 Authors of Cilium
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

package watchers

import (
	"sync"

	"github.com/cilium/cilium/pkg/k8s"
	"github.com/cilium/cilium/pkg/k8s/informer"
	"github.com/cilium/cilium/pkg/lock"
	"github.com/cilium/cilium/pkg/serializer"

	v1 "k8s.io/api/core/v1"
	"k8s.io/api/discovery/v1beta1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// endpointSlicesInit returns true if the cluster contains endpoint slices.
func (k *K8sWatcher) endpointSlicesInit(k8sClient kubernetes.Interface, serEps *serializer.FunctionQueue, swgEps *lock.StoppableWaitGroup) bool {
	var (
		hasEndpointSlices = make(chan struct{})
		once              sync.Once
	)

	_, endpointController := informer.NewInformer(
		cache.NewListWatchFromClient(k8sClient.DiscoveryV1beta1().RESTClient(),
			"endpointslices", v1.NamespaceAll, fields.Everything()),
		&v1beta1.EndpointSlice{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				once.Do(func() {
					// signalize that we have received an endpoint slice
					// so it means the cluster has endpoint slices enabled.
					close(hasEndpointSlices)
				})
				var valid, equal bool
				defer func() { k.K8sEventReceived(metricEndpointSlice, metricCreate, valid, equal) }()
				if k8sEP := k8s.CopyObjToV1EndpointSlice(obj); k8sEP != nil {
					valid = true
					swgEps.Add()
					serEps.Enqueue(func() error {
						defer swgEps.Done()
						k.K8sSvcCache.UpdateEndpointSlices(k8sEP, swgEps)
						k.K8sEventProcessed(metricEndpointSlice, metricCreate, true)
						return nil
					}, serializer.NoRetry)
				}
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				var valid, equal bool
				defer func() { k.K8sEventReceived(metricEndpointSlice, metricUpdate, valid, equal) }()
				if oldk8sEP := k8s.CopyObjToV1EndpointSlice(oldObj); oldk8sEP != nil {
					valid = true
					if newk8sEP := k8s.CopyObjToV1EndpointSlice(newObj); newk8sEP != nil {
						if k8s.EqualV1EndpointSlice(oldk8sEP, newk8sEP) {
							equal = true
							return
						}

						swgEps.Add()
						serEps.Enqueue(func() error {
							defer swgEps.Done()
							k.K8sSvcCache.UpdateEndpointSlices(newk8sEP, swgEps)
							k.K8sEventProcessed(metricEndpointSlice, metricUpdate, true)
							return nil
						}, serializer.NoRetry)
					}
				}
			},
			DeleteFunc: func(obj interface{}) {
				var valid, equal bool
				defer func() { k.K8sEventReceived(metricEndpointSlice, metricDelete, valid, equal) }()
				k8sEP := k8s.CopyObjToV1EndpointSlice(obj)
				if k8sEP == nil {
					deletedObj, ok := obj.(cache.DeletedFinalStateUnknown)
					if !ok {
						return
					}
					// Delete was not observed by the watcher but is
					// removed from kube-apiserver. This is the last
					// known state and the object no longer exists.
					k8sEP = k8s.CopyObjToV1EndpointSlice(deletedObj.Obj)
					if k8sEP == nil {
						return
					}
				}
				valid = true
				swgEps.Add()
				serEps.Enqueue(func() error {
					defer swgEps.Done()
					k.K8sSvcCache.DeleteEndpointSlices(k8sEP, swgEps)
					k.K8sEventProcessed(metricEndpointSlice, metricDelete, true)
					return nil
				}, serializer.NoRetry)
			},
		},
		k8s.ConvertToK8sEndpointSlice,
	)
	ecr := make(chan struct{})
	k.blockWaitGroupToSyncResources(ecr, swgEps, endpointController, K8sAPIGroupEndpointSliceV1Beta1Discovery)
	go endpointController.Run(ecr)
	k.k8sAPIGroups.addAPI(K8sAPIGroupEndpointSliceV1Beta1Discovery)

	if k8s.HasEndpointSlice(hasEndpointSlices, endpointController) {
		return true
	}

	// K8s is not running with endpoint slices enabled, stop the endpoint slice
	// controller to avoid watching for unnecessary stuff in k8s.
	k.k8sAPIGroups.removeAPI(K8sAPIGroupEndpointSliceV1Beta1Discovery)
	close(ecr)
	return false
}
