import * as React from "react";
import * as url from "url";

class LoginCallback extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      getAuthComplete: false,
      next: "/",
    }

  }

  componentDidMount() {
    if (window.localStorage.getItem("token")) {
      const session = window.localStorage.getItem("token");

      this.setState({
        getAuthComplete: true,
      });

      return;
  }

    (async () => {
      let next = "";
      try {
        const query = url.parse(window.location.href, true).query;
        const response = await this.requestSessionToken(query.code, query.state);
        window.localStorage.setItem("token", response.token);
        next = response.redirectUri === "" ? "/" : response.redirectUri;
      } catch (err) {
        console.log(err);
        this.props.history.replace("/error");
        return;
      }
      this.setState({
        getAuthComplete: true,
        next,
      });
    })();
  }

  componentDidUpdate(prevProps, prevState) {
    if (!this.state.getAuthComplete) {
      return;
    }

    const sessionToken = window.localStorage.getItem("token");
    if (!sessionToken) {
      console.log("Invalid or missing authorization credentials");
      this.props.history.replace("/error");
      return;
    }

    this.props.history.replace(this.state.next);
  }

  requestSessionToken = async (code, state) => {
    const uri = `${window.env.API_ENDPOINT}/login/callback`;
    const response = await fetch(uri, {
      method: "post",
      headers: {
        "Content-Type": "application/json",
      },  
      body: JSON.stringify({
        code,
        state,
      }),
    });

    if (!response.ok) {
      return;
    }

    const body = await response.json();
    return body;
  }

  render() {
    return (
      <>
      </>
    );
  }
}

export default LoginCallback;
