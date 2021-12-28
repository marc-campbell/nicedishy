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
  const [recent, setRecent] = useState({});

  const { id } = router.query

  useEffect( async () => {
    // hmm wait for the id
    // generate a nonce to use for the event source connection
    const nonce = await Utilities.fetchNonce();
    const source = new EventSource(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/dishy/stream?id=${id}&nonce=${nonce}`);
    source.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setDishy(data.dishy);
      setRecent(data.recent);

      if (isLoading) {
        setIsLoading(false);
      }
    }
  }, [])

  if (isLoading) {
    return <div />
  }

  let downloadSpeedData = [];
  for (const [when, stats] of Object.entries(recent)) {
    const x = new Date(when);
    downloadSpeedData.push({
      x: x,
      y: stats.downloadSpeed,
    });
  }

  return (
    <div>
      <h1>Dishy</h1>
      <h4>{dishy.name}</h4>
      <XYPlot height={200} width={900}>
        <HorizontalGridLines />
        <LineSeries data={downloadSpeedData} />
        <XAxis title="When" tickTotal={4} tickFormat={(t, i) => {
          var d = new Date(0);
          d.setUTCMilliseconds(t);
          return `${dayjs(d).format("MM-DD-YYYY")} @ ${dayjs(d).format("HH:mm:ss a")}`;
        }}/>
        <YAxis tickTotal={4} width={70} tickFormat={(t, i) => {
          console.log(t, 10);
          return Utilities.mbps(t, 10);
        }}/>
      </XYPlot><br />
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
