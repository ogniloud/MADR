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

// SignUpResponse is a struct that defines the response body for the sign-up endpoint.
type SignUpResponse struct {
	// ID of the user.
	//
	// example: 1
	ID int `json:"id"`

	// Email of the user.
	//
	// example: user@example.com
	Email string `json:"email"`
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
