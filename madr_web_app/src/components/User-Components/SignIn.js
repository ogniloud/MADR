import React, { useState } from 'react';
import './User-Components-Style/SignIn.css'
function SignIn() {
  const [signInData, setSignInData] = useState({
    email: '',
    password: '',
  });

  const handleSignInChange = (e) => {
    const { name, value } = e.target;
    setSignInData({
      ...signInData,
      [name]: value,
    });
  };

  const handleSignInSubmit = (e) => {
    e.preventDefault();
    // Handle sign-in form submission, e.g., sending data to a server for authentication.
  };

  return (
    <div className="container">
      <div className="left-side">
      </div>
      <div className="right-side">
        <p className="title">Sign in to your MADR account</p>
        <p className="subtitle">Let's start learning</p>
        <form onSubmit={handleSignInSubmit}>
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              name="email"
              value={signInData.email}
              onChange={handleSignInChange}
              placeholder=""
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              name="password"
              value={signInData.password}
              onChange={handleSignInChange}
              placeholder=""
              required
            />
          </div>
          <button className="submit" type="submit">
           Login
          </button>
        </form>
      </div>
    </div>
  );
}

export default SignIn;
