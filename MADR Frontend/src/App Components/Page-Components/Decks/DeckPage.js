import React, { useEffect, useState } from 'react';
import { Link, Route, Routes } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';
import {fetchUserDecks, deleteDeck} from "../API-Components/apiFunctions_decks";
import DeckDetail from './DeckDetail';
import './Styles/DeckPage.css';

const DecksPage = () => {
    const [createdDecks, setCreatedDecks] = useState([]);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            fetchDecks(token);
        }
    }, []);

    const fetchDecks = async (token) => {
        try {
            const decks = await fetchUserDecks(token);
            setCreatedDecks(decks);
        } catch (error) {
            console.error(error.message);
        }
    };

    const handleDeleteDeck = async (deckId) => {
        try {
            const token = localStorage.getItem('token');
            const userId = getUserId(token);
            await deleteDeck(deckId, userId, token);
            setCreatedDecks(prevDecks => prevDecks.filter(deck => deck.deck_id !== deckId));
            alert('Deck deleted successfully');
        } catch (error) {
            console.error(error.message);
            alert('Failed to delete deck');
        }
    };

    const getUserId = (token) => {
        if (token) {
            const decodedToken = jwtDecode(token);
            return decodedToken.user_id;
        }
        return null;
    };

    const handleDeleteConfirmation = (deckId, e) => {
        e.preventDefault();
        if (window.confirm('Are you sure you want to delete this deck?')) {
            handleDeleteDeck(deckId);
        }
    };

    const handleContextMenu = (e, deckId) => {
        e.preventDefault();
        handleDeleteConfirmation(deckId, e);
    };

    return (
        <div>
            <h2 className="title-deck">All Decks</h2>
            <div className="deck-container">
                {createdDecks.length > 0 &&
                    createdDecks.map((deck) => (
                        <div key={deck.deck_id} className="deck-card" onContextMenu={(e) => handleContextMenu(e, deck.deck_id)}>
                            <Link to={`/decks/${deck.deck_id}`}>
                                <span>{deck.name}</span>
                            </Link>
                        </div>
                    ))}
            </div>

            <Routes>
                <Route path="/:deckId/*" element={<DeckDetail />} />
            </Routes>
        </div>
    );
};

export default DecksPage;
