import React, { useEffect, useState } from 'react';
import * as url from "url";
import { Utilities } from "../../utils/utilities";
import Layout from "../../components/layout";
import { useRouter } from 'next/router'

export default function Page() {
  const router = useRouter();

  const onClickSignup = async (ev) => {
    ev.preventDefault();

    // ensure the user is logged out
    Utilities.logoutUser();

    // set or clear the next url from sessionstorage
    window.sessionStorage.removeItem('next');
    const query = url.parse(window.location.href, true).query;
    if (query && query.next) {
      window.sessionStorage.setItem('next', query.next);
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/login?next=${encodeURIComponent(url.parse(window.location.href, true).query.next)}`);
      if (!res.ok) {
        console.log("error")
        return;
      }

      const data = await res.json();
      window.location.href = data.redirectURL;
    } catch(err) {
      console.log(err);
    }
  }

  useEffect( () => {
    if (Utilities.getToken()) {
      router.push("/dishies");
    }
  })

  return (
    <>
      <h1>Create a NiceDishy account</h1>
      <p>
        By signing up, you are agreeing to our Terms of Service and Privacy Policy. We ask
        for read access to your Google profile in order to provide a complete experience
        here. We don&apos;t ask for permissions to change anything in your Google account.
      </p>
      <a href="#" width="80%" onClick={onClickSignup}>
        <img src="/images/btn_google_signin_dark_pressed_web@2x.png" alt="Sign in with Google" style={{width: "200px"}} />
      </a>
    </>
  );
}

Page.getLayout = function getLayout(page) {
  return (
    <Layout>
      {page}
    </Layout>
  );
}
