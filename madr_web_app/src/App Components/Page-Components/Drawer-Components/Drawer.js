import React from 'react';
import Drawer from '@mui/material/Drawer';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
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
    const [anchorElFollowers, setAnchorElFollowers] = React.useState(null);
    const [anchorElFollowings, setAnchorElFollowings] = React.useState(null);

    const handleClickFollowers = (event) => {
        setAnchorElFollowers(event.currentTarget);
    };

    const handleClickFollowings = (event) => {
        setAnchorElFollowings(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorElFollowers(null);
        setAnchorElFollowings(null);
    };

    return (
        <Drawer open={open} onClose={onClose}>
            <div className="drawer-content">
                {/* User info */}
                <div className="drawer-user-info">
                    <div>
                        <img src={defaultProfilePicture} alt="Profile"/>
                        <h2>{userInfo && userInfo.username}</h2>
                    </div>


                    <div className="dropdown-container">
                        <div className="dropdown-container-followers">
                            <button className="followers-button" onClick={handleClickFollowers}>
                                Followers ({followers.length})
                            </button>
                            <Menu
                                className="menu-dropdown"
                                anchorEl={anchorElFollowers}
                                open={Boolean(anchorElFollowers)}
                                onClose={handleClose}
                                anchorOrigin={{
                                    vertical: 'bottom',
                                    horizontal: 'left',
                                }}
                                transformOrigin={{
                                    vertical: 'top',
                                    horizontal: 'left',
                                }}
                            >
                                {followers.map((follower, index) => {
                                    return (
                                        <MenuItem key={index} onClick={handleClose}>
                                            {follower.Username}
                                        </MenuItem>
                                    );
                                })}
                            </Menu>
                        </div>
                        {/* Followings dropdown */}
                        <div className="dropdown-container-followings">
                            <button className="followings-button" onClick={handleClickFollowings}>
                                Following ({followings.length})
                            </button>
                            <Menu
                                className="menu-dropdown"
                                anchorEl={anchorElFollowings}
                                open={Boolean(anchorElFollowings)}
                                onClose={handleClose}
                                anchorOrigin={{
                                    vertical: 'bottom',
                                    horizontal: 'left',
                                }}
                                transformOrigin={{
                                    vertical: 'top',
                                    horizontal: 'left',
                                }}
                            >
                                {followings.map((following, index) => {
                                    return (
                                        <MenuItem key={index} onClick={handleClose}>
                                            {following.Username}
                                        </MenuItem>
                                    );
                                })}
                            </Menu>
                        </div>
                    </div>


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
