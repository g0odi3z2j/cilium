#include <iproute2/bpf_api.h>
#include <sys/socket.h>
#include <stdint.h>
#include <string.h>
#include <linux/icmpv6.h>
#include "common.h"
#include "lib/ipv6.h"
#include "lib/eth.h"
#include <lxc_config.h>

__BPF_MAP(cilium_lxc, BPF_MAP_TYPE_HASH, 0, sizeof(__u16), sizeof(struct lxc_info), PIN_GLOBAL_NS, 1024);

#ifndef DISABLE_SMAC_VERIFICATION
static inline int verify_src_mac(struct __sk_buff *skb)
{
	union macaddr src, valid = LXC_MAC;
	load_eth_saddr(skb, &src, 0);
	return compare_eth_addr(&src, &valid);
}
#else
static inline int verify_src_mac(struct __sk_buff *skb)
{
	return 0;
}
#endif

#ifndef DISABLE_SIP_VERIFICATION
static inline int verify_src_ip(struct __sk_buff *skb, int off)
{
	union v6addr src, valid = LXC_IP;
	load_ipv6_saddr(skb, off, &src);
	return compare_ipv6_addr(&src, &valid);
}
#else
static inline int verify_src_ip(struct __sk_buff *skb, int off)
{
	return 0;
}
#endif

static inline int verify_dst_mac(struct __sk_buff *skb)
{
	union macaddr dst, valid = ROUTER_MAC;
	int ret;

	load_eth_daddr(skb, &dst, 0);
	ret = compare_eth_addr(&dst, &valid);

#ifdef DEBUG
	if (unlikely(ret)) {
		char fmt[] = "skb %p: invalid dst MAC\n";
		trace_printk(fmt, sizeof(fmt), skb);
	}
#endif

	return ret;
}

static inline void debug_trace_packet(struct __sk_buff *skb)
{
#ifdef DEBUG
	char fmt[] = "skb %p len %d\n";

	trace_printk(fmt, sizeof(fmt), skb, skb->len);
#endif
}

static inline int do_redirect6(struct __sk_buff *skb, int nh_off)
{
	__u16 lxc_id;
	union v6addr dst;
	int node_id, *ifindex;

	debug_trace_packet(skb);

	if (verify_src_mac(skb) || verify_src_ip(skb, nh_off) ||
	    verify_dst_mac(skb))
		return -1;

	load_ipv6_daddr(skb, nh_off, &dst);
	lxc_id = derive_lxc_id(&dst);
	node_id = derive_node_id(&dst);
#ifdef DEBUG
	if (1) {
		char fmt[] = "lxc-id: %x node-id: %x\n";
		trace_printk(fmt, sizeof(fmt), lxc_id, node_id);
	}
#endif

	if (node_id != NODE_ID) {
		/* FIXME: Handle encapsulation case */
	} else {
		struct lxc_info *dst_lxc;

		dst_lxc = map_lookup_elem(&cilium_lxc, &lxc_id);
		if (dst_lxc) {
			__u64 tmp_mac = dst_lxc->mac;
			store_eth_daddr(skb, (char *) &tmp_mac, 0);

			if (decrement_ipv6_hoplimit(skb, nh_off)) {
				/* FIXME: Handle hoplimit == 0 */
			}

			char fmt[] = "Found destination container locally\n";
			trace_printk(fmt, sizeof(fmt));

			return redirect(dst_lxc->ifindex, 0);
		}
	}

	return -1;
}

static inline int handle_icmp6(struct __sk_buff *skb, int nh_off)
{
	char fmt[] = "ICMPv6 packet skb %p len %d type %d\n";
	char fmt1[] = "Redirect skb to Ifindex %d\n";
	union v6addr sip, router_ip;
	__u8 type;
	struct icmp6hdr icmp6hdr;
	union macaddr smac;
	union macaddr router_mac = ROUTER_MAC;
	__u8 opts[2] = { 2, 1 };

	type = load_byte(skb, ETH_HLEN + sizeof(struct ipv6hdr) + offsetof(struct icmp6hdr, icmp6_type));

	trace_printk(fmt, sizeof(fmt), skb, skb->len, type);

	if (type == 135) {
		/* skb->daddr = skb->saddr */
		load_ipv6_saddr(skb, nh_off, &sip);
		skb_store_bytes(skb, ETH_HLEN + offsetof(struct ipv6hdr, daddr), &sip, 16, 0);
		/* skb->saddr = router address, verifier rejects when initialized statically */
		router_ip.p1 = htonl(0xab);
		router_ip.p2 = htonl(0xcd);
		router_ip.p3 = htonl(0xab);
		router_ip.p4 = htonl(0xef);
		skb_store_bytes(skb, ETH_HLEN + offsetof(struct ipv6hdr, saddr), &router_ip, 16, 0);

		/* fill icmp6hdr */
		icmp6hdr.icmp6_type = 136;
		icmp6hdr.icmp6_code = 0;
		icmp6hdr.icmp6_dataun.un_data32[0] = 0;
		icmp6hdr.icmp6_router = 1;
		icmp6hdr.icmp6_solicited = 1;
		icmp6hdr.icmp6_override = 0;
		/* FIXME compute icmp6 checksum */
		icmp6hdr.icmp6_cksum = 0;
		skb_store_bytes(skb, ETH_HLEN + sizeof(struct ipv6hdr), &icmp6hdr, sizeof(icmp6hdr), 0);
		// ND_OPT_TARGET_LL_ADDR
		skb_store_bytes(skb, ETH_HLEN + sizeof(struct ipv6hdr) + sizeof(struct icmp6hdr) + sizeof(struct in6_addr), opts, sizeof(opts), 0);
		skb_store_bytes(skb, ETH_HLEN + sizeof(struct ipv6hdr) + sizeof(struct icmp6hdr) + sizeof(struct in6_addr) + 2, &router_mac, sizeof(router_mac), 0);

		/* dmac = smac, smac = router mac */
		load_eth_saddr(skb, &smac, 0);
		store_eth_daddr(skb, (char *) smac.addr, 0);
		store_eth_saddr(skb, (char *) &router_mac, 0);
		trace_printk(fmt1, sizeof(fmt1), skb->ifindex);
		return redirect(skb->ifindex, 0);

	}

	return -1;
}
__section("from-container")
int handle_ingress(struct __sk_buff *skb)
{
	int ret = 0, nh_off = ETH_HLEN;
	__u8 nexthdr;

	if (likely(skb->protocol == __constant_htons(ETH_P_IPV6))) {
		nexthdr = load_byte(skb, nh_off + offsetof(struct ipv6hdr, nexthdr));
		if (nexthdr == IPPROTO_ICMPV6)
			ret = handle_icmp6(skb, nh_off);
		else
			ret = do_redirect6(skb, nh_off);
	}
	return ret;
}
BPF_LICENSE("GPL");
