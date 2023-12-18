import React, { useState, useEffect } from 'react';
import { jwtDecode } from 'jwt-decode';

const WordMatch = () => {
    const [flashcards, setFlashcards] = useState(null);
    const [selectedPairs, setSelectedPairs] = useState({});
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [userId, setUserId] = useState(null);

    useEffect(() => {
        fetchFlashcardsData();
    }, []);

    const fetchFlashcardsData = async () => {
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                console.error('Error: User not authenticated.');
                return;
            }

            const decodedToken = jwtDecode(token);
            const user_id = decodedToken.user_id;

            const responseFlashcards = await fetch('http://localhost:8080/api/study/random_matching_deck', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    size: 5,
                    user_id: user_id,
                }),
            });

            if (responseFlashcards.ok) {
                const dataFlashcards = await responseFlashcards.json();
                setFlashcards(dataFlashcards.flashcards);
                setUserId(user_id);
            } else {
                console.error('Error fetching flashcards:', responseFlashcards.statusText);
            }
        } catch (error) {
            console.error('Error fetching flashcards:', error);
        }
    };

    const handlePairSelect = (flashcardId, side, value) => {
        setSelectedPairs({
            ...selectedPairs,
            [flashcardId]: {
                ...selectedPairs[flashcardId],
                [side]: value,
            },
        });
    };

    const handleCheckAnswers = async () => {
        try {
            let totalCorrectAnswers = 0;

            for (const flashcardId in selectedPairs) {
                const { word, answer } = flashcards.find(card => card.flashcard_id === Number(flashcardId));
                const selectedWord = selectedPairs[flashcardId].word;
                const selectedAnswer = selectedPairs[flashcardId].answer;

                if (selectedWord === answer && selectedAnswer === word) {
                    totalCorrectAnswers++;
                }
            }

            const mark = totalCorrectAnswers === Object.keys(selectedPairs).length ? 2 : 0;

            for (const flashcardId in selectedPairs) {
                const responseRate = await fetch('http://localhost:8080/api/study/rate', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        flashcard_id: Number(flashcardId), // Ensure flashcard_id is a number
                        mark: mark,
                        user_id: userId,
                    }),
                });

                if (!responseRate.ok) {
                    console.error('Error recording mark for flashcard', flashcardId, ':', responseRate.statusText);
                }
            }

            setSuccessMessage(
                totalCorrectAnswers === Object.keys(selectedPairs).length
                    ? 'All answers are correct! Marks recorded.'
                    : 'Some answers are incorrect. Try again.'
            );
            setErrorMessage('');
        } catch (error) {
            console.error('Error recording marks:', error);
        }
    };

    return (
        <div>
            <h2>Word Matching Exercise</h2>
            {flashcards ? (
                <div>
                    <div className="word-list">
                        {flashcards.map((card) => (
                            <div key={card.flashcard_id} className="word-item">
                                <p>{card.word}</p>
                            </div>
                        ))}
                    </div>
                    <div className="matching-options">
                        {flashcards.map((card) => (
                            <div key={card.flashcard_id} className="option-item">
                                <label>
                                    <input
                                        type="checkbox"
                                        value={card.answer}
                                        onChange={(e) => handlePairSelect(card.flashcard_id, 'word', e.target.value)}
                                    />
                                    {card.answer}
                                </label>
                                <label>
                                    <input
                                        type="checkbox"
                                        value={card.word}
                                        onChange={(e) => handlePairSelect(card.flashcard_id, 'answer', e.target.value)}
                                    />
                                    {card.word}
                                </label>
                            </div>
                        ))}
                    </div>
                    <button onClick={handleCheckAnswers}>Check Answers</button>
                    {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
                    {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
                </div>
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};

export default WordMatch;
