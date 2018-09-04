// Copyright 2016-2017 Authors of Cilium
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

package k8s

import (
	"fmt"
	"net"
	"time"

	"github.com/cilium/cilium/pkg/annotation"
	"github.com/cilium/cilium/pkg/comparator"
	"github.com/cilium/cilium/pkg/node"

	. "gopkg.in/check.v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/testing"
)

func (s *K8sSuite) TestUseNodeCIDR(c *C) {
	// Test IPv4
	node1 := v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node1",
			Annotations: map[string]string{
				annotation.V4CIDRName:   "10.254.0.0/16",
				annotation.CiliumHostIP: "10.254.0.1",
			},
		},
		Spec: v1.NodeSpec{
			PodCIDR: "10.2.0.0/16",
		},
	}

	// set buffer to 2 to prevent blocking when calling UseNodeCIDR
	// and we need to wait for the response of the channel.
	updateChan := make(chan bool, 2)
	k8sClient := &fake.Clientset{}
	k8sClient.AddReactor("get", "nodes",
		func(action testing.Action) (bool, runtime.Object, error) {
			name := action.(testing.GetAction).GetName()
			c.Assert(name, Equals, "node1")
			return true, node1.DeepCopy(), nil
		})
	k8sClient.AddReactor("update", "nodes",
		func(action testing.Action) (bool, runtime.Object, error) {
			n := action.(testing.UpdateAction).GetObject().(*v1.Node)
			n1copy := node1.DeepCopy()
			n1copy.Annotations[annotation.V4CIDRName] = "10.2.0.0/16"
			c.Assert(n, comparator.DeepEquals, n1copy)
			updateChan <- true
			return true, n1copy, nil
		})
	node1Cilium := ParseNode(&node1, node.FromAgentLocal)

	err := node.UseNodeCIDR(node1Cilium)
	c.Assert(err, IsNil)
	c.Assert(node.GetIPv4AllocRange().String(), Equals, "10.2.0.0/16")
	// IPv6 Node range is not checked because it shouldn't be changed.

	AnnotateNode(k8sClient, "node1",
		node.GetIPv4AllocRange(),
		node.GetIPv6NodeRange(),
		nil,
		nil,
		net.ParseIP("10.254.0.1"))

	c.Assert(err, IsNil)

	select {
	case <-updateChan:
	case <-time.Tick(10 * time.Second):
		c.Errorf("d.k8sClient.CoreV1().Nodes().Update() was not called")
		c.FailNow()
	}

	// Test IPv6
	node2 := v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node2",
			Annotations: map[string]string{
				annotation.V4CIDRName:   "10.254.0.0/16",
				annotation.CiliumHostIP: "10.254.0.1",
			},
		},
		Spec: v1.NodeSpec{
			PodCIDR: "aaaa:aaaa:aaaa:aaaa:beef:beef::/96",
		},
	}

	failAttempts := 0
	k8sClient = &fake.Clientset{}
	k8sClient.AddReactor("get", "nodes",
		func(action testing.Action) (bool, runtime.Object, error) {
			name := action.(testing.GetAction).GetName()
			c.Assert(name, Equals, "node2")
			return true, node2.DeepCopy(), nil
		})
	k8sClient.AddReactor("update", "nodes",
		func(action testing.Action) (bool, runtime.Object, error) {
			n := action.(testing.UpdateAction).GetObject().(*v1.Node)
			if failAttempts == 0 {
				failAttempts++
				return true, nil, fmt.Errorf("failing on purpose")
			}
			n2Copy := node2.DeepCopy()
			n2Copy.Annotations[annotation.V4CIDRName] = "10.254.0.0/16"
			n2Copy.Annotations[annotation.V6CIDRName] = "aaaa:aaaa:aaaa:aaaa:beef:beef::/96"
			c.Assert(n, comparator.DeepEquals, n2Copy)
			updateChan <- true
			return true, n2Copy, nil
		})

	node2Cilium := ParseNode(&node2, node.FromAgentLocal)
	err = node.UseNodeCIDR(node2Cilium)
	c.Assert(err, IsNil)

	// We use the node's annotation for the IPv4 and the PodCIDR for the
	// IPv6.
	c.Assert(node.GetIPv4AllocRange().String(), Equals, "10.254.0.0/16")
	c.Assert(node.GetIPv6NodeRange().String(), Equals, "aaaa:aaaa:aaaa:aaaa:beef:beef::/96")

	err = AnnotateNode(k8sClient, "node2",
		node.GetIPv4AllocRange(),
		node.GetIPv6NodeRange(),
		nil,
		nil,
		net.ParseIP("10.254.0.1"))

	c.Assert(err, IsNil)

	select {
	case <-updateChan:
	case <-time.Tick(10 * time.Second):
		c.Errorf("d.k8sClient.CoreV1().Nodes().Update() was not called")
		c.FailNow()
	}

}
