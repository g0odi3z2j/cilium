package daemon

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/noironetworks/cilium-net/bpf/policymap"
	"github.com/noironetworks/cilium-net/common"
	"github.com/noironetworks/cilium-net/common/types"

	"github.com/op/go-logging"
)

// findNode returns node and its parent or an error
func (d *Daemon) findNode(path string) (*types.PolicyNode, *types.PolicyNode, error) {
	var parent *types.PolicyNode

	if strings.HasPrefix(path, common.GlobalLabelPrefix) == false {
		return nil, nil, fmt.Errorf("Invalid path %s: must start with %s", path, common.GlobalLabelPrefix)
	}

	newPath := strings.Replace(path, common.GlobalLabelPrefix, "", 1)
	if newPath == "" {
		return d.policyTree.Root, nil, nil
	}

	current := d.policyTree.Root
	parent = nil

	for _, nodeName := range strings.Split(newPath, ".") {
		if nodeName == "" {
			continue
		}
		if child, ok := current.Children[nodeName]; ok {
			parent = current
			current = child
		} else {
			return nil, nil, fmt.Errorf("Unable to find child %s of node %s in path %s", nodeName, current.Name, path)
		}
	}

	return current, parent, nil
}

func (d *Daemon) evaluateConsumerSource(c *types.Consumable, ctx *types.SearchContext, srcID uint32) error {
	ctx.From = nil

	// Check if we have the source security context in our local
	// consumable cache
	srcConsumable := types.LookupConsumable(srcID)
	if srcConsumable != nil {
		ctx.From = srcConsumable.LabelList
	}

	// No cache entry or labels not available, do full lookup of labels
	// via KV store
	if ctx.From == nil {
		lbls, err := d.GetLabels(srcID)
		if err != nil {
			return err
		}

		// ID is not associated with anything, skip...
		if lbls == nil {
			return nil
		}

		ctx.From = make([]types.Label, len(lbls.Labels))

		idx := 0
		for k, v := range lbls.Labels {
			ctx.From[idx] = types.Label{Key: k, Value: v.Value, Source: v.Source}
			idx++
		}
	}

	decision := d.policyCanConsume(ctx)
	// Only accept rules get stored
	if decision == types.ACCEPT {
		c.AllowConsumerAndReverse(srcID)
	}

	return nil
}

// Must be called with endpointsMU held
func (d *Daemon) regenerateConsumable(c *types.Consumable) error {
	// Containers without a security label are not accessible
	if c.ID == 0 {
		log.Fatalf("Impossible: SecLabel == 0 when generating endpoint consumers")
		return nil
	}

	// Skip if policy for this consumable is already valid
	if c.Iteration == d.cacheIteration {
		log.Debugf("Policy for %d is already calculated, reusing...", c.ID)
		return nil
	}

	maxID, err := d.GetMaxID()
	if err != nil {
		return err
	}

	ctx := types.SearchContext{
		To: c.LabelList,
	}

	if d.conf.EnableTracing {
		ctx.Trace = types.TRACE_ENABLED
	}

	d.policyTreeMU.Lock()
	defer d.policyTreeMU.Unlock()

	// Mark all entries unused by denying them
	for k, _ := range c.Consumers {
		c.Consumers[k].DeletionMark = true
	}

	// Check access from reserved consumables first
	for _, id := range d.reservedConsumables {
		if err := d.evaluateConsumerSource(c, &ctx, id.ID); err != nil {
			// This should never really happen
			// FIXME: clear policy because it is inconsistent
			break
		}
	}

	// Iterate over all possible assigned search contexts
	idx := common.FirstFreeID
	for idx < maxID {
		if err := d.evaluateConsumerSource(c, &ctx, idx); err != nil {
			// FIXME: clear policy because it is inconsistent
			break
		}
		idx++
	}

	// Garbage collect all unused entries
	for _, val := range c.Consumers {
		if val.DeletionMark {
			val.DeletionMark = false
			c.BanConsumer(val.ID)
		}
	}

	// Result is valid until cache iteration advances
	c.Iteration = d.cacheIteration

	log.Debugf("New policy (iteration %d) for consumable %d: %+v\n", c.Iteration, c.ID, c.Consumers)

	return nil
}

func (d *Daemon) invalidateCache() {
	d.cacheIteration++
	if d.cacheIteration == 0 {
		d.cacheIteration = 1
	}
}

