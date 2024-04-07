import React, {useEffect, useRef, useState} from 'react';
import {Link, Route, Routes, useNavigate} from 'react-router-dom';
import {jwtDecode} from 'jwt-decode';


{/* API's for Flashcards.js */}
export const fetchRandomFlashcard = async (deckId, userId) => {
    try {
        const response = await fetch(`http://localhost:8080/api/study/random_deck`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                deck_id: parseInt(deckId), // Convert deck_id to a number
                user_id: userId,
            }),
        });

        if (response.ok) {
            const data = await response.json();
            return data.flashcard;
        } else {
            throw new Error('Error fetching flashcard');
        }
    } catch (error) {
        throw new Error('Error fetching flashcard');
    }
};

export const rateFlashcard = async (flashcardId, mark, userId) => {
    try {
        const response = await fetch('http://localhost:8080/api/study/rate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                flashcard_id: flashcardId,
                mark: mark,
                user_id: userId,
            }),
        });

        if (!response.ok) {
            throw new Error('Error recording mark');
        }
    } catch (error) {
        throw new Error('Error recording mark');
    }
};

export const deleteFlashcard = async (flashcardId, userId) => {
    try {
        const response = await fetch('http://localhost:8080/api/flashcards/delete_card', {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                flashcard_id: flashcardId,
                user_id: userId,
            }),
        });

        if (!response.ok) {
            throw new Error('Error deleting flashcard');
        }
    } catch (error) {
        throw new Error('Error deleting flashcard');
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

{/* API's for WordMatch.js */}

export const fetchRandomMatchingData = async (size, userId) => {
    try {
        const response = await fetch('http://localhost:8080/api/study/random_matching', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                size,
                user_id: userId,
            }),
        });

        if (response.ok) {
            return await response.json();
        } else {
            throw new Error('Failed to fetch matching data');
        }
    } catch (error) {
        throw new Error(`Error fetching matching data: ${error.message}`);
    }
};

{/* API's for WordMatch.js */}