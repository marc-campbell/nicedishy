import Navbar from './navbar'
import Footer from './footer'

export default function Layout({ children }) {
  return (
    <>
      <div style={{minHeight: "100vh", position: "relative"}}>
        <div className="col-lg-8 mx-auto" style={{paddingBottom: "200px"}}>
          <Navbar />
          <div className="py-md-5 p3">
            <main>{children}</main>
          </div>
        </div>
        <Footer />
      </div>
    </>

  )
}
