package model

import (
	"encoding/json"
	"io"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents the shop user model
type User struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Gender      string    `json:"gender"`
	AvatarURL   string    `json:"avatar_url"`
	IsVerified  bool      `json:"is_verified"`
	IsActive    bool      `json:"is_active"`
	LastLoginAt time.Time `json:"last_login_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"-"`
}

// ToJSON converts user to json string
func (u *User) ToJSON() string {
	b, _ := json.Marshal(u)
	return string(b)
}

// UserFromJSON decodes the input and return a User
func UserFromJSON(data io.Reader) *User {
	var user *User
	json.NewDecoder(data).Decode(&user)
	return user
}

// HashPassword generates a hash using bcrypt
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

// ComparePassword compares the hash
func ComparePassword(hash string, password string) bool {
	if len(password) == 0 || len(hash) == 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
