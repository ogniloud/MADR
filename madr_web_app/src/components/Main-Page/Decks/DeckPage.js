// DeckPage.js
import React, { useEffect, useState } from 'react';
import { Link, Route, Routes } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';
import DeckDetail from './DeckDetail';
import './DeckPage.css';

const DecksPage = () => {
    const [createdDecks, setCreatedDecks] = useState([]);

    useEffect(() => {
        const fetchUserDecks = async () => {
            try {
                const token = localStorage.getItem('token');
                if (token) {
                    const decodedToken = jwtDecode(token);
                    let retries = 3;

                    while (retries > 0) {
                        const response = await fetch('http://localhost:8080/api/flashcards/load', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify({ user_id: decodedToken.user_id }),
                        });

                        if (response.ok) {
                            const { decks } = await response.json();
                            setCreatedDecks(decks);
                            break;
                        } else {
                            console.error('Failed to load user decks:', response.statusText);

                            // Retry after a delay
                            await new Promise(resolve => setTimeout(resolve, 1000));
                            retries--;
                        }
                    }
                }
            } catch (error) {
                console.error('Error loading user decks:', error);
            }
        };
        fetchUserDecks();
    }, []);

    const handleDeleteDeck = async (deckId) => {
        try {
            const response = await fetch('http://localhost:8080/api/flashcards/delete_deck', {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    deck_id: deckId,
                    user_id: getUserId(),
                }),
            });

            if (response.ok) {
                // Remove the deleted deck from the list
                setCreatedDecks(prevDecks => prevDecks.filter(deck => deck.deck_id !== deckId));
                alert('Deck deleted successfully');
            } else {
                console.error('Failed to delete deck:', response.statusText);
                alert('Failed to delete deck');
            }
        } catch (error) {
            console.error('Error deleting deck:', error);
            alert('Error deleting deck');
        }
    };

    const getUserId = () => {
        const token = localStorage.getItem('token');
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
