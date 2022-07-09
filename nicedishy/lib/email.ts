import { getDB } from "./db";
import * as srs from "secure-random-string";

export async function queueEmail(fromAddress: string, toAddress: string, templateId: string, templateContext: any): Promise<void> {
  const id = srs.default({ length: 36 });

  const db = await getDB();
  await db.query(`insert into email_notification
  (id, queued_at, sent_at, error_at, from_address, to_address, template_id, marshalled_context)
  values
  ($1, $2, null, null, $3, $4, $5, $6)`,
    [id, new Date(), fromAddress, toAddress, templateId, JSON.stringify(templateContext)]);
}
