package model

import "time"

// User represents the shop user model
type User struct {
	ID            uint64    `json:"id"`
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	UserName      string    `json:"user_name"`
	Email         string    `json:"email"`
	Password      string    `json:"-"`
	EmailVerified bool      `json:"email_verified"`
	IsActive      bool      `json:"is_active"`
	LastLoginAt   time.Time `json:"last_login_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"-"`
}
