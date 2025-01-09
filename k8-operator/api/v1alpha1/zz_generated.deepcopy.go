//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSIamAuthDetails) DeepCopyInto(out *AWSIamAuthDetails) {
	*out = *in
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSIamAuthDetails.
func (in *AWSIamAuthDetails) DeepCopy() *AWSIamAuthDetails {
	if in == nil {
		return nil
	}
	out := new(AWSIamAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	out.ServiceAccount = in.ServiceAccount
	out.ServiceToken = in.ServiceToken
	out.UniversalAuth = in.UniversalAuth
	out.KubernetesAuth = in.KubernetesAuth
	out.AwsIamAuth = in.AwsIamAuth
	out.AzureAuth = in.AzureAuth
	out.GcpIdTokenAuth = in.GcpIdTokenAuth
	out.GcpIamAuth = in.GcpIamAuth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAuthDetails) DeepCopyInto(out *AzureAuthDetails) {
	*out = *in
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAuthDetails.
func (in *AzureAuthDetails) DeepCopy() *AzureAuthDetails {
	if in == nil {
		return nil
	}
	out := new(AzureAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaReference) DeepCopyInto(out *CaReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaReference.
func (in *CaReference) DeepCopy() *CaReference {
	if in == nil {
		return nil
	}
	out := new(CaReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicSecretDetails) DeepCopyInto(out *DynamicSecretDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicSecretDetails.
func (in *DynamicSecretDetails) DeepCopy() *DynamicSecretDetails {
	if in == nil {
		return nil
	}
	out := new(DynamicSecretDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPIdTokenAuthDetails) DeepCopyInto(out *GCPIdTokenAuthDetails) {
	*out = *in
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPIdTokenAuthDetails.
func (in *GCPIdTokenAuthDetails) DeepCopy() *GCPIdTokenAuthDetails {
	if in == nil {
		return nil
	}
	out := new(GCPIdTokenAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpIamAuthDetails) DeepCopyInto(out *GcpIamAuthDetails) {
	*out = *in
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpIamAuthDetails.
func (in *GcpIamAuthDetails) DeepCopy() *GcpIamAuthDetails {
	if in == nil {
		return nil
	}
	out := new(GcpIamAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericAwsIamAuth) DeepCopyInto(out *GenericAwsIamAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericAwsIamAuth.
func (in *GenericAwsIamAuth) DeepCopy() *GenericAwsIamAuth {
	if in == nil {
		return nil
	}
	out := new(GenericAwsIamAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericAzureAuth) DeepCopyInto(out *GenericAzureAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericAzureAuth.
func (in *GenericAzureAuth) DeepCopy() *GenericAzureAuth {
	if in == nil {
		return nil
	}
	out := new(GenericAzureAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericGcpIamAuth) DeepCopyInto(out *GenericGcpIamAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericGcpIamAuth.
func (in *GenericGcpIamAuth) DeepCopy() *GenericGcpIamAuth {
	if in == nil {
		return nil
	}
	out := new(GenericGcpIamAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericGcpIdTokenAuth) DeepCopyInto(out *GenericGcpIdTokenAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericGcpIdTokenAuth.
func (in *GenericGcpIdTokenAuth) DeepCopy() *GenericGcpIdTokenAuth {
	if in == nil {
		return nil
	}
	out := new(GenericGcpIdTokenAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericInfisicalAuthentication) DeepCopyInto(out *GenericInfisicalAuthentication) {
	*out = *in
	out.UniversalAuth = in.UniversalAuth
	out.KubernetesAuth = in.KubernetesAuth
	out.AwsIamAuth = in.AwsIamAuth
	out.AzureAuth = in.AzureAuth
	out.GcpIdTokenAuth = in.GcpIdTokenAuth
	out.GcpIamAuth = in.GcpIamAuth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericInfisicalAuthentication.
func (in *GenericInfisicalAuthentication) DeepCopy() *GenericInfisicalAuthentication {
	if in == nil {
		return nil
	}
	out := new(GenericInfisicalAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericKubernetesAuth) DeepCopyInto(out *GenericKubernetesAuth) {
	*out = *in
	out.ServiceAccountRef = in.ServiceAccountRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericKubernetesAuth.
func (in *GenericKubernetesAuth) DeepCopy() *GenericKubernetesAuth {
	if in == nil {
		return nil
	}
	out := new(GenericKubernetesAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericUniversalAuth) DeepCopyInto(out *GenericUniversalAuth) {
	*out = *in
	out.CredentialsRef = in.CredentialsRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericUniversalAuth.
func (in *GenericUniversalAuth) DeepCopy() *GenericUniversalAuth {
	if in == nil {
		return nil
	}
	out := new(GenericUniversalAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalDynamicSecret) DeepCopyInto(out *InfisicalDynamicSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalDynamicSecret.
func (in *InfisicalDynamicSecret) DeepCopy() *InfisicalDynamicSecret {
	if in == nil {
		return nil
	}
	out := new(InfisicalDynamicSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalDynamicSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalDynamicSecretLease) DeepCopyInto(out *InfisicalDynamicSecretLease) {
	*out = *in
	in.CreationTimestamp.DeepCopyInto(&out.CreationTimestamp)
	in.ExpiresAt.DeepCopyInto(&out.ExpiresAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalDynamicSecretLease.
func (in *InfisicalDynamicSecretLease) DeepCopy() *InfisicalDynamicSecretLease {
	if in == nil {
		return nil
	}
	out := new(InfisicalDynamicSecretLease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalDynamicSecretList) DeepCopyInto(out *InfisicalDynamicSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InfisicalDynamicSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalDynamicSecretList.
func (in *InfisicalDynamicSecretList) DeepCopy() *InfisicalDynamicSecretList {
	if in == nil {
		return nil
	}
	out := new(InfisicalDynamicSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalDynamicSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalDynamicSecretSpec) DeepCopyInto(out *InfisicalDynamicSecretSpec) {
	*out = *in
	in.ManagedSecretReference.DeepCopyInto(&out.ManagedSecretReference)
	out.Authentication = in.Authentication
	out.DynamicSecret = in.DynamicSecret
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalDynamicSecretSpec.
func (in *InfisicalDynamicSecretSpec) DeepCopy() *InfisicalDynamicSecretSpec {
	if in == nil {
		return nil
	}
	out := new(InfisicalDynamicSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalDynamicSecretStatus) DeepCopyInto(out *InfisicalDynamicSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Lease != nil {
		in, out := &in.Lease, &out.Lease
		*out = new(InfisicalDynamicSecretLease)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalDynamicSecretStatus.
func (in *InfisicalDynamicSecretStatus) DeepCopy() *InfisicalDynamicSecretStatus {
	if in == nil {
		return nil
	}
	out := new(InfisicalDynamicSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecret) DeepCopyInto(out *InfisicalPushSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecret.
func (in *InfisicalPushSecret) DeepCopy() *InfisicalPushSecret {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalPushSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecretDestination) DeepCopyInto(out *InfisicalPushSecretDestination) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecretDestination.
func (in *InfisicalPushSecretDestination) DeepCopy() *InfisicalPushSecretDestination {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecretDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecretList) DeepCopyInto(out *InfisicalPushSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InfisicalPushSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecretList.
func (in *InfisicalPushSecretList) DeepCopy() *InfisicalPushSecretList {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalPushSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecretSpec) DeepCopyInto(out *InfisicalPushSecretSpec) {
	*out = *in
	out.Destination = in.Destination
	out.Authentication = in.Authentication
	out.Push = in.Push
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecretSpec.
func (in *InfisicalPushSecretSpec) DeepCopy() *InfisicalPushSecretSpec {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecretStatus) DeepCopyInto(out *InfisicalPushSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagedSecrets != nil {
		in, out := &in.ManagedSecrets, &out.ManagedSecrets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecretStatus.
func (in *InfisicalPushSecretStatus) DeepCopy() *InfisicalPushSecretStatus {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecret) DeepCopyInto(out *InfisicalSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecret.
func (in *InfisicalSecret) DeepCopy() *InfisicalSecret {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretList) DeepCopyInto(out *InfisicalSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InfisicalSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretList.
func (in *InfisicalSecretList) DeepCopy() *InfisicalSecretList {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretSpec) DeepCopyInto(out *InfisicalSecretSpec) {
	*out = *in
	out.TokenSecretReference = in.TokenSecretReference
	out.Authentication = in.Authentication
	in.ManagedSecretReference.DeepCopyInto(&out.ManagedSecretReference)
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretSpec.
func (in *InfisicalSecretSpec) DeepCopy() *InfisicalSecretSpec {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretStatus) DeepCopyInto(out *InfisicalSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretStatus.
func (in *InfisicalSecretStatus) DeepCopy() *InfisicalSecretStatus {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretTemplate) DeepCopyInto(out *InfisicalSecretTemplate) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretTemplate.
func (in *InfisicalSecretTemplate) DeepCopy() *InfisicalSecretTemplate {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSecretReference) DeepCopyInto(out *KubeSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSecretReference.
func (in *KubeSecretReference) DeepCopy() *KubeSecretReference {
	if in == nil {
		return nil
	}
	out := new(KubeSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesAuthDetails) DeepCopyInto(out *KubernetesAuthDetails) {
	*out = *in
	out.ServiceAccountRef = in.ServiceAccountRef
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesAuthDetails.
func (in *KubernetesAuthDetails) DeepCopy() *KubernetesAuthDetails {
	if in == nil {
		return nil
	}
	out := new(KubernetesAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceAccountRef) DeepCopyInto(out *KubernetesServiceAccountRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceAccountRef.
func (in *KubernetesServiceAccountRef) DeepCopy() *KubernetesServiceAccountRef {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceAccountRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineIdentityScopeInWorkspace) DeepCopyInto(out *MachineIdentityScopeInWorkspace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineIdentityScopeInWorkspace.
func (in *MachineIdentityScopeInWorkspace) DeepCopy() *MachineIdentityScopeInWorkspace {
	if in == nil {
		return nil
	}
	out := new(MachineIdentityScopeInWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKubeSecretConfig) DeepCopyInto(out *ManagedKubeSecretConfig) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(InfisicalSecretTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKubeSecretConfig.
func (in *ManagedKubeSecretConfig) DeepCopy() *ManagedKubeSecretConfig {
	if in == nil {
		return nil
	}
	out := new(ManagedKubeSecretConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretPush) DeepCopyInto(out *SecretPush) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretPush.
func (in *SecretPush) DeepCopy() *SecretPush {
	if in == nil {
		return nil
	}
	out := new(SecretPush)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretScopeInWorkspace) DeepCopyInto(out *SecretScopeInWorkspace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretScopeInWorkspace.
func (in *SecretScopeInWorkspace) DeepCopy() *SecretScopeInWorkspace {
	if in == nil {
		return nil
	}
	out := new(SecretScopeInWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountDetails) DeepCopyInto(out *ServiceAccountDetails) {
	*out = *in
	out.ServiceAccountSecretReference = in.ServiceAccountSecretReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountDetails.
func (in *ServiceAccountDetails) DeepCopy() *ServiceAccountDetails {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTokenDetails) DeepCopyInto(out *ServiceTokenDetails) {
	*out = *in
	out.ServiceTokenSecretReference = in.ServiceTokenSecretReference
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTokenDetails.
func (in *ServiceTokenDetails) DeepCopy() *ServiceTokenDetails {
	if in == nil {
		return nil
	}
	out := new(ServiceTokenDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	out.CaRef = in.CaRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UniversalAuthDetails) DeepCopyInto(out *UniversalAuthDetails) {
	*out = *in
	out.CredentialsRef = in.CredentialsRef
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UniversalAuthDetails.
func (in *UniversalAuthDetails) DeepCopy() *UniversalAuthDetails {
	if in == nil {
		return nil
	}
	out := new(UniversalAuthDetails)
	in.DeepCopyInto(out)
	return out
}
