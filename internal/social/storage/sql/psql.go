package sql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/ogniloud/madr/internal/db"
	cardmodels "github.com/ogniloud/madr/internal/flashcards/models"
	usermodels "github.com/ogniloud/madr/internal/models"
	"github.com/ogniloud/madr/internal/social/models"
)

var (
	ErrUserNotCreator = errors.New("the user is not a group creator")
	ErrAlreadyCopied  = errors.New("deck already copied")
)

type Storage struct {
	Conn *db.PSQLDatabase
}

func (d *Storage) GetCreatedGroupsByUserId(ctx context.Context, id models.UserId) ([]models.GroupConfig, error) {
	rows, err := d.Conn.Query(ctx, `SELECT group_id, creator_id, name, time_created FROM groups WHERE creator_id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var groups []models.GroupConfig

	cfg := models.GroupConfig{}
	_, err = pgx.ForEachRow(rows, []any{&cfg.GroupId, &cfg.CreatorId, &cfg.Name, &cfg.TimeCreated}, func() error {
		groups = append(groups, cfg)

		return nil
	})

	if err != nil {
		return nil, err
	}

	groups1 := make([]models.GroupConfig, len(groups))
	copy(groups1, groups)
	return groups1, nil
}

func (d *Storage) GetGroupsByUserId(ctx context.Context, id models.UserId) ([]models.GroupConfig, error) {
	rows, err := d.Conn.Query(ctx, `
	SELECT group_id, creator_id, name, time_created
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

	groups := []models.GroupConfig{}

	cfg := models.GroupConfig{}
	_, err = pgx.ForEachRow(rows, []any{&cfg.GroupId, &cfg.CreatorId, &cfg.Name, &cfg.TimeCreated}, func() error {
		groups = append(groups, cfg)

		return nil
	})

	if err != nil {
		return nil, err
	}

	groups1 := make([]models.GroupConfig, len(groups))
	copy(groups1, groups)
	return groups1, nil
}

func (d *Storage) GetUsersByGroupId(ctx context.Context, id models.GroupId) (models.Members, error) {
	rows, err := d.Conn.Query(ctx, `
	SELECT user_id, time_joined
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
	SELECT group_id, creator_id, name, time_created
	FROM groups
	WHERE group_id=$1`, id)

	group := models.GroupConfig{}
	err := row.Scan(&group.GroupId, &group.CreatorId, &group.Name, &group.TimeCreated)

	if err != nil {
		return models.GroupConfig{}, fmt.Errorf("psql error: %w", err)
	}

	return group, nil
}

func (d *Storage) GetDecksByGroupId(ctx context.Context, id models.GroupId) ([]cardmodels.DeckConfig, error) {
	rows, err := d.Conn.Query(ctx, `SELECT dc.deck_id, user_id, name FROM group_decks
                 JOIN deck_config dc on dc.deck_id = group_decks.deck_id
                 WHERE group_id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("group decks not loaded: %w", err)
	}
	defer rows.Close()

	dcs := make([]cardmodels.DeckConfig, 0)
	var dc cardmodels.DeckConfig
	_, err = pgx.ForEachRow(rows, []any{&dc.DeckId, &dc.UserId, &dc.Name}, func() error {
		dcs = append(dcs, dc)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("group deck rows read fail: %w", err)
	}

	dcs1 := make([]cardmodels.DeckConfig, len(dcs))
	copy(dcs1, dcs)

	return dcs1, nil
}

func (d *Storage) GetInvitesByGroupId(ctx context.Context, id models.GroupId) (map[models.UserId]models.InviteInfo, error) {
	rows, err := d.Conn.Query(ctx, `
	SELECT group_id, invite_id, time_sent
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
	SELECT group_id, invite_id, time_sent
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
		d.Conn.Logger().Errorf("create group failed: %v", err)
		return 0, fmt.Errorf("psql error: %w", err)
	}

	err = d.addMember(ctx, id, groupId)
	if err != nil {
		d.Conn.Logger().Errorf("add member failed: %v", err)
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
		`INSERT INTO group_invites (group_id, user_id, time_sent) VALUES ($1, $2, now()) RETURNING group_id`, groupId, invitee)
	if err := row.Scan(&groupId); err != nil {
		return fmt.Errorf("send invite error: %w", err)
	}

	go func() {
		data := &models.Post{
			Type: models.Invite,
			Data: models.InviteData{
				InviteeId:   invitee,
				InviteeName: "TEST_NAME_INVITEE",
				GroupId:     groupId,
				GroupName:   "TEST_NAME_GROUP",
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := d.saveToFeed(ctx, invitee, data); err != nil {
			d.Conn.Logger().Errorf("invite: %v", err)
		}
	}()

	return nil
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
		return 0, ErrAlreadyCopied
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

func (d *Storage) GetFollowersByUserId(ctx context.Context, id models.UserId) ([]usermodels.UserInfo, error) {
	rows, err := d.Conn.Query(ctx, `SELECT follower_id, user_credentials.username, user_credentials.email FROM followers 
    JOIN user_credentials ON followers.follower_id = user_credentials.user_id 
                   WHERE followers.user_id=$1`, id)
	if err != nil {
		d.Conn.Logger().Errorf("couldn't get followers: %v", err)
		return nil, fmt.Errorf("couldn't get followers: %w", err)
	}
	defer func() { _ = rows.Close }()

	var info usermodels.UserInfo
	var infos []usermodels.UserInfo
	_, err = pgx.ForEachRow(rows, []any{&info.ID, &info.Username, &info.Email}, func() error {
		infos = append(infos, info)
		return nil
	})
	if err != nil {
		d.Conn.Logger().Errorf("foreach followers error: %v", err)
		return nil, fmt.Errorf("foreach followers error: %w", err)
	}

	infos2 := make([]usermodels.UserInfo, len(infos))
	copy(infos2, infos)

	return infos2, nil
}

func (d *Storage) GetFollowingsByUserId(ctx context.Context, id models.UserId) ([]usermodels.UserInfo, error) {
	rows, err := d.Conn.Query(ctx, `SELECT followers.user_id, user_credentials.username, user_credentials.email FROM followers 
    JOIN user_credentials ON followers.user_id = user_credentials.user_id 
                   WHERE followers.follower_id=$1`, id)
	if err != nil {
		d.Conn.Logger().Errorf("couldn't get followers: %v", err)
		return nil, fmt.Errorf("couldn't get followers: %w", err)
	}
	defer func() { _ = rows.Close }()

	var info usermodels.UserInfo
	var infos []usermodels.UserInfo
	_, err = pgx.ForEachRow(rows, []any{&info.ID, &info.Username, &info.Email}, func() error {
		infos = append(infos, info)
		return nil
	})
	if err != nil {
		d.Conn.Logger().Errorf("foreach followers error: %v", err)
		return nil, fmt.Errorf("foreach followers error: %w", err)
	}

	infos2 := make([]usermodels.UserInfo, len(infos))
	copy(infos2, infos)

	return infos2, nil
}

func (d *Storage) addMember(ctx context.Context, id models.UserId, groupId models.GroupId) error {
	row := d.Conn.QueryRow(ctx,
		`INSERT INTO group_members (group_id, user_id, time_joined) VALUES ($1, $2, now()) RETURNING group_id`, groupId, id)

	err := row.Scan(&groupId)
	if err != nil {
		return err
	}

	return d.ShareAllGroupDecks(ctx, id, groupId)
}

func (d *Storage) Follow(ctx context.Context, follower models.UserId, author models.UserId) error {
	_, err := d.Conn.Exec(ctx, `INSERT INTO followers (user_id, follower_id, time_followed) VALUES
                                                                ($1, $2, now())`, author, follower)
	if err != nil {
		d.Conn.Logger().Errorf("follow error: %v", err)
		return fmt.Errorf("follow error: %w", err)
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rows, err := d.Conn.Query(ctx, `SELECT follower_id FROM followers WHERE user_id=$1`, follower)
		if err != nil {
			d.Conn.Logger().Errorf("feed: select followers err: %v", err)
			return
		}
		defer rows.Close()

		followers, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.UserId, error) {
			var u models.UserId
			err := row.Scan(&u)
			return u, err
		})
		if err != nil {
			d.Conn.Logger().Errorf("collect followers err: %v", err)
			return
		}

		names, err := d.Conn.Query(ctx, `SELECT user_id, username FROM user_credentials WHERE 
                                      user_id=$1 OR user_id=$2`, follower, author)
		if err != nil {
			d.Conn.Logger().Errorf("feed: select followers err: %v", err)
			return
		}
		defer names.Close()

		m := make(map[models.UserId]string, 2)
		var u models.UserId
		var un string
		_, err = pgx.ForEachRow(names, []any{&u, &un}, func() error {
			m[u] = un
			return nil
		})
		if err != nil {
			d.Conn.Logger().Errorf("names collect failed err: %v", err)
			return
		}

		data := models.FollowingSubscribedData{
			FollowerId:   follower,
			FollowerName: m[follower],
			AuthorId:     author,
			AuthorName:   m[author],
		}
		b, _ := json.Marshal(&data)
		s := string(b)
		t := time.Now()
		_, err = d.Conn.CopyFrom(ctx,
			pgx.Identifier{"feed"},
			[]string{"user_id", "data", "timestamp"},
			pgx.CopyFromSlice(len(followers), func(i int) ([]any, error) {
				return []any{followers[i], s, t}, nil
			}),
		)
		if err != nil {
			d.Conn.Logger().Errorf("copy to feed failed: %v", err)
			return
		}
	}()

	return nil
}

func (d *Storage) Unfollow(ctx context.Context, follower models.UserId, author models.UserId) error {
	_, err := d.Conn.Exec(ctx, `DELETE FROM followers WHERE user_id=$1 AND follower_id=$2`, author, follower)
	if err != nil {
		d.Conn.Logger().Errorf("follow error: %v", err)
		return fmt.Errorf("follow error: %w", err)
	}

	return nil
}

func (d *Storage) ShareDeckGroup(ctx context.Context, owner models.UserId, groupId models.GroupId, deckId models.DeckId) error {
	_, err := d.Conn.Query(ctx, `SELECT group_id FROM groups WHERE group_id=$1 AND creator_id=$2`, groupId, owner)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrUserNotCreator
	} else if err != nil {
		d.Conn.Logger().Errorf("share deck error: %v", err)
		return fmt.Errorf("share deck error: %w", err)
	}

	_, err = d.Conn.Exec(ctx, `INSERT INTO group_decks VALUES ($1, $2, now())`, groupId, deckId)
	if err != nil {
		d.Conn.Logger().Errorf("group deck add error: %v", err)
		return fmt.Errorf("group deck add error: %w", err)
	}

	_, err = d.Conn.Exec(ctx, `INSERT INTO user_leitner(user_id, card_id, box, cool_down)
    (SELECT u, c, 0 as box, now() as cool_down FROM
            (SELECT user_id as u FROM group_members WHERE group_id=$1 AND user_id != $2) as gmu
            LEFT JOIN (SELECT flashcard.card_id as c FROM flashcard WHERE deck_id=$3) as fc ON TRUE
            WHERE NOT u IN (SELECT copier_id FROM copied_by WHERE deck_id=$4)
    );`, groupId, owner, deckId, deckId)
	if err != nil {
		d.Conn.Logger().Errorf("add user leitners to members failed: %v", err)
		return fmt.Errorf("add user leitners to members failed: %w", err)
	}

	return nil
}

func (d *Storage) DeleteDeckFromGroup(ctx context.Context, owner models.UserId, groupId models.GroupId, deckId models.DeckId) error {
	_, err := d.Conn.Query(ctx, `SELECT group_id FROM groups WHERE group_id=$1 AND creator_id=$2`, groupId, owner)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrUserNotCreator
	} else if err != nil {
		d.Conn.Logger().Errorf("share deck error: %v", err)
		return fmt.Errorf("share deck error: %w", err)
	}

	_, err = d.Conn.Exec(ctx, `DELETE FROM user_leitner WHERE
    (user_leitner.user_id, user_leitner.card_id) IN (
        SELECT * FROM (SELECT user_id FROM group_members WHERE group_id=$1 AND user_leitner.user_id!=$2)
                          JOIN (SELECT f.card_id FROM flashcard f WHERE deck_id=$3) ON TRUE)`, groupId, owner, deckId)
	if err != nil {
		d.Conn.Logger().Errorf("delete user leitners for group members failed: %v", err)
		return fmt.Errorf("delete user leitners for group members failed: %w", err)
	}

	_, err = d.Conn.Exec(ctx, `DELETE FROM group_decks WHERE group_id=$1 AND deck_id=$2`, groupId, deckId)
	if err != nil {
		d.Conn.Logger().Errorf("group deck delete error: %v", err)
		return fmt.Errorf("group deck add error: %w", err)
	}

	return nil
}

