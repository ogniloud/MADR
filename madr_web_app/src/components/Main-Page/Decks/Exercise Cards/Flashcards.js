import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';
import './Styles/Flashcards.css'

const Flashcards = () => {
    const { deck_id } = useParams();
    const [flashcard, setFlashcard] = useState(null);
    const [cardsSeen, setCardsSeen] = useState(0);
    const [userId, setUserId] = useState(null);

    useEffect(() => {
        // Fetch the initial flashcard when the component mounts
        fetchFlashcard();
        // Decode the token to get user_id
        const token = localStorage.getItem('token');
        if (token) {
            const decodedToken = jwtDecode(token);
            setUserId(decodedToken.user_id);
        }
    }, [deck_id, userId]);

    const fetchFlashcard = async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/study/random_deck`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    deck_id: parseInt(deck_id), // Convert deck_id to a number
                    user_id: userId,
                }),
            });

            if (response.ok) {
                const data = await response.json();
                setFlashcard(data.flashcard);
                setCardsSeen(cardsSeen + 1);
            } else {
                console.error('Error fetching flashcard:', response.statusText);
            }
        } catch (error) {
            console.error('Error fetching flashcard:', error);
        }
    };

    const handleMark = async (mark) => {
        try {
            const response = await fetch('http://localhost:8080/api/study/rate', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    flashcard_id: flashcard.id,
                    mark: mark,
                    user_id: userId,
                }),
            });

            if (response.ok) {
                // Fetch the next flashcard after recording the mark
                fetchFlashcard();
            } else {
                console.error('Error recording mark:', response.statusText);
            }
        } catch (error) {
            console.error('Error recording mark:', error);
        }
    };

    const handleCardClick = () => {
        // Toggle between front and back of the card
        setFlashcard((prevFlashcard) => ({
            ...prevFlashcard,
            showBack: !prevFlashcard.showBack,
        }));
    };

    return (
        <div className="ex-flash">
            <div><h2 className="x-flash-title">Flashcards</h2></div>

            <div className="flashcard-container" onClick={handleCardClick}>
                {flashcard && (
                    <div className={`flashcard ${flashcard.showBack ? 'show-back' : ''}`}>
                        <div className="card-content">{flashcard.showBack ? flashcard.backside.value : flashcard.word}</div>
                    </div>
                )}
            </div>
            <div className="buttons-container">
                <button onClick={() => handleMark(0)}>Can't recall it</button>
                <button onClick={() => handleMark(1)}>Yeah but...</button>
                <button onClick={() => handleMark(2)}>Yeah, perfectly</button>
            </div>
            <p>Cards Seen: {cardsSeen}</p>
        </div>
    );
};

export default Flashcards;
