import * as React from "react";
import * as url from "url";

import { Utilities } from "../../utils/utilities";

class Login extends React.Component {
  onClickLogin = async () => {
    // ensure the user is logged out
    Utilities.logoutUser();

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/login?next=${encodeURIComponent(url.parse(window.location.href, true).query.next)}`);
      if (!res.ok) {
        console.log("error")
        return;
      }

      const data = await res.json();
      window.location.href = data.redirectURL;
    } catch(err) {
      console.log(err);
    }
  }

  render() {
    return (
      <>
        <h1>Log in to NiceDishy</h1>
        <strong>You will be taken to Google to authenticate.</strong>
        <p>
          By logging in, you are agreeing to our Terms of Service and Privacy Policy. We ask
          for read access to your Google profile in order to provide a complete experience
          here. We don't ask for permissions to change anything in your Google account.
        </p>
        <button width="80%" onClick={this.onClickLogin}>Log In With Google</button>
      </>
    );
  }
}

export default Login;
