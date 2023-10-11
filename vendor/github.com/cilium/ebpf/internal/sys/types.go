// Code generated by internal/cmd/gentypes; DO NOT EDIT.

package sys

import (
	"unsafe"
)

type AdjRoomMode uint32

const (
	BPF_ADJ_ROOM_NET AdjRoomMode = 0
	BPF_ADJ_ROOM_MAC AdjRoomMode = 1
)

type AttachType uint32

const (
	BPF_CGROUP_INET_INGRESS            AttachType = 0
	BPF_CGROUP_INET_EGRESS             AttachType = 1
	BPF_CGROUP_INET_SOCK_CREATE        AttachType = 2
	BPF_CGROUP_SOCK_OPS                AttachType = 3
	BPF_SK_SKB_STREAM_PARSER           AttachType = 4
	BPF_SK_SKB_STREAM_VERDICT          AttachType = 5
	BPF_CGROUP_DEVICE                  AttachType = 6
	BPF_SK_MSG_VERDICT                 AttachType = 7
	BPF_CGROUP_INET4_BIND              AttachType = 8
	BPF_CGROUP_INET6_BIND              AttachType = 9
	BPF_CGROUP_INET4_CONNECT           AttachType = 10
	BPF_CGROUP_INET6_CONNECT           AttachType = 11
	BPF_CGROUP_INET4_POST_BIND         AttachType = 12
	BPF_CGROUP_INET6_POST_BIND         AttachType = 13
	BPF_CGROUP_UDP4_SENDMSG            AttachType = 14
	BPF_CGROUP_UDP6_SENDMSG            AttachType = 15
	BPF_LIRC_MODE2                     AttachType = 16
	BPF_FLOW_DISSECTOR                 AttachType = 17
	BPF_CGROUP_SYSCTL                  AttachType = 18
	BPF_CGROUP_UDP4_RECVMSG            AttachType = 19
	BPF_CGROUP_UDP6_RECVMSG            AttachType = 20
	BPF_CGROUP_GETSOCKOPT              AttachType = 21
	BPF_CGROUP_SETSOCKOPT              AttachType = 22
	BPF_TRACE_RAW_TP                   AttachType = 23
	BPF_TRACE_FENTRY                   AttachType = 24
	BPF_TRACE_FEXIT                    AttachType = 25
	BPF_MODIFY_RETURN                  AttachType = 26
	BPF_LSM_MAC                        AttachType = 27
	BPF_TRACE_ITER                     AttachType = 28
	BPF_CGROUP_INET4_GETPEERNAME       AttachType = 29
	BPF_CGROUP_INET6_GETPEERNAME       AttachType = 30
	BPF_CGROUP_INET4_GETSOCKNAME       AttachType = 31
	BPF_CGROUP_INET6_GETSOCKNAME       AttachType = 32
	BPF_XDP_DEVMAP                     AttachType = 33
	BPF_CGROUP_INET_SOCK_RELEASE       AttachType = 34
	BPF_XDP_CPUMAP                     AttachType = 35
	BPF_SK_LOOKUP                      AttachType = 36
	BPF_XDP                            AttachType = 37
	BPF_SK_SKB_VERDICT                 AttachType = 38
	BPF_SK_REUSEPORT_SELECT            AttachType = 39
	BPF_SK_REUSEPORT_SELECT_OR_MIGRATE AttachType = 40
	BPF_PERF_EVENT                     AttachType = 41
	BPF_TRACE_KPROBE_MULTI             AttachType = 42
	BPF_LSM_CGROUP                     AttachType = 43
	__MAX_BPF_ATTACH_TYPE              AttachType = 44
)

type Cmd uint32

const (
	BPF_MAP_CREATE                  Cmd = 0
	BPF_MAP_LOOKUP_ELEM             Cmd = 1
	BPF_MAP_UPDATE_ELEM             Cmd = 2
	BPF_MAP_DELETE_ELEM             Cmd = 3
	BPF_MAP_GET_NEXT_KEY            Cmd = 4
	BPF_PROG_LOAD                   Cmd = 5
	BPF_OBJ_PIN                     Cmd = 6
	BPF_OBJ_GET                     Cmd = 7
	BPF_PROG_ATTACH                 Cmd = 8
	BPF_PROG_DETACH                 Cmd = 9
	BPF_PROG_TEST_RUN               Cmd = 10
	BPF_PROG_RUN                    Cmd = 10
	BPF_PROG_GET_NEXT_ID            Cmd = 11
	BPF_MAP_GET_NEXT_ID             Cmd = 12
	BPF_PROG_GET_FD_BY_ID           Cmd = 13
	BPF_MAP_GET_FD_BY_ID            Cmd = 14
	BPF_OBJ_GET_INFO_BY_FD          Cmd = 15
	BPF_PROG_QUERY                  Cmd = 16
	BPF_RAW_TRACEPOINT_OPEN         Cmd = 17
	BPF_BTF_LOAD                    Cmd = 18
	BPF_BTF_GET_FD_BY_ID            Cmd = 19
	BPF_TASK_FD_QUERY               Cmd = 20
	BPF_MAP_LOOKUP_AND_DELETE_ELEM  Cmd = 21
	BPF_MAP_FREEZE                  Cmd = 22
	BPF_BTF_GET_NEXT_ID             Cmd = 23
	BPF_MAP_LOOKUP_BATCH            Cmd = 24
	BPF_MAP_LOOKUP_AND_DELETE_BATCH Cmd = 25
	BPF_MAP_UPDATE_BATCH            Cmd = 26
	BPF_MAP_DELETE_BATCH            Cmd = 27
	BPF_LINK_CREATE                 Cmd = 28
	BPF_LINK_UPDATE                 Cmd = 29
	BPF_LINK_GET_FD_BY_ID           Cmd = 30
	BPF_LINK_GET_NEXT_ID            Cmd = 31
	BPF_ENABLE_STATS                Cmd = 32
	BPF_ITER_CREATE                 Cmd = 33
	BPF_LINK_DETACH                 Cmd = 34
	BPF_PROG_BIND_MAP               Cmd = 35
)

