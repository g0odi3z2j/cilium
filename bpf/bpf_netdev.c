#include <node_config.h>

#include <iproute2/bpf_api.h>

#include <sys/socket.h>

#include <stdint.h>
#include <stdio.h>

#include "lib/common.h"
#include "lib/ipv6.h"
#include "lib/icmp6.h"
#include "lib/eth.h"
#include "lib/dbg.h"
#include "lib/l3.h"

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

__section("from-netdev")
int from_netdev(struct __sk_buff *skb)
{
	int ret;

	if (likely(skb->protocol == __constant_htons(ETH_P_IPV6))) {
		union v6addr dst = {};
		__u32 flowlabel;
		__u8 nexthdr;

		nexthdr = load_byte(skb, ETH_HLEN + offsetof(struct ipv6hdr, nexthdr));
		if (unlikely(nexthdr == IPPROTO_ICMPV6)) {
			ret = icmp6_handle(skb, ETH_HLEN);
			if (ret != LXC_REDIRECT)
				return ret;
		}

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
