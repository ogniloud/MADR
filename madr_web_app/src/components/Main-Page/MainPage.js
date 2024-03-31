import React, {useEffect, useRef, useState} from 'react';
import {Link, Route, Routes, useNavigate} from 'react-router-dom';
import Feed from './Feeds/FeedsPage'
import CreateDeck from './Decks/CreateDecks';
import {jwtDecode} from 'jwt-decode';
import './Styles/MainPage.css';
import defaultProfilePicture from './imgs/default-profile-picture.png';
import closeIcon from './imgs/close-circle.png';

const MainPage = () => {
    const [userInfo, setUserInfo] = useState(null);
    const [showPopup, setShowPopup] = useState(false);
    const [followers, setFollowers] = useState([]);
    const [followings, setFollowings] = useState([]);
    const [searchQuery, setSearchQuery] = useState('');
    const [searchResults, setSearchResults] = useState([]);
    const [searchClicked, setSearchClicked] = useState(false);
    const navigate = useNavigate();
    const searchResultsRef = useRef(null);

    useEffect(() => {
        const fetchUserData = async () => {
            const token = localStorage.getItem('token');
            if (token) {
                try {
                    const decodedToken = jwtDecode(token);
                    setUserInfo(decodedToken);

                    const responseFollowers = await fetch('http://localhost:8080/api/social/followers', {
                        method: 'POST',
                        headers: {
                            'Authorization': `Bearer ${token}`,
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({ user_id: decodedToken.user_id })
                    });

                    const dataFollowers = await responseFollowers.json();
                    setFollowers(dataFollowers.user_info || []);

                    const responseFollowings = await fetch('http://localhost:8080/api/social/followings', {
                        method: 'POST',
                        headers: {
                            'Authorization': `Bearer ${token}`,
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({ user_id: decodedToken.user_id })
                    });

                    const dataFollowings = await responseFollowings.json();
                    setFollowings(dataFollowings.user_info || []);
                } catch (error) {
                    console.error('Error fetching user data:', error);
                }
            }
        };

        fetchUserData();
    }, []);

    useEffect(() => {
        const handleOutsideClick = (event) => {
            if (searchResultsRef.current && !searchResultsRef.current.contains(event.target)) {
                setSearchClicked(false);
            }
        };

        document.addEventListener('click', handleOutsideClick);

        return () => {
            document.removeEventListener('click', handleOutsideClick);
        };
    }, []);

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/signin');
    };

    const togglePopup = () => {
        setShowPopup(!showPopup);
    };

    const stopPropagation = (event) => {
        event.stopPropagation();
    };

    const handleSearch = async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/social/search?q=${encodeURIComponent(searchQuery)}`);
            if (response.ok) {
                const searchData = await response.json();
                setSearchResults(searchData.users || []);
                setSearchClicked(true);
            } else {
                setSearchResults([]);
                setSearchClicked(true);
            }
        } catch (error) {
            console.error('Error fetching search results:', error);
        }
    };

    const handleKeyPress = (event) => {
        if (event.key === 'Enter') {
            handleSearch();
        }
    };

    const handleFollow = async (user) => {
        try {
            const isFollowing = followings.some((following) => following.userId === user.ID);
            const token = localStorage.getItem('token');
            if (token) {
                console.log('UserInfo:', userInfo);
                if (userInfo && userInfo.user_id) {
                    if (isFollowing) {
                        console.log('Unfollowing user:', user);
                        const requestBody = { author_id: user.ID, follower_id: userInfo.user_id };
                        console.log('Unfollow request body:', requestBody);
                        const response = await fetch('http://localhost:8080/api/social/unfollow', {
                            method: 'POST',
                            headers: {
                                'Authorization': `Bearer ${token}`,
                                'Content-Type': 'application/json'
                            },
                            body: JSON.stringify(requestBody)
                        });
                        console.log('Unfollow response:', response);
                        if (response.ok) {
                            setFollowings(followings.filter((following) => following.userId !== user.ID));
                        }
                    } else {
                        console.log('Following user:', user);
                        const requestBody = { author_id: user.ID, follower_id: userInfo.user_id };
                        console.log('Follow request body:', requestBody);
                        const response = await fetch('http://localhost:8080/api/social/follow', {
                            method: 'POST',
                            headers: {
                                'Authorization': `Bearer ${token}`,
                                'Content-Type': 'application/json'
                            },
                            body: JSON.stringify(requestBody)
                        });
                        console.log('Follow response:', response);
                        if (response.ok) {
                            setFollowings([...followings, { userId: user.ID, username: user.Username }]);
                        }
                    }
                } else {
                    console.log('User info is not available or missing user_id.');
                }
            }
        } catch (error) {
            console.error('Error following/unfollowing user:', error);
        }
    };



    return (
        <div className="main-page">
            <nav className="upper-part">
                <div className="main-page-search-container">
                    <input
                        type="text"
                        placeholder="Search friends"
                        value={searchQuery}
                        onChange={(e) => setSearchQuery(e.target.value)}
                        onKeyPress={handleKeyPress}
                    />
                    <button className="main-page-search-btn" onClick={handleSearch}>Search</button>
                </div>

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

                            <div className="popup-content-user-details">
                                <p className="follower-button-user-details">Followers: {followers.length}</p>
                                <ul>
                                    {followers.map((follower, index) => (
                                        <li key={index}>{follower.tierName}</li>
                                    ))}
                                </ul>
                                <p className="followings-button-user-details">Followings: {followings.length}</p>
                                <ul>
                                    {followings.map((following, index) => (
                                        <li key={index}>{following.username}</li>
                                    ))}
                                </ul>
                                <div className="group-button-user-details">
                                    <button className="group-dropbtn">Groups</button>
                                    <div className="group-dropdown-content">
                                        {/* Dropdown content here */}
                                    </div>
                                </div>

                            </div>

                            <button className="popup-logout-button" onClick={handleLogout}>
                                Logout
                            </button>
                        </div>
                    )}
                </div>
                <Link className="main-page-feed-button" to="/feed">
                    Feed
                </Link>

            </nav>

            {searchClicked && (
                <div className="main-page-search-results" ref={searchResultsRef}>
                    {searchResults.length > 0 ? (
                        <ul>
                            {searchResults.map((user, index) => (
                                <li key={index}>
                                    <p>Username: {user.Username}</p>
                                    <p>Email: {user.Email}</p>
                                    <button className="main-page-follow-unfollow-button" onClick={() => handleFollow(user)}>
                                        {followings.some((following) => following.userId === user.userId)
                                            ? 'Unfollow'
                                            : 'Follow'}
                                    </button>
                                </li>
                            ))}
                        </ul>
                    ) : (
                        <p>No users found</p>
                    )}
                </div>
            )}

            <Routes>
                <Route path="/create-deck" element={<CreateDeck/>}/>
                <Route path="/feed" element={<Feed/>}/>
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
