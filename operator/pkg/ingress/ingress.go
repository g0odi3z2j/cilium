// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package ingress

import (
	"context"

	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/cilium/cilium/operator/pkg/model/translation"
	ingressTranslation "github.com/cilium/cilium/operator/pkg/model/translation/ingress"
	ciliumv2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
)

const (
	ciliumIngressPrefix    = "cilium-ingress"
	ciliumIngressClassName = "cilium"
)

// ingressReconciler reconciles a Ingress object
type ingressReconciler struct {
	logger logrus.FieldLogger
	client client.Client

	maxRetries              int
	enforcedHTTPS           bool
	useProxyProtocol        bool
	secretsNamespace        string
	lbAnnotationPrefixes    []string
	sharedLBServiceName     string
	ciliumNamespace         string
	defaultLoadbalancerMode string
	defaultSecretNamespace  string
	defaultSecretName       string
	idleTimeoutSeconds      int

	sharedTranslator    translation.Translator
	dedicatedTranslator translation.Translator
}

func newIngressReconciler(
	logger logrus.FieldLogger,
	c client.Client,
	ciliumNamespace string,
	enforceHTTPS bool,
	useProxyProtocol bool,
	secretsNamespace string,
	lbAnnotationPrefixes []string,
	sharedLBServiceName string,
	defaultLoadbalancerMode string,
	defaultSecretNamespace string,
	defaultSecretName string,
	proxyIdleTimeoutSeconds int,
) *ingressReconciler {
	return &ingressReconciler{
		logger: logger,
		client: c,

		sharedTranslator:    ingressTranslation.NewSharedIngressTranslator(sharedLBServiceName, ciliumNamespace, secretsNamespace, enforceHTTPS, useProxyProtocol, proxyIdleTimeoutSeconds),
		dedicatedTranslator: ingressTranslation.NewDedicatedIngressTranslator(secretsNamespace, enforceHTTPS, useProxyProtocol, proxyIdleTimeoutSeconds),

		maxRetries:              10,
		enforcedHTTPS:           enforceHTTPS,
		useProxyProtocol:        useProxyProtocol,
		secretsNamespace:        secretsNamespace,
		lbAnnotationPrefixes:    lbAnnotationPrefixes,
		sharedLBServiceName:     sharedLBServiceName,
		ciliumNamespace:         ciliumNamespace,
		defaultLoadbalancerMode: defaultLoadbalancerMode,
		defaultSecretNamespace:  defaultSecretNamespace,
		defaultSecretName:       defaultSecretName,
		idleTimeoutSeconds:      proxyIdleTimeoutSeconds,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *ingressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkingv1.Ingress{}).
		// (LoadBalancer) Service resource with OwnerReference to the Ingress with dedicated loadbalancing mode
		Owns(&corev1.Service{}).
		// Endpoints resource with OwnerReference to the Ingress with dedicated loadbalancing mode
		Owns(&corev1.Endpoints{}).
		// CiliumEnvoyConfig resource with OwnerReference to the Ingress with dedicated loadbalancing mode
		Owns(&ciliumv2.CiliumEnvoyConfig{}).
		// Watching shared loadbalancer Service and reconcile all shared Cilium Ingresses
		Watches(&corev1.Service{}, r.enqueueSharedCiliumIngresses(), r.forSharedLoadbalancerService()).
		Complete(r)
}

func (r *ingressReconciler) enqueueSharedCiliumIngresses() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, _ client.Object) []reconcile.Request {
		ingressList := networkingv1.IngressList{}
		if err := r.client.List(ctx, &ingressList); err != nil {
			r.logger.WithError(err).Warn("Failed to list Ingresses")
			return nil
		}

		result := []reconcile.Request{}

		for _, i := range ingressList.Items {
			// Skip Ingresses that aren't managed by Cilium
			if !isCiliumManagedIngress(ctx, r.client, r.logger, i) {
				continue
			}

			// Skip Ingresses with dedicated loadbalancer mode
			if r.isEffectiveLoadbalancerModeDedicated(&i) {
				continue
			}

			result = append(result, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: i.Namespace,
					Name:      i.Name,
				},
			})
		}

		return result
	})
}

func (r *ingressReconciler) forSharedLoadbalancerService() builder.WatchesOption {
	return builder.WithPredicates(&isSharedLoadbalancerServicePredicate{namespace: r.ciliumNamespace, name: r.sharedLBServiceName})
}

var _ predicate.Predicate = &isSharedLoadbalancerServicePredicate{}

type isSharedLoadbalancerServicePredicate struct {
	namespace string
	name      string
}

func (r *isSharedLoadbalancerServicePredicate) Create(event event.CreateEvent) bool {
	return event.Object.GetNamespace() == r.namespace && event.Object.GetName() == r.name
}

func (r *isSharedLoadbalancerServicePredicate) Update(event event.UpdateEvent) bool {
	return event.ObjectNew.GetNamespace() == r.namespace && event.ObjectNew.GetName() == r.name
}

func (r *isSharedLoadbalancerServicePredicate) Delete(event event.DeleteEvent) bool {
	return event.Object.GetNamespace() == r.namespace && event.Object.GetName() == r.name
}

func (r *isSharedLoadbalancerServicePredicate) Generic(event event.GenericEvent) bool {
	return event.Object.GetNamespace() == r.namespace && event.Object.GetName() == r.name
}
