import { createDishyToken, listDishies } from "../../../../lib/dishy";
import { loadSession } from "../../../../lib/session";

export default async function handler(req, res) {
  if (req.method !== 'GET') {
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

  const allDishies = await listDishies(userId);
  for (const d of allDishies) {
    if (d.id === dishyId) {
      // get a token and return i
      const token = await createDishyToken(dishyId);
      res.status(200).json({token});
      return;
    }
  }

  res.status(403).json({message: "Not allowed to create token for dishy"});
}
