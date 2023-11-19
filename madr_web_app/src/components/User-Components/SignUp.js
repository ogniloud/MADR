import React, { useState } from 'react';
import { Link } from 'react-router-dom'; 
import './User-Components-Style/SignUp.css'
function Form() {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    password: '',
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async(e) => {
    e.preventDefault();
    
    try{
      const response = await fetch('http://localhost:8080/api/signup',{
        method:'POST',
        headers:{
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (response.ok){
        console.log('Registration Successful')
      } else{
        console.error('Registration Faild')
      }
    } catch(error){
      console.error('Error during registration', error);
    }

  };

  return (
    <div className="container">
      <div className="left-side">
        <p className='describtion'>
          MADR is a modern language learning tool. It provides various vocabulary enrichment methods that are scientifically proven to be effective and engaging. Find all the necessary resources in the web application.
        </p>
        <p className='madr_org'>madr.org</p>
        <p className='one_tab'> Just one browser tab</p>
      </div>

      <div className="right-side">
        <p className='title'>Create your MADR account</p>
        <p className='subtitle'>Let’s make learning more effective</p>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
           <label htmlFor="name">Name</label>
            <input type="text" id="name" name="name" value={formData.name} onChange={handleInputChange} placeholder="" required />
          </div>
          <div className="form-group">
            <label htmlFor="email">Email</label>
            <input type="email" id="email" name="email" value={formData.email} onChange={handleInputChange} placeholder="" required />
          </div>
          <div className="form-group">
            <label htmlFor="password">Password</label>
            <input type="password" id="password" name="password" value={formData.password} onChange={handleInputChange} placeholder="" required />
          </div>
          <button className="submit" type="submit">Create Account</button>
        </form>
        <p>Already have an account? <Link to='/signin'>Sign in</Link></p>
      </div>
    </div>
  );
}

export default Form;
