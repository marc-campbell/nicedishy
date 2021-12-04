import React, { useEffect, useState } from 'react';
import * as url from "url";
import { useRouter } from 'next/router'

function LoginCallback() {
  const router = useRouter();
  const [authComplete, setAuthComplete] = useState(false);
  const [nextUrl, setNextUrl] = useState("/");

  const requestSessionToken = async (code, state) => {
    const uri = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/login/callback`;
    const response = await fetch(uri, {
      method: "post",
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
      router.push("/dishies");
      return;
    }

    let nextUrl = "";
    try {
      const query = url.parse(window.location.href, true).query;
      const response = await requestSessionToken(query.code, query.state);
      if (!response ) {
        router.replace("/");
        return;
      }

      window.localStorage.setItem("token", response.token);
      nextUrl = !response.redirectUri ? "/dishies" : response.redirectUri;
    } catch (err) {
      console.log(err);
      router.replace("/error");
      return;
    }

    setAuthComplete(true);
    setNextUrl(nextUrl);

    router.push(nextUrl);
  })

  return (
    <>
    </>
  );
}

export default LoginCallback;
