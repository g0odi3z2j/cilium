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

package option

import (
	"errors"
)

const (
	PolicyTracing       = "PolicyTracing"
	ConntrackAccounting = "ConntrackAccounting"
	ConntrackLocal      = "ConntrackLocal"
	Conntrack           = "Conntrack"
	Debug               = "Debug"
	DebugLB             = "DebugLB"
	DropNotify          = "DropNotification"
	TraceNotify         = "TraceNotification"
	NAT46               = "NAT46"
	IngressPolicy       = "IngressPolicy"
	EgressPolicy        = "EgressPolicy"
	AlwaysEnforce       = "always"
	NeverEnforce        = "never"
	DefaultEnforcement  = "default"
)

var (
	ErrNAT46ReqIPv4 = errors.New("NAT46 requires IPv4 to be enabled")
	IPv4Disabled    = false
)

var (
	SpecConntrackAccounting = Option{
		Define:      "CONNTRACK_ACCOUNTING",
		Description: "Enable per flow (conntrack) statistics",
		Requires:    []string{Conntrack},
	}

	SpecConntrackLocal = Option{
		Define:      "CONNTRACK_LOCAL",
		Description: "Use endpoint dedicated tracking table instead of global one",
		Requires:    []string{Conntrack},
	}

	SpecConntrack = Option{
		Define:      "CONNTRACK",
		Description: "Enable stateful connection tracking",
	}

	SpecDebug = Option{
		Define:      "DEBUG",
		Description: "Enable debugging trace statements",
	}

	SpecDebugLB = Option{
		Define:      "LB_DEBUG",
		Description: "Enable debugging trace statements for load balancer",
	}

	SpecDropNotify = Option{
		Define:      "DROP_NOTIFY",
		Description: "Enable drop notifications",
	}

	SpecTraceNotify = Option{
		Define:      "TRACE_NOTIFY",
		Description: "Enable trace notifications",
	}

	SpecNAT46 = Option{
		Define:      "ENABLE_NAT46",
		Description: "Enable automatic NAT46 translation",
		Requires:    []string{Conntrack},
		Verify: func(key string, val bool) error {
			if IPv4Disabled {
				return ErrNAT46ReqIPv4
			}
			return nil
		},
	}

	IngressSpecPolicy = Option{
		Define:      "POLICY_INGRESS",
		Description: "Enable ingress policy enforcement",
	}

	EgressSpecPolicy = Option{
		Define:      "POLICY_EGRESS",
		Description: "Enable egress policy enforcement",
	}
)
