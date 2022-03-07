// Copyright (C) 2016 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//go:build !linux && !dragonfly && !freebsd && !netbsd && !openbsd && !darwin
// +build !linux,!dragonfly,!freebsd,!netbsd,!openbsd,!darwin

package server

import (
	"fmt"
	"net"
)

func setTcpMD5SigSockopt(l *net.TCPListener, address string, key string) error {
	return fmt.Errorf("setting md5 is not supported")
}

func setTcpTTLSockopt(conn *net.TCPConn, ttl int) error {
	return fmt.Errorf("setting ttl is not supported")
}

func setTcpMinTTLSockopt(conn *net.TCPConn, ttl int) error {
	return fmt.Errorf("setting min ttl is not supported")
}
