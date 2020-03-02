/* SPDX-License-Identifier: GPL-2.0 */
/* Copyright (C) 2016-2020 Authors of Cilium */

#ifndef __LIB_L3_H_
#define __LIB_L3_H_

#include "common.h"
#include "ipv6.h"
#include "ipv4.h"
#include "eps.h"
#include "eth.h"
#include "dbg.h"
#include "l4.h"
#include "icmp6.h"
#include "csum.h"

#ifdef ENABLE_IPV6
static __always_inline int ipv6_l3(struct __ctx_buff *ctx, int l3_off,
				   __u8 *smac, __u8 *dmac, __u8 direction)
{
	int ret;

	ret = ipv6_dec_hoplimit(ctx, l3_off);
	if (IS_ERR(ret))
		return ret;

	if (ret > 0) {
		/* Hoplimit was reached */
		return icmp6_send_time_exceeded(ctx, l3_off, direction);
	}

	if (smac && eth_store_saddr(ctx, smac, 0) < 0)
		return DROP_WRITE_ERROR;

	if (eth_store_daddr(ctx, dmac, 0) < 0)
		return DROP_WRITE_ERROR;

	return CTX_ACT_OK;
}
#endif /* ENABLE_IPV6 */

static __always_inline int ipv4_l3(struct __ctx_buff *ctx, int l3_off,
				   __u8 *smac, __u8 *dmac, struct iphdr *ip4)
{
	if (ipv4_dec_ttl(ctx, l3_off, ip4)) {
		/* FIXME: Send ICMP TTL */
		return DROP_INVALID;
	}

	if (smac && eth_store_saddr(ctx, smac, 0) < 0)
		return DROP_WRITE_ERROR;

	if (eth_store_daddr(ctx, dmac, 0) < 0)
		return DROP_WRITE_ERROR;

	return CTX_ACT_OK;
}

#ifdef ENABLE_IPV6
static __always_inline int ipv6_local_delivery(struct __ctx_buff *ctx, int l3_off,
					       __u32 seclabel,
					       struct endpoint_info *ep,
					       __u8 direction)
{
	int ret;

	cilium_dbg(ctx, DBG_LOCAL_DELIVERY, ep->lxc_id, seclabel);

	mac_t lxc_mac = ep->mac;
	mac_t router_mac = ep->node_mac;

	/* This will invalidate the size check */
	ret = ipv6_l3(ctx, l3_off, (__u8 *) &router_mac, (__u8 *) &lxc_mac, direction);
	if (ret != CTX_ACT_OK)
		return ret;

#if defined LOCAL_DELIVERY_METRICS
	/*
	 * Special LXC case for updating egress forwarding metrics.
	 * Note that the packet could still be dropped but it would show up
	 * as an ingress drop counter in metrics.
	 */
	update_metrics(ctx->len, direction, REASON_FORWARDED);
#endif

#if defined USE_BPF_PROG_FOR_INGRESS_POLICY && !defined FORCE_LOCAL_POLICY_EVAL_AT_SOURCE
	ctx->mark = (seclabel << 16) | MARK_MAGIC_IDENTITY;
	return redirect_peer(ep->ifindex, 0);
#else
	ctx->cb[CB_SRC_LABEL] = seclabel;
	ctx->cb[CB_IFINDEX] = ep->ifindex;
	tail_call(ctx, &POLICY_CALL_MAP, ep->lxc_id);
	return DROP_MISSED_TAIL_CALL;
#endif
}
#endif /* ENABLE_IPV6 */

static __always_inline int ipv4_local_delivery(struct __ctx_buff *ctx, int l3_off, int l4_off,
					       __u32 seclabel, struct iphdr *ip4,
					       struct endpoint_info *ep, __u8 direction)
{
	int ret;

	cilium_dbg(ctx, DBG_LOCAL_DELIVERY, ep->lxc_id, seclabel);

	mac_t lxc_mac = ep->mac;
	mac_t router_mac = ep->node_mac;

	ret = ipv4_l3(ctx, l3_off, (__u8 *) &router_mac, (__u8 *) &lxc_mac, ip4);
	if (ret != CTX_ACT_OK)
		return ret;

#if defined LOCAL_DELIVERY_METRICS
	/*
	 * Special LXC case for updating egress forwarding metrics.
	 * Note that the packet could still be dropped but it would show up
	 * as an ingress drop counter in metrics.
	 */
	update_metrics(ctx->len, direction, REASON_FORWARDED);
#endif

#if defined USE_BPF_PROG_FOR_INGRESS_POLICY && !defined FORCE_LOCAL_POLICY_EVAL_AT_SOURCE
	ctx->mark = (seclabel << 16) | MARK_MAGIC_IDENTITY;
	return redirect_peer(ep->ifindex, 0);
#else
	ctx->cb[CB_SRC_LABEL] = seclabel;
	ctx->cb[CB_IFINDEX] = ep->ifindex;
	tail_call(ctx, &POLICY_CALL_MAP, ep->lxc_id);
	return DROP_MISSED_TAIL_CALL;
#endif
}

static __always_inline __u8 get_encrypt_key(__u32 ctx)
{
	struct encrypt_key key = {.ctx = ctx};
	struct encrypt_config *cfg;

	cfg = map_lookup_elem(&ENCRYPT_MAP, &key);
	/* Having no key info for a context is the same as no encryption */
	if (!cfg)
		return 0;
	return cfg->encrypt_key;
}

static __always_inline __u8 get_min_encrypt_key(__u8 peer_key)
{
	__u8 local_key = get_encrypt_key(0);

	/* If both ends can encrypt/decrypt use smaller of the two this
	 * way both ends will have keys installed assuming key IDs are
	 * always increasing. However, we have to handle roll-over case
	 * and to do this safely we assume keys are no more than one ahead.
	 * We expect user/control-place to accomplish this. Notice zero
	 * will always be returned if either local or peer have the zero
	 * key indicating no encryption.
	 */
	if (peer_key == MAX_KEY_INDEX)
		return local_key == 1 ? peer_key : local_key;
	if (local_key == MAX_KEY_INDEX)
		return peer_key == 1 ? local_key : peer_key;
	return local_key < peer_key ? local_key : peer_key;
}

#endif
