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

package policy

import (
	"fmt"

	"github.com/cilium/cilium/pkg/policy/api"
)

type rule struct {
	api.Rule

	fromEntities []api.EndpointSelector
	toEntities   []api.EndpointSelector
}

func (r *rule) String() string {
	return fmt.Sprintf("%v", r.EndpointSelector)
}

// validate has a side effect of populating the fromEntities and toEntities
// slices to avoid superfluent map accesses
func (r *rule) validate() error {
	if r == nil || r.EndpointSelector.LabelSelector == nil {
		return fmt.Errorf("nil rule")
	}

	if len(r.EndpointSelector.MatchLabels) == 0 &&
		len(r.EndpointSelector.MatchExpressions) == 0 {
		return fmt.Errorf("empty EndpointSelector")
	}

	// resetting entity selector slices
	r.fromEntities = []api.EndpointSelector{}
	r.toEntities = []api.EndpointSelector{}
	entities := []api.Entity{}

	ingressEntityCounter := 0
	for _, rule := range r.Ingress {
		entities = append(entities, rule.FromEntities...)
		ingressEntityCounter += len(rule.FromEntities)
	}

	for _, rule := range r.Egress {
		entities = append(entities, rule.ToEntities...)
	}

	for j, entity := range entities {
		selector, ok := api.EntitySelectorMapping[entity]
		if !ok {
			return fmt.Errorf("unsupported entity: %s", entity)
		}

		if j < ingressEntityCounter {
			r.fromEntities = append(r.fromEntities, selector)
		} else {
			r.toEntities = append(r.toEntities, selector)
		}
	}

	return nil
}

func mergeL4Port(ctx *SearchContext, r api.PortRule, p api.PortProtocol, dir string, proto string, resMap L4PolicyMap) (int, error) {
	key := p.Port + "/" + proto
	v, ok := resMap[key]
	if !ok {
		resMap[key] = CreateL4Filter(r, p, dir, proto)
		return 1, nil
	}
	l4Filter := CreateL4Filter(r, p, dir, proto)
	if l4Filter.L7Parser != "" && v.L7Parser == "" {
		v.L7Parser = l4Filter.L7Parser
	} else if l4Filter.L7Parser != v.L7Parser {
		ctx.PolicyTrace("   Merge conflict: mismatching parsers %s/%s\n", l4Filter.L7Parser, v.L7Parser)
		return 0, fmt.Errorf("Cannot merge conflicting L7 parsers (%s/%s)", l4Filter.L7Parser, v.L7Parser)
	}
	if l4Filter.L7RedirectPort != 0 && v.L7RedirectPort == 0 {
		v.L7RedirectPort = l4Filter.L7RedirectPort
	} else if l4Filter.L7RedirectPort != v.L7RedirectPort {
		ctx.PolicyTrace("   Merge conflict: mismatching redirect ports %d/%d\n", l4Filter.L7RedirectPort, v.L7RedirectPort)
		return 0, fmt.Errorf("Cannot merge conflicting redirect ports (%d/%d)", l4Filter.L7RedirectPort, v.L7RedirectPort)
	}
	switch {
	case len(l4Filter.L7Rules.HTTP) > 0:
		if len(v.L7Rules.Kafka) > 0 {
			ctx.PolicyTrace("   Merge conflict: mismatching L7 rule types.\n")
			return 0, fmt.Errorf("Cannot merge conflicting L7 rule types")
		}
		v.L7Rules.HTTP = append(v.L7Rules.HTTP, l4Filter.L7Rules.HTTP...)
	case len(l4Filter.L7Rules.Kafka) > 0:
		if len(v.L7Rules.HTTP) > 0 {
			ctx.PolicyTrace("   Merge conflict: mismatching L7 rule types.\n")
			return 0, fmt.Errorf("Cannot merge conflicting L7 rule types")
		}
		v.L7Rules.Kafka = append(v.L7Rules.Kafka, l4Filter.L7Rules.Kafka...)
	default:
		ctx.PolicyTrace("   No L7 rules to merge.\n")
	}
	resMap[key] = v
	return 1, nil
}

func mergeL4(ctx *SearchContext, dir string, portRules []api.PortRule, resMap L4PolicyMap) (int, error) {
	found := 0
	var err error

	for _, r := range portRules {
		ctx.PolicyTrace("  Allows %s port %v\n", dir, r.Ports)

		if r.RedirectPort != 0 {
			ctx.PolicyTrace("    Redirect-To: %d\n", r.RedirectPort)
		}

		if r.Rules != nil {
			for _, l7 := range r.Rules.HTTP {
				ctx.PolicyTrace("      %+v\n", l7)
			}
		}

		for _, p := range r.Ports {
			var cnt int
			if p.Protocol != "" {
				cnt, err = mergeL4Port(ctx, r, p, dir, p.Protocol, resMap)
				if err != nil {
					return found, err
				}
				found += cnt
			} else {
				cnt, err = mergeL4Port(ctx, r, p, dir, "tcp", resMap)
				if err != nil {
					return found, err
				}
				found += cnt

				cnt, err = mergeL4Port(ctx, r, p, dir, "udp", resMap)
				if err != nil {
					return found, err
				}
				found += cnt
			}
		}
	}

	return found, nil
}

