import React, { useEffect, useState } from 'react';
import { jwtDecode } from 'jwt-decode';
import { useNavigate } from 'react-router-dom';
import './FeedPage.css';

const FeedsPage = () => {
    const [feedData, setFeedData] = useState([]);
    const [userInfo, setUserInfo] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchFeedData = async () => {
            const token = localStorage.getItem('token');
            if (token) {
                try {
                    const decodedToken = jwtDecode(token);
                    setUserInfo(decodedToken);

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
                        setFeedData(feedJson.feed || []);
                    } else {
                        console.error('Failed to fetch feed data. Status:', response.status);
                    }
                } catch (error) {
                    console.error('Error fetching feed data:', error);
                }
            }
        };

        fetchFeedData();
    }, []);

    const returnToHome = () => {
        navigate('/mainpage');
    };

    return (
        <div className="feeds-page">
            <div className="upper-part"></div>
            <p className="feed-page-title">Feed</p>
            <div className="feed-items-list">
                {feedData.map((item, index) => (
                    <div key={index} className="feed-item">
                        {/* Render feed item based on item.type */}
                        {item.type === 'invite_data' && (
                            <p>You received an invite from {item.invite_data.invitee_name} to group {item.invite_data.group_name}</p>
                        )}
                        {item.type === 'shared_from_group_data' && (
                            <p>New deck "{item.shared_from_group_data.deck_name}" added to group {item.shared_from_group_data.group_name}</p>
                        )}
                        {item.type === 'shared_from_following_data' && (
                            <p>New deck "{item.shared_from_following_data.deck_name}" shared by user {item.shared_from_following_data.author_name}</p>
                        )}
                        {item.type === 'following_subscribed_data' && (
                            <div className="feed-page-inside">
                                <p>{item.following_subscribed_data.follower_name} is following you.</p>
                                <p>You are following {item.following_subscribed_data.author_name}</p>
                            </div>
                        )}
                    </div>
                ))}
            </div>
            <p className="home-button" onClick={returnToHome}>Return to Home</p>
        </div>
    );
};

export default FeedsPage;
