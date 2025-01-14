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

type AcceleratorManufacturer string

const (
	AcceleratorManufacturer_nvidia              AcceleratorManufacturer = "nvidia"
	AcceleratorManufacturer_amd                 AcceleratorManufacturer = "amd"
	AcceleratorManufacturer_amazon_web_services AcceleratorManufacturer = "amazon-web-services"
	AcceleratorManufacturer_xilinx              AcceleratorManufacturer = "xilinx"
)

type AcceleratorName string

const (
	AcceleratorName_a100            AcceleratorName = "a100"
	AcceleratorName_v100            AcceleratorName = "v100"
	AcceleratorName_k80             AcceleratorName = "k80"
	AcceleratorName_t4              AcceleratorName = "t4"
	AcceleratorName_m60             AcceleratorName = "m60"
	AcceleratorName_radeon_pro_v520 AcceleratorName = "radeon-pro-v520"
	AcceleratorName_vu9p            AcceleratorName = "vu9p"
)

type AcceleratorType string

const (
	AcceleratorType_gpu       AcceleratorType = "gpu"
	AcceleratorType_fpga      AcceleratorType = "fpga"
	AcceleratorType_inference AcceleratorType = "inference"
)

type BareMetal string

const (
	BareMetal_included BareMetal = "included"
	BareMetal_excluded BareMetal = "excluded"
	BareMetal_required BareMetal = "required"
)

type BurstablePerformance string

const (
	BurstablePerformance_included BurstablePerformance = "included"
	BurstablePerformance_excluded BurstablePerformance = "excluded"
	BurstablePerformance_required BurstablePerformance = "required"
)

type CPUManufacturer string

const (
	CPUManufacturer_intel               CPUManufacturer = "intel"
	CPUManufacturer_amd                 CPUManufacturer = "amd"
	CPUManufacturer_amazon_web_services CPUManufacturer = "amazon-web-services"
)

type InstanceGeneration string

const (
	InstanceGeneration_current  InstanceGeneration = "current"
	InstanceGeneration_previous InstanceGeneration = "previous"
)

type InstanceMetadataEndpointState string

const (
	InstanceMetadataEndpointState_disabled InstanceMetadataEndpointState = "disabled"
	InstanceMetadataEndpointState_enabled  InstanceMetadataEndpointState = "enabled"
)

type InstanceMetadataHTTPTokensState string

const (
	InstanceMetadataHTTPTokensState_optional InstanceMetadataHTTPTokensState = "optional"
	InstanceMetadataHTTPTokensState_required InstanceMetadataHTTPTokensState = "required"
)

type InstanceRefreshStatus string

const (
	InstanceRefreshStatus_Pending    InstanceRefreshStatus = "Pending"
	InstanceRefreshStatus_InProgress InstanceRefreshStatus = "InProgress"
	InstanceRefreshStatus_Successful InstanceRefreshStatus = "Successful"
	InstanceRefreshStatus_Failed     InstanceRefreshStatus = "Failed"
	InstanceRefreshStatus_Cancelling InstanceRefreshStatus = "Cancelling"
	InstanceRefreshStatus_Cancelled  InstanceRefreshStatus = "Cancelled"
)

type LifecycleState string

const (
	LifecycleState_Pending                    LifecycleState = "Pending"
	LifecycleState_Pending_Wait               LifecycleState = "Pending:Wait"
	LifecycleState_Pending_Proceed            LifecycleState = "Pending:Proceed"
	LifecycleState_Quarantined                LifecycleState = "Quarantined"
	LifecycleState_InService                  LifecycleState = "InService"
	LifecycleState_Terminating                LifecycleState = "Terminating"
	LifecycleState_Terminating_Wait           LifecycleState = "Terminating:Wait"
	LifecycleState_Terminating_Proceed        LifecycleState = "Terminating:Proceed"
	LifecycleState_Terminated                 LifecycleState = "Terminated"
	LifecycleState_Detaching                  LifecycleState = "Detaching"
	LifecycleState_Detached                   LifecycleState = "Detached"
	LifecycleState_EnteringStandby            LifecycleState = "EnteringStandby"
	LifecycleState_Standby                    LifecycleState = "Standby"
	LifecycleState_Warmed_Pending             LifecycleState = "Warmed:Pending"
	LifecycleState_Warmed_Pending_Wait        LifecycleState = "Warmed:Pending:Wait"
	LifecycleState_Warmed_Pending_Proceed     LifecycleState = "Warmed:Pending:Proceed"
	LifecycleState_Warmed_Terminating         LifecycleState = "Warmed:Terminating"
	LifecycleState_Warmed_Terminating_Wait    LifecycleState = "Warmed:Terminating:Wait"
	LifecycleState_Warmed_Terminating_Proceed LifecycleState = "Warmed:Terminating:Proceed"
	LifecycleState_Warmed_Terminated          LifecycleState = "Warmed:Terminated"
	LifecycleState_Warmed_Stopped             LifecycleState = "Warmed:Stopped"
	LifecycleState_Warmed_Running             LifecycleState = "Warmed:Running"
	LifecycleState_Warmed_Hibernated          LifecycleState = "Warmed:Hibernated"
)

type LocalStorage string

const (
	LocalStorage_included LocalStorage = "included"
	LocalStorage_excluded LocalStorage = "excluded"
	LocalStorage_required LocalStorage = "required"
)

type LocalStorageType string

