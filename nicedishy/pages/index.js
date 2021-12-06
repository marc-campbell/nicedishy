import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import Link from 'next/link'
import { useRouter } from 'next/router'
import Footer from '../components/footer';

export default function Home() {
  const router = useRouter();

  const handleCompareClick = (e) => {
    e.preventDefault();
    router.push("/stats");
  }

  const handleMonitorClick = (e) => {
    e.preventDefault();
    router.push("/monitor");
  }

  const handleLearnClick = (e) => {
    e.preventDefault();
    router.push("/learn");
  }

  return (
    <div>
      <header id="header" className="navbar navbar-expand-lg navbar-end navbar-absolute-top navbar-light navbar-show-hide"
        data-hs-header-options='{
          "fixMoment": 1000,
          "fixEffect": "slide"
        }'>
        <div className="container navbar-topbar">
          <nav className="js-mega-menu navbar-nav-wrap">
            <button className="navbar-toggler ms-auto" type="button" data-bs-toggle="collapse" data-bs-target="#topbarNavDropdown" aria-controls="topbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
              <span className="d-flex justify-content-between align-items-center small">
                <span className="navbar-toggler-text">Topbar</span>

                <span className="navbar-toggler-default">
                  <i className="bi-chevron-down ms-2"></i>
                </span>
                <span className="navbar-toggler-toggled">
                  <i className="bi-chevron-up ms-2"></i>
                </span>
              </span>
            </button>

            <div id="topbarNavDropdown" className="navbar-nav-wrap-collapse collapse navbar-collapse navbar-topbar-collapse">
              <div className="navbar-toggler-wrapper">
                <div className="navbar-topbar-toggler d-flex justify-content-between align-items-center">
                  <span className="navbar-toggler-text small">Topbar</span>

                  <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#topbarNavDropdown" aria-controls="topbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
                    <i className="bi-x"></i>
                  </button>
                </div>
              </div>

            </div>
          </nav>
        </div>

        <div className="container">
          <nav className="js-mega-menu navbar-nav-wrap">
            <a className="navbar-brand" href="./index.html" aria-label="Front">
              <img className="navbar-brand-logo" src="/images/logo.png" alt="NiceDishy" />
            </a>

            <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
              <span className="navbar-toggler-default">
                <i className="bi-list"></i>
              </span>
              <span className="navbar-toggler-toggled">
                <i className="bi-x"></i>
              </span>
            </button>

            <div className="collapse navbar-collapse" id="navbarNavDropdown">
              <div className="navbar-absolute-top-scroller">
                <ul className="navbar-nav">

                <li className="nav-item">
                  <Link href="/docs"><div className="btn btn-transition">Docs</div></Link>
                </li>

                <li className="nav-item">
                  <Link href="/login"><div className="btn btn-transition btn-outline-primary">Sign Up</div></Link>
                </li>
              </ul>
            </div>
          </div>
        </nav>
      </div>
    </header>


      <main id="content" role="main">
        <div className="d-lg-flex position-relative">
          <div className="container d-lg-flex align-items-lg-center content-space-t-3 content-space-lg-0 min-vh-lg-100">
            <div className="w-100">
              <div className="row">
                <div className="col-lg-5">
                  <div className="mb-5">
                    <h1 className="display-4 mb-3">
                      Nice Dishy
                    </h1>

                    <p className="lead">Continuous monitoring for your Starlink Dishy, from anywhere.
  <br /><br />
  Compare the performance of your Dishy with everyone else!</p>
                  </div>

                  <div className="d-grid d-sm-flex gap-3">
                    <Link href="/login"><div className="btn btn-primary btn-transition px-6">Get started</div></Link>
                    <Link href="/stats"><div className="btn btn-outline-secondary btn-transition px-6">See the stats</div></Link>
                  </div>
                </div>
              </div>
            </div>

            <div className="col-lg-7 col-xl-6 d-none d-lg-block position-absolute top-0 end-0 pe-0" style={{"marginTop": "6.75rem"}}>

            </div>
          </div>
        </div>

        <div className="container content-space-2 content-space-t-xl-3 content-space-b-lg-3">
          <div className="w-md-150 w-lg-100 text-center mx-md-auto mb-5">
            <h2>Why join the amazing NiceDishy network?</h2>
          </div>

          <div className="row mb-5 mb-md-0">
            <div className="col-sm-6 col-lg-4 mb-4 mb-lg-0">
              <div className="card card-sm h-100">
                <div className="p-2" style={{textAlign: "center", marginTop: "10px"}}>
                  <img className="card-img" style={{width: "60px"}} src="/images/performance.png" alt="Compare Performance"/ >
                </div>

                <div className="card-body">
                  <h4 className="card-title">Compare</h4>
                  <p className="card-text">Compare the performance of your Dishy to folks around you or anywhere in the world.</p>
                  <p className="card-text">Compare your download, upload, latency, reliability and more to others near you or far away from you.</p>
                </div>

                <a className="card-footer card-link border-top" href="#" onClick={handleCompareClick}>Learn more <i className="bi-chevron-right small ms-1"></i></a>
              </div>
            </div>

            <div className="col-sm-6 col-lg-4 mb-4 mb-lg-0">
              <div className="card card-sm h-100">
                <div className="p-2" style={{textAlign: "center", marginTop: "10px"}}>
                  <img className="card-img" style={{width: "60px"}} src="/images/performance.png" alt="Compare Performance"/ >
                </div>

                <div className="card-body">
                  <h4 className="card-title">Monitor</h4>
                  <p className="card-text">Check on the status of your Dishy from anywhere.</p>

                </div>

                <a className="card-footer card-link border-top" href="#" onClick={handleMonitorClick}>Learn more <i className="bi-chevron-right small ms-1"></i></a>
              </div>
            </div>

            <div className="col-sm-6 col-lg-4">
              <div className="card card-sm h-100">
                <div className="p-2" style={{textAlign: "center", marginTop: "10px"}}>
                  <img className="card-img" style={{width: "60px"}} src="/images/performance.png" alt="Compare Performance"/ >
                </div>

                <div className="card-body">
                  <h4 className="card-title">Learn</h4>
                  <p className="card-text">Learn what to expect and understand if Starlink is working as expected.</p>

                </div>

                <a className="card-footer card-link border-top" href="#" onClick={handleLearnClick}>Learn more <i className="bi-chevron-right small ms-1"></i></a>
              </div>
            </div>
          </div>
        </div>

        <div className="position-relative bg-light rounded-2 mx-3 mx-lg-10">
          <div className="container content-space-2 content-space-lg-3">
            <div className="w-md-75 w-lg-50 text-center mx-md-auto mb-5">
              <h2>Global Stats</h2>
              <p>A quick look at the stats powering NiceDishy</p>
            </div>

            <div className="row">
              <div className="col-lg-12 mb-12 mb-lg-12">
                <div className="pe-lg-12">
                  <figure className="device-browser">
                    <div className="device-browser-header">
                      <div className="device-browser-header-btn-list">
                        <span className="device-browser-header-btn-list-btn"></span>
                        <span className="device-browser-header-btn-list-btn"></span>
                        <span className="device-browser-header-btn-list-btn"></span>
                      </div>
                      <div className="device-browser-header-browser-bar">nicedishy.com/stats/</div>
                    </div>

                    <div className="device-browser-frame">
                      <img className="device-browser-img" src="/images/stats-screen.png" alt="Stats" />
                    </div>
                  </figure>
                </div>
              </div>

            </div>
          </div>
        </div>

      </main>
      <Footer />
    </div>
  )
}
