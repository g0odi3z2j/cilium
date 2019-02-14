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

// +build !privileged_tests

package fqdn

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"regexp"
	"sort"
	"time"

	"github.com/cilium/cilium/pkg/checker"
	. "gopkg.in/check.v1"
)

type DNSCacheTestSuite struct{}

var _ = Suite(&DNSCacheTestSuite{})

func sortByName(entries []*cacheEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})
}

// TestUpdateLookup tests that we can insert DNS data and retrieve it. We
// iterate through time, ensuring that data is expired as appropriate. We also
// insert redundant DNS entries that should not change the output.
func (ds *DNSCacheTestSuite) TestUpdateLookup(c *C) {
	name := "test.com"
	now := time.Now()
	cache := NewDNSCache()
	endTimeSeconds := 4

	// Add 1 new entry "per second", and one with a redundant IP (with ttl/2).
	// The IP reflects the second in which it will expire, and should show up for
	// all now+ttl that is less than it.
	for i := 1; i <= endTimeSeconds; i++ {
		ttl := i
		cache.Update(now,
			name,
			[]net.IP{net.ParseIP(fmt.Sprintf("1.1.1.%d", i)), net.ParseIP(fmt.Sprintf("2.2.2.%d", i))},
			ttl)

		cache.Update(now,
			name,
			[]net.IP{net.ParseIP(fmt.Sprintf("1.1.1.%d", i))},
			ttl/2)
	}

	// lookup our entries
	//  - no redundant entries (the 1.1.1.x is not repeated)
	//  - with each step of secondsPastNow, fewer entries are returned
	for secondsPastNow := 1; secondsPastNow <= endTimeSeconds; secondsPastNow++ {
		ips := cache.lookupByTime(now.Add(time.Duration(secondsPastNow)*time.Second), name)
		c.Assert(len(ips), Equals, 2*(endTimeSeconds-secondsPastNow+1), Commentf("Incorrect number of IPs returned"))

		// Check that we returned each 1.1.1.x entry where x={1..endTimeSeconds}
		// These are sorted, and are in the first half of the array
		// Similarly, check the 2.2.2.x entries in the second half of the array
		j := secondsPastNow
		halfIndex := endTimeSeconds - secondsPastNow + 1
		for _, ip := range ips[:halfIndex] {
			c.Assert(ip.String(), Equals, fmt.Sprintf("1.1.1.%d", j), Commentf("Incorrect IP returned (j=%d, secondsPastNow=%d)", j, secondsPastNow))
			j++
		}
		j = secondsPastNow
		for _, ip := range ips[halfIndex:] {
			c.Assert(ip.String(), Equals, fmt.Sprintf("2.2.2.%d", j), Commentf("Incorrect IP returned (j=%d, secondsPastNow=%d)", j, secondsPastNow))
			j++
		}
	}
}

