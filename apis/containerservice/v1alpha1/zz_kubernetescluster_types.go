/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AciConnectorLinuxObservation struct {
}

type AciConnectorLinuxParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetName *string `json:"subnetName,omitempty" tf:"subnet_name,omitempty"`
}

type AddonProfileObservation struct {
}

type AddonProfileParameters struct {

	// +kubebuilder:validation:Optional
	AciConnectorLinux []AciConnectorLinuxParameters `json:"aciConnectorLinux,omitempty" tf:"aci_connector_linux,omitempty"`

	// +kubebuilder:validation:Optional
	AzurePolicy []AzurePolicyParameters `json:"azurePolicy,omitempty" tf:"azure_policy,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPApplicationRouting []HTTPApplicationRoutingParameters `json:"httpApplicationRouting,omitempty" tf:"http_application_routing,omitempty"`

	// +kubebuilder:validation:Optional
	IngressApplicationGateway []IngressApplicationGatewayParameters `json:"ingressApplicationGateway,omitempty" tf:"ingress_application_gateway,omitempty"`

	// +kubebuilder:validation:Optional
	KubeDashboard []KubeDashboardParameters `json:"kubeDashboard,omitempty" tf:"kube_dashboard,omitempty"`

	// +kubebuilder:validation:Optional
	OmsAgent []OmsAgentParameters `json:"omsAgent,omitempty" tf:"oms_agent,omitempty"`
}

type AllowedObservation struct {
}

type AllowedParameters struct {

	// +kubebuilder:validation:Required
	Day *string `json:"day" tf:"day,omitempty"`

	// +kubebuilder:validation:Required
	Hours []*int64 `json:"hours" tf:"hours,omitempty"`
}

type AutoScalerProfileObservation struct {
}

type AutoScalerProfileParameters struct {

	// +kubebuilder:validation:Optional
	BalanceSimilarNodeGroups *bool `json:"balanceSimilarNodeGroups,omitempty" tf:"balance_similar_node_groups,omitempty"`

	// +kubebuilder:validation:Optional
	EmptyBulkDeleteMax *string `json:"emptyBulkDeleteMax,omitempty" tf:"empty_bulk_delete_max,omitempty"`

	// +kubebuilder:validation:Optional
	Expander *string `json:"expander,omitempty" tf:"expander,omitempty"`

	// +kubebuilder:validation:Optional
	MaxGracefulTerminationSec *string `json:"maxGracefulTerminationSec,omitempty" tf:"max_graceful_termination_sec,omitempty"`

	// +kubebuilder:validation:Optional
	MaxNodeProvisioningTime *string `json:"maxNodeProvisioningTime,omitempty" tf:"max_node_provisioning_time,omitempty"`

	// +kubebuilder:validation:Optional
	MaxUnreadyNodes *int64 `json:"maxUnreadyNodes,omitempty" tf:"max_unready_nodes,omitempty"`

	// +kubebuilder:validation:Optional
	MaxUnreadyPercentage *float64 `json:"maxUnreadyPercentage,omitempty" tf:"max_unready_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	NewPodScaleUpDelay *string `json:"newPodScaleUpDelay,omitempty" tf:"new_pod_scale_up_delay,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleDownDelayAfterAdd *string `json:"scaleDownDelayAfterAdd,omitempty" tf:"scale_down_delay_after_add,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleDownDelayAfterDelete *string `json:"scaleDownDelayAfterDelete,omitempty" tf:"scale_down_delay_after_delete,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleDownDelayAfterFailure *string `json:"scaleDownDelayAfterFailure,omitempty" tf:"scale_down_delay_after_failure,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleDownUnneeded *string `json:"scaleDownUnneeded,omitempty" tf:"scale_down_unneeded,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleDownUnready *string `json:"scaleDownUnready,omitempty" tf:"scale_down_unready,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleDownUtilizationThreshold *string `json:"scaleDownUtilizationThreshold,omitempty" tf:"scale_down_utilization_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	ScanInterval *string `json:"scanInterval,omitempty" tf:"scan_interval,omitempty"`

	// +kubebuilder:validation:Optional
	SkipNodesWithLocalStorage *bool `json:"skipNodesWithLocalStorage,omitempty" tf:"skip_nodes_with_local_storage,omitempty"`

	// +kubebuilder:validation:Optional
	SkipNodesWithSystemPods *bool `json:"skipNodesWithSystemPods,omitempty" tf:"skip_nodes_with_system_pods,omitempty"`
}

