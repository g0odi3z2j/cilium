// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package bpf

import (
	"net"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
type BPFTestSuite struct{}

var _ = Suite(&BPFTestSuite{})

func (s *BPFTestSuite) TestEndpointKeyToString(c *C) {
	tests := []struct {
		ip string
	}{
		{"0.0.0.0"},
		{"192.0.2.3"},
		{"::"},
		{"fdff::ff"},
	}

	for _, tt := range tests {
		ip := net.ParseIP(tt.ip)
		k := NewEndpointKey(ip)
		c.Assert(k.ToIP().String(), Equals, tt.ip)
	}
}
