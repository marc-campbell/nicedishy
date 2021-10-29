NiceDishy is the Web and API side of the nicedishy.com service.
This runs in Kubernetes.

## Setting up a dev environment
1. Use codespaces
1. Build (`make`)
1. `skaffold dev`

## Local dev ports

## Generating some data

```
./bin/generate-data --token <token>
```

## Deploy the table schema to timescale

```
create table dishy_data (
  time timestamptz not null,
  dishy_id text not null,
  state text,
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