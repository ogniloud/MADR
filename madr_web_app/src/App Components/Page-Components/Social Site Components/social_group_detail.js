import React, {useEffect, useState} from 'react';
import {useNavigate, useParams} from 'react-router-dom';
import {jwtDecode} from "jwt-decode";
import {followersNotJoined, sendInvite} from "../API-Components/apiFunctions_groups";


const SocialGroupDetail = () =>{
    const {group_id} = useParams();
    const decodedToken = jwtDecode(localStorage.getItem('token'));
    const userId = parseInt(decodedToken.user_id);
    const navigate = useNavigate();
    const [listFollowers, setListFollowers] = useState([]);

    useEffect(() => {
        fetchFollowers()
    }, [])

    const fetchFollowers = () => {
        followersNotJoined(userId, parseInt(group_id), decodedToken).then((response) => {
            setListFollowers(response.followers)
            console.log(response)
        })
    }

    const returnToHome = () => {
        navigate('/social_group');
    };

    return <div>
        <button onClick={returnToHome} >Return to back</button>
        <h3>Choose a member to invite</h3>
        {listFollowers.map((item, index) => (
            <div onClick={() => {
                sendInvite(userId, parseInt(item.follower_id), parseInt(group_id), decodedToken)
            }} className="deck-details-flashcard flashcard-link">
                {item.follower_name}
            </div>
        ))}
    </div>
};

export default SocialGroupDetail