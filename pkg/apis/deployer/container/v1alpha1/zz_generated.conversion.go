// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/gardener/landscaper/pkg/apis/config"
	container "github.com/gardener/landscaper/pkg/apis/deployer/container"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Configuration)(nil), (*container.Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Configuration_To_container_Configuration(a.(*Configuration), b.(*container.Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.Configuration)(nil), (*Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_Configuration_To_v1alpha1_Configuration(a.(*container.Configuration), b.(*Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ContainerSpec)(nil), (*container.ContainerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(a.(*ContainerSpec), b.(*container.ContainerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.ContainerSpec)(nil), (*ContainerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(a.(*container.ContainerSpec), b.(*ContainerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProviderConfiguration)(nil), (*container.ProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration(a.(*ProviderConfiguration), b.(*container.ProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.ProviderConfiguration)(nil), (*ProviderConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(a.(*container.ProviderConfiguration), b.(*ProviderConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProviderStatus)(nil), (*container.ProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ProviderStatus_To_container_ProviderStatus(a.(*ProviderStatus), b.(*container.ProviderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*container.ProviderStatus)(nil), (*ProviderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_container_ProviderStatus_To_v1alpha1_ProviderStatus(a.(*container.ProviderStatus), b.(*ProviderStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Configuration_To_container_Configuration(in *Configuration, out *container.Configuration, s conversion.Scope) error {
	out.OCI = (*config.OCIConfiguration)(unsafe.Pointer(in.OCI))
	if err := Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(&in.DefaultImage, &out.DefaultImage, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(&in.InitContainer, &out.InitContainer, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(&in.SidecarContainer, &out.SidecarContainer, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Configuration_To_container_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_container_Configuration(in *Configuration, out *container.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_container_Configuration(in, out, s)
}

func autoConvert_container_Configuration_To_v1alpha1_Configuration(in *container.Configuration, out *Configuration, s conversion.Scope) error {
	out.OCI = (*config.OCIConfiguration)(unsafe.Pointer(in.OCI))
	if err := Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(&in.DefaultImage, &out.DefaultImage, s); err != nil {
		return err
	}
	if err := Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(&in.InitContainer, &out.InitContainer, s); err != nil {
		return err
	}
	if err := Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(&in.SidecarContainer, &out.SidecarContainer, s); err != nil {
		return err
	}
	return nil
}

// Convert_container_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_container_Configuration_To_v1alpha1_Configuration(in *container.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_container_Configuration_To_v1alpha1_Configuration(in, out, s)
}

func autoConvert_v1alpha1_ContainerSpec_To_container_ContainerSpec(in *ContainerSpec, out *container.ContainerSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	return nil
}

// Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec is an autogenerated conversion function.
func Convert_v1alpha1_ContainerSpec_To_container_ContainerSpec(in *ContainerSpec, out *container.ContainerSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ContainerSpec_To_container_ContainerSpec(in, out, s)
}

func autoConvert_container_ContainerSpec_To_v1alpha1_ContainerSpec(in *container.ContainerSpec, out *ContainerSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	return nil
}

// Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec is an autogenerated conversion function.
func Convert_container_ContainerSpec_To_v1alpha1_ContainerSpec(in *container.ContainerSpec, out *ContainerSpec, s conversion.Scope) error {
	return autoConvert_container_ContainerSpec_To_v1alpha1_ContainerSpec(in, out, s)
}

func autoConvert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration(in *ProviderConfiguration, out *container.ProviderConfiguration, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	return nil
}

// Convert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration(in *ProviderConfiguration, out *container.ProviderConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderConfiguration_To_container_ProviderConfiguration(in, out, s)
}

func autoConvert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in *container.ProviderConfiguration, out *ProviderConfiguration, s conversion.Scope) error {
	out.Image = in.Image
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.Args = *(*[]string)(unsafe.Pointer(&in.Args))
	return nil
}

// Convert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration is an autogenerated conversion function.
func Convert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in *container.ProviderConfiguration, out *ProviderConfiguration, s conversion.Scope) error {
	return autoConvert_container_ProviderConfiguration_To_v1alpha1_ProviderConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ProviderStatus_To_container_ProviderStatus(in *ProviderStatus, out *container.ProviderStatus, s conversion.Scope) error {
	out.PodName = in.PodName
	out.Message = in.Message
	out.Reason = in.Reason
	out.State = in.State
	out.Image = in.Image
	out.ImageID = in.ImageID
	return nil
}

// Convert_v1alpha1_ProviderStatus_To_container_ProviderStatus is an autogenerated conversion function.
func Convert_v1alpha1_ProviderStatus_To_container_ProviderStatus(in *ProviderStatus, out *container.ProviderStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderStatus_To_container_ProviderStatus(in, out, s)
}

func autoConvert_container_ProviderStatus_To_v1alpha1_ProviderStatus(in *container.ProviderStatus, out *ProviderStatus, s conversion.Scope) error {
	out.PodName = in.PodName
	out.Message = in.Message
	out.Reason = in.Reason
	out.State = in.State
	out.Image = in.Image
	out.ImageID = in.ImageID
	return nil
}

// Convert_container_ProviderStatus_To_v1alpha1_ProviderStatus is an autogenerated conversion function.
func Convert_container_ProviderStatus_To_v1alpha1_ProviderStatus(in *container.ProviderStatus, out *ProviderStatus, s conversion.Scope) error {
	return autoConvert_container_ProviderStatus_To_v1alpha1_ProviderStatus(in, out, s)
}
