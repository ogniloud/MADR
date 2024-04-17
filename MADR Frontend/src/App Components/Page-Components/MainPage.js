import React, { useEffect, useRef, useState } from 'react';
import { Link, Route, Routes, useNavigate } from 'react-router-dom';
import Feed from './Feeds/FeedsPage';
import CreateDeck from './Decks/CreateDecks';
import { jwtDecode } from 'jwt-decode';
import './Styles/MainPage.css';
import defaultProfilePicture from './resource/default-profile-picture.png';
import closeIcon from './resource/close-circle.png';
import { fetchFollowers, fetchFollowings, searchUsers, followUser, unfollowUser, createGroup } from './API-Components/apiFunctions_main_feeds';
import SocialGroup from "./Social Site Components/social_group";


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
    const [showGroupDialog, setShowGroupDialog] = useState(false);
    const [groupName, setGroupName] = useState('');
    const [groupCreationError, setGroupCreationError] = useState('');
    const [groupCreationSuccess, setGroupCreationSuccess] = useState('');
    const [groups, setGroups] = useState([]);

    useEffect(() => {
        const fetchUserData = async () => {
            const token = localStorage.getItem('token');
            if (token) {
                try {
                    const decodedToken = jwtDecode(token);
                    setUserInfo(decodedToken);

                    const followersData = await fetchFollowers(token, decodedToken.user_id);
                    setFollowers(followersData.user_info || []);

                    const followingsData = await fetchFollowings(token, decodedToken.user_id);
                    setFollowings(followingsData.user_info || []);
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
            const token = localStorage.getItem('token');
            if (token) {
                const searchData = await searchUsers(token, searchQuery);
                console.log('Search Data:', searchData);
                // Update search results with follow status
                const updatedSearchResults = searchData.map(user => {
                    // Check if the user is already being followed
                    const isFollowing = followings.some(following => following.userId === user.ID);
                    return {
                        ...user,
                        isFollowing: isFollowing
                    };
                });
                console.log('Updated Search Results:', updatedSearchResults);
                setSearchResults(updatedSearchResults);
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
            const token = localStorage.getItem('token');
            if (token && userInfo && userInfo.user_id) {
                if (user.isFollowing) {
                    // Unfollow user
                    console.log('Unfollowing user:', user);
                    const response = await unfollowUser(token, user.ID, userInfo.user_id);
                    console.log('Unfollow response:', response);
                    if (response) {
                        // Remove the unfollowed user from the followings state
                        setFollowings(followings.filter((following) => following.userId !== user.ID));
                        // Update the search results to reflect the unfollow action
                        setSearchResults(searchResults.map(result => ({
                            ...result,
                            isFollowing: result.ID === user.ID ? !user.isFollowing : result.isFollowing
                        })));
                        console.log('Updated search results after unfollow:', searchResults);
                    }
                } else {
                    // Follow user
                    console.log('Following user:', user);
                    const response = await followUser(token, user.ID, userInfo.user_id);
                    console.log('Follow response:', response);
                    if (response) {
                        // Add the followed user to the followings state
                        setFollowings([...followings, { userId: user.ID, username: user.Username }]);
                        // Update the search results to reflect the follow action
                        setSearchResults(searchResults.map(result => ({
                            ...result,
                            isFollowing: result.ID === user.ID ? !user.isFollowing : result.isFollowing
                        })));
                        console.log('Updated search results after follow:', searchResults);
                    }
                }
            } else {
                console.log('Token or user info is missing.');
            }
        } catch (error) {
            console.error('Error following/unfollowing user:', error);
        }
    };









    const handleCreateGroup = async () => {
        try {
            const token = localStorage.getItem('token');
            const userId = userInfo.user_id;

            if (token && userId && groupName.trim() !== '') {
                const response = await createGroup(token, userId, groupName);

                if (response && response.group_id) {
                    setShowGroupDialog(false);
                    setGroupCreationSuccess('Group created successfully!');
                    setGroupCreationError('');
                } else {
                    // Group creation failed
                    setShowGroupDialog(true);
                    setGroupCreationError('Failed to create group');
                    setGroupCreationSuccess('');
                }
            } else {

                setShowGroupDialog(true);
                setGroupCreationError('Invalid group name or user ID');

                setGroupCreationSuccess('');
            }
        } catch (error) {
            console.error('Error creating group:', error);
            setShowGroupDialog(true);
            setGroupCreationError('An error occurred while creating the group');
            setGroupCreationSuccess('');
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
                                    <img src={closeIcon} alt="Close" />
                                </button>
                            </div>

                            <div className="popup-content-user-profile">
                                <img src={defaultProfilePicture} alt="Profile" />
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
                                    {/* Add button to open group creation dialog */}
                                    <button className="create-group-button" onClick={() => setShowGroupDialog(true)}>Create Group</button>
                                </div>

                                <div>
                                    <Link className="view-groups-button" to="/social_group">
                                        View Groups
                                    </Link>
                                </div>

                            </div>



                            {/* Dialog box for group creation */}
                            {showGroupDialog && (
                                <div className="popup-overlay">
                                    <div className="creating-group-popup">
                                        <input className="Create-group-input"
                                            type="text"
                                            placeholder="Enter group name"
                                            value={groupName}
                                            onChange={(e) => setGroupName(e.target.value)}
                                        />
                                        <button className="Create-group-button" onClick={handleCreateGroup}>Create</button>
                                        <button className="Cancel-group-button" onClick={() => setShowGroupDialog(false)}>Cancel</button>

                                        {groupCreationError && <p>{groupCreationError}</p>}
                                        {groupCreationSuccess && !groupCreationError && <p>{groupCreationSuccess}</p>}



                                    </div>
                                </div>
                            )}


                            <button className="popup-logout-button" onClick={handleLogout}>
                                Logout
                            </button>
                        </div>
                    )}




                </div>
                <Link className="main-page-feed-button" to="/feed">
                    Feed
                </Link>

                {/* Render the SocialGroup component only if the user has groups */}
                {groups.length > 0 && (
                    <SocialGroup userId={userInfo && userInfo.user_id} />
                )}



            </nav>

            {/* Search results */}
            {searchClicked && (
                <div className="main-page-search-results" ref={searchResultsRef}>
                    {searchResults.length > 0 ? (
                        <ul>
                            {searchResults.map((user, index) => (
                                <li key={index}>
                                    <p>Username: {user.Username}</p>
                                    <p>Email: {user.Email}</p>
                                    {/* Display "Follow" or "Unfollow" based on user's follow status */}
                                    {user.isFollowing ? (
                                        <React.Fragment>
                                            <button className="main-page-follow-unfollow-button" onClick={() => handleFollow(user)}>
                                                Unfollow
                                            </button>
                                            {/* Display message indicating user is already followed */}
                                            <p className="user-following-message">You are already following this user</p>
                                        </React.Fragment>
                                    ) : (
                                        <button className="main-page-follow-unfollow-button" onClick={() => handleFollow(user)}>
                                            Follow
                                        </button>
                                    )}
                                </li>
                            ))}
                        </ul>
                    ) : (
                        <p>No users found</p>
                    )}

                </div>
            )}




            {/* Routes and lower-part section */}
            <Routes>
                <Route path="/create-deck" element={<CreateDeck />} />
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
