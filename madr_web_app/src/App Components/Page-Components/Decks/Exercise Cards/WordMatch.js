import React, {useEffect, useState} from 'react';
import {jwtDecode} from 'jwt-decode';
import {fetchRandomMatchingData, rateFlashcard} from "../../API-Components/apiFunctions_exe_cards";
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
                setUserId(parseInt(decodedToken.user_id));
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
            const dataMatchingData = await fetchRandomMatchingData(5, parseInt(userId));
            setMatchingData(dataMatchingData.matching);
        } catch (error) {
            console.error(error.message);
        }
    };

    const handlePairSelect = (property, value) => {
        setSelectedPairs({ ...selectedPairs, [property]: value });
    };

    const handleCheckAnswers = async () => {
        try {
            const isCorrect = Object.keys(selectedPairs).every((property) => {
                return matchingData.cards[property].answer === selectedPairs[property];
            });

            const mark = isCorrect ? 2 : 0;

            for (const property in selectedPairs) {
                await rateFlashcard(matchingData.cards[property].id, mark, userId);
            }

            setSuccessMessage(isCorrect ? 'Correct! Mark recorded.' : 'Incorrect. Try again.');
            setErrorMessage('');

            setSelectedPairs({});
            document.querySelectorAll('input[type="checkbox"]').forEach((checkbox) => {
                checkbox.checked = false;
            });

            fetchMatchingData();
        } catch (error) {
            console.error('Error recording mark:', error);
            setErrorMessage('Error recording mark. Please try again.');
            setSuccessMessage('');
        }
    };

    let s = '';
    return (
        <div className="wm-container">
            <h2 className="wm-title">Word Match </h2>
            <div>
                <p className="wm-subtitle">Match the words with suitable answers</p>
            </div>

            {matchingData ? (
                <div className="wm-exe-container">
                    <div className="wm-ex-box">
                        <div className="wm-word-list">
                            <div>
                                <h3 className="wm-word"> Word keys </h3>
                            </div>
                            {Object.keys(matchingData.cards).map((property, index) => (
                                <div key={index} className="wm-word-item">
                                    <label>
                                        <input className="wm-checkboxes-word"
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
                        <div className="wm-answer-list">
                            <div>
                                <h3 className="wm-word"> Answer keys </h3>
                            </div>
                            {Object.keys(matchingData.pairs).map((property, index) => (
                                <div key={index} className="wm-answer-item">
                                    <label>
                                        <input className="wm-checkboxes-answer"
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
                        <button className="wm-submit-button" onClick={handleCheckAnswers}>Check Answers</button>
                        {successMessage && (
                            <p className="wm-successMessage" style={{ color: 'green' }}>{successMessage}</p>
                        )}
                        {errorMessage && <p className="wm-errorMessage" style={{ color: 'red' }}>{errorMessage}</p>}
                    </div>
                </div>
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};

export default WordMatch;
