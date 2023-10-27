// Package models describes models for application.
//
// Documentation for MADR API.
//
// Schemes: http
// BasePath: /
// Version: 0.0.1
// Contact: Dmitriy Krasnov<dk.peyuaa@gmail.com>
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package models

// swagger:response signUpCreatedResponse
type SwaggerSignUpCreatedResponse struct {
	// Response with the credentials of the user.
	// in: body
	Body SignUpResponse
}

// swagger:response signUpBadRequestError
type SwaggerSignUpBadRequestError struct {
	// in: body
	Body GenericError
}

// swagger:response signUpConflictError
type SwaggerSignUpConflictError struct {
	// in: body
	Body GenericError
}

// swagger:response signUpInternalServerError
type SwaggerInternalServerError struct {
	// in: body
	Body GenericError
}

// swagger:response signInOkResponse
type SwaggerSignInOkResponse struct {
	// Response with the authorization token.
	// in: body
	Body SignInResponse
}

// swagger:response signInBadRequestError
type SwaggerSignInBadRequestError struct {
	// in: body
	Body GenericError
}

// swagger:response signInUnauthorizedError
type SwaggerSignInUnauthorizedError struct {
	// in: body
	Body GenericError
}

// swagger:response signInInternalServerError
type SwaggersignInInternalServerError struct {
	// in: body
	Body GenericError
}
