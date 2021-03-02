// Copyright 2021 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import "golang.org/x/sys/unix"

func sysconf(name int) (int64, error) {
	return unix.Sysconf(name)
}
