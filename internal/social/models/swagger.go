package models

import "github.com/ogniloud/madr/internal/ioutil"

// swagger:response getFollowersOkResponse
// SwaggerGetFollowersOkResponse is a struct that defines the response body for the getFollowers endpoint.
type SwaggerGetFollowersOkResponse struct {
	// in: body
	Body FollowersResponse
}

// swagger:response getFollowersBadRequestResponse
// SwaggerGetFollowersBadRequestResponse is a struct that defines the response body for the getFollowers endpoint.
// Returns a bad request error.
type SwaggerGetFollowersBadRequestResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response getFollowersInternalServerErrorResponse
// SwaggerGetFollowersInternalServerErrorResponse is a struct that defines the response body for the getFollowers endpoint.
// Returns an internal server error.
type SwaggerGetFollowersInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response getFollowingsOkResponse
// SwaggerGetFollowingsOkResponse is a struct that defines the response body for the Followings endpoint.
type SwaggerFollowingsOkResponse struct {
	// in: body
	Body FollowingsResponse
}

// swagger:response getFollowingsBadRequestResponse
// SwaggerGetFollowingsBadRequestResponse is a struct that defines the response body for the Followings endpoint.
type SwaggerFollowingsBadRequestResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response getFollowingsInternalServerErrorResponse
// SwaggerGetFollowingsInternalServerErrorResponse is a struct that defines the response body for the Followings endpoint.
type SwaggerFollowingsInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response followNoContentResponse
// SwaggerFollowNoContentResponse is a struct that defines the response body for the Follow endpoint.
type SwaggerFollowNoContentResponse struct{}

// swagger:response followBadRequestResponse
// SwaggerFollowBadRequestResponse is a struct that defines the response body for the Follow endpoint.
type SwaggerFollowBadRequestResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response followInternalServerErrorResponse
// SwaggerFollowInternalServerErrorResponse is a struct that defines the response body for the Follow endpoint.
type SwaggerFollowInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response unfollowNoContentResponse
// SwaggerUnfollowNoContentResponse is a struct that defines the response body for the Unfollow endpoint.
type SwaggerUnfollowNoContentResponse struct{}

// swagger:response unfollowBadRequestResponse
// SwaggerUnfollowBadRequestResponse is a struct that defines the response body for the Unfollow endpoint.
type SwaggerUnfollowBadRequestResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response unfollowInternalServerErrorResponse
// SwaggerUnfollowInternalServerErrorResponse is a struct that defines the response body for the Unfollow endpoint.
type SwaggerUnfollowInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response searchUserOkResponse
// SwaggerSearchUserOkResponse is a struct that defines the response body for the SearchUser endpoint.
type SwaggerSearchUserOkResponse struct {
	// in: body
	Body SearchUserResponse
}

// swagger:response searchUserInternalServerErrorResponse
// SwaggerSearchUserInternalServerErrorResponse is a struct that defines the response body for the SearchUser endpoint.
type SwaggerSearchUserInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response createGroupOkResponse
// SwaggerCreateGroupOkResponse is a struct that defines the response body for the CreateGroup endpoint.
type SwaggerCreateGroupOkResponse struct {
	// in: body
	Body CreateGroupResponse
}

// swagger:response createGroupBadRequestResponse
// SwaggerCreateGroupBadRequestResponse is a struct that defines the response body for the CreateGroup endpoint.
type SwaggerCreateGroupBadRequestResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response createGroupInternalServerErrorResponse
// SwaggerCreateGroupInternalServerErrorResponse is a struct that defines the response body for the CreateGroup endpoint.
type SwaggerCreateGroupInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response shareGroupDeckNoContentResponse
// SwaggerShareGroupDeckNoContentResponse is a struct that defines the response body for the ShareGroupDeck endpoint.
type SwaggerShareGroupDeckNoContentResponse struct{}

// swagger:response shareGroupDeckBadRequestResponse
// SwaggerShareGroupDeckBadRequestResponse is a struct that defines the response body for the ShareGroupDeck endpoint.
type SwaggerShareGroupDeckBadRequestResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response shareGroupDeckInternalServerErrorResponse
// SwaggerShareGroupDeckInternalServerErrorResponse is a struct that defines the response body for the ShareGroupDeck endpoint.
type SwaggerShareGroupDeckInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response getDecksByGroupIdOkResponse
// SwaggerGetDecksByGroupIdOkResponse is a struct that defines the response body for the GetDecksByGroupId endpoint.
type SwaggerGetDecksByGroupIdOkResponse struct {
	// in: body
	Body GetDecksByGroupIdResponse
}

// swagger:response getDecksByGroupIdBadRequestResponse
// SwaggerGetDecksByGroupIdBadRequestResponse is a struct that defines the response body for the GetDecksByGroupId endpoint.
type SwaggerGetDecksByGroupIdBadRequestResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response getDecksByGroupIdInternalServerErrorResponse
// SwaggerGetDecksByGroupIdInternalServerErrorResponse is a struct that defines the response body for the GetDecksByGroupId endpoint.
type SwaggerGetDecksByGroupIdInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response getGroupsByUserIdOkResponse
// SwaggerGetGroupsByUserIdOkResponse is a struct that defines the response body for the GetGroupsByUserId endpoint.
type SwaggerGetGroupsByUserIdOkResponse struct {
	// in: body
	Body GetGroupsByUserIdResponse
}

// swagger:response getGroupsByUserIdBadRequestResponse
// SwaggerGetGroupsByUserIdBadRequestResponse is a struct that defines the response body for the GetGroupsByUserId endpoint.
type SwaggerGetGroupsByUserIdBadRequestResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response getGroupsByUserIdInternalServerErrorResponse
// SwaggerGetGroupsByUserIdInternalServerErrorResponse is a struct that defines the response body for the GetGroupsByUserId endpoint.
type SwaggerGetGroupsByUserIdInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response feedOkResponse
// SwaggerFeedOkResponse is a struct that defines the response body for the Feed endpoint.
type SwaggerFeedOkResponse struct {
	// in: body
	Body FeedResponse
}

// swagger:response feedBadRequestResponse
// SwaggerFeedBadRequestResponse is a struct that defines the response body for the Feed endpoint.
type SwaggerFeedBadRequestResponse struct {
	// in: body
	Body ioutil.GenericError
}

// swagger:response feedInternalServerErrorResponse
// SwaggerFeedInternalServerErrorResponse is a struct that defines the response body for the Feed endpoint.
type SwaggerFeedInternalServerErrorResponse struct {
	// in: body
	Body ioutil.GenericError
}
