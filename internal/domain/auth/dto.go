package auth

type TokenInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenOutput struct {
	AccessToken string `json:"access_token"`
}
