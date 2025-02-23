package schemas

import (
	"reflect"

	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

var _ = Schema

func ExpandDataSourceInstanceGroupSpec(in map[string]interface{}) kops.InstanceGroupSpec {
	if in == nil {
		panic("expand InstanceGroupSpec failure, in is nil")
	}
	return kops.InstanceGroupSpec{
		Manager: func(in interface{}) kops.InstanceManager {
			return kops.InstanceManager(ExpandString(in))
		}(in["manager"]),
		Role: func(in interface{}) kops.InstanceGroupRole {
			return kops.InstanceGroupRole(ExpandString(in))
		}(in["role"]),
		Image: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["image"]),
		MinSize: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["min_size"]),
		MaxSize: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["max_size"]),
		Autoscale: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["autoscale"]),
		MachineType: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["machine_type"]),
		RootVolumeSize: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["root_volume_size"]),
		RootVolumeType: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["root_volume_type"]),
		RootVolumeIOPS: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["root_volume_iops"]),
		RootVolumeThroughput: func(in interface{}) *int32 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int32) *int32 {
					return &in
				}(int32(ExpandInt(in)))
			}(in)
		}(in["root_volume_throughput"]),
		RootVolumeOptimization: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["root_volume_optimization"]),
		RootVolumeEncryption: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["root_volume_encryption"]),
		RootVolumeEncryptionKey: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["root_volume_encryption_key"]),
		Volumes: func(in interface{}) []kops.VolumeSpec {
			return func(in interface{}) []kops.VolumeSpec {
				if in == nil {
					return nil
				}
				var out []kops.VolumeSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.VolumeSpec {
						if in == nil {
							return kops.VolumeSpec{}
						}
						return ExpandDataSourceVolumeSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["volumes"]),
		VolumeMounts: func(in interface{}) []kops.VolumeMountSpec {
			return func(in interface{}) []kops.VolumeMountSpec {
				if in == nil {
					return nil
				}
				var out []kops.VolumeMountSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.VolumeMountSpec {
						if in == nil {
							return kops.VolumeMountSpec{}
						}
						return ExpandDataSourceVolumeMountSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["volume_mounts"]),
		Subnets: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["subnets"]),
		Zones: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["zones"]),
		Hooks: func(in interface{}) []kops.HookSpec {
			return func(in interface{}) []kops.HookSpec {
				if in == nil {
					return nil
				}
				var out []kops.HookSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.HookSpec {
						if in == nil {
							return kops.HookSpec{}
						}
						return ExpandDataSourceHookSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["hooks"]),
		MaxPrice: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["max_price"]),
		SpotDurationInMinutes: func(in interface{}) *int64 {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *int64 {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in int64) *int64 {
					return &in
				}(int64(ExpandInt(in)))
			}(in)
		}(in["spot_duration_in_minutes"]),
		CPUCredits: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["cpu_credits"]),
		AssociatePublicIP: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["associate_public_ip"]),
		AdditionalSecurityGroups: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["additional_security_groups"]),
		CloudLabels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["cloud_labels"]),
		NodeLabels: func(in interface{}) map[string]string {
			return func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				if in, ok := in.(map[string]interface{}); ok {
					if len(in) > 0 {
						out := map[string]string{}
						for key, in := range in {
							out[key] = string(ExpandString(in))
						}
						return out
					}
				}
				return nil
			}(in)
		}(in["node_labels"]),
		FileAssets: func(in interface{}) []kops.FileAssetSpec {
			return func(in interface{}) []kops.FileAssetSpec {
				if in == nil {
					return nil
				}
				var out []kops.FileAssetSpec
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.FileAssetSpec {
						if in == nil {
							return kops.FileAssetSpec{}
						}
						return ExpandDataSourceFileAssetSpec(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["file_assets"]),
		Tenancy: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["tenancy"]),
		Kubelet: func(in interface{}) *kops.KubeletConfigSpec {
			return func(in interface{}) *kops.KubeletConfigSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.KubeletConfigSpec) *kops.KubeletConfigSpec {
					return &in
				}(func(in interface{}) kops.KubeletConfigSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceKubeletConfigSpec(in[0].(map[string]interface{}))
					}
					return kops.KubeletConfigSpec{}
				}(in))
			}(in)
		}(in["kubelet"]),
		Taints: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["taints"]),
		MixedInstancesPolicy: func(in interface{}) *kops.MixedInstancesPolicySpec {
			return func(in interface{}) *kops.MixedInstancesPolicySpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.MixedInstancesPolicySpec) *kops.MixedInstancesPolicySpec {
					return &in
				}(func(in interface{}) kops.MixedInstancesPolicySpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceMixedInstancesPolicySpec(in[0].(map[string]interface{}))
					}
					return kops.MixedInstancesPolicySpec{}
				}(in))
			}(in)
		}(in["mixed_instances_policy"]),
		AdditionalUserData: func(in interface{}) []kops.UserData {
			return func(in interface{}) []kops.UserData {
				if in == nil {
					return nil
				}
				var out []kops.UserData
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.UserData {
						if in == nil {
							return kops.UserData{}
						}
						return ExpandDataSourceUserData(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["additional_user_data"]),
		SuspendProcesses: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["suspend_processes"]),
		ExternalLoadBalancers: func(in interface{}) []kops.LoadBalancer {
			return func(in interface{}) []kops.LoadBalancer {
				if in == nil {
					return nil
				}
				var out []kops.LoadBalancer
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.LoadBalancer {
						if in == nil {
							return kops.LoadBalancer{}
						}
						return ExpandDataSourceLoadBalancer(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["external_load_balancers"]),
		DetailedInstanceMonitoring: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["detailed_instance_monitoring"]),
		IAM: func(in interface{}) *kops.IAMProfileSpec {
			return func(in interface{}) *kops.IAMProfileSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.IAMProfileSpec) *kops.IAMProfileSpec {
					return &in
				}(func(in interface{}) kops.IAMProfileSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceIAMProfileSpec(in[0].(map[string]interface{}))
					}
					return kops.IAMProfileSpec{}
				}(in))
			}(in)
		}(in["iam"]),
		SecurityGroupOverride: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["security_group_override"]),
		InstanceProtection: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["instance_protection"]),
		SysctlParameters: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["sysctl_parameters"]),
		RollingUpdate: func(in interface{}) *kops.RollingUpdate {
			return func(in interface{}) *kops.RollingUpdate {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.RollingUpdate) *kops.RollingUpdate {
					return &in
				}(func(in interface{}) kops.RollingUpdate {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceRollingUpdate(in[0].(map[string]interface{}))
					}
					return kops.RollingUpdate{}
				}(in))
			}(in)
		}(in["rolling_update"]),
		InstanceInterruptionBehavior: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["instance_interruption_behavior"]),
		CompressUserData: func(in interface{}) *bool {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in bool) *bool {
					return &in
				}(bool(ExpandBool(in)))
			}(in)
		}(in["compress_user_data"]),
		InstanceMetadata: func(in interface{}) *kops.InstanceMetadataOptions {
			return func(in interface{}) *kops.InstanceMetadataOptions {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.InstanceMetadataOptions) *kops.InstanceMetadataOptions {
					return &in
				}(func(in interface{}) kops.InstanceMetadataOptions {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceInstanceMetadataOptions(in[0].(map[string]interface{}))
					}
					return kops.InstanceMetadataOptions{}
				}(in))
			}(in)
		}(in["instance_metadata"]),
		UpdatePolicy: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["update_policy"]),
		WarmPool: func(in interface{}) *kops.WarmPoolSpec {
			return func(in interface{}) *kops.WarmPoolSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.WarmPoolSpec) *kops.WarmPoolSpec {
					return &in
				}(func(in interface{}) kops.WarmPoolSpec {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceWarmPoolSpec(in[0].(map[string]interface{}))
					}
					return kops.WarmPoolSpec{}
				}(in))
			}(in)
		}(in["warm_pool"]),
		Containerd: func(in interface{}) *kops.ContainerdConfig {
			return func(in interface{}) *kops.ContainerdConfig {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.ContainerdConfig) *kops.ContainerdConfig {
					return &in
				}(func(in interface{}) kops.ContainerdConfig {
					if in, ok := in.([]interface{}); ok && len(in) == 1 && in[0] != nil {
						return ExpandDataSourceContainerdConfig(in[0].(map[string]interface{}))
					}
					return kops.ContainerdConfig{}
				}(in))
			}(in)
		}(in["containerd"]),
		Packages: func(in interface{}) []string {
			return func(in interface{}) []string {
				if in == nil {
					return nil
				}
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
		}(in["packages"]),
		GuestAccelerators: func(in interface{}) []kops.AcceleratorConfig {
			return func(in interface{}) []kops.AcceleratorConfig {
				if in == nil {
					return nil
				}
				var out []kops.AcceleratorConfig
				for _, in := range in.([]interface{}) {
					out = append(out, func(in interface{}) kops.AcceleratorConfig {
						if in == nil {
							return kops.AcceleratorConfig{}
						}
						return ExpandDataSourceAcceleratorConfig(in.(map[string]interface{}))
					}(in))
				}
				return out
			}(in)
		}(in["guest_accelerators"]),
		MaxInstanceLifetime: func(in interface{}) *meta.Duration {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *meta.Duration {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in meta.Duration) *meta.Duration {
					return &in
				}(ExpandDuration(in))
			}(in)
		}(in["max_instance_lifetime"]),
		GCPProvisioningModel: func(in interface{}) *string {
			if in == nil {
				return nil
			}
			if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
				return nil
			}
			return func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in string) *string {
					return &in
				}(string(ExpandString(in)))
			}(in)
		}(in["gcp_provisioning_model"]),
	}
}

