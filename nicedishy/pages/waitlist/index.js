import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router'
import Layout from "../../components/layout";

export default function Page() {
  const router = useRouter();

  return (
    <>
      <div className="row">
        <div className="col-12">
          <h2>Waitlist update</h2>
          <p>Hey there, welcome. You are on the waitlist.</p>
          <p>Want to move up on the list? Complete the info below!</p>
          <div className="container">
            <form>
              <div className="row">
                <div className="col-8 offset-2">
                  <div className="form-group">
                    <label htmlFor="alreadyHaveDishy">Do you already have a Dishy?</label>
                    <select className="form-control" id="alreadyHaveDishy">
                      <option>Yep! Installed and working</option>
                      <option>No, I&apos;m still waiting</option>
                      <option>I have it, but it&apos;s not installed yet</option>
                    </select>
                  </div>
                </div>
              </div>
              <div className="row">
                <div className="col-8 offset-2">
                  <div className="form-group">
                    <label htmlFor="alreadyHaveDishy">How long have your been on the Starlink network?</label>
                    <select className="form-control" id="alreadyHaveDishy">
                      <option>Less than 3 months</option>
                      <option>3-6 months</option>
                      <option>Over 6 months (OG!)</option>
                    </select>
                  </div>
                </div>
              </div>
              <div className="row">
                <div className="col-8 offset-2">
                  <div className="form-group">
                    <label htmlFor="alreadyHaveDishy">Is Starlink your primary or backup connection?</label>
                    <select className="form-control" id="alreadyHaveDishy">
                      <option>Primary</option>
                      <option>Backup</option>
                    </select>
                  </div>
                </div>
              </div>
              <div className="row">
                <div className="col-8 offset-2">
                  <div className="form-group">
                    <label htmlFor="alreadyHaveDishy">Which of the following operating systems are you running at home and would be willing to install the NiceDishy agent on?</label>
                    <div className="form-check form-check-inline">
                      <input className="form-check-input" type="checkbox" id="inlineCheckbox1" value="option1" />
                      <label className="form-check-label" htmlFor="inlineCheckbox1">Windows</label>
                    </div>
                    <div className="form-check form-check-inline">
                      <input className="form-check-input" type="checkbox" id="inlineCheckbox2" value="option2" />
                      <label className="form-check-label" htmlFor="inlineCheckbox2">MacOS</label>
                    </div>
                    <div className="form-check form-check-inline">
                      <input className="form-check-input" type="checkbox" id="inlineCheckbox3" value="option3" />
                      <label className="form-check-label" htmlFor="inlineCheckbox3">Linux</label>
                    </div>
                  </div>
                </div>
              </div>
              <div className="row">
                <div className="col-8 offset-2">
                  <div className="form-group">
                    <label htmlFor="alreadyHaveDishy">Why do you want access to NiceDishy?</label>
                    <div className="form-check form-check-inline">
                      <input className="form-check-input" type="checkbox" id="inlineCheckbox1" value="option1" />
                      <label className="form-check-label" htmlFor="inlineCheckbox1">Monitoring my own dish</label>
                    </div>
                    <div className="form-check form-check-inline">
                      <input className="form-check-input" type="checkbox" id="inlineCheckbox2" value="option2" />
                      <label className="form-check-label" htmlFor="inlineCheckbox2">Compare my dish performance with others</label>
                    </div>
                    <div className="form-check form-check-inline">
                      <input className="form-check-input" type="checkbox" id="inlineCheckbox3" value="option3" />
                      <label className="form-check-label" htmlFor="inlineCheckbox3">Be a part of the Starlink community</label>
                    </div>
                  </div>
                </div>
              </div>
              <div className="row">
                <div className="col-8 offset-2 text-center">
                  <button type="submit" className="btn btn-primary">Save</button>
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
