import React, { useEffect } from 'react';
import Layout from "../components/layout";
import { useRouter } from 'next/router'
import { Utilities } from '../utils/utilities';

export default function Page() {
  const router = useRouter();

  useEffect(async () => {
    await Utilities.logoutUser();
  }, []);

  return (
    <div></div>
  );
}

Page.getLayout = function getLayout(page) {
  return (
    <Layout isLoggedIn>
      {page}
    </Layout>
  );
}


