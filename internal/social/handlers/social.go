package handlers

import (
	"net/http"

	"github.com/charmbracelet/log"

	"github.com/ogniloud/madr/internal/ioutil"
	"github.com/ogniloud/madr/internal/social/models"
	"github.com/ogniloud/madr/internal/social/storage"
)

type IService interface {
	storage.Storage
}

type Endpoints struct {
	logger *log.Logger
	ew     ioutil.ErrorWriter
	s      IService
}

func New(s IService, ew ioutil.ErrorWriter) Endpoints {
	return Endpoints{
		ew: ew,
		s:  s,
	}
}

// Followers /api/social/followers
func (e Endpoints) Followers(w http.ResponseWriter, r *http.Request) {
	reqBody := models.FollowersRequest{}
	respBody := models.FollowersResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	infos, err := e.s.GetFollowersByUserId(r.Context(), reqBody.UserId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.UserInfo = infos
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Followings /api/social/followings
func (e Endpoints) Followings(w http.ResponseWriter, r *http.Request) {
	reqBody := models.FollowingsRequest{}
	respBody := models.FollowingsResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	infos, err := e.s.GetFollowingsByUserId(r.Context(), reqBody.UserId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.UserInfo = infos
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e Endpoints) Follow(w http.ResponseWriter, r *http.Request) {
	reqBody := models.FollowRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reqBody.AuthorId == reqBody.FollowerId {
		e.ew.Error(w, "author_id is equal follower_id", http.StatusBadRequest)
		return
	}

	err := e.s.Follow(r.Context(), reqBody.FollowerId, reqBody.AuthorId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (e Endpoints) Unfollow(w http.ResponseWriter, r *http.Request) {
	reqBody := models.FollowRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reqBody.AuthorId == reqBody.FollowerId {
		e.ew.Error(w, "author_id is equal follower_id", http.StatusBadRequest)
		return
	}

	err := e.s.Unfollow(r.Context(), reqBody.FollowerId, reqBody.AuthorId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
