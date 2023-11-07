// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fqdn

import (
	"context"
	"net"
	"net/netip"
	"regexp"
	"sync"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/cilium/cilium/api/v1/models"
	"github.com/cilium/cilium/pkg/controller"
	"github.com/cilium/cilium/pkg/identity"
	"github.com/cilium/cilium/pkg/ip"
	ipcacheTypes "github.com/cilium/cilium/pkg/ipcache/types"
	"github.com/cilium/cilium/pkg/lock"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/policy/api"
	"github.com/cilium/cilium/pkg/source"
	"github.com/cilium/cilium/pkg/time"
)

// NameManager maintains state DNS names, via FQDNSelector or exact match for
// polling, need to be tracked. It is the main structure which relates the FQDN
// subsystem to the policy subsystem for plumbing the relation between a DNS
// name and the corresponding IPs which have been returned via DNS lookups.
// When DNS updates are given to a NameManager it update cached selectors as
// required via UpdateSelectors.
// DNS information is cached, respecting TTL.
type NameManager struct {
	lock.RWMutex

	// config is a copy from when this instance was initialized.
	// It is read-only once set
	config Config

	// allSelectors contains all FQDNSelectors which are present in all policy. We
	// use these selectors to map selectors --> IPs.
	allSelectors map[api.FQDNSelector]*regexp.Regexp

	// cache is a private copy of the pointer from config.
	cache *DNSCache

	bootstrapCompleted bool

	manager *controller.Manager
}

// GetModel returns the API model of the NameManager.
func (n *NameManager) GetModel() *models.NameManager {
	n.RWMutex.RLock()
	defer n.RWMutex.RUnlock()

	allSelectors := make([]*models.SelectorEntry, 0, len(n.allSelectors))
	for fqdnSel, regex := range n.allSelectors {
		pair := &models.SelectorEntry{
			SelectorString: fqdnSel.String(),
			RegexString:    regex.String(),
		}
		allSelectors = append(allSelectors, pair)
	}

	return &models.NameManager{
		FQDNPolicySelectors: allSelectors,
	}
}

// Lock must be held during any calls to RegisterForIdentityUpdatesLocked or
// UnregisterForIdentityUpdatesLocked.
func (n *NameManager) Lock() {
	n.RWMutex.Lock()
}

// Unlock must be called after calls to RegisterForIdentityUpdatesLocked or
// UnregisterForIdentityUpdatesLocked are done.
func (n *NameManager) Unlock() {
	n.RWMutex.Unlock()
}

// RegisterForIdentityUpdatesLocked exposes this FQDNSelector so that identities
// for IPs contained in a DNS response that matches said selector can be
// propagated back to the SelectorCache via `UpdateFQDNSelector`. All DNS names
// contained within the NameManager's cache are iterated over to see if they match
// the FQDNSelector. All IPs which correspond to the DNS names which match this
// Selector will be returned as CIDR identities, as other DNS Names which have
// already been resolved may match this FQDNSelector.
func (n *NameManager) RegisterForIdentityUpdatesLocked(selector api.FQDNSelector) {
	_, exists := n.allSelectors[selector]
	if exists {
		log.WithField("fqdnSelector", selector).Warning("FQDNSelector was already registered for updates, returning without any identities")
		return
	}

	// This error should never occur since the FQDNSelector has already been
	// validated, but account for it for good measure.
	regex, err := selector.ToRegex()
	if err != nil {
		log.WithError(err).WithField("fqdnSelector", selector).Error("FQDNSelector did not compile to valid regex")
		return
	}

	n.allSelectors[selector] = regex
}

// UnregisterForIdentityUpdatesLocked removes this FQDNSelector from the set of
// FQDNSelectors which are being tracked by the NameManager. No more updates
// for IPs which correspond to said selector are propagated.
func (n *NameManager) UnregisterForIdentityUpdatesLocked(selector api.FQDNSelector) {
	delete(n.allSelectors, selector)
}

