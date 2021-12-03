import { SideNav, Text, Flex, Box } from "@primer/components";
import * as React from "react";
import { NavLink, Route, Switch, withRouter } from "react-router-dom";
import Organizations from "./Organizations";

class Settings extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isLoading: true,
      selectedTab: "manifests",
    }
  }

  render() {
    return (
      <Flex alignSelf="center" alignItems="center" justifyContent="center">
        <Box display="flex">
          <Box width={240} paddingLeft={"10px"}>
            <SideNav bordered aria-label="Context Settings">
              <SideNav.Link as={NavLink} to={`/settings/profile`}>
                <Text>Profile</Text>
              </SideNav.Link>
              <SideNav.Link as={NavLink} to={`/settings/account`}>
                <Text>Account</Text>
              </SideNav.Link>
              <SideNav.Link as={NavLink} to={`/settings/organizations`}>
                <Text>Organizations</Text>
              </SideNav.Link>
            </SideNav>
          </Box>
          <Box marginLeft={"10px"} width={660}>
            <Switch>
              <Route exact path="/settings/organizations" render={props => <Organizations {...this.props} /> } />
            </Switch>
          </Box>
        </Box>
      </Flex>
    );
  }
}

export default withRouter(Settings);