func FlattenDataSourceInstanceGroupSpecInto(in kops.InstanceGroupSpec, out map[string]interface{}) {
	out["manager"] = func(in kops.InstanceManager) interface{} {
		return FlattenString(string(in))
	}(in.Manager)
	out["role"] = func(in kops.InstanceGroupRole) interface{} {
		return FlattenString(string(in))
	}(in.Role)
	out["image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Image)
	out["min_size"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MinSize)
	out["max_size"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.MaxSize)
	out["autoscale"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.Autoscale)
	out["machine_type"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.MachineType)
	out["root_volume_size"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.RootVolumeSize)
	out["root_volume_type"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.RootVolumeType)
	out["root_volume_iops"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.RootVolumeIOPS)
	out["root_volume_throughput"] = func(in *int32) interface{} {
		return func(in *int32) interface{} {
			if in == nil {
				return nil
			}
			return func(in int32) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.RootVolumeThroughput)
	out["root_volume_optimization"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.RootVolumeOptimization)
	out["root_volume_encryption"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.RootVolumeEncryption)
	out["root_volume_encryption_key"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.RootVolumeEncryptionKey)
	out["volumes"] = func(in []kops.VolumeSpec) interface{} {
		return func(in []kops.VolumeSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.VolumeSpec) interface{} {
					return FlattenDataSourceVolumeSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Volumes)
	out["volume_mounts"] = func(in []kops.VolumeMountSpec) interface{} {
		return func(in []kops.VolumeMountSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.VolumeMountSpec) interface{} {
					return FlattenDataSourceVolumeMountSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.VolumeMounts)
	out["subnets"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Subnets)
	out["zones"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Zones)
	out["hooks"] = func(in []kops.HookSpec) interface{} {
		return func(in []kops.HookSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.HookSpec) interface{} {
					return FlattenDataSourceHookSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.Hooks)
	out["max_price"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.MaxPrice)
	out["spot_duration_in_minutes"] = func(in *int64) interface{} {
		return func(in *int64) interface{} {
			if in == nil {
				return nil
			}
			return func(in int64) interface{} {
				return FlattenInt(int(in))
			}(*in)
		}(in)
	}(in.SpotDurationInMinutes)
	out["cpu_credits"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.CPUCredits)
	out["associate_public_ip"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.AssociatePublicIP)
	out["additional_security_groups"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.AdditionalSecurityGroups)
	out["cloud_labels"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.CloudLabels)
	out["node_labels"] = func(in map[string]string) interface{} {
		return func(in map[string]string) map[string]interface{} {
			if in == nil {
				return nil
			}
			out := map[string]interface{}{}
			for key, in := range in {
				out[key] = FlattenString(string(in))
			}
			return out
		}(in)
	}(in.NodeLabels)
	out["file_assets"] = func(in []kops.FileAssetSpec) interface{} {
		return func(in []kops.FileAssetSpec) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.FileAssetSpec) interface{} {
					return FlattenDataSourceFileAssetSpec(in)
				}(in))
			}
			return out
		}(in)
	}(in.FileAssets)
	out["tenancy"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Tenancy)
	out["kubelet"] = func(in *kops.KubeletConfigSpec) interface{} {
		return func(in *kops.KubeletConfigSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.KubeletConfigSpec) interface{} {
				return func(in kops.KubeletConfigSpec) []interface{} {
					return []interface{}{FlattenDataSourceKubeletConfigSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Kubelet)
	out["taints"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Taints)
	out["mixed_instances_policy"] = func(in *kops.MixedInstancesPolicySpec) interface{} {
		return func(in *kops.MixedInstancesPolicySpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.MixedInstancesPolicySpec) interface{} {
				return func(in kops.MixedInstancesPolicySpec) []interface{} {
					return []interface{}{FlattenDataSourceMixedInstancesPolicySpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.MixedInstancesPolicy)
	out["additional_user_data"] = func(in []kops.UserData) interface{} {
		return func(in []kops.UserData) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.UserData) interface{} {
					return FlattenDataSourceUserData(in)
				}(in))
			}
			return out
		}(in)
	}(in.AdditionalUserData)
	out["suspend_processes"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.SuspendProcesses)
	out["external_load_balancers"] = func(in []kops.LoadBalancer) interface{} {
		return func(in []kops.LoadBalancer) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.LoadBalancer) interface{} {
					return FlattenDataSourceLoadBalancer(in)
				}(in))
			}
			return out
		}(in)
	}(in.ExternalLoadBalancers)
	out["detailed_instance_monitoring"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.DetailedInstanceMonitoring)
	out["iam"] = func(in *kops.IAMProfileSpec) interface{} {
		return func(in *kops.IAMProfileSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.IAMProfileSpec) interface{} {
				return func(in kops.IAMProfileSpec) []interface{} {
					return []interface{}{FlattenDataSourceIAMProfileSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.IAM)
	out["security_group_override"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.SecurityGroupOverride)
	out["instance_protection"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.InstanceProtection)
	out["sysctl_parameters"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.SysctlParameters)
	out["rolling_update"] = func(in *kops.RollingUpdate) interface{} {
		return func(in *kops.RollingUpdate) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.RollingUpdate) interface{} {
				return func(in kops.RollingUpdate) []interface{} {
					return []interface{}{FlattenDataSourceRollingUpdate(in)}
				}(in)
			}(*in)
		}(in)
	}(in.RollingUpdate)
	out["instance_interruption_behavior"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.InstanceInterruptionBehavior)
	out["compress_user_data"] = func(in *bool) interface{} {
		return func(in *bool) interface{} {
			if in == nil {
				return nil
			}
			return func(in bool) interface{} {
				return FlattenBool(bool(in))
			}(*in)
		}(in)
	}(in.CompressUserData)
	out["instance_metadata"] = func(in *kops.InstanceMetadataOptions) interface{} {
		return func(in *kops.InstanceMetadataOptions) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.InstanceMetadataOptions) interface{} {
				return func(in kops.InstanceMetadataOptions) []interface{} {
					return []interface{}{FlattenDataSourceInstanceMetadataOptions(in)}
				}(in)
			}(*in)
		}(in)
	}(in.InstanceMetadata)
	out["update_policy"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.UpdatePolicy)
	out["warm_pool"] = func(in *kops.WarmPoolSpec) interface{} {
		return func(in *kops.WarmPoolSpec) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.WarmPoolSpec) interface{} {
				return func(in kops.WarmPoolSpec) []interface{} {
					return []interface{}{FlattenDataSourceWarmPoolSpec(in)}
				}(in)
			}(*in)
		}(in)
	}(in.WarmPool)
	out["containerd"] = func(in *kops.ContainerdConfig) interface{} {
		return func(in *kops.ContainerdConfig) interface{} {
			if in == nil {
				return nil
			}
			return func(in kops.ContainerdConfig) interface{} {
				return func(in kops.ContainerdConfig) []interface{} {
					return []interface{}{FlattenDataSourceContainerdConfig(in)}
				}(in)
			}(*in)
		}(in)
	}(in.Containerd)
	out["packages"] = func(in []string) interface{} {
		return func(in []string) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, FlattenString(string(in)))
			}
			return out
		}(in)
	}(in.Packages)
	out["guest_accelerators"] = func(in []kops.AcceleratorConfig) interface{} {
		return func(in []kops.AcceleratorConfig) []interface{} {
			var out []interface{}
			for _, in := range in {
				out = append(out, func(in kops.AcceleratorConfig) interface{} {
					return FlattenDataSourceAcceleratorConfig(in)
				}(in))
			}
			return out
		}(in)
	}(in.GuestAccelerators)
	out["max_instance_lifetime"] = func(in *meta.Duration) interface{} {
		return func(in *meta.Duration) interface{} {
			if in == nil {
				return nil
			}
			return func(in meta.Duration) interface{} {
				return FlattenDuration(in)
			}(*in)
		}(in)
	}(in.MaxInstanceLifetime)
	out["gcp_provisioning_model"] = func(in *string) interface{} {
		return func(in *string) interface{} {
			if in == nil {
				return nil
			}
			return func(in string) interface{} {
				return FlattenString(string(in))
			}(*in)
		}(in)
	}(in.GCPProvisioningModel)
}

func FlattenDataSourceInstanceGroupSpec(in kops.InstanceGroupSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenDataSourceInstanceGroupSpecInto(in, out)
	return out
}
