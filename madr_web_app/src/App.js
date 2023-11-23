import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Form from './components/User-Components/SignUp';
import SignIn from './components/User-Components/SignIn';
import MainPage from './components/Main-Page/MainPage';
import './Style-component/App.css';

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/signup" element={<Form />} />
          <Route path="/" element={<Form />} />
          <Route path="/signin" element={<SignIn />} />
          <Route path="/mainpage" element={<MainPage />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
