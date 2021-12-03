import { Box, Button, ButtonPrimary, Dropdown, Flash, Flex, Heading, TextInput } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

class NewContext extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isLoading: true,
      availableOwners: [],
      selectedOwner: "",
      selectedContextName: "",
      hasNameConflict: false,
      isCheckingName: false,
    }
  }

  componentDidMount = async () => {
    // get the available owners
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/orgs`, {
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
      const owners = [ Utilities.getClaim("login") ].concat(data.orgs.map((org) => {
        return org.login;
      }));
      const githubOwners = data.githubOrgs.map((org) => {
        return org;
      });

      const availableOrgs = owners.concat(githubOwners);

      owners.push(Utilities.getClaim("login"));

      this.setState({
        isLoading: false,
        availableOwners: availableOrgs,
        selectedOwner: Utilities.getClaim("login"),
      });
    } catch(err) {
      console.log(err);
    }
  }

  onChangeName = async (ev) => {
    this.setState({
      selectedContextName: ev.target.value,
      isCheckingName: true,
    });

    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/contexts/check-name`, {
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
      this.setState({
        hasNameConflict: !data.isAvailable,
        isCheckingName: false,
      });
    } catch(err) {
      console.error(err);
    }
  }

  onSelectOwner = async(owner) => {
    this.setState({
      selectedOwner: owner,
    });
  }

  onCreateClick = async() => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/contexts`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
        body: JSON.stringify({
          "owner": this.state.selectedOwner,
          "name": this.state.selectedContextName,
        }),
      });

      if (!res.ok) {
        console.log("error");
        return;
      }

      const data = await res.json();
      this.props.history.push(`/${data.context.owner}/${data.context.name}`);
    } catch(err) {
      console.error(err);
    }
  }

  render() {
    if (this.state.isLoading) {
      return <div />;
    }

    const ownersItems = this.state.availableOwners.map((owner) => {
      return (
        <Dropdown.Item onClick={this.onSelectOwner.bind(this, owner)} key={owner}>{owner}</Dropdown.Item>
      );
    });

    return (
      <div>
        <Flex alignSelf="center" paddingTop={50} paddingBottom={50} justifyContent="center">
          <Box width={900}>
            <Heading>Create a new context</Heading>
            <p>
              A context is a managed Kubernetes control plane, with a complete API server and
              a GitOps workflow.
            </p>
            <Box display="flex">
              <Box>
                <strong>Owner</strong><br />
                <Dropdown>
                  <Dropdown.Button>{this.state.selectedOwner}</Dropdown.Button>
                  <Dropdown.Menu direction="sw" >
                    {ownersItems}
                  </Dropdown.Menu>
                </Dropdown>
              </Box>
              <Box paddingTop={25} px={3}><h3>/</h3></Box>
              <Box width={200}>
                <strong>Context Name</strong><br />
                <TextInput aria-label="context-name" name="context-name" onChange={this.onChangeName} />
              </Box>
            </Box>
            <Box paddingTop="5px" className={classNames({"hidden": !this.state.hasNameConflict})}>
              <Flash variant="danger">The name {this.state.selectedOwner}/{this.state.selectedContextName} is not available</Flash>
            </Box>
            <Box paddingTop="10px">
              <ButtonPrimary
                onClick={this.onCreateClick}
                disabled={this.state.hasNameConflict || this.state.selectedContextName === ""}
              >Create context</ButtonPrimary>
            </Box>
          </Box>
        </Flex>
      </div>
    );
  }
}

export default withRouter(NewContext);
