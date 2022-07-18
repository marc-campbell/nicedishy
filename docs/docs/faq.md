# FAQ

#### How much does NiceDishy cost, and why is it free?

NiceDishy is free, and there I have no plans to charge for it. I've lived with a 6/1 Mbps DSL connection for a while and Starlink changed the game for me. Faster internet is simply not available, even though I'm just outside of Los Angeles. That said, Starlink is still early, and it's fun to watch it improve and grow. I've been noticing trends where my speeds get better (and sometimes slower). I built the original version of NiceDishy as a monitoring site for my device, but realized that adding some auth to it and a few changes would make it easy for others. 

I hope it's useful, but also (and very selfishly) having a larger set of data to compare my dish to is going to help identify trends or "is it [down, faster, slower, better, worse, different] for everyone or just me?"

#### The Starlink app has remote monitoring, why is NiceDishy better?

This is a new feature that's been introduced in the Starlink app. This feature only works if you have the Starlink router (round dish) or the rectangular dish. In addition, this doesn't provide alerts or speed tests from your device when you are remote. 

The value of NiceDishy is that these speed tests run regularly and allow you to see if your device is performing at the same levels as others.

#### Is there a Starlink outage?

One of the side effects of tracking this data for so many devices is that the service can start to detect when there are outages. Once a few more people use the site, I'm planning to add an outage map to show where service is degraded or down.

#### Will using this violate my ToS with Starlink?

I'm not reverse engineering or doing anything even remotely evil on your network or dish. I won't ship software or changes that put your service at risk in any way. The NiceDishy agent uses an eposed API on the dish -- the same one that the browser/app use. We've had zero reports of anyone having problems from SpaceX from using NiceDishy and will report here if this changes.
