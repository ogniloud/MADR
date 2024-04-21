import {jwtDecode} from 'jwt-decode';
import axios from 'axios';

{/* API's for CreateDecks.js */}
export const createDeck = async (deckData, token) => {
    try {
        const response = await axios.put(
            'http://${process.env.REACT_APP_API_HOST}/api/flashcards/new_deck',
            deckData,
            {
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${token}`,
                },
            }
        );
        return response.data;
    } catch (error) {
        throw new Error('Failed to create deck. Please try again.');
    }
};

{/* API's for CreateDecks.js */}



{/* API's for DeckPage.js */}
export const fetchUserDecks = async (token) => {
    try {
        const decodedToken = jwtDecode(token);

        const response = await fetch('http://${process.env.REACT_APP_API_HOST}/api/flashcards/load', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ user_id: decodedToken.user_id }),
        });

        if (response.ok) {
            const { decks } = await response.json();
            return decks;
        } else {
            throw new Error('Failed to load user decks');
        }
    } catch (error) {
        throw new Error('Error loading user decks');
    }
};

export const deleteDeck = async (deckId, userId, token) => {
    try {
        const response = await fetch('http://${process.env.REACT_APP_API_HOST}/api/flashcards/delete_deck', {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
                deck_id: deckId,
                user_id: userId,
            }),
        });

        if (!response.ok) {
            throw new Error('Failed to delete deck');
        }

        const data = await response.json();
        return data.ok;
    } catch (error) {
        throw new Error('Error deleting deck');
    }
};

export const checkIfShared = async (deckId, userId, token) => {
    try {
        const response = await fetch('http://${process.env.REACT_APP_API_HOST}/api/social/is_shared', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
                deck_id: deckId,
                user_id: userId,
            }),
        });

        if (!response.ok) {
            throw new Error('Failed to share deck');
        }

        return response.json()
    } catch (error) {
        throw new Error('Error check if shared');
    }
}

export const shareDeck = async (deckId, userId, token) => {
    try {
        const response = await fetch('http://${process.env.REACT_APP_API_HOST}/api/social/share', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
                deck_id: deckId,
                user_id: userId,
            }),
        });

        if (!response.ok) {
            throw new Error('Failed to share deck');
        }
    } catch (error) {
        throw new Error('Error sharing deck');
    }
}

export const checkIfSharedByGroups = async (deckId, userId, token) => {
    try {
        const response = await fetch('http://${process.env.REACT_APP_API_HOST}/api/social/groups_shared', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
                deck_id: deckId,
                creator_id: userId,
            }),
        });

        if (!response.ok) {
            throw new Error('Failed to share deck');
        }
        return response.json()
    } catch (error) {
        throw new Error('Error sharing deck');
    }
}

export const shareGroup = async (userId, groupId, deckId, token) => {
    try {
        const response = await fetch('http://${process.env.REACT_APP_API_HOST}/api/groups/share', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({
                deck_id: deckId,
                user_id: userId,
                group_id: groupId
            }),
        });

        if (!response.ok) {
            throw new Error('Failed to share deck');
        }
    } catch (error) {
        throw new Error('Error sharing deck');
    }
}

{/* API's for DeckPage.js */}




