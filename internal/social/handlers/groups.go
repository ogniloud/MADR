package handlers

import (
	"net/http"

	"github.com/ogniloud/madr/internal/ioutil"
	"github.com/ogniloud/madr/internal/social/models"
)

// swagger:route POST /api/group/share ShareGroupDeck
// Share a deck with a group.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http
//
//
//	Parameters:
//	+ name: request
//	  in: body
//	  required: true
//	  type: shareGroupDeckRequest
//
//	Responses:
//	204: shareGroupDeckNoContentResponse
//	400: shareGroupDeckBadRequestResponse
//	500: shareGroupDeckInternalServerErrorResponse
func (e Endpoints) ShareGroupDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.ShareGroupDeckRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.s.ShareDeckGroup(r.Context(), reqBody.UserId, reqBody.GroupId, reqBody.DeckId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route POST /api/group/decks GetDecksByGroupId
// Get decks by group id.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http
//
//
//	Parameters:
//	+ name: request
//	  in: body
//	  required: true
//	  type: getDecksByGroupIdRequest
//
//	Responses:
//	200: getDecksByGroupIdOkResponse
//	400: getDecksByGroupIdBadRequestResponse
//	500: getDecksByGroupIdInternalServerErrorResponse
func (e Endpoints) GetDecksByGroupId(w http.ResponseWriter, r *http.Request) {
	reqBody := models.GetDecksByGroupIdRequest{}
	respBody := models.GetDecksByGroupIdResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decks, err := e.s.GetDecksByGroupId(r.Context(), reqBody.GroupId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Decks = decks
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteGroupDeck /api/groups/delete_deck
func (e Endpoints) DeleteGroupDeck(w http.ResponseWriter, r *http.Request) {
	reqBody := models.DeleteGroupDeckRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.s.DeleteDeckFromGroup(r.Context(), reqBody.UserId, reqBody.GroupId, reqBody.DeckId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SearchGroupByName GET /api/groups/search?q=...
func (e Endpoints) SearchGroupByName(w http.ResponseWriter, r *http.Request) {
	respBody := models.SearchGroupByNameResponse{}

	name := r.Form.Get("q")

	groups, err := e.s.GetGroupsByName(r.Context(), name)
	if err != nil {
		e.logger.Errorf("query: %+v, error: %v", r.Form, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Groups = groups
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route PUT /api/groups/create CreateGroup
// Create a new group.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http
//
//
//	Parameters:
//	+ name: request
//	  in: body
//	  required: true
//	  type: createGroupRequest
//
//	Responses:
//	200: createGroupOkResponse
//	400: createGroupBadRequestResponse
//	500: createGroupInternalServerErrorResponse
func (e Endpoints) CreateGroup(w http.ResponseWriter, r *http.Request) {
	reqBody := models.CreateGroupRequest{}
	respBody := models.CreateGroupResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := e.s.CreateGroup(r.Context(), reqBody.UserId, reqBody.Name)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.GroupId = id
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/groups/groups GetGroupsByUserId
// Get groups by user id.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http
//
//
//	Parameters:
//	+ name: request
//	  in: body
//	  required: true
//	  type: getGroupsByUserIdRequest
//
//	Responses:
//	200: getGroupsByUserIdOkResponse
//	400: getGroupsByUserIdBadRequestResponse
//	500: getGroupsByUserIdInternalServerErrorResponse
func (e Endpoints) GetGroupsByUserId(w http.ResponseWriter, r *http.Request) {
	reqBody := models.GetGroupsByUserIdRequest{}
	respBody := models.GetGroupsByUserIdResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ids, err := e.s.GetGroupsByUserId(r.Context(), reqBody.UserId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Groups = ids
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetCreatedGroupsByUserId POST /api/groups/created_groups
func (e Endpoints) GetCreatedGroupsByUserId(w http.ResponseWriter, r *http.Request) {
	reqBody := models.GetCreatedGroupsByUserIdRequest{}
	respBody := models.GetCreatedGroupsByUserIdResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ids, err := e.s.GetCreatedGroupsByUserId(r.Context(), reqBody.UserId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Groups = ids
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// SendInvite /api/invite/send POST
func (e Endpoints) SendInvite(w http.ResponseWriter, r *http.Request) {
	reqBody := models.SendInviteRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.s.SendInvite(r.Context(), reqBody.CreatorId, reqBody.InviteeId, reqBody.GroupId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AcceptInvite /api/invite/accept POST
func (e Endpoints) AcceptInvite(w http.ResponseWriter, r *http.Request) {
	reqBody := models.AcceptInviteRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.s.AcceptInvite(r.Context(), reqBody.UserId, reqBody.GroupId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ChangeGroupName PUT /api/groups/change_name
func (e Endpoints) ChangeGroupName(w http.ResponseWriter, r *http.Request) {
	reqBody := models.ChangeGroupNameRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.s.ChangeGroupName(r.Context(), reqBody.CreatorId, reqBody.GroupId, reqBody.Name)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteGroup DELETE /api/groups/delete
func (e Endpoints) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	reqBody := models.DeleteGroupRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.s.DeleteGroup(r.Context(), reqBody.UserId, reqBody.GroupId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// QuitGroup DELETE /api/groups/quit
func (e Endpoints) QuitGroup(w http.ResponseWriter, r *http.Request) {
	reqBody := models.QuitGroupRequest{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := e.s.QuitGroup(r.Context(), reqBody.UserId, reqBody.GroupId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (e Endpoints) GetParticipantsByGroupId(w http.ResponseWriter, r *http.Request) {
	reqBody := models.GetParticipantsByGroupIdRequest{}
	respBody := models.GetParticipantsByGroupIdResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	infos, err := e.s.GetParticipantsByGroupId(r.Context(), reqBody.GroupId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Participants = infos
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e Endpoints) GetGroupsDeckShared(w http.ResponseWriter, r *http.Request) {
	reqBody := models.GetGroupsDeckSharedRequest{}
	respBody := models.GetGroupsDeckSharedResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groups, err := e.s.GetGroupsDeckShared(r.Context(), reqBody.CreatorId, reqBody.DeckId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Groups = groups
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e Endpoints) GetFollowersNotJoinedGroup(w http.ResponseWriter, r *http.Request) {
	reqBody := models.GetFollowersNotJoinedGroupRequest{}
	respBody := models.GetFollowersNotJoinedGroupResponse{}

	if err := ioutil.FromJSON(&reqBody, r.Body); err != nil {
		e.logger.Errorf("json not parsed: %v", err)
		e.ew.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	followers, err := e.s.GetFollowersNotJoinedGroup(r.Context(), reqBody.CreatorId, reqBody.GroupId)
	if err != nil {
		e.logger.Errorf("reqBody: %+v, error: %v", reqBody, err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody.Followers = followers
	if err := ioutil.ToJSON(respBody, w); err != nil {
		e.logger.Errorf("To json error: %v", err)
		e.ew.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
