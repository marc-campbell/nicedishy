import * as React from "react";
import { Button, Card, Elevation } from "@blueprintjs/core";
import { Grid, Row, Col } from "react-flexbox-grid";
import * as url from "url";

class LoginCallback extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      getAuthComplete: false,
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
      try {
        const query = url.parse(window.location.href, true).query;
        const sessionToken = await this.requestSessionToken(query.code);
        window.localStorage.setItem("token", sessionToken);
      } catch (err) {
        console.log(err);
        this.props.history.replace("/error");
        return;
      }
      this.setState({
        getAuthComplete: true,
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

    this.setState({
      getAuthComplete: false,
    });

    this.props.history.replace("/");
  }

  requestSessionToken = async (code) => {
    const uri = `${window.env.API_ENDPOINT}/login/callback`;
    const response = await fetch(uri, {
      method: "post",
      headers: {
        "Content-Type": "application/json",
      },  
      body: JSON.stringify({
        code,
      }),
    });

    if (!response.ok) {
      return;
    }

    const body = await response.json();
    return body.token;
  }

  render() {
    return (
      <>
      </>
    );
  }
}

export default LoginCallback;