// TestDelete tests that we can forcibly clear parts of the cache.
func (ds *DNSCacheTestSuite) TestDelete(c *C) {
	names := map[string]net.IP{
		"test1.com": net.ParseIP("2.2.2.1"),
		"test2.com": net.ParseIP("2.2.2.2"),
		"test3.com": net.ParseIP("2.2.2.3")}
	sharedIP := net.ParseIP("1.1.1.1")
	now := time.Now()
	cache := NewDNSCache()

	// Insert 3 records with 1 shared IP and 3 with different IPs
	cache.Update(now, "test1.com", []net.IP{sharedIP, names["test1.com"]}, 5)
	cache.Update(now, "test2.com", []net.IP{sharedIP, names["test2.com"]}, 5)
	cache.Update(now, "test3.com", []net.IP{sharedIP, names["test3.com"]}, 5)

	now = now.Add(time.Second)

	// Test that a non-matching ForceExpire doesn't do anything. All data should
	// still be present.
	nameMatch, err := regexp.Compile("^notatest.com$")
	c.Assert(err, IsNil)
	namesAffected := cache.ForceExpire(now, nameMatch)
	c.Assert(len(namesAffected), Equals, 0, Commentf("Incorrect count of names removed %v", namesAffected))
	for _, name := range []string{"test1.com", "test2.com", "test3.com"} {
		ips := cache.lookupByTime(now, name)
		c.Assert(len(ips), Equals, 2, Commentf("Wrong count of IPs returned (%v) for non-deleted name '%s'", ips, name))
	}

	// Delete a single name and check that
	// - It is returned in namesAffected
	// - Lookups for it show no data, but data remains for other names
	nameMatch, err = regexp.Compile("^test1.com$")
	c.Assert(err, IsNil)
	namesAffected = cache.ForceExpire(now, nameMatch)
	c.Assert(len(namesAffected), Equals, 1, Commentf("Incorrect count of names removed %v", namesAffected))
	c.Assert(namesAffected[0], Equals, "test1.com", Commentf("Incorrect affected name returned on forced expire: %s", namesAffected))
	ips := cache.lookupByTime(now, "test1.com")
	c.Assert(len(ips), Equals, 0, Commentf("IPs returned (%v) for deleted name 'test1.com'", ips))
	for _, name := range []string{"test2.com", "test3.com"} {
		ips = cache.lookupByTime(now, name)
		c.Assert(len(ips), Equals, 2, Commentf("Wrong count of IPs returned (%v) for non-deleted name '%s'", ips, name))
	}

	// Delete the whole cache. This should leave no data.
	namesAffected = cache.ForceExpire(now, nil)
	sort.Strings(namesAffected) // simplify the checks below
	c.Assert(len(namesAffected), Equals, 2, Commentf("Incorrect count of names removed %v", namesAffected))
	for i, name := range []string{"test2.com", "test3.com"} {
		c.Assert(namesAffected[i], Equals, name, Commentf("Incorrect affected name returned on forced expire"))
	}
	for name := range names {
		ips = cache.lookupByTime(now, name)
		c.Assert(len(ips), Equals, 0, Commentf("Returned IP data for %s after the cache was fully cleared: %v", name, ips))
	}
	dump := cache.Dump()
	c.Assert(len(dump), Equals, 0, Commentf("Returned cache entries from cache dump after the cache was fully cleared: %v", dump))
}

func (ds *DNSCacheTestSuite) TestForceExpiredByNames(c *C) {
	names := []string{"test1.com", "test2.com"}
	cache := NewDNSCache()
	for i := 1; i < 4; i++ {
		cache.Update(
			now,
			fmt.Sprintf("test%d.com", i),
			[]net.IP{net.ParseIP(fmt.Sprintf("1.1.1.%d", i))},
			5)
	}

	c.Assert(cache.forward, HasLen, 3)
	result := cache.ForceExpireByNames(time.Now(), names)
	c.Assert(result, checker.DeepEquals, names)
	c.Assert(result, HasLen, 2)
	c.Assert(cache.forward["test3.com"], Not(IsNil))

	invalidName := cache.ForceExpireByNames(now, []string{"invalid.name"})
	c.Assert(invalidName, HasLen, 0)
}

