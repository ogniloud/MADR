import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './User-Components-Style/SignUp.css';

function Form() {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
  });

  const navigate = useNavigate(); // Initialize navigate

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

 const handleSubmit = async (e) => {
  e.preventDefault();

  try {
    const response = await fetch('http://localhost:8080/api/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    });

    console.log('Signup Response:', response);

    if (response.ok) {
      console.log('Registration Successful');
      navigate('/signin'); // Redirect to the SignIn page
    } else {
      const data = await response.json().catch(() => null);
      console.error('Registration Failed:', data);

      // If data is null, it means the response body is not a valid JSON
      if (data === null) {
        console.error('Invalid JSON in response body');
      }
    }
  } catch (error) {
    console.error('Error during registration', error);
  }
};

  

  return (
    <div className="container">
      <div className="left-side">
        <p className="describtion">
          MADR is a modern language learning tool. It provides various vocabulary enrichment methods that are scientifically proven to be effective and engaging. Find all the necessary resources in the web application.
        </p>
        <p className="madr_org">madr.org</p>
        <p className="one_tab"> Just one browser tab</p>
      </div>

      <div className="right-side">
        <p className="title">Create your MADR account</p>
        <p className="subtitle">Let’s make learning more effective</p>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="name">User Name</label>
            <input type="text" id="name" name="username" value={formData.username} onChange={handleInputChange} placeholder="pick a suitable user name" required />
          </div>
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input type="email" id="email" name="email" value={formData.email} onChange={handleInputChange} placeholder="Enter your email" required />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input type="password" id="password" name="password" value={formData.password} onChange={handleInputChange} placeholder="Password should be more that 8 digit " required />
          </div>
          <button className="submit" type="submit">
            Create Account
          </button>
        </form>
        <p>
          Already have an account? <Link to="/signin">Sign in</Link>
        </p>
      </div>
    </div>
  );
}

export default Form;