const (
	LocalStorageType_hdd LocalStorageType = "hdd"
	LocalStorageType_ssd LocalStorageType = "ssd"
)

type MetricStatistic string

const (
	MetricStatistic_Average     MetricStatistic = "Average"
	MetricStatistic_Minimum     MetricStatistic = "Minimum"
	MetricStatistic_Maximum     MetricStatistic = "Maximum"
	MetricStatistic_SampleCount MetricStatistic = "SampleCount"
	MetricStatistic_Sum         MetricStatistic = "Sum"
)

type MetricType string

const (
	MetricType_ASGAverageCPUUtilization MetricType = "ASGAverageCPUUtilization"
	MetricType_ASGAverageNetworkIn      MetricType = "ASGAverageNetworkIn"
	MetricType_ASGAverageNetworkOut     MetricType = "ASGAverageNetworkOut"
	MetricType_ALBRequestCountPerTarget MetricType = "ALBRequestCountPerTarget"
)

type PredefinedLoadMetricType string

const (
	PredefinedLoadMetricType_ASGTotalCPUUtilization     PredefinedLoadMetricType = "ASGTotalCPUUtilization"
	PredefinedLoadMetricType_ASGTotalNetworkIn          PredefinedLoadMetricType = "ASGTotalNetworkIn"
	PredefinedLoadMetricType_ASGTotalNetworkOut         PredefinedLoadMetricType = "ASGTotalNetworkOut"
	PredefinedLoadMetricType_ALBTargetGroupRequestCount PredefinedLoadMetricType = "ALBTargetGroupRequestCount"
)

type PredefinedMetricPairType string

const (
	PredefinedMetricPairType_ASGCPUUtilization PredefinedMetricPairType = "ASGCPUUtilization"
	PredefinedMetricPairType_ASGNetworkIn      PredefinedMetricPairType = "ASGNetworkIn"
	PredefinedMetricPairType_ASGNetworkOut     PredefinedMetricPairType = "ASGNetworkOut"
	PredefinedMetricPairType_ALBRequestCount   PredefinedMetricPairType = "ALBRequestCount"
)

type PredefinedScalingMetricType string

const (
	PredefinedScalingMetricType_ASGAverageCPUUtilization PredefinedScalingMetricType = "ASGAverageCPUUtilization"
	PredefinedScalingMetricType_ASGAverageNetworkIn      PredefinedScalingMetricType = "ASGAverageNetworkIn"
	PredefinedScalingMetricType_ASGAverageNetworkOut     PredefinedScalingMetricType = "ASGAverageNetworkOut"
	PredefinedScalingMetricType_ALBRequestCountPerTarget PredefinedScalingMetricType = "ALBRequestCountPerTarget"
)

type PredictiveScalingMaxCapacityBreachBehavior string

const (
	PredictiveScalingMaxCapacityBreachBehavior_HonorMaxCapacity    PredictiveScalingMaxCapacityBreachBehavior = "HonorMaxCapacity"
	PredictiveScalingMaxCapacityBreachBehavior_IncreaseMaxCapacity PredictiveScalingMaxCapacityBreachBehavior = "IncreaseMaxCapacity"
)

type PredictiveScalingMode string

const (
	PredictiveScalingMode_ForecastAndScale PredictiveScalingMode = "ForecastAndScale"
	PredictiveScalingMode_ForecastOnly     PredictiveScalingMode = "ForecastOnly"
)

type RefreshStrategy string

const (
	RefreshStrategy_Rolling RefreshStrategy = "Rolling"
)

type ScalingActivityStatusCode string

const (
	ScalingActivityStatusCode_PendingSpotBidPlacement         ScalingActivityStatusCode = "PendingSpotBidPlacement"
	ScalingActivityStatusCode_WaitingForSpotInstanceRequestId ScalingActivityStatusCode = "WaitingForSpotInstanceRequestId"
	ScalingActivityStatusCode_WaitingForSpotInstanceId        ScalingActivityStatusCode = "WaitingForSpotInstanceId"
	ScalingActivityStatusCode_WaitingForInstanceId            ScalingActivityStatusCode = "WaitingForInstanceId"
	ScalingActivityStatusCode_PreInService                    ScalingActivityStatusCode = "PreInService"
	ScalingActivityStatusCode_InProgress                      ScalingActivityStatusCode = "InProgress"
	ScalingActivityStatusCode_WaitingForELBConnectionDraining ScalingActivityStatusCode = "WaitingForELBConnectionDraining"
	ScalingActivityStatusCode_MidLifecycleAction              ScalingActivityStatusCode = "MidLifecycleAction"
	ScalingActivityStatusCode_WaitingForInstanceWarmup        ScalingActivityStatusCode = "WaitingForInstanceWarmup"
	ScalingActivityStatusCode_Successful                      ScalingActivityStatusCode = "Successful"
	ScalingActivityStatusCode_Failed                          ScalingActivityStatusCode = "Failed"
	ScalingActivityStatusCode_Cancelled                       ScalingActivityStatusCode = "Cancelled"
)

type WarmPoolState string

const (
	WarmPoolState_Stopped    WarmPoolState = "Stopped"
	WarmPoolState_Running    WarmPoolState = "Running"
	WarmPoolState_Hibernated WarmPoolState = "Hibernated"
)

type WarmPoolStatus string

const (
	WarmPoolStatus_PendingDelete WarmPoolStatus = "PendingDelete"
)
