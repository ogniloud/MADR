package models

import (
	"time"

	"github.com/ogniloud/madr/internal/flashcards/models"
)

type (
	GroupId = int
	DeckId  = int
	UserId  = models.UserId
	Decks   = models.Decks
)

// GroupConfig contains information about a particular group.
type GroupConfig struct {
	// GroupId is a unique identifier of the group.
	//
	// required: true
	// example: 1
	GroupId GroupId `json:"group_id"`

	// CreatorId means that a user created the group and is authorized to share posts within it.
	//
	// required: true
	// example: 12
	CreatorId UserId `json:"creator_id"`

	// Name is a name of the group which the user assigned to it.
	//
	// required: true
	// example: vodka lovers
	Name string `json:"name"`

	// Timestamp when the group was created
	//
	// required: true
	// example: 2020-01-01T00:00:00Z
	TimeCreated time.Time `json:"time_created"`
}

// MemberInfo contains information about a particular group member.
type MemberInfo struct {
	MemberId UserId `json:"member_id"`

	// Timestamp when the member joined the group
	TimeJoined time.Time `json:"time_joined"`
}

// InviteInfo contains information about a particular user that got invite into the group.
type InviteInfo struct {
	GroupId GroupId `json:"group_id"`

	// Id of the user that received the invitation
	InvitedId UserId `json:"user_id"`

	// Timestamp when the invite was sent
	TimeInvited time.Time `json:"time_invited"`
}

// Groups is a map of groups which config can be obtained by id.
type Groups map[GroupId]GroupConfig

// Members is a map of members of a group.
type Members map[UserId]MemberInfo

type GroupsShared struct {
	DeckId    DeckId  `json:"deck_id"`
	GroupId   GroupId `json:"group_id"`
	GroupName string  `json:"group_name"`
	Shared    bool    `json:"shared"`
}

type GroupsFollowed struct {
	GroupId      GroupId `json:"group_id"`
	FollowerId   UserId  `json:"follower_id"`
	FollowerName string  `json:"follower_name"`
}
