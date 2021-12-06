import React, { useEffect, useState } from 'react';
import { Utilities } from "../utils/utilities";
import { useRouter } from 'next/router'
import Layout from "../components/layout";

export default function Page() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [dishies, setDishies] = useState([]);

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
        return;
      }

      const data = await res.json();
      return data;
    } catch (err) {
      console.error(err);
    }
  }

  const handleCardClick = async (dishyId) => {
    // TODO get token
    window.location.href='nicedishy://connected?token=ab'
  }

  useEffect( async () => {
    if (!Utilities.getToken()) {
      router.replace(`/login?next=/connect_device`);
      return;
    }

    const data = await fetchDishies();
    if (data.dishies.length === 0) {
      router.replace('/dishy/new');
      return;
    }
    setIsLoading(false);
    setDishies(data.dishies);
  }, []);

  if (isLoading) {
    return (
      <div>
        loading...
      </div>
    );
  }

  const cards = dishies.map((dishy) => {
    return (
      <a key={dishy.id} href="#" className="text-decoration-none" onClick={handleCardClick.bind(this, dishy.id)}>
        <div className="card card-hover" style={{width: "100%"}}>
          <div className="card-body">
            <div className="row">
              <div className="col-5" style={{textAlign: "center"}}>
                <h2 className="card-title">{dishy.name}</h2>
                <p className="card-text">No Data Received</p>
              </div>
              <div className="col-6" />
              <div className="col-1" style={{textAlign: "center"}}>
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
