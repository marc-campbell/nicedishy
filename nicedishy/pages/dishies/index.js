import React, { useEffect, useState } from 'react';
import { Utilities } from "../../utils/utilities";
import { useRouter } from 'next/router'
import Layout from "../../components/layout";
import DishyCard from '../../components/dishy-card';
import DishyPlaceholderCard from '../../components/dishy-placeholder-card';
import cookies from 'next-cookies';
import { loadSession } from "../../lib/session";
import { listDishies, getDishyStats, getDishySpeed } from "../../lib/dishy";

export default function Page({dishies, stats, speed}) {
  const router = useRouter();

  const handleAddAnotherClick = (ev) => {
    ev.preventDefault();
    router.push(`/dishy/new`);
  }

  let cards = dishies.map((dishy) => {
    let dishyStats, dishySpeed = {};
    for (const s of stats) {
      if (s.id === dishy.id) {
        dishyStats = s.stats;
      }
    };
    for (const s of speed) {
      if (s.id === dishy.id) {
        dishySpeed = s.speed;
      }
    };

    return <DishyCard key={dishy.id} dishy={dishy} stats={dishyStats} speed={dishySpeed} />
  });

  return (
    <>
      {cards}
      <br /><br />
      <div className="container">
        <div style={{textAlign: "center"}}>
          <a href="#" className="btn btn-outline-secondary" onClick={handleAddAnotherClick}>
            <i className="bi bi-plus-circle"></i>
            {' '}
            Add Another Dishy
          </a>
        </div>
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
      props:{},
    };
  }

  const dishies = await listDishies(sess.userId);

  const stats = await Promise.all(dishies.map(async (dishy) => {
    const stats = await getDishyStats(dishy.id);
    return {
      id: dishy.id,
      stats,
    };
  }));

  const speed = await Promise.all(dishies.map(async (dishy) => {
    const speed = await getDishySpeed(dishy.id);
    return {
      id: dishy.id,
      speed,
    }
  }));


  return {
    props: {
      dishies,
      stats,
      speed,
    },
  };
}
