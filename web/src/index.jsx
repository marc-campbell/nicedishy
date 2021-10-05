import * as React from "react";
import * as ReactDOM from "react-dom";
import CustomErrorBoundary from "./components/shared/ErrorBoundary";
import Root from "./Root";

ReactDOM.render((
  <CustomErrorBoundary>
    <Root />
  </CustomErrorBoundary>
), document.getElementById("app"));
