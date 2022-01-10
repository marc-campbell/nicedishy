import Head from 'next/head'
import Navbar from './navbar'
import Footer from './footer'

export default function Layout({ children }) {
  return (
    <>
      <Head>
        <link rel="preconnect" href="https://fonts.gstatic.com" />
        <link href="https://fonts.googleapis.com/css2?family=Newsreader:ital,wght@0,600;1,600&amp;display=swap" rel="stylesheet" />
        <link href="https://fonts.googleapis.com/css2?family=Mulish:ital,wght@0,300;0,500;0,600;0,700;1,300;1,500;1,600;1,700&amp;display=swap" rel="stylesheet" />
        <link href="https://fonts.googleapis.com/css2?family=Kanit:ital,wght@0,400;1,400&amp;display=swap" rel="stylesheet" />
      </Head>
      <div style={{minHeight: "100vh"}}>
        <div className="col-lg-8 mx-auto">
          <Navbar />
            <div className="py-md-5 p3">
              <main>{children}</main>
            </div>
          </div>
        <div style={{position: "absolute", left: "0", bottom: "0", width: "100%"}}>
          <Footer />
        </div>
      </div>
    </>

  )
}
