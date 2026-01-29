package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this provider.
// cf. https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md#external-name
// TODO: check how the read function in terraform provider works.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Compute resources
	"exoscale_ssh_key":              config.NameAsIdentifier,
	"exoscale_security_group":       config.IdentifierFromProvider,
	"exoscale_security_group_rule":  config.IdentifierFromProvider,
	"exoscale_anti_affinity_group":  config.IdentifierFromProvider,
	"exoscale_block_storage_volume": config.IdentifierFromProvider, // only works with terraform-provider-exoscale version > 0.67.1 TODO: upgrade terraform version in makefile once the new release is out.
	"exoscale_elastic_ip":           config.IdentifierFromProvider,
	"exoscale_private_network":      config.IdentifierFromProvider,
	"exoscale_compute_instance":     config.IdentifierFromProvider,
	"exoscale_instance_pool":        config.IdentifierFromProvider, // only works with terraform-provider-exoscale version > 0.67.1 TODO: upgrade terraform version in makefile once the new release is out.
	"exoscale_nlb":                  config.IdentifierFromProvider,
	"exoscale_nlb_service":          config.IdentifierFromProvider,
	"exoscale_sks_cluster":          config.IdentifierFromProvider, // only works with terraform-provider-exoscale version > 0.67.1 TODO: upgrade terraform version in makefile once the new release is out.
	"exoscale_sks_nodepool":         config.IdentifierFromProvider, // only works with terraform-provider-exoscale version > v0.67.2 TODO: upgrade terraform version in makefile once the new release is out.

	// DBAAS resources
	"exoscale_dbaas":             config.IdentifierFromProvider,
	"exoscale_dbaas_pg_user":     config.IdentifierFromProvider,
	"exoscale_dbaas_pg_database": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
