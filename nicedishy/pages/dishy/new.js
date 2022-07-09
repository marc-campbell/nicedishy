import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router'
import Layout from "../../components/layout";
import cookies from 'next-cookies';
import { loadSession } from '../../lib/session';

export default function Page({authToken}) {
  const router = useRouter();

  const [name, setName] = useState("");

  const handleSave = async () => {
    try {
      const res = await fetch(`/api/dishy`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          "Authorization": `Bearer ${authToken}`,
        },
        body: JSON.stringify({
          name,
        }),
      });

      if (!res.ok) {
        console.error('Error creating dishy');
        return;
      }

      const data = await res.json();
      router.push('/dishies');
    } catch(err) {
      console.error(err);
    }
  }

  return (
    <div>
      <h1>Let&apos;s Connect Dishy</h1>
      <p>
        We&apos;ve created a place for your Dishy to send and store performance data.
        All you need to do is give it a name.
      </p>
      <p>
        Stuck? Just name it after your location (maybe your city?).
        It&apos;s just a label and you can always change it later.
      </p>
      <div className="mb-3">
        <input type="email" className="form-control" placeholder="Mammoth Lakes, CA" value={name} onChange={e => setName(e.target.value)} />
        <a href="#" className="btn btn-primary" onClick={handleSave}>Let&apos;s Go</a>
      </div>
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

export async function getServerSideProps(ctx) {
  const c = cookies(ctx);
  const sess = await loadSession(c.auth);
  if (!sess) {
    return {
      redirect: {
        permanent: false,
        destination: "/login",
      },
      props: {},
    };
  }

  return {
    props: {
      authToken: c.auth,
    }
  }
}