type AzureActiveDirectoryObservation struct {
}

type AzureActiveDirectoryParameters struct {

	// +kubebuilder:validation:Optional
	AdminGroupObjectIds []*string `json:"adminGroupObjectIds,omitempty" tf:"admin_group_object_ids,omitempty"`

	// +kubebuilder:validation:Optional
	AzureRbacEnabled *bool `json:"azureRbacEnabled,omitempty" tf:"azure_rbac_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ClientAppID *string `json:"clientAppId,omitempty" tf:"client_app_id,omitempty"`

	// +kubebuilder:validation:Optional
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// +kubebuilder:validation:Optional
	ServerAppID *string `json:"serverAppId,omitempty" tf:"server_app_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServerAppSecretSecretRef *v1.SecretKeySelector `json:"serverAppSecretSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzurePolicyObservation struct {
}

type AzurePolicyParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type DefaultNodePoolObservation struct {
}

type DefaultNodePoolParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty" tf:"enable_auto_scaling,omitempty"`

	// +kubebuilder:validation:Optional
	EnableHostEncryption *bool `json:"enableHostEncryption,omitempty" tf:"enable_host_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	EnableNodePublicIP *bool `json:"enableNodePublicIp,omitempty" tf:"enable_node_public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	FipsEnabled *bool `json:"fipsEnabled,omitempty" tf:"fips_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KubeletConfig []KubeletConfigParameters `json:"kubeletConfig,omitempty" tf:"kubelet_config,omitempty"`

	// +kubebuilder:validation:Optional
	KubeletDiskType *string `json:"kubeletDiskType,omitempty" tf:"kubelet_disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	LinuxOsConfig []LinuxOsConfigParameters `json:"linuxOsConfig,omitempty" tf:"linux_os_config,omitempty"`

	// +kubebuilder:validation:Optional
	MaxCount *int64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`

	// +kubebuilder:validation:Optional
	MaxPods *int64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`

	// +kubebuilder:validation:Optional
	MinCount *int64 `json:"minCount,omitempty" tf:"min_count,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// +kubebuilder:validation:Optional
	NodeLabels map[string]*string `json:"nodeLabels,omitempty" tf:"node_labels,omitempty"`

	// +kubebuilder:validation:Optional
	NodePublicIPPrefixID *string `json:"nodePublicIpPrefixId,omitempty" tf:"node_public_ip_prefix_id,omitempty"`

	// +kubebuilder:validation:Optional
	NodeTaints []*string `json:"nodeTaints,omitempty" tf:"node_taints,omitempty"`

	// +kubebuilder:validation:Optional
	OnlyCriticalAddonsEnabled *bool `json:"onlyCriticalAddonsEnabled,omitempty" tf:"only_critical_addons_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	OrchestratorVersion *string `json:"orchestratorVersion,omitempty" tf:"orchestrator_version,omitempty"`

	// +kubebuilder:validation:Optional
	OsDiskSizeGb *int64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	OsSku *string `json:"osSku,omitempty" tf:"os_sku,omitempty"`

	// +kubebuilder:validation:Optional
	PodSubnetID *string `json:"podSubnetId,omitempty" tf:"pod_subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	UpgradeSettings []UpgradeSettingsParameters `json:"upgradeSettings,omitempty" tf:"upgrade_settings,omitempty"`

	// +kubebuilder:validation:Required
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`

	// +kubebuilder:validation:Optional
	VnetSubnetID *string `json:"vnetSubnetId,omitempty" tf:"vnet_subnet_id,omitempty"`
}

type HTTPApplicationRoutingObservation struct {
	HTTPApplicationRoutingZoneName *string `json:"httpApplicationRoutingZoneName,omitempty" tf:"http_application_routing_zone_name,omitempty"`
}

type HTTPApplicationRoutingParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type IngressApplicationGatewayIdentityObservation struct {
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type IngressApplicationGatewayIdentityParameters struct {
}

type IngressApplicationGatewayObservation struct {
	EffectiveGatewayID *string `json:"effectiveGatewayId,omitempty" tf:"effective_gateway_id,omitempty"`

	IngressApplicationGatewayIdentity []IngressApplicationGatewayIdentityObservation `json:"ingressApplicationGatewayIdentity,omitempty" tf:"ingress_application_gateway_identity,omitempty"`
}

type IngressApplicationGatewayParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayName *string `json:"gatewayName,omitempty" tf:"gateway_name,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetCidr *string `json:"subnetCidr,omitempty" tf:"subnet_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type KubeAdminConfigObservation struct {
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type KubeAdminConfigParameters struct {
}

type KubeConfigObservation struct {
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type KubeConfigParameters struct {
}

type KubeDashboardObservation struct {
}

type KubeDashboardParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type KubeletConfigObservation struct {
}

type KubeletConfigParameters struct {

	// +kubebuilder:validation:Optional
	AllowedUnsafeSysctls []*string `json:"allowedUnsafeSysctls,omitempty" tf:"allowed_unsafe_sysctls,omitempty"`

	// +kubebuilder:validation:Optional
	CPUCfsQuotaEnabled *bool `json:"cpuCfsQuotaEnabled,omitempty" tf:"cpu_cfs_quota_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	CPUCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty" tf:"cpu_cfs_quota_period,omitempty"`

	// +kubebuilder:validation:Optional
	CPUManagerPolicy *string `json:"cpuManagerPolicy,omitempty" tf:"cpu_manager_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ContainerLogMaxLine *int64 `json:"containerLogMaxLine,omitempty" tf:"container_log_max_line,omitempty"`

	// +kubebuilder:validation:Optional
	ContainerLogMaxSizeMb *int64 `json:"containerLogMaxSizeMb,omitempty" tf:"container_log_max_size_mb,omitempty"`

	// +kubebuilder:validation:Optional
	ImageGcHighThreshold *int64 `json:"imageGcHighThreshold,omitempty" tf:"image_gc_high_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	ImageGcLowThreshold *int64 `json:"imageGcLowThreshold,omitempty" tf:"image_gc_low_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	PodMaxPid *int64 `json:"podMaxPid,omitempty" tf:"pod_max_pid,omitempty"`

	// +kubebuilder:validation:Optional
	TopologyManagerPolicy *string `json:"topologyManagerPolicy,omitempty" tf:"topology_manager_policy,omitempty"`
}

type KubeletIdentityObservation struct {
}

type KubeletIdentityParameters struct {

	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// +kubebuilder:validation:Optional
	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type KubernetesClusterObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	KubeAdminConfig []KubeAdminConfigObservation `json:"kubeAdminConfig,omitempty" tf:"kube_admin_config,omitempty"`

	KubeConfig []KubeConfigObservation `json:"kubeConfig,omitempty" tf:"kube_config,omitempty"`

	PrivateFqdn *string `json:"privateFqdn,omitempty" tf:"private_fqdn,omitempty"`
}

type KubernetesClusterParameters struct {

	// +kubebuilder:validation:Optional
	APIServerAuthorizedIPRanges []*string `json:"apiServerAuthorizedIpRanges,omitempty" tf:"api_server_authorized_ip_ranges,omitempty"`

	// +kubebuilder:validation:Optional
	AddonProfile []AddonProfileParameters `json:"addonProfile,omitempty" tf:"addon_profile,omitempty"`

	// +kubebuilder:validation:Optional
	AutoScalerProfile []AutoScalerProfileParameters `json:"autoScalerProfile,omitempty" tf:"auto_scaler_profile,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticChannelUpgrade *string `json:"automaticChannelUpgrade,omitempty" tf:"automatic_channel_upgrade,omitempty"`

	// +kubebuilder:validation:Optional
	DNSPrefix *string `json:"dnsPrefix,omitempty" tf:"dns_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	DNSPrefixPrivateCluster *string `json:"dnsPrefixPrivateCluster,omitempty" tf:"dns_prefix_private_cluster,omitempty"`

	// +kubebuilder:validation:Required
	DefaultNodePool []DefaultNodePoolParameters `json:"defaultNodePool" tf:"default_node_pool,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	EnablePodSecurityPolicy *bool `json:"enablePodSecurityPolicy,omitempty" tf:"enable_pod_security_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	KubeletIdentity []KubeletIdentityParameters `json:"kubeletIdentity,omitempty" tf:"kubelet_identity,omitempty"`

	// +kubebuilder:validation:Optional
	KubernetesVersion *string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version,omitempty"`

	// +kubebuilder:validation:Optional
	LinuxProfile []LinuxProfileParameters `json:"linuxProfile,omitempty" tf:"linux_profile,omitempty"`

	// +kubebuilder:validation:Optional
	LocalAccountDisabled *bool `json:"localAccountDisabled,omitempty" tf:"local_account_disabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkProfile []NetworkProfileParameters `json:"networkProfile,omitempty" tf:"network_profile,omitempty"`

	// +kubebuilder:validation:Optional
	NodeResourceGroup *string `json:"nodeResourceGroup,omitempty" tf:"node_resource_group,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateClusterEnabled *bool `json:"privateClusterEnabled,omitempty" tf:"private_cluster_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateClusterPublicFqdnEnabled *bool `json:"privateClusterPublicFqdnEnabled,omitempty" tf:"private_cluster_public_fqdn_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateDNSZoneID *string `json:"privateDnsZoneId,omitempty" tf:"private_dns_zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateLinkEnabled *bool `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleBasedAccessControl []RoleBasedAccessControlParameters `json:"roleBasedAccessControl,omitempty" tf:"role_based_access_control,omitempty"`

	// +kubebuilder:validation:Optional
	ServicePrincipal []ServicePrincipalParameters `json:"servicePrincipal,omitempty" tf:"service_principal,omitempty"`

	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	WindowsProfile []WindowsProfileParameters `json:"windowsProfile,omitempty" tf:"windows_profile,omitempty"`
}

type LinuxOsConfigObservation struct {
}

type LinuxOsConfigParameters struct {

	// +kubebuilder:validation:Optional
	SwapFileSizeMb *int64 `json:"swapFileSizeMb,omitempty" tf:"swap_file_size_mb,omitempty"`

	// +kubebuilder:validation:Optional
	SysctlConfig []SysctlConfigParameters `json:"sysctlConfig,omitempty" tf:"sysctl_config,omitempty"`

	// +kubebuilder:validation:Optional
	TransparentHugePageDefrag *string `json:"transparentHugePageDefrag,omitempty" tf:"transparent_huge_page_defrag,omitempty"`

	// +kubebuilder:validation:Optional
	TransparentHugePageEnabled *string `json:"transparentHugePageEnabled,omitempty" tf:"transparent_huge_page_enabled,omitempty"`
}

type LinuxProfileObservation struct {
}

type LinuxProfileParameters struct {

	// +kubebuilder:validation:Required
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// +kubebuilder:validation:Required
	SSHKey []SSHKeyParameters `json:"sshKey" tf:"ssh_key,omitempty"`
}

type LoadBalancerProfileObservation struct {
	EffectiveOutboundIps []*string `json:"effectiveOutboundIps,omitempty" tf:"effective_outbound_ips,omitempty"`
}

type LoadBalancerProfileParameters struct {

	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	ManagedOutboundIPCount *int64 `json:"managedOutboundIpCount,omitempty" tf:"managed_outbound_ip_count,omitempty"`

	// +kubebuilder:validation:Optional
	OutboundIPAddressIds []*string `json:"outboundIpAddressIds,omitempty" tf:"outbound_ip_address_ids,omitempty"`

	// +kubebuilder:validation:Optional
	OutboundIPPrefixIds []*string `json:"outboundIpPrefixIds,omitempty" tf:"outbound_ip_prefix_ids,omitempty"`

	// +kubebuilder:validation:Optional
	OutboundPortsAllocated *int64 `json:"outboundPortsAllocated,omitempty" tf:"outbound_ports_allocated,omitempty"`
}

type MaintenanceWindowObservation struct {
}

type MaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	Allowed []AllowedParameters `json:"allowed,omitempty" tf:"allowed,omitempty"`

	// +kubebuilder:validation:Optional
	NotAllowed []NotAllowedParameters `json:"notAllowed,omitempty" tf:"not_allowed,omitempty"`
}

type NetworkProfileObservation struct {
}

type NetworkProfileParameters struct {

	// +kubebuilder:validation:Optional
	DNSServiceIP *string `json:"dnsServiceIp,omitempty" tf:"dns_service_ip,omitempty"`

	// +kubebuilder:validation:Optional
	DockerBridgeCidr *string `json:"dockerBridgeCidr,omitempty" tf:"docker_bridge_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerProfile []LoadBalancerProfileParameters `json:"loadBalancerProfile,omitempty" tf:"load_balancer_profile,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerSku *string `json:"loadBalancerSku,omitempty" tf:"load_balancer_sku,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkMode *string `json:"networkMode,omitempty" tf:"network_mode,omitempty"`

	// +kubebuilder:validation:Required
	NetworkPlugin *string `json:"networkPlugin" tf:"network_plugin,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkPolicy *string `json:"networkPolicy,omitempty" tf:"network_policy,omitempty"`

	// +kubebuilder:validation:Optional
	OutboundType *string `json:"outboundType,omitempty" tf:"outbound_type,omitempty"`

	// +kubebuilder:validation:Optional
	PodCidr *string `json:"podCidr,omitempty" tf:"pod_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceCidr *string `json:"serviceCidr,omitempty" tf:"service_cidr,omitempty"`
}

type NotAllowedObservation struct {
}

type NotAllowedParameters struct {

	// +kubebuilder:validation:Required
	End *string `json:"end" tf:"end,omitempty"`

	// +kubebuilder:validation:Required
	Start *string `json:"start" tf:"start,omitempty"`
}

type OmsAgentIdentityObservation struct {
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	UserAssignedIdentityID *string `json:"userAssignedIdentityId,omitempty" tf:"user_assigned_identity_id,omitempty"`
}

type OmsAgentIdentityParameters struct {
}

type OmsAgentObservation struct {
	OmsAgentIdentity []OmsAgentIdentityObservation `json:"omsAgentIdentity,omitempty" tf:"oms_agent_identity,omitempty"`
}

type OmsAgentParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`
}

type RoleBasedAccessControlObservation struct {
}

type RoleBasedAccessControlParameters struct {

	// +kubebuilder:validation:Optional
	AzureActiveDirectory []AzureActiveDirectoryParameters `json:"azureActiveDirectory,omitempty" tf:"azure_active_directory,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type SSHKeyObservation struct {
}

type SSHKeyParameters struct {

	// +kubebuilder:validation:Required
	KeyData *string `json:"keyData" tf:"key_data,omitempty"`
}

type ServicePrincipalObservation struct {
}

type ServicePrincipalParameters struct {

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`
}

type SysctlConfigObservation struct {
}

type SysctlConfigParameters struct {

	// +kubebuilder:validation:Optional
	FsAioMaxNr *int64 `json:"fsAioMaxNr,omitempty" tf:"fs_aio_max_nr,omitempty"`

	// +kubebuilder:validation:Optional
	FsFileMax *int64 `json:"fsFileMax,omitempty" tf:"fs_file_max,omitempty"`

	// +kubebuilder:validation:Optional
	FsInotifyMaxUserWatches *int64 `json:"fsInotifyMaxUserWatches,omitempty" tf:"fs_inotify_max_user_watches,omitempty"`

	// +kubebuilder:validation:Optional
	FsNrOpen *int64 `json:"fsNrOpen,omitempty" tf:"fs_nr_open,omitempty"`

	// +kubebuilder:validation:Optional
	KernelThreadsMax *int64 `json:"kernelThreadsMax,omitempty" tf:"kernel_threads_max,omitempty"`

	// +kubebuilder:validation:Optional
	NetCoreNetdevMaxBacklog *int64 `json:"netCoreNetdevMaxBacklog,omitempty" tf:"net_core_netdev_max_backlog,omitempty"`

	// +kubebuilder:validation:Optional
	NetCoreOptmemMax *int64 `json:"netCoreOptmemMax,omitempty" tf:"net_core_optmem_max,omitempty"`

	// +kubebuilder:validation:Optional
	NetCoreRmemDefault *int64 `json:"netCoreRmemDefault,omitempty" tf:"net_core_rmem_default,omitempty"`

	// +kubebuilder:validation:Optional
	NetCoreRmemMax *int64 `json:"netCoreRmemMax,omitempty" tf:"net_core_rmem_max,omitempty"`

	// +kubebuilder:validation:Optional
	NetCoreSomaxconn *int64 `json:"netCoreSomaxconn,omitempty" tf:"net_core_somaxconn,omitempty"`

	// +kubebuilder:validation:Optional
	NetCoreWmemDefault *int64 `json:"netCoreWmemDefault,omitempty" tf:"net_core_wmem_default,omitempty"`

	// +kubebuilder:validation:Optional
	NetCoreWmemMax *int64 `json:"netCoreWmemMax,omitempty" tf:"net_core_wmem_max,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4IPLocalPortRangeMax *int64 `json:"netIpv4IpLocalPortRangeMax,omitempty" tf:"net_ipv4_ip_local_port_range_max,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4IPLocalPortRangeMin *int64 `json:"netIpv4IpLocalPortRangeMin,omitempty" tf:"net_ipv4_ip_local_port_range_min,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4NeighDefaultGcThresh1 *int64 `json:"netIpv4NeighDefaultGcThresh1,omitempty" tf:"net_ipv4_neigh_default_gc_thresh1,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4NeighDefaultGcThresh2 *int64 `json:"netIpv4NeighDefaultGcThresh2,omitempty" tf:"net_ipv4_neigh_default_gc_thresh2,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4NeighDefaultGcThresh3 *int64 `json:"netIpv4NeighDefaultGcThresh3,omitempty" tf:"net_ipv4_neigh_default_gc_thresh3,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPFinTimeout *int64 `json:"netIpv4TcpFinTimeout,omitempty" tf:"net_ipv4_tcp_fin_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPKeepaliveIntvl *int64 `json:"netIpv4TcpKeepaliveIntvl,omitempty" tf:"net_ipv4_tcp_keepalive_intvl,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPKeepaliveProbes *int64 `json:"netIpv4TcpKeepaliveProbes,omitempty" tf:"net_ipv4_tcp_keepalive_probes,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPKeepaliveTime *int64 `json:"netIpv4TcpKeepaliveTime,omitempty" tf:"net_ipv4_tcp_keepalive_time,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPMaxSynBacklog *int64 `json:"netIpv4TcpMaxSynBacklog,omitempty" tf:"net_ipv4_tcp_max_syn_backlog,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPMaxTwBuckets *int64 `json:"netIpv4TcpMaxTwBuckets,omitempty" tf:"net_ipv4_tcp_max_tw_buckets,omitempty"`

	// +kubebuilder:validation:Optional
	NetIPv4TCPTwReuse *bool `json:"netIpv4TcpTwReuse,omitempty" tf:"net_ipv4_tcp_tw_reuse,omitempty"`

	// +kubebuilder:validation:Optional
	NetNetfilterNfConntrackBuckets *int64 `json:"netNetfilterNfConntrackBuckets,omitempty" tf:"net_netfilter_nf_conntrack_buckets,omitempty"`

	// +kubebuilder:validation:Optional
	NetNetfilterNfConntrackMax *int64 `json:"netNetfilterNfConntrackMax,omitempty" tf:"net_netfilter_nf_conntrack_max,omitempty"`

	// +kubebuilder:validation:Optional
	VMMaxMapCount *int64 `json:"vmMaxMapCount,omitempty" tf:"vm_max_map_count,omitempty"`

	// +kubebuilder:validation:Optional
	VMSwappiness *int64 `json:"vmSwappiness,omitempty" tf:"vm_swappiness,omitempty"`

	// +kubebuilder:validation:Optional
	VMVfsCachePressure *int64 `json:"vmVfsCachePressure,omitempty" tf:"vm_vfs_cache_pressure,omitempty"`
}

type UpgradeSettingsObservation struct {
}

type UpgradeSettingsParameters struct {

	// +kubebuilder:validation:Required
	MaxSurge *string `json:"maxSurge" tf:"max_surge,omitempty"`
}

type WindowsProfileObservation struct {
}

type WindowsProfileParameters struct {

	// +kubebuilder:validation:Optional
	AdminPasswordSecretRef *v1.SecretKeySelector `json:"adminPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// +kubebuilder:validation:Optional
	License *string `json:"license,omitempty" tf:"license,omitempty"`
}

// KubernetesClusterSpec defines the desired state of KubernetesCluster
type KubernetesClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KubernetesClusterParameters `json:"forProvider"`
}

// KubernetesClusterStatus defines the observed state of KubernetesCluster.
type KubernetesClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KubernetesClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesCluster is the Schema for the KubernetesClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesClusterList contains a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesCluster `json:"items"`
}

// Repository type metadata.
var (
	KubernetesCluster_Kind             = "KubernetesCluster"
	KubernetesCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KubernetesCluster_Kind}.String()
	KubernetesCluster_KindAPIVersion   = KubernetesCluster_Kind + "." + CRDGroupVersion.String()
	KubernetesCluster_GroupVersionKind = CRDGroupVersion.WithKind(KubernetesCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&KubernetesCluster{}, &KubernetesClusterList{})
}
