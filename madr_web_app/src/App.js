import React from 'react';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import Form from './App Components/User-Components/SignUp';
import SignIn from './App Components/User-Components/SignIn';
import MainPage from './App Components/Page-Components/MainPage';
import './App Components/Page-Components/Style-components/App.css';
import CreateDecks from "./App Components/Page-Components/Decks/CreateDecks";
import DecksPage from './App Components/Page-Components/Decks/DeckPage';
import DeckDetail from "./App Components/Page-Components/Decks/DeckDetail";
import AllWords from "./App Components/Page-Components/Decks/Browse Cards/AllWords";
import Flashcards from "./App Components/Page-Components/Decks/Exercise Cards/Flashcards";
import WordMatch from "./App Components/Page-Components/Decks/Exercise Cards/WordMatch";
import SocialGroup from "./App Components/Page-Components/Social Site Components/social_group";
import FeedsPage from "./App Components/Page-Components/Feeds/FeedsPage";
import Learned from "./App Components/Page-Components/Decks/Browse Cards/Learned";
import TheHottest from "./App Components/Page-Components/Decks/Browse Cards/TheHottest";
import Warm from "./App Components/Page-Components/Decks/Browse Cards/Warm";
import FillGaps from "./App Components/Page-Components/Decks/Exercise Cards/FillGaps";
import Texts from "./App Components/Page-Components/Decks/Exercise Cards/Texts";
import SocialGroupDetail from "./App Components/Page-Components/Social Site Components/social_group_detail";
import Drawer  from "./App Components/Page-Components/Drawer-Components/Drawer";



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
            <Route path="/feed" element={<FeedsPage />} />
            {/* Specific routes for individual App Components in DeckDetail */}
            <Route path="/decks/:deck_id/browse-cards/all-words" element={<AllWords />} />
            <Route path="/decks/:deck_id/browse-cards/the-hottest" element={<TheHottest />} />
            <Route path="/decks/:deck_id/browse-cards/learned" element={<Learned />} />
            <Route path="/decks/:deck_id/browse-cards/warm" element={<Warm />} />
            <Route path="/decks/:deck_id/exercise/flashcards" element={<Flashcards />} />
            <Route path="/decks/:deck_id/exercise/word-match" element={<WordMatch />} />
            <Route path="/decks/:deck_id/exercise/fill-gaps" element={<FillGaps />} />
            <Route path="/decks/:deck_id/exercise/texts" element={<Texts />} />


            {/* DeckDetail route without a wildcard */}
            <Route path="/decks/:deck_id" element={<DeckDetail />} />

              {/*  Social sites and user's profile*/}
            <Route path = "/social_group" element={<SocialGroup/>} />
            <Route path="/social_group/:group_id" element={<SocialGroupDetail />} />

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
