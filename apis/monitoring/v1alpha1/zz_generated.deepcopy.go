//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alarm) DeepCopyInto(out *Alarm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alarm.
func (in *Alarm) DeepCopy() *Alarm {
	if in == nil {
		return nil
	}
	out := new(Alarm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alarm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmList) DeepCopyInto(out *AlarmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alarm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmList.
func (in *AlarmList) DeepCopy() *AlarmList {
	if in == nil {
		return nil
	}
	out := new(AlarmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlarmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmObservation) DeepCopyInto(out *AlarmObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeUpdated != nil {
		in, out := &in.TimeUpdated, &out.TimeUpdated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmObservation.
func (in *AlarmObservation) DeepCopy() *AlarmObservation {
	if in == nil {
		return nil
	}
	out := new(AlarmObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmParameters) DeepCopyInto(out *AlarmParameters) {
	*out = *in
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.CompartmentIDRef != nil {
		in, out := &in.CompartmentIDRef, &out.CompartmentIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.CompartmentIDSelector != nil {
		in, out := &in.CompartmentIDSelector, &out.CompartmentIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DefinedTags != nil {
		in, out := &in.DefinedTags, &out.DefinedTags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsNotificationsPerMetricDimensionEnabled != nil {
		in, out := &in.IsNotificationsPerMetricDimensionEnabled, &out.IsNotificationsPerMetricDimensionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MessageFormat != nil {
		in, out := &in.MessageFormat, &out.MessageFormat
		*out = new(string)
		**out = **in
	}
	if in.MetricCompartmentID != nil {
		in, out := &in.MetricCompartmentID, &out.MetricCompartmentID
		*out = new(string)
		**out = **in
	}
	if in.MetricCompartmentIDInSubtree != nil {
		in, out := &in.MetricCompartmentIDInSubtree, &out.MetricCompartmentIDInSubtree
		*out = new(bool)
		**out = **in
	}
	if in.MetricCompartmentIDRef != nil {
		in, out := &in.MetricCompartmentIDRef, &out.MetricCompartmentIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.MetricCompartmentIDSelector != nil {
		in, out := &in.MetricCompartmentIDSelector, &out.MetricCompartmentIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.PendingDuration != nil {
		in, out := &in.PendingDuration, &out.PendingDuration
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.RepeatNotificationDuration != nil {
		in, out := &in.RepeatNotificationDuration, &out.RepeatNotificationDuration
		*out = new(string)
		**out = **in
	}
	if in.Resolution != nil {
		in, out := &in.Resolution, &out.Resolution
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(string)
		**out = **in
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(string)
		**out = **in
	}
	if in.Suppression != nil {
		in, out := &in.Suppression, &out.Suppression
		*out = make([]SuppressionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmParameters.
func (in *AlarmParameters) DeepCopy() *AlarmParameters {
	if in == nil {
		return nil
	}
	out := new(AlarmParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmSpec) DeepCopyInto(out *AlarmSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmSpec.
func (in *AlarmSpec) DeepCopy() *AlarmSpec {
	if in == nil {
		return nil
	}
	out := new(AlarmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlarmStatus) DeepCopyInto(out *AlarmStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlarmStatus.
func (in *AlarmStatus) DeepCopy() *AlarmStatus {
	if in == nil {
		return nil
	}
	out := new(AlarmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuppressionObservation) DeepCopyInto(out *SuppressionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuppressionObservation.
func (in *SuppressionObservation) DeepCopy() *SuppressionObservation {
	if in == nil {
		return nil
	}
	out := new(SuppressionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuppressionParameters) DeepCopyInto(out *SuppressionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.TimeSuppressFrom != nil {
		in, out := &in.TimeSuppressFrom, &out.TimeSuppressFrom
		*out = new(string)
		**out = **in
	}
	if in.TimeSuppressUntil != nil {
		in, out := &in.TimeSuppressUntil, &out.TimeSuppressUntil
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuppressionParameters.
func (in *SuppressionParameters) DeepCopy() *SuppressionParameters {
	if in == nil {
		return nil
	}
	out := new(SuppressionParameters)
	in.DeepCopyInto(out)
	return out
}
