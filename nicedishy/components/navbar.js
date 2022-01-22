import * as React from "react";
import { Utilities } from "../utils/utilities";
import { useRouter } from 'next/router'
import Image from 'next/image'

export default function Navbar() {
  const router = useRouter();

  const handleLogout = () => {
    Utilities.logoutUser();
  }

  const handleLogoClick = (ev) => {
    ev.preventDefault();
    if (Utilities.isLoggedIn()) {
      if (Utilities.isWaitlisted()) {
        router.push('/waitlist');
      } else {
        router.push('/dishies');
      }
      return;
    }
    router.push("/");
  }

  return (
      <nav className="navbar navbar-default">
        <div className="container-fluid">
          <a className="navbar-brand" href="#" onClick={handleLogoClick}>
            <Image src="/images/nicedishy-logo.png" alt="" height="48px" width="200px" />
          </a>
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>

          <div className="nav navbar-nav navbar-right">
            {
              Utilities.isLoggedIn() ? null : <div className="btn btn-link hidden" style={{textAlign: "right"}} onClick={handleLogout}>Logout</div>
            }
          </div>
        </div>
      </nav>
  );
}
