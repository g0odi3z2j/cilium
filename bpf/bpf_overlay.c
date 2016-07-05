#include <node_config.h>
#include <netdev_config.h>

#include <bpf/api.h>

#include <stdint.h>
#include <stdio.h>

#include "lib/common.h"
#include "lib/maps.h"
#include "lib/ipv6.h"
#include "lib/eth.h"
#include "lib/dbg.h"
#include "lib/l3.h"
#include "lib/geneve.h"
#include "lib/drop.h"

static inline int __inline__ do_l3_from_overlay(struct __sk_buff *skb, struct ipv6hdr *ip6,
						int nh_off, __u32 tunnel_id)
{
	union v6addr *dst = (union v6addr *) &ip6->daddr;
	__u32 node_id;

	cilium_trace_capture(skb, DBG_CAPTURE_FROM_OVERLAY, skb->ingress_ifindex);

	node_id = ipv6_derive_node_id(dst);

	if (unlikely(node_id != NODE_ID))
		return DROP_NON_LOCAL;
	else
		return local_delivery(skb, nh_off, ip6, dst, ntohl(tunnel_id));
}

__section("from-overlay")
int from_overlay(struct __sk_buff *skb)
{
	void *data_end = (void *) (long) skb->data_end;
	void *data = (void *) (long) skb->data;
	struct ipv6hdr *ip6 = data + ETH_HLEN;
	struct bpf_tunnel_key key = {};
	int ret = TC_ACT_OK;

	if (data + sizeof(*ip6) + ETH_HLEN > data_end) {
		ret = DROP_INVALID;
		goto error;
	}

	if (unlikely(skb_get_tunnel_key(skb, &key, sizeof(key), 0) < 0)) {
		ret = DROP_NO_TUNNEL_KEY;
		goto error;
	}

#ifdef ENCAP_GENEVE
	if (1) {
		uint8_t buf[MAX_GENEVE_OPT_LEN] = {};
		struct geneveopt_val geneveopt_val = {};

		if (unlikely(skb_get_tunnel_opt(skb, buf, sizeof(buf)) < 0)) {
			ret = DROP_NO_TUNNEL_OPT;
			goto error;
		}

		ret = parse_geneve_options(&geneveopt_val, buf);
		if (IS_ERR(ret))
			goto error;
	}
#endif

	if (likely(skb->protocol == __constant_htons(ETH_P_IPV6)))
		ret = do_l3_from_overlay(skb, ip6, ETH_HLEN, key.tunnel_id);
	else
		ret = DROP_UNKNOWN_L3;

	if (IS_ERR(ret)) {
error:
		return send_drop_notify_error(skb, ret, TC_ACT_SHOT);
	} else
		return ret;
}
BPF_LICENSE("GPL");
