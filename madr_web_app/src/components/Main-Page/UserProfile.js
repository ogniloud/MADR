import React from 'react';

const UserProfile = ({ userInfo }) => {
    return (
        <div className="user-profile">
            <h2>User Profile</h2>
            {userInfo ? (
                <div>
                    <p>User ID: {userInfo.id}</p>
                    <p>Username: {userInfo.username}</p>
                    <p>Email: {userInfo.email}</p>
                </div>
            ) : (
                <p>Loading user data...</p>
            )}
        </div>
    );
};

export default UserProfile;
