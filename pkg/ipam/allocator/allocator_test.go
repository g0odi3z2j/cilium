// SPDX-License-Identifier: Apache-2.0
// Copyright 2020 Authors of Cilium

// +build !privileged_tests

package allocator

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type AllocatorSuite struct{}

var _ = check.Suite(&AllocatorSuite{})
