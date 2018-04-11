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

package ipcache

import (
	"net"
	"reflect"
	"testing"

	"github.com/cilium/cilium/pkg/comparator"
	identityPkg "github.com/cilium/cilium/pkg/identity"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
type IPCacheTestSuite struct{}

var _ = Suite(&IPCacheTestSuite{})

func Test(t *testing.T) {
	TestingT(t)
}

func (s *IPCacheTestSuite) TestIPCache(c *C) {
	endpointIP := "10.0.0.15"
	identity := (identityPkg.NumericIdentity(68))

	// Assure sane state at start.
	c.Assert(len(IPIdentityCache.ipToIdentityCache), Equals, 0)
	c.Assert(len(IPIdentityCache.identityToIPCache), Equals, 0)

	// Deletion of key that doesn't exist doesn't cause panic.
	IPIdentityCache.delete(endpointIP)

	IPIdentityCache.Upsert(endpointIP, identity)

	// Assure both caches are updated..
	c.Assert(len(IPIdentityCache.ipToIdentityCache), Equals, 1)
	c.Assert(len(IPIdentityCache.identityToIPCache), Equals, 1)

	cachedIdentity, exists := IPIdentityCache.LookupByIP(endpointIP)
	c.Assert(cachedIdentity, Equals, identity)
	c.Assert(exists, Equals, true)

	IPIdentityCache.Upsert(endpointIP, identity)

	// No duplicates.
	c.Assert(len(IPIdentityCache.ipToIdentityCache), Equals, 1)
	c.Assert(len(IPIdentityCache.identityToIPCache), Equals, 1)

	IPIdentityCache.delete(endpointIP)

	// Assure deletion occurs across both mappings.
	c.Assert(len(IPIdentityCache.ipToIdentityCache), Equals, 0)
	c.Assert(len(IPIdentityCache.identityToIPCache), Equals, 0)

	_, exists = IPIdentityCache.LookupByIP(endpointIP)

	c.Assert(exists, Equals, false)

	IPIdentityCache.Upsert(endpointIP, identity)

	newIdentity := identityPkg.NumericIdentity(69)
	IPIdentityCache.Upsert(endpointIP, newIdentity)

	// Ensure that update of cache with new identity doesn't keep old identity-to-ip
	// mapping around.
	_, exists = IPIdentityCache.LookupByIdentity(identity)
	c.Assert(exists, Equals, false)

	cachedIPSet, exists := IPIdentityCache.LookupByIdentity(newIdentity)
	c.Assert(exists, Equals, true)
	for cachedIP := range cachedIPSet {
		c.Assert(cachedIP, Equals, endpointIP)
	}

	IPIdentityCache.delete(endpointIP)

	// Assure deletion occurs across both mappings.
	c.Assert(len(IPIdentityCache.ipToIdentityCache), Equals, 0)
	c.Assert(len(IPIdentityCache.identityToIPCache), Equals, 0)

	// Test mapping of multiple IPs to same identity.
	endpointIPs := []string{"192.168.0.1", "20.3.75.3", "27.2.2.2", "127.0.0.1", "127.0.0.1"}
	identities := []identityPkg.NumericIdentity{5, 67, 29, 29, 29}

	for index := range endpointIPs {
		IPIdentityCache.Upsert(endpointIPs[index], identities[index])
		cachedIdentity, _ := IPIdentityCache.LookupByIP(endpointIPs[index])
		c.Assert(cachedIdentity, Equals, identities[index])
	}

	expectedIPList := map[string]struct{}{
		"27.2.2.2":  {},
		"127.0.0.1": {},
	}

	cachedEndpointIPs, _ := IPIdentityCache.LookupByIdentity(29)
	c.Assert(reflect.DeepEqual(cachedEndpointIPs, expectedIPList), Equals, true)

	IPIdentityCache.delete("27.2.2.2")

	expectedIPList = map[string]struct{}{
		"127.0.0.1": {},
	}

	cachedEndpointIPs, _ = IPIdentityCache.LookupByIdentity(29)
	c.Assert(reflect.DeepEqual(cachedEndpointIPs, expectedIPList), Equals, true)

	cachedIdentity, exists = IPIdentityCache.LookupByIP("127.0.0.1")
	c.Assert(cachedIdentity, Equals, identityPkg.NumericIdentity(29))

	IPIdentityCache.delete("127.0.0.1")

	_, exists = IPIdentityCache.LookupByIdentity(29)
	c.Assert(exists, Equals, false)

	// Clean up.
	for index := range endpointIPs {
		IPIdentityCache.delete(endpointIPs[index])
		_, exists = IPIdentityCache.LookupByIP(endpointIPs[index])
		c.Assert(exists, Equals, false)

		_, exists = IPIdentityCache.LookupByIdentity(identities[index])
		c.Assert(exists, Equals, false)
	}

	c.Assert(len(IPIdentityCache.ipToIdentityCache), Equals, 0)
	c.Assert(len(IPIdentityCache.identityToIPCache), Equals, 0)

}

func (s *IPCacheTestSuite) TestKeyToIPNet(c *C) {
	// Valid IPv6.
	validIPv6Key := "cilium/state/ip/v1/default/f00d::a00:0:0:c164"

	_, expectedIPv6, err := net.ParseCIDR("f00d::a00:0:0:c164/128")
	c.Assert(err, IsNil)

	ipv6, isHost, err := keyToIPNet(validIPv6Key)
	c.Assert(ipv6, Not(IsNil))
	c.Assert(err, IsNil)
	c.Assert(isHost, Equals, true)
	c.Assert(ipv6, comparator.DeepEquals, expectedIPv6)

	// Valid IPv6 prefix.
	validIPv6Key = "cilium/state/ip/v1/default/f00d::a00:0:0:0/64"

	_, expectedIPv6, err = net.ParseCIDR("f00d::a00:0:0:0/64")
	c.Assert(err, IsNil)

	ipv6, isHost, err = keyToIPNet(validIPv6Key)
	c.Assert(ipv6, Not(IsNil))
	c.Assert(err, IsNil)
	c.Assert(isHost, Equals, false)
	c.Assert(ipv6, comparator.DeepEquals, expectedIPv6)

	// Valid IPv4.
	validIPv4Key := "cilium/state/ip/v1/default/10.0.114.197"
	_, expectedIPv4, err := net.ParseCIDR("10.0.114.197/32")
	c.Assert(err, IsNil)
	ipv4, isHost, err := keyToIPNet(validIPv4Key)
	c.Assert(ipv4, Not(IsNil))
	c.Assert(err, IsNil)
	c.Assert(isHost, Equals, true)
	c.Assert(ipv4, comparator.DeepEquals, expectedIPv4)

	// Valid IPv4 prefix.
	validIPv4Key = "cilium/state/ip/v1/default/10.0.114.0/24"
	_, expectedIPv4, err = net.ParseCIDR("10.0.114.0/24")
	c.Assert(err, IsNil)
	ipv4, isHost, err = keyToIPNet(validIPv4Key)
	c.Assert(ipv4, Not(IsNil))
	c.Assert(err, IsNil)
	c.Assert(isHost, Equals, false)
	c.Assert(ipv4, comparator.DeepEquals, expectedIPv4)

	// Invalid prefix.
	invalidPrefixKey := "cilium/state/foobar/v1/default/f00d::a00:0:0:c164"
	nilIP, isHost, err := keyToIPNet(invalidPrefixKey)
	c.Assert(nilIP, IsNil)
	c.Assert(err, Not(IsNil))
	c.Assert(isHost, Equals, false)

	// Invalid IP in key.
	invalidIPKey := "cilium/state/ip/v1/default/10.abfd.114.197"
	nilIP, isHost, err = keyToIPNet(invalidIPKey)
	c.Assert(nilIP, IsNil)
	c.Assert(err, Not(IsNil))
	c.Assert(isHost, Equals, false)

	// Invalid CIDR.
	invalidIPKey = "cilium/state/ip/v1/default/192.0.2.3/54"
	nilIP, isHost, err = keyToIPNet(invalidIPKey)
	c.Assert(nilIP, IsNil)
	c.Assert(err, Not(IsNil))
	c.Assert(isHost, Equals, false)
}
