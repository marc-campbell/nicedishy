import '../styles/globals.css'
import '../styles/home.css'
import posthog from 'posthog-js';

const App = ({ Component, pageProps }) => {
  if (typeof window !== undefined && process.env.NEXT_PUBLIC_POSTHOG_KEY) {
    posthog.init(process.env.NEXT_PUBLIC_POSTHOG_KEY, {api_host: 'https://app.posthog.com'})
  }

  const getLayout = Component.getLayout || ((page) => page);
  return getLayout(<Component {...pageProps} />);
}

export default App;
