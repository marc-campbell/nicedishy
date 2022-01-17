import Link from 'next/link'
import Footer from '../components/footer';

export default function Home() {

  return (
    <body id="page-top">
        <nav className="navbar navbar-expand-lg navbar-light fixed-top shadow-sm" id="mainNav">
            <div className="container px-5">
                <a className="navbar-brand fw-bold" href="#page-top">
                  <img className="navbar-brand-logo" src="/images/nicedishy-logo.png" alt="NiceDishy" />
                </a>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
                    Menu
                    <i className="bi-list"></i>
                </button>
                <div className="collapse navbar-collapse" id="navbarResponsive">
                    <ul className="navbar-nav ms-auto me-4 my-3 my-lg-0">
                        <li className="nav-item"><Link href="/docs"><a className="nav-link me-lg-3">Docs</a></Link></li>
                    </ul>
                    <button className="btn mb-2 mb-lg-0" data-bs-toggle="modal" data-bs-target="#feedbackModal">
                        <span className="d-flex align-items-center">
                            <i className="me-2"></i>
                            <Link href="/login">
                              <a><img src="/images/btn_google_signin_dark_pressed_web@2x.png" alt="Sign in with Google" style={{width: "200px"}} /></a>
                            </Link>
                        </span>
                    </button>
                </div>
            </div>
        </nav>

        <header className="masthead">
            <div className="container px-5">
                <div className="row gx-5 align-items-center">
                    <div className="col-lg-6">
                        <div className="mb-5 mb-lg-0 text-center text-lg-start">
                            <h1 className="display-1 lh-1 mb-3">
                              <img src="/images/nicedishy-word.png" />
                            </h1>
                            <p className="lead fw-normal text-muted mb-5">Continuous monitoring for your Starlink Dishy (Flatface or Squarepants), from anywhere.</p>
                            <p className="lead fw-normal text-muted mb-5">Compare the performance of your Dishy with everyone else and even get some alerts when something changes!</p>
                        </div>
                        <p className="lead fw-normal text-muted mb-5">  Compare the performance of your Dishy with everyone else!</p>
                        <div className="row">
                          <div className="col-12">
                            <Link href="/signup">
                              <a>
                                <img className="app-badge" src="/images/btn_google_signin_dark_pressed_web@2x.png" alt="Get Started" />
                              </a>
                            </Link>
                          </div>
                        </div>
                        <div className="row">
                          <div className="col-12">
                            (To get started, click the Google link)
                          </div>
                        </div>
                    </div>
                    <div className="col-lg-6">
                        <div className="masthead-device-mockup">
                            <img src="/images/nicedishy-icon.png" />
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
                    <div className="col-lg-8 order-lg-1 mb-5 mb-lg-0">
                        <div className="container-fluid px-5">
                            <div className="row gx-5">
                                <div className="col-md-6 mb-5">
                                    <div className="text-center">
                                        <img src="/images/7575373_news_events_mars_emoji.png" />
                                        <h3 className="font-alt">Compare</h3>
                                        <p className="text-muted mb-0">Compare the performance of your Dishy to folks around you or anywhere in the world</p>
                                    </div>
                                </div>
                                <div className="col-md-6 mb-5">
                                    <div className="text-center">
                                        <img src="/images/7734227_rotation_orbit_planet_mars.png" />
                                        <h3 className="font-alt">Monitor</h3>
                                        <p className="text-muted mb-0">Check on the status of your Dishy from anywhere.</p>
                                    </div>
                                </div>
                            </div>
                            <div className="row">
                                <div className="col-md-6 mb-5 mb-md-0">
                                    <div className="text-center">
                                        <img src="/images/7582125_research_explore_planet_exploration_mars_science.png" />
                                        <h3 className="font-alt">Learn</h3>
                                        <p className="text-muted mb-0">Learn what to expect and understand if Starlink is working as expected.</p>
                                    </div>
                                </div>
                                <div className="col-md-6">
                                    <div className="text-center">
                                        <img src="/images/7575375_emoji_image_mars_multimedia.png" />
                                        <h3 className="font-alt">Be Informed</h3>
                                        <p className="text-muted mb-0">Get an email when something changes (firmware, speed).</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="col-lg-4 order-lg-0">
                        <div className="features-device-mockup">
                          <img src="/images/screenshot-small.png" style={{maxWidth: "550px"}} />
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <section className="bg-light">
            <div className="container px-5">
                <div className="row gx-5 align-items-center justify-content-center justify-content-lg-between">
                    <div className="col-12 col-lg-5">
                        <h2 className="display-4 lh-1 mb-4">Let&apos;s build the most complete dataset of dishy stats!</h2>
                        <p className="lead fw-normal text-muted mb-5 mb-lg-0">This section is perfect for featuring some information about your application, why it was built, the problem it solves, or anything else! There&apos;s plenty of space for text here, so don&apos;t worry about writing too much.</p>
                    </div>
                    <div className="col-sm-8 col-md-6">
                        <div className="px-5 px-sm-0"><img className="img-fluid rounded-circle" src="https://source.unsplash.com/u8Jn2rzYIps/900x900" alt="..." /></div>
                    </div>
                </div>
            </div>
        </section>

        <Footer />

    </body>
  )
}
