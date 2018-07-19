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

package counter

import (
	"fmt"
	"net"

	"github.com/cilium/cilium/pkg/lock"
)

// PrefixLengthCounter tracks references to prefix lengths, limited by the
// maxUniquePrefixes count. Neither of the IPv4 or IPv6 counters nested within
// may contain more keys than the specified maximum number of unique prefixes.
type PrefixLengthCounter struct {
	lock.RWMutex

	v4 IntCounter
	v6 IntCounter

	maxUniquePrefixes int
}

// NewPrefixLengthCounter returns a new PrefixLengthCounter which limits
// insertions to the specified maximum number of unique prefix lengths.
func NewPrefixLengthCounter(maxUniquePrefixes int) *PrefixLengthCounter {
	return &PrefixLengthCounter{
		v4:                make(IntCounter),
		v6:                make(IntCounter),
		maxUniquePrefixes: maxUniquePrefixes,
	}
}

// checkLimits checks whether the specified new count of prefixes would exceed
// the specified limit on the maximum number of unique keys, and returns an
// error if it would exceed the limit.
func (p *PrefixLengthCounter) checkLimits(current, newCount int) error {
	if newCount > p.maxUniquePrefixes {
		return fmt.Errorf("Adding specified prefixes would result in too many prefix lengths (current: %d, result: %d, max: %d)",
			current, newCount, p.maxUniquePrefixes)
	}
	return nil
}

// Add increments references to prefix lengths for the specified IPNets to the
// counter. If the maximum number of unique prefix lengths would be exceeded,
// returns an error.
//
// Returns true if adding these prefixes results in an increase in the total
// number of unique prefix lengths in the counter.
func (p *PrefixLengthCounter) Add(prefixes []*net.IPNet) (bool, error) {
	p.Lock()
	defer p.Unlock()

	// Assemble a map of references that need to be added
	newV4Counter := p.v4.DeepCopy()
	newV6Counter := p.v6.DeepCopy()
	newV4Prefixes := false
	newV6Prefixes := false
	for _, prefix := range prefixes {
		ones, bits := prefix.Mask.Size()

		switch bits {
		case net.IPv4len * 8:
			if newV4Counter.Add(ones) {
				newV4Prefixes = true
			}
		case net.IPv6len * 8:
			if newV6Counter.Add(ones) {
				newV6Prefixes = true
			}
		default:
			return false, fmt.Errorf("Unsupported IPAddr bitlength %d", bits)
		}
	}

	// Check if they can be added given the limit in place
	if newV4Prefixes {
		if err := p.checkLimits(len(p.v4), len(newV4Counter)); err != nil {
			return false, err
		}
	}
	if newV6Prefixes {
		if err := p.checkLimits(len(p.v6), len(newV6Counter)); err != nil {
			return false, err
		}
	}

	// Set and return whether anything changed
	p.v4 = newV4Counter
	p.v6 = newV6Counter
	return newV4Prefixes || newV6Prefixes, nil
}

// Delete reduces references to prefix lengths in the the specified IPNets from
// the counter. Returns true if removing references to these prefix lengths
// would result in a decrese in the total number of unique prefix lengths in
// the counter.
func (p *PrefixLengthCounter) Delete(prefixes []*net.IPNet) (changed bool) {
	p.Lock()
	defer p.Unlock()

	for _, prefix := range prefixes {
		ones, bits := prefix.Mask.Size()
		switch bits {
		case net.IPv4len * 8:
			if p.v4.Delete(ones) {
				changed = true
			}
		case net.IPv6len * 8:
			if p.v6.Delete(ones) {
				changed = true
			}
		}
	}

	return changed
}

// ToBPFData converts the counter into a set of prefix lengths that the BPF
// datapath can use for LPM lookup.
func (p *PrefixLengthCounter) ToBPFData() (s6, s4 []int) {
	p.RLock()
	defer p.RUnlock()

	return p.v6.ToBPFData(), p.v4.ToBPFData()
}
