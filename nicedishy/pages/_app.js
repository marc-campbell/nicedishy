import '../styles/globals.css'
import frontStyles from '../styles/front.css'

function MyApp({ Component, pageProps }) {
  const getLayout = Component.getLayout || ((page) => page);
  return getLayout(<Component {...pageProps} />);
}

export default MyApp
