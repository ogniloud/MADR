import React, {useState} from 'react';
import './User-Components-Style/SignIn.css';
import {useNavigate} from 'react-router-dom';

function SignIn() {
  const [signInData, setSignInData] = useState({
    username: '',
    password: '',
  });

  const [error, setError] = useState(null); // State to handle sign-in errors
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
      const response = await fetch(`http://${process.env.REACT_APP_API_HOST}/api/signin`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(signInData),
      });

      if (response.ok) {
        const { authorization } = await response.json();
        localStorage.setItem('token', authorization);

        // Redirect to the main page
        navigate('/mainpage');
      } else {
        const errorData = await response.json(); // Parse the error response JSON
        setError(errorData.message || 'Sign-in failed'); // Set error message
      }
    } catch (error) {
      console.error('Error during sign-in:', error);
      setError('Error during sign-in. Please try again.'); // Set generic error message
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
              <label htmlFor="username">User Name</label>
              <input
                  type="text"
                  id="username"
                  name="username"
                  value={signInData.username}
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
                  placeholder="Enter your madr.org password"
                  required
              />
            </div>
            <button className="submit" type="submit">
              Login
            </button>
            {error && <div className="error-message">{error}</div>}
          </form>
        </div>
      </div>
  );
}

export default SignIn;