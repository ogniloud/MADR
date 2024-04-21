import {jwtDecode} from 'jwt-decode';


{/* API's for AllWords.js */}
export const fetchFlashcards = async (deckId) => {
    try {
        const token = localStorage.getItem('token');
        if (!token) {
            return [];
        }

        const decodedToken = jwtDecode(token);
        const user_id = decodedToken.user_id;

        const requestBody = {
            deck_id: Number(deckId),
        };

        const response = await fetch(`http://${process.env.REACT_APP_API_HOST}/api/flashcards/cards`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestBody),
        });

        if (response.ok) {
            const data = await response.json();
            return data.flashcards || [];
        } else {
            console.error('Error loading flashcards:', response.statusText);
            return [];
        }
    } catch (error) {
        console.error('Error loading flashcards:', error);
        return [];
    }
};

export const addFlashcard = async (flashcardData) => {
    try {
        const token = localStorage.getItem('token');
        if (!token) {
            return;
        }

        // Additional validation
        if (!flashcardData.word || !flashcardData.answer || !flashcardData.backsideValue) {
            throw new Error('Please fill in all fields.');
        }

        const decodedToken = jwtDecode(token);
        const user_id = decodedToken.user_id;

        const deckIdAsNumber = Number(flashcardData.deck_id);

        const requestBody = {
            word: flashcardData.word,
            answer: flashcardData.answer.replace(/^"|"$/g, ''), // Remove double quotes if present
            backside: {
                type: Number(flashcardData.backsideType),
                value: flashcardData.backsideValue,
            },
            deck_id: deckIdAsNumber,
            user_id,
        };

        console.log('Request Body:', JSON.stringify(requestBody));

        const response = await fetch(`http://${process.env.REACT_APP_API_HOST}/api/flashcards/add_card`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestBody),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(`Error: ${errorData.message}`);
        }
    } catch (error) {
        throw new Error('An unexpected error occurred.');
    }
};

export const deleteFlashcard = async (flashcardId) => {
    try {
        const token = localStorage.getItem('token');
        if (!token) {
            return;
        }

        const response = await fetch(`http://${process.env.REACT_APP_API_HOST}/api/flashcards/delete_card`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                flashcard_id: flashcardId,
                user_id: getUserId(),
            }),
        });

        if (!response.ok) {
            throw new Error('Error deleting word');
        }
    } catch (error) {
        throw new Error('Error deleting word');
    }
};

{/* API's for AllWords.js */}

