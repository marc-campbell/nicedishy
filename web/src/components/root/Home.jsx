import * as React from "react";
import { withRouter } from "react-router-dom";
import { Utilities } from "../../utilities/utilities";

class Home extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: true,
      dishies: [],
    }
  }

  componentDidMount = async() => {
    try {   
      const res = await fetch(`${window.env.API_ENDPOINT}/dishies`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
        },
      });

      if (!res.ok) {
        return;
      }

      const data = await res.json();
      
      this.setState({
        isLoading: false,
        dishies: data.dishies,
      });
    } catch (err) {
      console.error(err);
    }
  }

  render() {
    if (this.state.isLoading) {
      return (
        <div>
          loading...
        </div>
      );
    }

    const dishies = this.state.dishies.map((dishy) => {
      return (
        <div key={dishy.id}>
          <h3>{dishy.name}</h3>
        </div>
      );
    });

    return (
      <div>
        <h1>Dishy (or Dishies if you are lucky)</h1>
        {dishies}
        <div>
          <h3>Add a new dishy</h3> 
        </div>
      </div>
    );
  }
}

export default withRouter(Home);
