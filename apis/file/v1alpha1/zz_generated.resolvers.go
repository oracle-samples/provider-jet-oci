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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha11 "github.com/crossplane-contrib/provider-jet-oci/apis/core/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this StorageExport.
func (mg *StorageExport) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExportSetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ExportSetIDRef,
		Selector:     mg.Spec.ForProvider.ExportSetIDSelector,
		To: reference.To{
			List:    &StorageExportSetList{},
			Managed: &StorageExportSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExportSetID")
	}
	mg.Spec.ForProvider.ExportSetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ExportSetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FileSystemIDRef,
		Selector:     mg.Spec.ForProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &StorageFileSystemList{},
			Managed: &StorageFileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FileSystemID")
	}
	mg.Spec.ForProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FileSystemIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StorageExportSet.
func (mg *StorageExportSet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MountTargetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MountTargetIDRef,
		Selector:     mg.Spec.ForProvider.MountTargetIDSelector,
		To: reference.To{
			List:    &StorageMountTargetList{},
			Managed: &StorageMountTarget{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MountTargetID")
	}
	mg.Spec.ForProvider.MountTargetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MountTargetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StorageFileSystem.
func (mg *StorageFileSystem) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StorageMountTarget.
func (mg *StorageMountTarget) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha11.SubnetList{},
			Managed: &v1alpha11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StorageReplication.
func (mg *StorageReplication) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SourceIDRef,
		Selector:     mg.Spec.ForProvider.SourceIDSelector,
		To: reference.To{
			List:    &StorageFileSystemList{},
			Managed: &StorageFileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceID")
	}
	mg.Spec.ForProvider.SourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TargetIDRef,
		Selector:     mg.Spec.ForProvider.TargetIDSelector,
		To: reference.To{
			List:    &StorageFileSystemList{},
			Managed: &StorageFileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetID")
	}
	mg.Spec.ForProvider.TargetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StorageSnapshot.
func (mg *StorageSnapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FileSystemID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FileSystemIDRef,
		Selector:     mg.Spec.ForProvider.FileSystemIDSelector,
		To: reference.To{
			List:    &StorageFileSystemList{},
			Managed: &StorageFileSystem{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FileSystemID")
	}
	mg.Spec.ForProvider.FileSystemID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FileSystemIDRef = rsp.ResolvedReference

	return nil
}
