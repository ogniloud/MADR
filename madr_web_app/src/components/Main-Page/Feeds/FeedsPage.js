import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {fetchFeedData} from "../APIs/apiFunctions_main_feeds";
import './FeedPage.css';

const FeedsPage = () => {
    const [feedData, setFeedData] = useState([]);
    const [userInfo, setUserInfo] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            fetchFeed(token);
        }
    }, []);

    const fetchFeed = async (token) => {
        const data = await fetchFeedData(token);
        setFeedData(data);
    };

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
