import * as React from "react";
import { Route, Switch, withRouter } from "react-router-dom";
import ClientAuthorize from "../ClientAuthorize";
import { NavBar } from "../shared/NavBar";
import NotFound from "../static/NotFound";
import Context from "./Context";
import Home from "./Home";
import NewContext from "./NewContext";
import Settings from "./Settings";

class LoggedInRoot extends React.Component {
  render() {
    return (
      <div>
        <NavBar {...this.props} />
        <Switch>
          <Route exact path="/" component={Home} />
          <Route exact path="/client/authorize" render={props => <ClientAuthorize {...props} />} />
          <Route exact path="/new" render={props => <NewContext {...this.props} />} />
          <Route path="/settings" render={props => <Settings {...this.props} />} />
          <Route path="/:owner/:context" render={props => <Context {...this.props} />} />
          <Route component={NotFound} />
        </Switch>
      </div>
    );
  }
}

export default withRouter(LoggedInRoot);
