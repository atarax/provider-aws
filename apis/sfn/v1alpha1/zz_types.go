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
type ActivityListItem struct {
	ActivityARN *string `json:"activityARN,omitempty"`

	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type ActivityScheduledEventDetails struct {
	Resource *string `json:"resource,omitempty"`
}

// +kubebuilder:skipversion
type CloudWatchLogsLogGroup struct {
	LogGroupARN *string `json:"logGroupARN,omitempty"`
}

// +kubebuilder:skipversion
type ExecutionListItem struct {
	ExecutionARN *string `json:"executionARN,omitempty"`

	Name *string `json:"name,omitempty"`

	StartDate *metav1.Time `json:"startDate,omitempty"`

	StateMachineARN *string `json:"stateMachineARN,omitempty"`

	StopDate *metav1.Time `json:"stopDate,omitempty"`
}

// +kubebuilder:skipversion
type ExecutionStartedEventDetails struct {
	RoleARN *string `json:"roleARN,omitempty"`
}

// +kubebuilder:skipversion
type HistoryEvent struct {
	Timestamp *metav1.Time `json:"timestamp,omitempty"`
}

// +kubebuilder:skipversion
type LambdaFunctionScheduledEventDetails struct {
	Resource *string `json:"resource,omitempty"`
}

// +kubebuilder:skipversion
type LogDestination struct {
	CloudWatchLogsLogGroup *CloudWatchLogsLogGroup `json:"cloudWatchLogsLogGroup,omitempty"`
}

// +kubebuilder:skipversion
type LoggingConfiguration struct {
	Destinations []*LogDestination `json:"destinations,omitempty"`

	IncludeExecutionData *bool `json:"includeExecutionData,omitempty"`

	Level *string `json:"level,omitempty"`
}

// +kubebuilder:skipversion
type MapIterationEventDetails struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type MapRunListItem struct {
	ExecutionARN *string `json:"executionARN,omitempty"`

	StartDate *metav1.Time `json:"startDate,omitempty"`

	StateMachineARN *string `json:"stateMachineARN,omitempty"`

	StopDate *metav1.Time `json:"stopDate,omitempty"`
}

// +kubebuilder:skipversion
type StateEnteredEventDetails struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type StateExitedEventDetails struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type StateMachineListItem struct {
	CreationDate *metav1.Time `json:"creationDate,omitempty"`

	Name *string `json:"name,omitempty"`

	StateMachineARN *string `json:"stateMachineARN,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type TaskFailedEventDetails struct {
	Resource *string `json:"resource,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`
}

// +kubebuilder:skipversion
type TaskScheduledEventDetails struct {
	Region *string `json:"region,omitempty"`

	Resource *string `json:"resource,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`
}

// +kubebuilder:skipversion
type TaskStartFailedEventDetails struct {
	Resource *string `json:"resource,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`
}

// +kubebuilder:skipversion
type TaskStartedEventDetails struct {
	Resource *string `json:"resource,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`
}

// +kubebuilder:skipversion
type TaskSubmitFailedEventDetails struct {
	Resource *string `json:"resource,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`
}

// +kubebuilder:skipversion
type TaskSubmittedEventDetails struct {
	Resource *string `json:"resource,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`
}

// +kubebuilder:skipversion
type TaskSucceededEventDetails struct {
	Resource *string `json:"resource,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`
}

// +kubebuilder:skipversion
type TaskTimedOutEventDetails struct {
	Resource *string `json:"resource,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`
}

// +kubebuilder:skipversion
type TracingConfiguration struct {
	Enabled *bool `json:"enabled,omitempty"`
}
