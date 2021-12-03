import * as React from "react";
import {Box, Button, Heading, Link} from "@primer/components"
import * as url from "url";

import "../scss/components/login.scss";
import { Utilities } from "../utilities/utilities";

class ClientAuthorize extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
      dishies: [],
      isComplete: false,
    }
  }

  componentDidMount = async() => {
    try {   
      const res = await fetch(`${window.env.API_ENDPOINT}/dishies`, {
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
      
      this.setState({
        isLoading: false,
        dishies: data.dishies,
      });
    } catch (err) {
      console.error(err);
    }
  }

  onSelectDishy = async(id) => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/dishy/${id}/token`, {
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
      window.open(`nicedishy://token?token=${data.token}`)
      this.setState({
        isComplete: true,
      });
    } catch (err) {
      console.error(err);
    }
  }

  render() {
    if (this.state.isLoading) {
      return <div>loading</div>;
    }
    
    if (this.state.isComplete) {
      return <div>You can close this window.</div>
    }
    
    const dishies = this.state.dishies.map((dishy) => {
      return (
        <Link onClick={this.onSelectDishy.bind(this, dishy.id)} key={dishy.id}>
          <Box display="flex" borderColor="#000" bg="#5b6b7c" color="#fff" borderWidth={1} borderStyle="solid" p={3} minWidth={600}>
            <Box flex={1} fontWeight={600} fontSize="2em" color="#fff">{dishy.name}</Box>
            <Box width={60}>---&gt;</Box>
          </Box>
        </Link>
      );
    });

    return (
      <Box flexDirection="column" display="flex" alignSelf="center" paddingTop={50} paddingBottom={50} alignItems="center" justifyContent="center">
        <Box>
          <Heading>Select a dishy</Heading>
        </Box>
        <Box>
          <Box flexDirection="column" display="flex">
            {dishies}
          </Box>
        </Box>
      </Box>
    );
  }
}

export default ClientAuthorize;