func (d *Storage) GetGroupsByName(ctx context.Context, name string) ([]models.GroupConfig, error) {
	rows, err := d.Conn.Query(ctx, `SELECT * FROM groups WHERE name=$1
										UNION DISTINCT SELECT * FROM groups WHERE name LIKE $1`, name+"%")
	if err != nil {
		d.Conn.Logger().Errorf("groups not got: %v", err)
		return nil, fmt.Errorf("groups not got: %w", err)
	}

	var gcs []models.GroupConfig
	var t models.GroupConfig
	_, err = pgx.ForEachRow(rows, []any{&t.GroupId, &t.CreatorId, &t.Name, &t.TimeCreated}, func() error {
		gcs = append(gcs, t)
		return nil
	})
	if err != nil {
		d.Conn.Logger().Errorf("iteration error: %v", err)
		return nil, fmt.Errorf("iteration error: %w", err)
	}

	gcs1 := make([]models.GroupConfig, len(gcs))
	copy(gcs1, gcs)
	return gcs1, nil
}

func (d *Storage) ChangeGroupName(ctx context.Context, creatorId models.UserId, groupId models.GroupId, name string) error {
	_, err := d.Conn.Exec(ctx, `UPDATE groups SET name=$1 WHERE group_id=$2 AND creator_id=$3`,
		name, creatorId, groupId)
	if err != nil {
		d.Conn.Logger().Errorf("name not updated %v", err)
		return fmt.Errorf("name not updated %w", err)
	}

	return nil
}

