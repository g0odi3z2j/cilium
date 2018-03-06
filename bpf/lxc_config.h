/*
 *  Copyright (C) 2016-2017 Authors of Cilium
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program; if not, write to the Free Software
 *  Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 */
/*
 * This is just a dummy header with dummy values to allow for test
 * compilation without the full code generation engine backend.
 */

#define LXC_MAC { .addr = { 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff } }
#define LXC_IP 0xbe, 0xef, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x1, 0x1, 0x65, 0x82, 0xbc
#define LXC_IPV4 0x10203040
#define LXC_ID 0x1010
#define LXC_ID_NB 0x1010
#define LXC_NAT46
#ifndef SECLABEL
#define SECLABEL 0xfffff
#define SECLABEL_NB 0xfffff
#endif
#define POLICY_MAP cilium_policy_foo
#define NODE_MAC { .addr = { 0xde, 0xad, 0xbe, 0xef, 0xc0, 0xde } }
#define GENEVE_OPTS { 0xff, 0xff, 0x1, 0x1, 0x0, 0x0, 0x1, 0x1e }
#define DROP_NOTIFY
#define TRACE_NOTIFY
#define CT_MAP6 cilium_ct6_111
#define CT_MAP4 cilium_ct4_111
#define CT_MAP_SIZE 4096
#define CALLS_MAP cilium_calls_111
#define LB_L3
#define LB_L4
#define CONNTRACK
#define NR_CFG_L4_INGRESS 2
#define CFG_L4_INGRESS 0, 80, 8080, 0, 1, 80, 8080, 0, (), 0
#define NR_CFG_L4_EGRESS 1
#define CFG_L4_EGRESS 0, 80, 8080, 0, (), 0
#define POLICY_INGRESS
#define POLICY_EGRESS
#define ENABLE_IPv4
#define ALLOW_TO_HOST
#define HAVE_L4_POLICY

#ifndef SKIP_CIDR_PREFIXES
#define CIDR6_INGRESS_MAP cilium_cidr6_ingress_foo
#define CIDR4_INGRESS_MAP cilium_cidr4_ingress_foo
#define CIDR6_EGRESS_MAP cilium_cidr6_egress_foo
#define CIDR4_EGRESS_MAP cilium_cidr4_egress_foo
/* It appears that we can support around ninety prefixes in an unrolled loop
 * for LPM CIDR handling in older kernels along with the rest of the logic in
 * the datapath, hence the defines below. This number was arrived to by
 * adjusting the number of prefixes and running:
 *    $ make -C bpf && sudo test/bpf/verifier-test.sh
 *
 *  If you're from a future where all supported kernels include LPM map type,
 *  consider deprecating the hash-based CIDR lookup and removing the below.
 */
#define CIDR4_EGRESS_PREFIXES 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, \
20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
#define CIDR6_EGRESS_PREFIXES 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, \
80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62,   \
61, 60, 59, 58, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43,   \
42, 41, 40, 39, 38, 37, 36, 35, 34, 33, CIDR4_EGRESS_PREFIXES
#define CIDR6_INGRESS_PREFIXES CIDR6_EGRESS_PREFIXES
#define CIDR4_INGRESS_PREFIXES CIDR4_EGRESS_PREFIXES
#endif
