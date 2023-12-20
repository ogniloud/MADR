import React, {useEffect, useState} from 'react';
import {useParams} from 'react-router-dom';
import {jwtDecode} from 'jwt-decode';
import './Styles/AllWords.css';

const AllWords = () => {
    const { deck_id } = useParams();

    const [word, setWord] = useState('');
    const [answer, setAnswer] = useState('');
    const [backsideType, setBacksideType] = useState(0);
    const [backsideValue, setBacksideValue] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [flashcards, setFlashcards] = useState([]);

    useEffect(() => {

        loadFlashcards();
    }, [deck_id]);

    const loadFlashcards = async () => {
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                // Handle case where the user is not authenticated
                return;
            }

            const decodedToken = jwtDecode(token);
            const user_id = decodedToken.user_id;

            const requestBody = {
                deck_id: Number(deck_id),
            };

            const response = await fetch('http://localhost:8080/api/flashcards/cards', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(requestBody),
            });

            if (response.ok) {
                const data = await response.json();
                setFlashcards(data.flashcards);
            } else {
                console.error('Error loading flashcards:', response.statusText);
            }
        } catch (error) {
            console.error('Error loading flashcards:', error);
        }
    };

    const handleAddFlashcard = async () => {
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                // Handle case where the user is not authenticated
                return;
            }

            // Additional validation
            if (!word || !answer || !backsideValue) {
                setErrorMessage('Please fill in all fields.');
                return;
            }

            const decodedToken = jwtDecode(token);
            const user_id = decodedToken.user_id;

            // Convert deck_id to a number
            const deckIdAsNumber = Number(deck_id);

            const requestBody = {
                word,
                answer: answer.replace(/^"|"$/g, ''), // Remove double quotes if present
                backside: {
                    type: Number(backsideType),
                    value: backsideValue,
                },
                deck_id: deckIdAsNumber,
                user_id,
            };

            console.log('Request Body:', JSON.stringify(requestBody));

            const response = await fetch('http://localhost:8080/api/flashcards/add_card', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(requestBody),
            });

            console.log('Response:', response);

            if (response.ok) {
                setSuccessMessage('Flashcard added successfully!');
                setErrorMessage('');

                // Clear form fields after success
                setWord('');
                setAnswer('');
                setBacksideType(0);
                setBacksideValue('');
            } else {
                const errorData = await response.json();
                setSuccessMessage('');
                setErrorMessage(`Error: ${errorData.message}`);
            }
            window.location.reload()
        } catch (error) {
            console.error('Error adding flashcard:', error);
            setSuccessMessage('');
            setErrorMessage('An unexpected error occurred.');
        }
    };
    return (
        <div className="all-words-container">
            <div>
                <h2 className="all-words-title"> All Words</h2>
            </div>
            <div className="all-words-flashcard-list">
                <h3 className="all-words-subtitle">Browse Cards</h3>
                <div className="all-words-flashcard-grid">
                    {flashcards.map((flashcard, index) => (
                        <div key={flashcard.id} className="all-words-flashcard-item">
                            {flashcard.word}
                        </div>
                    ))}
                </div>
            </div>
            <div className="all-words-add-flashcards">
                <h3 className="all-words-box-title"> Create New Cards </h3>
                <label>
                    Word:
                    <input type="text" value={word} onChange={(e) => setWord(e.target.value)}/>
                </label>
                <label>
                    Answer:
                    <input type="text" value={answer} onChange={(e) => setAnswer(e.target.value)}/>
                </label>
                <label>
                    Backside Type:
                    <select value={backsideType} onChange={(e) => setBacksideType(Number(e.target.value))}>
                        <option value={0}>Type 0</option>
                    </select>
                </label>
                <label>
                    Backside Value:
                    <input type="text" value={backsideValue} onChange={(e) => setBacksideValue(e.target.value)}/>
                </label>
            </div>
            <button className='all-words-submit-button' onClick={handleAddFlashcard}>Add New Card</button>
            {successMessage && <p className="all-words-success-message">{successMessage}</p>}
            {errorMessage && <p className="all-words-error-message">{errorMessage}</p>}
        </div>
    );
};

export default AllWords;
