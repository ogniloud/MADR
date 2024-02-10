package storage

import (
	"context"

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

	// GetDecksByGroupId returns the decks shared within the group
	GetDecksByGroupId(ctx context.Context, id models.GroupId) (models.Decks, error)

	// GetInvitesByGroupId returns the users invited into the group
	GetInvitesByGroupId(ctx context.Context, id models.GroupId) ([]models.InviteInfo, error)
}