func (d *Storage) QuitGroup(ctx context.Context, userId models.UserId, groupId models.GroupId) error {
	_, err := d.Conn.Exec(ctx, `DELETE FROM user_leitner WHERE user_id=$1 AND
                               (card_id IN (SELECT f.card_id FROM group_decks
                                            JOIN public.flashcard f on group_decks.deck_id = f.deck_id
                                            WHERE group_id=$2))`, userId, groupId)
	if err != nil {
		d.Conn.Logger().Errorf("decks not deleted %v", err)
		return fmt.Errorf("decks not deleted %w", err)
	}

	_, err = d.Conn.Exec(ctx, `DELETE FROM group_members WHERE group_id=$1 AND user_id=$2`,
		groupId, userId)
	if err != nil {
		d.Conn.Logger().Errorf("member not deleted %v", err)
		return fmt.Errorf("member not deleted %w", err)
	}

	return nil
}

func (d *Storage) GetUsersByName(ctx context.Context, name string) ([]usermodels.UserInfo, error) {
	rows, err := d.Conn.Query(ctx, `SELECT user_id, username, email FROM user_credentials WHERE username LIKE $1`, "%"+name+"%")
	if err != nil {
		d.Conn.Logger().Errorf("Storage.GetUsersByName: %v", err)
		return nil, fmt.Errorf("can't get users by name: %w", err)
	}

	var u []usermodels.UserInfo

	for rows.Next() {
		var user usermodels.UserInfo
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			d.Conn.Logger().Errorf("Storage.GetUsersByName: %v", err)
			return nil, fmt.Errorf("can't scan user: %w", err)
		}
		u = append(u, user)
	}

	return u, nil
}

