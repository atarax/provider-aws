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

package functionurlconfig

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/lambda"
	svcsdk "github.com/aws/aws-sdk-go/service/lambda"
	svcsdkapi "github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/lambda/v1alpha1"
	awsclient "github.com/crossplane-contrib/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an FunctionURLConfig resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create FunctionURLConfig in AWS"
	errUpdate        = "cannot update FunctionURLConfig in AWS"
	errDescribe      = "failed to describe FunctionURLConfig"
	errDelete        = "failed to delete FunctionURLConfig"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.FunctionURLConfig)
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
	cr, ok := mg.(*svcapitypes.FunctionURLConfig)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetFunctionUrlConfigInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetFunctionUrlConfigWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateFunctionURLConfig(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

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
	cr, ok := mg.(*svcapitypes.FunctionURLConfig)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateFunctionUrlConfigInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateFunctionUrlConfigWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.AuthType != nil {
		cr.Spec.ForProvider.AuthType = resp.AuthType
	} else {
		cr.Spec.ForProvider.AuthType = nil
	}
	if resp.Cors != nil {
		f1 := &svcapitypes.CORS{}
		if resp.Cors.AllowCredentials != nil {
			f1.AllowCredentials = resp.Cors.AllowCredentials
		}
		if resp.Cors.AllowHeaders != nil {
			f1f1 := []*string{}
			for _, f1f1iter := range resp.Cors.AllowHeaders {
				var f1f1elem string
				f1f1elem = *f1f1iter
				f1f1 = append(f1f1, &f1f1elem)
			}
			f1.AllowHeaders = f1f1
		}
		if resp.Cors.AllowMethods != nil {
			f1f2 := []*string{}
			for _, f1f2iter := range resp.Cors.AllowMethods {
				var f1f2elem string
				f1f2elem = *f1f2iter
				f1f2 = append(f1f2, &f1f2elem)
			}
			f1.AllowMethods = f1f2
		}
		if resp.Cors.AllowOrigins != nil {
			f1f3 := []*string{}
			for _, f1f3iter := range resp.Cors.AllowOrigins {
				var f1f3elem string
				f1f3elem = *f1f3iter
				f1f3 = append(f1f3, &f1f3elem)
			}
			f1.AllowOrigins = f1f3
		}
		if resp.Cors.ExposeHeaders != nil {
			f1f4 := []*string{}
			for _, f1f4iter := range resp.Cors.ExposeHeaders {
				var f1f4elem string
				f1f4elem = *f1f4iter
				f1f4 = append(f1f4, &f1f4elem)
			}
			f1.ExposeHeaders = f1f4
		}
		if resp.Cors.MaxAge != nil {
			f1.MaxAge = resp.Cors.MaxAge
		}
		cr.Spec.ForProvider.CORS = f1
	} else {
		cr.Spec.ForProvider.CORS = nil
	}
	if resp.CreationTime != nil {
		cr.Status.AtProvider.CreationTime = resp.CreationTime
	} else {
		cr.Status.AtProvider.CreationTime = nil
	}
	if resp.FunctionArn != nil {
		cr.Status.AtProvider.FunctionARN = resp.FunctionArn
	} else {
		cr.Status.AtProvider.FunctionARN = nil
	}
	if resp.FunctionUrl != nil {
		cr.Status.AtProvider.FunctionURL = resp.FunctionUrl
	} else {
		cr.Status.AtProvider.FunctionURL = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.FunctionURLConfig)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateFunctionUrlConfigInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateFunctionUrlConfigWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.FunctionURLConfig)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteFunctionUrlConfigInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteFunctionUrlConfigWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.LambdaAPI, opts []option) *external {
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
	client         svcsdkapi.LambdaAPI
	preObserve     func(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.GetFunctionUrlConfigInput) error
	postObserve    func(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.GetFunctionUrlConfigOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.FunctionURLConfigParameters, *svcsdk.GetFunctionUrlConfigOutput) error
	isUpToDate     func(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.GetFunctionUrlConfigOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.CreateFunctionUrlConfigInput) error
	postCreate     func(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.CreateFunctionUrlConfigOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.DeleteFunctionUrlConfigInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.DeleteFunctionUrlConfigOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.UpdateFunctionUrlConfigInput) error
	postUpdate     func(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.UpdateFunctionUrlConfigOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.GetFunctionUrlConfigInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.FunctionURLConfig, _ *svcsdk.GetFunctionUrlConfigOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.FunctionURLConfigParameters, *svcsdk.GetFunctionUrlConfigOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.GetFunctionUrlConfigOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.CreateFunctionUrlConfigInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.FunctionURLConfig, _ *svcsdk.CreateFunctionUrlConfigOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.DeleteFunctionUrlConfigInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.FunctionURLConfig, _ *svcsdk.DeleteFunctionUrlConfigOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.FunctionURLConfig, *svcsdk.UpdateFunctionUrlConfigInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.FunctionURLConfig, _ *svcsdk.UpdateFunctionUrlConfigOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
