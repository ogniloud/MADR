import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';
import { fetchRandomFlashcard, rateFlashcard, deleteFlashcard } from "../../APIs/apiFunctions_exe_cards";
import './Styles/Flashcards.css';

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
            const fetchedFlashcard = await fetchRandomFlashcard(deck_id, userId);
            setFlashcard(fetchedFlashcard);
            setCardsSeen(cardsSeen + 1);
        } catch (error) {
            console.error(error.message);
        }
    };

    const handleMark = async (mark) => {
        try {
            await rateFlashcard(flashcard.id, mark, userId);
            // Fetch the next flashcard after recording the mark
            fetchFlashcard();
        } catch (error) {
            console.error(error.message);
        }
    };

    const handleDeleteFlashcard = async () => {
        try {
            await deleteFlashcard(flashcard.id, userId);
            // Fetch the next flashcard after successful deletion
            fetchFlashcard();
            alert('Flashcard deleted successfully');
        } catch (error) {
            console.error(error.message);
            alert('Error deleting flashcard');
        }
    };

    const handleContextMenu = (e) => {
        e.preventDefault();
        if (window.confirm('Are you sure you want to delete this flashcard?')) {
            handleDeleteFlashcard();
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
            <div><h2 className="ex-flash-title">Flashcards</h2></div>
            <div className="ex-flashcard-container" onClick={handleCardClick} onContextMenu={handleContextMenu}>
                {flashcard && (
                    <div className={`ex-flashcard ${flashcard.showBack ? 'ex-show-back' : ''}`}>
                        <div
                            className="card-content">{flashcard.showBack ? flashcard.backside.value : flashcard.word}</div>
                    </div>
                )}
            </div>
            <div className="ex-flashcard-buttons-container">
                <button className="ex-flashcard-button-1" onClick={() => handleMark(0)}>Can't recall it</button>
                <button className="ex-flashcard-button-1" onClick={() => handleMark(1)}>Yeah but...</button>
                <button className="ex-flashcard-button-1" onClick={() => handleMark(2)}>Yeah, perfectly</button>
            </div>
            <p>Cards Seen: {cardsSeen}</p>
        </div>
    );
};

export default Flashcards;
