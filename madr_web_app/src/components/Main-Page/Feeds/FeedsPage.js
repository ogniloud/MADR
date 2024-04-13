import React, {useEffect, useState} from 'react';
import {useNavigate} from 'react-router-dom';
import {acceptInvite, copyDeck, fetchFeedData} from "../APIs/apiFunctions_main_feeds";
import './FeedPage.css';
import {jwtDecode} from "jwt-decode";

const FeedsPage = () => {
    const [feedData, setFeedData] = useState([]);
    const [userInfo, setUserInfo] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            fetchFeed(token);

            const decodedToken = jwtDecode(token);
            setUserInfo(decodedToken);
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
                            <div>
                                <p>You received an invite from {item.invite_data.invitee_name} to
                                    group {item.invite_data.group_name}</p>
                                <button onClick={() => acceptInvite(
                                    null,
                                        userInfo.id,
                                        item.invite_data.group_id
                                    )} className="feed-button">ACCEPT</button>
                            </div>
                        )}
                        {item.type === 'shared_from_group_data' && (
                            <div>
                                <p>New deck "{item.shared_from_group_data.deck_name}" added to
                                    group {item.shared_from_group_data.group_name}</p>
                                <button onClick={() => acceptInvite(
                                    null,
                                    userInfo.id,
                                    item.shared_from_group_data.deck_id
                                )} className="feed-button">ACCEPT</button>
                            </div>
                        )}
                        {item.type === 'shared_from_following_data' && (
                            <div>
                                <p>New deck "{item.shared_from_following_data.deck_name}" shared by
                                    user {item.shared_from_following_data.author_name}</p>
                                <button onClick={() => copyDeck(
                                    null,
                                    userInfo.id,
                                    item.shared_from_following_data.deck_id
                                )} className={"feed-button"}>COPY</button>
                            </div>
                        )}
                        {item.type === 'following_subscribed_data' && (
                            <div className="feed-page-inside">
                                <p>{item.following_subscribed_data.follower_name} is followed to {item.following_subscribed_data.author_name}</p>
                            </div>
                        )}
                    </div>
                ))}
            </div>
            <button onClick={returnToHome}>Return to Home</button>
        </div>
    );
};

export default FeedsPage;
