import { BorderBox, Box, Flex } from "@primer/components";
import * as React from "react";
import { withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

class Nodes extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
      nodes: {},
      isEmpty: false,
    }
  }

  componentDidMount = async () => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/context/${this.props.match.params.owner}/${this.props.match.params.context}/nodes`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
      });

      if (!res.ok) {
        console.log("error");
        return;
      }

      const data = await res.json();
      this.setState({
        isLoading: false,
        nodes: data.nodes,
        joinToken: data.nodeJoinToken,
      });

    } catch (err) {
      console.error(err);
    }
  }

  render() {
    if (this.state.isLoading) {
      return <div />
    }

    if (!this.state.nodes) {
      return (
        <Flex alignSelf="center" alignItems="center" justifyContent="center">
          <Box display="flex">
            <Box width={900}>
              <BorderBox paddingLeft="10px" paddingRight="10px">
                <h3>Quick and easy command to add a node</h3>
                <pre>
                  $ kubectl krew install centralcontext<br />
                  $ kubectl centralcontext join {this.props.match.params.owner}/{this.props.match.params.context} --token {this.state.joinToken}
                </pre>
                <p>
                  The command above will install <code>kubelet</code> and configure it
                  to join this context. Once joined, pods and workloads will be scheduled
                  onto the new node.
                </p>
                The token in the above command is a secret, and should not be shared.
              </BorderBox>
              <BorderBox marginTop="10px" paddingLeft="10px" paddingRight="10px">
                <h3>...or get a script</h3>
                Coming soon
              </BorderBox>
              <BorderBox marginTop="10px" paddingLeft="10px" paddingRight="10px">
                <h3>...or add an existing node</h3>
                Coming soon
              </BorderBox>
            </Box>
          </Box>
        </Flex>
      );
    }

    return (
      <div>
        nodes
      </div>
    );
  }
}

export default withRouter(Nodes);
