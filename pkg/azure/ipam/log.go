// SPDX-License-Identifier: Apache-2.0
// Copyright 2020 Authors of Cilium

package ipam

import (
	"github.com/cilium/cilium/pkg/logging"
	"github.com/cilium/cilium/pkg/logging/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "azure")
