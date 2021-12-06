import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router'
import Layout from "../../components/layout";
import { Utilities } from '../../utils/utilities';

export default function Page() {
  const router = useRouter();

  const [name, setName] = useState("");

  const handleSave = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/dishy`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': Utilities.getToken(),
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
      <h1>Let's Connect Dishy</h1>
      <p>
        We've created a place for your Dishy to send and store performance data.
        All you need to do is give it a name.
      </p>
      <p>
        Stuck? Just name it after your location (maybe your city?).
        It's just a label and you can always change it later.
      </p>
      <div className="mb-3">
        <input type="email" className="form-control" placeholder="Mammoth Lakes, CA" value={name} onChange={e => setName(e.target.value)} />
        <a href="#" className="btn btn-primary" onClick={handleSave}>Let's Go</a>
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
