import { Box, Button, ButtonPrimary, Dropdown, Flash, Flex, Heading, TabNav, TextInput, UnderlineNav } from "@primer/components";
import * as React from "react";
import { NavLink, Route, Switch, withRouter } from "react-router-dom";
import ContextSettings from "./ContextSettings";
import Manifests from "./Manifests";
import Nodes from "./Nodes";
import Security from "./Security";
import Status from "./Status";
import Updates from "./Updates";
import UpdateRequest from "./UpdateRequest";
import AddGitHubRelease from "./AddGitHubRelease";

class Context extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isLoading: true,
      selectedTab: "manifests",
    }
  }

  render() {
    return (
      <div>
        <Flex px={20} alignSelf="center" justifyContent="center" width={900}>
          <UnderlineNav aria-label="Main">
            <UnderlineNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}`}>Manifests</UnderlineNav.Link>
            <UnderlineNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}/nodes`}>Nodes</UnderlineNav.Link>
            <UnderlineNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}/security`}>Security</UnderlineNav.Link>
            <UnderlineNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}/pulls`}>Pull Requests</UnderlineNav.Link>
            <UnderlineNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}/status`}>Status</UnderlineNav.Link>
            <UnderlineNav.Link as={NavLink} to={`/${this.props.match.params.owner}/${this.props.match.params.context}/settings`}>Settings</UnderlineNav.Link>
          </UnderlineNav>
        </Flex>
        <div style={{paddingTop: "20px", paddingBottom: "20px"}}>
          <Switch>
            <Route exact path="/:owner/:context" render={props => <Manifests {...this.props} /> } />
            <Route exact path="/:owner/:context/add/github-release" render={props => <AddGitHubRelease {...this.props} /> } />
            <Route path="/:owner/:context/nodes" render={props => <Nodes {...this.props} /> } />
            <Route path="/:owner/:context/security" render={props => <Security {...this.props} /> } />
            <Route path="/:owner/:context/pulls" render={props => <Updates {...this.props} /> } />
            <Route path="/:owner/:context/update/:number" render={props => <UpdateRequest {...this.props} /> } />
            <Route path="/:owner/:context/status" render={props => <Status {...this.props} /> } />
            <Route path="/:owner/:context/settings" render={props => <ContextSettings {...this.props} /> } />
          </Switch>
        </div>
      </div>
    );
  }
}

export default withRouter(Context);
