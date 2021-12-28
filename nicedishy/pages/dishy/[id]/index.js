import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router'
import Layout from "../../../components/layout";
import { Utilities } from '../../../utils/utilities';

export default function Page() {
  const router = useRouter();
  const { id } = router.query

  return (
    <div>
      <h1>Dishy</h1>

    </div>
  );
}

Page.getLayout = function getLayout(page) {
  return (
    <Layout>
      {page}
    </Layout>
  );
}
