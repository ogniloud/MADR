import React, { useState, useEffect } from 'react';
import { Link, Route, Routes, useNavigate } from 'react-router-dom';
import Feed from './Feeds/FeedsPage';
import CreateDeck from './Decks/CreateDecks';
import { jwtDecode } from 'jwt-decode';
import './Styles/MainPage.css';
import defaultProfilePicture from './imgs/default-profile-picture.png';
import closeIcon from './imgs/close-circle.png';

const MainPage = () => {
    const [userInfo, setUserInfo] = useState(null);
    const [showPopup, setShowPopup] = useState(false);
    const [followers, setFollowers] = useState([]);
    const [followings, setFollowings] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            const decodedToken = jwtDecode(token);
            setUserInfo(decodedToken);

            // Fetch followers data
            fetch('http://localhost:8080/api/social/followers', {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ userId: decodedToken.userId })
            })
                .then(response => response.json())
                .then(data => {
                    setFollowers(data.followers);
                })
                .catch(error => {
                    console.error('Error fetching followers:', error);
                });

            // Fetch followings data
            fetch('http://localhost:8080/api/social/followings', {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ userId: decodedToken.userId })
            })
                .then(response => response.json())
                .then(data => {
                    setFollowings(data.followings);
                })
                .catch(error => {
                    console.error('Error fetching followings:', error);
                });
        }
    }, []);

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/signin');
    };

    const togglePopup = () => {
        setShowPopup(!showPopup);
    };

    const handleSearch = () => {
        // Logic for searching friends
        console.log('Searching friends...');
    };

    const handleKeyPress = (event) => {
        if (event.key === 'Enter') {
            handleSearch();
        }
    };

    const stopPropagation = (event) => {
        event.stopPropagation();
    };

    return (
        <div className="main-page">
            <nav className="upper-part">
                <div className="user-info" onClick={togglePopup}>
                    <h2 className="title-user-name">{userInfo && userInfo.username}</h2>
                    {showPopup && (
                        <div className="popup-user-profile" onClick={stopPropagation}>
                            <div className="popup-header">
                                <h3 className="popup-user-name">{userInfo && userInfo.username}</h3>
                                <button className="popup-close-btn" onClick={togglePopup}>
                                    <img src={closeIcon} alt="Close"/>
                                </button>
                            </div>


                            <div className="popup-content-user-profile">
                                <img src={defaultProfilePicture} alt="Profile"/>
                            </div>

                            <div className={"popup-content-user-details"}>

                                <p className="follower-button-user-details">Followers: {followers.length}</p>
                                {/* Display followers' tier names */}
                                <ul>
                                    {followers.map((follower, index) => (
                                        <li key={index}>{follower.tierName}</li>
                                    ))}
                                </ul>
                                <p className="followings-button-user-details">Followings: {followings.length}</p>
                                {/* Display followings' names */}
                                <ul>
                                    {followings.map((following, index) => (
                                        <li key={index}>{following.username}</li>
                                    ))}
                                </ul>
                                <div className="group-button-user-details">
                                    <button className="group-dropbtn">Groups</button>
                                    <div className="group-dropdown-content">
                                        {/* Dropdown content here  - need to add them throught api endpoint, still not ready yet have to add them when they are ready -*/}
                                    </div>
                                </div>

                            </div>
                            <div className="popup-search-container">
                                <input type="text" placeholder="Search friends" onKeyPress={handleKeyPress}/>
                                <button className="popup-search-btn" onClick={handleSearch}>Search</button>
                            </div>
                        </div>
                    )}




                </div>
                <button className="logout-button" onClick={handleLogout}>
                    Logout
                </button>
            </nav>

            <Routes>
                <Route path="/create-deck" element={<CreateDeck/>}/>
                <Route path="/feed" element={<Feed />} />
            </Routes>

            <div className="lower-part">
                <Link className="create-deck" to="/create-deck">
                    Create Deck
                </Link>
                <Link className="all-cards" to="/decks">
                    All Decks
                </Link>
            </div>
        </div>
    );
};

export default MainPage;
