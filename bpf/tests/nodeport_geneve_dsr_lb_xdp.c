// SPDX-License-Identifier: (GPL-2.0-only OR BSD-2-Clause)
/* Copyright Authors of Cilium */

#include "common.h"

#include <bpf/ctx/xdp.h>
#include "pktgen.h"

/* Set ETH_HLEN to 14 to indicate that the packet has a 14 byte ethernet header */
#define ETH_HLEN 14

/* Enable code paths under test */
#define ENABLE_IPV4
#define ENABLE_NODEPORT
#define ENABLE_NODEPORT_ACCELERATION
#define ENABLE_DSR		1

#define DSR_ENCAP_IPIP		2
#define DSR_ENCAP_GENEVE	3
#define DSR_ENCAP_MODE		DSR_ENCAP_GENEVE

#define TUNNEL_PROTOCOL		TUNNEL_PROTOCOL_GENEVE
#define ENCAP_IFINDEX		42
#define TUNNEL_MODE

#define DISABLE_LOOPBACK_LB

/* Skip ingress policy checks, not needed to validate hairpin flow */
#define USE_BPF_PROG_FOR_INGRESS_POLICY
#undef FORCE_LOCAL_POLICY_EVAL_AT_SOURCE

#define CLIENT_IP		v4_ext_one
#define CLIENT_PORT		__bpf_htons(111)

#define FRONTEND_IP_LOCAL	v4_svc_one
#define FRONTEND_IP_REMOTE	v4_svc_two
#define FRONTEND_PORT		tcp_svc_one

#define LB_IP			v4_node_one
#define IPV4_DIRECT_ROUTING	LB_IP
#define BACKEND_NODE_IP		v4_node_two

#define DIRECT_ROUTING_IFINDEX	25

#define BACKEND_IP_LOCAL	v4_pod_one
#define BACKEND_IP_REMOTE	v4_pod_two
#define BACKEND_PORT		__bpf_htons(8080)

#define fib_lookup mock_fib_lookup

static volatile const __u8 *client_mac = mac_one;
/* this matches the default node_config.h: */
static volatile const __u8 lb_mac[ETH_ALEN]	= { 0xce, 0x72, 0xa7, 0x03, 0x88, 0x56 };
static volatile const __u8 *node_mac = mac_three;
static volatile const __u8 *local_backend_mac = mac_four;
static volatile const __u8 *backend_node_mac = mac_six;

static bool fail_fib;

#define ctx_redirect mock_ctx_redirect
static __always_inline __maybe_unused int
mock_ctx_redirect(const struct __ctx_buff *ctx __maybe_unused, int ifindex __maybe_unused,
		  __u32 flags __maybe_unused)
{
	if (ifindex != DIRECT_ROUTING_IFINDEX)
		return CTX_ACT_DROP;

	return CTX_ACT_REDIRECT;
}

long mock_fib_lookup(__maybe_unused void *ctx, struct bpf_fib_lookup *params,
		     __maybe_unused int plen, __maybe_unused __u32 flags)
{
	if (fail_fib)
		return BPF_FIB_LKUP_RET_NO_NEIGH;

	params->ifindex = DIRECT_ROUTING_IFINDEX;

	if (params->ipv4_dst == BACKEND_NODE_IP) {
		__bpf_memcpy_builtin(params->smac, (__u8 *)lb_mac, ETH_ALEN);
		__bpf_memcpy_builtin(params->dmac, (__u8 *)backend_node_mac, ETH_ALEN);
	} else {
		return CTX_ACT_DROP;
	}

	return 0;
}

#include <bpf_xdp.c>

#define FROM_NETDEV	0

struct {
	__uint(type, BPF_MAP_TYPE_PROG_ARRAY);
	__uint(key_size, sizeof(__u32));
	__uint(max_entries, 1);
	__array(values, int());
} entry_call_map __section(".maps") = {
	.values = {
		[FROM_NETDEV] = &cil_xdp_entry,
	},
};

/* Test that a SVC request to a local backend
 * - gets DNATed (but not SNATed)
 * - gets passed up from XDP to TC
 */