func (ds *DNSCacheTestSuite) TestReverseUpdateLookup(c *C) {
	names := map[string]net.IP{
		"test1.com": net.ParseIP("2.2.2.1"),
		"test2.com": net.ParseIP("2.2.2.2"),
		"test3.com": net.ParseIP("2.2.2.3")}
	sharedIP := net.ParseIP("1.1.1.1")
	now := time.Now()
	cache := NewDNSCache()

	// insert 2 records, with 1 shared IP
	cache.Update(now, "test1.com", []net.IP{sharedIP, names["test1.com"]}, 2)
	cache.Update(now, "test2.com", []net.IP{sharedIP, names["test2.com"]}, 4)

	// lookup within the TTL for both names should return 2 names for sharedIPs,
	// and one name for the 2.2.2.* IPs
	currentTime := now.Add(time.Second)
	lookupNames := cache.lookupIPByTime(currentTime, sharedIP)
	c.Assert(len(lookupNames), Equals, 2, Commentf("Incorrect number of names returned"))
	for _, name := range lookupNames {
		_, found := names[name]
		c.Assert(found, Equals, true, Commentf("Returned a DNS name that doesn't match IP"))
	}

	lookupNames = cache.lookupIPByTime(currentTime, names["test1.com"])
	c.Assert(len(lookupNames), Equals, 1, Commentf("Incorrect number of names returned"))
	c.Assert(lookupNames[0], Equals, "test1.com", Commentf("Returned a DNS name that doesn't match IP"))

	lookupNames = cache.lookupIPByTime(currentTime, names["test2.com"])
	c.Assert(len(lookupNames), Equals, 1, Commentf("Incorrect number of names returned"))
	c.Assert(lookupNames[0], Equals, "test2.com", Commentf("Returned a DNS name that doesn't match IP"))

	lookupNames = cache.lookupIPByTime(currentTime, names["test3.com"])
	c.Assert(len(lookupNames), Equals, 0, Commentf("Returned names for IP not in cache"))

	// lookup between 2-4 seconds later (test1.com has expired) for both names
	// should return 2 names for sharedIPs, and one name for the 2.2.2.* IPs
	currentTime = now.Add(3 * time.Second)
	lookupNames = cache.lookupIPByTime(currentTime, sharedIP)
	c.Assert(len(lookupNames), Equals, 1, Commentf("Incorrect number of names returned"))
	c.Assert(lookupNames[0], Equals, "test2.com", Commentf("Returned a DNS name that doesn't match IP"))

	lookupNames = cache.lookupIPByTime(currentTime, names["test1.com"])
	c.Assert(len(lookupNames), Equals, 0, Commentf("Incorrect number of names returned"))

	lookupNames = cache.lookupIPByTime(currentTime, names["test2.com"])
	c.Assert(len(lookupNames), Equals, 1, Commentf("Incorrect number of names returned"))
	c.Assert(lookupNames[0], Equals, "test2.com", Commentf("Returned a DNS name that doesn't match IP"))

	lookupNames = cache.lookupIPByTime(currentTime, names["test3.com"])
	c.Assert(len(lookupNames), Equals, 0, Commentf("Returned names for IP not in cache"))

	// lookup between after 4 seconds later (all have expired) for both names
	// should return no names in all cases.
	currentTime = now.Add(5 * time.Second)
	lookupNames = cache.lookupIPByTime(currentTime, sharedIP)
	c.Assert(len(lookupNames), Equals, 0, Commentf("Incorrect number of names returned"))

	lookupNames = cache.lookupIPByTime(currentTime, names["test1.com"])
	c.Assert(len(lookupNames), Equals, 0, Commentf("Incorrect number of names returned"))

	lookupNames = cache.lookupIPByTime(currentTime, names["test2.com"])
	c.Assert(len(lookupNames), Equals, 0, Commentf("Incorrect number of names returned"))

	lookupNames = cache.lookupIPByTime(currentTime, names["test3.com"])
	c.Assert(len(lookupNames), Equals, 0, Commentf("Returned names for IP not in cache"))
}

func (ds *DNSCacheTestSuite) TestJSONMarshal(c *C) {
	names := map[string]net.IP{
		"test1.com": net.ParseIP("2.2.2.1"),
		"test2.com": net.ParseIP("2.2.2.2"),
		"test3.com": net.ParseIP("2.2.2.3")}
	sharedIP := net.ParseIP("1.1.1.1")
	now := time.Now()
	cache := NewDNSCache()

	// insert 3 records with 1 shared IP and 3 with different IPs
	cache.Update(now, "test1.com", []net.IP{sharedIP}, 5)
	cache.Update(now, "test2.com", []net.IP{sharedIP}, 5)
	cache.Update(now, "test3.com", []net.IP{sharedIP}, 5)
	cache.Update(now, "test1.com", []net.IP{names["test1.com"]}, 5)
	cache.Update(now, "test2.com", []net.IP{names["test2.com"]}, 5)
	cache.Update(now, "test3.com", []net.IP{names["test3.com"]}, 5)

	// Marshal and unmarshal
	data, err := cache.MarshalJSON()
	c.Assert(err, IsNil)

	newCache := NewDNSCache()
	err = newCache.UnmarshalJSON(data)
	c.Assert(err, IsNil)

	// Marshalled data should have no duplicate entries Note: this is tightly
	// coupled with the implementation of DNSCache.MarshalJSON because the
	// unmarshalled instance will hide duplicates. We simply check the length
	// since we control the inserted data, and we test its correctness below.
	rawList := make([]*cacheEntry, 0)
	err = json.Unmarshal(data, &rawList)
	c.Assert(err, IsNil)
	c.Assert(len(rawList), Equals, 6)

	// Check that the unmarshalled instance contains all the data at now
	currentTime := now
	for name := range names {
		IPs := cache.lookupByTime(currentTime, name)
		c.Assert(len(IPs), Equals, 2, Commentf("Incorrect number of IPs returned for %s", name))
		c.Assert(IPs[0].String(), Equals, sharedIP.String(), Commentf("Returned an IP that doesn't match %s", name))
		c.Assert(IPs[1].String(), Equals, names[name].String(), Commentf("Returned an IP name that doesn't match %s", name))
	}

	// Check that the unmarshalled data expires correctly
	currentTime = now.Add(10 * time.Second)
	for name := range names {
		IPs := cache.lookupByTime(currentTime, name)
		c.Assert(len(IPs), Equals, 0, Commentf("Returned IPs that should be expired for %s", name))
	}
}

