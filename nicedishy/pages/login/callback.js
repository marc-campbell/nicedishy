import React, { useEffect, useState } from 'react';
import * as url from "url";
import { useRouter } from 'next/router'
import posthog from 'posthog-js';

function LoginCallback() {
  const router = useRouter();
  const [authComplete, setAuthComplete] = useState(false);
  const [nextUrl, setNextUrl] = useState("/");

  const requestSessionToken = async (code, state) => {
    const uri = `/api/login_callback`;
    const response = await fetch(uri, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        code,
        state,
      }),
    });

    if (!response.ok) {
      return "";
    }

    const body = await response.json();
    return body;
  }

  useEffect( async () => {
    if (window.localStorage.getItem("token")) {
      setAuthComplete(true);

      const next = sessionStorage.getItem("next");
      if (next) {
        router.push(decodeURIComponent(next));
      } else {
        router.push("/dishies");
      }
      return;
    }

    try {
      const query = url.parse(window.location.href, true).query;
      const response = await requestSessionToken(query.code, query.state);
      if (!response ) {
        router.replace("/");
        return;
      }

      setAuthComplete(true);

      if (process.env.NEXT_PUBLIC_POSTHOG_KEY) {
        posthog.init(process.env.NEXT_PUBLIC_POSTHOG_KEY, {
          api_host: 'https://app.posthog.com',
          loaded: function(posthog) {
            posthog.identify(
              `user:${response.userId}`,
              { email: response.emailAddress },
            );
            router.push(response.nextUrl);
          }
        });
      } else {
        router.push(response.nextUrl);
      }
    } catch (err) {
      console.log(err);
      router.replace("/error");
      return;
    }
  }, []);

  return (
    <>
    </>
  );
}

export default LoginCallback;
