// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned"
	corev1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/core/v1"
	fakecorev1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/core/v1/fake"
	discoveryv1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/discovery/v1"
	fakediscoveryv1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/discovery/v1/fake"
	discoveryv1beta1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/discovery/v1beta1"
	fakediscoveryv1beta1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/discovery/v1beta1/fake"
	metav1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/networking/v1"
	fakemetav1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/networking/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return &fakecorev1.FakeCoreV1{Fake: &c.Fake}
}

// DiscoveryV1beta1 retrieves the DiscoveryV1beta1Client
func (c *Clientset) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface {
	return &fakediscoveryv1beta1.FakeDiscoveryV1beta1{Fake: &c.Fake}
}

// DiscoveryV1 retrieves the DiscoveryV1Client
func (c *Clientset) DiscoveryV1() discoveryv1.DiscoveryV1Interface {
	return &fakediscoveryv1.FakeDiscoveryV1{Fake: &c.Fake}
}

// MetaV1 retrieves the MetaV1Client
func (c *Clientset) MetaV1() metav1.MetaV1Interface {
	return &fakemetav1.FakeMetaV1{Fake: &c.Fake}
}
