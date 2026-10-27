package repository

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup string = "compute"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_ssh_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SSHKey"
	})

	p.AddResourceConfigurator("exoscale_security_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SecurityGroup"
	})

	p.AddResourceConfigurator("exoscale_security_group_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SecurityGroupRules"

		r.References["security_group_id"] = config.Reference{
			TerraformName: "exoscale_security_group",
		}

		// Ignore the deprecated security_group field to avoid conflicts
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"security_group"},
		}
	})

	p.AddResourceConfigurator("exoscale_anti_affinity_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AntiAffinityGroup"
	})

	p.AddResourceConfigurator("exoscale_block_storage_volume", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BlockStorageVolume"
		r.TerraformResource.Schema["zone"].ForceNew = true
	})

	p.AddResourceConfigurator("exoscale_elastic_ip", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ElasticIP"
	})

	p.AddResourceConfigurator("exoscale_private_network", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PrivateNetwork"
	})

	p.AddResourceConfigurator("exoscale_compute_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Instance"

		r.References["anti_affinity_group_ids"] = config.Reference{
			TerraformName: "exoscale_anti_affinity_group",
		}
		r.References["block_storage_volume_ids"] = config.Reference{
			TerraformName: "exoscale_block_storage_volume",
		}
		r.References["elastic_ip_ids"] = config.Reference{
			TerraformName: "exoscale_elastic_ip",
		}
		r.References["network_interface.network_id"] = config.Reference{
			TerraformName: "exoscale_private_network",
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName: "exoscale_security_group",
		}
		r.References["ssh_keys"] = config.Reference{
			TerraformName: "exoscale_ssh_key",
		}
	})

	p.AddResourceConfigurator("exoscale_instance_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "InstancePool"

		r.References["affinity_group_ids"] = config.Reference{
			TerraformName: "exoscale_anti_affinity_group",
		}
		r.References["anti_affinity_group_ids"] = config.Reference{
			TerraformName: "exoscale_anti_affinity_group",
		}
		r.References["elastic_ip_ids"] = config.Reference{
			TerraformName: "exoscale_elastic_ip",
		}
		r.References["key_pair"] = config.Reference{
			TerraformName: "exoscale_ssh_key",
		}
		r.References["network_ids"] = config.Reference{
			TerraformName: "exoscale_private_network",
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName: "exoscale_security_group",
		}

		// Ignore the deprecated service_offering, virtual_machines field to avoid conflicts
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"service_offering", "virtual_machines", "affinity_group_ids"},
		}
	})

	p.AddResourceConfigurator("exoscale_nlb", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NLB"
	})

	p.AddResourceConfigurator("exoscale_nlb_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NLBService"

		r.References["instance_pool_id"] = config.Reference{
			TerraformName: "exoscale_instance_pool",
		}
		r.References["nlb_id"] = config.Reference{
			TerraformName: "exoscale_nlb",
		}
	})

	p.AddResourceConfigurator("exoscale_sks_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SKSCluster"
	})

	p.AddResourceConfigurator("exoscale_sks_nodepool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SKSNodepool"

		r.References["cluster_id"] = config.Reference{
			TerraformName: "exoscale_sks_cluster",
		}
		r.References["anti_affinity_group_ids"] = config.Reference{
			TerraformName: "exoscale_anti_affinity_group",
		}
		r.References["private_network_ids"] = config.Reference{
			TerraformName: "exoscale_private_network",
		}
		r.References["security_group_ids"] = config.Reference{
			TerraformName: "exoscale_security_group",
		}
	})
}
