#ifndef __LIB_L3_H_
#define __LIB_L3_H_

#include "common.h"
#include "ipv6.h"
#include "eth.h"
#include "dbg.h"

__BPF_MAP(cilium_lxc, BPF_MAP_TYPE_HASH, 0, sizeof(__u16), sizeof(struct lxc_info), PIN_GLOBAL_NS, 1024);

#ifdef ENCAP_IFINDEX
static inline int do_encapsulation(struct __sk_buff *skb, __u32 node_id)
{
	struct bpf_tunnel_key key = {};

	key.tunnel_id = 42;
	key.remote_ipv4 = node_id;
	key.tunnel_af = AF_INET;

	printk("Performing encapsulation to node %x on ifindex %u\n", node_id, ENCAP_IFINDEX);

	skb_set_tunnel_key(skb, &key, sizeof(key), 0);

	return redirect(ENCAP_IFINDEX, 0);
}
#endif

static inline void __inline__ __do_l3(struct __sk_buff *skb, int nh_off,
				     __u8 *smac, __u8 *dmac)
{
	if (smac)
		store_eth_saddr(skb, smac, 0);

	store_eth_daddr(skb, dmac, 0);

	if (decrement_ipv6_hoplimit(skb, nh_off)) {
		/* FIXME: Handle hoplimit == 0 */
	}
}

static inline int __inline__ do_l3(struct __sk_buff *skb, int nh_off,
				   union v6addr *dst)
{
	struct lxc_info *dst_lxc;
	__u16 lxc_id = derive_lxc_id(dst);

	printk("L3 on local node, lxc-id: %x\n", lxc_id);

	dst_lxc = map_lookup_elem(&cilium_lxc, &lxc_id);
	if (dst_lxc) {
		__u64 tmp_mac = dst_lxc->mac;
		union macaddr router_mac = NODE_MAC;

		printk("Found destination container locally\n");

		__do_l3(skb, nh_off, (__u8 *) &router_mac.addr,
			(__u8 *) &tmp_mac);

		return redirect(dst_lxc->ifindex, 0);
	} else {
		printk("Unknown container %x\n", lxc_id);
	}

	return TC_ACT_UNSPEC;
}

#endif