func (d *Storage) Feed(ctx context.Context, userId models.UserId, page int) (data []models.Post, err error) {
	rows, err := d.Conn.Query(ctx, `SELECT data FROM feed 
            WHERE user_id = $1 OFFSET $2*50 LIMIT 50`, userId, page)
	if err != nil {
		return nil, fmt.Errorf("get feed error: %w", err)
	}

	var entity string
	_, err = pgx.ForEachRow(rows, []any{&entity}, func() error {
		p := models.Post{}
		err := json.Unmarshal([]byte(entity), &p)

		if err != nil {
			d.Conn.Logger().Errorf("json broken: %+v, err: %+v", entity, err)
			return nil
		}

		data = append(data, p)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("")
	}

	data1 := make([]models.Post, len(data))
	copy(data1, data)

	return data1, nil
}

func (d *Storage) saveToFeed(ctx context.Context, userId models.UserId, posts ...*models.Post) error {
	t := time.Now().UTC()

	n, err := d.Conn.CopyFrom(ctx, pgx.Identifier{"feed"}, []string{"user_id", "data", "timestamp"},
		pgx.CopyFromSlice(len(posts), func(i int) ([]any, error) {
			b, _ := json.Marshal(posts[i])
			return []any{userId, string(b), t}, nil
		}))

	if err != nil {
		return fmt.Errorf("save to feed fail: %w", err)
	}

	d.Conn.Logger().Debugf("added %d rows to feed", n)

	return nil
}
