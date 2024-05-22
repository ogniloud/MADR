import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';
import { createDeck } from "../API-Components/apiFunctions_decks";
import './Styles/CreateDecks.css';

const CreateDecks = ({ fetchUserDecks }) => {
    const [deckName, setDeckName] = useState('');
    const [flashcards, setFlashcards] = useState([]);
    const [currentFlashcard, setCurrentFlashcard] = useState({ word: '', answer: '', backside: { type: 0, value: '' } });
    const [errorMessage, setErrorMessage] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const navigate = useNavigate();
    const decodedToken = jwtDecode(localStorage.getItem('token'));
    const userId = decodedToken.user_id;

    const addFlashcard = () => {
        setFlashcards([...flashcards, currentFlashcard]);
        setCurrentFlashcard({ word: '', answer: '', backside: { type: 0, value: '' } }); // Reset the current flashcard
        setSuccessMessage('Flashcard added successfully');
        setTimeout(() => setSuccessMessage(''), 2000); // Clear the message after 2 seconds
    };

    const handleFlashcardChange = (field, value) => {
        if (field === 'backside') {
            setCurrentFlashcard({ ...currentFlashcard, backside: { ...currentFlashcard.backside, value } });
        } else {
            setCurrentFlashcard({ ...currentFlashcard, [field]: value });
        }
    };

    const handleDeckNameChange = (value) => {
        setDeckName(value);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            let valid = true;
            flashcards.forEach((f) => {
                if (f.word === '' || f.answer === '' || f.backside.value === '') {
                    setErrorMessage("Some fields are empty");
                    valid = false;
                }
            });
            if (!valid) return;

            const deckData = {
                flashcards,
                name: deckName,
                user_id: userId,
            };

            const token = localStorage.getItem('token');
            await createDeck(deckData, token);

            setErrorMessage('');

            if (typeof fetchUserDecks === 'function') {
                await fetchUserDecks();
            }

            navigate('/decks');
        } catch (error) {
            console.error('Error creating deck:', error.message);
            setErrorMessage(error.message);
        }


    };


    const handleTitleClick = () => {
        navigate('/mainpage');
    };




    return (
        <div className="create-decks-page-container">
            <div><h2 className="cd-title" onClick={handleTitleClick}>Create Deck</h2></div>
            <div className="create-decks-main-container">
                <div className="create-decks-container">
                    <h2 className="create-deck-title">Create a New Deck</h2>
                    <form className="deck-form" onSubmit={handleSubmit}>
                        <label htmlFor="deckName">Deck Name:</label>
                        <input
                            required
                            type="text"
                            id="deckName"
                            value={deckName}
                            onChange={(e) => handleDeckNameChange(e.target.value)}
                        />
                        <h3 className="flashcards-do">With Flashcards</h3>
                        <div className="flashcard-form">
                            <label htmlFor="word">Word:</label>
                            <input
                                required
                                type="text"
                                id="word"
                                value={currentFlashcard.word}
                                onChange={(e) => handleFlashcardChange('word', e.target.value)}
                            />
                            <label htmlFor="answer">Answer:</label>
                            <input
                                required
                                type="text"
                                id="answer"
                                value={currentFlashcard.answer}
                                onChange={(e) => handleFlashcardChange('answer', e.target.value)}
                            />
                            <label htmlFor="backside">Backside:</label>
                            <input
                                required
                                type="text"
                                id="backside"
                                value={currentFlashcard.backside.value}
                                onChange={(e) => handleFlashcardChange('backside', e.target.value)}
                            />
                        </div>
                        {errorMessage && <div className="in-deck-error-message">{errorMessage}</div>}
                        <button className="in-deck-add-fc" type="button" onClick={addFlashcard}>Add Flashcard</button>
                        <button className="in-deck-add-dk" type="submit">Create Deck</button>
                    </form>
                    {successMessage && <div className="in-deck-success-message">{successMessage}</div>}
                </div>
                <div className="flashcards-grid-container">
                    <h3>Added Flashcards </h3>
                    <div className="in-deck-flashcards-grid">
                        {flashcards.map((flashcard, index) => (
                            <div key={index} className="flashcard-display">
                                <p><strong>Word:</strong> {flashcard.word}</p>
                                <p><strong>Answer:</strong> {flashcard.answer}</p>
                                <p><strong>Backside:</strong> {flashcard.backside.value}</p>
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    );
};

export default CreateDecks;
