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

package policy

import (
	"net"

	"github.com/cilium/cilium/pkg/policy/api"
)

// getPrefixesFromCIDR fetches all CIDRs referred to by the specified slice
// and returns them as regular golang CIDR objects.
func getPrefixesFromCIDR(cidrs api.CIDRSlice) []*net.IPNet {
	res := make([]*net.IPNet, 0, len(cidrs))
	for _, cidr := range cidrs {
		_, prefix, err := net.ParseCIDR(string(cidr))
		if err != nil {
			// Likely the CIDR is specified in host format.
			ip := net.ParseIP(string(cidr))
			if ip != nil {
				bits := net.IPv6len * 8
				if ip.To4() != nil {
					ip = ip.To4()
					bits = net.IPv4len * 8
				}
				prefix = &net.IPNet{
					IP:   ip,
					Mask: net.CIDRMask(bits, bits),
				}
			}
		}
		if prefix != nil {
			res = append(res, prefix)
		}
	}
	return res
}

// getPrefixesFromCIDRSet fetches all CIDRs referred to by the specified slice
// and returns them as regular golang CIDR objects.
func getPrefixesFromCIDRSet(rules api.CIDRRuleSlice) []*net.IPNet {
	cidrs := api.ComputeResultantCIDRSet(rules)
	return getPrefixesFromCIDR(cidrs)
}

// GetCIDRPrefixes runs through the specified 'rules' to find every reference
// to a CIDR in the rules, and returns a slice containing all of these CIDRs.
// Multiple rules referring to the same CIDR will result in multiple copies of
// the CIDR in the returned slice.
//
// Assumes that validation already occurred on 'rules'.
func GetCIDRPrefixes(rules api.Rules) []*net.IPNet {
	if len(rules) == 0 {
		return nil
	}
	res := make([]*net.IPNet, 0, 32)
	for _, r := range rules {
		for _, ir := range r.Ingress {
			if len(ir.FromCIDR) > 0 {
				res = append(res, getPrefixesFromCIDR(ir.FromCIDR)...)
			}
			if len(ir.FromCIDRSet) > 0 {
				res = append(res, getPrefixesFromCIDRSet(ir.FromCIDRSet)...)
			}
		}
		for _, er := range r.Egress {
			if len(er.ToCIDR) > 0 {
				res = append(res, getPrefixesFromCIDR(er.ToCIDR)...)
			}
			if len(er.ToCIDRSet) > 0 {
				res = append(res, getPrefixesFromCIDRSet(er.ToCIDRSet)...)
			}
		}
	}
	return res
}
