// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package webhook

import (
	"context"
	"errors"
	"net/http"

	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	admissionregv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type ValidationWebhookBuilder struct {
	name string
	mgr Manager
	obj runtime.Object
	validator interface{}
}

// CreateValidator defines functions to validate a create operation
type CreateValidator interface {
	ValidateCreate(ctx context.Context, req admission.Request) admission.Response
}

// UpdateValidator defines functions to validate a update operation
type UpdateValidator interface {
	ValidateUpdate(ctx context.Context, req admission.Request) admission.Response
}

// DeleteValidator defines functions to validate a delete operation
type DeleteValidator interface {
	ValidateDelete(ctx context.Context, req admission.Request) admission.Response
}

func NewValidationWebhookManagedBy(mgr Manager) *ValidationWebhookBuilder {
	return &ValidationWebhookBuilder{
		mgr: mgr,
	}
}

func (b *ValidationWebhookBuilder) WithValidator(val interface{}) *ValidationWebhookBuilder {
	b.validator = val
	return b
}

func (b *ValidationWebhookBuilder) WithName(name string) *ValidationWebhookBuilder {
	b.name = name
	return b
}

func (b *ValidationWebhookBuilder) For(obj runtime.Object) *ValidationWebhookBuilder {
	b.obj = obj
	return b
}

func (b *ValidationWebhookBuilder) Complete() error {
	if b.obj == nil {
		return errors.New("object is required")
	}
	if b.mgr == nil {
		return errors.New("manager is required")
	}
	if len(b.name) == 0 {
		return errors.New("name is required")
	}
	vwh := &ValidationWebhook{
		name: b.name,
		validator: b.validator,
	}
	return b.mgr.Register(vwh)
}

// ValidationWebhook decodes a admission request and validates the request.
type ValidationWebhook struct {
	name string

	decoder *admission.Decoder
	client client.Client
	obj runtime.Object
	validator interface{}
}

var _ admission.DecoderInjector = &ValidationWebhook{}
var _ inject.Client = &ValidationWebhook{}

// InjectDecoder injects the decoder into a validatingHandler.
func (h *ValidationWebhook) InjectDecoder(d *admission.Decoder) error {
	h.decoder = d
	if _, err := admission.InjectDecoderInto(d, h.validator); err != nil {
		return err
	}
	return nil
}

func (h *ValidationWebhook) Name() string {
	return h.name
}

func (h *ValidationWebhook) Handler() http.Handler {
	return &admission.Webhook{
		Handler: h,
	}
}

func (h *ValidationWebhook) InjectClient(client client.Client) error {
	h.client = client
	return nil
}

func (h *ValidationWebhook) WebhookConfiguration(namespace, serviceName string, servicePort int) (runtime.Object, error) {
	if h.client == nil {
		return nil, nil
	}

	// generate validation webhook
	cfg := &admissionregv1beta1.ValidatingWebhookConfiguration{}
	cfg.Name = h.name
	cfg.Namespace = namespace
	cfg.Webhooks = make([]admissionregv1beta1.ValidatingWebhook, 0)

	if _, ok := h.validator.(CreateValidator); ok {
		failurePolicy := admissionregv1beta1.Fail
		cfg.Webhooks = append(cfg.Webhooks, admissionregv1beta1.ValidatingWebhook{
			Name:                    "create",
			ClientConfig:            admissionregv1beta1.WebhookClientConfig{},
			Rules:                   []admissionregv1beta1.RuleWithOperations{
				{
					Operations: []admissionregv1beta1.OperationType{admissionregv1beta1.Create},
					Rule: admissionregv1beta1.Rule{
						APIGroups:   nil,
						APIVersions: nil,
						Resources:   nil,
					},
				},
			},
			FailurePolicy:           &failurePolicy,
			MatchPolicy:             nil,
			NamespaceSelector:       nil,
			ObjectSelector:          nil,
			SideEffects:             nil,
			TimeoutSeconds:          nil,
			AdmissionReviewVersions: nil,
		})
	}


	return cfg, nil
}

// Handle handles admission requests.
func (h *ValidationWebhook) Handle(ctx context.Context, req admission.Request) admission.Response {
	switch req.Operation {
	case admissionv1beta1.Create:
		val, ok := h.validator.(CreateValidator)
		if ok {
			return val.ValidateCreate(ctx, req)
		}
	case admissionv1beta1.Update:
		val, ok := h.validator.(UpdateValidator)
		if ok {
			return val.ValidateUpdate(ctx, req)
		}
	case admissionv1beta1.Delete:
		val, ok := h.validator.(DeleteValidator)
		if ok {
			return val.ValidateDelete(ctx, req)
		}

	}

	return admission.Allowed("")
}
