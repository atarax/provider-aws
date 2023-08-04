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

package api

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigatewayv2/v1alpha1"
	awsclient "github.com/crossplane-contrib/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an API resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create API in AWS"
	errUpdate        = "cannot update API in AWS"
	errDescribe      = "failed to describe API"
	errDelete        = "failed to delete API"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.API)
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
	cr, ok := mg.(*svcapitypes.API)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetApiInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetApiWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateAPI(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

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
	cr, ok := mg.(*svcapitypes.API)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateApiInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateApiWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.ApiEndpoint != nil {
		cr.Status.AtProvider.APIEndpoint = resp.ApiEndpoint
	} else {
		cr.Status.AtProvider.APIEndpoint = nil
	}
	if resp.ApiGatewayManaged != nil {
		cr.Status.AtProvider.APIGatewayManaged = resp.ApiGatewayManaged
	} else {
		cr.Status.AtProvider.APIGatewayManaged = nil
	}
	if resp.ApiId != nil {
		cr.Status.AtProvider.APIID = resp.ApiId
	} else {
		cr.Status.AtProvider.APIID = nil
	}
	if resp.ApiKeySelectionExpression != nil {
		cr.Spec.ForProvider.APIKeySelectionExpression = resp.ApiKeySelectionExpression
	} else {
		cr.Spec.ForProvider.APIKeySelectionExpression = nil
	}
	if resp.CorsConfiguration != nil {
		f4 := &svcapitypes.CORS{}
		if resp.CorsConfiguration.AllowCredentials != nil {
			f4.AllowCredentials = resp.CorsConfiguration.AllowCredentials
		}
		if resp.CorsConfiguration.AllowHeaders != nil {
			f4f1 := []*string{}
			for _, f4f1iter := range resp.CorsConfiguration.AllowHeaders {
				var f4f1elem string
				f4f1elem = *f4f1iter
				f4f1 = append(f4f1, &f4f1elem)
			}
			f4.AllowHeaders = f4f1
		}
		if resp.CorsConfiguration.AllowMethods != nil {
			f4f2 := []*string{}
			for _, f4f2iter := range resp.CorsConfiguration.AllowMethods {
				var f4f2elem string
				f4f2elem = *f4f2iter
				f4f2 = append(f4f2, &f4f2elem)
			}
			f4.AllowMethods = f4f2
		}
		if resp.CorsConfiguration.AllowOrigins != nil {
			f4f3 := []*string{}
			for _, f4f3iter := range resp.CorsConfiguration.AllowOrigins {
				var f4f3elem string
				f4f3elem = *f4f3iter
				f4f3 = append(f4f3, &f4f3elem)
			}
			f4.AllowOrigins = f4f3
		}
		if resp.CorsConfiguration.ExposeHeaders != nil {
			f4f4 := []*string{}
			for _, f4f4iter := range resp.CorsConfiguration.ExposeHeaders {
				var f4f4elem string
				f4f4elem = *f4f4iter
				f4f4 = append(f4f4, &f4f4elem)
			}
			f4.ExposeHeaders = f4f4
		}
		if resp.CorsConfiguration.MaxAge != nil {
			f4.MaxAge = resp.CorsConfiguration.MaxAge
		}
		cr.Spec.ForProvider.CORSConfiguration = f4
	} else {
		cr.Spec.ForProvider.CORSConfiguration = nil
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
	if resp.DisableExecuteApiEndpoint != nil {
		cr.Spec.ForProvider.DisableExecuteAPIEndpoint = resp.DisableExecuteApiEndpoint
	} else {
		cr.Spec.ForProvider.DisableExecuteAPIEndpoint = nil
	}
	if resp.DisableSchemaValidation != nil {
		cr.Spec.ForProvider.DisableSchemaValidation = resp.DisableSchemaValidation
	} else {
		cr.Spec.ForProvider.DisableSchemaValidation = nil
	}
	if resp.ImportInfo != nil {
		f9 := []*string{}
		for _, f9iter := range resp.ImportInfo {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		cr.Status.AtProvider.ImportInfo = f9
	} else {
		cr.Status.AtProvider.ImportInfo = nil
	}
	if resp.Name != nil {
		cr.Spec.ForProvider.Name = resp.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}
	if resp.ProtocolType != nil {
		cr.Spec.ForProvider.ProtocolType = resp.ProtocolType
	} else {
		cr.Spec.ForProvider.ProtocolType = nil
	}
	if resp.RouteSelectionExpression != nil {
		cr.Spec.ForProvider.RouteSelectionExpression = resp.RouteSelectionExpression
	} else {
		cr.Spec.ForProvider.RouteSelectionExpression = nil
	}
	if resp.Tags != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range resp.Tags {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		cr.Spec.ForProvider.Tags = f13
	} else {
		cr.Spec.ForProvider.Tags = nil
	}
	if resp.Version != nil {
		cr.Spec.ForProvider.Version = resp.Version
	} else {
		cr.Spec.ForProvider.Version = nil
	}
	if resp.Warnings != nil {
		f15 := []*string{}
		for _, f15iter := range resp.Warnings {
			var f15elem string
			f15elem = *f15iter
			f15 = append(f15, &f15elem)
		}
		cr.Status.AtProvider.Warnings = f15
	} else {
		cr.Status.AtProvider.Warnings = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.API)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateApiInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateApiWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.API)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteApiInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteApiWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ApiGatewayV2API, opts []option) *external {
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
	client         svcsdkapi.ApiGatewayV2API
	preObserve     func(context.Context, *svcapitypes.API, *svcsdk.GetApiInput) error
	postObserve    func(context.Context, *svcapitypes.API, *svcsdk.GetApiOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.APIParameters, *svcsdk.GetApiOutput) error
	isUpToDate     func(context.Context, *svcapitypes.API, *svcsdk.GetApiOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.API, *svcsdk.CreateApiInput) error
	postCreate     func(context.Context, *svcapitypes.API, *svcsdk.CreateApiOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.API, *svcsdk.DeleteApiInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.API, *svcsdk.DeleteApiOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.API, *svcsdk.UpdateApiInput) error
	postUpdate     func(context.Context, *svcapitypes.API, *svcsdk.UpdateApiOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.API, *svcsdk.GetApiInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.API, _ *svcsdk.GetApiOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.APIParameters, *svcsdk.GetApiOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.API, *svcsdk.GetApiOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.API, *svcsdk.CreateApiInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.API, _ *svcsdk.CreateApiOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.API, *svcsdk.DeleteApiInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.API, _ *svcsdk.DeleteApiOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.API, *svcsdk.UpdateApiInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.API, _ *svcsdk.UpdateApiOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