PKTGEN("xdp", "nodeport_geneve_dsr_lb_xdp1_local_backend")
int nodeport_geneve_dsr_lb_xdp1_local_backend_pktgen(struct __ctx_buff *ctx)
{
	struct pktgen builder;
	struct tcphdr *l4;
	struct ethhdr *l2;
	struct iphdr *l3;
	void *data;

	/* Init packet builder */
	pktgen__init(&builder, ctx);

	/* Push ethernet header */
	l2 = pktgen__push_ethhdr(&builder);
	if (!l2)
		return TEST_ERROR;

	ethhdr__set_macs(l2, (__u8 *)client_mac, (__u8 *)lb_mac);

	/* Push IPv4 header */
	l3 = pktgen__push_default_iphdr(&builder);
	if (!l3)
		return TEST_ERROR;

	l3->saddr = CLIENT_IP;
	l3->daddr = FRONTEND_IP_LOCAL;

	/* Push TCP header */
	l4 = pktgen__push_default_tcphdr(&builder);
	if (!l4)
		return TEST_ERROR;

	l4->source = CLIENT_PORT;
	l4->dest = FRONTEND_PORT;

	data = pktgen__push_data(&builder, default_data, sizeof(default_data));
	if (!data)
		return TEST_ERROR;

	/* Calc lengths, set protocol fields and calc checksums */
	pktgen__finish(&builder);

	return 0;
}

SETUP("xdp", "nodeport_geneve_dsr_lb_xdp1_local_backend")
int nodeport_geneve_dsr_lb_xdp1_local_backend_setup(struct __ctx_buff *ctx)
{
	__u16 revnat_id = 1;

	/* Register a fake LB backend matching our packet. */
	struct lb4_key lb_svc_key = {
		.address = FRONTEND_IP_LOCAL,
		.dport = FRONTEND_PORT,
		.scope = LB_LOOKUP_SCOPE_EXT,
	};
	/* Create a service with only one backend */
	struct lb4_service lb_svc_value = {
		.count = 1,
		.flags = SVC_FLAG_ROUTABLE,
		.rev_nat_index = revnat_id,
	};
	map_update_elem(&LB4_SERVICES_MAP_V2, &lb_svc_key, &lb_svc_value, BPF_ANY);
	/* We need to register both in the external and internal scopes for the
	 * packet to be redirected to a neighboring node
	 */
	lb_svc_key.scope = LB_LOOKUP_SCOPE_INT;
	map_update_elem(&LB4_SERVICES_MAP_V2, &lb_svc_key, &lb_svc_value, BPF_ANY);

	/* A backend between 1 and .count is chosen, since we have only one backend
	 * it is always backend_slot 1. Point it to backend_id 124.
	 */
	lb_svc_key.scope = LB_LOOKUP_SCOPE_EXT;
	lb_svc_key.backend_slot = 1;
	lb_svc_value.backend_id = 124;
	map_update_elem(&LB4_SERVICES_MAP_V2, &lb_svc_key, &lb_svc_value, BPF_ANY);

	/* Insert a reverse NAT entry for the above service */
	struct lb4_reverse_nat revnat_value = {
		.address = FRONTEND_IP_LOCAL,
		.port = FRONTEND_PORT,
	};
	map_update_elem(&LB4_REVERSE_NAT_MAP, &revnat_id, &revnat_value, BPF_ANY);

	struct lb4_backend backend = {
		.address = BACKEND_IP_LOCAL,
		.port = BACKEND_PORT,
		.proto = IPPROTO_TCP,
		.flags = BE_STATE_ACTIVE,
	};
	map_update_elem(&LB4_BACKEND_MAP, &lb_svc_value.backend_id, &backend, BPF_ANY);

	/* add local backend */
	struct endpoint_info ep_value = {};

	memcpy(&ep_value.mac, (__u8 *)local_backend_mac, ETH_ALEN);
	memcpy(&ep_value.node_mac, (__u8 *)node_mac, ETH_ALEN);

	struct endpoint_key ep_key = {
		.family = ENDPOINT_KEY_IPV4,
		.ip4 = BACKEND_IP_LOCAL,
	};
	map_update_elem(&ENDPOINTS_MAP, &ep_key, &ep_value, BPF_ANY);

	struct ipcache_key cache_key = {
		.lpm_key.prefixlen = IPCACHE_PREFIX_LEN(32),
		.family = ENDPOINT_KEY_IPV4,
		.ip4 = BACKEND_IP_LOCAL,
	};
	struct remote_endpoint_info cache_value = {
		.sec_identity = 112233,
	};
	map_update_elem(&IPCACHE_MAP, &cache_key, &cache_value, BPF_ANY);

	/* Jump into the entrypoint */
	tail_call_static(ctx, &entry_call_map, FROM_NETDEV);
	/* Fail if we didn't jump */
	return TEST_ERROR;
}

