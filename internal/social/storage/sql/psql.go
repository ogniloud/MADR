package sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/ogniloud/madr/internal/db"
	"github.com/ogniloud/madr/internal/social/models"
)

type Storage struct {
	Conn *db.PSQLDatabase
}

func (d *Storage) GetCreatedGroupsByUserId(ctx context.Context, id models.UserId) (models.Groups, error) {
	rows, err := d.Conn.Query(ctx, `SELECT * FROM groups WHERE creator_id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	groups := models.Groups{}

	cfg := models.GroupConfig{}
	_, err = pgx.ForEachRow(rows, []any{&cfg.GroupId, &cfg.CreatorId, &cfg.Name, &cfg.TimeCreated}, func() error {
		groups[cfg.GroupId] = cfg

		return nil
	})

	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (d *Storage) GetGroupsByUserId(ctx context.Context, id models.UserId) (models.Groups, error) {
	rows, err := d.Conn.Query(ctx, `
	SELECT *
	FROM groups
	WHERE group_id in (
		SELECT group_id
		FROM group_members
		WHERE user_id=$1
	)`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	groups := models.Groups{}

	cfg := models.GroupConfig{}
	_, err = pgx.ForEachRow(rows, []any{&cfg.GroupId, &cfg.CreatorId, &cfg.Name, &cfg.TimeCreated}, func() error {
		groups[cfg.GroupId] = cfg

		return nil
	})

	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (d *Storage) GetUsersByGroupId(ctx context.Context, id models.GroupId) (models.Members, error) {
	rows, err := d.Conn.Query(ctx, `
	SELECT *
	FROM group_members
	WHERE group_id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	members := models.Members{}

	cfg := models.MemberInfo{}
	_, err = pgx.ForEachRow(rows, []any{&cfg.MemberId, &cfg.TimeJoined}, func() error {
		members[cfg.MemberId] = cfg

		return nil
	})

	if err != nil {
		return nil, err
	}

	return members, nil
}

func (d *Storage) GetGroupByGroupId(ctx context.Context, id models.GroupId) (models.GroupConfig, error) {
	row := d.Conn.QueryRow(ctx, `
	SELECT *
	FROM groups
	WHERE group_id=$1`, id)

	group := models.GroupConfig{}
	err := row.Scan(&group.GroupId, &group.CreatorId, &group.Name, &group.TimeCreated)

	if err != nil {
		return models.GroupConfig{}, fmt.Errorf("psql error: %w", err)
	}

	return group, nil
}

func (d *Storage) GetDecksByGroupId(ctx context.Context, id models.GroupId) (models.Decks, error) {
	// TODO
	return nil, nil
}

func (d *Storage) GetInvitesByGroupId(ctx context.Context, id models.GroupId) ([]models.InviteInfo, error) {
	rows, err := d.Conn.Query(ctx, `
	SELECT *
	FROM group_invites
	WHERE group_id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	invites := make([]models.InviteInfo, 0)

	cfg := models.InviteInfo{}
	_, err = pgx.ForEachRow(rows, []any{&cfg.GroupId, &cfg.InvitedId, &cfg.TimeInvited}, func() error {
		invites = append(invites, cfg)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return invites[:len(invites):len(invites)], nil
}
