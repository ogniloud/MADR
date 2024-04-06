import React, {useEffect, useRef, useState} from 'react';
import {Link, Route, Routes, useNavigate} from 'react-router-dom';
import {jwtDecode} from 'jwt-decode';


{/* API's for MainPage.js */}

export const fetchFollowers = async (token, userId) => {
    try {
        const response = await fetch('http://localhost:8080/api/social/followers', {
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
        const response = await fetch('http://localhost:8080/api/social/followings', {
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
        const response = await fetch(`http://localhost:8080/api/social/search?q=${encodeURIComponent(query)}`, {
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
        const response = await fetch('http://localhost:8080/api/social/follow', {
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
        const response = await fetch('http://localhost:8080/api/social/unfollow', {
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

{/* API's for MainPage.js */}


{/* API's for FeedsPage.js */}
export const fetchFeedData = async (token) => {
    try {
        const decodedToken = jwtDecode(token);

        const response = await fetch('http://localhost:8080/api/social/feed', {
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


{/* API's for FeedsPage.js */}




