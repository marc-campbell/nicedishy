import * as React from "react";
import GaugeChart from 'react-gauge-chart';
import { Utilities } from "../utils/utilities";

export default function DishyPlaceholderCard() {
  return (
    <div className="card" style={{width: "100%", height: "300px"}}>
      <div className="card-body">
        <div className="row">
          <div className="col-6" style={{textAlign: "center"}}>
            <h2 className="card-title loading-placeholder-transparent" style={{textAlign: "left"}}>Dishy McFlatFace</h2>
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
                ({Utilities.mbps(400, 10)})
              </div>
              <div className="col-4 loading-placeholder">
                download speed<br />
                ({Utilities.mbps(400, 10)})
              </div>
            </div>
          </div>

          <div className="col-5 offset-1">
            <div className="row" style={{paddingTop: "60px"}}>
              <div className="col-12">
                <h4 className="loading-placeholder">More</h4>
              </div>
            </div>
            <div className="row">
              <div className="col-12 loading-placeholder">
                <i className="bi bi-clock-history"></i>{' '}Dashoard<br />
                <i className="bi bi-people-fill"></i>{' '}How do I compare?<br />
                <i className="bi bi-lightbulb"></i>{' '}Troubleshooting<br />
                <i className="bi bi-gear"></i>{' '}Settings
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
