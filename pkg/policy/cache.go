// Copyright 2016-2018 Authors of Cilium
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

import (
	"github.com/cilium/cilium/pkg/identity"
	"github.com/cilium/cilium/pkg/labels"
	"github.com/cilium/cilium/pkg/lock"
	"github.com/cilium/cilium/pkg/logging/logfields"
)

var (
	consumableCache = newConsumableCache()
)

type ConsumableCache struct {
	cacheMU lock.RWMutex // Protects the `cache` map
	cache   map[identity.NumericIdentity]*Consumable
	// List of consumables representing the reserved identities
	reserved []*Consumable
}

// GetConsumableCache returns the consumable cache. The cache is a list of all
// identities which are in use by local endpoints.
func GetConsumableCache() *ConsumableCache {
	return consumableCache
}

func newConsumableCache() *ConsumableCache {
	return &ConsumableCache{
		cache:    map[identity.NumericIdentity]*Consumable{},
		reserved: make([]*Consumable, 0),
	}
}

func (c *ConsumableCache) Lookup(id identity.NumericIdentity) *Consumable {
	c.cacheMU.RLock()
	v, _ := c.cache[id]
	c.cacheMU.RUnlock()
	return v
}

func (c *ConsumableCache) addReserved(elem *Consumable) {
	c.cacheMU.Lock()
	c.reserved = append(c.reserved, elem)
	c.cacheMU.Unlock()
}

// GetReservedIDs returns a slice of NumericIdentity present in the
// ConsumableCache.
func (c *ConsumableCache) GetReservedIDs() []identity.NumericIdentity {
	identities := []identity.NumericIdentity{}
	c.cacheMU.RLock()
	for _, id := range c.reserved {
		identities = append(identities, id.ID)
	}
	c.cacheMU.RUnlock()
	return identities
}

// ResolveIdentityLabels resolves a numeric identity to the identity's labels
// or nil
func ResolveIdentityLabels(id identity.NumericIdentity) labels.LabelArray {
	// Check if we have the source security context in our local
	// consumable cache
	if c := consumableCache.Lookup(id); c != nil {
		return c.LabelArray
	}

	if identity := identity.LookupIdentityByID(id); identity != nil {
		return identity.Labels.ToSlice()
	}

	return nil
}

// InitReserved must be called to initialize the Consumables that represent the reserved
// identities. This is because the reserved identities do not correspond to
// endpoints, and thus must be created explicitly, as opposed to during policy
// calculation, which is done for a specific endpoint when it is regenerated.
func InitReserved() {
	log.Info("Initializing reserved identities")
	for key, val := range identity.ReservedIdentities {
		log.WithField(logfields.Identity, key).Debug("Registering reserved identity")

		identity := identity.NewIdentity(val, labels.Labels{
			key: labels.NewLabel(val.String(), "", labels.LabelSourceReserved),
		})

		c := NewConsumable(val, identity)
		GetConsumableCache().addReserved(c)
	}
}
