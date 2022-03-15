import Link from 'next/link'
import Footer from '../components/footer';
import Image from 'next/image';
import Head from 'next/head';

export default function Home() {

  return (
    <body id="page-top">
      <Head>
        <title>NiceDishy - The missing bits for living with Starlink</title>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </Head>
      <div style={{position: "relative", minHeight: "100vh"}}>
        <div style={{paddingBottom: "200px"}}>
          <nav className="navbar navbar-expand-lg navbar-dark fixed-top shadow-sm bg-purple" id="mainNav">
            <div className="container px-5">
              <a className="navbar-brand fw-bold" href="#page-top">
                <Image className="navbar-brand-logo" src="/images/nicedishy-logo-white.png" height="48px" width="48px" alt="NiceDishy" />
              </a>
              <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
                Menu
                <i className="bi-list"></i>
              </button>
              <div className="collapse navbar-collapse" id="navbarResponsive">
                <ul className="navbar-nav ms-auto me-4 my-3 my-lg-0">
                  <li className="nav-item-light"><Link href="https://docs.nicedishy.com"><a className="nav-link me-lg-3">Docs</a></Link></li>
                  <li className="nav-item-light"><Link href="/login"><a className="nav-link me-lg-3">Login</a></Link></li>
                </ul>
              </div>
            </div>
          </nav>

          <header className="masthead">
            <div className="container px-5">
              <div className="row gx-5 align-items-center">
                <div className="col-lg-6">
                  <div className="mb-5 mb-lg-0 text-center text-lg-start">
                    <h1 className="display-1 lh-1 mb-3">
                      NiceDishy
                    </h1>
                    <p className="lead fw-normal text-muted mb-5">Continuous monitoring for your Starlink Dishy (Flatface or Squarepants), from anywhere.</p>
                    <p className="lead fw-normal text-muted mb-5">Compare the performance of your Dishy with everyone else and even get some alerts when something changes!</p>
                  </div>
                </div>
                <div className="col-lg-6">
                  <div className="masthead-device-mockup">
                    <Image src="/images/screenshot-small.png" maxwidth="550px" height="150px" width="400px" />
                  </div>
                  <div className="row mt-4">
                    <div className="col-12 text-center">
                      <Link href="/login">
                        <a>
                          <Image className="app-badge" src="/images/btn_google_signin_dark_pressed_web@2x.png" height="50px" width="200px" alt="Get Started" />
                        </a>
                      </Link>
                    </div>
                  </div>
                  <div className="row">
                    <div className="col-12 text-center">
                      Don’t have an account? <br />
                      <Link href="/signup">Click here to get started</Link>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </header>

          <aside className="text-center bg-gradient-primary-to-secondary">
              <div className="container px-5">
                  <div className="row gx-5 justify-content-center">
                      <div className="col-xl-12">
                          <div className="h2 fs-1 text-white mb-4">Woo. Think of it like the missing bits built by a Starlink fan.</div>
                      </div>
                  </div>
              </div>
          </aside>

          <section id="features">
              <div className="container px-5">
                  <div className="row gx-5 align-items-center">
                      <div className="col-lg-12 order-lg-1 mb-5 mb-lg-0">
                          <div className="container-fluid px-5">
                              <div className="row gx-5">
                                  <div className="col-md-6 mb-5">
                                    <div className="row">
                                      <div className="col-3">
                                        <Image src="/images/7575373_news_events_mars_emoji.png" height="120px" width="120px" />
                                      </div>
                                      <div className="col-9">
                                        <h3 className="font-alt">Compare</h3>
                                        <p className="text-muted mb-0">Compare the performance of your Dishy to folks around you or anywhere in the world</p>
                                      </div>
                                    </div>
                                  </div>
                                  <div className="col-md-6 mb-5">
                                    <div className="row">
                                      <div className="col-3">
                                        <Image src="/images/7734227_rotation_orbit_planet_mars.png" height="120px" width="120px" />
                                      </div>
                                      <div className="col-9">
                                        <h3 className="font-alt">Monitor</h3>
                                        <p className="text-muted mb-0">Check on the status of your Dishy from anywhere. Even when not home and not using the router.</p>
                                      </div>
                                    </div>
                                  </div>
                              </div>
                              <div className="row gx-5">
                                  <div className="col-md-6 mb-5">
                                    <div className="row">
                                      <div className="col-3">
                                        <Image src="/images/7582125_research_explore_planet_exploration_mars_science.png" height="120px" width="120px" />
                                      </div>
                                      <div className="col-9">
                                        <h3 className="font-alt">Learn</h3>
                                        <p className="text-muted mb-0">Learn what to expect and understand if Starlink is working as expected.</p>
                                      </div>
                                    </div>
                                  </div>
                                  <div className="col-md-6 mb-5">
                                    <div className="row">
                                      <div className="col-3">
                                        <Image src="/images/7575375_emoji_image_mars_multimedia.png" height="120px" width="120px" />
                                      </div>
                                      <div className="col-9">
                                        <h3 className="font-alt">Discover</h3>
                                        <p className="text-muted mb-0">Get an email when something changes (firmware, speed).</p>
                                      </div>
                                    </div>
                                  </div>
                              </div>
                          </div>
                      </div>
                  </div>
              </div>
          </section>
        </div>
        <Footer />
      </div>

    </body>
  )
}
