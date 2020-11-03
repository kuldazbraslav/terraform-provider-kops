package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func KubeletConfigSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_servers":                            OptionalString(),
			"anonymous_auth":                         OptionalBool(),
			"authorization_mode":                     OptionalString(),
			"bootstrap_kubeconfig":                   OptionalString(),
			"client_ca_file":                         OptionalString(),
			"tls_cert_file":                          OptionalString(),
			"tls_private_key_file":                   OptionalString(),
			"tls_cipher_suites":                      OptionalList(String()),
			"tls_min_version":                        OptionalString(),
			"kubeconfig_path":                        OptionalString(),
			"require_kubeconfig":                     OptionalBool(),
			"log_level":                              OptionalInt(),
			"pod_manifest_path":                      OptionalString(),
			"hostname_override":                      OptionalString(),
			"pod_infra_container_image":              OptionalString(),
			"seccomp_profile_root":                   OptionalString(),
			"allow_privileged":                       OptionalBool(),
			"enable_debugging_handlers":              OptionalBool(),
			"register_node":                          OptionalBool(),
			"node_status_update_frequency":           OptionalDuration(),
			"cluster_domain":                         OptionalString(),
			"cluster_dns":                            OptionalString(),
			"network_plugin_name":                    OptionalString(),
			"cloud_provider":                         OptionalString(),
			"kubelet_cgroups":                        OptionalString(),
			"runtime_cgroups":                        OptionalString(),
			"read_only_port":                         OptionalInt(),
			"system_cgroups":                         OptionalString(),
			"cgroup_root":                            OptionalString(),
			"configure_cbr_00":                       OptionalBool(),
			"hairpin_mode":                           OptionalString(),
			"babysit_daemons":                        OptionalBool(),
			"max_pods":                               OptionalInt(),
			"nvidia_gp_uss":                          OptionalInt(),
			"pod_cidr":                               OptionalString(),
			"resolver_config":                        OptionalString(),
			"reconcile_cidr":                         OptionalBool(),
			"register_schedulable":                   OptionalBool(),
			"serialize_image_pulls":                  OptionalBool(),
			"node_labels":                            OptionalMap(String()),
			"non_masquerade_cidr":                    OptionalString(),
			"enable_custom_metrics":                  OptionalBool(),
			"network_plugin_mtu":                     OptionalInt(),
			"image_gc_high_threshold_percent":        OptionalInt(),
			"image_gc_low_threshold_percent":         OptionalInt(),
			"image_pull_progress_deadline":           OptionalDuration(),
			"eviction_hard":                          OptionalString(),
			"eviction_soft":                          OptionalString(),
			"eviction_soft_grace_period":             OptionalString(),
			"eviction_pressure_transition_period":    OptionalDuration(),
			"eviction_max_pod_grace_period":          OptionalInt(),
			"eviction_minimum_reclaim":               OptionalString(),
			"volume_plugin_directory":                OptionalString(),
			"taints":                                 OptionalList(String()),
			"feature_gates":                          OptionalMap(String()),
			"kube_reserved":                          OptionalMap(String()),
			"kube_reserved_cgroup":                   OptionalString(),
			"system_reserved":                        OptionalMap(String()),
			"system_reserved_cgroup":                 OptionalString(),
			"enforce_node_allocatable":               OptionalString(),
			"runtime_request_timeout":                OptionalDuration(),
			"volume_stats_agg_period":                OptionalDuration(),
			"fail_swap_on":                           OptionalBool(),
			"experimental_allowed_unsafe_sysctls":    OptionalList(String()),
			"allowed_unsafe_sysctls":                 OptionalList(String()),
			"streaming_connection_idle_timeout":      OptionalDuration(),
			"docker_disable_shared_pid":              OptionalBool(),
			"root_dir":                               OptionalString(),
			"authentication_token_webhook":           OptionalBool(),
			"authentication_token_webhook_cache_ttl": OptionalDuration(),
			"cpucfs_quota":                           OptionalBool(),
			"cpucfs_quota_period":                    OptionalDuration(),
			"cpu_manager_policy":                     OptionalString(),
			"registry_pull_qps":                      OptionalInt(),
			"registry_burst":                         OptionalInt(),
			"topology_manager_policy":                OptionalString(),
			"rotate_certificates":                    OptionalBool(),
			"protect_kernel_defaults":                OptionalBool(),
			"cgroup_driver":                          OptionalString(),
		},
	}
}