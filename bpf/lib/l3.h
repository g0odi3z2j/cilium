#ifndef __LIB_L3_H_
#define __LIB_L3_H_

#include "common.h"
#include "ipv6.h"
#include "eth.h"
#include "dbg.h"
#include "l4.h"
#include "lxc_map.h"
#include "icmp6.h"
#include "nat46.h"

#ifdef ENCAP_IFINDEX
static inline int do_encapsulation(struct __sk_buff *skb, __u32 node_id,
				   __u32 seclabel)
{
	struct bpf_tunnel_key key = {};

	key.tunnel_id = 42;
	key.remote_ipv4 = node_id;
	key.tunnel_label = seclabel;

	printk("Performing encapsulation to node %x on ifindex %u seclabel 0x%x\n",
		node_id, ENCAP_IFINDEX, seclabel);

	skb_set_tunnel_key(skb, &key, sizeof(key), 0);

	return redirect(ENCAP_IFINDEX, 0);
}
#endif

static inline int __inline__ __do_l3(struct __sk_buff *skb, int nh_off,
				     __u8 *smac, __u8 *dmac)
{

	if (decrement_ipv6_hoplimit(skb, nh_off)) {
		return send_icmp6_time_exceeded(skb, nh_off);
	}

	if (smac)
		store_eth_saddr(skb, smac, 0);

	store_eth_daddr(skb, dmac, 0);

	return TC_ACT_OK;
}

#ifndef DISABLE_PORT_MAP
static inline void map_lxc_in(struct __sk_buff *skb, int off,
			      struct lxc_info *lxc)
{
	__u8 nexthdr = 0;
	int i;

	if (ipv6_load_nexthdr(skb, off, &nexthdr) < 0)
		return;

	off += sizeof(struct ipv6hdr);

#pragma unroll
	for (i = 0; i < PORTMAP_MAX; i++) {
		if (!lxc->portmap[i].to || !lxc->portmap[i].from)
			break;

		do_port_map_in(skb, off, nexthdr, &lxc->portmap[i]);
	}
}
#endif /* DISABLE_PORT_MAP */

static inline int __inline__ do_l3(struct __sk_buff *skb, int nh_off,
				   union v6addr *dst, __u32 seclabel)
{
	struct lxc_info *dst_lxc;
	__u16 lxc_id = derive_lxc_id(dst);
	int ret = 0;

	printk("L3 on local node, lxc-id: %x\n", lxc_id);

	dst_lxc = map_lookup_elem(&cilium_lxc, &lxc_id);
	if (dst_lxc) {
		union macaddr router_mac = NODE_MAC;
		mac_t tmp_mac = dst_lxc->mac;

		ret = __do_l3(skb, nh_off, (__u8 *) &router_mac.addr, (__u8 *) &tmp_mac);
		if (ret == TC_ACT_REDIRECT || ret == -1)
			return ret;

#ifndef DISABLE_PORT_MAP
		if (dst_lxc->portmap[0].to)
			map_lxc_in(skb, nh_off, dst_lxc);
#else
		printk("Port mapping disabled, skipping.\n");
#endif /* DISABLE_PORT_MAP */

		printk("Redirecting to ifindex %u\n", dst_lxc->ifindex);

#ifdef ENABLE_NAT46
		if (skb->tc_index == 0 &&
		    skb->protocol == __constant_htons(ETH_P_IPV6)) {
			ret = ipv6_to_ipv4(skb, nh_off);
			if (ret == -1)
				printk("v64 nat failed\n");

		}
		skb->tc_index = 0;
#endif

		skb->cb[0] = seclabel;
		skb->cb[1] = dst_lxc->ifindex;

		printk("ID: %d\n", ntohl(dst_lxc->sec_label));
		tail_call(skb, &cilium_jmp, ntohl(dst_lxc->sec_label));
		printk("Fall through\n");
		return BPF_H_DEFAULT;
	} else {
		printk("Unknown container %x\n", lxc_id);
	}

	return TC_ACT_UNSPEC;
}

#endif
