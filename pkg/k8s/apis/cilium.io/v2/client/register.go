// Copyright 2017-2020 Authors of Cilium
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

package client

import (
	"context"
	goerrors "errors"
	"fmt"
	"time"

	k8sconstv2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	k8sversion "github.com/cilium/cilium/pkg/k8s/version"
	"github.com/cilium/cilium/pkg/logging"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/option"
	"github.com/cilium/cilium/pkg/versioncheck"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	v1client "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/yaml"
)

const (
	// subsysK8s is the value for logfields.LogSubsys
	subsysK8s = "k8s"

	// CNPCRDName is the full name of the CNP CRD.
	CNPCRDName = k8sconstv2.CNPKindDefinition + "/" + k8sconstv2.CustomResourceDefinitionVersion

	// CCNPCRDName is the full name of the CCNP CRD.
	CCNPCRDName = k8sconstv2.CCNPKindDefinition + "/" + k8sconstv2.CustomResourceDefinitionVersion

	// CEPCRDName is the full name of the CEP CRD.
	CEPCRDName = k8sconstv2.CEPKindDefinition + "/" + k8sconstv2.CustomResourceDefinitionVersion

	// CIDCRDName is the full name of the CID CRD.
	CIDCRDName = k8sconstv2.CIDKindDefinition + "/" + k8sconstv2.CustomResourceDefinitionVersion

	// CNCRDName is the full name of the CN CRD.
	CNCRDName = k8sconstv2.CNKindDefinition + "/" + k8sconstv2.CustomResourceDefinitionVersion
)

var (
	// log is the k8s package logger object.
	log = logging.DefaultLogger.WithField(logfields.LogSubsys, subsysK8s)

	comparableCRDSchemaVersion = versioncheck.MustVersion(k8sconstv2.CustomResourceDefinitionSchemaVersion)
)

// CreateCustomResourceDefinitions creates our CRD objects in the Kubernetes
// cluster.
func CreateCustomResourceDefinitions(clientset apiextensionsclient.Interface) error {
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return createCNPCRD(clientset)
	})

	g.Go(func() error {
		return createCCNPCRD(clientset)
	})

	g.Go(func() error {
		return createCEPCRD(clientset)
	})

	g.Go(func() error {
		return createNodeCRD(clientset)
	})

	if option.Config.IdentityAllocationMode == option.IdentityAllocationModeCRD {
		g.Go(func() error {
			return createIdentityCRD(clientset)
		})
	}

	return g.Wait()
}

// GetPregeneratedCRD returns the pregenerated CRD based on the requested CRD
// name. The pregenerated CRDs are generated by the controller-gen tool and
// serialized into binary form by go-bindata. This function retrieves CRDs from
// the binary form.
func GetPregeneratedCRD(crdName string) apiextensionsv1.CustomResourceDefinition {
	var (
		err      error
		crdBytes []byte
	)

	scopedLog := log.WithField("crdName", crdName)

	switch crdName {
	case CNPCRDName:
		crdBytes, err = examplesCrdsCiliumnetworkpoliciesYamlBytes()
	case CCNPCRDName:
		crdBytes, err = examplesCrdsCiliumclusterwidenetworkpoliciesYamlBytes()
	case CEPCRDName:
		crdBytes, err = examplesCrdsCiliumendpointsYamlBytes()
	case CIDCRDName:
		crdBytes, err = examplesCrdsCiliumidentitiesYamlBytes()
	case CNCRDName:
		crdBytes, err = examplesCrdsCiliumnodesYamlBytes()
	default:
		scopedLog.Fatal("Pregenerated CRD does not exist")
	}

	if err != nil {
		scopedLog.WithError(err).Fatal("Error retrieving pregenerated CRD")
	}

	ciliumCRD := apiextensionsv1.CustomResourceDefinition{}
	err = yaml.Unmarshal(crdBytes, &ciliumCRD)
	if err != nil {
		scopedLog.WithError(err).Fatal("Error unmarshalling pregenerated CRD")
	}

	return ciliumCRD
}

// createCNPCRD creates and updates the CiliumNetworkPolicies CRD. It should be called
// on agent startup but is idempotent and safe to call again.
func createCNPCRD(clientset apiextensionsclient.Interface) error {
	ciliumCRD := GetPregeneratedCRD(CNPCRDName)

	return createUpdateCRD(clientset, CNPCRDName, constructV1CRD(k8sconstv2.CNPName, ciliumCRD))
}

