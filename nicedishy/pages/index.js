import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import Link from 'next/link'

export default function Home() {
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
              <img className="navbar-brand-logo" src="./assets/svg/logos/logo.svg" alt="Logo" />
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
                  <li className="hs-has-mega-menu nav-item">
                    <a id="landingsMegaMenu" className="hs-mega-menu-invoker nav-link dropdown-toggle active" aria-current="page" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">Landings</a>

                    <div className="hs-mega-menu dropdown-menu w-100" aria-labelledby="landingsMegaMenu" style={{"minWidth": "30rem"}}>
                      <div className="row">
                        <div className="col-lg-6 d-none d-lg-block">
                          <div className="navbar-dropdown-menu-banner" style={{"backgroundImage": "url(./assets/svg/components/shape-3.svg)"}}>
                            <div className="navbar-dropdown-menu-banner-content">
                              <div className="mb-4">
                                <span className="h2 d-block">Branding Works</span>
                                <p>Experience a level of our quality in both design &amp; customization works.</p>
                              </div>
                              <a className="btn btn-primary btn-transition" href="#">Learn more <i className="bi-chevron-right small"></i></a>
                            </div>
                          </div>
                        </div>

                        <div className="col-lg-6">
                          <div className="navbar-dropdown-menu-inner">
                            <div className="row">
                              <div className="col-sm mb-3 mb-sm-0">
                                <span className="dropdown-header">Classic</span>
                                <a className="dropdown-item active" href="./landing-classic-corporate.html">Corporate</a>
                                <a className="dropdown-item " href="./landing-classic-analytics.html">Analytics <span className="badge bg-primary rounded-pill ms-1">Hot</span></a>
                                <a className="dropdown-item " href="./landing-classic-studio.html">Studio</a>
                                <a className="dropdown-item " href="./landing-classic-marketing.html">Marketing</a>
                                <a className="dropdown-item " href="./landing-classic-advertisement.html">Advertisement</a>
                                <a className="dropdown-item " href="./landing-classic-consulting.html">Consulting</a>
                                <a className="dropdown-item " href="./landing-classic-portfolio.html">Portfolio</a>
                                <a className="dropdown-item " href="./landing-classic-software.html">Software</a>
                                <a className="dropdown-item " href="./landing-classic-business.html">Business</a>
                              </div>

                              <div className="col-sm">
                                <div className="mb-3">
                                  <span className="dropdown-header">App</span>
                                  <a className="dropdown-item " href="./landing-app-ui-kit.html">UI Kit</a>
                                  <a className="dropdown-item " href="./landing-app-saas.html">SaaS</a>
                                  <a className="dropdown-item " href="./landing-app-workflow.html">Workflow</a>
                                  <a className="dropdown-item " href="./landing-app-payment.html">Payment</a>
                                  <a className="dropdown-item " href="./landing-app-tool.html">Tool</a>
                                </div>

                                <span className="dropdown-header">Onepage</span>
                                <a className="dropdown-item " href="./landing-onepage-corporate.html">Corporate</a>
                                <a className="dropdown-item " href="./landing-onepage-saas.html">SaaS <span className="badge bg-primary rounded-pill ms-1">Hot</span></a>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </li>

                  <li className="hs-has-sub-menu nav-item">
                    <a id="companyMegaMenu" className="hs-mega-menu-invoker nav-link dropdown-toggle " href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">Company</a>

                    <div className="hs-sub-menu dropdown-menu" aria-labelledby="companyMegaMenu" style={{"minWidth": "14rem"}}>
                      <a className="dropdown-item " href="./page-about.html">About</a>
                      <a className="dropdown-item " href="./page-services.html">Services</a>
                      <a className="dropdown-item " href="./page-customer-stories.html">Customer Stories</a>
                      <a className="dropdown-item " href="./page-customer-story.html">Customer Story</a>
                      <a className="dropdown-item " href="./page-careers.html">Careers</a>
                      <a className="dropdown-item " href="./page-careers-overview.html">Careers Overview</a>
                      <a className="dropdown-item " href="./page-hire-us.html">Hire Us</a>
                      <a className="dropdown-item " href="./page-pricing.html">Pricing</a>
                      <a className="dropdown-item " href="./page-contacts-agency.html">Contacts: Agency</a>
                      <a className="dropdown-item " href="./page-contacts-startup.html">Contacts: Startup</a>
                    </div>
                  </li>

                  <li className="hs-has-sub-menu nav-item">
                    <a id="accountMegaMenu" className="hs-mega-menu-invoker nav-link dropdown-toggle " href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">Account</a>

                    <div className="hs-sub-menu dropdown-menu" aria-labelledby="accountMegaMenu" style={{"minWidth": "14rem"}}>
                      <div className="hs-has-sub-menu nav-item">
                        <a id="authenticationMegaMenu" className="hs-mega-menu-invoker dropdown-item dropdown-toggle " href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">Authentication</a>

                        <div className="hs-sub-menu dropdown-menu" aria-labelledby="authenticationMegaMenu" style={{"minWidth": "14rem"}}>
                          <a className="dropdown-item" href="#" data-bs-toggle="modal" data-bs-target="#signupModal">Signup Modal</a>
                          <div className="dropdown-divider"></div>
                          <a className="dropdown-item " href="./page-login.html">Login</a>
                          <a className="dropdown-item " href="./page-signup.html">Signup</a>
                          <a className="dropdown-item " href="./page-reset-password.html">Reset Password</a>
                          <div className="dropdown-divider"></div>
                          <a className="dropdown-item " href="./page-login-simple.html">Login Simple</a>
                          <a className="dropdown-item " href="./page-signup-simple.html">Signup Simple</a>
                          <a className="dropdown-item " href="./page-reset-password-simple.html">Reset Password Simple</a>
                        </div>
                      </div>

                      <a className="dropdown-item " href="./account-overview.html">Personal Info</a>
                      <a className="dropdown-item " href="./account-security.html">Security</a>
                      <a className="dropdown-item " href="./account-notifications.html">Notifications</a>
                      <a className="dropdown-item " href="./account-preferences.html">Preferences</a>
                      <a className="dropdown-item " href="./account-orders.html">Orders</a>
                      <a className="dropdown-item " href="./account-wishlist.html">Wishlist</a>
                      <a className="dropdown-item " href="./account-payments.html">Payments</a>
                      <a className="dropdown-item " href="./account-address.html">Address</a>
                      <a className="dropdown-item " href="./account-teams.html">Teams</a>
                    </div>
                  </li>

                  <li className="hs-has-sub-menu nav-item">
                    <a id="pagesMegaMenu" className="hs-mega-menu-invoker nav-link dropdown-toggle " href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">Pages</a>

                    <div className="hs-sub-menu dropdown-menu" aria-labelledby="pagesMegaMenu" style={{"minWidth": "14rem"}}>
                      <a className="dropdown-item " href="./page-faq.html">FAQ</a>
                      <a className="dropdown-item " href="./page-terms.html">Terms &amp; Conditions</a>
                      <a className="dropdown-item " href="./page-privacy.html">Privacy &amp; Policy</a>
                      <a className="dropdown-item " href="./page-coming-soon.html">Coming Soon</a>
                      <a className="dropdown-item " href="./page-maintenance-mode.html">Maintenance Mode</a>
                      <a className="dropdown-item " href="./page-status.html">Status</a>
                      <a className="dropdown-item " href="./page-invoice.html">Invoice</a>
                      <a className="dropdown-item " href="./page-error-404.html">Error 404</a>
                    </div>
                  </li>

                  <li className="hs-has-sub-menu nav-item">
                    <a id="blogMegaMenu" className="hs-mega-menu-invoker nav-link dropdown-toggle " href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">Blog</a>

                    <div className="hs-sub-menu dropdown-menu" aria-labelledby="blogMegaMenu" style={{"minWidth": "14rem"}}>
                      <a className="dropdown-item " href="./blog-journal.html">Journal</a>
                      <a className="dropdown-item " href="./blog-metro.html">Metro</a>
                      <a className="dropdown-item " href="./blog-newsroom.html">Newsroom</a>
                      <div className="dropdown-divider"></div>
                      <a className="dropdown-item " href="./blog-article.html">Article</a>
                      <a className="dropdown-item " href="./blog-author-profile.html">Author Profile</a>
                    </div>
                  </li>

                  <li className="hs-has-sub-menu nav-item">
                    <a id="portfolioMegaMenu" className="hs-mega-menu-invoker nav-link dropdown-toggle " href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">Portfolio</a>

                    <div className="hs-sub-menu dropdown-menu" aria-labelledby="portfolioMegaMenu" style={{"minWidth": "14rem"}}>
                      <a className="dropdown-item " href="./portfolio-grid.html">Grid</a>
                      <a className="dropdown-item " href="./portfolio-product-article.html">Product Article</a>
                      <a className="dropdown-item " href="./portfolio-case-studies-branding.html">Case Studies: Branding</a>
                      <a className="dropdown-item " href="./portfolio-case-studies-product.html">Case Studies: Product</a>
                    </div>
                  </li>

                  <li className="nav-item">
                    <a className="btn btn-primary btn-transition" href="/login">Log In</a>
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
                    <a className="btn btn-primary btn-transition px-6" href="./page-login-simple.html">Get started</a>
                    <a className="btn btn-link" href="#">See the stats <i className="bi-chevron-right small ms-1"></i></a>
                  </div>
                </div>
              </div>
            </div>

            <div className="col-lg-7 col-xl-6 d-none d-lg-block position-absolute top-0 end-0 pe-0" style={{"marginTop": "6.75rem"}}>
              <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" viewBox="0 0 1137.5 979.2">
                <path fill="#F9FBFF" d="M565.5,957.4c81.1-7.4,155.5-49.3,202.4-115.7C840,739.8,857,570,510.7,348.3C-35.5-1.5-4.2,340.3,2.7,389
                  c0.7,4.7,1.2,9.5,1.7,14.2l29.3,321c14,154.2,150.6,267.8,304.9,253.8L565.5,957.4z"/>
                <defs>
                  <path id="mainHeroSVG1" d="M1137.5,0H450.4l-278,279.7C22.4,430.6,24.3,675,176.8,823.5l0,0C316.9,960,537.7,968.7,688.2,843.6l449.3-373.4V0z"/>
                </defs>

                <g transform="matrix(1 0 0 1 0 0)">
                  <image width="750" height="750" transform="matrix(1.4462 0 0 1.4448 52.8755 0)"></image>
                </g>
              </svg>
            </div>
          </div>
        </div>

        <div className="container content-space-2 content-space-t-xl-3 content-space-b-lg-3">
          <div className="w-md-75 w-lg-50 text-center mx-md-auto mb-5">
            <h2>Benefits of connecting to NiceDishy</h2>
          </div>

          <div className="row mb-5 mb-md-0">
            <div className="col-sm-6 col-lg-4 mb-4 mb-lg-0">
              <div className="card card-sm h-100">
                <div className="p-2">
                  <img className="card-img" src="./assets/img/480x220/img1.jpg" alt="Image Description"/ >
                </div>

                <div className="card-body">
                  <h4 className="card-title">Compare</h4>
                  <p className="card-text">Compare the performance of your Dishy to folks around you or anywhere in the world.</p>

                  <ul className="list-pointer mb-0">
                    <li className="list-pointer-item">Advanced Analytics</li>
                    <li className="list-pointer-item">Digital Marketing</li>
                    <li className="list-pointer-item">Organization</li>
                  </ul>
                </div>

                <a className="card-footer card-link border-top" href="#">Learn more <i className="bi-chevron-right small ms-1"></i></a>
              </div>
            </div>

            <div className="col-sm-6 col-lg-4 mb-4 mb-lg-0">
              <div className="card card-sm h-100">
                <div className="p-2">
                  <img className="card-img" src="./assets/img/480x220/img2.jpg" alt="Image Description" />
                </div>

                <div className="card-body">
                  <h4 className="card-title">Monitor</h4>
                  <p className="card-text">Check on the status of your Dishy from anywhere.</p>

                  <ul className="list-pointer mb-0">
                    <li className="list-pointer-item">Cost Transformation</li>
                    <li className="list-pointer-item">Customer Experience</li>
                    <li className="list-pointer-item">Mergers and Acquisitions</li>
                  </ul>
                </div>

                <a className="card-footer card-link border-top" href="#">Learn more <i className="bi-chevron-right small ms-1"></i></a>
              </div>
            </div>

            <div className="col-sm-6 col-lg-4">
              <div className="card card-sm h-100">
                <div className="p-2">
                  <img className="card-img" src="./assets/img/480x220/img3.jpg" alt="Image Description" />
                </div>

                <div className="card-body">
                  <h4 className="card-title">Learn</h4>
                  <p className="card-text">Learn what to expect and understand if Starlink is working as expected.</p>

                  <ul className="list-pointer mb-0">
                    <li className="list-pointer-item">Enterprise Technology</li>
                    <li className="list-pointer-item">Private Equity</li>
                    <li className="list-pointer-item">Sustainability</li>
                  </ul>
                </div>

                <a className="card-footer card-link border-top" href="#">Learn more <i className="bi-chevron-right small ms-1"></i></a>
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

            <div className="text-center mb-10">
              <ul className="list-inline list-checked list-checked-primary">
                <li className="list-inline-item list-checked-item">Responsive</li>
                <li className="list-inline-item list-checked-item">5-star support</li>
                <li className="list-inline-item list-checked-item">Constant updates</li>
              </ul>
            </div>

            <div className="row">
              <div className="col-lg-7 mb-9 mb-lg-0">
                <div className="pe-lg-6">
                  <figure className="device-browser">
                    <div className="device-browser-header">
                      <div className="device-browser-header-btn-list">
                        <span className="device-browser-header-btn-list-btn"></span>
                        <span className="device-browser-header-btn-list-btn"></span>
                        <span className="device-browser-header-btn-list-btn"></span>
                      </div>
                      <div className="device-browser-header-browser-bar">www.htmlstream.com/front/</div>
                    </div>

                    <div className="device-browser-frame">
                      <img className="device-browser-img" src="./assets/img/1618x1010/img6.jpg" alt="Image Description" />
                    </div>
                  </figure>
                </div>
              </div>

              <div className="col-lg-5">
                <div className="mb-4">
                  <h2>Collaborative tools to design user experience</h2>
                  <p>We help businesses bring ideas to life in the digital world, by designing and implementing the technology tools that they need to win.</p>
                </div>

                <ul className="list-checked list-checked-primary mb-5">
                  <li className="list-checked-item">Less routine – more creativity</li>
                  <li className="list-checked-item">Hundreds of thousands saved</li>
                  <li className="list-checked-item">Scale budgets efficiently</li>
                </ul>

                <a className="btn btn-primary" href="#">Get started</a>

                <hr className="my-5" />

                <span className="d-block">Trusted by leading companies</span>

                <div className="row">
                  <div className="col py-3">
                    <img className="avatar avatar-4x3" src="./assets/svg/brands/fitbit-dark.svg" alt="Logo" />
                  </div>

                  <div className="col py-3">
                    <img className="avatar avatar-4x3" src="./assets/svg/brands/forbes-dark.svg" alt="Logo" />
                  </div>

                  <div className="col py-3">
                    <img className="avatar avatar-4x3" src="./assets/svg/brands/mailchimp-dark.svg" alt="Logo" />
                  </div>

                  <div className="col py-3">
                    <img className="avatar avatar-4x3" src="./assets/svg/brands/layar-dark.svg" alt="Logo" />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

      </main>
    </div>
  )
}