type FunctionId uint32

const (
	BPF_FUNC_unspec                         FunctionId = 0
	BPF_FUNC_map_lookup_elem                FunctionId = 1
	BPF_FUNC_map_update_elem                FunctionId = 2
	BPF_FUNC_map_delete_elem                FunctionId = 3
	BPF_FUNC_probe_read                     FunctionId = 4
	BPF_FUNC_ktime_get_ns                   FunctionId = 5
	BPF_FUNC_trace_printk                   FunctionId = 6
	BPF_FUNC_get_prandom_u32                FunctionId = 7
	BPF_FUNC_get_smp_processor_id           FunctionId = 8
	BPF_FUNC_skb_store_bytes                FunctionId = 9
	BPF_FUNC_l3_csum_replace                FunctionId = 10
	BPF_FUNC_l4_csum_replace                FunctionId = 11
	BPF_FUNC_tail_call                      FunctionId = 12
	BPF_FUNC_clone_redirect                 FunctionId = 13
	BPF_FUNC_get_current_pid_tgid           FunctionId = 14
	BPF_FUNC_get_current_uid_gid            FunctionId = 15
	BPF_FUNC_get_current_comm               FunctionId = 16
	BPF_FUNC_get_cgroup_classid             FunctionId = 17
	BPF_FUNC_skb_vlan_push                  FunctionId = 18
	BPF_FUNC_skb_vlan_pop                   FunctionId = 19
	BPF_FUNC_skb_get_tunnel_key             FunctionId = 20
	BPF_FUNC_skb_set_tunnel_key             FunctionId = 21
	BPF_FUNC_perf_event_read                FunctionId = 22
	BPF_FUNC_redirect                       FunctionId = 23
	BPF_FUNC_get_route_realm                FunctionId = 24
	BPF_FUNC_perf_event_output              FunctionId = 25
	BPF_FUNC_skb_load_bytes                 FunctionId = 26
	BPF_FUNC_get_stackid                    FunctionId = 27
	BPF_FUNC_csum_diff                      FunctionId = 28
	BPF_FUNC_skb_get_tunnel_opt             FunctionId = 29
	BPF_FUNC_skb_set_tunnel_opt             FunctionId = 30
	BPF_FUNC_skb_change_proto               FunctionId = 31
	BPF_FUNC_skb_change_type                FunctionId = 32
	BPF_FUNC_skb_under_cgroup               FunctionId = 33
	BPF_FUNC_get_hash_recalc                FunctionId = 34
	BPF_FUNC_get_current_task               FunctionId = 35
	BPF_FUNC_probe_write_user               FunctionId = 36
	BPF_FUNC_current_task_under_cgroup      FunctionId = 37
	BPF_FUNC_skb_change_tail                FunctionId = 38
	BPF_FUNC_skb_pull_data                  FunctionId = 39
	BPF_FUNC_csum_update                    FunctionId = 40
	BPF_FUNC_set_hash_invalid               FunctionId = 41
	BPF_FUNC_get_numa_node_id               FunctionId = 42
	BPF_FUNC_skb_change_head                FunctionId = 43
	BPF_FUNC_xdp_adjust_head                FunctionId = 44
	BPF_FUNC_probe_read_str                 FunctionId = 45
	BPF_FUNC_get_socket_cookie              FunctionId = 46
	BPF_FUNC_get_socket_uid                 FunctionId = 47
	BPF_FUNC_set_hash                       FunctionId = 48
	BPF_FUNC_setsockopt                     FunctionId = 49
	BPF_FUNC_skb_adjust_room                FunctionId = 50
	BPF_FUNC_redirect_map                   FunctionId = 51
	BPF_FUNC_sk_redirect_map                FunctionId = 52
	BPF_FUNC_sock_map_update                FunctionId = 53
	BPF_FUNC_xdp_adjust_meta                FunctionId = 54
	BPF_FUNC_perf_event_read_value          FunctionId = 55
	BPF_FUNC_perf_prog_read_value           FunctionId = 56
	BPF_FUNC_getsockopt                     FunctionId = 57
	BPF_FUNC_override_return                FunctionId = 58
	BPF_FUNC_sock_ops_cb_flags_set          FunctionId = 59
	BPF_FUNC_msg_redirect_map               FunctionId = 60
	BPF_FUNC_msg_apply_bytes                FunctionId = 61
	BPF_FUNC_msg_cork_bytes                 FunctionId = 62
	BPF_FUNC_msg_pull_data                  FunctionId = 63
	BPF_FUNC_bind                           FunctionId = 64
	BPF_FUNC_xdp_adjust_tail                FunctionId = 65
	BPF_FUNC_skb_get_xfrm_state             FunctionId = 66
	BPF_FUNC_get_stack                      FunctionId = 67
	BPF_FUNC_skb_load_bytes_relative        FunctionId = 68
	BPF_FUNC_fib_lookup                     FunctionId = 69
	BPF_FUNC_sock_hash_update               FunctionId = 70
	BPF_FUNC_msg_redirect_hash              FunctionId = 71
	BPF_FUNC_sk_redirect_hash               FunctionId = 72
	BPF_FUNC_lwt_push_encap                 FunctionId = 73
	BPF_FUNC_lwt_seg6_store_bytes           FunctionId = 74
	BPF_FUNC_lwt_seg6_adjust_srh            FunctionId = 75
	BPF_FUNC_lwt_seg6_action                FunctionId = 76
	BPF_FUNC_rc_repeat                      FunctionId = 77
	BPF_FUNC_rc_keydown                     FunctionId = 78
	BPF_FUNC_skb_cgroup_id                  FunctionId = 79
	BPF_FUNC_get_current_cgroup_id          FunctionId = 80
	BPF_FUNC_get_local_storage              FunctionId = 81
	BPF_FUNC_sk_select_reuseport            FunctionId = 82
	BPF_FUNC_skb_ancestor_cgroup_id         FunctionId = 83
	BPF_FUNC_sk_lookup_tcp                  FunctionId = 84
	BPF_FUNC_sk_lookup_udp                  FunctionId = 85
	BPF_FUNC_sk_release                     FunctionId = 86
	BPF_FUNC_map_push_elem                  FunctionId = 87
	BPF_FUNC_map_pop_elem                   FunctionId = 88
	BPF_FUNC_map_peek_elem                  FunctionId = 89
	BPF_FUNC_msg_push_data                  FunctionId = 90
	BPF_FUNC_msg_pop_data                   FunctionId = 91
	BPF_FUNC_rc_pointer_rel                 FunctionId = 92
	BPF_FUNC_spin_lock                      FunctionId = 93
	BPF_FUNC_spin_unlock                    FunctionId = 94
	BPF_FUNC_sk_fullsock                    FunctionId = 95
	BPF_FUNC_tcp_sock                       FunctionId = 96
	BPF_FUNC_skb_ecn_set_ce                 FunctionId = 97
	BPF_FUNC_get_listener_sock              FunctionId = 98
	BPF_FUNC_skc_lookup_tcp                 FunctionId = 99
	BPF_FUNC_tcp_check_syncookie            FunctionId = 100
	BPF_FUNC_sysctl_get_name                FunctionId = 101
	BPF_FUNC_sysctl_get_current_value       FunctionId = 102
	BPF_FUNC_sysctl_get_new_value           FunctionId = 103
	BPF_FUNC_sysctl_set_new_value           FunctionId = 104
	BPF_FUNC_strtol                         FunctionId = 105
	BPF_FUNC_strtoul                        FunctionId = 106
	BPF_FUNC_sk_storage_get                 FunctionId = 107
	BPF_FUNC_sk_storage_delete              FunctionId = 108
	BPF_FUNC_send_signal                    FunctionId = 109
	BPF_FUNC_tcp_gen_syncookie              FunctionId = 110
	BPF_FUNC_skb_output                     FunctionId = 111
	BPF_FUNC_probe_read_user                FunctionId = 112
	BPF_FUNC_probe_read_kernel              FunctionId = 113
	BPF_FUNC_probe_read_user_str            FunctionId = 114
	BPF_FUNC_probe_read_kernel_str          FunctionId = 115
	BPF_FUNC_tcp_send_ack                   FunctionId = 116
	BPF_FUNC_send_signal_thread             FunctionId = 117
	BPF_FUNC_jiffies64                      FunctionId = 118
	BPF_FUNC_read_branch_records            FunctionId = 119
	BPF_FUNC_get_ns_current_pid_tgid        FunctionId = 120
	BPF_FUNC_xdp_output                     FunctionId = 121
	BPF_FUNC_get_netns_cookie               FunctionId = 122
	BPF_FUNC_get_current_ancestor_cgroup_id FunctionId = 123
	BPF_FUNC_sk_assign                      FunctionId = 124
	BPF_FUNC_ktime_get_boot_ns              FunctionId = 125
	BPF_FUNC_seq_printf                     FunctionId = 126
	BPF_FUNC_seq_write                      FunctionId = 127
	BPF_FUNC_sk_cgroup_id                   FunctionId = 128
	BPF_FUNC_sk_ancestor_cgroup_id          FunctionId = 129
	BPF_FUNC_ringbuf_output                 FunctionId = 130
	BPF_FUNC_ringbuf_reserve                FunctionId = 131
	BPF_FUNC_ringbuf_submit                 FunctionId = 132
	BPF_FUNC_ringbuf_discard                FunctionId = 133
	BPF_FUNC_ringbuf_query                  FunctionId = 134
	BPF_FUNC_csum_level                     FunctionId = 135
	BPF_FUNC_skc_to_tcp6_sock               FunctionId = 136
	BPF_FUNC_skc_to_tcp_sock                FunctionId = 137
	BPF_FUNC_skc_to_tcp_timewait_sock       FunctionId = 138
	BPF_FUNC_skc_to_tcp_request_sock        FunctionId = 139
	BPF_FUNC_skc_to_udp6_sock               FunctionId = 140
	BPF_FUNC_get_task_stack                 FunctionId = 141
	BPF_FUNC_load_hdr_opt                   FunctionId = 142
	BPF_FUNC_store_hdr_opt                  FunctionId = 143
	BPF_FUNC_reserve_hdr_opt                FunctionId = 144
	BPF_FUNC_inode_storage_get              FunctionId = 145
	BPF_FUNC_inode_storage_delete           FunctionId = 146
	BPF_FUNC_d_path                         FunctionId = 147
	BPF_FUNC_copy_from_user                 FunctionId = 148
	BPF_FUNC_snprintf_btf                   FunctionId = 149
	BPF_FUNC_seq_printf_btf                 FunctionId = 150
	BPF_FUNC_skb_cgroup_classid             FunctionId = 151
	BPF_FUNC_redirect_neigh                 FunctionId = 152
	BPF_FUNC_per_cpu_ptr                    FunctionId = 153
	BPF_FUNC_this_cpu_ptr                   FunctionId = 154
	BPF_FUNC_redirect_peer                  FunctionId = 155
	BPF_FUNC_task_storage_get               FunctionId = 156
	BPF_FUNC_task_storage_delete            FunctionId = 157
	BPF_FUNC_get_current_task_btf           FunctionId = 158
	BPF_FUNC_bprm_opts_set                  FunctionId = 159
	BPF_FUNC_ktime_get_coarse_ns            FunctionId = 160
	BPF_FUNC_ima_inode_hash                 FunctionId = 161
	BPF_FUNC_sock_from_file                 FunctionId = 162
	BPF_FUNC_check_mtu                      FunctionId = 163
	BPF_FUNC_for_each_map_elem              FunctionId = 164
	BPF_FUNC_snprintf                       FunctionId = 165
	BPF_FUNC_sys_bpf                        FunctionId = 166
	BPF_FUNC_btf_find_by_name_kind          FunctionId = 167
	BPF_FUNC_sys_close                      FunctionId = 168
	BPF_FUNC_timer_init                     FunctionId = 169
	BPF_FUNC_timer_set_callback             FunctionId = 170
	BPF_FUNC_timer_start                    FunctionId = 171
	BPF_FUNC_timer_cancel                   FunctionId = 172
	BPF_FUNC_get_func_ip                    FunctionId = 173
	BPF_FUNC_get_attach_cookie              FunctionId = 174
	BPF_FUNC_task_pt_regs                   FunctionId = 175
	BPF_FUNC_get_branch_snapshot            FunctionId = 176
	BPF_FUNC_trace_vprintk                  FunctionId = 177
	BPF_FUNC_skc_to_unix_sock               FunctionId = 178
	BPF_FUNC_kallsyms_lookup_name           FunctionId = 179
	BPF_FUNC_find_vma                       FunctionId = 180
	BPF_FUNC_loop                           FunctionId = 181
	BPF_FUNC_strncmp                        FunctionId = 182
	BPF_FUNC_get_func_arg                   FunctionId = 183
	BPF_FUNC_get_func_ret                   FunctionId = 184
	BPF_FUNC_get_func_arg_cnt               FunctionId = 185
	BPF_FUNC_get_retval                     FunctionId = 186
	BPF_FUNC_set_retval                     FunctionId = 187
	BPF_FUNC_xdp_get_buff_len               FunctionId = 188
	BPF_FUNC_xdp_load_bytes                 FunctionId = 189
	BPF_FUNC_xdp_store_bytes                FunctionId = 190
	BPF_FUNC_copy_from_user_task            FunctionId = 191
	BPF_FUNC_skb_set_tstamp                 FunctionId = 192
	BPF_FUNC_ima_file_hash                  FunctionId = 193
	BPF_FUNC_kptr_xchg                      FunctionId = 194
	BPF_FUNC_map_lookup_percpu_elem         FunctionId = 195
	BPF_FUNC_skc_to_mptcp_sock              FunctionId = 196
	BPF_FUNC_dynptr_from_mem                FunctionId = 197
	BPF_FUNC_ringbuf_reserve_dynptr         FunctionId = 198
	BPF_FUNC_ringbuf_submit_dynptr          FunctionId = 199
	BPF_FUNC_ringbuf_discard_dynptr         FunctionId = 200
	BPF_FUNC_dynptr_read                    FunctionId = 201
	BPF_FUNC_dynptr_write                   FunctionId = 202
	BPF_FUNC_dynptr_data                    FunctionId = 203
	BPF_FUNC_tcp_raw_gen_syncookie_ipv4     FunctionId = 204
	BPF_FUNC_tcp_raw_gen_syncookie_ipv6     FunctionId = 205
	BPF_FUNC_tcp_raw_check_syncookie_ipv4   FunctionId = 206
	BPF_FUNC_tcp_raw_check_syncookie_ipv6   FunctionId = 207
	BPF_FUNC_ktime_get_tai_ns               FunctionId = 208
	BPF_FUNC_user_ringbuf_drain             FunctionId = 209
	__BPF_FUNC_MAX_ID                       FunctionId = 210
)

