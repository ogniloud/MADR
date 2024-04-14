package models

import (
	"encoding/json"
	"fmt"
	"strings"

	cardmodels "github.com/ogniloud/madr/internal/flashcards/models"
	"github.com/ogniloud/madr/internal/models"
)

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
	Groups []GroupConfig `json:"groups"`
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
	Groups []GroupConfig `json:"groups"`
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
	// Decks is a list of decks.
	//
	// required: true
	Decks []cardmodels.DeckConfig `json:"decks"`
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
//
//swagger:model getFollowersRequest
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
//
//swagger:model getFollowingsRequest
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
//
//swagger:model followRequest
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

// ShareGroupDeckRequest contains values for sharing decks.
//
// swagger:model shareGroupDeckRequest
type ShareGroupDeckRequest struct {
	// UserId is an ID of the user in a storage.
	//
	// required: true
	// example: 189
	UserId UserId `json:"user_id"`

	// GroupId is an ID of a group in the storage.
	//
	// required: true
	// example: 222
	GroupId GroupId `json:"group_id"`

	// DeckId is an ID of a deck in the storage.
	//
	// required: true
	// example: 3
	DeckId DeckId `json:"deck_id"`
}

type ShareGroupDeckResponse struct{}

type DeleteGroupDeckRequest struct {
	UserId  UserId  `json:"user_id"`
	GroupId GroupId `json:"group_id"`
	DeckId  DeckId  `json:"deck_id"`
}

type DeleteGroupDeckResponse struct{}

type SearchGroupByNameRequest struct{} // GET

type SearchGroupByNameResponse struct {
	Groups []GroupConfig `json:"groups"`
}

type ChangeGroupNameRequest struct {
	CreatorId UserId  `json:"creator_id"`
	GroupId   GroupId `json:"group_id"`
	Name      string  `json:"name"`
}

type ChangeGroupNameResponse struct{}

type QuitGroupRequest struct {
	UserId  UserId  `json:"user_id"`
	GroupId GroupId `json:"group_id"`
}

type QuitGroupResponse struct{}

// SearchUserResponse is a struct that defines the response body for
// searching users by name.
type SearchUserResponse struct {
	Users []UserInfo `json:"users"`
}

// UserInfo is a struct that defines the user info model.
type UserInfo struct {
	ID       int
	Username string
	Email    string
}

func (u *UserInfo) Scan(v any) error {
	switch vt := v.(type) {
	case []byte:
		err := json.Unmarshal(vt, u)
		if err != nil {
			return err
		}
	case string:
		err := json.NewDecoder(strings.NewReader(vt)).Decode(u)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("type not allowed")
	}

	return nil
}

// FeedRequest is a struct that defines the request body for getting a user's feed.
// swagger:model feedRequest
type FeedRequest struct {
	UserId UserId `json:"user_id"`
}

type FeedResponse struct {
	Feed []Post `json:"feed"`
}

type ShareWithFollowersRequest struct {
	UserId UserId `json:"user_id"`
	DeckId DeckId `json:"deck_id"`
}

type CheckIfSharedWithFollowersRequest struct {
	UserId UserId `json:"user_id"`
	DeckId DeckId `json:"deck_id"`
}

type CheckIfSharedWithFollowersResponse struct {
	Ok bool `json:"ok"`
}

type GetParticipantsByGroupIdRequest struct {
	GroupId GroupId `json:"group_id"`
}

type GetParticipantsByGroupIdResponse struct {
	Participants []UserInfo `json:"participants"`
}
