import { test, expect } from '@playwright/test';

test('signup', async ({ page }) => {

  // Go to okteto preview url for PR
  await page.goto('/');

  // Click text=Click here to get started
  await page.locator('text=Click here to get started').click();

  // Go to /signup
  await page.goto('/signup');

  // Click img[alt="Sign in with Google"]
  await page.locator('img[alt="Sign in with Google"]').click();
  // await expect(page).toHaveURL('https://accounts.google.com/o/oauth2/v2/auth/identifier?access_type=offline&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email&response_type=code&client_id=807829907877-k47sersiq73tpcqfi68r6671a1t6kmen.apps.googleusercontent.com&redirect_uri=https%3A%2F%2Fnicedishy-marccampbell.cloud.okteto.net%2Flogin%2Fcallback&flowName=GeneralOAuthFlow');

  // Click [aria-label="Email or phone"]
  await page.locator('[aria-label="Email or phone"]').click();

  // Fill [aria-label="Email or phone"]
  await page.locator('[aria-label="Email or phone"]').fill('nicedishy.tests@gmail.com');

  // Click button:has-text("Next")
  await page.locator('button:has-text("Next")').click();
  // await expect(page).toHaveURL('https://accounts.google.com/signin/v2/challenge/pwd?access_type=offline&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email&response_type=code&client_id=807829907877-k47sersiq73tpcqfi68r6671a1t6kmen.apps.googleusercontent.com&redirect_uri=https%3A%2F%2Fnicedishy-marccampbell.cloud.okteto.net%2Flogin%2Fcallback&flowName=GeneralOAuthFlow&cid=1&navigationDirection=forward&TL=AKqFyY-bFcdRAMEx-oTPjkVlb3aXITseninIE8Lp3xYTTRcmLNmSEDHS8o-5_AXR');

  // Click [aria-label="Enter your password"]
  await page.locator('[aria-label="Enter your password"]').click();

  // Fill [aria-label="Enter your password"]
  await page.locator('[aria-label="Enter your password"]').fill('dwNPf9_ngV_H99aArjJcM7');

  // Click button:has-text("Next")
  await page.locator('button:has-text("Next")').click();
  // await expect(page).toHaveURL('https://accounts.youtube.com/accounts/SetSID');

  // // Go to /login/callback?code=4%2F0AdQt8qhmR0leK5joO5fUQsu7sIRyBJGq9_3dGARkY18z9jG--7OJ0pAIMZVJ-_PMBdRQig&scope=email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+openid&authuser=0&prompt=consent
  // await page.goto('/login/callback?code=4%2F0AdQt8qhmR0leK5joO5fUQsu7sIRyBJGq9_3dGARkY18z9jG--7OJ0pAIMZVJ-_PMBdRQig&scope=email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+openid&authuser=0&prompt=consent');

  // // Go to /dishy/new
  await page.screenshot({ path: 'screenshot.png' });
  await page.goto('/dishy/new');

  // // Click [placeholder="Mammoth Lakes\, CA"]
  // await page.locator('[placeholder="Mammoth Lakes\\, CA"]').click();

  // // Fill [placeholder="Mammoth Lakes\, CA"]
  // await page.locator('[placeholder="Mammoth Lakes\\, CA"]').fill('Ashville, NC');

  // // Click text=Let's Go
  // await page.locator('text=Let\'s Go').click();
  // await expect(page).toHaveURL('/dishies');

  // // Click text=Get Connected
  // await page.locator('text=Get Connected').click();

  // // Click text=To connect your Dishy, download and install our app on a laptop or workstation t
  // await page.locator('text=To connect your Dishy, download and install our app on a laptop or workstation t').click();

});