/* Benchmarks
 * These are here to help gauge the relative costs of operations in DNSCache.
 * Note: some are on arrays `size` elements, so the benchmark "op time" is too
 * large.
 */

var (
	now         = time.Now()
	size        = uint32(1000) // size of array to operate on
	entriesOrig = makeEntries(now, 1+size/3, 1+size/3, 1+size/3)
	ipsOrig     = makeIPs(size)
)

// makeIPs generates count sequential IPv4 IPs
func makeIPs(count uint32) []net.IP {
	ips := make([]net.IP, 0, count)
	for i := uint32(0); i < count; i++ {
		ips = append(ips, net.IPv4(byte(i>>24), byte(i>>16), byte(i>>8), byte(i>>0)))
	}
	return ips
}

func makeEntries(now time.Time, live, redundant, expired uint32) (entries []*cacheEntry) {
	liveTTL := 120
	redundantTTL := 60

	for ; live > 0; live-- {
		ip := net.IPv4(byte(live>>24), byte(live>>16), byte(live>>8), byte(live>>0))

		entries = append(entries, &cacheEntry{
			Name:           fmt.Sprintf("live-%s", ip.String()),
			LookupTime:     now,
			ExpirationTime: now.Add(time.Duration(liveTTL) * time.Second),
			TTL:            liveTTL,
			IPs:            []net.IP{ip}})

		if redundant > 0 {
			redundant--
			entries = append(entries, &cacheEntry{
				Name:           fmt.Sprintf("redundant-%s", ip.String()),
				LookupTime:     now,
				ExpirationTime: now.Add(time.Duration(redundantTTL) * time.Second),
				TTL:            redundantTTL,
				IPs:            []net.IP{ip}})
		}

		if expired > 0 {
			expired--
			entries = append(entries, &cacheEntry{
				Name:           fmt.Sprintf("expired-%s", ip.String()),
				LookupTime:     now.Add(-time.Duration(liveTTL) * time.Second),
				ExpirationTime: now.Add(-time.Second),
				TTL:            liveTTL,
				IPs:            []net.IP{ip}})
		}
	}

	rand.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})

	return entries
}

// Note: each "op" works on size things
func (ds *DNSCacheTestSuite) BenchmarkGetIPs(c *C) {
	c.StopTimer()
	now := time.Now()
	cache := NewDNSCache()
	cache.Update(now, "test.com", []net.IP{net.ParseIP("1.2.3.4")}, 60)
	entries := cache.forward["test.com"]
	for _, entry := range entriesOrig {
		cache.updateWithEntryIPs(entries, entry)
	}
	c.StartTimer()

	for i := 0; i < c.N; i++ {
		entries.getIPs(now)
	}
}

// Note: each "op" works on size things
func (ds *DNSCacheTestSuite) BenchmarkUpdateIPs(c *C) {
	for i := 0; i < c.N; i++ {
		c.StopTimer()
		now := time.Now()
		cache := NewDNSCache()
		cache.Update(now, "test.com", []net.IP{net.ParseIP("1.2.3.4")}, 60)
		entries := cache.forward["test.com"]
		c.StartTimer()

		for _, entry := range entriesOrig {
			cache.updateWithEntryIPs(entries, entry)
			cache.removeExpired(entries, now, time.Time{})
		}
	}
}

func (ds *DNSCacheTestSuite) BenchmarkIPString(c *C) {
	for i := 0; i < c.N; i++ {
		_ = net.IPv4(byte(i>>24), byte(i>>16), byte(i>>8), byte(i>>0)).String()
	}
}

func (ds *DNSCacheTestSuite) BenchmarkParseIPSimple(c *C) {
	ip := ipsOrig[0].String()
	for i := 0; i < c.N; i++ {
		_ = net.ParseIP(ip)
	}
}

