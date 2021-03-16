// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package mock

import (
	json "encoding/json"

	runtime "k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfiguration) DeepCopyInto(out *ProviderConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Phase != nil {
		in, out := &in.Phase, &out.Phase
		*out = new(v1alpha1.ExecutionPhase)
		**out = **in
	}
	if in.ProviderStatus != nil {
		in, out := &in.ProviderStatus, &out.ProviderStatus
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Export != nil {
		in, out := &in.Export, &out.Export
		*out = new(json.RawMessage)
		if **in != nil {
			in, out := *in, *out
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfiguration.
func (in *ProviderConfiguration) DeepCopy() *ProviderConfiguration {
	if in == nil {
		return nil
	}
	out := new(ProviderConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProviderConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