func (r *rule) resolveL4Policy(ctx *SearchContext, state *traceState, result *L4Policy) (*L4Policy, error) {
	if !r.EndpointSelector.Matches(ctx.To) {
		ctx.PolicyTraceVerbose("  Rule %d %s: no match\n", state.ruleID, r)
		return nil, nil
	}

	state.selectedRules++
	ctx.PolicyTrace("* Rule %d %s: match\n", state.ruleID, r)
	found := 0

	if !ctx.EgressL4Only {
		for _, r := range r.Ingress {
			cnt, err := mergeL4(ctx, "Ingress", r.ToPorts, result.Ingress)
			if err != nil {
				return nil, err
			}
			found += cnt
		}
	}

	if !ctx.IngressL4Only {
		for _, r := range r.Egress {
			cnt, err := mergeL4(ctx, "Egress", r.ToPorts, result.Egress)
			if err != nil {
				return nil, err
			}
			found += cnt
		}
	}

	if found > 0 {
		return result, nil
	}

	ctx.PolicyTrace("    No L4 rules\n")
	return nil, nil
}

func mergeL3(ctx *SearchContext, dir string, ipRules []api.CIDR, resMap *L3PolicyMap) int {
	found := 0

	for _, r := range ipRules {
		strCIDR := string(r)
		ctx.PolicyTrace("  Allows %s IP %s\n", dir, strCIDR)

		found += resMap.Insert(strCIDR)
	}

	return found
}

func (r *rule) resolveL3Policy(ctx *SearchContext, state *traceState, result *L3Policy) *L3Policy {
	if !r.EndpointSelector.Matches(ctx.To) {
		ctx.PolicyTraceVerbose("  Rule %d %s: no match\n", state.ruleID, r)
		return nil
	}

	state.selectedRules++
	ctx.PolicyTrace("* Rule %d %s: match\n", state.ruleID, r)
	found := 0

	for _, r := range r.Ingress {
		found += mergeL3(ctx, "Ingress", r.FromCIDR, &result.Ingress)
	}
	for _, r := range r.Egress {
		found += mergeL3(ctx, "Egress", r.ToCIDR, &result.Egress)
	}

	if found > 0 {
		return result
	}

	ctx.PolicyTrace("    No L3 rules\n")
	return nil
}

func (r *rule) canReach(ctx *SearchContext, state *traceState) api.Decision {
	entitiesDecision := r.canReachEntities(ctx, state)

	if !r.EndpointSelector.Matches(ctx.To) {
		if entitiesDecision == api.Undecided {
			ctx.PolicyTraceVerbose("  Rule %d %s: no match for %+v\n", state.ruleID, r, ctx.To)
		} else {
			state.selectedRules++
			ctx.PolicyTrace("* Rule %d %s: match\n", state.ruleID, r)
		}
		return entitiesDecision
	}

	state.selectedRules++
	ctx.PolicyTrace("* Rule %d %s: match\n", state.ruleID, r)

	for _, r := range r.Ingress {
		for _, sel := range r.FromRequires {
			ctx.PolicyTrace("    Requires from labels %+v", sel)
			if !sel.Matches(ctx.From) {
				ctx.PolicyTrace("-     Labels %v not found\n", ctx.From)
				return api.Denied
			}
			ctx.PolicyTrace("+     Found all required labels\n")
		}
	}

	// separate loop is needed as failure to meet FromRequires always takes
	// precedence over FromEndpoints
	for _, r := range r.Ingress {
		for _, sel := range r.FromEndpoints {
			ctx.PolicyTrace("    Allows from labels %+v", sel)
			if sel.Matches(ctx.From) {
				ctx.PolicyTrace("+     Found all required labels\n")
				return api.Allowed
			}

			ctx.PolicyTrace("      Labels %v not found\n", ctx.From)
		}
	}

	for _, entitySelector := range r.fromEntities {
		if entitySelector.Matches(ctx.From) {
			ctx.PolicyTrace("+     Found all required labels to match entity %s\n", entitySelector.String())
			return api.Allowed
		}

	}

	return entitiesDecision
}

func (r *rule) canReachEntities(ctx *SearchContext, state *traceState) api.Decision {
	for _, entitySelector := range r.toEntities {
		if entitySelector.Matches(ctx.To) {
			ctx.PolicyTrace("+     Found all required labels to match entity %s\n", entitySelector.String())
			return api.Allowed
		}
	}

	return api.Undecided
}
