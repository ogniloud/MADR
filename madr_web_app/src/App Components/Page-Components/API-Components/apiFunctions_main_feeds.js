import {jwtDecode} from 'jwt-decode';


{/* API's for MainPage.js */}

export const fetchFollowers = async (token, userId) => {
    try {
        const response = await fetch(`${process.env.REACT_APP_API_HOST}/api/social/followers`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ user_id: userId })
        });
        return response.json();
    } catch (error) {
        console.error('Error fetching followers:', error);
        throw error;
    }
};

export const fetchFollowings = async (token, userId) => {
    try {
        const response = await fetch(`${process.env.REACT_APP_API_HOST}/api/social/followings`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ user_id: userId })
        });
        return response.json();
    } catch (error) {
        console.error('Error fetching followings:', error);
        throw error;
    }
};

export const searchUsers = async (token, query) => {
    try {
        const response = await fetch(`${process.env.REACT_APP_API_HOST}/api/social/search?q=${encodeURIComponent(query)}`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        if (response.ok) {
            const searchData = await response.json();
            return searchData.users || [];
        } else {
            console.error('Error searching users:', response.statusText);
            return [];
        }
    } catch (error) {
        console.error('Error searching users:', error);
        throw error;
    }
};

export const followUser = async (token, userId, followerId) => {
    try {
        const requestBody = { author_id: userId, follower_id: followerId };
        const response = await fetch(`${process.env.REACT_APP_API_HOST}/api/social/follow`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestBody)
        });
        return response.ok;
    } catch (error) {
        console.error('Error following user:', error);
        throw error;
    }
};

export const unfollowUser = async (token, userId, followerId) => {
    try {
        const requestBody = { author_id: userId, follower_id: followerId };
        const response = await fetch(`${process.env.REACT_APP_API_HOST}/api/social/unfollow`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestBody)
        });
        return response.ok;
    } catch (error) {
        console.error('Error unfollowing user:', error);
        throw error;
    }
};

export const createGroup = async (token, userId, groupName) => {
    try {
        const requestBody = {
            name: groupName,
            user_id: userId,
        };

        const response = await fetch(`${process.env.REACT_APP_API_HOST}/api/groups/create`, {
            method: 'PUT',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestBody),
        });

        if (response.ok) {
            return await response.json();
        } else {
            console.error('Failed to create group:', response.statusText);
            return null;
        }
    } catch (error) {
        console.error('Error creating group:', error);
        throw error;
    }
};

{/* API's for MainPage.js */}


{/* API's for FeedsPage.js */}
export const fetchFeedData = async (token) => {
    try {
        const decodedToken = jwtDecode(token);

        const response = await fetch(`${process.env.REACT_APP_API_HOST}/api/social/feed`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ user_id: decodedToken.user_id })
        });

        if (response.ok) {
            const feedJson = await response.json();
            return feedJson.feed || [];
        } else {
            console.error('Failed to fetch feed data. Status:', response.status);
            return [];
        }
    } catch (error) {
        console.error('Error fetching feed data:', error);
        return [];
    }
};

    export const acceptInvite = async (token, user_id, group_id) => {
        try {
            const requestBody = { user_id: user_id, group_id: group_id };
            const response = await fetch(`${process.env.REACT_APP_API_HOST}/api/invite/accept`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(requestBody)
            });
            return response.ok;
        } catch (error) {
            console.error('Error unfollowing user:', error);
            throw error;
        }
    }

export const copyDeck = async (token, copier_id, deck_id) => {
    try {
        const requestBody = { copier_id: copier_id, deck_id: deck_id };
        const response = await fetch(`${process.env.REACT_APP_API_HOST}/api/social/copy`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestBody)
        });
        return response.ok;
    } catch (error) {
        console.error('Error unfollowing user:', error);
        throw error;
    }
}

{/* API's for FeedsPage.js */}




