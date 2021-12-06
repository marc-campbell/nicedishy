import Navbar from './navbar'
import Footer from './footer'

export default function Layout({ children }) {
  return (
    <>
      <div className="col-lg-8 mx-auto">
        <Navbar />
          <div className="py-md-5 p3">
            <main>{children}</main>
          </div>
        <Footer />
      </div>
    </>
  )
}
