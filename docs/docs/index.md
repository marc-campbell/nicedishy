# What is NiceDishy

NiceDishy is a service that can receive metrics and data rom your [Starlink](https://starlink.com) dish to help you monitor and manage your dish.

!!! Beta
    The service is currently in private beta. There might be some bumps and issues.
    
## Features

NiceDishy is a *free* hosted service and installable agent (see "[How NiceDishy Works](/how-nicedishy-workds)") that ships your Starlink performance data
to a remote server. Using this data, the service can provide a lot more insights into how your dish is performing, what changes have been rolled out recently, 
and just help you stay better connected to what's happening with the technology you have installed.

### Monitoring Stats

The service collects anonymized stats from your dish that never can include any personally indentifying information such as browsing history or chracteristics of your local network.
The service collects what's available on the Starlink app's "Stats" page and send that to our servers.

### Monitoring Speed

The service also runs periodic fast.com speed tests on your network and send the download and upload results to our servers.

### The Dataset

To start, I make all of your data available to you. All of it. I don't keep anything secret or hold anything back. This is valuable right here, you can 
build a history of your Starlink device performance.

But NiceDishy goes a little further and triggers events such as email notifications when certain things happen. For example, when you get a new firmware version, 
we'll email you to let you know. 

The real power comes from combining these features into the community dataset. When you get a new firmware, unless you are the first to receive it, we'll already 
have some basic stastics about the performance characteristics of the new firmware. 

Or when your dish is slowing down and you are wondering if you need to climb on the roof or stow/unstow, you can also take a look at nearby dishies to see if 
anyone else is seeing problems. Sort of a "is it down for everyone or just me" for Starlink.
