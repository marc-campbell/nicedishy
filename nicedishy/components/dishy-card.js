import * as React from "react";
import { useRouter } from 'next/router'
import GaugeChart from 'react-gauge-chart';
import { Utilities, secondsAgo } from "../utils/utilities";
import Link from 'next/link';

export default function DishyCard({dishy}) {
  const router = useRouter();

  const handleDownloadClick = (e) => {
    e.preventDefault();
    router.push('/download');
  }

  const handleSettingsClick = (dishyId) => {
    router.push(`/dishy/${dishyId}/settings`);
  }

  let isConnected = false;
  let hasEverConnected = !!dishy.lastMetricAt;
  if (hasEverConnected) {
    isConnected = secondsAgo(dishy.lastMetricAt) < 86400;
  }

  if (!hasEverConnected || !isConnected) {
    return (
      <div key={dishy.id} className="card" style={{width: "100%", height: "300px"}}>
        <div className="card-body">
          <div className="row">
            <div className="col-6" style={{textAlign: "center"}}>
              <h2 className="card-title" style={{textAlign: "left"}}>{dishy.name}</h2>
              <div className="row" style={{paddingTop: "30px"}}>
                <div className="col-4">
                  <div style={{fontSize: "3em", lineHeight: "1.2em", paddingTop: "10px"}} className="loading-placeholder">00 ms</div>
                </div>
                <div className="col-4">
                  <GaugeChart
                    percent={40}
                    needleColor="#aaaaaa"
                    needleBaseColor="#aaaaaa"
                    colors={["#aaaaaa", "#dddddd"]}
                    hideText={true}
                    animate={false}
                  />
                </div>
                <div className="col-4">
                  <GaugeChart
                    percent={60}
                    needleColor="#aaaaaa"
                    needleBaseColor="#aaaaaa"
                    colors={["#aaaaaa", "#dddddd"]}
                    hideText={true}
                    animate={false}
                  />
                </div>
              </div>
              <div className="row" style={{paddingTop: "30px"}}>
                <div className="col-4 loading-placeholder">
                  ping
                </div>
                <div className="col-4 loading-placeholder">
                  upload speed<br />
                  ({Utilities.mbps(0, 10)})
                </div>
                <div className="col-4 loading-placeholder">
                  download speed<br />
                  ({Utilities.mbps(0, 10)})
                </div>
              </div>
            </div>

            <div className="col-5 offset-1">
              <div className="row" style={{paddingTop: "60px"}}>
                <div className="col-12">
                  {!hasEverConnected ? (<h4>Get Connected</h4>) : (<h4>Reconnect Your Dish</h4>) }
                </div>
              </div>
              <div className="row">
                <div className="col-12">
                  {!hasEverConnected ?
                    (
                      <p>To connect your Dishy, download and install our app on a laptop or workstation that’s connected to the Dishy network.</p>
                    )
                    :
                    (
                      <p>Your device is no longer sending data to NiceDishy.com. Check the installation or reinstall using the links below.</p>
                    )}
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
        </div>
      </div>
    )
  }

  return (
    <div key={dishy.id} className="card" style={{width: "100%", height: "300px"}}>
      <div className="card-body">
        <div className="row">
          <div className="col-6" style={{textAlign: "center"}}>
            <h2 className="card-title" style={{textAlign: "left"}}>{dishy.name}</h2>
            <div className="row" style={{paddingTop: "30px"}}>
              <div className="col-4">
                <div style={{fontSize: "3em", lineHeight: "1.2em", paddingTop: "10px"}}>{dishy.latestStats.popPingLatencyMs.toFixed(0)}ms</div>
              </div>
              <div className="col-4">
                <GaugeChart
                  percent={dishy.latestSpeeds.uploadSpeed / 40000000}
                  colors={["#FF5F6D", "#00ff00"]}
                  hideText={true}
                />
              </div>
              <div className="col-4">
                <GaugeChart
                  percent={dishy.latestSpeeds.downloadSpeed / 300000000}
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
                ({Utilities.mbps(dishy.latestSpeeds.uploadSpeed, 10)})
              </div>
              <div className="col-4">
                download speed<br />
                ({Utilities.mbps(dishy.latestSpeeds.downloadSpeed, 10)})
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
                <Link href={`${process.env.NEXT_PUBLIC_DASHBOARD_ENDPOINT}/d/${dishy.id}`} passHref><a target="_blank"><i className="bi bi-clock-history"></i>{' '}Dashboard</a></Link><br />
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
}
