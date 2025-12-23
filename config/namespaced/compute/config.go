package repository

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

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
			Type: "github.com/exoscale/provider-exoscale/apis/namespaced/compute/v1alpha1.SecurityGroup",
		}

		// Ignore the deprecated security_group field to avoid conflicts
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"security_group"},
		}

		// // Custom external name config to handle the composite ID format
		// r.ExternalName = config.IdentifierFromProvider
		// r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		// 	// The actual ID from Terraform state
		// 	id, ok := tfstate["id"].(string)
		// 	if !ok {
		// 		return "", errors.New("id in tfstate is not string")
		// 	}
		// 	return id, nil
		// }

		// r.ExternalName.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]any, setup map[string]any) (string, error) {
		// 	// For import: combine security_group_id with the rule ID
		// 	sgID, ok := parameters["security_group_id"].(string)
		// 	if !ok {
		// 		return "", errors.New("security_group_id is required")
		// 	}
		// 	return fmt.Sprintf("%s/%s", sgID, externalName), nil
		// }
	})
}