func (d *Daemon) regenerateEndpoint(e *types.Endpoint) error {
	if e.Consumable != nil {
		return d.regenerateConsumable(e.Consumable)
	} else {
		return nil
	}
}

func (d *Daemon) triggerPolicyUpdates(added []uint32) {
	d.endpointsMU.Lock()
	defer d.endpointsMU.Unlock()

	if len(added) == 0 {
		log.Debugf("Full policy recalculation triggered")
		d.invalidateCache()
	} else {
		log.Debugf("Partial policy recalculation triggered: %d\n", added)
		// FIXME: Invalidate only cache that is affected
		d.invalidateCache()
	}

	for _, ep := range d.endpoints {
		d.regenerateEndpoint(ep)
	}
}

// policyCanConsume calculates if the ctx allows the consumer to be consumed.
func (d *Daemon) policyCanConsume(ctx *types.SearchContext) types.ConsumableDecision {
	return d.policyTree.Allows(ctx)
}

// PolicyCanConsume calculates if the ctx allows the consumer to be consumed. This public
// function returns a SearchContextReply with the consumable decision and the tracing log
// if ctx.Trace was set.
func (d *Daemon) PolicyCanConsume(ctx *types.SearchContext) (*types.SearchContextReply, error) {
	buffer := new(bytes.Buffer)
	if ctx.Trace != types.TRACE_DISABLED {
		ctx.Logging = logging.NewLogBackend(buffer, "", 0)
	}
	scr := types.SearchContextReply{}
	scr.Decision = d.policyTree.Allows(ctx)
	if ctx.Trace != types.TRACE_DISABLED {
		scr.Logging = buffer.Bytes()
	}
	return &scr, nil
}

func (d *Daemon) PolicyAdd(path string, node *types.PolicyNode) error {
	if node.Name == "" {
		node.Name = path
	}

	log.Debugf("Policy Add Request: %+v", node)

	d.policyTreeMU.Lock()
	parentNode, parent, err := d.findNode(path)
	if err != nil {
		d.policyTreeMU.Unlock()
		return err
	}
	if parent == nil {
		d.policyTree.Root = node
	} else {
		if parent == nil {
			d.policyTree.Root = node
		} else {
			parentNode.Children[node.Name] = node
		}
	}
	d.policyTreeMU.Unlock()

	d.triggerPolicyUpdates([]uint32{})

	return nil
}

// PolicyDelete deletes the policy set in path from the policy tree.
func (d *Daemon) PolicyDelete(path string) error {
	log.Debugf("Policy Delete Request: %s", path)

	d.policyTreeMU.Lock()
	node, parent, err := d.findNode(path)
	if err != nil {
		d.policyTreeMU.Unlock()
		return err
	}
	if parent == nil {
		d.policyTree.Root = &types.PolicyNode{}
	} else {
		if parent == nil {
			d.policyTree.Root = &types.PolicyNode{}
		} else {
			delete(parent.Children, node.Name)
		}
	}
	d.policyTreeMU.Unlock()

	d.triggerPolicyUpdates([]uint32{})

	return nil
}

// PolicyGet returns the policy of the given path.
func (d *Daemon) PolicyGet(path string) (*types.PolicyNode, error) {
	log.Debugf("Policy Get Request: %s", path)
	node, _, err := d.findNode(path)
	return node, err
}

func (d *Daemon) PolicyInit() error {
	for k, v := range types.ResDec {
		lbl := types.SecCtxLabel{
			ID:       uint32(v),
			RefCount: 1,
			Labels:   map[string]*types.Label{},
		}
		policyMapPath := fmt.Sprintf("%sreserved_%d", common.PolicyMapPath, uint32(v))

		lbl.Labels[k] = &types.Label{Key: k, Source: common.ReservedLabelSource}

		policyMap, err := policymap.OpenMap(policyMapPath)
		if err != nil {
			return fmt.Errorf("Could not create policy BPF map '%s': %s", policyMapPath, err)
		}

		if c := types.GetConsumable(uint32(v), &lbl); c == nil {
			return fmt.Errorf("Unable to initialize consumable for %v", lbl)
		} else {
			d.reservedConsumables = append(d.reservedConsumables, c)
			c.AddMap(policyMap)
		}
	}

	return nil
}
