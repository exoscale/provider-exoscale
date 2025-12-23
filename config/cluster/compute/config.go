package repository

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_ssh_key", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "SSHKey"
	})

	p.AddResourceConfigurator("exoscale_security_group", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "SecurityGroup"
	})

	p.AddResourceConfigurator("exoscale_security_group_rule", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "SecurityGroupRules"

		r.References["security_group_id"] = config.Reference{
			Type: "github.com/exoscale/provider-exoscale/apis/cluster/compute/v1alpha1.SecurityGroup",
		}

		// Ignore the deprecated security_group field to avoid conflicts
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"security_group"},
		}
	})
}
