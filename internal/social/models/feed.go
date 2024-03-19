package models

import (
	"encoding/json"
	"fmt"
)

type FeedPostType int

const (
	Invite = FeedPostType(iota)
	SharedFromGroup
	SharedFromFollowing
	FollowingSubscribed
)

type FeedData interface {
	Type() FeedPostType
}

type Post struct {
	// Type is a number for type of the data. Use Data.Type()
	// to set correct number to the field.
	Type FeedPostType `json:"type"`
	Data FeedData     `json:"data"`
}

func (p *Post) UnmarshalJSON(b []byte) error {
	m := map[string]any{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	temp, ok := m["type"]
	if !ok {
		return &json.UnmarshalTypeError{Field: "type"}
	}

	t, ok := temp.(float64)
	if !ok {
		return &json.UnmarshalTypeError{Value: fmt.Sprint(temp), Field: "type"}
	}

	data, ok := m["data"]
	if !ok {
		return &json.UnmarshalTypeError{Field: "data"}
	}

	bb, _ := json.Marshal(data)

	p.Type = FeedPostType(t)

	switch p.Type {
	case Invite:
		data := InviteData{}
		if err := json.Unmarshal(bb, &data); err != nil {
			return err
		}
		p.Data = data
	case SharedFromGroup:
		data := SharedFromGroupData{}
		if err := json.Unmarshal(bb, &data); err != nil {
			return err
		}
		p.Data = data
	case SharedFromFollowing:
		data := SharedFromFollowingData{}
		if err := json.Unmarshal(bb, &data); err != nil {
			return err
		}
		p.Data = data
	case FollowingSubscribed:
		data := FollowingSubscribedData{}
		if err := json.Unmarshal(bb, &data); err != nil {
			return err
		}
		p.Data = data
	}

	return nil
}

type InviteData struct {
	InviteeId   UserId  `json:"invitee_id"`
	InviteeName string  `json:"invitee_name"`
	GroupId     GroupId `json:"group_id"`
	GroupName   string  `json:"group_name"`
}

func (d InviteData) Type() FeedPostType {
	return Invite
}

type SharedFromGroupData struct {
	GroupId   GroupId `json:"group_id"`
	GroupName string  `json:"group_name"`
	DeckId    DeckId  `json:"deck_id"`
	DeckName  string  `json:"deck_name"`
}

func (d SharedFromGroupData) Type() FeedPostType {
	return SharedFromGroup
}

type SharedFromFollowingData struct {
	AuthorId   UserId `json:"author_id"`
	AuthorName string `json:"group_name"`
	DeckId     DeckId `json:"deck_id"`
	DeckName   string `json:"deck_name"`
}

func (d SharedFromFollowingData) Type() FeedPostType {
	return SharedFromFollowing
}

type FollowingSubscribedData struct {
	FollowerId   UserId `json:"following_id"`
	FollowerName string `json:"following_name"`
	AuthorId     UserId `json:"author_id"`
	AuthorName   string `json:"author_name"`
}

func (d FollowingSubscribedData) Type() FeedPostType {
	return FollowingSubscribed
}
