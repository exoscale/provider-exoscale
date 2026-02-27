# Provider Exoscale

`provider-exoscale` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code generation
tools and exposes XRM-conformant managed resources for the
[Exoscale](https://www.exoscale.com/) API.

It enables you to manage Exoscale cloud infrastructure declaratively from
Kubernetes using the Crossplane ecosystem. Define compute instances, managed
databases, network resources, and more as Kubernetes custom resources, and let
Crossplane handle provisioning and lifecycle management.

## Supported Resources

| Category | Resources |
|----------|-----------|
| **Compute** | Instance, InstancePool, SKSCluster, SKSNodepool, SecurityGroup, SecurityGroupRules, SSHKey, ElasticIP, PrivateNetwork, NLB, NLBService, BlockStorageVolume, AntiAffinityGroup |
| **Database (DBaaS)** | DBAASService (PostgreSQL, MySQL, Kafka, OpenSearch, Grafana, Valkey), DBAASUserPG, DBAASDatabasePG, DBAASUserMySQL, DBAASDatabaseMySQL, DBAASUserKafka, DBAASUserOpenSearch |
| **IAM** | IAMRole, IAMAPIKey, IAMOrgPolicy |

## Getting Started

### Prerequisites

- An existing Kubernetes cluster
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) installed and configured
- [Helm](https://helm.sh/docs/intro/install/) installed
- An [Exoscale](https://portal.exoscale.com/register) account with API credentials

### Install Crossplane

```bash
$> helm repo add crossplane-stable https://charts.crossplane.io/stable
$> helm repo update

$> helm install crossplane crossplane-stable/crossplane \
   --namespace crossplane-system \
   --create-namespace

$> kubectl wait deployment crossplane \
   --namespace crossplane-system \
   --for=condition=Available \
   --timeout=120s
```

### Install the Exoscale Provider

```bash
$> cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-exoscale
spec:
  package: docker.io/exoscale/provider-exoscale:latest
EOF

$> kubectl wait provider provider-exoscale \
   --for=condition=Healthy \
   --timeout=120s
```

Verify the installation:

```bash
$> kubectl get providers
$> kubectl get crds | grep exoscale
```

### Configure the Provider

Create an [IAM API Key](https://community.exoscale.com/documentation/iam/quick-start/) in the Exoscale console, then configure the provider with your credentials:

```bash
$> export EXOSCALE_API_KEY=<your-api-key>
$> export EXOSCALE_API_SECRET=<your-api-secret>

$> kubectl create secret generic exoscale-credentials \
   --namespace crossplane-system \
   --from-literal=credentials="{\"key\": \"$EXOSCALE_API_KEY\", \"secret\": \"$EXOSCALE_API_SECRET\"}"

$> cat <<EOF | kubectl apply -f -
apiVersion: exoscale.m.exoscale.ch/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: exoscale-credentials
      namespace: crossplane-system
      key: credentials
EOF
```

### Create Your First Resource

Once the provider is configured, you can start provisioning Exoscale resources.
Ready-to-use examples for every supported resource are available in the
[`examples/namespaced/`](examples/namespaced/) directory.

<<<<<<< HEAD
For instance, to create a security group:

```bash
$> kubectl apply -f examples/namespaced/compute/v1alpha1/securitygroup.yaml
```

=======
>>>>>>> 32cf914 (feat: update doc)
Monitor the status of your managed resources:

```bash
$> kubectl get managed -A
```
<<<<<<< HEAD
=======

To create an instance:

```bash
$> kubectl apply -f examples/namespaced/compute/v1alpha1/instance.yaml
$> kubectl wait instance.compute.exoscale.m.exoscale.ch/my-instance \
   --namespace crossplane-system \
   --for=condition=Ready \
   --timeout=120s

$> VM_PUBLIC_IP=$(kubectl get instance.compute.exoscale.m.exoscale.ch/my-instance -n crossplane-system -o json | jq '.status.atProvider.publicIpAddress' -r)
$> watch curl $VM_PUBLIC_IP ## might need 1-2 min

## clean up
$> kubectl delete -f examples/namespaced/compute/v1alpha1/instance.yaml
```
>>>>>>> 32cf914 (feat: update doc)

## Developing

> Based on the [Upjet documentation](https://github.com/crossplane/upjet/tree/de59389582c8675a9b2a72a840e084285a3dfb90/docs).

Run the code-generation pipeline:

```bash
$> make generate
```

Run the provider locally against an existing Kubernetes cluster:

```bash
$> make run
```

Check deployed resources:

```bash
$> watch kubectl get managed -A
```

### Updating Examples and Tests

When making changes to resource definitions in the `apis/` directory (e.g. adding
a new resource, renaming a field, or changing defaults), make sure to:

1. Update the corresponding example manifests in
   [`examples/namespaced/`](examples/namespaced/) so they stay in sync with the
   current API schema.
2. Review and update the end-to-end test manifests in
   [`cluster/test/`](cluster/test/) to cover the changes. These manifests are
   used by `make e2e` and must reflect the latest resource specifications.

### End-to-End Tests

```bash
$> export EXOSCALE_API_KEY=<your-api-key>
$> export EXOSCALE_API_SECRET=<your-api-secret>

$> mkdir -p .work
$> cat > .work/uptest_datasource.yaml << EOF
zone: ch-gva-2
suffix: local
EOF

$> make e2e \
   PROVIDER_NAME=provider-exoscale \
   UPTEST_EXAMPLE_LIST=$(find cluster/test/*/*.yaml | tr '\n' ',') \
   UPTEST_CLOUD_CREDENTIALS="{\"key\": \"$EXOSCALE_API_KEY\", \"secret\": \"$EXOSCALE_API_SECRET\"}" \
   UPTEST_DATASOURCE_PATH=./.work/uptest_datasource.yaml
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/exoscale/provider-exoscale/issues).
