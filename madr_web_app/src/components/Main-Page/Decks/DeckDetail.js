import React from 'react';
import { useParams, Link, Route, Routes } from 'react-router-dom';
import AllWords from './Browse Cards/AllWords';
import TheHottest from './Browse Cards/TheHottest';
import Warm from './Browse Cards/Warm';
import Learned from './Browse Cards/Learned';
import Flashcards from './Exercise Cards/Flashcards';
import Texts from './Exercise Cards/Texts';
import WordMatch from './Exercise Cards/WordMatch';
import FillGaps from './Exercise Cards/FillGaps';
import './DeckDetails.css';

const DeckDetail = () => {
    const { deck_id } = useParams();

    return (
        <div className="deck-details-container">
            <h2 className="deck-details-title">Deck Details</h2>

            <div className="deck-details-section">
                <h2 className="title-Browse-Cards">Browse Cards</h2>
                <div className="deck-details-flashcard">
                    <Link to={`/decks/${deck_id}/browse-cards/all-words`} className="flashcard-link">All Words</Link>
                </div>
                <div className="deck-details-flashcard">
                    <Link to={`/decks/${deck_id}/browse-cards/the-hottest`} className="flashcard-link">The Hottest</Link>
                </div>
                <div className="deck-details-flashcard">
                    <Link to={`/decks/${deck_id}/browse-cards/warm`} className="flashcard-link">Warm</Link>
                </div>
                <div className="deck-details-flashcard">
                    <Link to={`/decks/${deck_id}/browse-cards/learned`} className="flashcard-link">Learned</Link>
                </div>
            </div>

            <div className="deck-details-section">
                <h2 className="title-Exercise">Exercise</h2>
                <div className="deck-details-flashcard">
                    <Link to={`/decks/${deck_id}/exercise/flashcards`} className="flashcard-link">Flashcards</Link>
                </div>
                <div className="deck-details-flashcard">
                    <Link to={`/decks/${deck_id}/exercise/texts`} className="flashcard-link">Texts</Link>
                </div>
                <div className="deck-details-flashcard">
                    <Link to={`/decks/${deck_id}/exercise/word-match`} className="flashcard-link">Word Match</Link>
                </div>
                <div className="deck-details-flashcard">
                    <Link to={`/decks/${deck_id}/exercise/fill-gaps`} className="flashcard-link">Fill Gaps</Link>
                </div>
            </div>

            <Routes>
                {/* Browse Cards Routes */}
                <Route path="/decks/:deck_id/browse-cards/all-words" element={<AllWords />} />
                <Route path="/decks/:deck_id/browse-cards/the-hottest" element={<TheHottest />} />
                <Route path="/decks/:deck_id/browse-cards/warm" element={<Warm />} />
                <Route path="/decks/:deck_id/browse-cards/learned" element={<Learned />} />

                {/* Exercise Routes */}
                <Route path="/decks/:deck_id/exercise/flashcards" element={<Flashcards />} />
                <Route path="/decks/:deck_id/exercise/texts" element={<Texts />} />
                <Route path="/decks/:deck_id/exercise/word-match" element={<WordMatch />} />
                <Route path="/decks/:deck_id/exercise/fill-gaps" element={<FillGaps />} />
            </Routes>
        </div>
    );
};

export default DeckDetail;
