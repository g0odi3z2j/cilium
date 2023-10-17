// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package types

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	peerpb "github.com/cilium/cilium/api/v1/peer"
	"github.com/cilium/cilium/pkg/hubble/defaults"
)

func TestFromChangeNotification(t *testing.T) {
	tests := []struct {
		name string
		arg  *peerpb.ChangeNotification
		want *Peer
	}{
		{
			name: "nil",
			arg:  nil,
			want: (*Peer)(nil),
		}, {
			name: "without address",
			arg: &peerpb.ChangeNotification{
				Name: "name",
			},
			want: &Peer{
				Name:    "name",
				Address: nil,
			},
		}, {
			name: "with IPv4 address but without port",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "192.0.2.1",
			},
			want: &Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: defaults.ServerPort,
				},
			},
		}, {
			name: "with an IPv4 address and port",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "192.0.2.1:4000",
			},
			want: &Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 4000,
				},
			},
		}, {
			name: "with an IPv4 address and a bad port",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "192.0.2.1:4x",
			},
			want: &Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: defaults.ServerPort,
				},
			},
		}, {
			name: "with an IPv6 address but without port",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "2001:db8::68",
			},
			want: &Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("2001:db8::68"),
					Port: defaults.ServerPort,
				},
			},
		}, {
			name: "with an IPv6 address and port",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "[2001:db8::68]:4000",
			},
			want: &Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("2001:db8::68"),
					Port: 4000,
				},
			},
		}, {
			name: "with an IPv6 address and a bad port",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "[2001:db8::68]:4x",
			},
			want: &Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("2001:db8::68"),
					Port: defaults.ServerPort,
				},
			},
		}, {
			name: "with a unix domain socket",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "unix:///var/run/hubble.sock",
			},
			want: &Peer{
				Name: "name",
				Address: &net.UnixAddr{
					Name: "unix:///var/run/hubble.sock",
					Net:  "unix",
				},
			},
		}, {
			name: "with a unix domain socket without prefix",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "/var/run/hubble.sock",
			},
			want: &Peer{
				Name: "name",
				Address: &net.UnixAddr{
					Name: "/var/run/hubble.sock",
					Net:  "unix",
				},
			},
		}, {
			name: "with an IPv4 address and port and TLS enabled",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "192.0.2.1:4000",
				Tls:     &peerpb.TLS{},
			},
			want: &Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 4000,
				},
				TLSEnabled: true,
			},
		}, {
			name: "with an IPv4 address and port and TLS enabled with server name",
			arg: &peerpb.ChangeNotification{
				Name:    "name",
				Address: "192.0.2.1:4000",
				Tls: &peerpb.TLS{
					ServerName: "name.default.hubble-grpc.cilium.io",
				},
			},
			want: &Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 4000,
				},
				TLSEnabled:    true,
				TLSServerName: "name.default.hubble-grpc.cilium.io",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromChangeNotification(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
func TestEqual(t *testing.T) {
	tests := []struct {
		name  string
		a     Peer
		b     Peer
		equal bool
	}{
		{
			name:  "empty",
			a:     Peer{},
			b:     Peer{},
			equal: true,
		},
		{
			name: "without address",
			a: Peer{
				Name:    "name",
				Address: nil,
			},
			b: Peer{
				Name:    "name",
				Address: nil,
			},
			equal: true,
		},
		{
			name: "without address different name",
			a: Peer{
				Name:    "foo",
				Address: nil,
			},
			b: Peer{
				Name:    "bar",
				Address: nil,
			},
			equal: false,
		},
		{
			name: "IPv4 no port",
			a: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP: net.ParseIP("192.0.2.1"),
				},
			},
			b: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP: net.ParseIP("192.0.2.1"),
				},
			},
			equal: true,
		},
		{
			name: "different IPv4 no port",
			a: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP: net.ParseIP("192.0.2.1"),
				},
			},
			b: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP: net.ParseIP("192.0.2.2"),
				},
			},
			equal: false,
		},
		{
			name: "IPv4 same port",
			a: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 2222,
				},
			},
			b: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 2222,
				},
			},
			equal: true,
		},
		{
			name: "IPv4 different port",
			a: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 2222,
				},
			},
			b: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 1111,
				},
			},
			equal: false,
		},
		{
			name: "IPv6",
			a: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("2001:db8::68"),
					Port: 4000,
				},
			},
			b: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("2001:db8::68"),
					Port: 4000,
				},
			},
			equal: true,
		},
		{
			name: "different IPv6",
			a: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("2001:db8::68"),
					Port: 4000,
				},
			},
			b: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("2001:db7::68"),
					Port: 4000,
				},
			},
		},
		{
			name: "unix domain socket",
			a: Peer{
				Name: "name",
				Address: &net.UnixAddr{
					Name: "unix:///var/run/hubble.sock",
					Net:  "unix",
				},
			},
			b: Peer{
				Name: "name",
				Address: &net.UnixAddr{
					Name: "unix:///var/run/hubble.sock",
					Net:  "unix",
				},
			},
			equal: true,
		},
		{
			name: "TLS enabled",
			a: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 4000,
				},
				TLSEnabled:    true,
				TLSServerName: "name.default.hubble-grpc.cilium.io",
			},
			b: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 4000,
				},
				TLSEnabled:    true,
				TLSServerName: "name.default.hubble-grpc.cilium.io",
			},
			equal: true,
		},
		{
			name: "TLS enabled different server name",
			a: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 4000,
				},
				TLSEnabled:    true,
				TLSServerName: "name.default.hubble-grpc.cilium.io",
			},
			b: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 4000,
				},
				TLSEnabled:    true,
				TLSServerName: "name.other.hubble-grpc.cilium.io",
			},
			equal: false,
		},
		{
			name: "TLS enabled and not enabled",
			a: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 4000,
				},
				TLSEnabled:    true,
				TLSServerName: "name.default.hubble-grpc.cilium.io",
			},
			b: Peer{
				Name: "name",
				Address: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 4000,
				},
			},
			equal: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aeqb := tt.a.Equal(tt.b)
			beqa := tt.b.Equal(tt.a)

			assert.Equalf(t, aeqb, beqa, "equal isn't commutative (a.Equal(b) is %t, but b.Equal(a) is %t)", aeqb, beqa)
			assert.Equalf(t, tt.equal, aeqb, "exepectd a.Equal(b) to be %t", tt.equal)
		})
	}
}