// NewNameManager creates an initialized NameManager.
// When config.Cache is nil, the global fqdn.DefaultDNSCache is used.
func NewNameManager(config Config) *NameManager {

	if config.Cache == nil {
		config.Cache = NewDNSCache(0)
	}

	if config.UpdateSelectors == nil {
		config.UpdateSelectors = func(ctx context.Context, selectorsWithIPs map[api.FQDNSelector][]net.IP, selectorsWithoutIPs []api.FQDNSelector) (*sync.WaitGroup, []*identity.Identity, map[netip.Prefix]*identity.Identity, error) {
			return &sync.WaitGroup{}, nil, nil, nil
		}
	}
	if config.GetEndpointsDNSInfo == nil {
		config.GetEndpointsDNSInfo = func() []EndpointDNSInfo {
			return nil
		}
	}

	return &NameManager{
		config:       config,
		allSelectors: make(map[api.FQDNSelector]*regexp.Regexp),
		cache:        config.Cache,
		manager:      controller.NewManager(),
	}
}

// UpdateGenerateDNS inserts the new DNS information into the cache. If the IPs
// have changed for a name they will be reflected in updatedDNSIPs.
func (n *NameManager) UpdateGenerateDNS(ctx context.Context, lookupTime time.Time, updatedDNSIPs map[string]*DNSIPRecords) (wg *sync.WaitGroup, usedIdentities []*identity.Identity, newlyAllocatedIdentities map[netip.Prefix]*identity.Identity, err error) {
	n.RWMutex.Lock()
	defer n.RWMutex.Unlock()

	// Update IPs in n
	fqdnSelectorsToUpdate, updatedDNSNames := n.updateDNSIPs(lookupTime, updatedDNSIPs)
	for dnsName, IPs := range updatedDNSNames {
		log.WithFields(logrus.Fields{
			"matchName":             dnsName,
			"IPs":                   IPs,
			"fqdnSelectorsToUpdate": fqdnSelectorsToUpdate,
		}).Debug("Updated FQDN with new IPs")
	}

	namesMissingIPs, selectorIPMapping := n.MapSelectorsToIPsLocked(fqdnSelectorsToUpdate)
	if len(namesMissingIPs) != 0 {
		log.WithField(logfields.DNSName, namesMissingIPs).
			Debug("No IPs to insert when generating DNS name selected by ToFQDN rule")
	}

	return n.config.UpdateSelectors(ctx, selectorIPMapping, namesMissingIPs)
}

// ForceGenerateDNS unconditionally regenerates all rules that refer to DNS
// names in namesToRegen. These names are FQDNs and toFQDNs.matchPatterns or
// matchNames that match them will cause these rules to regenerate.
// Note: This is used only when DNS entries are cleaned up, not when new results
// are ingested.
func (n *NameManager) ForceGenerateDNS(ctx context.Context, namesToRegen []string) (wg *sync.WaitGroup, err error) {
	n.RWMutex.Lock()
	defer n.RWMutex.Unlock()

	affectedFQDNSels := make(sets.Set[api.FQDNSelector], 0)
	for _, dnsName := range namesToRegen {
		for fqdnSel, fqdnRegEx := range n.allSelectors {
			if fqdnRegEx.MatchString(dnsName) {
				affectedFQDNSels.Insert(fqdnSel)
			}
		}
	}

	namesMissingIPs, selectorIPMapping := n.MapSelectorsToIPsLocked(affectedFQDNSels)
	if len(namesMissingIPs) != 0 {
		log.WithField(logfields.DNSName, namesMissingIPs).
			Debug("No IPs to insert when generating DNS name selected by ToFQDN rule")
	}

	// Emit the new rules.
	// Ignore newly allocated IDs (3rd result) as this is only used for deletes.
	wg, _, _, err = n.config.UpdateSelectors(ctx, selectorIPMapping, namesMissingIPs)
	return wg, err
}

func (n *NameManager) CompleteBootstrap() {
	n.Lock()
	n.bootstrapCompleted = true
	n.Unlock()
}

