// SPDX-License-Identifier: (GPL-2.0-only OR BSD-2-Clause)
/* Copyright Authors of Cilium */

#include "common.h"

#include <bpf/ctx/skb.h>
#include "pktgen.h"

#define LXC_IPV4 (__be32)v4_pod_one
#include "config_replacement.h"

/* Set ETH_HLEN to 14 to indicate that the packet has a 14 byte ethernet header */
#define ETH_HLEN 14

/* Enable code paths under test */
#define ENABLE_IPV4
#define ENABLE_NODEPORT
#define ENABLE_EGRESS_GATEWAY
#define ENABLE_MASQUERADE_IPV4
#define ENCAP_IFINDEX 0

#define SECCTX_FROM_IPCACHE 1

#include "bpf_lxc.c"

#include "lib/egressgw.h"

#define FROM_CONTAINER 0

struct {
	__uint(type, BPF_MAP_TYPE_PROG_ARRAY);
	__uint(key_size, sizeof(__u32));
	__uint(max_entries, 1);
	__array(values, int());
} entry_call_map __section(".maps") = {
	.values = {
		[FROM_CONTAINER] = &cil_from_container,
	},
};

/* Test that a packet matching an egress gateway policy on the from-container
 * program gets redirected to the gateway node.
 */
PKTGEN("tc", "tc_egressgw_redirect")
int egressgw_redirect_pktgen(struct __ctx_buff *ctx)
{
	return egressgw_pktgen(ctx, (struct egressgw_test_ctx) {
			.test = TEST_REDIRECT,
		});
}

SETUP("tc", "tc_egressgw_redirect")
int egressgw_redirect_setup(struct __ctx_buff *ctx)
{
	add_egressgw_policy_entry(CLIENT_IP, EXTERNAL_SVC_IP & 0xffffff, 24, GATEWAY_NODE_IP, 0);

	/* Avoid policy drop */
	add_allow_all_egress_policy();

	/* Jump into the entrypoint */
	tail_call_static(ctx, &entry_call_map, FROM_CONTAINER);
	/* Fail if we didn't jump */
	return TEST_ERROR;
}

CHECK("tc", "tc_egressgw_redirect")
int egressgw_redirect_check(const struct __ctx_buff *ctx)
{
	void *data, *data_end;
	__u32 *status_code;

	test_init();

	data = (void *)(long)ctx_data(ctx);
	data_end = (void *)(long)ctx->data_end;

	if (data + sizeof(__u32) > data_end)
		test_fatal("status code out of bounds");

	status_code = data;
	assert(*status_code == CTX_ACT_REDIRECT);

	del_allow_all_egress_policy();
	del_egressgw_policy_entry(CLIENT_IP, EXTERNAL_SVC_IP & 0xffffff, 24);

	test_finish();
}

/* Test that a packet matching an excluded CIDR egress gateway policy on the
 * from-container program does not get redirected to the gateway node.
 */
PKTGEN("tc", "tc_egressgw_skip_excluded_cidr_redirect")
int egressgw_skip_excluded_cidr_redirect_pktgen(struct __ctx_buff *ctx)
{
	return egressgw_pktgen(ctx, (struct egressgw_test_ctx) {
			.test = TEST_REDIRECT,
		});
}

SETUP("tc", "tc_egressgw_skip_excluded_cidr_redirect")
int egressgw_skip_excluded_cidr_redirect_setup(struct __ctx_buff *ctx)
{
	add_egressgw_policy_entry(CLIENT_IP, EXTERNAL_SVC_IP & 0xffffff, 24, GATEWAY_NODE_IP, 0);
	add_egressgw_policy_entry(CLIENT_IP, EXTERNAL_SVC_IP, 32, EGRESS_GATEWAY_EXCLUDED_CIDR, 0);

	/* Avoid policy drop */
	add_allow_all_egress_policy();

	/* Jump into the entrypoint */
	tail_call_static(ctx, &entry_call_map, FROM_CONTAINER);
	/* Fail if we didn't jump */
	return TEST_ERROR;
}

CHECK("tc", "tc_egressgw_skip_excluded_cidr_redirect")
int egressgw_skip_excluded_cidr_redirect_check(const struct __ctx_buff *ctx)
{
	void *data, *data_end;
	__u32 *status_code;

	test_init();

	data = (void *)(long)ctx_data(ctx);
	data_end = (void *)(long)ctx->data_end;

	if (data + sizeof(__u32) > data_end)
		test_fatal("status code out of bounds");

	status_code = data;
	assert(*status_code == CTX_ACT_OK);

	del_allow_all_egress_policy();
	del_egressgw_policy_entry(CLIENT_IP, EXTERNAL_SVC_IP & 0xffffff, 24);
	del_egressgw_policy_entry(CLIENT_IP, EXTERNAL_SVC_IP, 32);

	test_finish();
}

/* Test that a packet matching an egress gateway policy without a gateway on the
 * from-container program does not get redirected to the gateway node.
 */
PKTGEN("tc", "tc_egressgw_skip_no_gateway_redirect")
int egressgw_skip_no_gateway_redirect_pktgen(struct __ctx_buff *ctx)
{

	return egressgw_pktgen(ctx, (struct egressgw_test_ctx) {
			.test = TEST_REDIRECT_SKIP_NO_GATEWAY,
		});
}

SETUP("tc", "tc_egressgw_skip_no_gateway_redirect")
int egressgw_skip_no_gateway_redirect_setup(struct __ctx_buff *ctx)
{
	add_egressgw_policy_entry(CLIENT_IP, EXTERNAL_SVC_IP, 32, EGRESS_GATEWAY_NO_GATEWAY, 0);

	/* Avoid policy drop */
	add_allow_all_egress_policy();

	/* Jump into the entrypoint */
	tail_call_static(ctx, &entry_call_map, FROM_CONTAINER);
	/* Fail if we didn't jump */
	return TEST_ERROR;
}

CHECK("tc", "tc_egressgw_skip_no_gateway_redirect")
int egressgw_skip_no_gateway_redirect_check(const struct __ctx_buff *ctx)
{
	void *data, *data_end;
	__u32 *status_code;

	struct metrics_value *entry = NULL;
	struct metrics_key key = {};

	test_init();

	data = (void *)(long)ctx_data(ctx);
	data_end = (void *)(long)ctx->data_end;

	if (data + sizeof(__u32) > data_end)
		test_fatal("status code out of bounds");

	status_code = data;
	assert(*status_code == CTX_ACT_DROP);

	key.reason = (__u8)-DROP_NO_EGRESS_GATEWAY;
	key.dir = METRIC_EGRESS;
	entry = map_lookup_elem(&METRICS_MAP, &key);
	if (!entry)
		test_fatal("metrics entry not found");
	assert(entry->count == 1);

	del_allow_all_egress_policy();
	del_egressgw_policy_entry(CLIENT_IP, EXTERNAL_SVC_IP, 32);

	test_finish();
}
