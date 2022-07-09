const {OAuth2Client} = require('google-auth-library');
import { signUpForNewsletter } from '../../lib/newsletter';

export default async function handler(req, res) {
  if (req.method !== 'POST') {
    res.status(405).send({ error: 'Method not allowed' });
    return;
  }

  const emailAddress = req.body.emailAddress;
  await signUpForNewsletter(emailAddress, true);

  res.status(200).send({redirectURL: authorizeUrl});
}
