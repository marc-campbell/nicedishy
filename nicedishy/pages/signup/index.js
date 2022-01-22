import React, { useEffect, useState } from 'react';
import * as url from "url";
import { Utilities } from "../../utils/utilities";
import Layout from "../../components/layout";
import { useRouter } from 'next/router'
import Link from 'next/link'
import Image from 'next/image';

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

  useEffect( async () => {
    if (Utilities.getToken()) {
      // validate that the token is still valid
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/whoami`, {
         method: 'GET',
         headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
          },
        });

        if (!res.ok) {
          console.error("error");
          return;
        }

        const data = await res.json();
        if (data.user) {
          if (data.user.isWaitlisted) {
            router.push("/waitlist");
          } else {
            router.push("/dishies");
          }
          return;
        }
      } catch(err) {
        console.log(err);
      }
      Utilities.logoutUser();
    }
  })

  return (
    <>
    <div className="row">
      <div className="col-6 offset-3">
        <h1>Create a NiceDishy account</h1>
        <p>
          By signing up, you are agreeing to our <Link href="/tos">Terms of Service</Link> and <Link href="/privacy">Privacy Policy</Link>.
          We ask for read access to your Google profile in order to provide a complete experience
          here. We don&apos;t ask for permissions to change anything in your Google account.
        </p>
        <p>
          Note: you are signing up for the waitlist. We will be opening the service to additional
          folks as soon as possible, but right now, signing up will add you to the waitlist.
        </p>
        <p>
          The best way to get to the top of the waitlist is to sign up and answer the quick questions
          that follow. We promise that it&apos;s quick, and will be picking the best use cases next.
        </p>
      </div>
    </div>
    <div className="row">
      <div className="col-12 text-center">
        <a href="#" width="80%" onClick={onClickSignup}>
          <Image src="/images/btn_google_signin_dark_pressed_web@2x.png" alt="Sign in with Google" width="200px" height="50px" />
        </a>
      </div>
    </div>
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
