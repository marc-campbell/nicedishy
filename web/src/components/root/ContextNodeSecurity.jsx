import { FormGroup, Heading, Grid, Box, ButtonDanger } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

class ContextNodeSecurity extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
      settings: {},
    }
  }

  componentDidMount = async() => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/context/${this.props.match.params.owner}/${this.props.match.params.context}/settings`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
      });

      if (res.status === 401) {
        this.props.history.replace(`/login?next=`);
        return;
      }

      if (!res.ok) {
        console.log("error");
        return;
      }

      const data = await res.json();
      this.setState({
        isLoading: false,
        settings: data,
      });

    } catch(err) {
      console.error(err);
    }
  }

  render() {
    if (this.state.isLoading) {
      return <div />;
    }

    let tokens = this.state.settings.nodeJoinTokens.map((token) => {
      return (
        [
          <Box key={`${token.id}-name`}>
            <strong>{token.name}</strong><br />
            ****************
          </Box>,
          <Box key={`${token.id}-lastused`}>
            {
              token.lastUsedAt ? `Last used at ${token.lastUsedAt}` : "Never used"
            }
          </Box>
        ]
      );
    });

    return (
      <div>
        <Heading fontSize={2}>Node Security</Heading>
        <FormGroup>
          <FormGroup.Label htmlFor="allow-nodes">
            <input type="checkbox" id="allow-nodes" style={{marginRight: "7px"}} />
            Allow new nodes to join this context<br />
            <span style={{marginLeft: "20px", fontSize: "0.8em", fontWeight: "400"}}>
              When checked, new nodes that are installed using the install script will
              automatically join the context.
            </span>
          </FormGroup.Label>
        </FormGroup>
        <FormGroup>
          <FormGroup.Label htmlFor="join-tokens">
            Join tokens
          </FormGroup.Label>
          <Grid marginLeft="7px" gridTemplateColumns="repeat(3, auto)" gridGap={1}>
            {tokens}
          </Grid>
        </FormGroup>
      </div>
    );
  }
}

export default withRouter(ContextNodeSecurity);
