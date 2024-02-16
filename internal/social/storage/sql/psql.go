package sql

import (
	"context"
	"errors"
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

func (d *Storage) GetDecksByGroupId(ctx context.Context, id models.GroupId) ([]models.DeckId, error) {
	rows, err := d.Conn.Query(ctx, `SELECT deck_id FROM group_decks WHERE group_id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("group decks not loaded: %w", err)
	}
	defer rows.Close()

	vals, err := rows.Values()
	if err != nil {
		return nil, fmt.Errorf("group deck rows read fail: %w", err)
	}

	ids := make([]models.DeckId, len(vals))
	for k, v := range vals {
		ids[k], _ = v.(models.DeckId)
	}

	return ids, nil
}

func (d *Storage) GetInvitesByGroupId(ctx context.Context, id models.GroupId) (map[models.UserId]models.InviteInfo, error) {
	rows, err := d.Conn.Query(ctx, `
	SELECT *
	FROM group_invites
	WHERE group_id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	invites := make(map[models.UserId]models.InviteInfo)

	cfg := models.InviteInfo{}
	_, err = pgx.ForEachRow(rows, []any{&cfg.GroupId, &cfg.InvitedId, &cfg.TimeInvited}, func() error {
		invites[cfg.InvitedId] = cfg

		return nil
	})

	if err != nil {
		return nil, err
	}

	return invites, nil
}

func (d *Storage) GetInvitesByUserId(ctx context.Context, id models.UserId) (map[models.GroupId]models.InviteInfo, error) {
	rows, err := d.Conn.Query(ctx, `
	SELECT *
	FROM group_invites
	WHERE user_id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	invites := make(map[models.GroupId]models.InviteInfo)

	cfg := models.InviteInfo{}
	_, err = pgx.ForEachRow(rows, []any{&cfg.GroupId, &cfg.InvitedId, &cfg.TimeInvited}, func() error {
		invites[cfg.GroupId] = cfg

		return nil
	})

	if err != nil {
		return nil, err
	}

	return invites, nil
}

func (d *Storage) CreateGroup(ctx context.Context, id models.UserId, name string) (models.GroupId, error) {
	row := d.Conn.QueryRow(ctx,
		`INSERT INTO groups (creator_id, name, time_created) VALUES ($1, $2, now()) RETURNING group_id`,
		id, name,
	)

	groupId := 0

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("psql error: %w", err)
	}

	err = d.addMember(ctx, id, groupId)
	if err != nil {
		return 0, err
	}

	return groupId, nil
}

func (d *Storage) DeleteGroup(ctx context.Context, id models.UserId, groupId models.GroupId) error {
	row := d.Conn.QueryRow(ctx,
		`DELETE FROM groups WHERE creator_id=$1 AND group_id=$2 RETURNING name`, id, groupId)

	s := ""
	return row.Scan(&s)
}

func (d *Storage) AcceptInvite(ctx context.Context, id models.UserId, groupId models.GroupId) error {
	invites, err := d.GetInvitesByUserId(ctx, id)
	if err != nil {
		return fmt.Errorf("no invites sent to user %v, error: %w", id, err)
	}

	_, present := invites[groupId]
	if !present {
		return fmt.Errorf("no invites from the group %v", groupId)
	}

	row := d.Conn.QueryRow(ctx,
		`DELETE FROM group_invites WHERE user_id=$1 AND group_id=$2 RETURNING user_id`, id, groupId)

	err = row.Scan(&id)
	if err != nil {
		return err
	}

	return d.addMember(ctx, id, groupId)
}

func (d *Storage) SendInvite(ctx context.Context, creatorId models.UserId, invitee models.UserId, groupId models.GroupId) error {
	group, err := d.GetGroupByGroupId(ctx, groupId)
	if err != nil {
		return fmt.Errorf("couldn't get the group the invite is supposed to be sent from, error: %w", err)
	}
	if group.CreatorId != creatorId {
		return fmt.Errorf("the user doesn't have rights to send invites to the group")
	}

	row := d.Conn.QueryRow(ctx,
		`INSERT INTO group_invites VALUES ($1, $2, now()) RETURNING group_id`, groupId, invitee)
	return row.Scan(&groupId)
}

func (d *Storage) ShareAllGroupDecks(ctx context.Context, id models.UserId, groupId models.GroupId) error {
	return nil
}

func (d *Storage) CheckIfCopied(ctx context.Context, copier models.UserId, deckId models.DeckId) (bool, error) {
	row := d.Conn.QueryRow(ctx, `SELECT deck_id FROM copied_by WHERE copier_id=$1 AND deck_id=$2`, copier, deckId)
	id := 0

	if err := row.Scan(&id); err == nil {
		return true, nil
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return false, fmt.Errorf("fatal db error: %w", err)
	}

	return false, nil
}

func (d *Storage) DeepCopyDeck(ctx context.Context, copier models.UserId, deckId models.DeckId) (models.DeckId, error) {
	if ok, err := d.CheckIfCopied(ctx, copier, deckId); ok {
		return 0, errors.New("deck already copied")
	} else if err != nil {
		return 0, err
	}

	tx, err := d.Conn.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("copy transaction fail: %w", err)
	}

	// scan name in order to insert a new deck record
	name := "default_copied_deck_name"
	row := tx.QueryRow(ctx, `SELECT name FROM deck_config WHERE deck_id=$1`, deckId)
	err = row.Scan(&name)
	if err != nil {
		d.Conn.Logger().Errorf("Name wasn't selected: %v", err)
	}

	// insert a new deck and return an id
	id := 0
	row = tx.QueryRow(ctx, `INSERT INTO deck_config(user_id, name) VALUES ($1, $2) RETURNING deck_id`, copier, name)
	err = row.Scan(&id)
	if err != nil {
		defer func() {
			if err := tx.Rollback(ctx); err != nil {
				d.Conn.Logger().Errorf("Transaction rollback failed: %v", err)
			}
		}()
		return 0, err
	}

	// copy all the flashcards
	_, err = tx.Exec(ctx, `INSERT INTO flashcard(word, backside, deck_id, answer)
SELECT f.word, f.backside, $1, f.answer FROM flashcard AS f
WHERE f.deck_id = $2`, id, deckId)
	if err != nil {
		return 0, fmt.Errorf("couldn't copy flashcards: %w", err)
	}

	// create a new record about copying
	_, err = tx.Exec(ctx, `INSERT INTO copied_by VALUES ($1, $2, now())`, copier, deckId)
	if err != nil {
		defer func() {
			if err := tx.Rollback(ctx); err != nil {
				d.Conn.Logger().Errorf("Transaction rollback failed: %v", err)
			}
		}()
		return 0, err
	}

	if err := tx.Commit(ctx); err != nil {
		d.Conn.Logger().Errorf("deep copy not commited: %v", err)
		return 0, fmt.Errorf("deep copy not commited: %w", err)
	}

	return id, nil
}

func (d *Storage) addMember(ctx context.Context, id models.UserId, groupId models.GroupId) error {
	row := d.Conn.QueryRow(ctx,
		`INSERT INTO group_members VALUES ($1, $2, now()) RETURNING group_id`, groupId, id)

	err := row.Scan(&groupId)
	if err != nil {
		return err
	}

	return d.ShareAllGroupDecks(ctx, id, groupId)
}
