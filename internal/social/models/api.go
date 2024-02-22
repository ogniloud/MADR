package models

import "github.com/ogniloud/madr/internal/models"

// GetCreatedGroupsByUserIdRequest is a struct that defines the request body for
// loading groups the user is a creator of.
//
//swagger:model getCreatedGroupsByUserId
type GetCreatedGroupsByUserIdRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`
}

// GetCreatedGroupsByUserIdResponse is a struct that defines the response body for
// loading groups the user is a creator of.
type GetCreatedGroupsByUserIdResponse struct {
	// Groups is a map from group id to group config.
	//
	// required: true
	Groups Groups `json:"groups"`
}

// GetGroupsByUserIdRequest is a struct that defines the request body for
// loading groups the user is a member of.
//
//swagger:model getGroupsByUserIdRequest
type GetGroupsByUserIdRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`
}

// GetGroupsByUserIdResponse is a struct that defines the request body for
// loading groups the user is a member of.
type GetGroupsByUserIdResponse struct {
	// Groups is a map from group id to group config.
	//
	// required: true
	Groups Groups `json:"groups"`
}

// GetUsersByGroupIdRequest is a struct that defines the request body for
// loading members of a group.
//
//swagger:model getUsersByGroupIdRequest
type GetUsersByGroupIdRequest struct {
	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 189
	GroupId GroupId `json:"group_id"`
}

// GetUsersByGroupIdResponse is a struct that defines the request body for
// loading members of a group.
type GetUsersByGroupIdResponse struct {
	// Members is a map from user id to member info.
	//
	// required: true
	Members Members `json:"members"`
}

// GetGroupByGroupIdRequest is a struct that defines the request body for
// loading a config of the group.
//
//swagger:model getGroupByGroupIdRequest
type GetGroupByGroupIdRequest struct {
	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 189
	GroupId GroupId `json:"group_id"`
}

// GetGroupByGroupIdResponse is a struct that defines the request body for
// loading members of a group.
type GetGroupByGroupIdResponse struct {
	// GroupConfig is a struct containing the info about a group.
	//
	// required: true
	GroupConfig GroupConfig `json:"group_config"`
}

// GetDecksByGroupIdRequest is a struct that defines the request body for
// loading decks of a group.
//
//swagger:model getDecksByGroupIdRequest
type GetDecksByGroupIdRequest struct {
	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 189
	GroupId GroupId `json:"group_id"`
}

// GetDecksByGroupIdResponse is a struct that defines the request body for
// loading decks of a group.
type GetDecksByGroupIdResponse struct {
	// Decks is an array of deck ids.
	//
	// required: true
	Decks []DeckId `json:"decks"`
}

// GetInvitesByGroupIdRequest is a struct that defines the request body for
// loading invites to a group.
//
//swagger:model getInvitesByGroupIdRequest
type GetInvitesByGroupIdRequest struct {
	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 189
	GroupId GroupId `json:"group_id"`
}

// GetInvitesByGroupIdResponse is a struct that defines the request body for
// loading invites to a group.
type GetInvitesByGroupIdResponse struct {
	// Invites is a map from user id to invite information.
	//
	// required: true
	Invites map[UserId]InviteInfo `json:"invites_from_group"`
}

