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

package vpcendpointserviceconfiguration

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/ec2"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/ec2/v1alpha1"
	awsclient "github.com/crossplane-contrib/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an VPCEndpointServiceConfiguration resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create VPCEndpointServiceConfiguration in AWS"
	errUpdate        = "cannot update VPCEndpointServiceConfiguration in AWS"
	errDescribe      = "failed to describe VPCEndpointServiceConfiguration"
	errDelete        = "failed to delete VPCEndpointServiceConfiguration"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.VPCEndpointServiceConfiguration)
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
	cr, ok := mg.(*svcapitypes.VPCEndpointServiceConfiguration)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeVpcEndpointServiceConfigurationsInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeVpcEndpointServiceConfigurationsWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.ServiceConfigurations) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateVPCEndpointServiceConfiguration(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.VPCEndpointServiceConfiguration)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateVpcEndpointServiceConfigurationInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateVpcEndpointServiceConfigurationWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.ClientToken != nil {
		cr.Status.AtProvider.ClientToken = resp.ClientToken
	} else {
		cr.Status.AtProvider.ClientToken = nil
	}
	if resp.ServiceConfiguration != nil {
		f1 := &svcapitypes.ServiceConfiguration{}
		if resp.ServiceConfiguration.AcceptanceRequired != nil {
			f1.AcceptanceRequired = resp.ServiceConfiguration.AcceptanceRequired
		}
		if resp.ServiceConfiguration.AvailabilityZones != nil {
			f1f1 := []*string{}
			for _, f1f1iter := range resp.ServiceConfiguration.AvailabilityZones {
				var f1f1elem string
				f1f1elem = *f1f1iter
				f1f1 = append(f1f1, &f1f1elem)
			}
			f1.AvailabilityZones = f1f1
		}
		if resp.ServiceConfiguration.BaseEndpointDnsNames != nil {
			f1f2 := []*string{}
			for _, f1f2iter := range resp.ServiceConfiguration.BaseEndpointDnsNames {
				var f1f2elem string
				f1f2elem = *f1f2iter
				f1f2 = append(f1f2, &f1f2elem)
			}
			f1.BaseEndpointDNSNames = f1f2
		}
		if resp.ServiceConfiguration.GatewayLoadBalancerArns != nil {
			f1f3 := []*string{}
			for _, f1f3iter := range resp.ServiceConfiguration.GatewayLoadBalancerArns {
				var f1f3elem string
				f1f3elem = *f1f3iter
				f1f3 = append(f1f3, &f1f3elem)
			}
			f1.GatewayLoadBalancerARNs = f1f3
		}
		if resp.ServiceConfiguration.ManagesVpcEndpoints != nil {
			f1.ManagesVPCEndpoints = resp.ServiceConfiguration.ManagesVpcEndpoints
		}
		if resp.ServiceConfiguration.NetworkLoadBalancerArns != nil {
			f1f5 := []*string{}
			for _, f1f5iter := range resp.ServiceConfiguration.NetworkLoadBalancerArns {
				var f1f5elem string
				f1f5elem = *f1f5iter
				f1f5 = append(f1f5, &f1f5elem)
			}
			f1.NetworkLoadBalancerARNs = f1f5
		}
		if resp.ServiceConfiguration.PayerResponsibility != nil {
			f1.PayerResponsibility = resp.ServiceConfiguration.PayerResponsibility
		}
		if resp.ServiceConfiguration.PrivateDnsName != nil {
			f1.PrivateDNSName = resp.ServiceConfiguration.PrivateDnsName
		}
		if resp.ServiceConfiguration.PrivateDnsNameConfiguration != nil {
			f1f8 := &svcapitypes.PrivateDNSNameConfiguration{}
			if resp.ServiceConfiguration.PrivateDnsNameConfiguration.Name != nil {
				f1f8.Name = resp.ServiceConfiguration.PrivateDnsNameConfiguration.Name
			}
			if resp.ServiceConfiguration.PrivateDnsNameConfiguration.State != nil {
				f1f8.State = resp.ServiceConfiguration.PrivateDnsNameConfiguration.State
			}
			if resp.ServiceConfiguration.PrivateDnsNameConfiguration.Type != nil {
				f1f8.Type = resp.ServiceConfiguration.PrivateDnsNameConfiguration.Type
			}
			if resp.ServiceConfiguration.PrivateDnsNameConfiguration.Value != nil {
				f1f8.Value = resp.ServiceConfiguration.PrivateDnsNameConfiguration.Value
			}
			f1.PrivateDNSNameConfiguration = f1f8
		}
		if resp.ServiceConfiguration.ServiceId != nil {
			f1.ServiceID = resp.ServiceConfiguration.ServiceId
		}
		if resp.ServiceConfiguration.ServiceName != nil {
			f1.ServiceName = resp.ServiceConfiguration.ServiceName
		}
		if resp.ServiceConfiguration.ServiceState != nil {
			f1.ServiceState = resp.ServiceConfiguration.ServiceState
		}
		if resp.ServiceConfiguration.ServiceType != nil {
			f1f12 := []*svcapitypes.ServiceTypeDetail{}
			for _, f1f12iter := range resp.ServiceConfiguration.ServiceType {
				f1f12elem := &svcapitypes.ServiceTypeDetail{}
				if f1f12iter.ServiceType != nil {
					f1f12elem.ServiceType = f1f12iter.ServiceType
				}
				f1f12 = append(f1f12, f1f12elem)
			}
			f1.ServiceType = f1f12
		}
		if resp.ServiceConfiguration.SupportedIpAddressTypes != nil {
			f1f13 := []*string{}
			for _, f1f13iter := range resp.ServiceConfiguration.SupportedIpAddressTypes {
				var f1f13elem string
				f1f13elem = *f1f13iter
				f1f13 = append(f1f13, &f1f13elem)
			}
			f1.SupportedIPAddressTypes = f1f13
		}
		if resp.ServiceConfiguration.Tags != nil {
			f1f14 := []*svcapitypes.Tag{}
			for _, f1f14iter := range resp.ServiceConfiguration.Tags {
				f1f14elem := &svcapitypes.Tag{}
				if f1f14iter.Key != nil {
					f1f14elem.Key = f1f14iter.Key
				}
				if f1f14iter.Value != nil {
					f1f14elem.Value = f1f14iter.Value
				}
				f1f14 = append(f1f14, f1f14elem)
			}
			f1.Tags = f1f14
		}
		cr.Status.AtProvider.ServiceConfiguration = f1
	} else {
		cr.Status.AtProvider.ServiceConfiguration = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.VPCEndpointServiceConfiguration)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateModifyVpcEndpointServiceConfigurationInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.ModifyVpcEndpointServiceConfigurationWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.VPCEndpointServiceConfiguration)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	return e.delete(ctx, mg)

}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.EC2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		delete:         nopDelete,
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
	client         svcsdkapi.EC2API
	preObserve     func(context.Context, *svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.DescribeVpcEndpointServiceConfigurationsInput) error
	postObserve    func(context.Context, *svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput) *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput
	lateInitialize func(*svcapitypes.VPCEndpointServiceConfigurationParameters, *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput) error
	isUpToDate     func(*svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.CreateVpcEndpointServiceConfigurationInput) error
	postCreate     func(context.Context, *svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.CreateVpcEndpointServiceConfigurationOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	delete         func(context.Context, cpresource.Managed) error
	preUpdate      func(context.Context, *svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.ModifyVpcEndpointServiceConfigurationInput) error
	postUpdate     func(context.Context, *svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.ModifyVpcEndpointServiceConfigurationOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.DescribeVpcEndpointServiceConfigurationsInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.VPCEndpointServiceConfiguration, _ *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.VPCEndpointServiceConfiguration, list *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput) *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput {
	return list
}

func nopLateInitialize(*svcapitypes.VPCEndpointServiceConfigurationParameters, *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.DescribeVpcEndpointServiceConfigurationsOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.CreateVpcEndpointServiceConfigurationInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.VPCEndpointServiceConfiguration, _ *svcsdk.CreateVpcEndpointServiceConfigurationOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopDelete(context.Context, cpresource.Managed) error {
	return nil
}
func nopPreUpdate(context.Context, *svcapitypes.VPCEndpointServiceConfiguration, *svcsdk.ModifyVpcEndpointServiceConfigurationInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.VPCEndpointServiceConfiguration, _ *svcsdk.ModifyVpcEndpointServiceConfigurationOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
