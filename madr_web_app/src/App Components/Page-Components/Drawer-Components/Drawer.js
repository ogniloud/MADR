import React from 'react';
import Drawer from '@mui/material/Drawer';
import Button from '@mui/material/Button';
import './Styles/Drawer.css';


const UserDrawer = ({
                        open,
                        onClose,
                        userInfo,
                        handleLogout,
                        followers,
                        followings,
                        handleFollow,
                        handleCreateGroup,
                        handleViewGroups,
                        groups,
                        showGroupDialog,
                        setShowGroupDialog,
                        groupName,
                        setGroupName,
                        groupCreationError,
                        setGroupCreationError,
                        groupCreationSuccess,
                        setGroupCreationSuccess,
                        defaultProfilePicture

                    }) => {
    return (
        <Drawer open={open} onClose={onClose}>
            <div className="drawer-content">
                {/* User info */}
                <div className="drawer-user-info">
                    <div>
                        <img src={defaultProfilePicture} alt="Profile"/>
                        <h2>{userInfo && userInfo.username}</h2>
                    </div>
                    <p className="follower-button">
                        <span className= "follower-button-name">Followers</span>
                        <span className= "follower-button-number">{followers.length}</span>
                    </p>
                    {/* Followings */}
                    <p className="following-button">
                        <span className= "following-button-name">Followings</span>
                        <span className= "following-button-number">{followings.length}</span>
                    </p>

                    {/* Create Group */}
                    <button className="create-group-button" onClick={() => setShowGroupDialog(true)}>Create Group
                    </button>
                    {/* View Groups */}
                    <button className="View-group-button" onClick={handleViewGroups}>View Groups</button>
                    {/* Logout */}
                    <button className="logout-button" onClick={handleLogout}>Logout</button>
                </div>
                {/* Other content  */}
            </div>
        </Drawer>
    );
};

export default UserDrawer;