type HdrStartOff uint32

const (
	BPF_HDR_START_MAC HdrStartOff = 0
	BPF_HDR_START_NET HdrStartOff = 1
)

type LinkType uint32

const (
	BPF_LINK_TYPE_UNSPEC         LinkType = 0
	BPF_LINK_TYPE_RAW_TRACEPOINT LinkType = 1
	BPF_LINK_TYPE_TRACING        LinkType = 2
	BPF_LINK_TYPE_CGROUP         LinkType = 3
	BPF_LINK_TYPE_ITER           LinkType = 4
	BPF_LINK_TYPE_NETNS          LinkType = 5
	BPF_LINK_TYPE_XDP            LinkType = 6
	BPF_LINK_TYPE_PERF_EVENT     LinkType = 7
	BPF_LINK_TYPE_KPROBE_MULTI   LinkType = 8
	BPF_LINK_TYPE_STRUCT_OPS     LinkType = 9
	MAX_BPF_LINK_TYPE            LinkType = 10
)

type MapType uint32

const (
	BPF_MAP_TYPE_UNSPEC                MapType = 0
	BPF_MAP_TYPE_HASH                  MapType = 1
	BPF_MAP_TYPE_ARRAY                 MapType = 2
	BPF_MAP_TYPE_PROG_ARRAY            MapType = 3
	BPF_MAP_TYPE_PERF_EVENT_ARRAY      MapType = 4
	BPF_MAP_TYPE_PERCPU_HASH           MapType = 5
	BPF_MAP_TYPE_PERCPU_ARRAY          MapType = 6
	BPF_MAP_TYPE_STACK_TRACE           MapType = 7
	BPF_MAP_TYPE_CGROUP_ARRAY          MapType = 8
	BPF_MAP_TYPE_LRU_HASH              MapType = 9
	BPF_MAP_TYPE_LRU_PERCPU_HASH       MapType = 10
	BPF_MAP_TYPE_LPM_TRIE              MapType = 11
	BPF_MAP_TYPE_ARRAY_OF_MAPS         MapType = 12
	BPF_MAP_TYPE_HASH_OF_MAPS          MapType = 13
	BPF_MAP_TYPE_DEVMAP                MapType = 14
	BPF_MAP_TYPE_SOCKMAP               MapType = 15
	BPF_MAP_TYPE_CPUMAP                MapType = 16
	BPF_MAP_TYPE_XSKMAP                MapType = 17
	BPF_MAP_TYPE_SOCKHASH              MapType = 18
	BPF_MAP_TYPE_CGROUP_STORAGE        MapType = 19
	BPF_MAP_TYPE_REUSEPORT_SOCKARRAY   MapType = 20
	BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE MapType = 21
	BPF_MAP_TYPE_QUEUE                 MapType = 22
	BPF_MAP_TYPE_STACK                 MapType = 23
	BPF_MAP_TYPE_SK_STORAGE            MapType = 24
	BPF_MAP_TYPE_DEVMAP_HASH           MapType = 25
	BPF_MAP_TYPE_STRUCT_OPS            MapType = 26
	BPF_MAP_TYPE_RINGBUF               MapType = 27
	BPF_MAP_TYPE_INODE_STORAGE         MapType = 28
	BPF_MAP_TYPE_TASK_STORAGE          MapType = 29
	BPF_MAP_TYPE_BLOOM_FILTER          MapType = 30
	BPF_MAP_TYPE_USER_RINGBUF          MapType = 31
)

