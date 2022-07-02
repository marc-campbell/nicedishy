const {OAuth2Client} = require('google-auth-library');
import { createSession } from '../../lib/session';
import { getOrCreateUser } from '../../lib/user';
import Cookies from 'cookies';

export default async function handler(req, res) {
  if (req.method !== 'POST') {
    res.status(405).send({ error: 'Method not allowed' });
    return;
  }

  const code = req.body.code;

  const oAuth2Client = new OAuth2Client(
    process.env.GOOGLE_AUTH_CLIENT_ID,
    process.env.GOOGLE_AUTH_CLIENT_SECRET,
    process.env.GOOGLE_AUTH_REDIRECT_URI,
  );

  const r = await oAuth2Client.getToken(code);
  oAuth2Client.setCredentials(r.tokens);

  const accessToken = await oAuth2Client.getAccessToken();

  // create a session
  const result = await oAuth2Client.request({
    url: `https://www.googleapis.com/oauth2/v3/userinfo`,
  });

  const u = await getOrCreateUser(result.data.email, result.data.picture);

  const sess = await createSession(u.id, accessToken.token);

  const token = await sess.getToken();
  const cookies = new Cookies(req, res)

  cookies.set('auth', token, {
    httpOnly: true,
  });

  res.status(200).send({
    nextUrl: '/dishies',
  });
}
