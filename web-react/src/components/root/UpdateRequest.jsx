import { Box, Flex, BorderBox, ButtonPrimary } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { withRouter, Link } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

class UpdateRequest extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
      updateRequest: {},
    }
  }

  componentDidMount = async () => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/context/${this.props.match.params.owner}/${this.props.match.params.context}/update/${this.props.match.params.number}`, {
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
        updateRequest: data.updateRequest,
      });
  } catch (err) {
      console.error(err);
    }
  }

  onMergeClick = async() => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/context/${this.props.match.params.owner}/${this.props.match.params.context}/pull/${this.props.match.params.number}/merge`, {
        method: "POST",
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
      console.log(data);
    } catch (err) {
      console.error(err);
    }
  }

  render() {
    return (
      <Flex alignSelf="center" alignItems="center" justifyContent="center">
        <Box display="flex">
          <Box width={900}>
            <BorderBox>
              <Box style={{backgroundColor: "#ccc"}} paddingLeft="10px" paddingRight="10px">
                Update Request
              </Box>
              <Box>
                USER wants to merge 1 commits into <code>{this.state.updateRequest.baseContextBranch}</code> from <code>{this.state.updateRequest.headContextBranch}</code>
              </Box>
              <BorderBox>
                <ButtonPrimary onClick={this.onMergeClick}>Merge</ButtonPrimary>
              </BorderBox>
            </BorderBox>
          </Box>
        </Box>
      </Flex>
    );
  }
}

export default withRouter(UpdateRequest);
