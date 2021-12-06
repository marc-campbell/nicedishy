import Navbar from './navbar'
import Footer from './footer'

export default function Layout({ children }) {
  return (
    <>
      <Navbar />
      <div className="col-lg-8 mx-auto p-3 py-md-5">
        <main>{children}</main>
      </div>
      <Footer />
    </>
  )
}
