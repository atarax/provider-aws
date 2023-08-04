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

// Code generated by ack-generate. DO NOT EDIT.

package documentationversion

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigateway/v1alpha1"
	awsclient "github.com/crossplane-contrib/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an DocumentationVersion resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create DocumentationVersion in AWS"
	errUpdate        = "cannot update DocumentationVersion in AWS"
	errDescribe      = "failed to describe DocumentationVersion"
	errDelete        = "failed to delete DocumentationVersion"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.DocumentationVersion)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.DocumentationVersion)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetDocumentationVersionInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetDocumentationVersionWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateDocumentationVersion(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, diff, err := e.isUpToDate(ctx, cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.DocumentationVersion)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateDocumentationVersionInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateDocumentationVersionWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.CreatedDate != nil {
		cr.Status.AtProvider.CreatedDate = &metav1.Time{*resp.CreatedDate}
	} else {
		cr.Status.AtProvider.CreatedDate = nil
	}
	if resp.Description != nil {
		cr.Spec.ForProvider.Description = resp.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Version != nil {
		cr.Status.AtProvider.Version = resp.Version
	} else {
		cr.Status.AtProvider.Version = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.DocumentationVersion)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateDocumentationVersionInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateDocumentationVersionWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.DocumentationVersion)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteDocumentationVersionInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteDocumentationVersionWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.APIGatewayAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.APIGatewayAPI
	preObserve     func(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.GetDocumentationVersionInput) error
	postObserve    func(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.DocumentationVersion, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.DocumentationVersionParameters, *svcsdk.DocumentationVersion) error
	isUpToDate     func(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.DocumentationVersion) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.CreateDocumentationVersionInput) error
	postCreate     func(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.DocumentationVersion, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.DeleteDocumentationVersionInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.DeleteDocumentationVersionOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.UpdateDocumentationVersionInput) error
	postUpdate     func(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.DocumentationVersion, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.GetDocumentationVersionInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.DocumentationVersion, _ *svcsdk.DocumentationVersion, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.DocumentationVersionParameters, *svcsdk.DocumentationVersion) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.DocumentationVersion) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.CreateDocumentationVersionInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.DocumentationVersion, _ *svcsdk.DocumentationVersion, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.DeleteDocumentationVersionInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.DocumentationVersion, _ *svcsdk.DeleteDocumentationVersionOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.DocumentationVersion, *svcsdk.UpdateDocumentationVersionInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.DocumentationVersion, _ *svcsdk.DocumentationVersion, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
