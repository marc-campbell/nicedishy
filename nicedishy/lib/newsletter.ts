const mailchimp = require('@mailchimp/mailchimp_marketing');
const md5 = require("md5")

export async function signUpForNewsletter(email: string, requireOptIn: boolean): Promise<void> {
  const listId = 'f8c85f2a58';

  mailchimp.setConfig({
    apiKey: process.env.MAILCHIMP_API_KEY,
    server: 'us20',
  });

  const subscriberHash = md5(email.toLowerCase());

  const response = await mailchimp.lists.setListMember(
    listId,
    subscriberHash,
    {
      email_address: email,
      status_if_new: requireOptIn ? 'pending' : 'subscribed',
    }
  );
}
