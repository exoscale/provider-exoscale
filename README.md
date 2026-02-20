# Provider Exoscale

`provider-exoscale` is a [Crossplane](https://crossplane.io/) provider
exoscale that is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the Exoscale
API.

## Getting Started

This exoscale serves as a starting point for generating a new [Crossplane Provider](https://docs.crossplane.io/latest/packages/providers/) using the [`upjet`](https://github.com/crossplane/upjet) tooling. Please follow the guide linked below to generate a new Provider:

https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md

## Developing

Run code-generation pipeline:
```Bash
$> make generate
```

Check deployed resources
```Bash
&> watch kubectl get managed -A
```

Run against an existing Kubernetes cluster:

```Bash
make run
```

Run e2e test
```Bash
$> export EXOSCALE_API_KEY=...
$> export EXOSCALE_API_SECRET=...
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

Build, push, and install:

```Bash
make all
```

Build binary:

```Bash
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/exoscale/provider-exoscale/issues).
