import * as React from 'react';
import {Box, Flex, Button, Heading} from '@primer/components'

import "../scss/components/login.scss";
import { Utilities } from '../utilities/utilities';

class Login extends React.Component {
  onClickLogin = async () => {
    Utilities.logoutUser();

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
    return (
      <Flex alignSelf="center" paddingTop={50} paddingBottom={50} alignItems="center" justifyContent="center">
        <Box display="flex">
          <Box width={800}>
            <Heading>Log in to NiceDishy</Heading>
            <strong>You will be taken to Google to authenticate.</strong>
            <p>
              By logging in, you are agreeing to our Terms of Service and Privacy Policy. We ask
              for read access to your Google profile in order to provide a complete experience
              here. We don't ask for permissions to change anything in your Google account.
            </p>
          </Box>
          <Box width={100}></Box>
          <Box width={300} marginTop={30}>
            <Button width="80%" onClick={this.onClickLogin}>Log In With Google</Button>
          </Box>
        </Box>
      </Flex>
    );
  }
}

export default Login;
