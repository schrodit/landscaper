// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package installations

import (
	"context"
	"net/http"

	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/gardener/landscaper/pkg/apis/core"
	"github.com/gardener/landscaper/pkg/apis/core/validation"
	"github.com/gardener/landscaper/pkg/utils/kubernetes/webhook"
)

// Validator decodes a admission request and validates the request.
type Validator struct {
	decoder *admission.Decoder
}

var _ admission.DecoderInjector = &Validator{}
var _ webhook.CreateValidator = &Validator{}
var _ webhook.UpdateValidator = &Validator{}

// InjectDecoder injects the decoder into a validatingHandler.
func (v *Validator) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}

func (v *Validator) ValidateCreate(_ context.Context, req admission.Request) admission.Response {
	obj := &core.Installation{}
	err := v.decoder.Decode(req, obj)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	if err := validation.ValidateInstallation(obj); err != nil {
		return admission.Denied(err.Error())
	}
	return admission.Allowed("")
}

func (v *Validator) ValidateUpdate(_ context.Context, req admission.Request) admission.Response {
	obj := &core.Installation{}
	err := v.decoder.Decode(req, obj)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	if err := validation.ValidateInstallation(obj); err != nil {
		return admission.Denied(err.Error())
	}
	return admission.Allowed("")
}
