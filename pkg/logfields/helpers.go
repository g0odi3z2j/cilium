// Copyright 2017 Authors of Cilium
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

package logfields

import (
	"fmt"

	// Initialize logrus package with debug log level if needed
	_ "github.com/cilium/cilium/pkg/debugdetection"
)

// Repr formats an object with the Printf %+v formatter
func Repr(s interface{}) string {
	return fmt.Sprintf("%+v", s)
}
