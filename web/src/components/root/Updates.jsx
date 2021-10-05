import { Box, Flex, BorderBox, FormGroup } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { withRouter, Link } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

const icons = {
  open: require("../../images/icon-pull-open.png"),
  closed: require("../../images/icon-pull-merged.png"),
  merged: require("../../images/icon-pull-merged.png"),
};

class Updates extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      updateRequests: [],
      isLoading: true,
    }
  }

  componentDidMount = async () => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/context/${this.props.match.params.owner}/${this.props.match.params.context}/pulls`, {
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
        updateRequests: data.updateRequests,
      });
    } catch (err) {
      console.error(err);
    }
  }

  render() {
    if (this.state.isLoading) {
      return <div />
    }

    const rows = this.state.updateRequests.map((updateRequest) => {
      let icon = icons[updateRequest.status];
      if (updateRequest.status === "closed" && updateRequest.mergedAt !== "") {
        icon = icons["merged"];
      }

      return (
        <Box key={updateRequest.id} paddingTop={"6px"}>
          <Flex>
            <Box>
              <FormGroup.Label htmlFor={updateRequest.id}>
                <input type="checkbox" id={updateRequest.id} style={{marginRight: "7px"}} />
              </FormGroup.Label>
            </Box>
            <Box paddingRight={"8px"}>
              <img src={icon} style={{width: "20px"}} />
            </Box>
            <Box>
              <Link to={`/${this.props.match.params.owner}/${this.props.match.params.context}/update/${updateRequest.number}`}>
                <strong>{updateRequest.title}</strong><br />
              </Link>
              #{updateRequest.number} opened on {updateRequest.createdAt} by USER
            </Box>
          </Flex>
        </Box>
      )
    });

    return (
      <Flex alignSelf="center" alignItems="center" justifyContent="center">
        <Box display="flex">
          <Box width={900}>
            <BorderBox>
              <Box style={{backgroundColor: "#ccc"}} paddingLeft="10px" paddingRight="10px">
                Updates
              </Box>
              {rows}
            </BorderBox>
          </Box>
        </Box>
      </Flex>
    );
  }
}

export default withRouter(Updates);
