// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fqdn

import (
	"context"
	"fmt"
	"net"
	"net/netip"
	"testing"

	"github.com/cilium/cilium/pkg/defaults"
	"github.com/cilium/cilium/pkg/fqdn/dns"
	"github.com/cilium/cilium/pkg/fqdn/re"
	"github.com/cilium/cilium/pkg/policy/api"
	testipcache "github.com/cilium/cilium/pkg/testutils/ipcache"
	"github.com/cilium/cilium/pkg/time"
)

// BenchmarkUpdateGenerateDNS tests updating a large number of selectors.
//
// Run it like
// go test -benchmem -run=^$ -bench ^BenchmarkUpdateGeneratedDNS$ github.com/cilium/cilium/pkg/fqdn -benchtime=4x -count=10
func BenchmarkUpdateGenerateDNS(b *testing.B) {

	// For every i in range 1 .. K, create selectors
	// - "$K.example.com"
	// - "*.$K.example.com"
	// as well as *.example.com.
	//
	// Then, update the generated DNS N * K times, setting i to N % K, and updating
	// - $i.example.com
	// - foo.$i.example.com
	// with a random new IP

	numSelectors := 1000

	re.InitRegexCompileLRU(defaults.FQDNRegexCompileLRUSize)

	nameManager := NewNameManager(Config{
		MinTTL:  1,
		Cache:   NewDNSCache(0),
		IPCache: testipcache.NewMockIPCache(),
	})

	for i := 0; i < numSelectors; i++ {
		nameManager.RegisterForIPUpdatesLocked(api.FQDNSelector{
			MatchName: fmt.Sprintf("%d.example.com", i),
		})
		nameManager.RegisterForIPUpdatesLocked(api.FQDNSelector{
			MatchPattern: fmt.Sprintf("*.%d.example.com", i),
		})
	}
	nameManager.RegisterForIPUpdatesLocked(api.FQDNSelector{
		MatchPattern: "*.example.com",
	})

	t := time.Now() // doesn't matter, just need a stable base
	ip := netip.MustParseAddr("10.0.0.0")

	b.ResetTimer() // Don't benchmark adding selectors, just evaluating them
	for i := 0; i < b.N*numSelectors; i++ {
		t = t.Add(1 * time.Second)
		ip = ip.Next()

		k := i % numSelectors
		nameManager.UpdateGenerateDNS(context.Background(), t, map[string]*DNSIPRecords{
			dns.FQDN(fmt.Sprintf("%d.example.com", k)): {
				TTL: 60,
				IPs: []net.IP{ip.AsSlice()},
			}})

		nameManager.UpdateGenerateDNS(context.Background(), t, map[string]*DNSIPRecords{
			dns.FQDN(fmt.Sprintf("example.%d.example.com", k)): {
				TTL: 60,
				IPs: []net.IP{ip.AsSlice()},
			}})
	}
}
