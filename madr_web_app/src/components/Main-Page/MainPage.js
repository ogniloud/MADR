import React, { useState, useEffect } from 'react';
import { Link, Route, Routes, useNavigate } from 'react-router-dom';
import Feed from './Feeds/FeedsPage';
import CreateDeck from './Decks/CreateDecks';
import { jwtDecode } from 'jwt-decode';
import './Styles/MainPage.css'

const MainPage = () => {
    const [userInfo, setUserInfo] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            const decodedToken = jwtDecode(token);
            setUserInfo(decodedToken);
        }
    }, []);

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/signin');
    };

    return (
        <div className="main-page">
            <nav className="upper-part">
                <div className="user-info">
                    <h2 className="title-user-name" >{userInfo && userInfo.username}</h2>
                    <button className="logout-button" onClick={handleLogout}>Logout</button>
                </div>
            </nav>

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
