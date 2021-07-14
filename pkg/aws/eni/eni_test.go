// SPDX-License-Identifier: Apache-2.0
// Copyright 2019 Authors of Cilium

// +build !privileged_tests

package eni

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type ENISuite struct{}

var _ = check.Suite(&ENISuite{})
