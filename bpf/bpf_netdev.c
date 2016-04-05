#define NODE_MAC { .addr = { 0xde, 0xad, 0xbe, 0xef, 0xc0, 0xde } }

#include <node_config.h>

#include <iproute2/bpf_api.h>

#include <sys/socket.h>

#include <stdint.h>
#include <stdio.h>

#include "lib/common.h"
#include "lib/ipv6.h"
#include "lib/ipv4.h"
#include "lib/icmp6.h"
#include "lib/eth.h"
#include "lib/dbg.h"
#include "lib/l3.h"
#include "lib/nat46.h"
#include "lib/arp.h"

static inline int is_node_subnet(const union v6addr *dst)
{
	union v6addr node = { . addr = ROUTER_IP };
	int tmp;

	tmp = dst->p1 - node.p1;
	if (!tmp) {
		tmp = dst->p2 - node.p2;
		if (!tmp) {
			tmp = dst->p3 - node.p3;
			if (!tmp) {
				__u32 a = ntohl(dst->p4);
				__u32 b = ntohl(node.p4);
				tmp = (a & 0xFFFF0000) - (b & 0xFFFF0000);
			}
		}
	}

	return !tmp;
}

/*
 * respond to arp request for target IPV4_GW with HOST_IFINDEX_MAC
 */
__section_tail(CILIUM_MAP_PROTO, CILIUM_MAP_PROTO_ARP) int arp_respond(struct __sk_buff *skb)
{
	union macaddr smac = {};
	__be32 sip = 0;
	__be16 arpop = __constant_htons(ARPOP_REPLY);
	__be32 responder_ip = IPV4_GW;
	union macaddr responder_mac = HOST_IFINDEX_MAC;

	load_eth_saddr(skb, smac.addr, 0);
	if (skb_load_bytes(skb, 28, &sip, sizeof(sip)) < 0)
		return TC_ACT_SHOT;

	skb_store_bytes(skb, 20, &arpop, sizeof(arpop), 0);
	skb_store_bytes(skb, 32, &smac, sizeof(smac), 0);
	skb_store_bytes(skb, 38, &sip, sizeof(sip), 0);

	skb_store_bytes(skb, 22, &responder_mac, 6, 0);
	skb_store_bytes(skb, 28, &responder_ip, 4, 0);

	store_eth_saddr(skb, responder_mac.addr, 0);
	store_eth_daddr(skb, smac.addr, 0);
	printk("arp_respond on ifindex %d\n", skb->ifindex);

	return redirect(skb->ifindex, 0);
}

__section("from-netdev")
int from_netdev(struct __sk_buff *skb)
{
#ifdef ENABLE_ARP_RESPONDER
	union macaddr responder_mac = HOST_IFINDEX_MAC;
	if (arp_check(skb, IPV4_GW, &responder_mac) == 1) {
		tail_call(skb, &cilium_proto, CILIUM_MAP_PROTO_ARP);
		return TC_ACT_SHOT;
	}
#endif

#ifdef ENABLE_NAT46
	/* First try to do v46 nat */
	if (skb->protocol == __constant_htons(ETH_P_IP)) {
		union v6addr sp = NAT46_SRC_PREFIX;
		union v6addr dp = HOST_IP;
		__u32 dst = 0;

		if (ipv4_load_daddr(skb, ETH_HLEN, &dst) < 0)
			return TC_ACT_SHOT;

		if ((dst & IPV4_MASK) != IPV4_RANGE)
			return TC_ACT_OK;

		if (ipv4_to_ipv6(skb, 14, &sp, &dp) < 0) {
			printk("ipv4_to_ipv6 failed\n");
			return TC_ACT_SHOT;
		}
		skb->tc_index = 1;
	}
#endif

	if (likely(skb->protocol == __constant_htons(ETH_P_IPV6))) {
		union v6addr dst = {};
		__u32 flowlabel = 0;
#ifdef HANDLE_NS
		__u8 nexthdr;

		nexthdr = load_byte(skb, ETH_HLEN + offsetof(struct ipv6hdr, nexthdr));
		if (unlikely(nexthdr == IPPROTO_ICMPV6)) {
			int ret = icmp6_handle(skb, ETH_HLEN);
			if (ret != REDIRECT_TO_LXC)
				return ret;
		}
#endif
		printk("IPv6 packet from netdev skb %p len %d\n", skb, skb->len);

		load_ipv6_daddr(skb, ETH_HLEN, &dst);
		ipv6_load_flowlabel(skb, ETH_HLEN, &flowlabel);

		if (is_node_subnet(&dst)) {
			printk("Targeted for a local container, src label: %d\n",
				ntohl(flowlabel));

			return do_l3(skb, ETH_HLEN, &dst, ntohl(flowlabel));
		}
	}

	return TC_ACT_OK;
}

BPF_LICENSE("GPL");
