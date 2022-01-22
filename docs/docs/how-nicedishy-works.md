# How NiceDishy works
  
NiceDishy works with a little helper agent that you run on a computer somewhere on your Starlink network. This service doesn't require the Starlink router, so even if you have your own router, our agent is going to work.

The agent serves as a bridge between your device and our service. On a regular basis, the agent will run various queries to retrieve the status of your dish or run an occaisional speed test in the background. We've designed these to be silent and out of the way and should never interfere with normal operations of your computer, network, or Starlink service. 

## Stats

Your Starlink dish (not the router), has a [GRPC](https://grpc.io) server that serves some basic stats. Our agent will query these GRPC stats regularly and send some of them to the NiceDishy service. These are anomymous stats and never can or will include any information about what you are using your internet connection for. NiceDishy just queries for some aggregate numbers and relay them to our API so that it can write them into a time series database and build our metrics. Querying stats locally is a very low effort operation and doesn't use any data on your internet conenction. By default, it queries these stats every minute, and this is configurable in the preferences/settings page of the agent.

If you are curious to know exactly what data is being sent with the stats, the mock JSON object below contains the fields that are sent along with a description of each:

```json
{
  "uptimeSeconds": number of seconds since the last reboot,
  "hardwareVersion": the version of the hardware, as reported by the dish,
  "softwareVersion": the firware version running on the device,
  "snr": the signal-to-noise ratio,
  "downlinkThroughputBps": the actual download throughput used (not speed, but how much used),
  "uplinkThroughputBps": the actual upload throughput used (not speed, but how much used),
  "popPingLatencyMs": the ping latency rate,
  "popPingDropRate": the drop rate of pings,
  "percentObstructed": the percent of time your dish is obstructed,
  "secondsObstructed": the number of seconds your dish has been obstructed,
}
```

## Speed

Your Starlink device doesn't have the ability to report on available download and upload speed. But regular speed tests is important to know if the service and installation are improving over time. 

We've implemented as regular speed test into the agent. By default, it will run a download and upload speed test to fast.com every hour. This time interval is configurable in the preferences/settings page of the agent application.

If you are curious to know exactly what data is being sent with the speed tests, the mock JSON object below contains the fields that are sebt along with a description of each:

```json
{
  "downloadSpeed": the download speed from fast.com, in bps,
  "uploadSpeed": the upload speed to fast.com, in bps,
}
```

!!! note
    Unfortunately, fast.com doesn't provide an easy API to run speed tests. To solve this, we've implemented our own logic to perform speed tests using fast.com's servers and 
    attempting to mimic their logic. Sometimes our results are different from the ones you'll see on fast.com directly. I are continuing to tune the code here
    and hope to get the same results as more people useh NiceDishy.
