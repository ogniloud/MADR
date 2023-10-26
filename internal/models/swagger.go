package models

// swagger:response signUpResponse
type SwaggerSignUpResponse struct {
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
