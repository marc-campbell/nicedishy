import * as React from "react";
import { withRouter } from "react-router-dom";
import {Box, Flex, Dropdown, Avatar, StyledOcticon} from '@primer/components';
import { Utilities } from "../../utilities/utilities";

const logo = require("../../images/icon-white.png");

// NavBar will only read from session... for performance reasons, we don't
// make http requests from this component
export class NavBar extends React.Component {

  onNavItemClick = (tag) => {
    if (tag === "logout") {
      Utilities.logoutUser();
      return;
    }
  }

  render() {
    return (
      <Flex flexDirection="row" width="100%" height={60} px={40} paddingTop="17px" backgroundColor="#000" color="#fff">
        <Box display="flex" flexGrow={1}>
          <img src={logo} className="logo" />
        </Box>
        <Box paddingLeft={20}>
          <Dropdown>
            <summary>
              <Avatar src="https://avatars.githubusercontent.com/primer" />
              <Dropdown.Caret/>
            </summary>
            <Dropdown.Menu direction='sw'>
              <Dropdown.Item onClick={this.onNavItemClick.bind(this, "settings")}>Settings</Dropdown.Item>
              <Dropdown.Item onClick={this.onNavItemClick.bind(this, "logout")}>Logout</Dropdown.Item>
            </Dropdown.Menu>
          </Dropdown>
        </Box>
      </Flex>
    )
  }
}

export default withRouter(NavBar);
