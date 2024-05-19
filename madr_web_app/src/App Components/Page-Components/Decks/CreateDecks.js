import React, {useState} from 'react';
import {jwtDecode} from 'jwt-decode';
import {useNavigate} from 'react-router-dom';
import {createDeck} from "../API-Components/apiFunctions_decks";
import './Styles/CreateDecks.css';


const CreateDecks = ({ fetchUserDecks }) => {
    const [deckName, setDeckName] = useState('');
    const [flashcards, setFlashcards] = useState([{ word: '', answer: '', backside: { type: 0, value: '' } }]);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate(); // Initialize navigate
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

    return (

        <div className="create-decks-page-container">
            <div><h2 className="cd-title">Create Deck</h2></div>

            <div className="create-decks-container">
                <h2 className="create-deck-title">Create a New Deck</h2>
                <form className="deck-form">
                    <label htmlFor="deckName">Deck Name:</label>
                    <input
                        required
                        type="text"
                        id="deckName"
                        value={deckName}
                        onChange={(e) => handleDeckNameChange(e.target.value)}
                    />

                    <h3 className="flashcards-do">Flashcards:</h3>
                    {flashcards.map((flashcard, index) => (
                        <div key={index} className="flashcard-form">
                            <label htmlFor={`word${index}`}>Word:</label>
                            <input
                                required
                                type="text"
                                id={`word${index}`}
                                value={flashcard.word}
                                onChange={(e) => handleFlashcardChange(index, 'word', e.target.value)}
                            />

                            <label htmlFor={`answer${index}`}>Answer:</label>
                            <input
                                required
                                type="text"
                                id={`answer${index}`}
                                value={flashcard.answer}
                                onChange={(e) => handleFlashcardChange(index, 'answer', e.target.value)}
                            />

                            <label htmlFor={`backside${index}`}>Backside:</label>
                            <input
                                required
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
                </form>
            </div>
        </div>
    );
};

export default CreateDecks;