CHECK("xdp", "nodeport_geneve_dsr_lb_xdp1_local_backend")
int nodeport_geneve_dsr_lb_xdp1_local_backend_check(const struct __ctx_buff *ctx)
{
	void *data, *data_end;
	__u32 *status_code;
	struct tcphdr *l4;
	struct ethhdr *l2;
	struct iphdr *l3;
	__u32 *meta;

	test_init();

	data = (void *)(long)ctx_data(ctx);
	data_end = (void *)(long)ctx->data_end;

	status_code = data;
	if (data + sizeof(__u32) > data_end)
		test_fatal("status code out of bounds");

	meta = (void *)status_code + sizeof(__u32);
	if ((void *)meta + sizeof(__u32) > data_end)
		test_fatal("meta out of bounds");

	l2 = (void *)meta + sizeof(__u32);
	if ((void *)l2 + sizeof(struct ethhdr) > data_end)
		test_fatal("l2 out of bounds");

	l3 = (void *)l2 + sizeof(struct ethhdr);
	if ((void *)l3 + sizeof(struct iphdr) > data_end)
		test_fatal("l3 out of bounds");

	l4 = (void *)l3 + sizeof(struct iphdr);
	if ((void *)l4 + sizeof(struct tcphdr) > data_end)
		test_fatal("l4 out of bounds");

	assert(*status_code == CTX_ACT_OK);

	assert((*meta & XFER_PKT_NO_SVC) == XFER_PKT_NO_SVC);

	if (memcmp(l2->h_source, (__u8 *)client_mac, ETH_ALEN) != 0)
		test_fatal("src MAC is not the client MAC");
	if (memcmp(l2->h_dest, (__u8 *)lb_mac, ETH_ALEN) != 0)
		test_fatal("dst MAC is not the LB MAC");

	if (l3->saddr != CLIENT_IP)
		test_fatal("src IP has changed");

	if (l3->daddr != BACKEND_IP_LOCAL)
		test_fatal("dst IP hasn't been NATed to local backend IP");

	if (l4->source != CLIENT_PORT)
		test_fatal("src port has changed");

	if (l4->dest != BACKEND_PORT)
		test_fatal("dst TCP port hasn't been NATed to backend port");

	test_finish();
}

/* Test that a SVC request that is LBed to a DSR remote backend
 * - gets DNATed,
 * - has tunnel encapsulation header added,
 * - has DSR option inserted
 */
PKTGEN("xdp", "nodeport_geneve_dsr_lb_xdp2_fwd")
int nodeport_geneve_dsr_lb_xdp2_fwd_pktgen(struct __ctx_buff *ctx)
{
	struct pktgen builder;
	struct tcphdr *l4;
	struct ethhdr *l2;
	struct iphdr *l3;
	void *data;

	/* Init packet builder */
	pktgen__init(&builder, ctx);

	/* Push ethernet header */
	l2 = pktgen__push_ethhdr(&builder);
	if (!l2)
		return TEST_ERROR;

	ethhdr__set_macs(l2, (__u8 *)client_mac, (__u8 *)lb_mac);

	/* Push IPv4 header */
	l3 = pktgen__push_default_iphdr(&builder);
	if (!l3)
		return TEST_ERROR;

	l3->saddr = CLIENT_IP;
	l3->daddr = FRONTEND_IP_REMOTE;

	/* Push TCP header */
	l4 = pktgen__push_default_tcphdr(&builder);
	if (!l4)
		return TEST_ERROR;

	l4->source = CLIENT_PORT;
	l4->dest = FRONTEND_PORT;

	data = pktgen__push_data(&builder, default_data, sizeof(default_data));
	if (!data)
		return TEST_ERROR;

	/* Calc lengths, set protocol fields and calc checksums */
	pktgen__finish(&builder);

	return 0;
}