type ProgType uint32

const (
	BPF_PROG_TYPE_UNSPEC                  ProgType = 0
	BPF_PROG_TYPE_SOCKET_FILTER           ProgType = 1
	BPF_PROG_TYPE_KPROBE                  ProgType = 2
	BPF_PROG_TYPE_SCHED_CLS               ProgType = 3
	BPF_PROG_TYPE_SCHED_ACT               ProgType = 4
	BPF_PROG_TYPE_TRACEPOINT              ProgType = 5
	BPF_PROG_TYPE_XDP                     ProgType = 6
	BPF_PROG_TYPE_PERF_EVENT              ProgType = 7
	BPF_PROG_TYPE_CGROUP_SKB              ProgType = 8
	BPF_PROG_TYPE_CGROUP_SOCK             ProgType = 9
	BPF_PROG_TYPE_LWT_IN                  ProgType = 10
	BPF_PROG_TYPE_LWT_OUT                 ProgType = 11
	BPF_PROG_TYPE_LWT_XMIT                ProgType = 12
	BPF_PROG_TYPE_SOCK_OPS                ProgType = 13
	BPF_PROG_TYPE_SK_SKB                  ProgType = 14
	BPF_PROG_TYPE_CGROUP_DEVICE           ProgType = 15
	BPF_PROG_TYPE_SK_MSG                  ProgType = 16
	BPF_PROG_TYPE_RAW_TRACEPOINT          ProgType = 17
	BPF_PROG_TYPE_CGROUP_SOCK_ADDR        ProgType = 18
	BPF_PROG_TYPE_LWT_SEG6LOCAL           ProgType = 19
	BPF_PROG_TYPE_LIRC_MODE2              ProgType = 20
	BPF_PROG_TYPE_SK_REUSEPORT            ProgType = 21
	BPF_PROG_TYPE_FLOW_DISSECTOR          ProgType = 22
	BPF_PROG_TYPE_CGROUP_SYSCTL           ProgType = 23
	BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE ProgType = 24
	BPF_PROG_TYPE_CGROUP_SOCKOPT          ProgType = 25
	BPF_PROG_TYPE_TRACING                 ProgType = 26
	BPF_PROG_TYPE_STRUCT_OPS              ProgType = 27
	BPF_PROG_TYPE_EXT                     ProgType = 28
	BPF_PROG_TYPE_LSM                     ProgType = 29
	BPF_PROG_TYPE_SK_LOOKUP               ProgType = 30
	BPF_PROG_TYPE_SYSCALL                 ProgType = 31
)

