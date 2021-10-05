import * as React from "react";

export default class Avatar extends React.Component {
  render() {
    return (
      <div className="avatar-wrapper" style={{ backgroundImage: `url(${this.props.imageUrl || ""})` }}></div>
    );
  }
}
