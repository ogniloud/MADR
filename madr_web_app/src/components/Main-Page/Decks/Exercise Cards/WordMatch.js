import React, { useState, useEffect } from 'react';
import { jwtDecode } from 'jwt-decode';

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
            const responseMatchingData = await fetch('http://localhost:8080/api/study/random_matching', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    size: 5,
                    user_id: userId,
                }),
            });

            if (responseMatchingData.ok) {
                const dataMatchingData = await responseMatchingData.json();
                setMatchingData(dataMatchingData.matching);
            } else {
                console.error('Error fetching matching data:', responseMatchingData.statusText);
            }
        } catch (error) {
            console.error('Error fetching matching data:', error);
        }
    };

    const handlePairSelect = (property, value) => {
        setSelectedPairs({ ...selectedPairs, [property]: value });
    };


    const handleCheckAnswers = async () => {
        try {
            let isCorrect = true;

            for (const property in selectedPairs) {
                if (selectedPairs[property] !== matchingData.pairs[property]) {
                    isCorrect = false;
                    break;
                }
            }

            const mark = isCorrect ? 2 : 0;

            const responseRate = await fetch('http://localhost:8080/api/study/rate', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    mark: mark,
                    user_id: userId,
                }),
            });

            if (responseRate.ok) {
                setSuccessMessage(isCorrect ? 'Correct! Mark recorded.' : 'Incorrect. Try again.');
                setErrorMessage('');
                // Optionally, you can fetch new matching data here for the next round
            } else {
                console.error('Error recording mark:', responseRate.statusText);
            }
        } catch (error) {
            console.error('Error recording mark:', error);
        }
    };

    return (
        <div>
            <h2>Word Matching Exercise</h2>
            {matchingData ? (
                <div>
                    <div className="word-list">
                        {Object.keys(matchingData.cards).map((property, index) => (
                            <div key={index} className="word-item">
                                <p>{matchingData.cards[property].word}</p>
                            </div>
                        ))}
                    </div>
                    <div className="matching-options">
                        {Object.keys(matchingData.pairs).map((property, index) => (
                            <div key={index} className="option-item">
                                <label>
                                    <input
                                        type="checkbox"
                                        value={matchingData.pairs[property]}
                                        onChange={(e) => handlePairSelect(property, e.target.value)}
                                    />
                                    {matchingData.pairs[property]}
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


/// needed to be fixed