type RetCode uint32

const (
	BPF_OK                      RetCode = 0
	BPF_DROP                    RetCode = 2
	BPF_REDIRECT                RetCode = 7
	BPF_LWT_REROUTE             RetCode = 128
	BPF_FLOW_DISSECTOR_CONTINUE RetCode = 129
)

type SkAction uint32

const (
	SK_DROP SkAction = 0
	SK_PASS SkAction = 1
)

type StackBuildIdStatus uint32

const (
	BPF_STACK_BUILD_ID_EMPTY StackBuildIdStatus = 0
	BPF_STACK_BUILD_ID_VALID StackBuildIdStatus = 1
	BPF_STACK_BUILD_ID_IP    StackBuildIdStatus = 2
)

type StatsType uint32

const (
	BPF_STATS_RUN_TIME StatsType = 0
)

type XdpAction uint32

const (
	XDP_ABORTED  XdpAction = 0
	XDP_DROP     XdpAction = 1
	XDP_PASS     XdpAction = 2
	XDP_TX       XdpAction = 3
	XDP_REDIRECT XdpAction = 4
)

type BtfInfo struct {
	Btf       Pointer
	BtfSize   uint32
	Id        BTFID
	Name      Pointer
	NameLen   uint32
	KernelBtf uint32
}

type FuncInfo struct {
	InsnOff uint32
	TypeId  uint32
}

