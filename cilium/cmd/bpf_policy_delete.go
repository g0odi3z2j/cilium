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

package cmd

import (
	"github.com/spf13/cobra"
)

// bpf_policy_deleteCmd represents the bpf_policy_delete command
var bpf_policy_deleteCmd = &cobra.Command{
	Use:    "delete <endpoint id> <identity>",
	Short:  "Delete a policy entry",
	PreRun: requireEndpointID,
	Run: func(cmd *cobra.Command, args []string) {
		updatePolicyKey(cmd, args, false)
	},
}

func init() {
	bpf_policyCmd.AddCommand(bpf_policy_deleteCmd)
}
