package models

// swagger:model signUpRequest
// SignUpRequest is a struct that defines the request body for the sign-up endpoint.
type SignUpRequest struct {
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

// SignInResponse is a struct that defines the response body for the sign-in endpoint.
type SignInResponse struct {
	// Authorization token.
	//
	// example: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9s
	Authorization string `json:"authorization"`
}
