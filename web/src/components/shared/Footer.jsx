import * as React from "react";
import { Link, withRouter } from "react-router-dom";
import { getBuildVersion } from "../../utilities/utilities";
import { Layout } from 'antd';
const { Footer } = Layout;

import "../../scss/components/footer.scss";

export class NiceDishyFooter extends React.Component {

  render() {
    return (
      <Footer style={{ textAlign: 'center' }}>NiceDishy &copy; 2021<br /></Footer>
    );
  }
}

export default withRouter(NiceDishyFooter);