// Note: each "op" works on size things
func (ds *DNSCacheTestSuite) BenchmarkParseIP(c *C) {
	c.StopTimer()
	ips := make([]string, 0, len(ipsOrig))
	for _, ip := range ipsOrig {
		ips = append(ips, ip.String())
	}
	c.StartTimer()

	for i := 0; i < c.N; i++ {
		for _, ipStr := range ips {
			_ = net.ParseIP(ipStr)
		}
	}
}

// JSON Marshal/Unmarshal benchmarks
var numIPsPerEntry = 10 // number of IPs to generate in each entry

func (ds *DNSCacheTestSuite) BenchmarkMarshalJSON10(c *C)    { benchmarkMarshalJSON(c, 10) }
func (ds *DNSCacheTestSuite) BenchmarkMarshalJSON100(c *C)   { benchmarkMarshalJSON(c, 100) }
func (ds *DNSCacheTestSuite) BenchmarkMarshalJSON1000(c *C)  { benchmarkMarshalJSON(c, 1000) }
func (ds *DNSCacheTestSuite) BenchmarkMarshalJSON10000(c *C) { benchmarkMarshalJSON(c, 10000) }

func (ds *DNSCacheTestSuite) BenchmarkUnmarshalJSON10(c *C)    { benchmarkUnmarshalJSON(c, 10) }
func (ds *DNSCacheTestSuite) BenchmarkUnmarshalJSON100(c *C)   { benchmarkUnmarshalJSON(c, 100) }
func (ds *DNSCacheTestSuite) BenchmarkUnmarshalJSON1000(c *C)  { benchmarkUnmarshalJSON(c, 1000) }
func (ds *DNSCacheTestSuite) BenchmarkUnmarshalJSON10000(c *C) { benchmarkUnmarshalJSON(c, 10000) }

// BenchmarkMarshalJSON100Repeat2 tests whether repeating the whole
// serialization is notably slower than a single run.
func (ds *DNSCacheTestSuite) BenchmarkMarshalJSON100Repeat2(c *C) {
	benchmarkMarshalJSON(c, 50)
	benchmarkMarshalJSON(c, 50)
}

func (ds *DNSCacheTestSuite) BenchmarkMarshalJSON1000Repeat2(c *C) {
	benchmarkMarshalJSON(c, 500)
	benchmarkMarshalJSON(c, 500)
}

// benchmarkMarshalJSON benchmarks the cost of creating a json representation
// of DNSCache. Each benchmark "op" is on numDNSEntries.
// Note: It assumes the JSON only uses data in DNSCache.forward when generating
// the data. Changes to the implementation need to also change this benchmark.
func benchmarkMarshalJSON(c *C, numDNSEntries int) {
	c.StopTimer()
	ips := makeIPs(uint32(numIPsPerEntry))

	cache := NewDNSCache()
	for i := 0; i < numDNSEntries; i++ {
		// TTL needs to be far enough in the future that the entry is serialized
		cache.Update(time.Now(), fmt.Sprintf("domain-%v.com", i), ips, 86400)
	}
	c.StartTimer()

	for i := 0; i < c.N; i++ {
		_, err := cache.MarshalJSON()
		c.Assert(err, IsNil)
	}
}

// benchmarkUnmarshalJSON benchmarks the cost of parsing a json representation
// of DNSCache. Each benchmark "op" is on numDNSEntries.
// Note: It assumes the JSON only uses data in DNSCache.forward when generating
// the data. Changes to the implementation need to also change this benchmark.
func benchmarkUnmarshalJSON(c *C, numDNSEntries int) {
	c.StopTimer()
	ips := makeIPs(uint32(numIPsPerEntry))

	cache := NewDNSCache()
	for i := 0; i < numDNSEntries; i++ {
		// TTL needs to be far enough in the future that the entry is serialized
		cache.Update(time.Now(), fmt.Sprintf("domain-%v.com", i), ips, 86400)
	}

	data, err := cache.MarshalJSON()
	c.Assert(err, IsNil)

	emptyCaches := make([]*DNSCache, c.N)
	for i := 0; i < c.N; i++ {
		emptyCaches[i] = NewDNSCache()
	}
	c.StartTimer()

	for i := 0; i < c.N; i++ {
		err := emptyCaches[i].UnmarshalJSON(data)
		c.Assert(err, IsNil)
	}
}

