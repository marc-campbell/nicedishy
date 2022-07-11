import { createDashboard } from "../../../../lib/dashboard";
import { getDishyUnsafe } from "../../../../lib/dishy";

export default async function handler(req, res) {
  if (req.method !== 'PUT') {
    res.status(405).send({ error: 'Method not allowed' });
    return;
  }

  const authHeader = req.headers["authorization"];
  if (!authHeader) {
    res.status(401).json({message: "Unauthorized"});
    return
  }

  // if (authHeader !== process.env.INTERNAL_AUTH_TOKEN) {
  //   res.status(401).json({message: "Unauthorized"});
  //   return;
  // }

  const dishyId = req.query.id;

  const dishy = await getDishyUnsafe(dishyId);
  if (!dishy) {
    res.status(404).json({message: "Dishy not found"});
    return;
  }

  await createDashboard(dishy.id, dishy.name)

  res.status(204).send();
}
