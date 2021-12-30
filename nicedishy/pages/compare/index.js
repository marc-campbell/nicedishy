import React from 'react';
import Layout from "../../components/layout";

export default function Page() {
  return (
    <div>
      <h1>Compare Dishy</h1>

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