type LineInfo struct {
	InsnOff     uint32
	FileNameOff uint32
	LineOff     uint32
	LineCol     uint32
}

type LinkInfo struct {
	Type   LinkType
	Id     LinkID
	ProgId uint32
	_      [4]byte
	Extra  [32]uint8
}

type MapInfo struct {
	Type                  uint32
	Id                    uint32
	KeySize               uint32
	ValueSize             uint32
	MaxEntries            uint32
	MapFlags              MapFlags
	Name                  ObjName
	Ifindex               uint32
	BtfVmlinuxValueTypeId TypeID
	NetnsDev              uint64
	NetnsIno              uint64
	BtfId                 uint32
	BtfKeyTypeId          TypeID
	BtfValueTypeId        TypeID
	_                     [4]byte
	MapExtra              uint64
}

type ProgInfo struct {
	Type                 uint32
	Id                   uint32
	Tag                  [8]uint8
	JitedProgLen         uint32
	XlatedProgLen        uint32
	JitedProgInsns       uint64
	XlatedProgInsns      Pointer
	LoadTime             uint64
	CreatedByUid         uint32
	NrMapIds             uint32
	MapIds               Pointer
	Name                 ObjName
	Ifindex              uint32
	_                    [4]byte /* unsupported bitfield */
	NetnsDev             uint64
	NetnsIno             uint64
	NrJitedKsyms         uint32
	NrJitedFuncLens      uint32
	JitedKsyms           uint64
	JitedFuncLens        uint64
	BtfId                BTFID
	FuncInfoRecSize      uint32
	FuncInfo             Pointer
	NrFuncInfo           uint32
	NrLineInfo           uint32
	LineInfo             Pointer
	JitedLineInfo        uint64
	NrJitedLineInfo      uint32
	LineInfoRecSize      uint32
	JitedLineInfoRecSize uint32
	NrProgTags           uint32
	ProgTags             uint64
	RunTimeNs            uint64
	RunCnt               uint64
	RecursionMisses      uint64
	VerifiedInsns        uint32
	AttachBtfObjId       BTFID
	AttachBtfId          TypeID
	_                    [4]byte
}

type SkLookup struct {
	Cookie         uint64
	Family         uint32
	Protocol       uint32
	RemoteIp4      [4]uint8
	RemoteIp6      [16]uint8
	RemotePort     uint16
	_              [2]byte
	LocalIp4       [4]uint8
	LocalIp6       [16]uint8
	LocalPort      uint32
	IngressIfindex uint32
	_              [4]byte
}

type XdpMd struct {
	Data           uint32
	DataEnd        uint32
	DataMeta       uint32
	IngressIfindex uint32
	RxQueueIndex   uint32
	EgressIfindex  uint32
}

type BtfGetFdByIdAttr struct{ Id uint32 }

