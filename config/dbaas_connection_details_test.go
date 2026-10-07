package config

import (
	"reflect"
	"testing"

	ujresource "github.com/crossplane/upjet/v2/pkg/resource"
	dbaasv1alpha1 "github.com/exoscale/provider-exoscale/apis/namespaced/dbaas/v1alpha1"
)

func TestDBAASServiceConnectionDetails(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		attributes map[string]any
		want       map[string][]byte
	}{
		"postgres": {
			attributes: map[string]any{
				"ca_certificate": "certificate",
				"pg": map[string]any{
					"admin_password": "p@ss&word",
					"admin_username": "app user",
				},
				"type": "pg",
				"uri":  "postgres://ignored:credentials@database.example:21699/defaultdb?sslmode=require",
			},
			want: map[string][]byte{
				"clusterCA": []byte("certificate"),
				"dbname":    []byte("defaultdb"),
				"endpoint":  []byte("database.example"),
				"host":      []byte("database.example"),
				"jdbc-uri":  []byte("jdbc:postgresql://database.example:21699/defaultdb?password=p%40ss%26word&sslmode=require&user=app+user"),
				"password":  []byte("p@ss&word"),
				"port":      []byte("21699"),
				"uri":       []byte("postgres://database.example:21699/defaultdb?sslmode=require"),
				"username":  []byte("app user"),
			},
		},
		"scheme-less Kafka URI": {
			attributes: map[string]any{
				"type": "kafka",
				"uri":  "kafka.example:21701",
			},
			want: map[string][]byte{
				"endpoint": []byte("kafka.example"),
				"host":     []byte("kafka.example"),
				"port":     []byte("21701"),
				"uri":      []byte("kafka.example:21701"),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := dbaasServiceConnectionDetails(test.attributes)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("got %#v, want %#v", got, test.want)
			}
		})
	}
}

func TestDBAASResourceConnectionDetails(t *testing.T) {
	t.Parallel()

	user, err := dbaasUserConnectionDetails(map[string]any{
		"username": "application",
		"password": "secret",
	})
	if err != nil {
		t.Fatal(err)
	}
	if want := map[string][]byte{"username": []byte("application"), "password": []byte("secret")}; !reflect.DeepEqual(user, want) {
		t.Fatalf("got user details %#v, want %#v", user, want)
	}

	database, err := dbaasDatabaseConnectionDetails(map[string]any{"database_name": "application"})
	if err != nil {
		t.Fatal(err)
	}
	if want := map[string][]byte{"dbname": []byte("application")}; !reflect.DeepEqual(database, want) {
		t.Fatalf("got database details %#v, want %#v", database, want)
	}
}

func TestDBAASConnectionDetailsConfiguration(t *testing.T) {
	t.Parallel()

	attributes := map[string]any{
		"ca_certificate": "certificate",
		"pg": map[string]any{
			"admin_password": "secret",
			"admin_username": "application",
		},
		"type": "pg",
		"uri":  "postgres://database.example:21699/defaultdb?sslmode=require",
	}
	configured := GetProviderNamespaced().Resources["exoscale_dbaas"]
	got, err := ujresource.GetConnectionDetails(attributes, &dbaasv1alpha1.DBAASService{}, configured)
	if err != nil {
		t.Fatal(err)
	}

	want := map[string][]byte{
		"attribute.ca_certificate":    []byte("certificate"),
		"attribute.pg.admin_password": []byte("secret"),
		"clusterCA":                   []byte("certificate"),
		"dbname":                      []byte("defaultdb"),
		"endpoint":                    []byte("database.example"),
		"host":                        []byte("database.example"),
		"jdbc-uri":                    []byte("jdbc:postgresql://database.example:21699/defaultdb?password=secret&sslmode=require&user=application"),
		"password":                    []byte("secret"),
		"port":                        []byte("21699"),
		"uri":                         []byte("postgres://database.example:21699/defaultdb?sslmode=require"),
		"username":                    []byte("application"),
	}
	if !reflect.DeepEqual(map[string][]byte(got), want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
