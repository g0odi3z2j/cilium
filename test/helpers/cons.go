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

package helpers

import (
	"fmt"
	"time"
)

var (
	timeout     time.Duration = 300 //WithTimeout helper translate it to seconds
	basePath                  = "/vagrant/"
	networkName               = "cilium-net"
)

//GetFilePath returns the absolute path of the provided filename
func GetFilePath(filename string) string {
	return fmt.Sprintf("%s%s", basePath, filename)
}
