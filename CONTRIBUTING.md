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


```
create table session (
  id text not null primary key,
  user_id text not null,
  expire_at timestamptz not null,
  access_token text not null
);

create table session_nonce (
  id text not null primary key,
  expire_at timestamptz not null,
  session_id text not null
);

create table oauth_state (
  id text not null primary key,
  created_at timestamptz not null,
  next text not null
);

create table google_user (
  id text not null primary key,
  email_address text not null,
  avatar_url text not null,
  created_at timestamptz not null,
  last_login_at timestamptz not null,
  is_waitlisted boolean not null
);

create table dishy (
  id text not null primary key,
  user_id text not null,
  created_at timestamptz not null,
  name text not null,
  last_metric_at timestamptz,
  last_geocheck_at timestamptz
);

create table dishy_geo (
  time timestamptz not null,
  id text not null,
  ip_address text not null,
  continent text not null,
  country text not null,
  region text not null,
  city text not null,
  org text not null,
  latitude double precision not null,
  longitude double precision not null
);
select create_hypertable('dishy_geo', 'time');

create table dishy_token (
  token_sha text not null primary key,
  dishy_id text not null
);

create table dishy_data (
  time timestamptz not null,
  dishy_id text not null,
  ip_address text not null,
  snr integer,
  downlink_throughput_bps double precision,
  uplink_throughput_bps double precision,
  pop_ping_latency_ms double precision,
  pop_ping_drop_rate double precision,
  percent_obstructed double precision,
  seconds_obstructed double precision,
  software_version text not null,
  hardware_version text not null
);
SELECT create_hypertable('dishy_data', 'time'); 

create table dishy_speed (
  time timestamptz not null,
  dishy_id text not null,
  ip_address text not null,
  download_speed double precision,
  upload_speed double precision,
  software_version text not null,
  hardware_version text not null
);
SELECT create_hypertable('dishy_speed', 'time'); 

create table google_user_waitlist (
  id text not null primary key,
  created_at timestamptz,
  already_have_dishy text not null,
  how_long_with_dishy text not null,
  primary_or_backup text not null,
  operating_systems text[] not null,
  why_access text[] not null
);

create table waitlist_approved (
  id text not null primary key,
  approved_at timestamptz
);

create table email_notification (
  id text not null primary key,
  queued_at timestamptz not null,
  sent_at timestamptz,
  error_at timestamptz,
  from_address text not null,
  to_address text not null,
  template_id text not null,
  marshalled_context text not null
);
```


## Flux

Flux was installed with the following command (useful when upgrading):

```

```

### Operations

To see 

## Deploying

All CI is handled by GitHub Actions, and the resulting manifests are pushed to https://github.com/marc-campbell/nicedishy-gitops. Weave Flux is in the cluster and syncs from this repo.