SETUP("xdp", "nodeport_geneve_dsr_lb_xdp2_fwd")
int nodeport_geneve_dsr_lb_xdp2_fwd_setup(struct __ctx_buff *ctx)
{
	__u16 revnat_id = 2;

	/* Register a fake LB backend matching our packet. */
	struct lb4_key lb_svc_key = {
		.address = FRONTEND_IP_REMOTE,
		.dport = FRONTEND_PORT,
		.scope = LB_LOOKUP_SCOPE_EXT,
	};
	/* Create a service with only one backend */
	struct lb4_service lb_svc_value = {
		.count = 1,
		.flags = SVC_FLAG_ROUTABLE,
		.rev_nat_index = revnat_id,
	};
	map_update_elem(&LB4_SERVICES_MAP_V2, &lb_svc_key, &lb_svc_value, BPF_ANY);
	/* We need to register both in the external and internal scopes for the
	 * packet to be redirected to a neighboring node
	 */
	lb_svc_key.scope = LB_LOOKUP_SCOPE_INT;
	map_update_elem(&LB4_SERVICES_MAP_V2, &lb_svc_key, &lb_svc_value, BPF_ANY);

	/* A backend between 1 and .count is chosen, since we have only one backend
	 * it is always backend_slot 1.
	 */
	lb_svc_key.scope = LB_LOOKUP_SCOPE_EXT;
	lb_svc_key.backend_slot = 1;
	lb_svc_value.backend_id = 125;
	map_update_elem(&LB4_SERVICES_MAP_V2, &lb_svc_key, &lb_svc_value, BPF_ANY);

	/* Insert a reverse NAT entry for the above service */
	struct lb4_reverse_nat revnat_value = {
		.address = FRONTEND_IP_REMOTE,
		.port = FRONTEND_PORT,
	};
	map_update_elem(&LB4_REVERSE_NAT_MAP, &revnat_id, &revnat_value, BPF_ANY);

	struct lb4_backend backend = {
		.address = BACKEND_IP_REMOTE,
		.port = BACKEND_PORT,
		.proto = IPPROTO_TCP,
		.flags = BE_STATE_ACTIVE,
	};
	map_update_elem(&LB4_BACKEND_MAP, &lb_svc_value.backend_id, &backend, BPF_ANY);

	struct ipcache_key cache_key = {
		.lpm_key.prefixlen = IPCACHE_PREFIX_LEN(32),
		.family = ENDPOINT_KEY_IPV4,
		.ip4 = BACKEND_IP_REMOTE,
	};
	struct remote_endpoint_info cache_value = {
		.sec_identity = 112233,
		.tunnel_endpoint = BACKEND_NODE_IP,
	};
	map_update_elem(&IPCACHE_MAP, &cache_key, &cache_value, BPF_ANY);

	/* Jump into the entrypoint */
	tail_call_static(ctx, &entry_call_map, FROM_NETDEV);
	/* Fail if we didn't jump */
	return TEST_ERROR;
}

