import { createDishy } from "../../../lib/dishy";
import { loadSession } from "../../../lib/session";

export default async function handler(req, res) {
  if (req.method !== 'POST') {
    res.status(405).send({ error: 'Method not allowed' });
    return;
  }

  const authHeader = req.headers["authorization"];
  if (!authHeader) {
    res.status(401).json({message: "Unauthorized"});
    return
  }

  const tokenParts = authHeader.split(" ");
  if (tokenParts.length !== 2) {
    res.status(401).json({message: "Unauthorized"});
    return;
  }

  const token = tokenParts[1];
  const sess = await loadSession(token);

  const dishy = await createDishy(sess.userId, req.body.name);

  res.status(200).json({
    dishy,
  });
}
