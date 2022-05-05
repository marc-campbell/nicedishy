NiceDishy is the Web and API side of the nicedishy.com service.
This runs in Kubernetes.

## install nats in production

```
$ flux create source helm nats --interval=1h --url https://nats-io.github.io/k8s/helm/charts/
✚ generating HelmRepository source
► applying HelmRepository source
✔ source created
◎ waiting for HelmRepository source reconciliation
✔ HelmRepository source reconciliation completed
✔ fetched revision: e4f79d4f184bc1b4a2ed6ca0095f5d848f53dcc843e637bb6f047c6f7f58d962

$ flux create helmrelease nats --interval=1h --release-name=nats --target-namespace=nats --source=HelmRepository/nats --chart=nats --crds=CreateReplace
✚ generating HelmRelease
► applying HelmRelease
✔ HelmRelease created
◎ waiting for HelmRelease reconciliation
✔ HelmRelease nats is ready
✔ applied revision 0.10.0
```

## Deploy the table schema to timescale




## Flux

Flux was installed with the following command (useful when upgrading):

```

```

### Operations

To see 

## Deploying

All CI is handled by GitHub Actions, and the resulting manifests are pushed to https://github.com/marc-campbell/nicedishy. Weave Flux is in the cluster and syncs from this repo.

