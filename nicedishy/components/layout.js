import Navbar from './navbar'
import Footer from './footer'

export default function Layout({ children }) {
  return (
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
  )
}
