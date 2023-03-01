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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type DescribedAccess struct {
	HomeDirectory *string `json:"homeDirectory,omitempty"`

	HomeDirectoryMappings []*HomeDirectoryMapEntry `json:"homeDirectoryMappings,omitempty"`

	HomeDirectoryType *string `json:"homeDirectoryType,omitempty"`

	Policy *string `json:"policy,omitempty"`
	// The full POSIX identity, including user ID (Uid), group ID (Gid), and any
	// secondary groups IDs (SecondaryGids), that controls your users' access to
	// your Amazon EFS file systems. The POSIX permissions that are set on files
	// and directories in your file system determine the level of access your users
	// get when transferring files into and out of your Amazon EFS file systems.
	PosixProfile *PosixProfile `json:"posixProfile,omitempty"`

	Role *string `json:"role,omitempty"`
}

// +kubebuilder:skipversion
type DescribedAgreement struct {
	AccessRole *string `json:"accessRole,omitempty"`

	ARN *string `json:"arn,omitempty"`

	BaseDirectory *string `json:"baseDirectory,omitempty"`

	ServerID *string `json:"serverID,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type DescribedCertificate struct {
	ARN *string `json:"arn,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type DescribedConnector struct {
	AccessRole *string `json:"accessRole,omitempty"`

	ARN *string `json:"arn,omitempty"`

	LoggingRole *string `json:"loggingRole,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	URL *string `json:"url,omitempty"`
}

// +kubebuilder:skipversion
type DescribedExecution struct {
	ExecutionRole *string `json:"executionRole,omitempty"`
	// The full POSIX identity, including user ID (Uid), group ID (Gid), and any
	// secondary groups IDs (SecondaryGids), that controls your users' access to
	// your Amazon EFS file systems. The POSIX permissions that are set on files
	// and directories in your file system determine the level of access your users
	// get when transferring files into and out of your Amazon EFS file systems.
	PosixProfile *PosixProfile `json:"posixProfile,omitempty"`
}

// +kubebuilder:skipversion
type DescribedHostKey struct {
	ARN *string `json:"arn,omitempty"`

	DateImported *metav1.Time `json:"dateImported,omitempty"`

	HostKeyFingerprint *string `json:"hostKeyFingerprint,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type DescribedProfile struct {
	ARN *string `json:"arn,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type DescribedSecurityPolicy struct {
	SecurityPolicyName *string `json:"securityPolicyName,omitempty"`
}

// +kubebuilder:skipversion
type DescribedServer struct {
	ARN *string `json:"arn,omitempty"`

	Certificate *string `json:"certificate,omitempty"`

	Domain *string `json:"domain,omitempty"`
	// The virtual private cloud (VPC) endpoint settings that are configured for
	// your file transfer protocol-enabled server. With a VPC endpoint, you can
	// restrict access to your server and resources only within your VPC. To control
	// incoming internet traffic, invoke the UpdateServer API and attach an Elastic
	// IP address to your server's endpoint.
	//
	// After May 19, 2021, you won't be able to create a server using EndpointType=VPC_ENDPOINT
	// in your Amazon Web Servicesaccount if your account hasn't already done so
	// before May 19, 2021. If you have already created servers with EndpointType=VPC_ENDPOINT
	// in your Amazon Web Servicesaccount on or before May 19, 2021, you will not
	// be affected. After this date, use EndpointType=VPC.
	//
	// For more information, see https://docs.aws.amazon.com/transfer/latest/userguide/create-server-in-vpc.html#deprecate-vpc-endpoint.
	EndpointDetails *EndpointDetails `json:"endpointDetails,omitempty"`

	EndpointType *string `json:"endpointType,omitempty"`

	HostKeyFingerprint *string `json:"hostKeyFingerprint,omitempty"`
	// Returns information related to the type of user authentication that is in
	// use for a file transfer protocol-enabled server's users. A server can have
	// only one method of authentication.
	IdentityProviderDetails *IdentityProviderDetails `json:"identityProviderDetails,omitempty"`
	// Returns information related to the type of user authentication that is in
	// use for a file transfer protocol-enabled server's users. For AWS_DIRECTORY_SERVICE
	// or SERVICE_MANAGED authentication, the Secure Shell (SSH) public keys are
	// stored with a user on the server instance. For API_GATEWAY authentication,
	// your custom authentication method is implemented by using an API call. The
	// server can have only one method of authentication.
	IdentityProviderType *string `json:"identityProviderType,omitempty"`

	LoggingRole *string `json:"loggingRole,omitempty"`

	PostAuthenticationLoginBanner *string `json:"postAuthenticationLoginBanner,omitempty"`

	PreAuthenticationLoginBanner *string `json:"preAuthenticationLoginBanner,omitempty"`
	// The protocol settings that are configured for your server.
	ProtocolDetails *ProtocolDetails `json:"protocolDetails,omitempty"`

	Protocols []*string `json:"protocols,omitempty"`

	SecurityPolicyName *string `json:"securityPolicyName,omitempty"`

	ServerID *string `json:"serverID,omitempty"`
	// Describes the condition of a file transfer protocol-enabled server with respect
	// to its ability to perform file operations. There are six possible states:
	// OFFLINE, ONLINE, STARTING, STOPPING, START_FAILED, and STOP_FAILED.
	//
	// OFFLINE indicates that the server exists, but that it is not available for
	// file operations. ONLINE indicates that the server is available to perform
	// file operations. STARTING indicates that the server's was instantiated, but
	// the server is not yet available to perform file operations. Under normal
	// conditions, it can take a couple of minutes for the server to be completely
	// operational. Both START_FAILED and STOP_FAILED are error conditions.
	State *string `json:"state,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	UserCount *int64 `json:"userCount,omitempty"`
	// Container for the WorkflowDetail data type. It is used by actions that trigger
	// a workflow to begin execution.
	WorkflowDetails *WorkflowDetails `json:"workflowDetails,omitempty"`
}

