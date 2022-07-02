import Cookies from 'cookies';

export default async function handler(req, res) {
  if (req.method !== 'POST') {
    res.status(405).send({ error: 'Method not allowed' });
    return;
  }

  const cookies = new Cookies(req, res)

  cookies.set('auth', null, {
    httpOnly: true,
  });

  res.status(200).send({});
}
