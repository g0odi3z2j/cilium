// Copyright 2019-2020 Authors of Cilium
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

package ipam

import (
	"context"
	"fmt"

	"github.com/cilium/cilium/pkg/azure/types"
	"github.com/cilium/cilium/pkg/ipam"
	ipamTypes "github.com/cilium/cilium/pkg/ipam/types"
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/cilium/cilium/pkg/math"

	"github.com/sirupsen/logrus"
)

// Node represents a node representing an Azure instance
type Node struct {
	// node contains the general purpose fields of a node
	node *ipam.Node

	// manager is the Azure node manager responsible for this node
	manager *InstancesManager
}

// UpdatedNode is called when an update to the CiliumNode is received.
func (n *Node) UpdatedNode(obj *v2.CiliumNode) {
}

// PopulateStatusFields fills in the status field of the CiliumNode custom
// resource with Azure specific information
func (n *Node) PopulateStatusFields(k8sObj *v2.CiliumNode) {
	k8sObj.Status.Azure.Interfaces = []types.AzureInterface{}
	n.manager.instances.ForeachInterface(n.node.InstanceID(), func(instanceID, interfaceID string, interfaceObj ipamTypes.InterfaceRevision) error {
		iface, ok := interfaceObj.Resource.(*types.AzureInterface)
		if ok {
			k8sObj.Status.Azure.Interfaces = append(k8sObj.Status.Azure.Interfaces, *(iface.DeepCopy()))
		}
		return nil
	})
}

// PrepareIPRelease prepares the release of IPs
func (n *Node) PrepareIPRelease(excessIPs int, scopedLog *logrus.Entry) *ipam.ReleaseAction {
	return &ipam.ReleaseAction{}
}

// ReleaseIPs performs the IP release operation
func (n *Node) ReleaseIPs(ctx context.Context, r *ipam.ReleaseAction) error {
	return fmt.Errorf("not implemented")
}

// PrepareIPAllocation returns the number of IPs that can be allocated/created.
func (n *Node) PrepareIPAllocation(scopedLog *logrus.Entry) (a *ipam.AllocationAction, err error) {
	a = &ipam.AllocationAction{}
	err = n.manager.instances.ForeachInterface(n.node.InstanceID(), func(instanceID, interfaceID string, interfaceObj ipamTypes.InterfaceRevision) error {
		iface, ok := interfaceObj.Resource.(*types.AzureInterface)
		if !ok {
			return fmt.Errorf("invalid interface object")
		}

		scopedLog.WithFields(logrus.Fields{
			"id":           iface.ID,
			"numAddresses": len(iface.Addresses),
		}).Debug("Considering interface for allocation")

		availableOnInterface := math.IntMax(types.InterfaceAddressLimit-len(iface.Addresses), 0)
		if availableOnInterface <= 0 {
			return nil
		}

		a.AvailableInterfaces++

		if a.InterfaceID == "" {
			scopedLog.WithFields(logrus.Fields{
				"id":                   iface.ID,
				"availableOnInterface": availableOnInterface,
			}).Debug("Interface has IPs available")

			preferredPoolIDs := []ipamTypes.PoolID{}
			for _, address := range iface.Addresses {
				if address.Subnet != "" {
					preferredPoolIDs = append(preferredPoolIDs, ipamTypes.PoolID(address.Subnet))
				}
			}

			poolID, available := n.manager.getAllocator().FirstPoolWithAvailableQuota(preferredPoolIDs)
			if poolID != ipamTypes.PoolNotExists {
				scopedLog.WithFields(logrus.Fields{
					"subnetID":           poolID,
					"availableAddresses": available,
				}).Debug("Subnet has IPs available")

				a.InterfaceID = iface.ID
				a.PoolID = poolID
				a.AvailableForAllocation = math.IntMin(available, availableOnInterface)
			}
		}
		return nil
	})

	return
}

// AllocateIPs performs the Azure IP allocation operation
func (n *Node) AllocateIPs(ctx context.Context, a *ipam.AllocationAction) error {
	ips, err := n.manager.getAllocator().AllocateMany(a.PoolID, a.AvailableForAllocation)
	if err != nil {
		return err
	}

	err = n.manager.api.AssignPrivateIpAddresses(ctx, string(a.PoolID), a.InterfaceID, ips)
	if err != nil {
		n.manager.getAllocator().ReleaseMany(a.PoolID, ips)
		return err
	}

	return nil
}

// CreateInterface is called to create a new interface. This operation is
// currently not supported on Azure.
func (n *Node) CreateInterface(ctx context.Context, allocation *ipam.AllocationAction, scopedLog *logrus.Entry) (int, string, error) {
	return 0, "", fmt.Errorf("not implemented")
}

// ResyncInterfacesAndIPs is called to retrieve and interfaces and IPs as known
// to the Azure API and return them
func (n *Node) ResyncInterfacesAndIPs(ctx context.Context, scopedLog *logrus.Entry) (ipamTypes.AllocationMap, error) {
	if n.node.InstanceID() == "" {
		return nil, nil
	}

	available := ipamTypes.AllocationMap{}
	n.manager.instances.ForeachAddress(n.node.InstanceID(), func(instanceID, interfaceID, ip, poolID string, addressObj ipamTypes.Address) error {
		address, ok := addressObj.(types.AzureAddress)
		if !ok {
			log.WithField("ip", ip).Warning("Not an Azure address object, ignoring IP")
			return nil
		}

		if address.State == types.StateSucceeded {
			available[address.IP] = ipamTypes.AllocationIP{Resource: interfaceID}
		} else {
			log.WithFields(logrus.Fields{
				"ip":    address.IP,
				"state": address.State,
			}).Warning("Ignoring potentially available IP due to non-successful state")
		}
		return nil
	})

	return available, nil
}
