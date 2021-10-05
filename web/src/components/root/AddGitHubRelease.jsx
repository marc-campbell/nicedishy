import { Box, Button, ButtonPrimary, Dropdown, FormGroup, Flex, Heading, TextInput } from "@primer/components";
import * as React from "react";
import classNames from "classnames";
import { withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

class AddGitHubRelease extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      repo: "",
      isValid: true,
      releases: [],
      selectedRelease: "",
      selectedFilenames: [],
      directory: "",
    };
  }

  onChangeTargetDir = (ev) => {
    this.setState({
      directory: ev.target.value,
    });
  }

  onChangeRepo = async (ev) => {
    this.setState({
      repo: ev.target.value,
      releases: [],
      isValid: false,
      selectedRelease: "",
      selectedFilenames: [],
    });

    const parts = ev.target.value.split("/");
    if (parts.length < 2) {
      return;
    }
    if (parts[1].length === 0) {
      return;
    }

    // try the repo input
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/github/check-repo`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
        body: JSON.stringify({
          name: ev.target.value,
        }),
      })

      if (!res.ok) {
        console.log("error");
        return;
      }

      const data = await res.json();
      if (!data.isValid) {
        this.setState({
          isValid: false,
        });
        return;
      }

      const selectedRelease = data.releases.length > 0 ? data.releases[0].name : "";

      this.setState({
        isValid: true,
        releases: data.releases,
        selectedRelease: selectedRelease,
      })

    } catch (err) {
      console.error(err);
    }
  }

  onSelectRelease = async(release) => {
    this.setState({
      selectedRelease: release,
    })
  }

  onSelectFilename = async(filename, ev) => {
    const selectedFilenames = this.state.selectedFilenames;

    if (ev.target.checked) {
      selectedFilenames.push(filename);
    } else {
      const i = selectedFilenames.indexOf(filename);
      if (i > -1) {
        selectedFilenames.splice(i, 1);
      }
    }
    this.setState({
      selectedFilenames,
    });

  }

  onClickSave = async() => {
    try {
      const res = await fetch(`${window.env.API_ENDPOINT}/context/${this.props.match.params.owner}/${this.props.match.params.context}/link/github-release`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
        body: JSON.stringify({
          selectedFilenames: this.state.selectedFilenames,
          repo: this.state.repo,
          directory: this.state.directory,
        })
      });

      if (res.status === 401) {
        this.props.history.replace(`/login`);
        return;
      }

      if (!res.ok) {
        console.log("error");
        return;
      }

      const data = await res.json();
      console.log(data);

    } catch(err) {
      console.error(err);
    }
  }

  render() {
    const releases = this.state.releases.map((release) => {
      return (
        <Dropdown.Item key={release.name} onClick={this.onSelectRelease.bind(this, release.name)}>{release.name}</Dropdown.Item>
      );
    });

    let files = [];
    if (this.state.selectedRelease !== "") {
      const release = this.state.releases.find((release) => {
        return (release.name === this.state.selectedRelease);
      });

      files = release.assets.map((asset) => {
        return (
          <FormGroup.Label key={asset.filename} htmlFor={asset.filename}>
            <input type="checkbox" id={asset.filename} style={{marginRight: "7px"}} onClick={this.onSelectFilename.bind(this, asset.filename)} />{asset.filename}
          </FormGroup.Label>
        );
      })
    }


    return (
      <Flex alignSelf="center" alignItems="center" justifyContent="center">
        <Box display="flex">
          <Box width={900}>
            <p>
              This will add one or more files from a GitHub Release to this repo.
              CentralContext will watch the GitHub releases and send Pull Requests to your repo when there
              are new versions available.
            </p>
            <FormGroup>
              <FormGroup.Label htmlFor="target-dir">Target Directory</FormGroup.Label>
              <TextInput id="target-dir" placeholder="/sealed-secrets" onChange={this.onChangeTargetDir} value={this.state.directory} />
            </FormGroup>
            <FormGroup>
              <FormGroup.Label htmlFor="github-repo">GitHub Repo</FormGroup.Label>
              <TextInput id="github-repo" placeholder="bitnami-labs/sealed-secrets" onChange={this.onChangeRepo} value={this.state.repo} />
            </FormGroup>
            <FormGroup>
              <FormGroup.Label htmlFor="file-selector">File Selector</FormGroup.Label>
                The following files were found in the
                <Dropdown>
                  <Dropdown.Button>{this.state.selectedRelease}</Dropdown.Button>
                  <Dropdown.Menu direction="sw">
                    {releases}
                  </Dropdown.Menu>
                </Dropdown>
                release. To add a link, select the files and CentralContext will watch
                for new releases containing these filenames and will update them.
              {files}
            </FormGroup>
            <ButtonPrimary disabled={this.state.selectedFilenames.length === 0} onClick={this.onClickSave}>Save & Create Link</ButtonPrimary>
          </Box>
        </Box>
      </Flex>
    );
  }
}

export default withRouter(AddGitHubRelease);
