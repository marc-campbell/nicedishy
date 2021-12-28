import React, { useEffect, useState } from 'react';
import { Utilities, secondsAgo } from "../../utils/utilities";
import { useRouter } from 'next/router'
import Layout from "../../components/layout";
import {XYPlot, LineSeries} from 'react-vis'
import GaugeChart from 'react-gauge-chart';

export default function Page() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [dishies, setDishies] = useState([]);

  const handleAddAnotherClick = () => {
    router.push(`/dishy/new`);
  }

  const handleSettingsClick = (dishyId) => {
    router.push(`/dishy/${dishyId}/settings`);
  }

  const fetchDishies = async() => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/dishies`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
      });

      if (res.status === 401) {
        router.push('/login?next=/dishies');
        return;
      }

      if (!res.ok) {
        return;
      }

      const data = await res.json();
      return data;
    } catch (err) {
      console.error(err);
    }
  }

  const fetchNonce = async() => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/nonce`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
      });

      if (res.status === 401) {
        router.push('/login?next=/dishies');
        return;
      }

      if (!res.ok) {
        return;
      }

      const data = await res.json();
      return data.nonce;
    } catch (err) {
      console.error(err);
    }
  }

  useEffect( async () => {
    // generate a nonce to use for the event source connection
    const nonce = await fetchNonce();
    const source = new EventSource(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/dishies/stream?nonce=${nonce}`);
    source.onmessage = (event) => {
      console.log(event);
    }
  }, [])

  useEffect( async () => {
    const data = await fetchDishies();
    if (!data) {
      return;
    }

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
    let isConnected = false;
    let hasEverConnected = !!dishy.lastMetricAt;
    if (hasEverConnected) {
      isConnected = secondsAgo(dishy.lastMetricAt) < 86400;
    }

    if (!hasEverConnected) {
      return (
        <div key={dishy.id} className="card" style={{width: "100%"}}>
          <div className="card-body">
            <div className="row">
              <div className="col-5" style={{textAlign: "center"}}>
                <h2 className="card-title">{dishy.name}</h2>
                Never Connected
              </div>
              <div className="col-3" />
              <div className="col-4" style={{textAlign: "center"}}>
                <p>
                  To connect your Dishy, download and install our app on a laptop or workstation that’s connected to the Dishy network.
                </p>
                <a href="#" className="btn btn-primary" onClick={handleDownloadClick}>Download</a>
                {' '}
                <a href="#" className="btn btn-outline-primary">Docs</a>
                <br /><br />
                <a href="#" onClick={handleSettingsClick.bind(this, dishy.id)}>
                  Settings <i className="bi bi-gear"></i>
                </a>
              </div>
            </div>
          </div>
        </div>
      )
    }

    if (!isConnected) {
      return (
        <div key={dishy.id} className="card" style={{width: "100%"}}>
          <div className="card-body">
            <div className="row">
              <div className="col-5" style={{textAlign: "center"}}>
                <h2 className="card-title">{dishy.name}</h2>
                Disconnected
              </div>
              <div className="col-3" />
              <div className="col-4" style={{textAlign: "center"}}>
                <p>
                  To connect your Dishy, download and install our app on a laptop or workstation that’s connected to the Dishy network.
                </p>
                <a href="#" className="btn btn-primary" onClick={handleDownloadClick}>Download</a>
                {' '}
                <a href="#" className="btn btn-outline-primary">Docs</a>
                <br /><br />
                <a href="#" onClick={handleSettingsClick.bind(this, dishy.id)}>
                  Settings <i className="bi bi-gear"></i>
                </a>
              </div>
            </div>
          </div>
        </div>
      )
    }

    return (
      <div key={dishy.id} className="card" style={{width: "100%"}}>
        <div className="card-body">
          <div className="row">
            <div className="col-6" style={{textAlign: "center"}}>
              <h2 className="card-title" style={{textAlign: "left"}}>{dishy.name}</h2>
              <div className="row" style={{paddingTop: "30px"}}>
                <div className="col-4">
                  <div style={{fontSize: "3em", lineHeight: "1.2em", paddingTop: "10px"}}>{dishy.latest.popPingLatencyMs.toFixed(0)}ms</div>
                </div>
                <div className="col-4">
                  <GaugeChart
                    percent={dishy.latest.uploadSpeed / 40000000}
                    colors={["#FF5F6D", "#00ff00"]}
                    hideText={true}
                  />
                </div>
                <div className="col-4">
                  <GaugeChart
                    percent={dishy.latest.downloadSpeed / 300000000}
                    colors={["#FF5F6D", "#00ff00"]}
                    hideText={true}
                  />
                </div>
              </div>
              <div className="row" style={{paddingTop: "30px"}}>
                <div className="col-4">
                  ping
                </div>
                <div className="col-4">
                  upload speed<br />
                  ({Utilities.mbps(dishy.latest.uploadSpeed, 10)})
                </div>
                <div className="col-4">
                  download speed<br />
                  ({Utilities.mbps(dishy.latest.downloadSpeed, 10)})
                </div>
              </div>
            </div>

            <div className="col-5 offset-1">
              <div className="row" style={{paddingTop: "60px"}}>
                <div className="col-12">
                  <h4>More</h4>
                </div>
              </div>
              <div className="row">
                <div className="col-12">
                  <a href={`/dishy/${dishy.id}`}><i className="bi bi-clock-history"></i>{' '}Recent stats from my dishy</a><br />
                  <a href={`/dishy/${dishy.id}/compare`}><i className="bi bi-people-fill"></i>{' '}How do I compare?</a><br />
                  <a href={`/dishy/${dishy.id}/troubleshooting`}><i className="bi bi-lightbulb"></i>{' '}Troubleshooting</a><br />
                  <a href={`/dishy/${dishy.id}/settings`}><i className="bi bi-gear"></i>{' '}Settings</a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    )
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
