import React, { useState, useEffect } from 'react';
import {useNavigate, useParams} from 'react-router-dom';
import {jwtDecode} from 'jwt-decode';
import { fetchRandomFlashcard, rateFlashcard, deleteFlashcard } from "../../API-Components/apiFunctions_exe_cards";
import ReactCardFlip from 'react-card-flip';
import './Styles/Flashcards.css';

const Flashcards = () => {
    const { deck_id } = useParams();
    const navigate = useNavigate();
    const [flashcard, setFlashcard] = useState(null);
    const [cardsSeen, setCardsSeen] = useState(0);
    const [userId, setUserId] = useState(null);
    const [isFlipped, setIsFlipped] = useState(false);
    const [backsideIndex, setBacksideIndex] = useState(0);

    useEffect(() => {
        fetchFlashcard();
        const token = localStorage.getItem('token');
        if (token) {
            const decodedToken = jwtDecode(token);
            setUserId(parseInt(decodedToken.user_id));
        }
    }, [deck_id, userId]);

    const fetchFlashcard = async () => {
        try {
            const fetchedFlashcard = await fetchRandomFlashcard(parseInt(deck_id), userId);
            setFlashcard(fetchedFlashcard);
            setCardsSeen(cardsSeen + 1);
            setIsFlipped(false); // Reset flip state for new card
            setBacksideIndex(0); // Reset backside index for new card
        } catch (error) {
            console.error(error.message);
        }
    };

    const handleMark = async (mark) => {
        try {
            await rateFlashcard(flashcard.id, mark, userId);
            fetchFlashcard();
        } catch (error) {
            console.error(error.message);
        }
    };

    const handleDeleteFlashcard = async () => {
        try {
            await deleteFlashcard(flashcard.id, userId);
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

    const handleClick = (e) => {
        e.preventDefault();
        setIsFlipped(!isFlipped);
    };

    const handleNextBackside = () => {
        if (backsideIndex < flashcard.multiple_backside.length - 1) {
            setBacksideIndex(backsideIndex + 1);
        }
    };

    const handlePrevBackside = () => {
        if (backsideIndex > 0) {
            setBacksideIndex(backsideIndex - 1);
        }
    };
    const handleReturnToDeckDetailsPageClick = () => {
        navigate(`/decks/${deck_id}`);
    };


    return (
        <div className="ex-flash">
            <div><h2 className="ex-flash-title" onClick={handleReturnToDeckDetailsPageClick}>Flashcards</h2></div>
            <div className="ex-flashcard-container" onContextMenu={handleContextMenu}>
                {flashcard && (
                    <ReactCardFlip isFlipped={isFlipped} flipDirection="horizontal">
                        <div className="ex-flashcard" onClick={handleClick}>
                            <div className="card-content">{flashcard.word}</div>
                        </div>

                        <div className="ex-flashcard ex-show-back" onClick={handleClick}>
                            <div className="card-content">{flashcard.multiple_backside[backsideIndex].value}</div>
                        </div>
                    </ReactCardFlip>
                )}
                {isFlipped && flashcard.multiple_backside.length > 1 && (
                    <div className="carousel-controls">
                        <button
                            className={`carousel-button-prev ${backsideIndex === 0 ? 'disabled' : ''}`}
                            onClick={handlePrevBackside}
                            style={{ backgroundColor: backsideIndex === 0 ? 'gray' : '' }}
                        >
                            Previous
                        </button>
                        <button
                            className={`carousel-button-nxt ${backsideIndex === flashcard.multiple_backside.length - 1 ? 'disabled' : ''}`}
                            onClick={handleNextBackside}
                            style={{ backgroundColor: backsideIndex === flashcard.multiple_backside.length - 1 ? 'gray' : '' }}
                        >
                            Next
                        </button>
                        <div className="carousel-slide-number">
                            {backsideIndex + 1}/{flashcard.multiple_backside.length}
                        </div>
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
