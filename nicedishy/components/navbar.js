import * as React from "react";
import { Utilities } from "../utils/utilities";
import { useRouter } from 'next/router'

export default function Navbar() {
  const router = useRouter();

  const handleLogout = () => {
    Utilities.logoutUser();
  }

  const handleLogoClick = () => {
    if (Utilities.isLoggedIn()) {
      router.push('/dishies');
      return;
    }
    router.push("/");
  }

  return (
      <nav className="navbar navbar-default">
        <div className="container-fluid">
          <a className="navbar-brand" href="#" onClick={handleLogoClick}>
            <img src="/images/nicedishy-logo.png" alt="" height="48" />
          </a>
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>

          <div className="nav navbar-nav navbar-right">
            <div className="btn btn-link" style={{textAlign: "right"}} onClick={handleLogout}>Logout</div>
          </div>
        </div>
      </nav>
  );
}
