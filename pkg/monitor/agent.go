// Copyright 2018 Authors of Cilium
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

package monitor

import (
	"encoding/json"
	"fmt"

	"github.com/cilium/cilium/pkg/policy/api"
)

// AgentNotify is a notification from the agent
type AgentNotify struct {
	Type AgentNotification
	Text string
}

// AgentNotification specifies the type of agent notification
type AgentNotification uint32

const (
	AgentNotifyUnspec AgentNotification = iota
	AgentNotifyGeneric
	AgentNotifyStart
	AgentNotifyEndpointRegenerateSuccess
	AgentNotifyEndpointRegenerateFail
	AgentNotifyPolicyUpdated
	AgentNotifyPolicyDeleted
)

var notifyTable = map[AgentNotification]string{
	AgentNotifyUnspec:                    "unspecified",
	AgentNotifyGeneric:                   "Message",
	AgentNotifyStart:                     "Cilium agent started",
	AgentNotifyEndpointRegenerateSuccess: "Endpoint regenerated",
	AgentNotifyEndpointRegenerateFail:    "Failed endpoint regeneration",
	AgentNotifyPolicyUpdated:             "Policy updated",
	AgentNotifyPolicyDeleted:             "Policy deleted",
}

func resolveAgentType(t AgentNotification) string {
	if n, ok := notifyTable[t]; ok {
		return n
	}

	return fmt.Sprintf("%d", t)
}

// DumpInfo dumps an agent notification
func (n *AgentNotify) DumpInfo() {
	fmt.Printf(">> %s: %s\n", resolveAgentType(n.Type), n.Text)
}

type PolicyUpdateNotification struct {
	Labels    []string `json:"labels,omitempty"`
	Revision  uint64   `json:"revision,omitempty"`
	RuleCount int      `json:"rule_count"`
}

func PolicyUpdateRepr(rules api.Rules, revision uint64) (string, error) {
	labels := make([]string, 0, len(rules))
	for _, r := range rules {
		labels = append(labels, r.Labels.GetModel()...)
	}

	notification := PolicyUpdateNotification{
		Labels:    labels,
		Revision:  revision,
		RuleCount: len(rules),
	}

	resp, err := json.Marshal(notification)

	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func PolicyDeleteRepr(deleted int, labels []string, revision uint64) (string, error) {
	notification := PolicyUpdateNotification{
		Labels:    labels,
		Revision:  revision,
		RuleCount: deleted,
	}
	resp, err := json.Marshal(notification)

	if err != nil {
		return "", err
	}

	return string(resp), nil
}
