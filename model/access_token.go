package model

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

// tokens information
const (
	AccessTokenType       = "bearer"
	AccessTokenGrantType  = "access_token"
	RefreshTokenGrantType = "refresh_token"
)

// AccessToken is the user access token
type AccessToken struct {
	ID          uuid.UUID `json:"id"`
	Token       string    `json:"token,omitempty"`
	UserID      string    `json:"user_id"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
}

// AccessData holds auth access information
type AccessData struct {
	ClientID     string `json:"client_id"`
	UserID       string `json:"user_id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	RedirectURI  string `json:"redirect_uri"`
	ExpiresAt    int    `json:"expires_at"`
	Scope        string `json:"scope"`
}

// AccessResponse is the response sent to the client
type AccessResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// PreSave sets the uuid and isActive flag
func (t *AccessToken) PreSave() {
	t.ID = uuid.New()
	t.IsActive = true
}

// ToJSON converts token to json string
func (t *AccessToken) ToJSON() string {
	b, _ := json.Marshal(t)
	return string(b)
}

// AccessTokenFromJSON decodes the input and returns the AccessToken
func AccessTokenFromJSON(data io.Reader) (*AccessToken, error) {
	var t *AccessToken
	err := json.NewDecoder(data).Decode(&t)
	return t, err
}
