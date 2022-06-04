import React, { useEffect, useState } from 'react';
import Link from 'next/link'
import { Utilities } from "../../utils/utilities";
import Layout from "../../components/layout";
import { useRouter } from 'next/router'
import Image from 'next/image';
import cookies from 'next-cookies';
import { loadSession } from "../../lib/session";

export default function Page() {
  const router = useRouter();

  const onClickLogin = async (ev) => {
    ev.preventDefault();

    // ensure the user is logged out
    Utilities.logoutUser();

    try {
      const res = await fetch(`/api/login`);
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
      async function fetchData() {
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
            router.push("/dishies");
            return;
          }
        } catch(err) {
          console.log(err);
        }
        Utilities.logoutUser();
      }
      fetchData();
    }
  })

  return (
    <>
      <div className="row">
        <div className="col-6 offset-3">
          <h1>Create a NiceDishy account</h1>
          <p>
          By signing up, you are agreeing to our <Link href="/terms">Terms of Service</Link> and <Link href="/privacy">Privacy Policy</Link>. I ask for read access to your Google profile in order to provide a complete experience here. I don&apos;t ask for permissions to change anything in your Google account.
          </p>
        </div>
      </div>
      <div className="row">
        <div className="col-12 text-center">
          <a href="#" width="80%" onClick={onClickLogin}>
            <Image src="/images/btn_google_signin_dark_pressed_web@2x.png" alt="Sign in with Google" width="200px" height="50px" />
          </a>
        </div>
      </div>
    </>
  );
}

Page.getLayout = function getLayout(page) {
  return (
    <Layout isLoggedIn>
      {page}
    </Layout>
  );
}

export async function getServerSideProps(ctx) {
  const c = cookies(ctx);
  const sess = await loadSession(c.auth);
  if (sess) {
    return {
      redirect: {
        permanent: false,
        destination: "/dishies",
      },
      props: {},
    };
  }

  return {
    props: {},
  };
}
