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
    // generate a nonce to use for the event source connection
    const nonce = await Utilities.fetchNonce(router, '/dishies');
    const source = new EventSource(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/dishies/stream?nonce=${nonce}`);
    source.onmessage = (event) => {
      setIsLoading(false);
      setDishies(JSON.parse(event.data));
    }
  }, [])


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
