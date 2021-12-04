import React, { useEffect, useState } from 'react';
import { Utilities } from "../../utils/utilities";
import { useRouter } from 'next/router'

function Dishies() {
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

  return (
    <div>
      <h1>Dishy (or Dishies if you are lucky)</h1>
      {
        dishies.map((dishy) => {
          return (
            <div key={dishy.id}>
              <h3>{dishy.name}</h3>
            </div>
          );
        })
      }
      <div>
        <h3>Add a new dishy</h3>
      </div>
    </div>
  );
}

export default Dishies;
