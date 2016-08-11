//
// Copyright 2016 Authors of Cilium
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
//
package types

import (
	"bytes"
	"net"
	"testing"
	"time"

	"github.com/noironetworks/cilium-net/bpf/policymap"
	"github.com/noironetworks/cilium-net/common/addressing"

	. "gopkg.in/check.v1"
)

var (
	IPv6Addr, _ = addressing.NewCiliumIPv6("beef:beef:beef:beef:aaaa:aaaa:1111:1112")
	IPv4Addr, _ = addressing.NewCiliumIPv4("10.11.12.13")
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type EndpointSuite struct{}

var _ = Suite(&EndpointSuite{})

func (s *EndpointSuite) TestEndpointID(c *C) {
	e := Endpoint{IPv6: IPv6Addr, IPv4: IPv4Addr}
	e.SetID()
	c.Assert(e.ID, Equals, uint16(4370)) //"0x1112"
	c.Assert(bytes.Compare(e.IPv6, IPv6Addr) == 0, Equals, true)
	c.Assert(bytes.Compare(e.IPv4, IPv4Addr) == 0, Equals, true)
}

func (s *EndpointSuite) TestOrderEndpointAsc(c *C) {
	eps := []Endpoint{
		Endpoint{ID: 5},
		Endpoint{ID: 1000},
		Endpoint{ID: 1},
		Endpoint{ID: 3},
		Endpoint{ID: 2},
	}
	epsWant := []Endpoint{
		Endpoint{ID: 1},
		Endpoint{ID: 2},
		Endpoint{ID: 3},
		Endpoint{ID: 5},
		Endpoint{ID: 1000},
	}
	OrderEndpointAsc(eps)
	c.Assert(eps, DeepEquals, epsWant)
}

func (s *EndpointSuite) TestDeepCopy(c *C) {
	ipv4, err := addressing.NewCiliumIPv4("127.0.0.1")
	c.Assert(err, IsNil)
	ipv6, err := addressing.NewCiliumIPv6("::1")
	c.Assert(err, IsNil)
	epWant := &Endpoint{
		ID:               12,
		DockerID:         "123",
		DockerNetworkID:  "1234",
		DockerEndpointID: "12345",
		IfName:           "lxcifname",
		LXCMAC:           MAC{1, 2, 3, 4, 5, 6},
		IPv6:             ipv6,
		IPv4:             ipv4,
		IfIndex:          4,
		NodeMAC:          MAC{1, 2, 3, 4, 5, 6},
		NodeIP:           net.ParseIP("192.168.0.1"),
		PortMap:          make([]EPPortMap, 2),
		Opts:             NewBoolOptions(&EndpointOptionLibrary),
	}
	cpy := epWant.DeepCopy()
	c.Assert(*cpy, DeepEquals, *epWant)
	epWant.SecLabel = &SecCtxLabel{
		ID: 1,
		Labels: Labels{
			"io.cilium.kubernetes": NewLabel("io.cilium.kubernetes", "", "cilium"),
		},
		Containers: map[string]time.Time{
			"1234": time.Now(),
		},
	}
	epWant.Consumable = &Consumable{
		ID:        123,
		Iteration: 3,
		Labels:    nil,
		LabelList: []Label{
			*NewLabel("io.cilium.kubernetes", "", "cilium"),
		},
		Maps: map[int]*policymap.PolicyMap{
			0: &policymap.PolicyMap{},
		},
		Consumers: map[string]*Consumer{
			"foo": NewConsumer(12),
		},
		ReverseRules: map[uint32]*Consumer{
			12: NewConsumer(12),
		},
	}
	epWant.PolicyMap = &policymap.PolicyMap{}
	cpy = epWant.DeepCopy()
	c.Assert(*cpy.SecLabel, DeepEquals, *epWant.SecLabel)
	c.Assert(*cpy.Consumable, DeepEquals, *epWant.Consumable)
	c.Assert(*cpy.PolicyMap, DeepEquals, *epWant.PolicyMap)

	epWant.Consumable.Labels = &SecCtxLabel{
		ID: 1,
		Labels: Labels{
			"io.cilium.kubernetes": NewLabel("io.cilium.kubernetes", "", "cilium"),
		},
		Containers: map[string]time.Time{
			"1234": time.Now(),
		},
	}

	epWant.PolicyMap = &policymap.PolicyMap{}
	cpy = epWant.DeepCopy()

	c.Assert(*cpy.Consumable.Labels, DeepEquals, *epWant.Consumable.Labels)

	cpy.Consumable.Labels.Containers["1234"] = time.Now()
	c.Assert(*cpy.Consumable.Labels, Not(DeepEquals), *epWant.Consumable.Labels)
}
