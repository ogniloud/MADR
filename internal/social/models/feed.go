package models

type FeedPostType string

const (
	Invite              = "invite_data"
	SharedFromGroup     = "shared_from_group_data"
	SharedFromFollowing = "shared_from_following_data"
	FollowingSubscribed = "following_subscribed_data"
)

type Post struct {
	// PostId is a unique identifier of the post.
	//
	// example: 1
	Type FeedPostType `json:"type"`

	// InviteData contains information about the invitation.
	InviteData *InviteData `json:"invite_data"`

	// SharedFromGroupData contains information about the shared deck from the group.
	SharedFromGroupData *SharedFromGroupData `json:"shared_from_group_data"`

	// SharedFromFollowingData contains information about the shared deck from the following.
	SharedFromFollowingData *SharedFromFollowingData `json:"shared_from_following_data"`

	// FollowingSubscribedData contains information about the following subscription.
	FollowingSubscribedData *FollowingSubscribedData `json:"following_subscribed_data"`
}

type InviteData struct {
	// InviteeId is a unique identifier of the user that received the invitation.
	//
	// example: 1
	InviteeId UserId `json:"invitee_id"`

	// InviteeName is a name of the user that received the invitation.
	//
	// example: John Doe
	InviteeName string `json:"invitee_name"`

	// GroupId is a unique identifier of the group.
	//
	// example: 1
	GroupId GroupId `json:"group_id"`

	// GroupName is a name of the group.
	//
	// example: vodka lovers
	GroupName string `json:"group_name"`
}

type SharedFromGroupData struct {
	// GroupId is a unique identifier of the group.
	//
	// example: 1
	GroupId GroupId `json:"group_id"`

	// GroupName is a name of the group.
	//
	// example: vodka lovers
	GroupName string `json:"group_name"`

	// DeckId is a unique identifier of the deck.
	//
	// example: 1
	DeckId DeckId `json:"deck_id"`

	// DeckName is a name of the deck.
	//
	// example: Bangladesh famous words
	DeckName string `json:"deck_name"`
}

type SharedFromFollowingData struct {
	// AuthorId is a unique identifier of the author.
	//
	// example: 1
	AuthorId UserId `json:"author_id"`

	// AuthorName is a name of the author.
	//
	// example: John Doe
	AuthorName string `json:"author_name"`

	// DeckId is a unique identifier of the deck.
	//
	// example: 1
	DeckId DeckId `json:"deck_id"`

	// DeckName is a name of the deck.
	//
	// example: russian curse words
	DeckName string `json:"deck_name"`
}

type FollowingSubscribedData struct {
	// FollowerId is a unique identifier of the follower.
	//
	// example: 1
	FollowerId UserId `json:"follower_id"`

	// FollowerName is a name of the follower.
	//
	// example: John Doe
	FollowerName string `json:"follower_name"`

	// AuthorId is a unique identifier of the author.
	//
	// example: 1
	AuthorId UserId `json:"author_id"`

	// AuthorName is a name of the author.
	//
	// example: Dima Krasnov
	AuthorName string `json:"author_name"`
}
