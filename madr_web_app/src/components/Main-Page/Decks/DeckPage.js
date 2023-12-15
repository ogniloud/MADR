// DeckPage.js
import React, { useEffect, useState } from 'react';
import { Link, Route, Routes } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';
import DeckDetail from './DeckDetail';
import './DeckPage.css';

const Decks = () => {
    const [createdDecks, setCreatedDecks] = useState([]);

    useEffect(() => {
        const fetchUserDecks = async () => {
            try {
                const token = localStorage.getItem('token');
                if (token) {
                    const decodedToken = jwtDecode(token);
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
                    } else {
                        console.error('Failed to load user decks:', response.statusText);
                    }
                }
            } catch (error) {
                console.error('Error loading user decks:', error);
            }
        };

        fetchUserDecks();
    }, []);

    return (
        <div>
            <h2>All Decks</h2>
            <div className="deck-container">
                {/* Display all decks in a grid view */}
                {createdDecks.length > 0 &&
                    createdDecks.map((deck) => (
                        <div key={deck.deck_id} className="deck-card">
                            <Link to={`/decks/${deck.deck_id}`}>
                                <span>{deck.name}</span>
                            </Link>
                        </div>
                    ))}
            </div>

            {/* Use the Routes component to define routes */}
            <Routes>
                {/* Add a route for deck details */}
                <Route path="/:deckId/*" element={<DeckDetail />} />
            </Routes>
        </div>
    );
};

export default Decks;
