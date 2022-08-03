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
  longitude double precision not null,
  timezone_id text,
  timezone_abbr text,
  timezone_offset int,
  timezone_utc text
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
  hardware_version text not null,
  user_agent text
);
SELECT create_hypertable('dishy_data', 'time'); 

create table dishy_data_hourly (
  time_start timestamptz not null,
  dishy_id text not null,
  snr integer,
  downlink_throughput_bps double precision,
  uplink_throughput_bps double precision,
  pop_ping_latency_ms double precision,
  pop_ping_drop_rate double precision,
  percent_obstructed double precision,
  seconds_obstructed double precision,
  primary key (time_start, dishy_id)
);
SELECT create_hypertable('dishy_data_hourly', 'time_start'); 


create table dishy_speed (
  time timestamptz not null,
  dishy_id text not null,
  ip_address text not null,
  download_speed double precision,
  upload_speed double precision,
  software_version text not null,
  hardware_version text not null,
  user_agent text
);
SELECT create_hypertable('dishy_speed', 'time'); 

create table dishy_speed_hourly (
  time_start timestamptz not null,
  dishy_id text not null,
  download_speed double precision,
  upload_speed double precision,
  primary key (time_start, dishy_id)
);
SELECT create_hypertable('dishy_speed_hourly', 'time_start'); 

create table dishy_speed_fourhour (
  time_start timestamptz not null,
  dishy_id text not null,
  download_speed double precision,
  upload_speed double precision,
  primary key (time_start, dishy_id)
);
SELECT create_hypertable('dishy_speed_fourhour', 'time_start'); 

create table dishy_data_fourhour (
  time_start timestamptz not null,
  dishy_id text not null,
  snr integer,
  downlink_throughput_bps double precision,
  uplink_throughput_bps double precision,
  pop_ping_latency_ms double precision,
  pop_ping_drop_rate double precision,
  percent_obstructed double precision,
  seconds_obstructed double precision,
  primary key (time_start, dishy_id)
);
SELECT create_hypertable('dishy_data_fourhour', 'time_start'); 

create table dishy_speed_daily (
  time_start timestamptz not null,
  dishy_id text not null,
  download_speed double precision,
  upload_speed double precision,
  primary key (time_start, dishy_id)
);
SELECT create_hypertable('dishy_speed_daily', 'time_start'); 

create table dishy_data_daily (
  time_start timestamptz not null,
  dishy_id text not null,
  snr integer,
  downlink_throughput_bps double precision,
  uplink_throughput_bps double precision,
  pop_ping_latency_ms double precision,
  pop_ping_drop_rate double precision,
  percent_obstructed double precision,
  seconds_obstructed double precision,
  primary key (time_start, dishy_id)
);
SELECT create_hypertable('dishy_data_daily', 'time_start'); 


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

create table dishy_disconnected_queue (
  dishy_id text not null primary key,
  send_at timestamp not null
);

create table dishy_report_weekly (
  dishy_id text not null,
  week_start timestamptz not null,
  week_end timestamptz not null,
  report_context text not null,
  is_generating boolean not null,
  primary key (dishy_id, week_start)
);




CREATE OR REPLACE FUNCTION download_speed(
  IN id TEXT, 
  IN min_time BIGINT,
  IN max_time BIGINT
)
RETURNS table (
    time_start timestamptz, 
    download_speed double precision
)
LANGUAGE plpgsql
AS $$
DECLARE 
    tabname varchar;
    interval bigint;
    timecolumn varchar;
BEGIN
    interval := max_time - min_time;
    tabname := 'dishy_speed';
    timecolumn := 'time';

    IF interval >= 1209600000 THEN
      tabname := 'dishy_speed_daily';      
      timecolumn := 'time_start';
    ELSIF interval >= 604800000 THEN
      tabname := 'dishy_speed_fourhour';
      timecolumn := 'time_start';
    ELSIF interval >= 172800000 THEN 
      tabname := 'dishy_speed_hourly';
      timecolumn := 'time_start';
    END IF;

    return query EXECUTE '
        select 
            ' 
            || quote_ident(timecolumn) 
            || ' as time_start,
            download_speed
        from '
        || quote_ident(tabname) 
        || ' where dishy_id = $1 and '
        || quote_ident(timecolumn) 
        || ' >= to_timestamp($2/1000) and '
        || quote_ident(timecolumn) 
        || ' < to_timestamp($3/1000) order by '
        || quote_ident(timecolumn)
        || ' desc;'
        using id, min_time, max_time;
END;
$$;


CREATE OR REPLACE FUNCTION upload_speed(
  IN id TEXT, 
  IN min_time BIGINT,
  IN max_time BIGINT
)
RETURNS table (
    time_start timestamptz, 
    upload_speed double precision
)
LANGUAGE plpgsql
AS $$
DECLARE 
    tabname varchar;
    interval bigint;
    timecolumn varchar;
BEGIN
    interval := max_time - min_time;
    tabname := 'dishy_speed';
    timecolumn := 'time';

    IF interval >= 1209600000 THEN
      tabname := 'dishy_speed_daily';      
      timecolumn := 'time_start';
    ELSIF interval >= 604800000 THEN
      tabname := 'dishy_speed_fourhour';
      timecolumn := 'time_start';
    ELSIF interval >= 172800000 THEN 
      tabname := 'dishy_speed_hourly';
      timecolumn := 'time_start';
    END IF;

    return query EXECUTE '
        select 
            ' 
            || quote_ident(timecolumn) 
            || ' as time_start,
            upload_speed
        from '
        || quote_ident(tabname) 
        || ' where dishy_id = $1 and '
        || quote_ident(timecolumn) 
        || ' >= to_timestamp($2/1000) and '
        || quote_ident(timecolumn) 
        || ' < to_timestamp($3/1000) order by '
        || quote_ident(timecolumn) 
        || ' desc;'
        using id, min_time, max_time;
END;
$$;

q

% select time_start,download_speed from download_speed('rVe0QHfn7N6NM-XbuSItiU5v2bklb3-WJHwb', 1655292003329, 1657884003329);
select time_start as "time",
download_speed from download_speed('28suwqfvNtHj31El3d6JniEnwgb', 1657886376257, 1658491176257) order by 1;

delete from dishy_data; delete from dishy_data_hourly; delete from dishy_data_fourhour; delete from dishy_speed; delete from dishy_speed_daily; delete from dishy_speed_hourly; update dishy set last_metric_at = null; delete from dishy_speed_fourhour;


select * from download_speed('zVByq3ggfnyYhVzs2iblkhcCDLWx4uVfig_p', 1659446478000, 1659532878000);
        select time as time_start,
        download_speed
        from dishy_speed
        where dishy_id = 'zVByq3ggfnyYhVzs2iblkhcCDLWx4uVfig_p'
        and time >= to_timestamp(1659446478000/1000) 
        and time < to_timestamp(1659532878000/1000) order by time desc;

        

  select * from table_name where time >= to_timestamp(1659446478000/1000) and time < to_timestamp(1659532878000/1000) order by time desc;
