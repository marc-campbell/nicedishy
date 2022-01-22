import Link from 'next/link'
import { useState } from 'react';

export default function Footer() {
  const [emailAddress, setEmailAddress] = useState('');

  const handleSubscribeClick = async () => {
    const uri = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/subscribe`;
    const response = await fetch(uri, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        emailAddress,
      }),
    });

    if (!response.ok) {
      return;
    }
  }

  return (
    <footer className="bg-light text-center py-5">
        <div className="container px-1">
            <div className="text-black-50 small">
                <div className="row">

                  <div className="col-2">
                    <h5>Dishy Owners</h5>
                    <ul className="nav flex-column">
                      <li className="nav-item mb-2"><a href="#" className="nav-link p-0 text-muted">Join NiceDishy</a></li>
                      <li className="nav-item mb-2"><Link href="/compare"><a className="nav-link p-0 text-muted">Compare My Dishy</a></Link></li>
                      <li className="nav-item mb-2"><a href="#" className="nav-link p-0 text-muted">Is Everything Ok?</a></li>
                    </ul>
                  </div>

                  <div className="col-2">
                    <h5>Dishy Waitlisters</h5>
                    <ul className="nav flex-column">
                      <li className="nav-item mb-2"><Link href="https://docs.nicedishy.com/prepare"><a className="nav-link p-0 text-muted">Get Prepared</a></Link></li>
                    </ul>
                  </div>

                  <div className="col-2">
                    <h5>Everyone Else</h5>
                    <ul className="nav flex-column">
                      <li className="nav-item mb-2"><Link href="/stats"><a className="nav-link p-0 text-muted">See Some Stats</a></Link></li>
                      <li className="nav-item mb-2"><a href="#" className="nav-link p-0 text-muted">Request Data</a></li>
                    </ul>
                  </div>

                  <div className="col-4 offset-1">
                    <form>
                      <h5>Subscribe to our newsletter</h5>
                      <p>A semi-regular newsletter about Dishy.</p>
                      <div className="d-flex w-100 gap-2">
                        <label htmlFor="newsletter1" className="visually-hidden">Email address</label>
                        <input value={emailAddress} onChange={e => setEmailAddress(e.target.value)} type="text" className="form-control" placeholder="Email address" />
                        <button className="btn btn-outline-secondary" type="button" onClick={handleSubscribeClick}>Subscribe</button>
                      </div>
                    </form>
                  </div>
                </div>

            </div>
        </div>
    </footer>
  );
}