// updateDNSIPs updates the IPs for each DNS name in updatedDNSIPs.
// It returns:
// affectedSelectors: a set of all FQDNSelectors which match DNS Names whose
// corresponding set of IPs has changed.
// updatedNames: a map of DNS names to all the valid IPs we store for each.
func (n *NameManager) updateDNSIPs(lookupTime time.Time, updatedDNSIPs map[string]*DNSIPRecords) (affectedSelectors sets.Set[api.FQDNSelector], updatedNames map[string][]net.IP) {
	updatedNames = make(map[string][]net.IP, len(updatedDNSIPs))
	affectedSelectors = make(sets.Set[api.FQDNSelector], len(updatedDNSIPs))

	for dnsName, lookupIPs := range updatedDNSIPs {
		addrs := ip.MustAddrsFromIPs(lookupIPs.IPs)
		updated := n.updateIPsForName(lookupTime, dnsName, addrs, lookupIPs.TTL)

		// The IPs didn't change. No more to be done for this dnsName
		if !updated && n.bootstrapCompleted {
			log.WithFields(logrus.Fields{
				"dnsName":   dnsName,
				"lookupIPs": lookupIPs,
			}).Debug("FQDN: IPs didn't change for DNS name")
			continue
		}

		// Upsert in to the ipcache metadata layer
		n.upsertMetadata(addrs)

		// record the IPs that were different
		updatedNames[dnsName] = lookupIPs.IPs

		// accumulate the new selectors affected by new IPs
		if len(n.allSelectors) == 0 {
			log.WithFields(logrus.Fields{
				"dnsName":   dnsName,
				"lookupIPs": lookupIPs,
			}).Debug("FQDN: No selectors registered for updates")
		}
		for fqdnSel, fqdnRegex := range n.allSelectors {
			matches := fqdnRegex.MatchString(dnsName)
			if matches {
				affectedSelectors.Insert(fqdnSel)
			}
		}
	}

	return affectedSelectors, updatedNames
}

// updateIPsName will update the IPs for dnsName. It always retains a copy of
// newIPs.
// updated is true when the new IPs differ from the old IPs
func (n *NameManager) updateIPsForName(lookupTime time.Time, dnsName string, newIPs []netip.Addr, ttl int) (updated bool) {
	cacheIPs := n.cache.Lookup(dnsName)

	if n.config.MinTTL > ttl {
		ttl = n.config.MinTTL
	}

	n.cache.Update(lookupTime, dnsName, newIPs, ttl)
	sortedNewIPs := n.cache.Lookup(dnsName) // DNSCache returns IPs sorted

	// The 0 checks below account for an unlike race condition where this
	// function is called with already expired data and if other cache data
	// from before also expired.
	return (len(cacheIPs) == 0 && len(sortedNewIPs) == 0) || !ip.SortedIPListsAreEqual(sortedNewIPs, cacheIPs)
}

var ipcacheResource = ipcacheTypes.NewResourceID(ipcacheTypes.ResourceKindDaemon, "", "fqdn-name-manager")

// upsertMetadata adds an entry in the ipcache metadata layer for the set of IPs.
// Returns the ipcache queue revision (to pass to .WaitForRevision()).
func (n *NameManager) upsertMetadata(ips []netip.Addr) uint64 {
	prefixes := make([]netip.Prefix, 0, len(ips))
	for _, ip := range ips {
		prefixes = append(prefixes, netip.PrefixFrom(ip, ip.BitLen()))
	}
	return n.config.IPCache.UpsertPrefixes(prefixes, source.Generated, ipcacheResource)
}

// maybeRemoveMetadata removes the ipcache metadata from every IP in maybeRemoved,
// as long as that IP is not still in the dns cache.
func (n *NameManager) maybeRemoveMetadata(maybeRemoved sets.Set[netip.Addr]) {
	// Need to take an RLock here so that no DNS updates are processed.
	// Otherwise, we might accidentally remove an IP that is newly inserted.
	n.RWMutex.RLock()
	defer n.RWMutex.RUnlock()

	n.cache.RLock()
	prefixes := make([]netip.Prefix, 0, len(maybeRemoved))
	for ip := range maybeRemoved {
		if !n.cache.ipExistsLocked(ip) {
			prefixes = append(prefixes, netip.PrefixFrom(ip, ip.BitLen()))
		}
	}
	n.cache.RUnlock()

	log.WithField(logfields.Prefix, prefixes).Debug("Removing fqdn entry from ipcache metadata layer")
	n.config.IPCache.RemovePrefixes(prefixes, source.Generated, ipcacheResource)
}
