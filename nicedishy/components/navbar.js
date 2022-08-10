import * as React from "react";
import {useState, useEffect} from 'react'
import { Utilities } from "../utils/utilities";
import { useRouter } from 'next/router'
import Image from 'next/image'
import cookies from 'next-cookies';
import { loadSession } from "../lib/session";
import Link from 'next/link';

export default function Navbar({isLoggedIn}) {
  const router = useRouter();

  const handleLogout = () => {
    Utilities.logoutUser();
  }

  const handleLogoClick = (ev) => {
    ev.preventDefault();
    if (isLoggedIn) {
      router.push('/dishies');
      return;
    }
    router.push("/");
  }

  return (
      <nav className="navbar navbar-default">
        <div className="container-fluid">
          <Link href="#" passHref>
            <a className="navbar-brand" onClick={handleLogoClick}>
              <Image src="/images/nicedishy-logo.png" alt="" height="48px" width="200px" />
            </a>
          </Link>
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>

          <div className="nav navbar-nav navbar-right">
            {
              isLoggedIn ? <Link href="/logout" onCLick={handleLogout} passHref><span className="btn btn-link hidden" style={{textAlign: "right"}} onClick={handleLogout}>Logout</span></Link> : null
            }
          </div>
        </div>
      </nav>
  );
}
