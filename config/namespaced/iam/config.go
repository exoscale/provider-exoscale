package dbaas

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const shortGroup string = "iam"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_iam_role", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IAMRole"
	})

	p.AddResourceConfigurator("exoscale_iam_api_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IAMAPIKey"

		r.References["role_id"] = config.Reference{
			Type: "github.com/exoscale/provider-exoscale/apis/namespaced/iam/v1alpha1.IAMRole",
		}
	})
}
