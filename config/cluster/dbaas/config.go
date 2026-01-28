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
}
