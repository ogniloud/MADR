import React, {useEffect, useState} from 'react';
import {Link, Route, Routes, useNavigate, useParams} from 'react-router-dom';
import {jwtDecode} from 'jwt-decode';
import AllWords from './Browse Cards/AllWords';
import TheHottest from './Browse Cards/TheHottest';
import Warm from './Browse Cards/Warm';
import Learned from './Browse Cards/Learned';
import Flashcards from './Exercise Cards/Flashcards';
import Texts from './Exercise Cards/Texts';
import WordMatch from './Exercise Cards/WordMatch';
import FillGaps from './Exercise Cards/FillGaps';
import './Styles/DeckDetails.css';
import {checkIfShared, shareDeck} from "../APIs/apiFunctions_decks";

const DeckDetail = () => {
    const { deck_id } = useParams();
    const navigate = useNavigate();
    const [deckToDelete, setDeckToDelete] = useState(null);
    const decodedToken = jwtDecode(localStorage.getItem('token'));
    const userId = decodedToken.user_id;
    const [isShared, setIsShared] = useState(null)

    useEffect(() => {
        handleShare()
    }, [deck_id]);

    const handleShare = () => {
        checkIfShared(parseInt(deck_id), parseInt(userId), decodedToken).then((response) => {
            setIsShared(response.ok)
            console.log(response)
        })
    }

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

            <div className="">
                {isShared === true && (
                    <div className="deck-details-flashcard flashcard-link">
                        Shared!
                    </div>
                ) || isShared === false && (
                    <div onClick={() => {
                        shareDeck(parseInt(deck_id), parseInt(userId), decodedToken)
                        setIsShared(true)
                    }} className="deck-details-flashcard flashcard-link">
                        Share
                    </div>
                )}

            </div>
        </div>
    );
};

export default DeckDetail;
