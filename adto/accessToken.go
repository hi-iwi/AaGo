package adto

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"` // expires in seconds
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"` // OAuth 2.0 / Bearer
	Scopes       string `json:"scopes"`
}
