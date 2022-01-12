import React, { useEffect, useState } from 'react';
import Layout from "../../components/layout";

export default function Page() {

  return (
    <>
      <h1>NiceDishy Documentation</h1>

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
