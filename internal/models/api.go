package models

// SignUpRequest is a struct that defines the request body for the sign-up endpoint.
type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpResponse is a struct that defines the response body for the sign-up endpoint.
type SignUpResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// SignInRequest is a struct that defines the request body for the sign-in endpoint.
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignInResponse is a struct that defines the response body for the sign-in endpoint.
type SignInResponse struct {
	Authorization string `json:"authorization"`
}