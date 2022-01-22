import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router'
import Layout from "../../components/layout";
import { Utilities } from '../../utils/utilities';

export default function Page() {
  const router = useRouter();
  const [alreadyHaveDishy, setAlreadyHaveDishy] = useState("yes");
  const [howLongWithDishy, setHowLongWithDishy] = useState("less-than-3-months");
  const [primaryOrBackup, setPrimaryOrBackup] = useState("primary");
  const [operatingSystems, setOperatingSystems] = useState([]);
  const [whyAccess, setWhyAccess] = useState([]);
  const [isSaved, setIsSaved] = useState(false);
  const [saveError, setSaveError] = useState("");

  useEffect(() => {
    async function fetchData() {
      if (!Utilities.getToken()) {
        router.push('/signup');
        return;
      }
      // validate that the token is still valid
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/whoami`, {
          method: 'GET',
          headers: {
          "Content-Type": "application/json",
          "Authorization": Utilities.getToken(),
          },
        });

        if (res.status === 401) {
          router.put("/signup");
          return;
        }
      } catch(err) {
        console.log(err);
      }
    }
    fetchData();;
  }, []);

  const handleSaveClick = async (ev) => {
    ev.preventDefault();

    setSaveError("");
    const payload = {
      alreadyHaveDishy: alreadyHaveDishy,
      howLongWithDishy: howLongWithDishy,
      primaryOrBackup: primaryOrBackup,
      operatingSystems: operatingSystems,
      whyAccess: whyAccess,
    };

    const uri = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/api/v1/waitlist`;
    const response = await fetch(uri, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        "Authorization": Utilities.getToken(),
      },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      setSaveError("oh no, unable to save. try again?");
      return "";
    }

    setIsSaved(true);
    const body = await response.json();
    return body;
  }

  const handleChangeDropdown = (name, ev) => {
    switch (name) {
      case 'alreadyHaveDishy':
        setAlreadyHaveDishy(ev.target.value);
        break;
      case 'howLongWithDishy':
        setHowLongWithDishy(ev.target.value);
        break;
      case 'primaryOrBackup':
        setPrimaryOrBackup(ev.target.value);
        break;
    }
  }

  const handleChangeOperatingSystem = (os, ev) => {
    if (ev.target.checked) {
      setOperatingSystems([...operatingSystems, os]);
    } else {
      setOperatingSystems(operatingSystems.filter(o => o !== os));
    }
  }

  const handleChangeWhyAccess = (why, ev) => {
    if  (ev.target.checked) {
      setWhyAccess([...whyAccess, why]);
    } else {
      setWhyAccess(whyAccess.filter(w => w !== why));
    }
  }

  return (
    <>
      <div className="row">
        <div className="col-12">
          <h2>Waitlist update</h2>
          <p>Hey there, welcome. You are on the waitlist.</p>
          <p>Want to move up on the list? Complete the info below!</p>
          <div className="row mx-4">
            <form>
              <div className="row">
                <div className="col-12">
                  <div className="form-group">
                    <label htmlFor="alreadyHaveDishy"><strong>Do you already have a Dishy?</strong></label>
                    <select className="form-control" id="alreadyHaveDishy" onChange={handleChangeDropdown.bind(this, "alreadyHaveDishy")}>
                      <option value="yes">Yep! Installed and working</option>
                      <option value="no">No, I&apos;m still waiting</option>
                      <option value="not-installed">I have it, but it&apos;s not installed yet</option>
                    </select>
                  </div>
                </div>
              </div>
              <div className="row py-4">
                <div className="col-12">
                  <div className="form-group">
                    <label htmlFor="alreadyHaveDishy"><strong>How long have your been on the Starlink network?</strong></label>
                    <select className="form-control" id="howLongWithDishy" onChange={handleChangeDropdown.bind(this, "howLongWithDishy")}>
                      <option value="less-than-three-months">Less than 3 months</option>
                      <option value="three-six-months">3-6 months</option>
                      <option value="over-six-months">Over 6 months (OG!)</option>
                    </select>
                  </div>
                </div>
              </div>
              <div className="row py-4">
                <div className="col-12">
                  <div className="form-group">
                    <label htmlFor="primaryOrBackup"><strong>Is Starlink your primary or backup connection?</strong></label>
                    <select className="form-control" id="primaryOrBackup" onChange={handleChangeDropdown.bind(this, "primaryOrBackup")}>
                      <option value="primary">Primary</option>
                      <option value="backup">Backup</option>
                    </select>
                  </div>
                </div>
              </div>
              <div className="row py-4">
                <div className="col-12">
                  <div className="form-group">
                    <label htmlFor="operatingSystems"><strong>Which of the following operating systems are you running at home and would be willing to install the NiceDishy agent on?</strong></label>
                    <div className="form-check">
                      <input className="form-check-input" type="checkbox" id="operatingSystemsWindows" value="windows" onChange={handleChangeOperatingSystem.bind(this, "windows")} />
                      <label className="form-check-label" htmlFor="operatingSystemsWindows">Windows</label>
                    </div>
                    <div className="form-check">
                      <input className="form-check-input" type="checkbox" id="opeatingSystemsMacos" value="macos" onChange={handleChangeOperatingSystem.bind(this, "macos")} />
                      <label className="form-check-label" htmlFor="inlineCopeatingSystemsMacosheckbox2">MacOS</label>
                    </div>
                    <div className="form-check">
                      <input className="form-check-input" type="checkbox" id="operatingSystemsLinux" value="linux" onChange={handleChangeOperatingSystem.bind(this, "linux")} />
                      <label className="form-check-label" htmlFor="operatingSystemsLinux">Linux</label>
                    </div>
                  </div>
                </div>
              </div>
              <div className="row py-4">
                <div className="col-12">
                  <div className="form-group">
                    <label htmlFor="whyAccess"><strong>Why do you want access to NiceDishy?</strong></label>
                    <div className="form-check">
                      <input className="form-check-input" type="checkbox" id="whyMonitoring" value="monitoring" onChange={handleChangeWhyAccess.bind(this, "monitoring")}  />
                      <label className="form-check-label" htmlFor="whyMonitoring">Monitoring my own dish</label>
                    </div>
                    <div className="form-check">
                      <input className="form-check-input" type="checkbox" id="whyCompare" value="compare" onChange={handleChangeWhyAccess.bind(this, "compare")} />
                      <label className="form-check-label" htmlFor="whyCompare">Compare my dish performance with others</label>
                    </div>
                    <div className="form-check">
                      <input className="form-check-input" type="checkbox" id="whyCommunity" value="community" onChange={handleChangeWhyAccess.bind(this, "community")} />
                      <label className="form-check-label" htmlFor="whyCommunity">Be a part of the Starlink community</label>
                    </div>
                  </div>
                </div>
              </div>
              <div className="row">
                <div className={`alert alert-success fade ${isSaved ? "show" : "hide"}`} role="alert">
                  We got it! Your responses have been saved. Stay tuned for the next steps.
                </div>
                <div className={`alert alert-danger fade ${saveError ? "show" : "hide"}`} role="alert">
                  {saveError}
                </div>
                <div className="col-8 offset-2 text-center">
                  <button className="btn btn-primary" onClick={handleSaveClick}>Save</button>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </>
  );
}

Page.getLayout = function getLayout(page) {
  return (
    <Layout>
      {page}
    </Layout>
  );
}
