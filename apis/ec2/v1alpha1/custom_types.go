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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// +kubebuilder:storageversion

// CustomVolumeParameters contains the additional fields for VolumeParameters.
type CustomVolumeParameters struct {
	// The identifier of the AWS Key Management Service (AWS KMS) customer master
	// key (CMK) to use for Amazon EBS encryption. If this parameter is not specified,
	// your AWS managed CMK for EBS is used. If KmsKeyId is specified, the encrypted
	// state must be true.
	//
	// You can specify the CMK using any of the following:
	//
	//    * Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//    * Key alias. For example, alias/ExampleAlias.
	//
	//    * Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//    * Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS authenticates the CMK asynchronously. Therefore, if you specify an ID,
	// alias, or ARN that is not valid, the action can appear to complete, but eventually
	// fails.
	// +optional
	// +crossplane:generate:reference:type=github.com/crossplane/provider-aws/apis/kms/v1alpha1.Key
	// +crossplane:generate:reference:refFieldName=KMSKeyIDRef
	// +crossplane:generate:reference:selectorFieldName=KMSKeyIDSelector
	KMSKeyID *string `json:"kmsKeyId,omitempty"`

	// KMSKeyIDRef is a reference to a KMS Key used to set KMSKeyID.
	// +optional
	KMSKeyIDRef *xpv1.Reference `json:"kmsKeyIdRef,omitempty"`

	// KMSKeyIDSelector selects a reference to a KMS Key used to set KMSKeyID.
	// +optional
	KMSKeyIDSelector *xpv1.Selector `json:"kmsKeyIdSelector,omitempty"`
}

// CustomVPCPeeringConnectionParameters are custom parameters for VPCPeeringConnection
type CustomVPCPeeringConnectionParameters struct {
	// The ID of the requester VPC. You must specify this parameter in the request.
	VPCID *string `json:"vpcID,omitempty"`
	// VPCIDRef is a reference to an API used to set
	// the VPCID.
	// +optional
	VPCIDRef *xpv1.Reference `json:"vpcIDRef,omitempty"`
	// VPCIDSelector selects references to API used
	// to set the VPCID.
	// +optional
	VPCIDSelector *xpv1.Selector `json:"vpcIDSelector,omitempty"`
	// The ID of the VPC with which you are creating the VPC peering connection.
	// You must specify this parameter in the request.
	PeerVPCID *string `json:"peerVPCID,omitempty"`
	// PeerVPCIDRef is a reference to an API used to set
	// the PeerVPCID.
	// +optional
	PeerVPCIDRef *xpv1.Reference `json:"peerVPCIDRef,omitempty"`
	// PeerVPCIDSelector selects references to API used
	// to set the PeerVPCID.
	// +optional
	PeerVPCIDSelector *xpv1.Selector `json:"peerVPCIDSelector,omitempty"`
	// Automatically accepts the peering connection. If this is not set, the peering connection
	// will be created, but will be in pending-acceptance state. This will only lead to an active
	// connection if both VPCs are in the same tenant.
	AcceptRequest bool `json:"acceptRequest,omitempty"`
}
