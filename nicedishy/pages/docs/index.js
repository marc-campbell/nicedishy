import React, { useEffect, useState } from 'react';
import Layout from "../../components/layout";

export default function Page() {

  return (
    <>
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

export async function getServerSideProps(ctx) {
  return {
    redirect: {
      permanent: false,
      destination: "https://docs.nicedishy.com",
    },
    props: {},
  };
}
