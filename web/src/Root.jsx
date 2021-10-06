import { hot } from "react-hot-loader/root";
import React, { Component } from "react";
import { createBrowserHistory } from "history";
import { Switch, Route, Redirect, Router } from "react-router-dom";
import { Helmet } from "react-helmet";
import UnsupportedBrowser from "./components/static/UnsupportedBrowser";
import { Utilities } from "./utilities/utilities";
import Login from "./components/Login";
import LoginCallback from "./components/LoginCallback";
import LoggedInRoot from "./components/root/LoggedInRoot";
import {ThemeProvider} from '@primer/components'

import Footer from "./components/shared/Footer";
import NavBar from "./components/shared/NavBar";

import "./scss/index.scss";
let history = createBrowserHistory();

class ProtectedRoute extends Component {
  render() {
    const redirectURL = `/login?next=${this.props.location.pathname}${this.props.location.search}`;

    return (
      <Route path={this.props.path} render={(innerProps) => {
        if (Utilities.isLoggedIn()) {
          if (this.props.component) {
            return <this.props.component {...innerProps} />;
          }
          return this.props.render(innerProps);
        }
        return <Redirect to={redirectURL} />;
      }} />
    );
  }
}

class Root extends Component {
  state = {
    themeState: {
      navbarLogo: null,
    },
  };
  /**
   * Sets the Theme State for the whole application
   * @param {Object} newThemeState - Object to set for new theme state
   * @param {Function} callback - callback to run like in setState()'s callback
   */
  setThemeState = (newThemeState, callback) => {
    this.setState({
      themeState: { ...newThemeState }
    }, callback);
  }

  /**
   * Gets the current theme state of the app
   * @return {Object}
   */
  getThemeState = () => {
    return this.state.themeState;
  }

  /**
   * Clears the current theme state to nothing
   */
  clearThemeState = () => {
    /**
     * Reference object to a blank theme state
     */
    const EMPTY_THEME_STATE = {
      navbarLogo: null,
    };

    this.setState({
      themeState: { ...EMPTY_THEME_STATE }
    });
  }


  render() {
    return (
      <div className="flex-column flex1">
        <Helmet>
          <meta httpEquiv="Cache-Control" content="no-cache, no-store, must-revalidate" />
          <meta httpEquiv="Pragma" content="no-cache" />
          <meta httpEquiv="Expires" content="0" />
        </Helmet>
        <ThemeProvider>
          <Router history={history}>
            <Switch>
              <Route exact path="/login" render={props => <Login {...props} />} />
              <Route exact path="/login/callback" render={props => <LoginCallback {...props} />} />
              <Route path="/unsupported" component={UnsupportedBrowser} />
              <Route path="/" render={props =>
                Utilities.isLoggedIn() ? <LoggedInRoot {...props} /> : <Redirect to={`/login?next=${props.location.pathname}`} />}
              />
            </Switch>
            <Footer />
          </Router>
        </ThemeProvider>
      </div>
    );
  }
}
export default hot(Root);
