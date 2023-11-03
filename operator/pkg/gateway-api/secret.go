// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package gateway_api

import (
	"context"

	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"

	"github.com/cilium/cilium/operator/pkg/gateway-api/helpers"
	"github.com/cilium/cilium/pkg/logging/logfields"
)

const (
	owningSecretName      = "io.cilium.gateway/owning-secret-name"
	owningSecretNamespace = "io.cilium.gateway/owning-secret-namespace"
)

// secretSyncer syncs Gateway API secrets to dedicated namespace.
type secretSyncer struct {
	client.Client
	Scheme *runtime.Scheme

	SecretsNamespace string
	controllerName   string
}

func newSecretSyncReconciler(mgr ctrl.Manager, secretsNamespace string) *secretSyncer {
	return &secretSyncer{
		Client:           mgr.GetClient(),
		Scheme:           mgr.GetScheme(),
		SecretsNamespace: secretsNamespace,
		controllerName:   controllerName,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *secretSyncer) SetupWithManager(mgr ctrl.Manager) error {
	hasMatchingControllerFn := hasMatchingController(context.Background(), r.Client, r.controllerName)
	return ctrl.NewControllerManagedBy(mgr).
		// Source Secrets outside of the secrets namespace
		For(&corev1.Secret{}, r.notInSecretsNamespace()).
		// Synced Secrets in the secrets namespace
		Watches(&corev1.Secret{}, enqueueOwningSecretFromLabels(), r.deletedOrChangedInSecretsNamespace()).
		Watches(&gatewayv1.Gateway{}, r.enqueueRequestForGatewayTLS(), builder.WithPredicates(predicate.NewPredicateFuncs(hasMatchingControllerFn))).
		Complete(r)
}

func (r *secretSyncer) notInSecretsNamespace() builder.Predicates {
	return builder.WithPredicates(predicate.NewPredicateFuncs(func(object client.Object) bool {
		return object.GetNamespace() != r.SecretsNamespace
	}))
}

func (r *secretSyncer) enqueueRequestForGatewayTLS() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []reconcile.Request {
		scopedLog := log.WithFields(logrus.Fields{
			logfields.Controller: "secrets",
			logfields.Resource:   obj.GetName(),
		})

		gw, ok := obj.(*gatewayv1.Gateway)
		if !ok {
			return nil
		}

		var reqs []reconcile.Request
		for _, l := range gw.Spec.Listeners {
			if l.TLS == nil {
				continue
			}
			for _, cert := range l.TLS.CertificateRefs {
				if !helpers.IsSecret(cert) {
					continue
				}
				s := types.NamespacedName{
					Namespace: helpers.NamespaceDerefOr(cert.Namespace, gw.Namespace),
					Name:      string(cert.Name),
				}
				reqs = append(reqs, reconcile.Request{NamespacedName: s})
				scopedLog.WithField("secret", s).Debug("Enqueued secret for gateway")
			}
		}
		return reqs
	})
}

func enqueueOwningSecretFromLabels() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(_ context.Context, o client.Object) []reconcile.Request {
		labels := o.GetLabels()

		if labels == nil {
			return nil
		}

		owningSecretNamespace, owningSecretNamespacePresent := labels[owningSecretNamespace]
		owningSecretName, owningSecretNamePresent := labels[owningSecretName]

		if !owningSecretNamespacePresent || !owningSecretNamePresent {
			return nil
		}

		return []reconcile.Request{
			{
				NamespacedName: types.NamespacedName{
					Namespace: owningSecretNamespace,
					Name:      owningSecretName,
				},
			},
		}
	})
}

func (r *secretSyncer) deletedOrChangedInSecretsNamespace() builder.Predicates {
	return builder.WithPredicates(&deletedOrChangedInSecretsNamespaceStruct{
		secretsNamespace: r.SecretsNamespace,
	})
}

var _ predicate.Predicate = &deletedOrChangedInSecretsNamespaceStruct{}

type deletedOrChangedInSecretsNamespaceStruct struct {
	secretsNamespace string
}

func (r *deletedOrChangedInSecretsNamespaceStruct) Create(event.CreateEvent) bool {
	return false
}

func (r *deletedOrChangedInSecretsNamespaceStruct) Update(event event.UpdateEvent) bool {
	return event.ObjectOld.GetNamespace() == r.secretsNamespace
}

func (r *deletedOrChangedInSecretsNamespaceStruct) Delete(event event.DeleteEvent) bool {
	return event.Object.GetNamespace() == r.secretsNamespace
}

func (r *deletedOrChangedInSecretsNamespaceStruct) Generic(event.GenericEvent) bool {
	return false
}
