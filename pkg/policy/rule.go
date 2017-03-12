// Copyright 2016-2017 Authors of Cilium
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

package policy

type PolicyRule interface {
	Resolve(node *Node) error
	SHA256Sum() (string, error)
	CoverageSHA256Sum() (string, error)

	// Must return true if a rule allows merging with other rules within a node.
	// Certain rules are not additive and require strict ordering, such rules
	// may never be merged in a node as merging may occur in undefined order.
	IsMergeable() bool
}
