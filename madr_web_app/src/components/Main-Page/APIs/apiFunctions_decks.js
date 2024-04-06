import React, {useEffect, useRef, useState} from 'react';
import {Link, Route, Routes, useNavigate} from 'react-router-dom';
import {jwtDecode} from 'jwt-decode';
import axios from 'axios';

{/* API's for CreateDecks.js */}
export const createDeck = async (deckData, token) => {
    try {
        const response = await axios.put(
            'http://localhost:8080/api/flashcards/new_deck',
            deckData,
            {
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${token}`,
                },
            }
        );
        return response.data;
    } catch (error) {
        throw new Error('Failed to create deck. Please try again.');
    }
};

{/* API's for CreateDecks.js */}



{/* API's for DeckPage.js */}
export const fetchUserDecks = async (token) => {
    try {
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
            return decks;
        } else {
            throw new Error('Failed to load user decks');
        }
    } catch (error) {
        throw new Error('Error loading user decks');
    }
};

export const deleteDeck = async (deckId, userId, token) => {
    try {
        const response = await fetch('http://localhost:8080/api/flashcards/delete_deck', {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
                deck_id: deckId,
                user_id: userId,
            }),
        });

        if (!response.ok) {
            throw new Error('Failed to delete deck');
        }
    } catch (error) {
        throw new Error('Error deleting deck');
    }
};

{/* API's for DeckPage.js */}




