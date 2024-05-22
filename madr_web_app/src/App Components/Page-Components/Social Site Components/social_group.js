import './Styles/social_group.css'
import { useNavigate } from "react-router-dom";
import React, { useEffect, useState } from 'react';
import { getCreatedGroupsByUserId, getGroupsByUserId } from "../API-Components/apiFunctions_groups";
import { jwtDecode } from "jwt-decode";
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
                const fetchedGroups = await getGroupsByUserId(userId);
                const fetchedCreatedGroups = await getCreatedGroupsByUserId(userId);
                setGroups(fetchedGroups);
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
        navigate('/mainpage'); // Use navigate to navigate to 'MainPage'
    };

    const renderGroupCards = (groupList, messageIfEmpty) => {
        if (groupList.length > 0) {
            return (
                <div className="group-cards">
                    {groupList.map((group) => (
                        <div key={group.group_id} className="group-card">
                            <h3>Group Name: {group.name}</h3>
                            <p>Creator ID: {group.creator_id}</p>
                            <p>Created At: {group.time_created}</p>
                            <button onClick={()=>{navigate('/social_group/'+group.group_id)}}>Invite someone...</button>
                        </div>
                    ))}
                </div>
            );
        } else {
            return <p>{messageIfEmpty}</p>;
        }
    };

    return (
        <div className="social-group-container">
            <div className="upper-part">
                <h2 className="social-group-title">My Groups</h2>
            </div>

            <div className="groups-container" style={{ '--card-height': '50px', '--num-cards': groups.length + createdGroups.length }}>
                {/* Add the "Return to Home" button */}
                <button className="social-group-home-button" onClick={returnToHome}>Return to Home</button>

                {renderGroupCards(groups.concat(createdGroups), "No groups.")}
            </div>
        </div>
    );
};

export default SocialGroup;
