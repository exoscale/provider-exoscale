package dbaas

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const shortGroup string = "dbaas"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("exoscale_dbaas", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DBAASService"
	})

	p.AddResourceConfigurator("exoscale_dbaas_pg_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DBAASUserPG"

		r.References["service"] = config.Reference{
			Type: "github.com/exoscale/provider-exoscale/apis/cluster/dbaas/v1alpha1.DBAASService",
		}
	})

	p.AddResourceConfigurator("exoscale_dbaas_pg_database", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DBAASDatabasePG"

		r.References["service"] = config.Reference{
			Type: "github.com/exoscale/provider-exoscale/apis/cluster/dbaas/v1alpha1.DBAASService",
		}
	})

	p.AddResourceConfigurator("exoscale_dbaas_mysql_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DBAASUserMySQL"

		r.References["service"] = config.Reference{
			Type: "github.com/exoscale/provider-exoscale/apis/cluster/dbaas/v1alpha1.DBAASService",
		}
	})
}
