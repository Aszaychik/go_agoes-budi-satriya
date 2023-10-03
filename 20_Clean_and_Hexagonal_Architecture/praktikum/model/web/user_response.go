package web

type UserCreateResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}