import { BorderBox, ButtonDanger, Dropdown, Flash, Flex, FormGroup, Grid, Heading, TextInput } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

class ContextSettingsDetails extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      name: "",
      version: "",
    }
  }

  onDeleteContextClick = async() => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/context/${this.props.match.params.owner}/${this.props.match.params.context}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
      });

      if (res.status === 401) {
        this.props.history.replace(`/login?next=/`);
        return;
      }

      if (!res.ok) {
        console.log("error")
        return;
      }

      const data = await res.json();
      if (data.success) {
        this.props.history.replace("/");
      }
    } catch (err) {
      console.error(err);
    }
  }

  render() {
    return (
      <div>
        <Heading fontSize={2}>Settings</Heading>
        <FormGroup>
          <FormGroup.Label htmlFor="context-name">Context name</FormGroup.Label>
          <TextInput id="context-name" value={this.state.name} />
        </FormGroup>
        <FormGroup>
          <FormGroup.Label htmlFor="kubernetes-version">Kubernetes version</FormGroup.Label>
          <Dropdown>
            <Dropdown.Button>1.21.x</Dropdown.Button>
            <Dropdown.Menu direction='sw'>
              <Dropdown.Item>1.20.x</Dropdown.Item>
              <Dropdown.Item>1.21.x</Dropdown.Item>
            </Dropdown.Menu>
          </Dropdown>
        </FormGroup>
        <Heading color="#d73a49" fontSize={2}>Here be dragons</Heading>
        <BorderBox borderColor="border.danger" p="10px">
          <Grid marginLeft="7px" gridTemplateColumns="repeat(2, auto)" gridGap={1}>
            <Grid>
              <strong>Delete this context</strong><br />
              Once you delete, you cannot recover. This is permanent. Think carefully.
            </Grid>
            <Grid>
              <ButtonDanger onClick={this.onDeleteContextClick} variant="small">Delete this context</ButtonDanger>
            </Grid>
          </Grid>
        </BorderBox>
      </div>
    );
  }
}

export default withRouter(ContextSettingsDetails);
