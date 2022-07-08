import * as srs from "secure-random-string";
import * as jwt from "jsonwebtoken";
import { getDB } from "./db";

/* tslint:disable:variable-name */
interface Claims {
  session_id: string;
  email: string;
}

export class Session {
  public id: string;
  public expiresAt: Date;
  public userId: string;
  public accessToken: string;

  constructor() {
    this.id = '';
    this.expiresAt = new Date();
    this.userId = '';
    this.accessToken = '';
  }

  public async getToken(): Promise<string> {
    const user: any = {};

    try {
      const claims: Claims = {
        session_id: this.id,
        email: user.email,
      };

      const token = jwt.sign(claims, process.env.SESSION_KEY!);
      return token;
    } catch (err) {
      console.error(err);
      return "";
    }
  }
}

export async function createSession(userId: string, accessToken: string): Promise<Session | undefined> {
  const sess = new Session();
  sess.id = srs.default({ length: 36 });

  const expiresAt = new Date();
  expiresAt.setDate(expiresAt.getDate() + 1);

  sess.expiresAt = expiresAt;
  sess.userId = userId;
  sess.accessToken = accessToken;

  try {
    const db = await getDB();
    await db.query(`insert into session (id, expire_at, user_id, access_token) values ($1, $2, $3, $4)`,
      [sess.id, sess.expiresAt, sess.userId, sess.accessToken]);

    return sess;
  } catch (err) {
    console.error(err);
    return;
  }
}

export async function deleteSession(id: string): Promise<void> {
  try {
    const db = await getDB();
    await db.query(`delete from session where id = $1`, [id]);
  } catch (err) {
    console.error(err);
  }
}

export async function loadSession(token: string): Promise<Session | undefined> {
  try {
    const claims = await jwt.verify(token, process.env.SESSION_KEY!) as jwt.JwtPayload;

    const db = await getDB();
    const result = await db.query(`select id, expire_at, user_id, access_token from session where id = $1`, [claims.session_id]);
    if (!result.rows) {
      return;
    }

    const row = result.rows[0];

    const sess = new Session();
    sess.id = claims.session_id;
    sess.expiresAt = row.expire_at;
    sess.userId = row.user_id;
    sess.accessToken = row.access_token;

    return sess;
  } catch (err) {
    console.error(err);
    return;
  }
}

