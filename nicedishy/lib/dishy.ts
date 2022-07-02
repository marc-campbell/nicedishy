import { getDB } from "./db";

export interface Dishy {
  id: string;
  name: string;
  createdAt: string;
  lastMetricAt: string;
  lastGeocheckAt: string;
}

export interface DishyStats {
  snr: number;
  downlinkThroughputBps: number;
  uplinkThroughputBps: number;
  popPingLatencyMs: number;
  popPingDropRate: number;
  percentObstructed: number;
  obstructedSeconds: number;
}

export interface DishySpeed {
  downloadSpeed: number;
  uploadSpeed: number;
}

export async function listDishies(userId: string): Promise<Dishy[]> {
  const db = await getDB();

  const result = await db.query(`select id, name, created_at, last_metric_at, last_geocheck_at from dishy where user_id = $1`, [userId]);
  if (!result.rows) {
    return [];
  }

  return result.rows.map((row: any) => ({
    id: row.id,
    name: row.name,
    createdAt: new Date(row.created_at).toISOString(),
    lastMetricAt: new Date(row.last_metric_at).toISOString(),
    lastGeocheckAt: new Date(row.last_geocheck_at).toISOString(),
  }));
}

export async function getDishySpeed(id: string): Promise<DishySpeed | undefined> {
  const db = await getDB();

  const result = await db.query(`select download_speed, upload_speed from dishy_speed where dishy_id = $1 and download_speed is not null
order by time desc limit 1`, [id]);
  if (!result.rows) {
    return undefined;
  }

  const row = result.rows[0];

  return {
    downloadSpeed: row.download_speed,
    uploadSpeed: row.upload_speed,
  };
}

export async function getDishyStats(id: string): Promise<DishyStats | undefined> {
  const db = await getDB();

  const result = await db.query(`select
snr, downlink_throughput_bps, uplink_throughput_bps, pop_ping_latency_ms,
pop_ping_drop_rate, percent_obstructed, seconds_obstructed
from dishy_data
where dishy_id = $1
and downlink_throughput_bps is not null
order by time desc limit 1`, [id]);

  if (!result.rows) {
    return;
  }

  const row = result.rows[0];

  return {
    snr: row.snr,
    downlinkThroughputBps: row.downlink_throughput_bps,
    uplinkThroughputBps: row.uplink_throughput_bps,
    popPingLatencyMs: row.pop_ping_latency_ms,
    popPingDropRate: row.pop_ping_drop_rate,
    percentObstructed: row.percent_obstructed,
    obstructedSeconds: row.seconds_obstructed,
  };
}