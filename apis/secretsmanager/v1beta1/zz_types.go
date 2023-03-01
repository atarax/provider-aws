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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)



// +kubebuilder:skipversion
type Filter struct {
	
	Key *string `json:"key,omitempty"`
	
	Values []*string `json:"values,omitempty"`
}

// +kubebuilder:skipversion
type ReplicaRegionType struct {
	
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	
	Region *string `json:"region,omitempty"`
}

// +kubebuilder:skipversion
type ReplicationStatusType struct {
	
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	
	LastAccessedDate *metav1.Time `json:"lastAccessedDate,omitempty"`
	
	Region *string `json:"region,omitempty"`
	
	Status *string `json:"status,omitempty"`
	
	StatusMessage *string `json:"statusMessage,omitempty"`
}

// +kubebuilder:skipversion
type RotationRulesType struct {
	
	AutomaticallyAfterDays *int64 `json:"automaticallyAfterDays,omitempty"`
	
	Duration *string `json:"duration,omitempty"`
	
	ScheduleExpression *string `json:"scheduleExpression,omitempty"`
}

// +kubebuilder:skipversion
type SecretListEntry struct {
	
	ARN *string `json:"arn,omitempty"`
	
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`
	
	DeletedDate *metav1.Time `json:"deletedDate,omitempty"`
	
	Description *string `json:"description,omitempty"`
	
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	
	LastAccessedDate *metav1.Time `json:"lastAccessedDate,omitempty"`
	
	LastChangedDate *metav1.Time `json:"lastChangedDate,omitempty"`
	
	LastRotatedDate *metav1.Time `json:"lastRotatedDate,omitempty"`
	
	Name *string `json:"name,omitempty"`
	
	NextRotationDate *metav1.Time `json:"nextRotationDate,omitempty"`
	
	OwningService *string `json:"owningService,omitempty"`
	
	PrimaryRegion *string `json:"primaryRegion,omitempty"`
	
	RotationEnabled *bool `json:"rotationEnabled,omitempty"`
	
	RotationLambdaARN *string `json:"rotationLambdaARN,omitempty"`
	// A structure that defines the rotation configuration for the secret.
	RotationRules *RotationRulesType `json:"rotationRules,omitempty"`
	
	SecretVersionsToStages map[string][]*string `json:"secretVersionsToStages,omitempty"`
	
	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type SecretVersionsListEntry struct {
	
	LastAccessedDate *metav1.Time `json:"lastAccessedDate,omitempty"`
	
	VersionStages []*string `json:"versionStages,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	
	Key *string `json:"key,omitempty"`
	
	Value *string `json:"value,omitempty"`
}