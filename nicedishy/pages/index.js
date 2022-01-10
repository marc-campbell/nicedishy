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
                        <li className="nav-item"><a className="nav-link me-lg-3" href="#features">Docs</a></li>
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
                            <p className="lead fw-normal text-muted mb-5">  Compare the performance of your Dishy with everyone else!</p>
                            <div className="d-flex flex-column flex-lg-row align-items-center">
                                <a className="me-lg-3 mb-4 mb-lg-0" href="#!"><img className="app-badge" src="assets/img/google-play-badge.svg" alt="..." /></a>
                                <a href="#!"><img className="app-badge" src="assets/img/app-store-badge.svg" alt="..." /></a>
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
                    <div className="col-xl-8">
                        <div className="h2 fs-1 text-white mb-4">Woo. It's free and it's exciting.</div>
                        <img src="assets/img/tnw-logo.svg" alt="..." style={{height: "3rem"}} />
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
                                        <i className="bi-phone icon-feature text-gradient d-block mb-3"></i>
                                        <h3 className="font-alt">Compare</h3>
                                        <p className="text-muted mb-0">Compare the performance of your Dishy to folks around you or anywhere in the world</p>
                                    </div>
                                </div>
                                <div className="col-md-6 mb-5">
                                    <div className="text-center">
                                        <i className="bi-camera icon-feature text-gradient d-block mb-3"></i>
                                        <h3 className="font-alt">Monitor</h3>
                                        <p className="text-muted mb-0">heck on the status of your Dishy from anywhere.</p>
                                    </div>
                                </div>
                            </div>
                            <div className="row">
                                <div className="col-md-6 mb-5 mb-md-0">
                                    <div className="text-center">
                                        <i className="bi-gift icon-feature text-gradient d-block mb-3"></i>
                                        <h3 className="font-alt">Learn</h3>
                                        <p className="text-muted mb-0">Learn what to expect and understand if Starlink is working as expected.</p>
                                    </div>
                                </div>
                                <div className="col-md-6">
                                    <div className="text-center">
                                        <i className="bi-patch-check icon-feature text-gradient d-block mb-3"></i>
                                        <h3 className="font-alt">?? </h3>
                                        <p className="text-muted mb-0">?? ??? ?? ? ?</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="col-lg-4 order-lg-0">
                        <div className="features-device-mockup">
                            <svg className="circle" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg">
                                <defs>
                                    <linearGradient id="circleGradient" gradientTransform="rotate(45)">
                                        <stop className="gradient-start-color" offset="0%"></stop>
                                        <stop className="gradient-end-color" offset="100%"></stop>
                                    </linearGradient>
                                </defs>
                                <circle cx="50" cy="50" r="50"></circle></svg
                            ><svg className="shape-1 d-none d-sm-block" viewBox="0 0 240.83 240.83" xmlns="http://www.w3.org/2000/svg">
                                <rect x="-32.54" y="78.39" width="305.92" height="84.05" rx="42.03" transform="translate(120.42 -49.88) rotate(45)"></rect>
                                <rect x="-32.54" y="78.39" width="305.92" height="84.05" rx="42.03" transform="translate(-49.88 120.42) rotate(-45)"></rect></svg
                            ><svg className="shape-2 d-none d-sm-block" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg"><circle cx="50" cy="50" r="50"></circle></svg>
                            <div className="device-wrapper">
                                <div className="device" data-device="iPhoneX" data-orientation="portrait" data-color="black">
                                    <div className="screen bg-black">
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
        <section className="bg-light">
            <div className="container px-5">
                <div className="row gx-5 align-items-center justify-content-center justify-content-lg-between">
                    <div className="col-12 col-lg-5">
                        <h2 className="display-4 lh-1 mb-4">Let's build the most complete dataset of dishy stats!</h2>
                        <p className="lead fw-normal text-muted mb-5 mb-lg-0">This section is perfect for featuring some information about your application, why it was built, the problem it solves, or anything else! There's plenty of space for text here, so don't worry about writing too much.</p>
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
