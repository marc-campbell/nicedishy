import React, { useEffect, useState } from 'react';
import { Utilities } from "../../utils/utilities";
import { useRouter } from 'next/router'
import Layout from "../../components/layout";

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

  useEffect( async () => {
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

  const handleDownloadClick = (e) => {
    e.preventDefault();
    router.push('/download');
  }

  const cards = dishies.map((dishy) => {
    return (
      <div key={dishy.id} className="card" style={{width: "100%"}}>
        <div className="card-body">
          <div className="row">
            <div className="col-5" style={{textAlign: "center"}}>
              <h2 className="card-title">{dishy.name}</h2>
              <p className="card-text">Not connected</p>
            </div>
            <div className="col-3" />
            <div className="col-4" style={{textAlign: "center"}}>
              <p>
                To connect your Dishy, download and install our app on a laptop or workstation that’s connected to the Dishy network.
              </p>
              <a href="#" className="btn btn-primary" onClick={handleDownloadClick}>Download</a>
              {' '}
              <a href="#" className="btn btn-outline-primary">Docs</a>
            </div>
          </div>
        </div>
      </div>
    );
  });

  return (
    <>
      {cards}
      <br /><br />
      <div className="container">
        <div style={{textAlign: "center"}}>
          <a href="#" className="btn btn-outline-secondary">
            <i class="bi bi-plus-circle"></i>
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
