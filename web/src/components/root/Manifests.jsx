import { BorderBox, Box, Grid, Flex, Dropdown, Link } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

const icons = {
  blob: require("../../images/blob-icon.png"),
  tree: require("../../images/tree-icon.png"),
};

class Manifests extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
      tree: {},
      isEmpty: false,
    }
  }

  componentDidMount = async () => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/tree?owner=${this.props.match.params.owner}&context=${this.props.match.params.context}`, {
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
      this.setState({
        isLoading: false,
        tree: data.tree,
        isEmpty: data.isEmpty,
      });

    } catch (err) {
      console.error(err);
    }
  }

  onClickAddGitHubRelease = async () => {
    this.props.history.push(`/${this.props.match.params.owner}/${this.props.match.params.context}/add/github-release`)
  }

  onClickGitItem = async (t) => {
    if (t.type === "blob") {
      this.props.history.push(`/${this.props.match.params.owner}/${this.props.match.params.context}/blob/master/${t.path}`);
    } else {
      console.log("unknown type");
    }
  }

  render() {
    if (this.state.isLoading) {
      return <div />
    }

    if (this.state.isEmpty) {
      return (
        <Flex alignSelf="center" alignItems="center" justifyContent="center">
          <Box display="flex">
            <Box width={900}>
              <BorderBox marginTop="10px" paddingLeft="10px" paddingRight="10px">
                <h3>...or push existing manifests from the command line</h3>
                <pre>
                  git remote add origin git@centralcontext.com:retagged/test.git<br />
                  git branch -M main<br />
                  git push -u origin main<br />
                </pre>
              </BorderBox>
            </Box>
          </Box>
        </Flex>
      );
    }

    const rows = this.state.tree.map((t) => {
      return (
        [
          <Box key={`${t.path}-type`} style={{flexShrink: 0, paddingLeft: "16px", paddingTop: "4px", paddingBottom: "4px"}}>
            <img src={icons[t.type]} style={{width: "20px"}} />
          </Box>,
          <Box key={`${t.path}-path`} style={{flexShrink: 0, paddingLeft: "16px", paddingTop: "4px", paddingBottom: "4px"}}>
            <Link onClick={this.onClickGitItem.bind(this, t)}>{t.path}</Link>
          </Box>,
          <Box key={`${t.path}-sha`} style={{flexShrink: 0, paddingLeft: "16px", paddingTop: "4px", paddingBottom: "4px"}}>
            Last commit message
          </Box>,
          <Box key={`${t.path}-size`} style={{flexShrink: 0, paddingLeft: "16px", paddingTop: "4px", paddingBottom: "4px"}}>
            24 hours ago
          </Box>,
        ]
      )
    });

    return (
      <Flex alignSelf="center" alignItems="center" justifyContent="center">
        <Box display="flex">
          <Box width={900}>
            <Box alignSelf="right" justifyContent="right">
              <Dropdown>
                <Dropdown.Button onClick={this.onClickAddGitHubRelease}>+ GitHub Release</Dropdown.Button>
              </Dropdown>
            </Box>
            <BorderBox>
              <Grid gridTemplateColumns="repeat(4, auto)" gridGap={3}>
                {rows.flat()}
              </Grid>
            </BorderBox>
          </Box>
        </Box>
      </Flex>
    );
  }
}

export default withRouter(Manifests);
