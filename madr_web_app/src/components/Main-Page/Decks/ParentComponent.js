
import React, { useState, useEffect } from 'react';
import CreateDecks from './CreateDecks';
import axios from 'axios';

const ParentComponent = () => {
    const [userData, setUserData] = useState({ user_id: null, username: '' });
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/api/user');
                setUserData({
                    user_id: response.data.id,
                    username: response.data.username,
                });
                setLoading(false);
            } catch (error) {
                console.error('Error fetching user data:', error.message);
                setLoading(false);
            }
        };

        fetchUserData();
    }, []);

    if (loading) {
        return <div>Loading...</div>;
    }

    return <CreateDecks userId={userData.user_id} />;
};

export default ParentComponent;
