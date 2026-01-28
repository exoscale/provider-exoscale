package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	clustercompute "github.com/exoscale/provider-exoscale/config/cluster/compute"
	clusterdbaas "github.com/exoscale/provider-exoscale/config/cluster/dbaas"
	namespacedcompute "github.com/exoscale/provider-exoscale/config/namespaced/compute"
	namespaceddbaas "github.com/exoscale/provider-exoscale/config/namespaced/dbaas"
)

const (
	resourcePrefix = "exoscale"
	modulePath     = "github.com/exoscale/provider-exoscale"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("exoscale.exoscale.ch"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		clustercompute.Configure,
		clusterdbaas.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("exoscale.m.exoscale.ch"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		namespacedcompute.Configure,
		namespaceddbaas.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
