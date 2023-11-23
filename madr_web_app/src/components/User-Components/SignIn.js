import React, { useState } from 'react';
import './User-Components-Style/SignIn.css';
import { useNavigate } from 'react-router-dom';

function SignIn() {
  const [signInData, setSignInData] = useState({
    email: '',
    password: '',
  });

  const navigate = useNavigate(); // Initialize navigate

  const handleSignInChange = (e) => {
    const { name, value } = e.target;
    setSignInData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSignInSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/api/signin', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(signInData),
      });

      if (response.ok) {
        // Handle successful sign-in
        console.log('Sign-in successful');
        navigate('/mainpage'); // Redirect to the main page
      } else {
        // Handle unsuccessful sign-in, e.g., display an error message
        console.error('Sign-in failed');
      }
    } catch (error) {
      console.error('Error during sign-in:', error);
    }
  };

  return (
    <div className="container">
      <div className="left-side"></div>
      <div className="right-side">
        <p className="title">Sign in to your MADR account</p>
        <p className="subtitle">Let's start learning</p>
        <form onSubmit={handleSignInSubmit}>
          <div className="form-group">
            <label htmlFor="email">User Name</label>
            <input
              type="name"
              id="name"
              name="name"
              value={signInData.name}
              onChange={handleSignInChange}
              placeholder="Enter your madr.org username"
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
              placeholder="Enter your madr.org passowrd"
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
