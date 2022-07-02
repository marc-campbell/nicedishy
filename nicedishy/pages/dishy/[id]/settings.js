import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router'
import Layout from "../../../components/layout";
import { Utilities } from '../../../utils/utilities';
import cookies from 'next-cookies';
import { loadSession } from "../../../lib/session";
import { listDishies } from '../../../lib/dishy';

export default function Page({dishy, settings}) {
  const router = useRouter();

  const [name, setName] = useState(dishy.name);

  const handleSave = async () => {

  }

  const handleDelete = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/dishy/${id}`, {
        method: 'DELETE',
        headers: {
          'Authorization': Utilities.getToken(),
        },
      });

      if (res.status !== 204) {
        console.error('Error deleting dishy');
        return;
      }

      router.push('/dishies');
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div>
      <h1>Dishy Settings</h1>
      <div className="container">
        <div className="row" style={{paddingTop: "60px", paddingBottom: "60px"}}>
          <div className="col-lg-5">
            <h3>Dishy Settings</h3>
            <form>
              <div className="form-group">
                <label htmlFor="name">Name</label>
                <input type="name" className="form-control" id="name" placeholder="Mammoth Lakes, CA" value={name} onChange={e => setName(e.target.value)}/>
              </div>
            </form>
          </div>
          <div className="row" style={{paddingTop: "60px", paddingBottom: "60px"}}>
            <div className="col-lg-12">
              <h3>Danger Zone</h3>
              <div className="row">
                <div className="col-lg-10">
                  Delete dishy. This is a permanent action and cannot be undone. Tread carefully!
                </div>
                <div className="col-lg-2">
                  <div className="btn btn-danger" onClick={handleDelete}>Delete</div>
                </div>
              </div>
            </div>
          </div>
        </div>
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
      props:{},
    };
  }

  const dishies = await listDishies(sess.userId);
  const dishy = dishies.find(d => d.id === ctx.query.id);
  if (!dishy) {
    return {
      redirect: {
        permanent: false,
        destination: "/dishies",
      },
      props:{},
    };
  }

  return {
    props: {
      id: ctx.query.id,
      dishy,
    },
  }
}
