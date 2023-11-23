import React, { useState } from 'react';
import { Link, Route, Routes } from 'react-router-dom';
import DecksPage from './Decks/DeckPage';
import Feed from './Feeds/FeedsPage';
import AllCards from './AllCards';
import UserProfile from './UserProfile';
import CreateDeck from './Decks/CreateDecks';
import './Styles/MainPage.css'


const MainPage = () => {
  // State to manage created decks
  const [createdDecks, setCreatedDecks] = useState([]);

  // Function to add a new deck
  const addDeck = (newDeck) => {
    setCreatedDecks([...createdDecks, newDeck]);
  };

  return (
    <div className="main-page">
      <nav className='upper-part'>
        <Link className='decks-bar' to="/decks">Decks</Link>
        <Link className='feed-bar' to="/feed">Feed</Link>
        <Link className='profile-bar' to="/profile">Profile</Link>
      </nav>

      <Routes>
        <Route path="decks" element={<DecksPage createdDecks={createdDecks} />} />
        <Route path="feed" element={<Feed />} />
        <Route path="profile" element={<UserProfile />} />
      </Routes>

      <div className='lower-part'>
        <Link className='create-deck' to="/create-deck">Create Deck</Link>
        <Link className='all-cards' to="/all-cards">All Cards</Link>
      </div>

      <Routes>
        <Route path="create-deck" element={<CreateDeck addDeck={addDeck} />} />
        <Route path="all-cards" element={<AllCards addDeck={addDeck} />} />
      </Routes>
    </div>
  );
};

export default MainPage;
