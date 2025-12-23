package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this provider.
// cf. https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md#external-name
var ExternalNameConfigs = map[string]config.ExternalName{
	"exoscale_ssh_key":             config.NameAsIdentifier,       // we use terraform import exoscale_ssh_key.my_ssh_key ssh-key-name
	"exoscale_security_group":      config.IdentifierFromProvider, // we use terraform import exoscale_security_group.my_security_group security-group-id
	"exoscale_security_group_rule": config.IdentifierFromProvider, // config.TemplatedStringAsIdentifier("", "{{ .parameters.security_group_id }}/{{ .external_name }}"), // we use terraform import exoscale_security_group_rule.my_security_group_rule <security-group-ID>/<security-group-rule-ID>
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