// +kubebuilder:skipversion
type DescribedUser struct {
	ARN *string `json:"arn,omitempty"`

	HomeDirectory *string `json:"homeDirectory,omitempty"`

	HomeDirectoryMappings []*HomeDirectoryMapEntry `json:"homeDirectoryMappings,omitempty"`

	HomeDirectoryType *string `json:"homeDirectoryType,omitempty"`

	Policy *string `json:"policy,omitempty"`
	// The full POSIX identity, including user ID (Uid), group ID (Gid), and any
	// secondary groups IDs (SecondaryGids), that controls your users' access to
	// your Amazon EFS file systems. The POSIX permissions that are set on files
	// and directories in your file system determine the level of access your users
	// get when transferring files into and out of your Amazon EFS file systems.
	PosixProfile *PosixProfile `json:"posixProfile,omitempty"`

	Role *string `json:"role,omitempty"`

	SshPublicKeys []*SshPublicKey `json:"sshPublicKeys,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	UserName *string `json:"userName,omitempty"`
}

// +kubebuilder:skipversion
type DescribedWorkflow struct {
	ARN *string `json:"arn,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	WorkflowID *string `json:"workflowID,omitempty"`
}

// +kubebuilder:skipversion
type EndpointDetails struct {
	AddressAllocationIDs []*string `json:"addressAllocationIDs,omitempty"`

	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`

	SubnetIDs []*string `json:"subnetIDs,omitempty"`

	VPCEndpointID *string `json:"vpcEndpointID,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}

// +kubebuilder:skipversion
type HomeDirectoryMapEntry struct {
	Entry *string `json:"entry,omitempty"`

	Target *string `json:"target,omitempty"`
}

// +kubebuilder:skipversion
type IdentityProviderDetails struct {
	DirectoryID *string `json:"directoryID,omitempty"`

	Function *string `json:"function,omitempty"`

	InvocationRole *string `json:"invocationRole,omitempty"`

	URL *string `json:"url,omitempty"`
}

// +kubebuilder:skipversion
type ListedAccess struct {
	HomeDirectory *string `json:"homeDirectory,omitempty"`

	HomeDirectoryType *string `json:"homeDirectoryType,omitempty"`

	Role *string `json:"role,omitempty"`
}

// +kubebuilder:skipversion
type ListedAgreement struct {
	ARN *string `json:"arn,omitempty"`

	ServerID *string `json:"serverID,omitempty"`
}

// +kubebuilder:skipversion
type ListedCertificate struct {
	ARN *string `json:"arn,omitempty"`
}

// +kubebuilder:skipversion
type ListedConnector struct {
	ARN *string `json:"arn,omitempty"`

	URL *string `json:"url,omitempty"`
}

// +kubebuilder:skipversion
type ListedHostKey struct {
	ARN *string `json:"arn,omitempty"`

	DateImported *metav1.Time `json:"dateImported,omitempty"`

	Fingerprint *string `json:"fingerprint,omitempty"`
}

// +kubebuilder:skipversion
type ListedProfile struct {
	ARN *string `json:"arn,omitempty"`
}

// +kubebuilder:skipversion
type ListedServer struct {
	ARN *string `json:"arn,omitempty"`

	Domain *string `json:"domain,omitempty"`

	EndpointType *string `json:"endpointType,omitempty"`
	// Returns information related to the type of user authentication that is in
	// use for a file transfer protocol-enabled server's users. For AWS_DIRECTORY_SERVICE
	// or SERVICE_MANAGED authentication, the Secure Shell (SSH) public keys are
	// stored with a user on the server instance. For API_GATEWAY authentication,
	// your custom authentication method is implemented by using an API call. The
	// server can have only one method of authentication.
	IdentityProviderType *string `json:"identityProviderType,omitempty"`

	LoggingRole *string `json:"loggingRole,omitempty"`

	ServerID *string `json:"serverID,omitempty"`
	// Describes the condition of a file transfer protocol-enabled server with respect
	// to its ability to perform file operations. There are six possible states:
	// OFFLINE, ONLINE, STARTING, STOPPING, START_FAILED, and STOP_FAILED.
	//
	// OFFLINE indicates that the server exists, but that it is not available for
	// file operations. ONLINE indicates that the server is available to perform
	// file operations. STARTING indicates that the server's was instantiated, but
	// the server is not yet available to perform file operations. Under normal
	// conditions, it can take a couple of minutes for the server to be completely
	// operational. Both START_FAILED and STOP_FAILED are error conditions.
	State *string `json:"state,omitempty"`

	UserCount *int64 `json:"userCount,omitempty"`
}

// +kubebuilder:skipversion
type ListedUser struct {
	ARN *string `json:"arn,omitempty"`

	HomeDirectory *string `json:"homeDirectory,omitempty"`

	HomeDirectoryType *string `json:"homeDirectoryType,omitempty"`

	Role *string `json:"role,omitempty"`

	SshPublicKeyCount *int64 `json:"sshPublicKeyCount,omitempty"`

	UserName *string `json:"userName,omitempty"`
}

// +kubebuilder:skipversion
type ListedWorkflow struct {
	ARN *string `json:"arn,omitempty"`

	WorkflowID *string `json:"workflowID,omitempty"`
}

// +kubebuilder:skipversion
type LoggingConfiguration struct {
	LoggingRole *string `json:"loggingRole,omitempty"`
}

// +kubebuilder:skipversion
type PosixProfile struct {
	Gid *int64 `json:"gid,omitempty"`

	SecondaryGids []*int64 `json:"secondaryGids,omitempty"`

	Uid *int64 `json:"uid,omitempty"`
}

// +kubebuilder:skipversion
type ProtocolDetails struct {
	As2Transports []*string `json:"as2Transports,omitempty"`

	PassiveIP *string `json:"passiveIP,omitempty"`

	SetStatOption *string `json:"setStatOption,omitempty"`

	TLSSessionResumptionMode *string `json:"tlsSessionResumptionMode,omitempty"`
}

// +kubebuilder:skipversion
type SshPublicKey struct {
	DateImported *metav1.Time `json:"dateImported,omitempty"`

	SshPublicKeyBody *string `json:"sshPublicKeyBody,omitempty"`

	SshPublicKeyID *string `json:"sshPublicKeyID,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type UserDetails struct {
	ServerID *string `json:"serverID,omitempty"`

	UserName *string `json:"userName,omitempty"`
}

// +kubebuilder:skipversion
type WorkflowDetail struct {
	ExecutionRole *string `json:"executionRole,omitempty"`

	WorkflowID *string `json:"workflowID,omitempty"`
}

// +kubebuilder:skipversion
type WorkflowDetails struct {
	OnPartialUpload []*WorkflowDetail `json:"onPartialUpload,omitempty"`

	OnUpload []*WorkflowDetail `json:"onUpload,omitempty"`
}
