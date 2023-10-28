
import React from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Form from './components/Form'
import SignIn from './components/SignIn'
import './Style-component/App.css'
import './Style-component/SignUp.css'




function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/signup" element={<Form />} />
          <Route path="/signin" element={<SignIn />} />
          <Route path="/" element={<Form />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
