NiceDishy is the Web and API side of the nicedishy.com service.
This runs in Kubernetes.


## Generating some data

```
./bin/generate-data --token <token>
```

## Deploy the table schema to timescale


```
create table session (
  id text not null primary key,
  user_id text not null,
  expire_at timestamptz not null,
  access_token text not null
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
  last_login_at timestamptz not null
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
  seconds_obstructed double precision
);
SELECT create_hypertable('dishy_data', 'time'); 
```


## Flux

Flux was installed with the following command (useful when upgrading):

```

```

### Operations

To see 

## Deploying

All CI is handled by GitHub Actions, and the resulting manifests are pushed to https://github.com/marc-campbell/nicedishy-gitops. Weave Flux is in the cluster and syncs from this repo.

