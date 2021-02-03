// Copyright 2019 Authors of Hubble
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !privileged_tests

package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_httpPlugin_HelpText(t *testing.T) {
	plugin := &httpPlugin{}
	expected := `http - HTTP metrics
Metrics related to the HTTP protocol

Metrics:
  http_requests_total           - Count of HTTP requests by methods.
  http_responses_total          - Count of HTTP responses by methods and status codes.
  http_request_duration_seconds - Median, 90th and 99th percentile of request duration.

Options:
 sourceContext          := { identity | namespace | pod | pod-short | dns }
 destinationContext     := { identity | namespace | pod | pod-short | dns }`
	assert.Equal(t, expected, plugin.HelpText())
}
