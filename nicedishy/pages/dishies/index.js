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


    const latencyData = [
      {x: 0, y: 8},
      {x: 1, y: 5},
      {x: 2, y: 4},
      {x: 3, y: 9},
      {x: 4, y: 1},
      {x: 5, y: 7},
      {x: 6, y: 6},
      {x: 7, y: 3},
      {x: 8, y: 2},
      {x: 9, y: 0}
    ];

    const uploadData = [
      {x: 0, y: 8},
      {x: 1, y: 5},
      {x: 2, y: 4},
      {x: 3, y: 9},
      {x: 4, y: 1},
      {x: 5, y: 7},
      {x: 6, y: 6},
      {x: 7, y: 3},
      {x: 8, y: 2},
      {x: 9, y: 0}
    ];

    const downloadData = [
      {x: 0, y: 9},
      {x: 1, y: 2},
      {x: 2, y: 5},
      {x: 3, y: 1},
      {x: 4, y: 8},
      {x: 5, y: 3},
      {x: 6, y: 4},
      {x: 7, y: 6},
      {x: 8, y: 1},
      {x: 9, y: 1}
    ];

    return (
      <div key={dishy.id} className="card" style={{width: "100%"}}>
        <div className="card-body">
          <div className="row">
            <div className="col-6" style={{textAlign: "center"}}>
              <h2 className="card-title">{dishy.name}</h2>
              <div className="row">
                <div className="col-4">
                  <span style={{fontSize: "4em"}}>24ms</span><br /><br />
                  ping
                </div>
                <div className="col-4">
                  <GaugeChart
                    percent={0.45}
                    colors={["#FF5F6D", "#00ff00"]}
                    hideText={true}
                  />
                  uplink
                </div>
                <div className="col-4">
                <GaugeChart
                    percent={0.9}
                    colors={["#FF5F6D", "#00ff00"]}
                    hideText={true}
                  />downlink
                </div>
              </div>
            </div>

            <div className="col-3" style={{textAlign: "center", paddingTop: "60px"}}>
              <XYPlot height={100} width={400}>
                <LineSeries data={latencyData} />
              </XYPlot><br />
              Latency
            </div>
            <div className="col-3" style={{textAlign: "center", paddingTop: "60px"}}>
              <XYPlot height={100} width={400}>
                <LineSeries data={downloadData} />
                <LineSeries data={uploadData} />
              </XYPlot><br />
              Upload / Download
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
