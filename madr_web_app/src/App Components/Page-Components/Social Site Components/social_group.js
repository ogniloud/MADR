
import './Styles/social_group.css'
import {useNavigate} from "react-router-dom"; // Import useNavigation
import React, {useEffect, useState} from 'react';
import {getCreatedGroupsByUserId, getGroupsByUserId} from "../API-Components/apiFunctions_groups";
import {jwtDecode} from "jwt-decode";
import InfiniteScroll from "react-infinite-scroll-component";

const SocialGroup = () => {
    const navigate = useNavigate();
    const decodedToken = jwtDecode(localStorage.getItem('token'));
    const userId = parseInt(decodedToken.user_id);
    const [groups, setGroups] = useState([]);
    const [createdGroups, setCreatedGroups] = useState([]);
    const [isLoading, setIsLoading] = useState(true);


    useEffect(() => {
        const fetchGroups = async () => {
            try {
                console.log('Fetching groups...', userId);
                const fetchedGroups = await getGroupsByUserId(parseInt(userId));
                console.log('Fetched groups:', fetchedGroups);
                setGroups(fetchedGroups);

                console.log('Fetching created groups...', userId);
                const fetchedCreatedGroups = await getCreatedGroupsByUserId(parseInt(userId));
                console.log('Fetched created groups:', fetchedCreatedGroups);
                setCreatedGroups(fetchedCreatedGroups);
                setIsLoading(false);
            } catch (error) {
                console.error('Error fetching groups:', error);
                setIsLoading(false);
            }
        };

        fetchGroups();
    }, [userId]);

    console.log('Groups:', groups);
    console.log('Is loading:', isLoading);

    const returnToHome = () => {
        navigation.navigate('MainPage'); // Use navigation.navigate to navigate to 'MainPage'
    };




    return (
        <div className="social-group-container">
            <div className="upper-part">
                <h2 className="social-group-title">My Groups</h2>
            </div>

            <div className="groups-container">
                {/* Add the "Return to Home" button */}
                <button className="social-group-home-button" onClick={returnToHome}>Return to Home</button>

                {groups.length > 0 && (
                    <div className="group-cards">
                        {groups.map((group) => (
                            <div key={group.group_id} className="group-card">
                                <h3>Group Name: {group.name}</h3>
                                <p>ID: {group.group_id}</p>
                                <p>Creator ID: {group.creator_id}</p>
                                <p>Created At: {group.time_created}</p>
                            </div>
                        ))}
                    </div>
                )}

                {createdGroups.length > 0 ? (
                    <div className="group-cards">
                        {createdGroups.map((group) => (
                            <div key={group.group_id} className="group-card">
                                <h3>Group Name: {group.name}</h3>
                                <p>ID: {group.group_id}</p>
                                <p>Creator ID: {group.creator_id}</p>
                                <p>Created At: {group.time_created}</p>
                            </div>
                        ))}
                    </div>
                ) : (
                    groups.length === 0 && <p>No groups.</p>
                )}
            </div>

            <h2>My Groups</h2>
            {groups.length > 0 ? (
                <div className="group-cards">
                    {groups.map((group) => (
                        <div key={group.group_id} className="group-card">
                            <h3>Group Name: {group.name}</h3>
                            <p>ID: {group.group_id}</p>
                            <p>Creator ID: {group.creator_id}</p>
                            <p>Created At: {group.time_created}</p>
                        </div>
                    ))}
                </div>
            ) : (
                <p>You aren't part of any group.</p>
            )}
            <h3>Your created groups</h3>
            {createdGroups.length > 0 ? (
                <div className="group-cards">
                    {createdGroups.map((group) => (
                        <div key={group.group_id} className="group-card">
                            <h3>Group Name: {group.name}</h3>
                            <p>ID: {group.group_id}</p>
                            <p>Creator ID: {group.creator_id}</p>
                            <p>Created At: {group.time_created}</p>
                        </div>
                    ))}
                </div>
            ) : (
                <p>No groups.</p>
            )}
        </div>
    );
};

export default SocialGroup;
