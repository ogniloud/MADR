import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';


// components
import Form from './components/User-Components/SignUp';
import SignIn from './components/User-Components/SignIn';
import MainPage from './components/Main-Page/MainPage';
import './components/Style-component/App.css';
import CreateDecks from "./components/Main-Page/Decks/CreateDecks";
import DecksPage from './components/Main-Page/Decks/DeckPage';
import DeckDetail from "./components/Main-Page/Decks/DeckDetail";
import AllWords from "./components/Main-Page/Decks/Browse Cards/AllWords";
import Flashcards from "./components/Main-Page/Decks/Exercise Cards/Flashcards";
import WordMatch from "./components/Main-Page/Decks/Exercise Cards/WordMatch";

function App() {
  return (
      <div className="App">
        <Router>
          <Routes>
            <Route path="/signup" element={<Form />} />
            <Route path="/" element={<Form />} />
            <Route path="/signin" element={<SignIn />} />
            <Route path="/mainpage" element={<MainPage />} />
            <Route path="/create-deck" element={<CreateDecks />} />
            <Route path="/decks" element={<DecksPage />} />

            {/* Specific routes for individual components in DeckDetail */}
            <Route path="/decks/:deck_id/browse-cards/all-words" element={<AllWords />} />
            <Route path="/decks/:deck_id/exercise/flashcards" element={<Flashcards />} />
            <Route path="/decks/:deck_id/exercise/word-match" element={<WordMatch />} />

            {/* DeckDetail route without a wildcard */}
            <Route path="/decks/:deck_id" element={<DeckDetail />} />
          </Routes>
        </Router>
      </div>
  );
}

const NotFound = () => (
    <div>
      <h1>404 - Not Found</h1>
      <p>The page you are looking for does not exist.</p>
    </div>
);

export default App;
