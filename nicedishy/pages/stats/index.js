import React, { useEffect, useState } from 'react';
import Layout from "../../components/layout";
import { Utilities } from '../../utils/utilities';
import dayjs from "dayjs";

export default function Page() {
  const [stats, setStats] = useState({});
  const [wasConnected, setWasConnected] = useState(false);
  const [isConnected, setIsConnected] = useState(false);
  const [disconnctedAt, setDisconnectedAt] = useState("");

  useEffect( async() => {
    const source = new EventSource(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/stats/public/stream`);
    source.onmessage = (event) => {
      setStats(JSON.parse(event.data));
    }
    source.onopen = () => {
      setIsConnected(true);
      setWasConnected(true);
    }
    source.onerror = () => {
      setDisconnectedAt(new Date().toISOString());
      setIsConnected(false);
    }
  }, []);

  return (
    <div>
      <div className="row">
        <div className="col-8">
          <h1>Browse the stats</h1>
        </div>
        <div className="col-4">
          <p style={{textAlign: "right", paddingTop: "20px", marginRight: "40px"}}>
            {isConnected ?
              <span><i className="bi bi-circle-fill" style={{color: "#00aa00"}}></i>{" "}connected</span>
              :
              wasConnected ?
                <span><i className="bi bi-circle-fill" style={{color: "#aa0000"}}></i>{" "}disconnected</span>
                :
                <i className="bi bi-arrow-repeat" style={{color: "#aaaaaa"}}></i>
          }
          </p>
        </div>
      </div>
      <div className={`alert alert-danger ${wasConnected && !isConnected ? "" : "d-none"}`} role="alert">
        Disconnected from server at {dayjs(disconnctedAt).format("HH:mm:ss a on MMM DD, YYYY")}. These stats are no longer automatically updating.
      </div>
      <div className="container">
        <h5>Numbers</h5>
        <div className="row dashboard-row">
          <div className="col-md-4 px-md-5 text-center">
            <div className="dashboard-card dashboard-card-blue">
              <div className="big-number">{stats.connectedDishyCount}</div>
              dishies currently sending data<br />
              ({stats.allTimeDishyCount} seen all time)
            </div>
          </div>
          <div className="col-md-4 px-md-5 text-center">
          <div className="dashboard-card dashboard-card-blue">
              <div className="big-number">{stats.newDishyCount}</div>
              new dishies connected for the first time this week
            </div>
          </div>
          <div className="col-md-4 px-md-5 text-center">
            <div className="dashboard-card">
              <div className="big-number">?</div>
              somethings<br />
            </div>
          </div>
        </div>
      </div>

      <div className="container">
        <h5>Speed</h5>
        <div className="row dashboard-row">
          <div className="col-md-6 px-md-5 text-center">
            <div className="dashboard-card">
              <div className="big-number">{Utilities.mbps(stats.highestDownloadSpeed, 10)}</div>
              The highest download speed we&apos;ve seen and confirmed in the past week
            </div>
          </div>
          <div className="col-md-6 px-md-5 text-center">
            <div className="dashboard-card">
              <div className="big-number">{Utilities.mbps(stats.averageDownloadSpeed, 10)}</div>
              The average download speed in the past week
            </div>
          </div>
        </div>
        <div className="row dashboard-row">
          <div className="col-md-6 px-md-5 text-center">
            <div className="dashboard-card">
              <div className="big-number">{stats.lowestPingTime ? stats.lowestPingTime.toFixed(1) : ""} ms</div>
              The lowest (best) ping time we&apos;ve seen in the past week
            </div>
          </div>
          <div className="col-md-6 px-md-5 text-center">
            <div className="dashboard-card">
              <div className="big-number">{stats.averagePingTime ? stats.averagePingTime.toFixed(1) : ""} ms</div>
              The average ping time we&apos;ve seen in the past week
            </div>
          </div>
        </div>
      </div>

      <div className="container">
        <h5>Where</h5>
        <div className="row dashboard-row">
          <div className="col-md-4 px-md-5 text-center">
            <div className="dashboard-card">
              We have dishies from <em>nn</em> countries.<br />
            </div>
          </div>
          <div className="col-md-4 px-md-5 text-center">
            <div className="dashboard-card">
              The most popular country is <em>nn</em> with <em>nn</em> dishies.<br />
            </div>
          </div>
          <div className="col-md-4 px-md-5 text-center">
            <div className="dashboard-card">
              The fastest area recently seems to be <em>nn</em> with an average download speed of <em>nn</em>.
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
