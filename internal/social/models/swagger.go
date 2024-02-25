package models

import "github.com/ogniloud/madr/internal/ioutil"

// swagger:model getFollowersRequest
// SwaggerGetFollowersRequest is a struct that defines the request body for the getFollowers endpoint.
type SwaggerGetFollowersRequest struct {
	// in: body
	Body FollowersRequest
}

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

// swagger:model getFollowingsRequest
// SwaggerGetFollowingsRequest is a struct that defines the request body for the Followings endpoint.
type SwaggerFollowingsRequest struct {
	// in: body
	Body FollowingsRequest
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

// swagger:model followRequest
// SwaggerFollowRequest is a struct that defines the request body for the Follow endpoint.
type SwaggerFollowersRequest struct {
	// in: body
	Body FollowRequest
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

// swagger:model unfollowRequest
// SwaggerUnfollowRequest is a struct that defines the request body for the Unfollow endpoint.
type SwaggerUnfollowRequest struct {
	// in: body
	Body UnfollowRequest
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
