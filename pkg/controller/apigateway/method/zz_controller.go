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

package method

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigateway/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an Method resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Method in AWS"
	errUpdate        = "cannot update Method in AWS"
	errDescribe      = "failed to describe Method"
	errDelete        = "failed to delete Method"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Method)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := connectaws.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.Method)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetMethodInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetMethodWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateMethod(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	upToDate := true
	diff := ""
	if !meta.WasDeleted(cr) { // There is no need to run isUpToDate if the resource is deleted
		upToDate, diff, err = e.isUpToDate(ctx, cr, resp)
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
		}
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.Method)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GeneratePutMethodInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.PutMethodWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.ApiKeyRequired != nil {
		cr.Spec.ForProvider.APIKeyRequired = resp.ApiKeyRequired
	} else {
		cr.Spec.ForProvider.APIKeyRequired = nil
	}
	if resp.AuthorizationScopes != nil {
		f1 := []*string{}
		for _, f1iter := range resp.AuthorizationScopes {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		cr.Spec.ForProvider.AuthorizationScopes = f1
	} else {
		cr.Spec.ForProvider.AuthorizationScopes = nil
	}
	if resp.AuthorizationType != nil {
		cr.Spec.ForProvider.AuthorizationType = resp.AuthorizationType
	} else {
		cr.Spec.ForProvider.AuthorizationType = nil
	}
	if resp.AuthorizerId != nil {
		cr.Spec.ForProvider.AuthorizerID = resp.AuthorizerId
	} else {
		cr.Spec.ForProvider.AuthorizerID = nil
	}
	if resp.HttpMethod != nil {
		cr.Spec.ForProvider.HTTPMethod = resp.HttpMethod
	} else {
		cr.Spec.ForProvider.HTTPMethod = nil
	}
	if resp.OperationName != nil {
		cr.Spec.ForProvider.OperationName = resp.OperationName
	} else {
		cr.Spec.ForProvider.OperationName = nil
	}
	if resp.RequestModels != nil {
		f6 := map[string]*string{}
		for f6key, f6valiter := range resp.RequestModels {
			var f6val string
			f6val = *f6valiter
			f6[f6key] = &f6val
		}
		cr.Spec.ForProvider.RequestModels = f6
	} else {
		cr.Spec.ForProvider.RequestModels = nil
	}
	if resp.RequestParameters != nil {
		f7 := map[string]*bool{}
		for f7key, f7valiter := range resp.RequestParameters {
			var f7val bool
			f7val = *f7valiter
			f7[f7key] = &f7val
		}
		cr.Spec.ForProvider.RequestParameters = f7
	} else {
		cr.Spec.ForProvider.RequestParameters = nil
	}
	if resp.RequestValidatorId != nil {
		cr.Spec.ForProvider.RequestValidatorID = resp.RequestValidatorId
	} else {
		cr.Spec.ForProvider.RequestValidatorID = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.Method)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateMethodInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateMethodWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Method)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteMethodInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteMethodWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
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
	preObserve     func(context.Context, *svcapitypes.Method, *svcsdk.GetMethodInput) error
	postObserve    func(context.Context, *svcapitypes.Method, *svcsdk.Method, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.MethodParameters, *svcsdk.Method) error
	isUpToDate     func(context.Context, *svcapitypes.Method, *svcsdk.Method) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.Method, *svcsdk.PutMethodInput) error
	postCreate     func(context.Context, *svcapitypes.Method, *svcsdk.Method, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Method, *svcsdk.DeleteMethodInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Method, *svcsdk.DeleteMethodOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.Method, *svcsdk.UpdateMethodInput) error
	postUpdate     func(context.Context, *svcapitypes.Method, *svcsdk.Method, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Method, *svcsdk.GetMethodInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Method, _ *svcsdk.Method, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.MethodParameters, *svcsdk.Method) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.Method, *svcsdk.Method) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.Method, *svcsdk.PutMethodInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Method, _ *svcsdk.Method, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Method, *svcsdk.DeleteMethodInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Method, _ *svcsdk.DeleteMethodOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.Method, *svcsdk.UpdateMethodInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.Method, _ *svcsdk.Method, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
