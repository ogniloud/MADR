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
        <div>
            <h2>Deck Details</h2>
            <nav>
                <ul>
                    <li>
                        <Link to={`/decks/${deck_id}/browse-cards/all-words`}>All Words</Link>
                    </li>
                    <li>
                        <Link to={`/decks/${deck_id}/browse-cards/the-hottest`}>The Hottest</Link>
                    </li>
                    <li>
                        <Link to={`/decks/${deck_id}/browse-cards/warm`}>Warm</Link>
                    </li>
                    <li>
                        <Link to={`/decks/${deck_id}/browse-cards/learned`}>Learned</Link>
                    </li>
                    <li>
                        <Link to={`/decks/${deck_id}/exercise/flashcards`}>Flashcards</Link>
                    </li>
                    <li>
                        <Link to={`/decks/${deck_id}/exercise/texts`}>Texts</Link>
                    </li>
                    <li>
                        <Link to={`/decks/${deck_id}/exercise/word-match`}>Word Match</Link>
                    </li>
                    <li>
                        <Link to={`/decks/${deck_id}/exercise/fill-gaps`}>Fill Gaps</Link>
                    </li>
                </ul>
            </nav>

            <Routes>
                {/* Browse Cards Routes */}
                <Route path="browse-cards/all-words" element={<AllWords />} />
                <Route path="browse-cards/the-hottest" element={<TheHottest />} />
                <Route path="browse-cards/warm" element={<Warm />} />
                <Route path="browse-cards/learned" element={<Learned />} />

                {/* Exercise Routes */}
                <Route path="exercise/flashcards" element={<Flashcards />} />
                <Route path="exercise/texts" element={<Texts />} />
                <Route path="exercise/word-match" element={<WordMatch />} />
                <Route path="exercise/fill-gaps" element={<FillGaps />} />
            </Routes>
        </div>
    );
};

export default DeckDetail;
