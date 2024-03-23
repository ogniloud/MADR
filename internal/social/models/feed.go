package models

type FeedPostType string

const (
	Invite              = "invite_data"
	SharedFromGroup     = "shared_from_group_data"
	SharedFromFollowing = "shared_from_following_data"
	FollowingSubscribed = "following_subscribed_data"
)

type Post struct {
	Type                    FeedPostType             `json:"type"`
	InviteData              *InviteData              `json:"invite_data"`
	SharedFromGroupData     *SharedFromGroupData     `json:"shared_from_group_data"`
	SharedFromFollowingData *SharedFromFollowingData `json:"shared_from_following_data"`
	FollowingSubscribedData *FollowingSubscribedData `json:"following_subscribed_data"`
}

type InviteData struct {
	InviteeId   UserId  `json:"invitee_id"`
	InviteeName string  `json:"invitee_name"`
	GroupId     GroupId `json:"group_id"`
	GroupName   string  `json:"group_name"`
}

type SharedFromGroupData struct {
	GroupId   GroupId `json:"group_id"`
	GroupName string  `json:"group_name"`
	DeckId    DeckId  `json:"deck_id"`
	DeckName  string  `json:"deck_name"`
}

type SharedFromFollowingData struct {
	AuthorId   UserId `json:"author_id"`
	AuthorName string `json:"author_name"`
	DeckId     DeckId `json:"deck_id"`
	DeckName   string `json:"deck_name"`
}

type FollowingSubscribedData struct {
	FollowerId   UserId `json:"follower_id"`
	FollowerName string `json:"follower_name"`
	AuthorId     UserId `json:"author_id"`
	AuthorName   string `json:"author_name"`
}
