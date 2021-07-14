// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/discovery/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDiscoveryV1beta1 struct {
	*testing.Fake
}

func (c *FakeDiscoveryV1beta1) EndpointSlices(namespace string) v1beta1.EndpointSliceInterface {
	return &FakeEndpointSlices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDiscoveryV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
