package types

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Available privileges for policy nodes to define
type Privilege byte

const (
	ALLOW Privilege = iota
	REQUIRES
	DROP_PRIVILEGES
)

type ConsumableDecision byte

const (
	ACCEPT ConsumableDecision = iota
	DENY
	UNDECIDED
)

func (d *ConsumableDecision) String() string {
	switch *d {
	case ACCEPT:
		return "accept"
		break
	case DENY:
		return "deny"
		break
	case UNDECIDED:
		return "undecided"
		break
	}

	return "unknown"
}

type LabelSelector struct {
	Source string
	Key    string
	Value  string
}

func (l *LabelSelector) String() string {
	if l.Value != "" {
		return fmt.Sprintf("%s=%s", l.Key, l.Value)
	} else {
		return l.Key
	}
}

func (l *LabelSelector) UnmarshalJSON(data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))

	if l == nil {
		return fmt.Errorf("Cannot unmarhshal to nil pointer")
	}

	if len(data) == 0 {
		return fmt.Errorf("Invalid LabelSelector: empty data")
	}

	if bytes.Contains(data, []byte(`"source":`)) {
		var aux struct {
			Source string `json:"source"`
			Key    string `json:"key" binding:"required"`
			Value  string `json:"value"`
		}

		if err := decoder.Decode(&aux); err != nil {
			return fmt.Errorf("Decode of LabelSelector failed: %+v", err)
		}

		if aux.Key == "" {
			return fmt.Errorf("Invalid LabelSelector: must provide a label key")
		}

		l.Source = aux.Source
		l.Key = aux.Key
		l.Value = aux.Value
	} else {
		// This is a short form in which only a string to be interpreted
		// as a cilium label key is provided
		var aux string

		if err := decoder.Decode(&aux); err != nil {
			return fmt.Errorf("Decode of LabelSelector as string failed: %+v", err)
		}

		if aux == "" {
			return fmt.Errorf("Invalid LabelSelector: must provide a label key")
		}

		l.Source = "cilium"
		l.Key = aux
		l.Value = ""
	}

	return nil
}

// FIXME: Write test cases
func (l *LabelSelector) Expand(node *PolicyNode) string {
	return fmt.Sprintf("%s.%s", node.FullName(), l)
}

type SearchContext struct {
	From map[string]string
	To   map[string]string
}

// Base type for all PolicyRule* types
type PolicyRuleBase struct {
	Coverage []LabelSelector `json:"Coverage,omitempty"`
}

type AllowRule struct {
	Inverted bool
	Label    LabelSelector
}

// Allow the following consumers
type PolicyRuleConsumers struct {
	PolicyRuleBase
	Allow []AllowRule `json:"Allow"`
}

func (c *PolicyRuleConsumers) Allows(ctx *SearchContext) ConsumableDecision {
	//decision := UNDECIDED
	//for _, allowedLabel := range c.Allow {
	//	if val, ok := ctx.from[allowedLabel]; ok {
	//		decision = val.Decision
	//	}
	//}

	return UNDECIDED
}

// Any further consumer requires the specified list of
// labels in order to consume
type PolicyRuleRequires struct {
	PolicyRuleBase
	Requires []string `json:"Requires"`
}

type Port struct {
	proto  string `json:"Protocol"`
	number int    `json:"Number"`
}

type PolicyRulePorts struct {
	PolicyRuleBase
	Ports []Port `json:"Ports"`
}

// Do not allow further rules of specified type
type PolicyRuleDropPrivileges struct {
	PolicyRuleBase
	DropPrivileges []Privilege `json:"Drop-privileges"`
}

// Node to define hierarchy of rules
type PolicyNode struct {
	Name     string
	Parent   *PolicyNode            `json:"-"`
	Rules    []interface{}          `json:"Rules,omitempty"`
	Children map[string]*PolicyNode `json:"Children,omitempty"`
}

// Overall policy tree
type PolicyTree struct {
	Root PolicyNode
}

func (pn *PolicyNode) FullName() string {
	if pn.Parent != nil {
		s := fmt.Sprintf("%s.%s", pn.Parent.FullName(), pn.Name)
		fmt.Printf("Building FullName: %s\n", s)
		return s
	}

	fmt.Printf("Building FullName: io.cilium\n")
	return "io.cilium"
}

func (pn *PolicyNode) UnmarshalJSON(data []byte) error {
	var policyNode struct {
		Name     string                 `json:"Name,omitempty"`
		Rules    []*json.RawMessage     `json:"Rules,omitempty"`
		Children map[string]*PolicyNode `json:"Children,omitempty"`
	}

	decoder := json.NewDecoder(bytes.NewReader(data))

	if err := decoder.Decode(&policyNode); err != nil {
		return fmt.Errorf("Decode of PolicyNode failed: %+v", err)
	}

	pn.Name = policyNode.Name
	pn.Children = policyNode.Children

	// Fill out "Name" field of children which ommited it in the JSON
	for k, _ := range policyNode.Children {
		pn.Children[k].Name = k
		pn.Children[k].Parent = pn
	}

	for _, rawMsg := range policyNode.Rules {
		var om map[string]*json.RawMessage

		if err := json.Unmarshal(*rawMsg, &om); err != nil {
			return err
		}

		if _, ok := om["Allow"]; ok {
			var pr_c PolicyRuleConsumers

			if err := json.Unmarshal(*rawMsg, &pr_c); err != nil {
				return err
			}

			pn.Rules = append(pn.Rules, pr_c)
		} else if _, ok := om["Requires"]; ok {
			var pr_r PolicyRuleRequires

			if err := json.Unmarshal(*rawMsg, &pr_r); err != nil {
				return err
			}

			pn.Rules = append(pn.Rules, pr_r)
		} else {
			return fmt.Errorf("Unknown policy rule object: %+v", om)
		}
	}

	return nil
}
