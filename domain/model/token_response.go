package model

// TokenResponse is a struct that handle the "unmarshall" information from HTTP Response Body when asking for a new Token.
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}