CHECK("xdp", "nodeport_geneve_dsr_lb_xdp2_fwd")
int nodeport_geneve_dsr_lb_xdp_fwd_check(__maybe_unused const struct __ctx_buff *ctx)
{
	struct geneve_dsr_opt4 *dsr_opt;
	struct ethhdr *l2, *inner_l2;
	struct iphdr *l3, *inner_l3;
	struct tcphdr *tcp_inner;
	struct genevehdr *geneve;
	void *data, *data_end;
	__u32 *status_code;
	struct udphdr *udp;

	test_init();

	data = (void *)(long)ctx_data(ctx);
	data_end = (void *)(long)ctx->data_end;

	if (data + sizeof(__u32) > data_end)
		test_fatal("status code out of bounds");

	status_code = data;

	assert(*status_code == CTX_ACT_REDIRECT);

	l2 = data + sizeof(__u32);
	if ((void *)l2 + sizeof(*l2) > data_end)
		test_fatal("l2 out of bounds");

	l3 = (void *)l2 + sizeof(*l2);
	if ((void *)l3 + sizeof(*l3) > data_end)
		test_fatal("l3 out of bounds");

	udp = (void *)l3 + sizeof(*l3);
	if ((void *)udp + sizeof(*udp) > data_end)
		test_fatal("udp out of bounds");

	geneve = (void *)udp + sizeof(*udp);
	if ((void *)geneve + sizeof(*geneve) > data_end)
		test_fatal("geneve out of bounds");

	dsr_opt = (void *)geneve + sizeof(*geneve);
	if ((void *)dsr_opt + sizeof(*dsr_opt) > data_end)
		test_fatal("dsr opt out of bounds");
	if ((void *)dsr_opt + geneve->opt_len * 4 > data_end)
		test_fatal("geneve opts out of bounds");

	inner_l2 = (void *)dsr_opt + geneve->opt_len * 4;
	if ((void *)inner_l2 + sizeof(*inner_l2) > data_end)
		test_fatal("l2 out of bounds");

	inner_l3 = (void *)inner_l2 + sizeof(*inner_l2);
	if ((void *)inner_l3 + sizeof(*inner_l3) > data_end)
		test_fatal("l3 out of bounds");

	tcp_inner = (void *)inner_l3 + sizeof(*inner_l3);
	if ((void *)tcp_inner + sizeof(*tcp_inner) > data_end)
		test_fatal("tcp out of bounds");

	if (memcmp(l2->h_source, (__u8 *)lb_mac, ETH_ALEN) != 0)
		test_fatal("src MAC is not the LB MAC");
	if (memcmp(l2->h_dest, (__u8 *)backend_node_mac, ETH_ALEN) != 0)
		test_fatal("dst MAC is not the backend node MAC");

	if (l2->h_proto != bpf_htons(ETH_P_IP))
		test_fatal("l2 doesn't have correct proto type");

	if (l3->protocol != IPPROTO_UDP)
		test_fatal("outer IP doesn't have correct L4 protocol");

	if (l3->saddr != IPV4_DIRECT_ROUTING)
		test_fatal("outerSrcIP is not correct");

	if (l3->daddr != BACKEND_NODE_IP)
		test_fatal("outerDstIP is not correct");

	if (udp->dest != bpf_htons(TUNNEL_PORT))
		test_fatal("outerDstPort is not tunnel port");

	__be32 sec_id;

	memcpy(&sec_id, &geneve->vni, 4);
	if (tunnel_vni_to_sec_identity(sec_id) != WORLD_ID)
		test_fatal("geneve has unexpected SrcSecID");

	if (geneve->opt_len * 4 != sizeof(*dsr_opt))
		test_fatal("geneve has unexpected opt length");

	if (dsr_opt->hdr.opt_class != bpf_htons(DSR_GENEVE_OPT_CLASS))
		test_fatal("geneve opt has unexpected class");
	if (dsr_opt->hdr.type != DSR_GENEVE_OPT_TYPE)
		test_fatal("geneve opt has unexpected type");
	if (dsr_opt->hdr.length != DSR_IPV4_GENEVE_OPT_LEN)
		test_fatal("geneve opt has unexpected length");
	if (dsr_opt->addr != FRONTEND_IP_REMOTE)
		test_fatal("geneve opt has unexpected svc IP");
	if (dsr_opt->port != FRONTEND_PORT)
		test_fatal("geneve opt has unexpected svc port");

	if (inner_l2->h_proto != bpf_htons(ETH_P_IP))
		test_fatal("inner l2 doesn't have correct proto type");

	if (inner_l3->protocol != IPPROTO_TCP)
		test_fatal("inner IP doesn't have correct L4 protocol");

	if (inner_l3->saddr != CLIENT_IP)
		test_fatal("innerSrcIP has changed");

	if (inner_l3->daddr != BACKEND_IP_REMOTE)
		test_fatal("innerDstIP hasn't been NATed to remote backend IP");

	if (tcp_inner->source != CLIENT_PORT)
		test_fatal("innerSrcPort has changed");

	if (tcp_inner->dest != BACKEND_PORT)
		test_fatal("innerDstPort hasn't been NATed to backend port");

	test_finish();
}
