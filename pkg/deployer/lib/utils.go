// SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package lib

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kutil "github.com/gardener/landscaper/pkg/utils/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	lsv1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
	lsv1alpha1helper "github.com/gardener/landscaper/apis/core/v1alpha1/helper"
)

// HandleAnnotationsAndGeneration is meant to be called at the beginning of a deployer's reconcile loop.
// If a reconcile is needed due to the reconcile annotation or a change in the generation, it will set the phase to Init and remove the reconcile annotation.
// It will also remove the timeout annotation if it is set.
// Returns:
//   - the modified deployitem
//   - an error, if updating the deployitem failed, nil otherwise
func HandleAnnotationsAndGeneration(ctx context.Context,
	log logr.Logger,
	kubeClient client.Client,
	di *lsv1alpha1.DeployItem,
	deployerInfo lsv1alpha1.DeployerInformation) error {
	changedMeta := false
	hasReconcileAnnotation := lsv1alpha1helper.HasOperation(di.ObjectMeta, lsv1alpha1.ReconcileOperation)
	if hasReconcileAnnotation || di.Status.ObservedGeneration != di.Generation {
		// reconcile necessary due to one of
		// - reconcile annotation
		// - outdated generation
		log.V(5).Info("reconcile required, setting observed generation, phase, and last change reconcile timestamp", "reconcileAnnotation", hasReconcileAnnotation, "observedGeneration", di.Status.ObservedGeneration, "generation", di.Generation)
		di.Status.ObservedGeneration = di.Generation
		di.Status.Phase = lsv1alpha1.ExecutionPhaseInit
		now := metav1.Now()
		di.Status.LastReconcileTime = &now

		log.V(7).Info("updating status")
		if err := kubeClient.Status().Update(ctx, di); err != nil {
			return err
		}
		log.V(7).Info("successfully updated status")
	}
	if hasReconcileAnnotation {
		log.V(5).Info("removing reconcile annotation")
		changedMeta = true
		delete(di.ObjectMeta.Annotations, lsv1alpha1.OperationAnnotation)
	}
	if metav1.HasAnnotation(di.ObjectMeta, string(lsv1alpha1helper.ReconcileTimestamp)) {
		log.V(5).Info("removing timestamp annotation")
		changedMeta = true
		delete(di.ObjectMeta.Annotations, lsv1alpha1.ReconcileTimestampAnnotation)
	}
	if di.Status.Deployer.Identity != deployerInfo.Identity {
		di.Status.Deployer = deployerInfo
		changedMeta = true
	}

	if changedMeta {
		log.V(7).Info("updating metadata")
		if err := kubeClient.Update(ctx, di); err != nil {
			return err
		}
		log.V(7).Info("successfully updated metadata")
	}

	return nil
}

// ShouldReconcile returns true if the given deploy item should be reconciled
func ShouldReconcile(di *lsv1alpha1.DeployItem) bool {
	if di.Status.Phase == lsv1alpha1.ExecutionPhaseInit || di.Status.Phase == lsv1alpha1.ExecutionPhaseProgressing || di.Status.Phase == lsv1alpha1.ExecutionPhaseDeleting {
		return true
	}

	return false
}

// GetKubeconfigFromTargetConfig fetches the kubeconfig from a given config.
// If the config defines the target from a secret that secret is read from all provided clients.
func GetKubeconfigFromTargetConfig(ctx context.Context, config *lsv1alpha1.KubernetesClusterTargetConfig, kubeClients ...client.Client) ([]byte, error) {
	if config.Kubeconfig.StrVal != nil {
		return []byte(*config.Kubeconfig.StrVal), nil
	}
	if config.Kubeconfig.SecretRef == nil {
		return nil, errors.New("kubeconfig not defined")
	}

	ref := config.Kubeconfig.SecretRef
	var errList []error
	for _, kubeClient := range kubeClients {
		secret := &corev1.Secret{}
		if err := kubeClient.Get(ctx, ref.NamespacedName(), secret); err != nil {
			if !apierrors.IsNotFound(err) {
				errList = append(errList, err)
			}
			continue
		}

		if len(ref.Key) == 0 {
			ref.Key = lsv1alpha1.DefaultKubeconfigKey
		}

		kubeconfig, ok := secret.Data[ref.Key]
		if !ok {
			errList = append(errList, fmt.Errorf("secret found but key %q not found", ref.Key))
			continue
		}
		return kubeconfig, nil
	}
	if len(errList) != 0 {
		return nil, utilerrors.NewAggregate(errList)
	}
	return nil, apierrors.NewNotFound(schema.GroupResource{
		Resource: "secret",
	}, ref.Name)
}

// SetProviderStatus sets the provider specific status for a deploy item.
func SetProviderStatus(di *lsv1alpha1.DeployItem, status runtime.Object, scheme *runtime.Scheme) error {
	rawStatus, err := kutil.ConvertToRawExtension(status, scheme)
	if err != nil {
		return err
	}
	di.Status.ProviderStatus = rawStatus
	return nil
}