// GetInvitesByUserIdRequest is a struct that defines the request body for
// loading invites to a user.
//
//swagger:model getInvitesByUserIdRequest
type GetInvitesByUserIdRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`
}

// GetInvitesByUserIdResponse is a struct that defines the request body for
// loading invites to a user.
type GetInvitesByUserIdResponse struct {
	// Invites is a map from group id to invite information.
	//
	// required: true
	Invites map[GroupId]InviteInfo `json:"invites_to_user"`
}

// CreateGroupRequest is a struct that defines the request body for
// creating a new group.
//
//swagger:model createGroupRequest
type CreateGroupRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// Name is the new group's name
	//
	// required: true
	// example: Eduard
	Name string `json:"name"`
}

// CreateGroupResponse is a struct that defines the request body for
// creating a new group.
type CreateGroupResponse struct {
	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 189
	GroupId GroupId `json:"group_id"`
}

// DeleteGroupRequest is a struct that defines the request body for
// deleting a group.
//
//swagger:model deleteGroupRequest
type DeleteGroupRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 189
	GroupId GroupId `json:"group_id"`
}

// DeleteGroupResponse is a struct that defines the request body for
// creating a new group.
type DeleteGroupResponse struct{}

// AcceptInviteRequest is a struct that defines the request body for
// accepting an invite to group.
//
//swagger:model acceptInviteRequest
type AcceptInviteRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 189
	GroupId GroupId `json:"group_id"`
}

// AcceptInviteResponse is a struct that defines the request body for
// sending an invite to group.
type AcceptInviteResponse struct{}

// SendInviteRequest is a struct that defines the request body for
// sending an invite to group.
//
//swagger:model sendInviteRequest
type SendInviteRequest struct {
	// CreatorId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	CreatorId UserId `json:"creator_id"`

	// CreatorId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	InviteeId UserId `json:"invitee_id"`

	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 189
	GroupId GroupId `json:"group_id"`
}

// SendInviteResponse is a struct that defines the request body for
// sending an invite to group.
type SendInviteResponse struct{}

// ShareAllGroupDecksRequest is a struct that defines the request body for
// sharing all the group decks to a new user.
//
//swagger:model shareAllGroupDecksRequest
type ShareAllGroupDecksRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 189
	GroupId GroupId `json:"group_id"`
}

// ShareAllGroupDecksResponse is a struct that defines the request body for
// sharing all the group decks to a new user.
type ShareAllGroupDecksResponse struct{}

// DeepCopyDeckRequest is a struct that defines the request body for
// deep copying a deck (with all its contents recursively copied).
//
//swagger:model deepCopyDeckRequest
type DeepCopyDeckRequest struct {
	// CopierId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	CopierId UserId `json:"copier_id"`

	// DeckId is an ID of the deck to copy.
	//
	// required: true
	// example: 189
	DeckId DeckId `json:"deck_id"`
}

// DeepCopyDeckResponse is a struct that defines the request body for
// deep copying a deck (with all its contents recursively copied).
type DeepCopyDeckResponse struct {
	// DeckId is an ID of the newly copied deck.
	//
	// required: true
	// example: 189
	DeckId DeckId `json:"deck_id"`
}

// FollowersRequest contains an id of the user to get their followers.
type FollowersRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`
}

// FollowersResponse contains ids and names of the followers.
type FollowersResponse struct {
	UserInfo []models.UserInfo `json:"user_info"`
}

// FollowingsRequest contains an id of the user to get their followings.
type FollowingsRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`
}

// FollowingsResponse contains ids and names of the followings.
type FollowingsResponse struct {
	UserInfo []models.UserInfo `json:"user_info"`
}

// FollowRequest contains an id of the user that following to author
type FollowRequest struct {
	// FollowerId is an ID of the user following the author.
	//
	// required: true
	// example: 189
	FollowerId UserId `json:"follower_id"`

	// AuthorId is an ID of the supplier of content.
	//
	// required: true
	// example: 189
	AuthorId UserId `json:"author_id"`
}

// FollowResponse contains ids and names of the followings.
type FollowResponse struct{}

// UnfollowRequest contains an id of the user that unfollowing from the author
type UnfollowRequest struct {
	// FollowerId is an ID of the follower.
	//
	// required: true
	// example: 189
	FollowerId UserId `json:"follower_id"`

	// AuthorId is an ID of the supplier of content.
	//
	// required: true
	// example: 183
	AuthorId UserId `json:"author_id"`
}

// UnfollowResponse contains ids and names of the followings.
type UnfollowResponse struct{}