import React, { useEffect, useState } from 'react';
import { jwtDecode } from 'jwt-decode';
import './FeedPage.css';

const FeedsPage = () => {
    const [feedData, setFeedData] = useState([]);
    const [userInfo, setUserInfo] = useState(null);

    useEffect(() => {
        const fetchFeedData = async () => {
            const token = localStorage.getItem('token');
            console.log('Token:', token); // Debugging: Log token
            if (token) {
                try {
                    // Decode token to get user_id
                    const decodedToken = jwtDecode(token);
                    console.log('Decoded token:', decodedToken);
                    setUserInfo(decodedToken);
                    console.log('User ID:', decodedToken.user_id);

                    // Fetch feed data
                    const response = await fetch('http://localhost:8080/api/social/feed', {
                        method: 'POST',
                        headers: {
                            'Authorization': `Bearer ${token}`,
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({ UserId: decodedToken.user_id })
                    });

                    console.log('Fetch response:', response); // Debugging: Log fetch response

                    if (response.ok) {
                        const feedJson = await response.json();
                        console.log('Fetched feed data:', feedJson); // Debugging: Log fetched feed data
                        setFeedData(feedJson.feed || []);
                    } else {
                        console.error('Failed to fetch feed data. Status:', response.status);
                        // You can provide better error handling here, e.g., set an error state
                    }
                } catch (error) {
                    console.error('Error fetching feed data:', error);
                    // You can provide better error handling here, e.g., set an error state
                }
            } else {
                console.error('Token not found in localStorage');
                // You can provide better error handling here, e.g., set an error state
            }
        };

        fetchFeedData();
    }, []);

    const renderFeedItem = (item, index) => {
        switch (item.type) {
            case 0:
                return (
                    <div key={index}>
                        {/* Render invite data */}
                        <p>{item.data.invitee_name} received an invite from {item.data.invitee_id} to group {item.data.group_name}</p>
                    </div>
                );
            case 1:
                return (
                    <div key={index}>
                        <p>New deck "{item.data.deck_name}" added to group {item.data.group_name}</p>
                    </div>
                );
            default:
                return null;
        }
    };

    return (
        <div className="feeds-page">
            <h1>Feed</h1>
            {feedData.map((item, index) => renderFeedItem(item, index))}
        </div>
    );
};

export default FeedsPage;
