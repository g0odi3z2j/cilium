/* Dummy configuration for test compilation */

#define NODE_ID 0x10001
#define NODE_MAC { .addr = { 0xde, 0xad, 0xbe, 0xef, 0xc0, 0xde } }
#define ROUTER_IP { 0xbe, 0xef, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x1, 0x0, 0x0 }
#define ENCAP_IFINDEX 1
#define HOST_IFINDEX 1
#define HOST_IP { 0xbe, 0xef, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x2, 0xf, 0xff, 0xff }
#define HOST_IFINDEX_MAC { . addr = { 0xce, 0x72, 0xa7, 0x03, 0x88, 0x56 } }
#define NAT46_SRC_PREFIX { .addr = { 0xbe, 0xef, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0 } }
#define NAT46_DST_PREFIX { .addr = { 0xbe, 0xef, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0 } }
