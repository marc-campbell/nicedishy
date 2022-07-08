import React, { useEffect, useState } from 'react';
import { Utilities } from "../utils/utilities";
import { useRouter } from 'next/router'
import Layout from "../components/layout";
import { loadSession } from '../lib/session';
import { listDishies } from '../lib/dishy';
import cookies from 'next-cookies';

export default function Page({dishies, authToken}) {
  const router = useRouter();
  const [hasRedirected, setHasRedirected] = useState(false);

  const fetchToken = async(dishyId) => {
    try {
      const res = await fetch(`/api/dishy/${dishyId}/token`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${authToken}`,
        },
      });

      if (!res.ok) {
        return;
      }

      const data = await res.json();

      return data.token;
    } catch (err) {
      console.error(err);
    }
  }

  const handleGoView = () => {
    router.push("/dishies");
  }

  const handleCardClick = async (dishyId) => {
    const token = await fetchToken(dishyId);
    window.location.href = `nicedishy://connected?token=${token}`;
    setHasRedirected(true);
  }

  if (hasRedirected) {
    return (
      <div>
        Ok! You&apos;re connected and data should start coming in.<br />
        <a href="#" className="btn btn-primary" onClick={handleGoView}>
            <i className="bi bi-plus-circle"></i>
            {' '}
            Go view the stats
          </a>
      </div>
    )
  }

  const cards = dishies.map((dishy) => {
    return (
      <a key={dishy.id} href="#" className="text-decoration-none" onClick={handleCardClick.bind(this, dishy.id)}>
        <div className="card card-hover" style={{width: "100%"}}>
          <div className="card-body">
            <div className="row">
              <div className="col-5">
                <h2 className="card-title">{dishy.name}</h2>
              </div>
              <div className="col-6" />
              <div className="col-1">
                <i className="bi bi-chevron-compact-right"></i>
              </div>
            </div>
          </div>
        </div>
      </a>
    );
  });

  return (
    <>
      <div className="container">
        <h2>Connect Device</h2>
        <p>Choose a device to connect this Dishy to</p>
        {cards}
      </div>
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

  const dishies = await listDishies(sess.userId);

  return {
    props: {
      dishies,
      authToken: c.auth,
    }
  }
}
