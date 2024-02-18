package models

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
	Groups Groups `json:"decks"`
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

// GetGroupsByUserIdRequest is a struct that defines the request body for
// loading groups the user is a member of.
type GetGroupsByUserIdResponse struct {
	// Groups is a map from group id to group config.
	//
	// required: true
	Groups Groups `json:"decks"`
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

// GetUsersByGroupIdRequest is a struct that defines the request body for
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

// GetGroupByGroupIdRequest is a struct that defines the request body for
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

// GetDecksByGroupIdRequest is a struct that defines the request body for
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

// GetInvitesByGroupIdRequest is a struct that defines the request body for
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

// GetInvitesByUserIdRequest is a struct that defines the request body for
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
	// example: fefhjkehsfuhgwefvgewhufgvhejwgvfhewghfjewfhjewsvfhgewhf
	Name string `json:"name"`
}

// CreateGroupRequest is a struct that defines the request body for
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

// DeleteGroupRequest is a struct that defines the request body for
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

// AcceptInviteRequest is a struct that defines the request body for
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

// SendInviteRequest is a struct that defines the request body for
// sending an invite to group.
type SendInviteResponse struct{}

// ShareAllGroupDecksRequest is a struct that defines the request body for
// sharing all the group decks to a new user.
//
//swagger:model shareAllGroupDecksRequest
type ShareAllGroupDecksRequest struct {
	// CreatorId is an ID of the user in a storage.
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

// ShareAllGroupDecksRequest is a struct that defines the request body for
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

// DeepCopyDeckRequest is a struct that defines the request body for
// deep copying a deck (with all its contents recursively copied).
type DeepCopyDeckResponse struct {
	// DeckId is an ID of the newly copied deck.
	//
	// required: true
	// example: 189
	DeckId DeckId `json:"deck_id"`
}
