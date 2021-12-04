import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router'
import Layout from "../../components/layout";
import { FormGroup, TextInput, Button } from '@primer/components';

export default function Page() {
  const router = useRouter();

  const [name, setName] = useState("");

  const handleSave = async () => {
    console.log('here');
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
      <FormGroup>
        <TextInput id="example-text" value={name} onChange={e => setName(e.target.value)} />
        <Button onClick={handleSave}>Let's Go</Button>
      </FormGroup>
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