func BtfGetFdById(attr *BtfGetFdByIdAttr) (*FD, error) {
	fd, err := BPF(BPF_BTF_GET_FD_BY_ID, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type BtfGetNextIdAttr struct {
	Id     BTFID
	NextId BTFID
}

func BtfGetNextId(attr *BtfGetNextIdAttr) error {
	_, err := BPF(BPF_BTF_GET_NEXT_ID, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type BtfLoadAttr struct {
	Btf         Pointer
	BtfLogBuf   Pointer
	BtfSize     uint32
	BtfLogSize  uint32
	BtfLogLevel uint32
	_           [4]byte
}

func BtfLoad(attr *BtfLoadAttr) (*FD, error) {
	fd, err := BPF(BPF_BTF_LOAD, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type EnableStatsAttr struct{ Type uint32 }

func EnableStats(attr *EnableStatsAttr) (*FD, error) {
	fd, err := BPF(BPF_ENABLE_STATS, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type IterCreateAttr struct {
	LinkFd uint32
	Flags  uint32
}

func IterCreate(attr *IterCreateAttr) (*FD, error) {
	fd, err := BPF(BPF_ITER_CREATE, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type LinkCreateAttr struct {
	ProgFd      uint32
	TargetFd    uint32
	AttachType  AttachType
	Flags       uint32
	TargetBtfId TypeID
	_           [28]byte
}

func LinkCreate(attr *LinkCreateAttr) (*FD, error) {
	fd, err := BPF(BPF_LINK_CREATE, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type LinkCreateIterAttr struct {
	ProgFd      uint32
	TargetFd    uint32
	AttachType  AttachType
	Flags       uint32
	IterInfo    Pointer
	IterInfoLen uint32
	_           [20]byte
}

func LinkCreateIter(attr *LinkCreateIterAttr) (*FD, error) {
	fd, err := BPF(BPF_LINK_CREATE, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type LinkCreateKprobeMultiAttr struct {
	ProgFd           uint32
	TargetFd         uint32
	AttachType       AttachType
	Flags            uint32
	KprobeMultiFlags uint32
	Count            uint32
	Syms             Pointer
	Addrs            Pointer
	Cookies          Pointer
}

func LinkCreateKprobeMulti(attr *LinkCreateKprobeMultiAttr) (*FD, error) {
	fd, err := BPF(BPF_LINK_CREATE, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type LinkCreatePerfEventAttr struct {
	ProgFd     uint32
	TargetFd   uint32
	AttachType AttachType
	Flags      uint32
	BpfCookie  uint64
	_          [24]byte
}

func LinkCreatePerfEvent(attr *LinkCreatePerfEventAttr) (*FD, error) {
	fd, err := BPF(BPF_LINK_CREATE, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type LinkCreateTracingAttr struct {
	ProgFd      uint32
	TargetFd    uint32
	AttachType  AttachType
	Flags       uint32
	TargetBtfId BTFID
	_           [4]byte
	Cookie      uint64
	_           [16]byte
}

func LinkCreateTracing(attr *LinkCreateTracingAttr) (*FD, error) {
	fd, err := BPF(BPF_LINK_CREATE, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type LinkUpdateAttr struct {
	LinkFd    uint32
	NewProgFd uint32
	Flags     uint32
	OldProgFd uint32
}

func LinkUpdate(attr *LinkUpdateAttr) error {
	_, err := BPF(BPF_LINK_UPDATE, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapCreateAttr struct {
	MapType               MapType
	KeySize               uint32
	ValueSize             uint32
	MaxEntries            uint32
	MapFlags              MapFlags
	InnerMapFd            uint32
	NumaNode              uint32
	MapName               ObjName
	MapIfindex            uint32
	BtfFd                 uint32
	BtfKeyTypeId          TypeID
	BtfValueTypeId        TypeID
	BtfVmlinuxValueTypeId TypeID
	MapExtra              uint64
}

func MapCreate(attr *MapCreateAttr) (*FD, error) {
	fd, err := BPF(BPF_MAP_CREATE, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type MapDeleteBatchAttr struct {
	InBatch   Pointer
	OutBatch  Pointer
	Keys      Pointer
	Values    Pointer
	Count     uint32
	MapFd     uint32
	ElemFlags uint64
	Flags     uint64
}

func MapDeleteBatch(attr *MapDeleteBatchAttr) error {
	_, err := BPF(BPF_MAP_DELETE_BATCH, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapDeleteElemAttr struct {
	MapFd uint32
	_     [4]byte
	Key   Pointer
	Value Pointer
	Flags uint64
}

func MapDeleteElem(attr *MapDeleteElemAttr) error {
	_, err := BPF(BPF_MAP_DELETE_ELEM, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapFreezeAttr struct{ MapFd uint32 }

func MapFreeze(attr *MapFreezeAttr) error {
	_, err := BPF(BPF_MAP_FREEZE, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapGetFdByIdAttr struct{ Id uint32 }

func MapGetFdById(attr *MapGetFdByIdAttr) (*FD, error) {
	fd, err := BPF(BPF_MAP_GET_FD_BY_ID, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type MapGetNextIdAttr struct {
	Id     uint32
	NextId uint32
}

func MapGetNextId(attr *MapGetNextIdAttr) error {
	_, err := BPF(BPF_MAP_GET_NEXT_ID, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapGetNextKeyAttr struct {
	MapFd   uint32
	_       [4]byte
	Key     Pointer
	NextKey Pointer
}

func MapGetNextKey(attr *MapGetNextKeyAttr) error {
	_, err := BPF(BPF_MAP_GET_NEXT_KEY, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapLookupAndDeleteBatchAttr struct {
	InBatch   Pointer
	OutBatch  Pointer
	Keys      Pointer
	Values    Pointer
	Count     uint32
	MapFd     uint32
	ElemFlags uint64
	Flags     uint64
}

func MapLookupAndDeleteBatch(attr *MapLookupAndDeleteBatchAttr) error {
	_, err := BPF(BPF_MAP_LOOKUP_AND_DELETE_BATCH, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapLookupAndDeleteElemAttr struct {
	MapFd uint32
	_     [4]byte
	Key   Pointer
	Value Pointer
	Flags uint64
}

func MapLookupAndDeleteElem(attr *MapLookupAndDeleteElemAttr) error {
	_, err := BPF(BPF_MAP_LOOKUP_AND_DELETE_ELEM, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapLookupBatchAttr struct {
	InBatch   Pointer
	OutBatch  Pointer
	Keys      Pointer
	Values    Pointer
	Count     uint32
	MapFd     uint32
	ElemFlags uint64
	Flags     uint64
}

func MapLookupBatch(attr *MapLookupBatchAttr) error {
	_, err := BPF(BPF_MAP_LOOKUP_BATCH, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapLookupElemAttr struct {
	MapFd uint32
	_     [4]byte
	Key   Pointer
	Value Pointer
	Flags uint64
}

func MapLookupElem(attr *MapLookupElemAttr) error {
	_, err := BPF(BPF_MAP_LOOKUP_ELEM, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapUpdateBatchAttr struct {
	InBatch   Pointer
	OutBatch  Pointer
	Keys      Pointer
	Values    Pointer
	Count     uint32
	MapFd     uint32
	ElemFlags uint64
	Flags     uint64
}

func MapUpdateBatch(attr *MapUpdateBatchAttr) error {
	_, err := BPF(BPF_MAP_UPDATE_BATCH, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type MapUpdateElemAttr struct {
	MapFd uint32
	_     [4]byte
	Key   Pointer
	Value Pointer
	Flags uint64
}

func MapUpdateElem(attr *MapUpdateElemAttr) error {
	_, err := BPF(BPF_MAP_UPDATE_ELEM, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type ObjGetAttr struct {
	Pathname  Pointer
	BpfFd     uint32
	FileFlags uint32
}

func ObjGet(attr *ObjGetAttr) (*FD, error) {
	fd, err := BPF(BPF_OBJ_GET, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type ObjGetInfoByFdAttr struct {
	BpfFd   uint32
	InfoLen uint32
	Info    Pointer
}

func ObjGetInfoByFd(attr *ObjGetInfoByFdAttr) error {
	_, err := BPF(BPF_OBJ_GET_INFO_BY_FD, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type ObjPinAttr struct {
	Pathname  Pointer
	BpfFd     uint32
	FileFlags uint32
}

func ObjPin(attr *ObjPinAttr) error {
	_, err := BPF(BPF_OBJ_PIN, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type ProgAttachAttr struct {
	TargetFd     uint32
	AttachBpfFd  uint32
	AttachType   uint32
	AttachFlags  uint32
	ReplaceBpfFd uint32
}

func ProgAttach(attr *ProgAttachAttr) error {
	_, err := BPF(BPF_PROG_ATTACH, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type ProgBindMapAttr struct {
	ProgFd uint32
	MapFd  uint32
	Flags  uint32
}

func ProgBindMap(attr *ProgBindMapAttr) error {
	_, err := BPF(BPF_PROG_BIND_MAP, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type ProgDetachAttr struct {
	TargetFd    uint32
	AttachBpfFd uint32
	AttachType  uint32
}

func ProgDetach(attr *ProgDetachAttr) error {
	_, err := BPF(BPF_PROG_DETACH, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type ProgGetFdByIdAttr struct{ Id uint32 }

func ProgGetFdById(attr *ProgGetFdByIdAttr) (*FD, error) {
	fd, err := BPF(BPF_PROG_GET_FD_BY_ID, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type ProgGetNextIdAttr struct {
	Id     uint32
	NextId uint32
}

func ProgGetNextId(attr *ProgGetNextIdAttr) error {
	_, err := BPF(BPF_PROG_GET_NEXT_ID, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type ProgLoadAttr struct {
	ProgType           ProgType
	InsnCnt            uint32
	Insns              Pointer
	License            Pointer
	LogLevel           LogLevel
	LogSize            uint32
	LogBuf             Pointer
	KernVersion        uint32
	ProgFlags          uint32
	ProgName           ObjName
	ProgIfindex        uint32
	ExpectedAttachType AttachType
	ProgBtfFd          uint32
	FuncInfoRecSize    uint32
	FuncInfo           Pointer
	FuncInfoCnt        uint32
	LineInfoRecSize    uint32
	LineInfo           Pointer
	LineInfoCnt        uint32
	AttachBtfId        TypeID
	AttachBtfObjFd     uint32
	CoreReloCnt        uint32
	FdArray            Pointer
	CoreRelos          Pointer
	CoreReloRecSize    uint32
	_                  [4]byte
}

func ProgLoad(attr *ProgLoadAttr) (*FD, error) {
	fd, err := BPF(BPF_PROG_LOAD, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type ProgQueryAttr struct {
	TargetFd        uint32
	AttachType      AttachType
	QueryFlags      uint32
	AttachFlags     uint32
	ProgIds         Pointer
	ProgCount       uint32
	_               [4]byte
	ProgAttachFlags uint64
}

func ProgQuery(attr *ProgQueryAttr) error {
	_, err := BPF(BPF_PROG_QUERY, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type ProgRunAttr struct {
	ProgFd      uint32
	Retval      uint32
	DataSizeIn  uint32
	DataSizeOut uint32
	DataIn      Pointer
	DataOut     Pointer
	Repeat      uint32
	Duration    uint32
	CtxSizeIn   uint32
	CtxSizeOut  uint32
	CtxIn       Pointer
	CtxOut      Pointer
	Flags       uint32
	Cpu         uint32
	BatchSize   uint32
	_           [4]byte
}

func ProgRun(attr *ProgRunAttr) error {
	_, err := BPF(BPF_PROG_TEST_RUN, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	return err
}

type RawTracepointOpenAttr struct {
	Name   Pointer
	ProgFd uint32
	_      [4]byte
}

func RawTracepointOpen(attr *RawTracepointOpenAttr) (*FD, error) {
	fd, err := BPF(BPF_RAW_TRACEPOINT_OPEN, unsafe.Pointer(attr), unsafe.Sizeof(*attr))
	if err != nil {
		return nil, err
	}
	return NewFD(int(fd))
}

type CgroupLinkInfo struct {
	CgroupId   uint64
	AttachType AttachType
	_          [4]byte
}

type IterLinkInfo struct {
	TargetName    Pointer
	TargetNameLen uint32
}

type NetNsLinkInfo struct {
	NetnsIno   uint32
	AttachType AttachType
}

type RawTracepointLinkInfo struct {
	TpName    Pointer
	TpNameLen uint32
	_         [4]byte
}

type TracingLinkInfo struct {
	AttachType  AttachType
	TargetObjId uint32
	TargetBtfId TypeID
}

type XDPLinkInfo struct{ Ifindex uint32 }
