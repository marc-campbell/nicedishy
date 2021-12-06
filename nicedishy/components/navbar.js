import * as React from "react";
import { Utilities } from "../utils/utilities";
import Link from 'next/link'

export default function Navbar() {
  const handleLogout = () => {
    Utilities.logoutUser();
  }

  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light">
      <div className="container-fluid">
        <Link className="navbar-brand" href="/">
          <img src="/images/icon-white.png" alt="" width="30" height="24" />
        </Link>
        <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className="collapse navbar-collapse" id="navbarSupportedContent">
          <ul className="navbar-nav me-auto mb-2 mb-lg-0">
            <li className="nav-item">
              <a className="nav-link" href="#" onClick={handleLogout}>Logout</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
}
