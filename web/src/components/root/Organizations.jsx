import { Box, Button, ButtonPrimary, Dropdown, Flash, Flex, Heading, TextInput } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

class Organizations extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
      orgs: [],
    }
  }

  componentDidMount = async () => {
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
      console.log(data);
    } catch (err) {
      console.error(err)
    }
  }

  onAddMissingOrgsClick = async () => {

  }

  render() {
    return (
      <div>
        Orgs<br />
        <p>
          CentralContext can import and use any organization that you have defined on GitHub. To get started with a new org, create it on GitHub first.
        </p>
        <p>
          If you are a member of any orgs in GitHub that are not on this list, click the button below to add them to your scope!
        </p>
        <Button onClick={this.onAddMissingOrgsClick}>Add Missing Org(s)</Button>
        <br /><br />

      </div>
    );
  }
}

export default withRouter(Organizations);
