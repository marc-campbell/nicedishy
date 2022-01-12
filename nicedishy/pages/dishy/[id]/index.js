import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router'
import Layout from "../../../components/layout";
import { Utilities } from '../../../utils/utilities';
import {XYPlot, LineSeries, HorizontalGridLines, XAxis, YAxis} from 'react-vis'
import dayjs from "dayjs";

export default function Page() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(true);
  const [dishy, setDishy] = useState(null);
  const [stats, setStats] = useState({});
  const [speeds, setSpeeds] = useState({});
  const [versions, setVersions] = useState({});

  const { id } = router.query

  useEffect( async () => {
    // hmm wait for the id
    // generate a nonce to use for the event source connection
    const nonce = await Utilities.fetchNonce();
    const source = new EventSource(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/dishy/stream?id=${id}&nonce=${nonce}`);
    source.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setDishy(data.dishy);
      setStats(data.stats);
      setSpeeds(data.speeds);
      setVersions(data.versions);

      if (isLoading) {
        setIsLoading(false);
      }
    }
  }, [])

  if (isLoading) {
    return <div />
  }

  let downloadSpeedData = [];
  for (const [when, speed] of Object.entries(speeds)) {
    const x = new Date(when);
    downloadSpeedData.push({
      x: x,
      y: speed.downloadSpeed,
    });
  }

  let uploadSpeedData = [];
  for (const [when, speed] of Object.entries(speeds)) {
    const x = new Date(when);
    uploadSpeedData.push({
      x: x,
      y: speed.uploadSpeed,
    });
  }

  return (
    <div>
      <h1>{dishy.name}</h1>

      <h3>Versions</h3>
      <div className="container">
        <div className="row">
          <div className="col-4">
            <h4>Current Software Revision{' '}<i className="bi bi-info-circle"></i></h4>
          </div>
          <div className="col-8">
            {versions.software}
          </div>
        </div>
        <div className="row">
          <div className="col-4">
            <h4>Current Hardware Revision{' '}<i className="bi bi-info-circle"></i></h4>
          </div>
          <div className="col-8">
            {versions.hardware}
          </div>
        </div>
      </div>

      <h3>Speed</h3>
      <div className="container">
        <div className="row">
          <div className="col-12">
            <h4>Download Speed</h4>
            <XYPlot height={200} width={900}>
              <HorizontalGridLines />
              <LineSeries data={downloadSpeedData} />
              <XAxis tickTotal={4} tickFormat={(t, i) => {
                var d = new Date(0);
                d.setUTCMilliseconds(t);
                return `${dayjs(d).format("MM-DD-YYYY")} @ ${dayjs(d).format("HH:mm:ss a")}`;
              }}/>
              <YAxis tickTotal={4} width={70} tickFormat={(t, i) => {
                return Utilities.mbps(t, 10);
              }}/>
            </XYPlot>
          </div>
        </div>

        <div className="row">
          <div className="col-12">
            <h4>Upload Speed</h4>
            <XYPlot height={200} width={900}>
              <HorizontalGridLines />
              <LineSeries data={uploadSpeedData} />
              <XAxis tickTotal={4} tickFormat={(t, i) => {
                var d = new Date(0);
                d.setUTCMilliseconds(t);
                return `${dayjs(d).format("MM-DD-YYYY")} @ ${dayjs(d).format("HH:mm:ss a")}`;
              }}/>
              <YAxis tickTotal={4} width={70} tickFormat={(t, i) => {
                return Utilities.mbps(t, 10);
              }}/>
            </XYPlot>
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

// This is here to disable https://nextjs.org/docs/advanced-features/automatic-static-optimization
Page.getInitialProps = async (ctx) => {
  return {};
}
