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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// FlowLogParameters defines the desired state of FlowLog
type FlowLogParameters struct {
	// Region is which region the FlowLog will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to ensure idempotency (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string `json:"clientToken,omitempty"`
	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs across
	// accounts.
	DeliverCrossAccountRole *string `json:"deliverCrossAccountRole,omitempty"`
	// The destination options.
	DestinationOptions *DestinationOptionsRequest `json:"destinationOptions,omitempty"`
	// The destination for the flow log data. The meaning of this parameter depends
	// on the destination type.
	//
	//    * If the destination type is cloud-watch-logs, specify the ARN of a CloudWatch
	//    Logs log group. For example: arn:aws:logs:region:account_id:log-group:my_group
	//    Alternatively, use the LogGroupName parameter.
	//
	//    * If the destination type is s3, specify the ARN of an S3 bucket. For
	//    example: arn:aws:s3:::my_bucket/my_subfolder/ The subfolder is optional.
	//    Note that you can't use AWSLogs as a subfolder name.
	//
	//    * If the destination type is kinesis-data-firehose, specify the ARN of
	//    a Kinesis Data Firehose delivery stream. For example: arn:aws:firehose:region:account_id:deliverystream:my_stream
	LogDestination *string `json:"logDestination,omitempty"`
	// The type of destination for the flow log data.
	//
	// Default: cloud-watch-logs
	LogDestinationType *string `json:"logDestinationType,omitempty"`
	// The fields to include in the flow log record. List the fields in the order
	// in which they should appear. If you omit this parameter, the flow log is
	// created using the default format. If you specify this parameter, you must
	// include at least one field. For more information about the available fields,
	// see Flow log records (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records)
	// in the Amazon VPC User Guide or Transit Gateway Flow Log records (https://docs.aws.amazon.com/vpc/latest/tgw/tgw-flow-logs.html#flow-log-records)
	// in the Amazon Web Services Transit Gateway Guide.
	//
	// Specify the fields using the ${field-id} format, separated by spaces. For
	// the CLI, surround this parameter value with single quotes on Linux or double
	// quotes on Windows.
	LogFormat *string `json:"logFormat,omitempty"`
	// The name of a new or existing CloudWatch Logs log group where Amazon EC2
	// publishes your flow logs.
	//
	// This parameter is valid only if the destination type is cloud-watch-logs.
	LogGroupName *string `json:"logGroupName,omitempty"`
	// The maximum interval of time during which a flow of packets is captured and
	// aggregated into a flow log record. The possible values are 60 seconds (1
	// minute) or 600 seconds (10 minutes). This parameter must be 60 seconds for
	// transit gateway resource types.
	//
	// When a network interface is attached to a Nitro-based instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances),
	// the aggregation interval is always 60 seconds or less, regardless of the
	// value that you specify.
	//
	// Default: 600
	MaxAggregationInterval *int64 `json:"maxAggregationInterval,omitempty"`
	// The type of traffic to monitor (accepted traffic, rejected traffic, or all
	// traffic). This parameter is not supported for transit gateway resource types.
	// It is required for the other resource types.
	TrafficType             *string `json:"trafficType,omitempty"`
	CustomFlowLogParameters `json:",inline"`
}

// FlowLogSpec defines the desired state of FlowLog
type FlowLogSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FlowLogParameters `json:"forProvider"`
}

// FlowLogObservation defines the observed state of FlowLog
type FlowLogObservation struct {
	// The date and time the flow log was created.
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// The status of the logs delivery (SUCCESS | FAILED).
	DeliverLogsStatus *string `json:"deliverLogsStatus,omitempty"`
	// The ID of the flow log.
	FlowLogID *string `json:"flowLogID,omitempty"`
	// The status of the flow log (ACTIVE).
	FlowLogStatus *string `json:"flowLogStatus,omitempty"`
	// The ID of the resource being monitored.
	ResourceID *string `json:"resourceID,omitempty"`
	// The tags for the flow log.
	Tags []*Tag `json:"tags,omitempty"`
}

// FlowLogStatus defines the observed state of FlowLog.
type FlowLogStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FlowLogObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlowLog is the Schema for the FlowLogs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FlowLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlowLogSpec   `json:"spec"`
	Status            FlowLogStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlowLogList contains a list of FlowLogs
type FlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlowLog `json:"items"`
}

// Repository type metadata.
var (
	FlowLogKind             = "FlowLog"
	FlowLogGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlowLogKind}.String()
	FlowLogKindAPIVersion   = FlowLogKind + "." + GroupVersion.String()
	FlowLogGroupVersionKind = GroupVersion.WithKind(FlowLogKind)
)

func init() {
	SchemeBuilder.Register(&FlowLog{}, &FlowLogList{})
}
