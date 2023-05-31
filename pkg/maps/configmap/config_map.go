// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package configmap

import (
	"fmt"

	"github.com/cilium/cilium/pkg/bpf"
	"github.com/cilium/cilium/pkg/ebpf"
)

const (
	// MapName name of map used to pin map for datapath
	MapName = "cilium_runtime_config"

	// MaxEntries represents the maximum number of config entries.
	// Initially defined as 256, so that downgrade from a future version having more than one
	// entry works without necessarily resizing the map. Entries not known by the datapath
	// version are simply ignored.
	MaxEntries = 256
)

// Index is the index to the runtime config array.
type Index uint32

// All supported indices in one place.
// Must be in sync with RUNTIME_CONFIG_ enum in bpf/lib/common.h
const (
	UTimeOffset Index = iota
	UsedEntries
	AgentLiveness
)

// Value is the generic datapath runtime config value.
type Value uint64

// String pretty print the Index
func (k *Index) String() string {
	return fmt.Sprintf("%d", uint32(*k))
}

func (k *Index) DeepCopyMapKey() bpf.MapKey { return new(Index) }

// String pretty print the config Value.
func (v *Value) String() string {
	return fmt.Sprintf("%d", uint64(*v))
}

func (v *Value) DeepCopyMapValue() bpf.MapValue { return new(Value) }

// Map provides access to the eBPF map configmap.
type Map interface {
	// Update writes the given uint64 value to the bpf map at the given index.
	Update(index Index, val uint64) error
}

type configMap struct {
	bpfMap *bpf.Map
}

func newConfigMap() *configMap {
	var index Index
	var value Value

	return &configMap{
		bpfMap: bpf.NewMap(MapName,
			ebpf.Array,
			&index,
			&value,
			MaxEntries,
			0,
		),
	}
}

func (m *configMap) init() error {
	if err := m.bpfMap.OpenOrCreate(); err != nil {
		return fmt.Errorf("failed to init bpf map: %w", err)
	}

	return nil
}

func (m *configMap) close() error {
	if err := m.bpfMap.Close(); err != nil {
		return fmt.Errorf("failed to close bpf map: %w", err)
	}

	return nil
}

func (m *configMap) Update(index Index, val uint64) error {
	value := Value(val)
	return m.bpfMap.Update(&index, &value)
}
