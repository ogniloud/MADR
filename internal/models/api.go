package models

// swagger:model signUpRequest
// SignUpRequest is a struct that defines the request body for the sign-up endpoint.
type SignUpRequest struct {
	// Username of the user.
	//
	// required: true
	// example: user123
	Username string `json:"username"`

	// Email of the user.
	//
	// required: true
	// example: user@example.com
	Email string `json:"email"`

	// Password of the user.
	//
	// required: true
	// example: myVerySecurePassword123
	Password string `json:"password"`
}

// swagger:model signInRequest
// SignInRequest is a struct that defines the request body for the sign-in endpoint.
type SignInRequest struct {
	// Username of the user.
	//
	// required: true
	// example: user123
	Username string `json:"username"`

	// Password of the user.
	//
	// required: true
	// example: myVerySecurePassword123
	Password string `json:"password"`
}

// SignInResponse is a struct that defines the response body for the sign-in endpoint.
type SignInResponse struct {
	// Authorization token.
	//
	// example: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjkxOTU0Nzk4MDksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoidXNlcjEyMyJ9.fHSoS6ZFf1TU4AmcqNeqpEDo6hdU6uLr2-PRAd0MKzAKDvDtGafuV6X6W8HSXAgwraXZ0_3qS8CmrUQW6am8Hg
	Authorization string `json:"authorization"`
}
