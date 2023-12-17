import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import {jwtDecode}from 'jwt-decode';

const AllWords = () => {
    const { deck_id } = useParams();

    const [word, setWord] = useState('');
    const [answer, setAnswer] = useState('');
    const [backsideType, setBacksideType] = useState(0);
    const [backsideValue, setBacksideValue] = useState('');
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    useEffect(() => {
        // You can use deck_id here for any initialization or additional data fetching
    }, [deck_id]);

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

            const requestBody = {
                word: word,
                answer: answer.replace(/^"|"$/g, ''), // Remove double quotes if present
                backside: {
                    type: Number(backsideType),
                    value: backsideValue,
                },
                deck_id: deck_id,
                user_id: user_id,
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
            } else {
                const errorData = await response.json();
                setSuccessMessage('');
                setErrorMessage(`Error: ${errorData.message}`);
            }
        } catch (error) {
            console.error('Error adding flashcard:', error);
            setSuccessMessage('');
            setErrorMessage('An unexpected error occurred.');
        }
    };


    return (
        <div>
            <h2>All Words</h2>
            <div>
                <label>
                    Word:
                    <input type="text" value={word} onChange={(e) => setWord(e.target.value)} />
                </label>
            </div>
            <div>
                <label>
                    Answer:
                    <input type="text" value={answer} onChange={(e) => setAnswer(e.target.value)} />
                </label>
            </div>
            <div>
                <label>
                    Backside Type:
                    <select value={backsideType} onChange={(e) => setBacksideType(Number(e.target.value))}>
                        <option value={0}>Type 0</option>
                        {/* Add more options if needed */}
                    </select>
                </label>
            </div>
            <div>
                <label>
                    Backside Value:
                    <input type="text" value={backsideValue} onChange={(e) => setBacksideValue(e.target.value)} />
                </label>
            </div>
            <button onClick={handleAddFlashcard}>Add New Card</button>
            {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
            {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
        </div>
    );
};

export default AllWords;