func (ds *DNSCacheTestSuite) TestTTLCleanupEntries(c *C) {
	cache := NewDNSCache()
	cache.Update(now, "test.com", []net.IP{net.ParseIP("1.2.3.4")}, 3)
	c.Assert(len(cache.cleanup), Equals, 1)
	entries, _ := cache.cleanupExpiredEntries(time.Now().Add(5 * time.Second))
	c.Assert(entries, HasLen, 1)
	c.Assert(cache.cleanup, HasLen, 0)
	c.Assert(cache.Lookup("test.com"), HasLen, 0)
}

func (ds *DNSCacheTestSuite) TestTTLCleanupWithoutForward(c *C) {
	cache := NewDNSCache()
	now := time.Now()
	cache.cleanup[now.Unix()] = []string{"test.com"}
	// To make sure that all entries are validated correctly
	cache.lastCleanup = time.Now().Add(-1 * time.Minute)
	entries, _ := cache.cleanupExpiredEntries(time.Now().Add(5 * time.Second))
	c.Assert(entries, HasLen, 0)
	c.Assert(cache.cleanup, HasLen, 0)
}

func (ds *DNSCacheTestSuite) TestOverlimitEntriesWithValidLimit(c *C) {
	limit := 5
	cache := NewDNSCacheWithLimit(limit)

	cache.Update(now, "foo.bar", []net.IP{net.ParseIP("1.1.1.1")}, 1)
	cache.Update(now, "bar.foo", []net.IP{net.ParseIP("2.1.1.1")}, 1)
	for i := 1; i < limit+2; i++ {
		cache.Update(now, "test.com", []net.IP{net.ParseIP(fmt.Sprintf("1.1.1.%d", i))}, i)
	}
	c.Assert(cache.cleanupOverLimitEntries(), checker.DeepEquals, []string{"test.com"})
	c.Assert(cache.Lookup("test.com"), HasLen, limit)
	c.Assert(cache.LookupIP(net.ParseIP("1.1.1.1")), checker.DeepEquals, []string{"foo.bar"})
	c.Assert(cache.forward["test.com"]["1.1.1.1"], IsNil)
	c.Assert(cache.Lookup("foo.bar"), HasLen, 1)
	c.Assert(cache.Lookup("bar.foo"), HasLen, 1)
	c.Assert(cache.overLimit, HasLen, 0)
}

func (ds *DNSCacheTestSuite) TestOverlimitEntriesWithoutLimit(c *C) {
	limit := 0
	cache := NewDNSCacheWithLimit(limit)
	for i := 0; i < 5; i++ {
		cache.Update(now, "test.com", []net.IP{net.ParseIP(fmt.Sprintf("1.1.1.%d", i))}, i)
	}
	c.Assert(cache.cleanupOverLimitEntries(), checker.DeepEquals, []string{})
	c.Assert(cache.Lookup("test.com"), HasLen, 4)
}

func (ds *DNSCacheTestSuite) TestGCOverlimitAfterTTLCleanup(c *C) {
	limit := 5
	cache := NewDNSCacheWithLimit(limit)

	// Make sure that the cleanup takes all the changes from 1 minute ago.
	cache.lastCleanup = time.Now().Add(-1 * time.Minute)
	for i := 1; i < limit+2; i++ {
		cache.Update(now, "test.com", []net.IP{net.ParseIP(fmt.Sprintf("1.1.1.%d", i))}, 1)
	}

	c.Assert(cache.Lookup("test.com"), HasLen, limit+1)
	c.Assert(cache.overLimit, HasLen, 1)

	result, _ := cache.cleanupExpiredEntries(time.Now().Add(5 * time.Second))
	c.Assert(result, checker.DeepEquals, []string{"test.com"})

	// Due all entries are deleted on TTL, the overlimit should return 0 entries.
	c.Assert(cache.cleanupOverLimitEntries(), checker.DeepEquals, []string{})
}

func (ds *DNSCacheTestSuite) TestOverlimitAfterDeleteForwardEntry(c *C) {
	// Validate if something delete the forward entry no invalid key access on
	// CG operation
	dnsCache := NewDNSCache()
	dnsCache.overLimit["test.com"] = true
	c.Assert(dnsCache.cleanupOverLimitEntries(), checker.DeepEquals, []string{})
}
