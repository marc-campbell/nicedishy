import Link from 'next/link'
import { useState } from 'react';
import Modal from 'react-modal';

Modal.setAppElement('#__next');

export default function Footer() {
  const [showSubscribe, setShowSubscribe] = useState(false);

  const onSubscribeClick = () => {
    setShowSubscribe(true);
  }

  const handleCloseNewsletterModal = () => {
    setShowSubscribe(false);
  }

  return (
    <>
      <footer className="bg-purple text-center py-5">
        <div className="container px-1">
          <div className="text-black-50 small">
            <div className="row">

              <div className="col-2">
                <h5 className="header-light">Dishy Owners</h5>
                <ul className="nav flex-column">
                  <li className="nav-item mb-2 nav-item-light"><Link href="/signup" className="nav-link p-0 text-muted">Join NiceDishy</Link></li>
                  {/* <li className="nav-item mb-2 nav-item-light"><Link href="/compare"><a className="nav-link p-0 text-muted">Compare My Dishy</a></Link></li>
                  <li className="nav-item mb-2 nav-item-light"><Link href="/outages"><a className="nav-link p-0 text-muted">Is Everything Ok?</a></Link></li> */}
                </ul>
              </div>

              <div className="col-2">
                <h5 className="header-light">Dishy Waitlisters</h5>
                <ul className="nav flex-column">
                  <li className="nav-item mb-2 nav-item-light"><Link href="https://docs.nicedishy.com/prepare"><a className="nav-link p-0 text-muted">Get Prepared</a></Link></li>
                </ul>
              </div>

              <div className="col-2">
                <h5 className="header-light">Everyone Else</h5>
                <ul className="nav flex-column">
                  {/* <li className="nav-item mb-2 nav-item-light"><Link href={`${process.env.NEXT_PUBLIC_DASHBOARD_ENDPOINT}/d/${process.env.NEXT_PUBLIC_GRAFANA_PUBLIC_FOLDER}/the-data`}><a className="nav-link p-0 text-muted">See Some Stats</a></Link></li> */}
                </ul>
              </div>

              <div className="col-4 offset-1">
                <h5 className="header-light">Subscribe to our newsletter</h5>
                <div className="btn btn-outline-light" onClick={onSubscribeClick}>Subscribe</div>
              </div>
            </div>

          </div>
        </div>
      </footer>
      <Modal
        isOpen={showSubscribe}
        onRequestClose={handleCloseNewsletterModal}
        className="newsletter-modal"
        contentLabel="Subscribe to the NiceDishy newsletter"
        style={{
          overlay: {
            position: 'fixed',
            top: 0,
            left: 0,
            right: 0,
            bottom: 0,
            backgroundColor: 'rgba(0, 0, 0, 0.75)'
          }
        }}>
        <iframe src="https://nicedishy.substack.com/embed" width="480" height="320" style={{border: "1px solid #EEE", background: "white"}} frameBorder="0" scrolling="no"></iframe>
      </Modal>
    </>
  );
}
