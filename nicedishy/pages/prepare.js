import React from 'react';
import Layout from "../components/layout";

export default function Page() {

  return (
    <div>
      <h1>Prepare for your Starlink</h1>
      <p>
        Are you getting ready to put a Starlink dish somewhere on your house or yard?
        What do you need to know before it&apos;s time?
      </p>
      <h5>Order timing & phases</h5>
      <p>
        When you preorder (pay $99 USD or equivilent), you are just reserving your spot.
        You&apos;ll get an email saying that Starlink is preparing to ship your Dish.
        If you don&apos;t respond to the email, it&apos;s ok.
        The dish will still send, assuming you have a valid payment method on file.
      </p>
      <h5>Finding a location</h5>
      <p>
        Probably the hardest part of Starlink for must of us is finding a good place to install it.
        Luckily, the service has improved to be a little more resilient to bad locations recently.
      </p>
      <h5>Accessories?</h5>
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


