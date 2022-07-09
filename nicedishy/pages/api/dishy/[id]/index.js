import { createDishyToken, deleteDishy, getDishy, listDishies } from "../../../../lib/dishy";
import { loadSession } from "../../../../lib/session";

export default async function handler(req, res) {
  if (req.method !== 'DELETE') {
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

  const userId = sess.userId;
  const dishyId = req.query.id;

  const dishy = await getDishy(userId, dishyId);
  if (!dishy) {
    res.status(404).json({message: "Dishy not found"});
    return;
  }

  await deleteDishy(userId, dishyId);

  res.status(204).send();
}
