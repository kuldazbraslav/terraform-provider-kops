package structures

import (
	"reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandKubeProxyConfig(in map[string]interface{}) kops.KubeProxyConfig {
	if in == nil {
		panic("expand KubeProxyConfig failure, in is nil")
	}
	return kops.KubeProxyConfig{
		Image: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["image"]),
		CPURequest: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cpu_request"]),
		CPULimit: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cpu_limit"]),
		MemoryRequest: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["memory_request"]),
		MemoryLimit: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["memory_limit"]),
		LogLevel: func(in interface{}) int32 {
			value := int32(ExpandInt(in))
			return value
		}(in["log_level"]),
		ClusterCIDR: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["cluster_cidr"]),
		HostnameOverride: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["hostname_override"]),
		BindAddress: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["bind_address"]),
		Master: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["master"]),
		MetricsBindAddress: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["metrics_bind_address"]),
		Enabled: func(in interface{}) *bool {
			value := func(in interface{}) *bool {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in bool) *bool {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(bool(ExpandBool(in)))
				return tmp
			}(in)
			return value
		}(in["enabled"]),
		ProxyMode: func(in interface{}) string {
			value := string(ExpandString(in))
			return value
		}(in["proxy_mode"]),
		IPVSExcludeCIDRS: func(in interface{}) []string {
			value := func(in interface{}) []string {
				var out []string
				for _, in := range in.([]interface{}) {
					out = append(out, string(ExpandString(in)))
				}
				return out
			}(in)
			return value
		}(in["ip_vs_exclude_cidr_s"]),
		IPVSMinSyncPeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["ip_vs_min_sync_period"]),
		IPVSScheduler: func(in interface{}) *string {
			value := func(in interface{}) *string {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in string) *string {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(string(ExpandString(in)))
				return tmp
			}(in)
			return value
		}(in["ip_vs_scheduler"]),
		IPVSSyncPeriod: func(in interface{}) *v1.Duration {
			value := func(in interface{}) *v1.Duration {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in v1.Duration) *v1.Duration {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(ExpandDuration(in))
				return tmp
			}(in)
			return value
		}(in["ip_vs_sync_period"]),
		FeatureGates: func(in interface{}) map[string]string {
			value := func(in interface{}) map[string]string {
				if in == nil {
					return nil
				}
				out := map[string]string{}
				for key, in := range in.(map[string]interface{}) {
					out[key] = string(ExpandString(in))
				}
				return out
			}(in)
			return value
		}(in["feature_gates"]),
		ConntrackMaxPerCore: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["conntrack_max_per_core"]),
		ConntrackMin: func(in interface{}) *int32 {
			value := func(in interface{}) *int32 {
				if in == nil {
					return nil
				}
				if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
					return nil
				}
				tmp := func(in int32) *int32 {
					if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
						return nil
					}
					return &in
				}(int32(ExpandInt(in)))
				return tmp
			}(in)
			return value
		}(in["conntrack_min"]),
	}
}

func FlattenKubeProxyConfig(in kops.KubeProxyConfig) map[string]interface{} {
	return map[string]interface{}{
		"image": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Image),
		"cpu_request": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CPURequest),
		"cpu_limit": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.CPULimit),
		"memory_request": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.MemoryRequest),
		"memory_limit": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.MemoryLimit),
		"log_level": func(in int32) interface{} {
			value := FlattenInt(int(in))
			return value
		}(in.LogLevel),
		"cluster_cidr": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ClusterCIDR),
		"hostname_override": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.HostnameOverride),
		"bind_address": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.BindAddress),
		"master": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.Master),
		"metrics_bind_address": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.MetricsBindAddress),
		"enabled": func(in *bool) interface{} {
			value := func(in *bool) interface{} {
				if in == nil {
					return nil
				}
				return func(in bool) interface{} {
					return FlattenBool(bool(in))
				}(*in)
			}(in)
			return value
		}(in.Enabled),
		"proxy_mode": func(in string) interface{} {
			value := FlattenString(string(in))
			return value
		}(in.ProxyMode),
		"ip_vs_exclude_cidr_s": func(in []string) interface{} {
			value := func(in []string) []interface{} {
				var out []interface{}
				for _, in := range in {
					out = append(out, FlattenString(string(in)))
				}
				return out
			}(in)
			return value
		}(in.IPVSExcludeCIDRS),
		"ip_vs_min_sync_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.IPVSMinSyncPeriod),
		"ip_vs_scheduler": func(in *string) interface{} {
			value := func(in *string) interface{} {
				if in == nil {
					return nil
				}
				return func(in string) interface{} {
					return FlattenString(string(in))
				}(*in)
			}(in)
			return value
		}(in.IPVSScheduler),
		"ip_vs_sync_period": func(in *v1.Duration) interface{} {
			value := func(in *v1.Duration) interface{} {
				if in == nil {
					return nil
				}
				return func(in v1.Duration) interface{} {
					return FlattenDuration(in)
				}(*in)
			}(in)
			return value
		}(in.IPVSSyncPeriod),
		"feature_gates": func(in map[string]string) interface{} {
			value := func(in map[string]string) map[string]interface{} {
				if in == nil {
					return nil
				}
				// TODO
				return nil
			}(in)
			return value
		}(in.FeatureGates),
		"conntrack_max_per_core": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConntrackMaxPerCore),
		"conntrack_min": func(in *int32) interface{} {
			value := func(in *int32) interface{} {
				if in == nil {
					return nil
				}
				return func(in int32) interface{} {
					return FlattenInt(int(in))
				}(*in)
			}(in)
			return value
		}(in.ConntrackMin),
	}
}