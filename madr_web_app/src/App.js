import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Form from './components/User-Components/SignUp';
import SignIn from './components/User-Components/SignIn';
import MainPage from './components/Main-Page/MainPage';
import './components/Style-component/App.css';
import CreateDecks from "./components/Main-Page/Decks/CreateDecks";
import DecksPage from './components/Main-Page/Decks/DeckPage';
import DeckDetail from "./components/Main-Page/Decks/DeckDetail";


function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/signup" element={<Form />} />
          <Route path="/" element={<Form />} />
          <Route path="/signin" element={<SignIn />} />
          <Route path="/mainpage" element={<MainPage />} />
          <Route path="/create-deck" element={<CreateDecks/>} />
          <Route path="/decks" element={<DecksPage />} />
          <Route path="/decks/:deck_id/*" element={<DeckDetail />} />

        </Routes>
      </Router>
    </div>
  );
}

export default App;
