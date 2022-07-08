import * as srs from "secure-random-string";

import { getDB } from "./db";

export interface User {
  id: string;
  email: string;
  avatarUrl: string;
  createdAt: number;
  lastLoginAt: number;
}

export async function getOrCreateUser(email: string, avatarUrl: string): Promise<User> {
  try {
    const db = await getDB();
    const result = await db.query(`select id from google_user where email_address = $1`, [email]);
    if (result.rowCount > 0) {
      return getUser(result.rows[0].id);
    }

    const id = srs.default({ length: 36 });
    const createdAt = new Date();
    const lastLoginAt = new Date();

    await db.query(`insert into google_user (id, email_address, avatar_url, created_at, last_login_at) values ($1, $2, $3, $4, $5)`,
      [id, email, avatarUrl, createdAt, lastLoginAt]);

    return getUser(id);
  } catch (err) {
    console.error(err);
    throw err;
  }
}

export async function getUser(id: string): Promise<User> {
  try {
    const db = await getDB();

    const result = await db.query(`select id, email_address, avatar_url, created_at, last_login_at from google_user where id = $1`, [id]);
    if (!result.rows) {
      throw new Error(`User ${id} not found`);
    }

    const row = result.rows[0];
    return {
      id: row.id,
      email: row.email_address,
      avatarUrl: row.avatar_url,
      createdAt: row.created_at,
      lastLoginAt: row.last_login_at,
    };
  } catch (err) {
    console.error(err);
    throw err;
  }
}
