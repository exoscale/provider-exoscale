package config

import (
	"fmt"
	"net/url"
	"strings"

	commonv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	connectionKeyDBName  = "dbname"
	connectionKeyHost    = "host"
	connectionKeyJDBCURI = "jdbc-uri"
	connectionKeyURI     = "uri"
)

func configureDBAASConnectionDetails(p *ujconfig.Provider) {
	p.AddResourceConfigurator("exoscale_dbaas", func(r *ujconfig.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = dbaasServiceConnectionDetails
	})

	for _, resource := range []string{
		"exoscale_dbaas_kafka_user",
		"exoscale_dbaas_mysql_user",
		"exoscale_dbaas_opensearch_user",
		"exoscale_dbaas_pg_user",
	} {
		p.AddResourceConfigurator(resource, func(r *ujconfig.Resource) {
			r.Sensitive.AdditionalConnectionDetailsFn = dbaasUserConnectionDetails
		})
	}

	for _, resource := range []string{
		"exoscale_dbaas_mysql_database",
		"exoscale_dbaas_pg_database",
	} {
		p.AddResourceConfigurator(resource, func(r *ujconfig.Resource) {
			r.Sensitive.AdditionalConnectionDetailsFn = dbaasDatabaseConnectionDetails
		})
	}
}

func dbaasServiceConnectionDetails(attr map[string]any) (map[string][]byte, error) {
	details := map[string][]byte{}
	addConnectionDetail(details, commonv1.ResourceCredentialsSecretCAKey, stringAttribute(attr, "ca_certificate"))

	serviceType := stringAttribute(attr, "type")
	service, _ := attr[serviceType].(map[string]any)
	username := stringAttribute(service, "admin_username")
	password := stringAttribute(service, "admin_password")
	addConnectionDetail(details, commonv1.ResourceCredentialsSecretUserKey, username)
	addConnectionDetail(details, commonv1.ResourceCredentialsSecretPasswordKey, password)

	rawURI := stringAttribute(attr, "uri")
	if rawURI == "" {
		return details, nil
	}

	parseURI := rawURI
	hasScheme := strings.Contains(parseURI, "://")
	if !hasScheme {
		parseURI = "//" + parseURI
	}
	parsed, err := url.Parse(parseURI)
	if err != nil {
		return nil, fmt.Errorf("cannot parse DBaaS service URI %q: %w", rawURI, err)
	}
	parsed.User = nil
	if hasScheme {
		details[connectionKeyURI] = []byte(parsed.String())
	} else {
		details[connectionKeyURI] = []byte(rawURI)
	}

	host := parsed.Hostname()
	addConnectionDetail(details, commonv1.ResourceCredentialsSecretEndpointKey, host)
	addConnectionDetail(details, connectionKeyHost, host)
	addConnectionDetail(details, commonv1.ResourceCredentialsSecretPortKey, parsed.Port())

	if dbName := strings.TrimPrefix(parsed.Path, "/"); dbName != "" {
		details[connectionKeyDBName] = []byte(dbName)
	}

	if serviceType == "pg" {
		jdbc := *parsed
		jdbc.Scheme = "postgresql"
		query := jdbc.Query()
		if username != "" {
			query.Set("user", username)
		}
		if password != "" {
			query.Set("password", password)
		}
		jdbc.RawQuery = query.Encode()
		details[connectionKeyJDBCURI] = []byte("jdbc:" + jdbc.String())
	}

	return details, nil
}

func dbaasUserConnectionDetails(attr map[string]any) (map[string][]byte, error) {
	details := map[string][]byte{}
	addConnectionDetail(details, commonv1.ResourceCredentialsSecretUserKey, stringAttribute(attr, "username"))
	addConnectionDetail(details, commonv1.ResourceCredentialsSecretPasswordKey, stringAttribute(attr, "password"))
	return details, nil
}

func dbaasDatabaseConnectionDetails(attr map[string]any) (map[string][]byte, error) {
	details := map[string][]byte{}
	addConnectionDetail(details, connectionKeyDBName, stringAttribute(attr, "database_name"))
	return details, nil
}

func stringAttribute(attr map[string]any, key string) string {
	value, _ := attr[key].(string)
	return value
}

func addConnectionDetail(details map[string][]byte, key, value string) {
	if value != "" {
		details[key] = []byte(value)
	}
}
