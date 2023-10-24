package models

// SignUpRequest is a struct that defines the request body for the sign-up endpoint.
type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// SignInRequest is a struct that defines the request body for the sign-in endpoint.
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
