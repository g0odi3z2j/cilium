// Code generated by internal/cmd/gensections.awk; DO NOT EDIT.

package ebpf

// Code in this file is derived from libbpf, available under BSD-2-Clause.

import "github.com/cilium/ebpf/internal/sys"

var elfSectionDefs = []libbpfElfSectionDef{
	{"socket", sys.BPF_PROG_TYPE_SOCKET_FILTER, 0, _SEC_NONE},
	{"sk_reuseport/migrate", sys.BPF_PROG_TYPE_SK_REUSEPORT, sys.BPF_SK_REUSEPORT_SELECT_OR_MIGRATE, _SEC_ATTACHABLE},
	{"sk_reuseport", sys.BPF_PROG_TYPE_SK_REUSEPORT, sys.BPF_SK_REUSEPORT_SELECT, _SEC_ATTACHABLE},
	{"kprobe+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_NONE},
	{"uprobe+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_NONE},
	{"uprobe.s+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_SLEEPABLE},
	{"kretprobe+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_NONE},
	{"uretprobe+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_NONE},
	{"uretprobe.s+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_SLEEPABLE},
	{"kprobe.multi+", sys.BPF_PROG_TYPE_KPROBE, sys.BPF_TRACE_KPROBE_MULTI, _SEC_NONE},
	{"kretprobe.multi+", sys.BPF_PROG_TYPE_KPROBE, sys.BPF_TRACE_KPROBE_MULTI, _SEC_NONE},
	{"uprobe.multi+", sys.BPF_PROG_TYPE_KPROBE, sys.BPF_TRACE_UPROBE_MULTI, _SEC_NONE},
	{"uretprobe.multi+", sys.BPF_PROG_TYPE_KPROBE, sys.BPF_TRACE_UPROBE_MULTI, _SEC_NONE},
	{"uprobe.multi.s+", sys.BPF_PROG_TYPE_KPROBE, sys.BPF_TRACE_UPROBE_MULTI, _SEC_SLEEPABLE},
	{"uretprobe.multi.s+", sys.BPF_PROG_TYPE_KPROBE, sys.BPF_TRACE_UPROBE_MULTI, _SEC_SLEEPABLE},
	{"ksyscall+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_NONE},
	{"kretsyscall+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_NONE},
	{"usdt+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_USDT},
	{"usdt.s+", sys.BPF_PROG_TYPE_KPROBE, 0, _SEC_USDT | _SEC_SLEEPABLE},
	{"tc/ingress", sys.BPF_PROG_TYPE_SCHED_CLS, sys.BPF_TCX_INGRESS, _SEC_NONE},
	{"tc/egress", sys.BPF_PROG_TYPE_SCHED_CLS, sys.BPF_TCX_EGRESS, _SEC_NONE},
	{"tcx/ingress", sys.BPF_PROG_TYPE_SCHED_CLS, sys.BPF_TCX_INGRESS, _SEC_NONE},
	{"tcx/egress", sys.BPF_PROG_TYPE_SCHED_CLS, sys.BPF_TCX_EGRESS, _SEC_NONE},
	{"tc", sys.BPF_PROG_TYPE_SCHED_CLS, 0, _SEC_NONE},
	{"classifier", sys.BPF_PROG_TYPE_SCHED_CLS, 0, _SEC_NONE},
	{"action", sys.BPF_PROG_TYPE_SCHED_ACT, 0, _SEC_NONE},
	{"tracepoint+", sys.BPF_PROG_TYPE_TRACEPOINT, 0, _SEC_NONE},
	{"tp+", sys.BPF_PROG_TYPE_TRACEPOINT, 0, _SEC_NONE},
	{"raw_tracepoint+", sys.BPF_PROG_TYPE_RAW_TRACEPOINT, 0, _SEC_NONE},
	{"raw_tp+", sys.BPF_PROG_TYPE_RAW_TRACEPOINT, 0, _SEC_NONE},
	{"raw_tracepoint.w+", sys.BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE, 0, _SEC_NONE},
	{"raw_tp.w+", sys.BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE, 0, _SEC_NONE},
	{"tp_btf+", sys.BPF_PROG_TYPE_TRACING, sys.BPF_TRACE_RAW_TP, _SEC_ATTACH_BTF},
	{"fentry+", sys.BPF_PROG_TYPE_TRACING, sys.BPF_TRACE_FENTRY, _SEC_ATTACH_BTF},
	{"fmod_ret+", sys.BPF_PROG_TYPE_TRACING, sys.BPF_MODIFY_RETURN, _SEC_ATTACH_BTF},
	{"fexit+", sys.BPF_PROG_TYPE_TRACING, sys.BPF_TRACE_FEXIT, _SEC_ATTACH_BTF},
	{"fentry.s+", sys.BPF_PROG_TYPE_TRACING, sys.BPF_TRACE_FENTRY, _SEC_ATTACH_BTF | _SEC_SLEEPABLE},
	{"fmod_ret.s+", sys.BPF_PROG_TYPE_TRACING, sys.BPF_MODIFY_RETURN, _SEC_ATTACH_BTF | _SEC_SLEEPABLE},
	{"fexit.s+", sys.BPF_PROG_TYPE_TRACING, sys.BPF_TRACE_FEXIT, _SEC_ATTACH_BTF | _SEC_SLEEPABLE},
	{"freplace+", sys.BPF_PROG_TYPE_EXT, 0, _SEC_ATTACH_BTF},
	{"lsm+", sys.BPF_PROG_TYPE_LSM, sys.BPF_LSM_MAC, _SEC_ATTACH_BTF},
	{"lsm.s+", sys.BPF_PROG_TYPE_LSM, sys.BPF_LSM_MAC, _SEC_ATTACH_BTF | _SEC_SLEEPABLE},
	{"lsm_cgroup+", sys.BPF_PROG_TYPE_LSM, sys.BPF_LSM_CGROUP, _SEC_ATTACH_BTF},
	{"iter+", sys.BPF_PROG_TYPE_TRACING, sys.BPF_TRACE_ITER, _SEC_ATTACH_BTF},
	{"iter.s+", sys.BPF_PROG_TYPE_TRACING, sys.BPF_TRACE_ITER, _SEC_ATTACH_BTF | _SEC_SLEEPABLE},
	{"syscall", sys.BPF_PROG_TYPE_SYSCALL, 0, _SEC_SLEEPABLE},
	{"xdp.frags/devmap", sys.BPF_PROG_TYPE_XDP, sys.BPF_XDP_DEVMAP, _SEC_XDP_FRAGS},
	{"xdp/devmap", sys.BPF_PROG_TYPE_XDP, sys.BPF_XDP_DEVMAP, _SEC_ATTACHABLE},
	{"xdp.frags/cpumap", sys.BPF_PROG_TYPE_XDP, sys.BPF_XDP_CPUMAP, _SEC_XDP_FRAGS},
	{"xdp/cpumap", sys.BPF_PROG_TYPE_XDP, sys.BPF_XDP_CPUMAP, _SEC_ATTACHABLE},
	{"xdp.frags", sys.BPF_PROG_TYPE_XDP, sys.BPF_XDP, _SEC_XDP_FRAGS},
	{"xdp", sys.BPF_PROG_TYPE_XDP, sys.BPF_XDP, _SEC_ATTACHABLE_OPT},
	{"perf_event", sys.BPF_PROG_TYPE_PERF_EVENT, 0, _SEC_NONE},
	{"lwt_in", sys.BPF_PROG_TYPE_LWT_IN, 0, _SEC_NONE},
	{"lwt_out", sys.BPF_PROG_TYPE_LWT_OUT, 0, _SEC_NONE},
	{"lwt_xmit", sys.BPF_PROG_TYPE_LWT_XMIT, 0, _SEC_NONE},
	{"lwt_seg6local", sys.BPF_PROG_TYPE_LWT_SEG6LOCAL, 0, _SEC_NONE},
	{"sockops", sys.BPF_PROG_TYPE_SOCK_OPS, sys.BPF_CGROUP_SOCK_OPS, _SEC_ATTACHABLE_OPT},
	{"sk_skb/stream_parser", sys.BPF_PROG_TYPE_SK_SKB, sys.BPF_SK_SKB_STREAM_PARSER, _SEC_ATTACHABLE_OPT},
	{"sk_skb/stream_verdict", sys.BPF_PROG_TYPE_SK_SKB, sys.BPF_SK_SKB_STREAM_VERDICT, _SEC_ATTACHABLE_OPT},
	{"sk_skb", sys.BPF_PROG_TYPE_SK_SKB, 0, _SEC_NONE},
	{"sk_msg", sys.BPF_PROG_TYPE_SK_MSG, sys.BPF_SK_MSG_VERDICT, _SEC_ATTACHABLE_OPT},
	{"lirc_mode2", sys.BPF_PROG_TYPE_LIRC_MODE2, sys.BPF_LIRC_MODE2, _SEC_ATTACHABLE_OPT},
	{"flow_dissector", sys.BPF_PROG_TYPE_FLOW_DISSECTOR, sys.BPF_FLOW_DISSECTOR, _SEC_ATTACHABLE_OPT},
	{"cgroup_skb/ingress", sys.BPF_PROG_TYPE_CGROUP_SKB, sys.BPF_CGROUP_INET_INGRESS, _SEC_ATTACHABLE_OPT},
	{"cgroup_skb/egress", sys.BPF_PROG_TYPE_CGROUP_SKB, sys.BPF_CGROUP_INET_EGRESS, _SEC_ATTACHABLE_OPT},
	{"cgroup/skb", sys.BPF_PROG_TYPE_CGROUP_SKB, 0, _SEC_NONE},
	{"cgroup/sock_create", sys.BPF_PROG_TYPE_CGROUP_SOCK, sys.BPF_CGROUP_INET_SOCK_CREATE, _SEC_ATTACHABLE},
	{"cgroup/sock_release", sys.BPF_PROG_TYPE_CGROUP_SOCK, sys.BPF_CGROUP_INET_SOCK_RELEASE, _SEC_ATTACHABLE},
	{"cgroup/sock", sys.BPF_PROG_TYPE_CGROUP_SOCK, sys.BPF_CGROUP_INET_SOCK_CREATE, _SEC_ATTACHABLE_OPT},
	{"cgroup/post_bind4", sys.BPF_PROG_TYPE_CGROUP_SOCK, sys.BPF_CGROUP_INET4_POST_BIND, _SEC_ATTACHABLE},
	{"cgroup/post_bind6", sys.BPF_PROG_TYPE_CGROUP_SOCK, sys.BPF_CGROUP_INET6_POST_BIND, _SEC_ATTACHABLE},
	{"cgroup/bind4", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_INET4_BIND, _SEC_ATTACHABLE},
	{"cgroup/bind6", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_INET6_BIND, _SEC_ATTACHABLE},
	{"cgroup/connect4", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_INET4_CONNECT, _SEC_ATTACHABLE},
	{"cgroup/connect6", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_INET6_CONNECT, _SEC_ATTACHABLE},
	{"cgroup/sendmsg4", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_UDP4_SENDMSG, _SEC_ATTACHABLE},
	{"cgroup/sendmsg6", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_UDP6_SENDMSG, _SEC_ATTACHABLE},
	{"cgroup/recvmsg4", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_UDP4_RECVMSG, _SEC_ATTACHABLE},
	{"cgroup/recvmsg6", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_UDP6_RECVMSG, _SEC_ATTACHABLE},
	{"cgroup/getpeername4", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_INET4_GETPEERNAME, _SEC_ATTACHABLE},
	{"cgroup/getpeername6", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_INET6_GETPEERNAME, _SEC_ATTACHABLE},
	{"cgroup/getsockname4", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_INET4_GETSOCKNAME, _SEC_ATTACHABLE},
	{"cgroup/getsockname6", sys.BPF_PROG_TYPE_CGROUP_SOCK_ADDR, sys.BPF_CGROUP_INET6_GETSOCKNAME, _SEC_ATTACHABLE},
	{"cgroup/sysctl", sys.BPF_PROG_TYPE_CGROUP_SYSCTL, sys.BPF_CGROUP_SYSCTL, _SEC_ATTACHABLE},
	{"cgroup/getsockopt", sys.BPF_PROG_TYPE_CGROUP_SOCKOPT, sys.BPF_CGROUP_GETSOCKOPT, _SEC_ATTACHABLE},
	{"cgroup/setsockopt", sys.BPF_PROG_TYPE_CGROUP_SOCKOPT, sys.BPF_CGROUP_SETSOCKOPT, _SEC_ATTACHABLE},
	{"cgroup/dev", sys.BPF_PROG_TYPE_CGROUP_DEVICE, sys.BPF_CGROUP_DEVICE, _SEC_ATTACHABLE_OPT},
	{"struct_ops+", sys.BPF_PROG_TYPE_STRUCT_OPS, 0, _SEC_NONE},
	{"struct_ops.s+", sys.BPF_PROG_TYPE_STRUCT_OPS, 0, _SEC_SLEEPABLE},
	{"sk_lookup", sys.BPF_PROG_TYPE_SK_LOOKUP, sys.BPF_SK_LOOKUP, _SEC_ATTACHABLE},
	{"netfilter", sys.BPF_PROG_TYPE_NETFILTER, sys.BPF_NETFILTER, _SEC_NONE},
}
