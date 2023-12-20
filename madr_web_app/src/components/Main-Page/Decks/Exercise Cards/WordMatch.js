import React, {useEffect, useState} from 'react';
import {jwtDecode} from 'jwt-decode';
import './Styles/WordMatch.css';

const WordMatch = () => {
    const [matchingData, setMatchingData] = useState(null);
    const [selectedPairs, setSelectedPairs] = useState({});
    const [successMessage, setSuccessMessage] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const [userId, setUserId] = useState(null);

    useEffect(() => {
        fetchUserIdAndMatchingData();
    }, [userId]);

    const fetchUserIdAndMatchingData = async () => {
        try {
            const token = localStorage.getItem('token');
            if (token) {
                const decodedToken = jwtDecode(token);
                setUserId(decodedToken.user_id);
                fetchMatchingData();
            } else {
                console.error('Error: User not authenticated.');
            }
        } catch (error) {
            console.error('Error fetching user_id:', error);
        }
    };

    const fetchMatchingData = async () => {
        try {
            const responseMatchingData = await fetch(
                'http://localhost:8080/api/study/random_matching',
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        size: 5,
                        user_id: userId,
                    }),
                }
            );

            if (responseMatchingData.ok) {
                const dataMatchingData = await responseMatchingData.json();
                setMatchingData(dataMatchingData.matching);
            } else {
                console.error(
                    'Error fetching matching data:',
                    responseMatchingData.statusText
                );
            }
        } catch (error) {
            console.error('Error fetching matching data:', error);
        }
    };

    const handlePairSelect = (property, value) => {
        setSelectedPairs({ ...selectedPairs, [property]: value });
    };

    const handleCheckAnswers = async () => {
        console.log('selectedPairs:', selectedPairs);
        console.log('matchingData.cards:', matchingData.cards);
        console.log('matchingData.pairs:', matchingData.pairs);
        try {
            const isCorrect = Object.keys(selectedPairs).every((property) => {
                return matchingData.cards[property].answer === selectedPairs[property]

                console.log('isCorrect:', isCorrect);
                console.log('mark:', mark);
            });


            const mark = isCorrect ? 2 : 0;

            for (const property in selectedPairs) {
                const responseRate = await fetch('http://localhost:8080/api/study/rate', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        mark: mark,
                        flashcard_id: matchingData.cards[property].id,
                        user_id: userId,
                    }),
                });

                if (!responseRate.ok) {
                    console.error('Error recording mark:', responseRate.statusText);
                    return;
                }
            }

            setSuccessMessage(
                isCorrect ? 'Correct! Mark recorded.' : 'Incorrect. Try again.'
            );
            setErrorMessage('');

            // Clear selected pairs and reset checkboxes
            setSelectedPairs({});
            document
                .querySelectorAll('input[type="checkbox"]')
                .forEach((checkbox) => {
                    checkbox.checked = false;
                });

            // Optionally, you can fetch new matching data here for the next round
            fetchMatchingData();
        } catch (error) {
            console.error('Error recording mark:', error);
            setErrorMessage('Error recording mark. Please try again.');
            setSuccessMessage('');
        }
    };

    let s = "";
    return (
        <div className="wm-container">
            <h2 className="wm-title">Word Match </h2>
            <div>
                <p className="wm-subtitle">Match the words with the suitable answers</p>
            </div>

            {matchingData ? (
                <div className="wm-exe-container">
                    <div className="wm-ex-box">
                        <div className="word-list">
                            {Object.keys(matchingData.cards).map((property, index) => (
                                <div key={index} className="word-item">
                                    <label>
                                        <input
                                            type="checkbox"
                                            value={matchingData.cards[property].word}
                                            onChange={() =>
                                                s = matchingData.cards[property].word
                                            }
                                        />
                                        {matchingData.cards[property].word}
                                    </label>
                                </div>
                            ))}
                        </div>
                        <div className="answer-list">
                            {Object.keys(matchingData.pairs).map((property, index) => (
                                <div key={index} className="answer-item">
                                    <label>
                                        <input
                                            type="checkbox"
                                            value={matchingData.pairs[property]}
                                            onChange={() =>
                                                handlePairSelect(
                                                    s,
                                                    matchingData.pairs[property]
                                                )
                                            }
                                        />
                                        {matchingData.pairs[property]}
                                    </label>
                                </div>
                            ))}
                        </div>
                    </div>

                    <div className="wm-button">
                        <button onClick={handleCheckAnswers}>Check Answers</button>
                        {successMessage && (
                            <p style={{ color: 'green' }}>{successMessage}</p>
                        )}
                        {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
                    </div>
                </div>
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};

export default WordMatch;
