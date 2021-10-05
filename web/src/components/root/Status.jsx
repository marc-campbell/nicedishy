import { Box, Button, ButtonPrimary, Dropdown, Flash, Flex, Heading, TextInput } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

class Status extends React.Component {
  constructor(props) {
    super(props);

  }

  render() {
    return (
      <div>
        status
      </div>
    );
  }
}

export default withRouter(Status);
