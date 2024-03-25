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

func New(s IService, ew ioutil.ErrorWriter, logger *log.Logger) Endpoints {
	return Endpoints{
		ew:     ew,
		s:      s,
		logger: logger,
	}
}

// swagger:route POST /api/social/followers getFollowers
// Returns a list of followers of the user.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Scheme: http
//
//
//	Parameters:
//	+ name: request
//	  in: body
//	  description: request body
//	  required: true
//	  type: getFollowersRequest
//
//	Responses:
//	200: getFollowersOkResponse
//	400: getFollowersBadRequestResponse
//	500: getFollowersInternalServerErrorResponse
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

// swagger:route POST /api/social/followings getFollowings
// Returns a list of followings of the user.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Scheme: http
//
//
//	Parameters:
//	+ name: request
//	  in: body
//	  description: request body
//	  required: true
//	  type: getFollowingsRequest
//
//	Responses:
//	200: getFollowingsOkResponse
//	400: getFollowingsBadRequestResponse
//	500: getFollowingsInternalServerErrorResponse
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

// swagger:route POST /api/social/follow follow
// Follows the user.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Scheme: http
//
//
//	Parameters:
//	+ name: request
//	  in: body
//	  description: request body
//	  required: true
//	  type: followRequest
//
//	Responses:
//	204: followNoContentResponse
//	400: followBadRequestResponse
//	500: followInternalServerErrorResponse
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

// swagger:route POST /api/social/unfollow unfollow
// Unfollows the user.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Scheme: http
//
//
//	Parameters:
//	+ name: request
//	  in: body
//	  description: request body
//	  required: true
//	  type: followRequest
//
//	Responses:
//	204: unfollowNoContentResponse
//	400: unfollowBadRequestResponse
//	500: unfollowInternalServerErrorResponse
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

// DeepCopyDeck /api/social/copy
func (e Endpoints) DeepCopyDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.DeepCopyDeckRequest{}
	respBody := models.DeepCopyDeckResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deckId, err := e.s.DeepCopyDeck(r.Context(), reqBody.CopierId, reqBody.DeckId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.DeckId = deckId
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route GET /api/social/search search
// Returns a list of users by name.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Scheme: http
//
//
//	Parameters:
//	+ name: q
//	  in: query
//	  description: search query
//	  required: true
//	  type: string
//
//	Responses:
//	200: searchUserOkResponse
//	500: searchUserInternalServerErrorResponse
func (e Endpoints) SearchUser(w http.ResponseWriter, r *http.Request) {
	respBody := models.SearchUserResponse{}

	err := r.ParseForm()
	if err != nil {
		e.logger.Errorf("parse form error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := r.Form.Get("q")

	users, err := e.s.GetUsersByName(r.Context(), name)
	if err != nil {
		e.logger.Errorf("query: %+v, error: %v", r.Form, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, user := range users {
		respBody.Users = append(respBody.Users, models.UserInfo(user))
	}

	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/social/feed Feed
// Returns user's feed.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Scheme: http
//
//
//	Parameters:
//	+ name: request
//	  in: body
//	  description: request body
//	  required: true
//	  type: feedRequest
//
//	Responses:
//	200: feedOkResponse
//	400: feedBadRequestResponse
//	500: feedInternalServerErrorResponse
func (e Endpoints) Feed(w http.ResponseWriter, r *http.Request) {
	reqBody := models.FeedRequest{}
	respBody := models.FeedResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	feed, err := e.s.Feed(r.Context(), reqBody.UserId, 0)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Feed = feed
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
