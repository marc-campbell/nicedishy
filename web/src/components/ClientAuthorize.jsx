import * as React from 'react';
import {Box, Flex, Button, Heading} from '@primer/components'
import * as url from "url";

import "../scss/components/login.scss";
import { Utilities } from '../utilities/utilities';

class ClientAuthorize extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
      dishies: [],
    }
  }

  componentDidMount = async() => {
    if (!Utilities.isLoggedIn()) {
      return;
    }

    try {
      const query = url.parse(window.location.href, true).query;
      const res = await fetch(`${window.env.API_ENDPOINT}/auth-cli?code=${query.code}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
      });

      if (!res.ok) {
        return;
      }

      const data = await res.json();
      console.log(data);
    } catch (err) {
      console.error(err);
    }
  }

  onClickLogin = async () => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/login`);
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
    if (this.state.isLoading) {
      return <div>loading</div>;
    }

    return (
      <Flex alignSelf="center" paddingTop={50} paddingBottom={50} alignItems="center" justifyContent="center">
        <Box display="flex">
          <Box width={800}>
            <Heading>Log in to CentralContext</Heading>
            <strong>You will be taken to GitHub to authenticate.</strong>
            <p>
              By logging in, you are agreeing to our Terms of Service and Privacy Policy. We ask
              for read access to your GitHub profile in order to provide a complete experience
              here. We don't ask for permissions to change anything or to read your code on GitHub.
            </p>
          </Box>
          <Box width={100}></Box>
          <Box width={300} marginTop={30}>
            <Button width="80%" onClick={this.onClickLogin}>Log In With GitHub</Button>
          </Box>
        </Box>
      </Flex>
    );
  }
}

export default ClientAuthorize;
