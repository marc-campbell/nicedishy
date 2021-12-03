import { Box, Button, ButtonPrimary, Dropdown, Flash, Flex, Heading, SideNav, Text } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { NavLink, Route, Switch, withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";
import ContextSettingsDetails from "./ContextSettingsDetails";
import ContextNodeSecurity from "./ContextNodeSecurity";

class ContextSettings extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {

    return (
      <Flex alignSelf="center" alignItems="center" justifyContent="center">
        <Box display="flex">
          <Box width={240} paddingLeft={"10px"}>
            <SideNav bordered aria-label="Context Settings">
              <SideNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}/settings`}>
                <Text>Context Settings</Text>
              </SideNav.Link>
              <SideNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}/settings/node-security`}>
                <Text>Node Security</Text>
              </SideNav.Link>
              <SideNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}/settings/encryption`}>
                <Text>Encryption</Text>
              </SideNav.Link>
              <SideNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}/settings/networking`}>
                <Text>Networking (CNI)</Text>
              </SideNav.Link>
            </SideNav>
          </Box>
          <Box marginLeft={"10px"} width={660}>
            <Switch>
              <Route exact path="/:owner/:context/settings" render={props => <ContextSettingsDetails {...this.props} /> } />
              <Route exact path="/:owner/:context/settings/node-security" render={props => <ContextNodeSecurity {...this.props} /> } />
            </Switch>
          </Box>
        </Box>
      </Flex>
    );
  }
}

export default withRouter(ContextSettings);
