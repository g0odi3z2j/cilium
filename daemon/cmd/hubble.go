// Copyright 2020 Authors of Cilium
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
	"context"
	"fmt"
	"net"
	"time"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	"github.com/cilium/cilium/api/v1/models"
	observerpb "github.com/cilium/cilium/api/v1/observer"
	v1 "github.com/cilium/cilium/pkg/hubble/api/v1"
	"github.com/cilium/cilium/pkg/hubble/listener"
	"github.com/cilium/cilium/pkg/hubble/metrics"
	"github.com/cilium/cilium/pkg/hubble/observer"
	"github.com/cilium/cilium/pkg/hubble/observer/observeroption"
	"github.com/cilium/cilium/pkg/hubble/parser"
	"github.com/cilium/cilium/pkg/hubble/peer"
	"github.com/cilium/cilium/pkg/hubble/server"
	"github.com/cilium/cilium/pkg/hubble/server/serveroption"
	"github.com/cilium/cilium/pkg/identity"
	identitymodel "github.com/cilium/cilium/pkg/identity/model"
	"github.com/cilium/cilium/pkg/ipcache"
	"github.com/cilium/cilium/pkg/loadbalancer"
	"github.com/cilium/cilium/pkg/logging"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/option"

	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
)

func (d *Daemon) getHubbleStatus(ctx context.Context) *models.HubbleStatus {
	if !option.Config.EnableHubble {
		return &models.HubbleStatus{State: models.HubbleStatusStateDisabled}
	}

	if d.hubbleObserver == nil {
		return &models.HubbleStatus{
			State: models.HubbleStatusStateWarning,
			Msg:   "Server not initialized",
		}
	}

	req := &observerpb.ServerStatusRequest{}
	status, err := d.hubbleObserver.ServerStatus(ctx, req)
	if err != nil {
		return &models.HubbleStatus{State: models.HubbleStatusStateFailure, Msg: err.Error()}
	}

	metricsState := models.HubbleStatusMetricsStateDisabled
	if option.Config.HubbleMetricsServer != "" {
		// TODO: The metrics package should be refactored to be able report its actual state
		metricsState = models.HubbleStatusMetricsStateOk
	}

	hubbleStatus := &models.HubbleStatus{
		State: models.StatusStateOk,
		Observer: &models.HubbleStatusObserver{
			CurrentFlows: int64(status.NumFlows),
			MaxFlows:     int64(status.MaxFlows),
			SeenFlows:    int64(status.SeenFlows),
			Uptime:       strfmt.Duration(time.Duration(status.UptimeNs)),
		},
		Metrics: &models.HubbleStatusMetrics{
			State: metricsState,
		},
	}

	return hubbleStatus
}

