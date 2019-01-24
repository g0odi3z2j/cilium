// Copyright 2018 Authors of Cilium
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

// +build !privileged_test

package ipam

import (
	"net"

	"github.com/cilium/cilium/pkg/datapath/fake"

	. "gopkg.in/check.v1"
)

func (s *IPAMSuite) TestAllocatedIPDump(c *C) {
	ipam := NewIPAM(fake.NewNodeAddressing(), Configuration{EnableIPv4: true, EnableIPv6: true})

	err := ipam.AllocateInternalIPs()
	c.Assert(err, IsNil)

	allocv4, allocv6 := ipam.Dump()
	// Test the format of the dumped ip addresses
	for i := 0; i < len(allocv4); i++ {
		c.Assert(net.ParseIP(allocv4[i]), NotNil)
	}
	for i := 0; i < len(allocv6); i++ {
		c.Assert(net.ParseIP(allocv6[i]), NotNil)
	}
}
