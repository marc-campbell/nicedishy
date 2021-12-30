import React from 'react';
import Layout from "../../../components/layout";

export default function Page() {

  return (
    <div>
      test
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


