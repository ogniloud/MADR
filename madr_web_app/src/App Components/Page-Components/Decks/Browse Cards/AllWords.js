import React, {useEffect, useState} from 'react';
import {useParams} from 'react-router-dom';
import {jwtDecode} from 'jwt-decode';
import {addFlashcard, deleteFlashcard, fetchFlashcards} from "../../API-Components/apiFunctions_browse_cards";
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
            const fetchedFlashcards = await fetchFlashcards(deck_id);
            setFlashcards(fetchedFlashcards);
        } catch (error) {
            console.error(error.message);
        }
    };

    const handleAddFlashcard = async () => {
        try {
            await addFlashcard({ word, answer, backsideType, backsideValue, deck_id });
            setSuccessMessage('Flashcard added successfully!');
            setErrorMessage('');

            // Clear form fields after success
            setWord('');
            setAnswer('');
            setBacksideType(0);
            setBacksideValue('');
            loadFlashcards(); // Reload flashcards
        } catch (error) {
            console.error(error.message);
            setSuccessMessage('');
            setErrorMessage(error.message);
        }
    };

    const handleDeleteWord = async (flashcardId) => {
        try {
            await deleteFlashcard(flashcardId);
            setFlashcards(prevFlashcards => prevFlashcards.filter(flashcard => flashcard.id !== flashcardId));
            setSuccessMessage('Word deleted successfully');
        } catch (error) {
            console.error(error.message);
            setErrorMessage('Error deleting word');
        }
    };


    const getUserId = () => {
        const token = localStorage.getItem('token');
        if (token) {
            const decodedToken = jwtDecode(token);
            return decodedToken.user_id;
        }
        return null;
    };

    const handleContextMenu = (e, flashcardId) => {
        e.preventDefault();
        if (window.confirm('Are you sure you want to delete this word?')) {
            handleDeleteWord(flashcardId);
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
                    {flashcards.map((flashcard) => (
                        <div key={flashcard.id} className="all-words-flashcard-item" onContextMenu={(e) => handleContextMenu(e, flashcard.id)}>
                            {flashcard.word}
                        </div>
                    ))}
                </div>
            </div>
            <form className="all-words-add-flashcards">
                <h3 className="all-words-box-title"> Create New Cards </h3>
                <label>
                    Word:
                    <input required type="text" value={word} onChange={(e) => setWord(e.target.value)}/>
                </label>
                <label>
                    Answer:
                    <input required type="text" value={answer} onChange={(e) => setAnswer(e.target.value)}/>
                </label>
                <label>
                    Backside Type:
                    <select value={backsideType} onChange={(e) => setBacksideType(Number(e.target.value))}>
                        <option value={0}>Plain text</option>
                        <option value={1}>Definition</option>
                        <option value={2}>Translation</option>
                        <option value={3}>Base64 PNG Image</option> {/*может сделаем, но пусть будет>*/}
                    </select>
                </label>
                <label>
                Backside Value:
                    <input required type="text" value={backsideValue} onChange={(e) => setBacksideValue(e.target.value)}/>
                </label>
            </form>
            <button className='all-words-submit-button' onClick={handleAddFlashcard}>Add New Card</button>
            {successMessage && <p className="all-words-success-message">{successMessage}</p>}
            {errorMessage && <p className="all-words-error-message">{errorMessage}</p>}
        </div>
    );
};

export default AllWords;
