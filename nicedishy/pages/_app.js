import { usePostHog } from 'next-use-posthog'

import '../styles/globals.css'
import '../styles/home.css'

const App = ({ Component, pageProps }) => {
  if (process.env.NEXT_PUBLIC_POSTHOG_KEY) {
    usePostHog(process.env.NEXT_PUBLIC_POSTHOG_KEY, {
      api_host: 'https://app.posthog.com',
    })
  }

  const getLayout = Component.getLayout || ((page) => page);
  return getLayout(<Component {...pageProps} />);
}

export default App;
