// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package main

/*
#include "proxylib/types.h"
*/
import "C"

import (
	"github.com/sirupsen/logrus"

	"github.com/cilium/cilium/proxylib/libcilium"
)

func init() {
	logrus.Info("proxylib: Initializing library")
}

// OnNewConnection is used to register a new connection of protocol 'proto'.
// Note that the 'origBuf' and replyBuf' type '*[]byte' corresponds to 'InjectBuf' type, but due to
// cgo export restrictions we can't use the go type in the prototype.
//
//export OnNewConnection
func OnNewConnection(instanceId uint64, proto string, connectionId uint64, ingress bool, srcId, dstId uint32, srcAddr, dstAddr, policyName string, origBuf, replyBuf *[]byte) C.FilterResult {
	return C.FilterResult(libcilium.OnNewConnection(instanceId, proto, connectionId, ingress, srcId, dstId, srcAddr, dstAddr, policyName, origBuf, replyBuf))
}

// Each connection is assumed to be called from a single thread, so accessing connection metadata
// does not need protection.
//
// OnData gets all the unparsed data the datapath has received so far. The data is provided to the parser
// associated with the connection, and the parser is expected to find if the data frame contains enough data
// to make a PASS/DROP decision for the whole data frame. Note that the whole data frame need not be received,
// if the decision including the length of the data frame in bytes can be determined based on the beginning of
// the data frame only (e.g., headers including the length of the data frame). The parser returns a decision
// with the number of bytes on which the decision applies. If more data is available, then the parser will be
// called again with the remaining data. Parser needs to return MORE if a decision can't be made with
// the available data, including the minimum number of additional bytes that is needed before the parser is
// called again.
//
// The parser can also inject at arbitrary points in the data stream. This is indecated by an INJECT operation
// with the number of bytes to be injected. The actual bytes to be injected are provided via an Inject()
// callback prior to returning the INJECT operation. The Inject() callback operates on a limited size buffer
// provided by the datapath, and multiple INJECT operations may be needed to inject large amounts of data.
// Since we get the data on one direction at a time, any frames to be injected in the reverse direction
// are placed in the reverse direction buffer, from where the datapath injects the data before calling
// us again for the reverse direction input.
//
//export OnData
func OnData(connectionId uint64, reply, endStream bool, data *[][]byte, filterOps *[][2]int64) C.FilterResult {
	return C.FilterResult(libcilium.OnData(connectionId, reply, endStream, data, filterOps))
}

// Make this more general connection event callback
//
//export Close
func Close(connectionId uint64) {
	libcilium.Close(connectionId)
}

// OpenModule is called before any other APIs.
// Called concurrently by different filter instances.
// Returns a library instance ID that must be passed to all other API calls.
// Calls with the same parameters will return the same instance.
// Zero return value indicates an error.
//
//export OpenModule
func OpenModule(params [][2]string, debug bool) uint64 {
	return libcilium.OpenModule(params, debug)
}

//export CloseModule
func CloseModule(id uint64) {
	libcilium.CloseModule(id)
}

// Must have empty main
func main() {}