// createCCNPCRD creates and updates the CiliumClusterwideNetworkPolicy CRD. It
// should be called on agent startup but is idempotent and safe to call again.
func createCCNPCRD(clientset apiextensionsclient.Interface) error {
	ciliumCRD := GetPregeneratedCRD(CCNPCRDName)

	return createUpdateCRD(clientset, CCNPCRDName, constructV1CRD(k8sconstv2.CCNPName, ciliumCRD))
}

// createCEPCRD creates and updates the CiliumEndpoint CRD. It should be called
// on agent startup but is idempotent and safe to call again.
func createCEPCRD(clientset apiextensionsclient.Interface) error {
	ciliumCRD := GetPregeneratedCRD(CEPCRDName)

	return createUpdateCRD(clientset, CEPCRDName, constructV1CRD(k8sconstv2.CEPName, ciliumCRD))
}

// createNodeCRD creates and updates the CiliumNode CRD. It should be called on
// agent startup but is idempotent and safe to call again.
func createNodeCRD(clientset apiextensionsclient.Interface) error {
	ciliumCRD := GetPregeneratedCRD(CNCRDName)

	return createUpdateCRD(clientset, CNCRDName, constructV1CRD(k8sconstv2.CNName, ciliumCRD))
}

// createIdentityCRD creates and updates the CiliumIdentity CRD. It should be
// called on agent startup but is idempotent and safe to call again.
func createIdentityCRD(clientset apiextensionsclient.Interface) error {
	ciliumCRD := GetPregeneratedCRD(CIDCRDName)

	return createUpdateCRD(clientset, CIDCRDName, constructV1CRD(k8sconstv2.CIDName, ciliumCRD))
}

// createUpdateCRD ensures the CRD object is installed into the K8s cluster. It
// will create or update the CRD and its validation schema as necessary. This
// function only accepts v1 CRD objects, and defers to its v1beta1 variant if
// the cluster only supports v1beta1 CRDs. This allows us to convert all our
// CRDs into v1 form and only perform conversions on-demand, simplifying the
// code.
func createUpdateCRD(
	clientset apiextensionsclient.Interface,
	crdName string,
	crd *apiextensionsv1.CustomResourceDefinition,
) error {
	scopedLog := log.WithField("name", crdName)

	if !k8sversion.Capabilities().APIExtensionsV1CRD {
		log.Infof("K8s apiserver does not support v1 CRDs, falling back to v1beta1")

		return createUpdateV1beta1CRD(
			scopedLog,
			clientset.ApiextensionsV1beta1(),
			crdName,
			crd,
		)
	}

	v1CRDClient := clientset.ApiextensionsV1()
	clusterCRD, err := v1CRDClient.CustomResourceDefinitions().Get(
		context.TODO(),
		crd.ObjectMeta.Name,
		metav1.GetOptions{})
	if errors.IsNotFound(err) {
		scopedLog.Info("Creating CRD (CustomResourceDefinition)...")

		clusterCRD, err = v1CRDClient.CustomResourceDefinitions().Create(
			context.TODO(),
			crd,
			metav1.CreateOptions{})
		// This occurs when multiple agents race to create the CRD. Since another has
		// created it, it will also update it, hence the non-error return.
		if errors.IsAlreadyExists(err) {
			return nil
		}
	}
	if err != nil {
		return err
	}

	if err := updateV1CRD(scopedLog, crd, clusterCRD, v1CRDClient); err != nil {
		return err
	}
	if err := waitForV1CRD(scopedLog, crdName, clusterCRD, v1CRDClient); err != nil {
		return err
	}

	scopedLog.Info("CRD (CustomResourceDefinition) is installed and up-to-date")

	return nil
}

func constructV1CRD(
	name string,
	template apiextensionsv1.CustomResourceDefinition,
) *apiextensionsv1.CustomResourceDefinition {
	return &apiextensionsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				k8sconstv2.CustomResourceDefinitionSchemaVersionKey: k8sconstv2.CustomResourceDefinitionSchemaVersion,
			},
		},
		Spec: apiextensionsv1.CustomResourceDefinitionSpec{
			Group: k8sconstv2.CustomResourceDefinitionGroup,
			Names: apiextensionsv1.CustomResourceDefinitionNames{
				Kind:       template.Spec.Names.Kind,
				Plural:     template.Spec.Names.Plural,
				ShortNames: template.Spec.Names.ShortNames,
				Singular:   template.Spec.Names.Singular,
			},
			Scope:    template.Spec.Scope,
			Versions: template.Spec.Versions,
		},
	}
}

