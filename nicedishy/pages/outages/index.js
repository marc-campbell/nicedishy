import React from 'react';
import Layout from "../../components/layout";

export default function Page() {
  return (
    <div>
      <h1>Current Outages</h1>
      <p>coming soon.</p>
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
