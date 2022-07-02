const {OAuth2Client} = require('google-auth-library');

export default async function handler(req, res) {
  if (req.method !== 'GET') {
    res.status(405).send({ error: 'Method not allowed' });
    return;
  }

  // create the redirect
  const oAuth2Client = new OAuth2Client(
    process.env.GOOGLE_AUTH_CLIENT_ID,
    process.env.GOOGLE_AUTH_CLIENT_SECRET,
    process.env.GOOGLE_AUTH_REDIRECT_URI,
  );

  // Generate the url that will be used for the consent dialog.
  const authorizeUrl = oAuth2Client.generateAuthUrl({
    access_type: 'offline',
    scope: 'https://www.googleapis.com/auth/userinfo.email',
  });

  res.status(200).send({redirectURL: authorizeUrl});
}