func (d *Daemon) launchHubble() {
	logger := logging.DefaultLogger.WithField(logfields.LogSubsys, "hubble")
	if !option.Config.EnableHubble {
		logger.Info("Hubble server is disabled")
		return
	}

	payloadParser, err := parser.New(logger, d, d, d, ipcache.IPIdentityCache, d)
	if err != nil {
		logger.WithError(err).Error("Failed to initialize Hubble")
		return
	}
	d.hubbleObserver, err = observer.NewLocalServer(payloadParser, logger,
		observeroption.WithMaxFlows(option.Config.HubbleFlowBufferSize),
		observeroption.WithMonitorBuffer(option.Config.HubbleEventQueueSize),
		observeroption.WithCiliumDaemon(d))
	if err != nil {
		logger.WithError(err).Error("Failed to initialize Hubble")
		return
	}
	go d.hubbleObserver.Start()
	d.monitorAgent.GetMonitor().RegisterNewListener(d.ctx, listener.NewHubbleListener(d.hubbleObserver))

	// configure a local hubble instance that serves more gRPC services
	sockPath := "unix://" + option.Config.HubbleSocketPath
	localSrv, err := server.NewServer(logger,
		serveroption.WithUnixSocketListener(sockPath),
		serveroption.WithHealthService(),
		serveroption.WithObserverService(d.hubbleObserver),
		serveroption.WithPeerService(peer.NewService(d.nodeDiscovery.Manager)),
	)
	if err != nil {
		logger.WithError(err).Error("Failed to initialize local Hubble server")
		return
	}
	logger.WithField("address", sockPath).Info("Starting local Hubble server")
	if err := localSrv.Serve(); err != nil {
		logger.WithError(err).Error("Failed to start local Hubble server")
		return
	}
	go func() {
		<-d.ctx.Done()
		localSrv.Stop()
	}()

	// configure another hubble instance that serve fewer gRPC services
	address := option.Config.HubbleListenAddress
	if address != "" {
		// TODO: remove warning once mutual TLS has been implemented
		logger.WithField("address", address).Warn("Hubble server will be exposing its API insecurely on this address")
		srv, err := server.NewServer(logger,
			serveroption.WithTCPListener(address),
			serveroption.WithHealthService(),
			serveroption.WithObserverService(d.hubbleObserver),
		)
		if err != nil {
			logger.WithError(err).Error("Failed to initialize Hubble server")
			return
		}
		logger.WithField("address", address).Info("Starting Hubble server")
		if err := srv.Serve(); err != nil {
			logger.WithError(err).Error("Failed to start Hubble server")
			return
		}
		go func() {
			<-d.ctx.Done()
			srv.Stop()
		}()
	}

	if option.Config.HubbleMetricsServer != "" {
		logger.WithFields(logrus.Fields{
			"address": option.Config.HubbleMetricsServer,
			"metrics": option.Config.HubbleMetrics,
		}).Info("Starting Hubble Metrics server")
		if err := metrics.EnableMetrics(log, option.Config.HubbleMetricsServer, option.Config.HubbleMetrics); err != nil {
			logger.WithError(err).Warn("Failed to initialize Hubble metrics server")
			return
		}
	}
}

// GetIdentity looks up identity by ID from Cilium's identity cache. Hubble uses the identity info
// to populate source and destination labels of flows.
//
//  - IdentityGetter: https://github.com/cilium/hubble/blob/04ab72591faca62a305ce0715108876167182e04/pkg/parser/getters/getters.go#L40
func (d *Daemon) GetIdentity(securityIdentity uint32) (*models.Identity, error) {
	ident := d.identityAllocator.LookupIdentityByID(context.Background(), identity.NumericIdentity(securityIdentity))
	if ident == nil {
		return nil, fmt.Errorf("identity %d not found", securityIdentity)
	}
	return identitymodel.CreateModel(ident), nil
}

// GetEndpointInfo returns endpoint info for a given IP address. Hubble uses this function to populate
// fields like namespace and pod name for local endpoints.
//
//  - EndpointGetter: https://github.com/cilium/hubble/blob/04ab72591faca62a305ce0715108876167182e04/pkg/parser/getters/getters.go#L34
func (d *Daemon) GetEndpointInfo(ip net.IP) (endpoint v1.EndpointInfo, ok bool) {
	ep := d.endpointManager.LookupIP(ip)
	if ep == nil {
		return nil, false
	}
	return ep, true
}

// GetNamesOf implements DNSGetter.GetNamesOf. It looks up DNS names of a given IP from the
// FQDN cache of an endpoint specified by sourceEpID.
//
//  - DNSGetter: https://github.com/cilium/hubble/blob/04ab72591faca62a305ce0715108876167182e04/pkg/parser/getters/getters.go#L27
func (d *Daemon) GetNamesOf(sourceEpID uint32, ip net.IP) []string {
	ep := d.endpointManager.LookupCiliumID(uint16(sourceEpID))
	if ep == nil {
		return nil
	}
	return ep.DNSHistory.LookupIP(ip)
}

// GetServiceByAddr looks up service by IP/port. Hubble uses this function to annotate flows
// with service information.
//
//  - ServiceGetter: https://github.com/cilium/hubble/blob/04ab72591faca62a305ce0715108876167182e04/pkg/parser/getters/getters.go#L52
func (d *Daemon) GetServiceByAddr(ip net.IP, port uint16) (flowpb.Service, bool) {
	addr := loadbalancer.L3n4Addr{
		IP: ip,
		L4Addr: loadbalancer.L4Addr{
			Port: port,
		},
	}
	namespace, name, ok := d.svc.GetServiceNameByAddr(addr)
	if !ok {
		return flowpb.Service{}, false
	}
	return flowpb.Service{
		Namespace: namespace,
		Name:      name,
	}, true
}
