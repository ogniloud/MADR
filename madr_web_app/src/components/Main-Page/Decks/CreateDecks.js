import React, { useState } from 'react';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';
import { useNavigate } from 'react-router-dom'; // Import the useNavigate hook
import './CreateDecks.css';

const CreateDecks = ({ fetchUserDecks }) => {
    const [deckName, setDeckName] = useState('');
    const [flashcards, setFlashcards] = useState([
        { word: '', answer: '', backside: { type: 0, value: '' } },
    ]);
    const [errorMessage, setErrorMessage] = useState('');

    const navigate = useNavigate(); // Initialize navigate

    // Get the user ID from the decoded token
    const decodedToken = jwtDecode(localStorage.getItem('token'));
    const userId = decodedToken.user_id;

    const addFlashcard = () => {
        setFlashcards([...flashcards, { word: '', answer: '', backside: { type: 0, value: '' } }]);
    };

    const handleFlashcardChange = (index, field, value) => {
        const updatedFlashcards = [...flashcards];
        if (field === 'backside') {
            updatedFlashcards[index][field].value = value;
        } else {
            updatedFlashcards[index][field] = value;
        }
        setFlashcards(updatedFlashcards);
    };

    const handleDeckNameChange = (value) => {
        setDeckName(value);
    };

    const handleSubmit = async () => {
        try {
            const response = await axios.put(
                'http://localhost:8080/api/flashcards/new_deck',
                {
                    flashcards,
                    name: deckName,
                    user_id: userId, // Pass the user ID to the server
                },
                {
                    headers: {
                        'Content-Type': 'application/json',
                        Authorization: `Bearer ${localStorage.getItem('token')}`, // Passing  the JWT token to the server
                    },
                }
            );

            console.log('Deck created successfully:', response.data);
            setErrorMessage('');

            if (typeof fetchUserDecks === 'function') {
                await fetchUserDecks();
            }


            // Navigate to MainPage after successful deck creation
            navigate('/decks');
        } catch (error) {
            console.error('Error creating deck:', error.message);
            setErrorMessage('Failed to create the deck. Please try again.');
        }
    };

    return (
        <div className="create-decks-container">
            <h2>Create a New Deck</h2>
            <div className="deck-form">
                <label htmlFor="deckName">Deck Name:</label>
                <input
                    type="text"
                    id="deckName"
                    value={deckName}
                    onChange={(e) => handleDeckNameChange(e.target.value)}
                />

                <h3>Flashcards:</h3>
                {flashcards.map((flashcard, index) => (
                    <div key={index} className="flashcard-form">
                        <label htmlFor={`word${index}`}>Word:</label>
                        <input
                            type="text"
                            id={`word${index}`}
                            value={flashcard.word}
                            onChange={(e) => handleFlashcardChange(index, 'word', e.target.value)}
                        />

                        <label htmlFor={`answer${index}`}>Answer:</label>
                        <input
                            type="text"
                            id={`answer${index}`}
                            value={flashcard.answer}
                            onChange={(e) => handleFlashcardChange(index, 'answer', e.target.value)}
                        />

                        <label htmlFor={`backside${index}`}>Backside:</label>
                        <input
                            type="text"
                            id={`backside${index}`}
                            value={flashcard.backside.value}
                            onChange={(e) => handleFlashcardChange(index, 'backside', e.target.value)}
                        />
                    </div>
                ))}

                {errorMessage && <div className="error-message">{errorMessage}</div>}

                <button onClick={addFlashcard}>Add Flashcard</button>
                <button onClick={handleSubmit}>Create Deck</button>
            </div>
        </div>
    );
};

export default CreateDecks;
