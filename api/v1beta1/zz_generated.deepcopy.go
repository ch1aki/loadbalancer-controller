// +build !ignore_autogenerated

/*


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

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSBackend) DeepCopyInto(out *AWSBackend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSBackend.
func (in *AWSBackend) DeepCopy() *AWSBackend {
	if in == nil {
		return nil
	}
	out := new(AWSBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSBackend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSBackendAction) DeepCopyInto(out *AWSBackendAction) {
	*out = *in
	in.TargetGroup.DeepCopyInto(&out.TargetGroup)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSBackendAction.
func (in *AWSBackendAction) DeepCopy() *AWSBackendAction {
	if in == nil {
		return nil
	}
	out := new(AWSBackendAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSBackendCredentials) DeepCopyInto(out *AWSBackendCredentials) {
	*out = *in
	if in.AccesskeyID != nil {
		in, out := &in.AccesskeyID, &out.AccesskeyID
		*out = new(v1.EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretAccesskey != nil {
		in, out := &in.SecretAccesskey, &out.SecretAccesskey
		*out = new(v1.EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSBackendCredentials.
func (in *AWSBackendCredentials) DeepCopy() *AWSBackendCredentials {
	if in == nil {
		return nil
	}
	out := new(AWSBackendCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSBackendDestination) DeepCopyInto(out *AWSBackendDestination) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSBackendDestination.
func (in *AWSBackendDestination) DeepCopy() *AWSBackendDestination {
	if in == nil {
		return nil
	}
	out := new(AWSBackendDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSBackendList) DeepCopyInto(out *AWSBackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSBackendList.
func (in *AWSBackendList) DeepCopy() *AWSBackendList {
	if in == nil {
		return nil
	}
	out := new(AWSBackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSBackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSBackendSpec) DeepCopyInto(out *AWSBackendSpec) {
	*out = *in
	in.Credentials.DeepCopyInto(&out.Credentials)
	out.VPC = in.VPC
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]Identifier, len(*in))
		copy(*out, *in)
	}
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make([]Listener, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSBackendSpec.
func (in *AWSBackendSpec) DeepCopy() *AWSBackendSpec {
	if in == nil {
		return nil
	}
	out := new(AWSBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSBackendTarget) DeepCopyInto(out *AWSBackendTarget) {
	*out = *in
	out.Destination = in.Destination
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSBackendTarget.
func (in *AWSBackendTarget) DeepCopy() *AWSBackendTarget {
	if in == nil {
		return nil
	}
	out := new(AWSBackendTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSBackendTargetGroup) DeepCopyInto(out *AWSBackendTargetGroup) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]AWSBackendTarget, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSBackendTargetGroup.
func (in *AWSBackendTargetGroup) DeepCopy() *AWSBackendTargetGroup {
	if in == nil {
		return nil
	}
	out := new(AWSBackendTargetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendEndpoint) DeepCopyInto(out *BackendEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendEndpoint.
func (in *BackendEndpoint) DeepCopy() *BackendEndpoint {
	if in == nil {
		return nil
	}
	out := new(BackendEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendListener) DeepCopyInto(out *BackendListener) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendListener.
func (in *BackendListener) DeepCopy() *BackendListener {
	if in == nil {
		return nil
	}
	out := new(BackendListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendStatus) DeepCopyInto(out *BackendStatus) {
	*out = *in
	out.Endpoint = in.Endpoint
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make([]BackendListener, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendStatus.
func (in *BackendStatus) DeepCopy() *BackendStatus {
	if in == nil {
		return nil
	}
	out := new(BackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalBackend) DeepCopyInto(out *ExternalBackend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalBackend.
func (in *ExternalBackend) DeepCopy() *ExternalBackend {
	if in == nil {
		return nil
	}
	out := new(ExternalBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalBackend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalBackendList) DeepCopyInto(out *ExternalBackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalBackendList.
func (in *ExternalBackendList) DeepCopy() *ExternalBackendList {
	if in == nil {
		return nil
	}
	out := new(ExternalBackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalBackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalBackendSpec) DeepCopyInto(out *ExternalBackendSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalBackendSpec.
func (in *ExternalBackendSpec) DeepCopy() *ExternalBackendSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalBackendStatus) DeepCopyInto(out *ExternalBackendStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalBackendStatus.
func (in *ExternalBackendStatus) DeepCopy() *ExternalBackendStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalBackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identifier) DeepCopyInto(out *Identifier) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identifier.
func (in *Identifier) DeepCopy() *Identifier {
	if in == nil {
		return nil
	}
	out := new(Identifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listener) DeepCopyInto(out *Listener) {
	*out = *in
	in.DefaultAction.DeepCopyInto(&out.DefaultAction)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listener.
func (in *Listener) DeepCopy() *Listener {
	if in == nil {
		return nil
	}
	out := new(Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Loadbalancer) DeepCopyInto(out *Loadbalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Loadbalancer.
func (in *Loadbalancer) DeepCopy() *Loadbalancer {
	if in == nil {
		return nil
	}
	out := new(Loadbalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Loadbalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerList) DeepCopyInto(out *LoadbalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Loadbalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerList.
func (in *LoadbalancerList) DeepCopy() *LoadbalancerList {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadbalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerSpec) DeepCopyInto(out *LoadbalancerSpec) {
	*out = *in
	in.AWSBackend.DeepCopyInto(&out.AWSBackend)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerSpec.
func (in *LoadbalancerSpec) DeepCopy() *LoadbalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadbalancerStatus) DeepCopyInto(out *LoadbalancerStatus) {
	*out = *in
	in.Backend.DeepCopyInto(&out.Backend)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadbalancerStatus.
func (in *LoadbalancerStatus) DeepCopy() *LoadbalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadbalancerStatus)
	in.DeepCopyInto(out)
	return out
}