func needsUpdateV1(clusterCRD *apiextensionsv1.CustomResourceDefinition) bool {
	if clusterCRD.Spec.Versions[0].Schema == nil {
		// no validation detected
		return true
	}
	v, ok := clusterCRD.Labels[k8sconstv2.CustomResourceDefinitionSchemaVersionKey]
	if !ok {
		// no schema version detected
		return true
	}

	clusterVersion, err := versioncheck.Version(v)
	if err != nil || clusterVersion.LT(comparableCRDSchemaVersion) {
		// version in cluster is either unparsable or smaller than current version
		return true
	}

	return false
}

func updateV1CRD(
	scopedLog *logrus.Entry,
	crd, clusterCRD *apiextensionsv1.CustomResourceDefinition,
	client v1client.CustomResourceDefinitionsGetter,
) error {
	scopedLog.Debug("Checking if CRD (CustomResourceDefinition) needs update...")

	if crd.Spec.Versions[0].Schema != nil && needsUpdateV1(clusterCRD) {
		scopedLog.Info("Updating CRD (CustomResourceDefinition)...")

		// Update the CRD with the validation schema.
		err := wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
			var err error
			clusterCRD, err = client.CustomResourceDefinitions().Get(
				context.TODO(),
				crd.ObjectMeta.Name,
				metav1.GetOptions{})
			if err != nil {
				return false, err
			}

			// This seems too permissive but we only get here if the version is
			// different per needsUpdate above. If so, we want to update on any
			// validation change including adding or removing validation.
			if needsUpdateV1(clusterCRD) {
				scopedLog.Debug("CRD validation is different, updating it...")

				clusterCRD.ObjectMeta.Labels = crd.ObjectMeta.Labels
				clusterCRD.Spec = crd.Spec

				// Even though v1 CRDs omit this field by default (which also
				// means it's false) it is still carried over from the previous
				// CRD. Therefore, we must set this to false explicitly because
				// the apiserver will carry over the old value (true).
				clusterCRD.Spec.PreserveUnknownFields = false

				_, err := client.CustomResourceDefinitions().Update(
					context.TODO(),
					clusterCRD,
					metav1.UpdateOptions{})
				if err == nil {
					return true, nil
				}

				scopedLog.WithError(err).Debug("Unable to update CRD validation")

				return false, err
			}

			return true, nil
		})
		if err != nil {
			scopedLog.WithError(err).Error("Unable to update CRD")
			return err
		}
	}

	return nil
}

func waitForV1CRD(
	scopedLog *logrus.Entry,
	crdName string,
	crd *apiextensionsv1.CustomResourceDefinition,
	client v1client.CustomResourceDefinitionsGetter,
) error {
	scopedLog.Debug("Waiting for CRD (CustomResourceDefinition) to be available...")

	err := wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		for _, cond := range crd.Status.Conditions {
			switch cond.Type {
			case apiextensionsv1.Established:
				if cond.Status == apiextensionsv1.ConditionTrue {
					return true, nil
				}
			case apiextensionsv1.NamesAccepted:
				if cond.Status == apiextensionsv1.ConditionFalse {
					err := goerrors.New(cond.Reason)
					scopedLog.WithError(err).Error("Name conflict for CRD")
					return false, err
				}
			}
		}

		var err error
		if crd, err = client.CustomResourceDefinitions().Get(
			context.TODO(),
			crd.ObjectMeta.Name,
			metav1.GetOptions{}); err != nil {
			return false, err
		}
		return false, err
	})
	if err != nil {
		if deleteErr := client.CustomResourceDefinitions().Delete(
			context.TODO(),
			crd.ObjectMeta.Name,
			metav1.DeleteOptions{},
		); deleteErr != nil {
			return fmt.Errorf("unable to delete k8s %s CRD %s. Deleting CRD due: %s",
				crdName, deleteErr, err)
		}
		return err
	}

	return nil
}
