const { Pool } = require('pg')

let pool: any;

export async function getDB() {
  if (pool) {
    return pool;
  }

  const p = new Pool({
    connectionString: process.env.DATABASE_URL,
  });

  await p.connect();

  pool = p;
  return pool;
}

