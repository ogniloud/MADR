package storage

import (
	"context"

	cardmodels "github.com/ogniloud/madr/internal/flashcards/models"
	usermodels "github.com/ogniloud/madr/internal/models"
	"github.com/ogniloud/madr/internal/social/models"
)

// Storage is an interface for accessing a database.
type Storage interface {
	// GetCreatedGroupsByUserId returns all the groups the user created
	GetCreatedGroupsByUserId(ctx context.Context, id models.UserId) (models.Groups, error)

	// GetGroupsByUserId returns all the groups the user pertains to
	GetGroupsByUserId(ctx context.Context, id models.UserId) (models.Groups, error)

	// GetUsersByGroupId returns all users pertaining to the group
	GetUsersByGroupId(ctx context.Context, id models.GroupId) (models.Members, error)

	// GetGroupByGroupId returns the group with the given id
	GetGroupByGroupId(ctx context.Context, id models.GroupId) (models.GroupConfig, error)

	// TODO
	GetDecksByGroupId(ctx context.Context, id models.GroupId) ([]cardmodels.DeckConfig, error)

	// GetInvitesByGroupId returns invites from the given group
	GetInvitesByGroupId(ctx context.Context, id models.GroupId) (map[models.UserId]models.InviteInfo, error)

	// GetInvitesByUserId returns invites to the given user
	GetInvitesByUserId(ctx context.Context, id models.UserId) (map[models.GroupId]models.InviteInfo, error)

	// Creates new group with id being the owner
	CreateGroup(ctx context.Context, id models.UserId, name string) (models.GroupId, error)

	// Deletes the group where id belongs to its creator
	DeleteGroup(ctx context.Context, id models.UserId, group_id models.GroupId) error

	// User accepts an invite to a group and becomes a member of the group
	AcceptInvite(ctx context.Context, id models.UserId, group_id models.GroupId) error

	// Sends invite to group to the user
	SendInvite(ctx context.Context, creator_id models.UserId, invitee models.UserId, group_id models.GroupId) error

	// TODO
	ShareAllGroupDecks(ctx context.Context, id models.UserId, group_id models.GroupId) error

	DeepCopyDeck(ctx context.Context, copier models.UserId, deckId models.DeckId) (models.DeckId, error)

	GetFollowersByUserId(ctx context.Context, id models.UserId) ([]usermodels.UserInfo, error)

	GetFollowingsByUserId(ctx context.Context, id models.UserId) ([]usermodels.UserInfo, error)

	Follow(ctx context.Context, follower models.UserId, author models.UserId) error

	Unfollow(ctx context.Context, follower models.UserId, author models.UserId) error

	ShareDeckGroup(ctx context.Context, owner models.UserId, groupId models.GroupId, deckId models.DeckId) error

	DeleteDeckFromGroup(ctx context.Context, owner models.UserId, groupId models.GroupId, deckId models.DeckId) error

	GetGroupsByName(ctx context.Context, name string) ([]models.GroupConfig, error)
}
