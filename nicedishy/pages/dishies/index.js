import React, { useEffect, useState } from 'react';
import { Utilities } from "../../utils/utilities";
import { useRouter } from 'next/router'
import Layout from "../../components/layout";
import DishyCard from '../../components/dishy-card';
import DishyPlaceholderCard from '../../components/dishy-placeholder-card';

export default function Page() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(true);
  const [dishies, setDishies] = useState([]);

  const handleAddAnotherClick = () => {
    router.push(`/dishy/new`);
  }

  useEffect( async () => {
    if (!Utilities.getToken()) {
      router.replace(`/login?next=/dishies`);
      return;
    }

    const data = await fetchDishies();
    if (data.dishies.length === 0) {
      router.replace('/dishy/new');
      return;
    }

    setIsLoading(false);
    setDishies(data.dishies);

    setInterval(async () => {
      const data = await fetchDishies();
      if (data.dishies.length === 0) {
        router.replace('/dishy/new');
        return;
      }

      setIsLoading(false);
      setDishies(data.dishies);
    }, 3000);
  }, [])

  const fetchDishies = async() => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/dishies`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
      });

      if (!res.ok) {
        if (res.status === 401) {
          router.push(`/login?next=${encodeURIComponent(router.pathname)}`);
        }
        return;
      }

      const data = await res.json();
      return data;
    } catch (err) {
      console.error(err);
    }
  }

  let cards = [];
  if (isLoading) {
    cards = [<DishyPlaceholderCard key="loading" />];
  } else {
    cards = dishies.map((dishy) => {
      return <DishyCard key={dishy.id} dishy={dishy} />
    });
  }

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
