import React, { useEffect, useState } from 'react';
import { getGroupsByUserId } from "../API-Components/apiFunctions_groups";
import { useNavigate } from 'react-router-dom';
import './Styles/social_group.css';

const SocialGroup = ({ userId }) => {
    const [groups, setGroups] = useState([]);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const fetchGroups = async () => {
            try {
                console.log('Fetching groups...');
                const fetchedGroups = await getGroupsByUserId(userId);
                console.log('Fetched groups:', fetchedGroups);
                setGroups(fetchedGroups);
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

    if (isLoading) {
        return <div>Loading...</div>;
    }

    return (
        <div className="social-group-container">
            <h2>My Groups</h2>
            {groups.length > 0 ? (
                <div className="group-cards">
                    {groups.map((group) => (
                        <div key={group.group_id} className="group-card">
                            <h3>Group Name: {group.name}</h3>
                            <p>Creator ID: {group.creator_id}</p>
                            <p>Created At: {group.time_created}</p>
                        </div>
                    ))}
                </div>
            ) : (
                <p>You aren't part of any group.</p>
            )}
        </div>
    );
};

export default SocialGroup;