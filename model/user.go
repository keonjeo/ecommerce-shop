package model

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	BCRYPT_COST               = 14
	USER_EMAIL_MAX_LENGTH     = 128
	USER_PASSWORD_MAX_LENGTH  = 72
	USER_USERNAME_MAX_RUNES   = 64
	USER_POSITION_MAX_RUNES   = 128
	USER_FIRST_NAME_MAX_RUNES = 64
	USER_LAST_NAME_MAX_RUNES  = 64
	USER_AUTH_DATA_MAX_LENGTH = 128
	USER_NAME_MAX_LENGTH      = 64
	USER_NAME_MIN_LENGTH      = 1
)

var reservedNames = []string{
	"api",
	"admin",
	"signup",
	"login",
	"oauth",
	"error",
	"help",
}

// User represents the shop user model
type User struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Gender         string    `json:"gender"`
	AvatarURL      string    `json:"avatar_url"`
	Active         bool      `json:"active"`
	EmailVerified  bool      `json:"email_verified"`
	FailedAttempts int       `json:"failed_attempts"`
	LastLoginAt    time.Time `json:"last_login_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"-"`
}

// ToJSON converts user to json string
func (u *User) ToJSON() string {
	b, _ := json.Marshal(u)
	return string(b)
}

// UserFromJSON decodes the input and return a User
func UserFromJSON(data io.Reader) (*User, error) {
	var user *User
	err := json.NewDecoder(data).Decode(&user)
	return user, err
}

// Validate checks if user is valid
func (u *User) Validate() error {
	return v.ValidateStruct(u,
		v.Field(&u.Email, v.Required, is.Email),
		v.Field(&u.Password, v.Required),
	)
}

// Validate checks if user is valid
func (u *User) _Validate() *AppError {
	if len(u.Email) > USER_EMAIL_MAX_LENGTH || len(u.Email) == 0 {
		return NewAppError("GetUsers", "api.user.get_users.invalid_email.app_error", nil, "", http.StatusInternalServerError)

	}
	if len(u.Password) > USER_PASSWORD_MAX_LENGTH || len(u.Password) == 0 {
		return NewAppError("GetUsers", "api.user.get_users.invalid_password.app_error", nil, "", http.StatusInternalServerError)
	}

	return nil

	// if len(u.ID) <= 0 {
	// 	return InvalidUserError("id", "")
	// }

	// if u.CreatedAt == 0 {
	// 	return InvalidUserError("created_at", u.Id)
	// }

	// if u.UpdatedAt == 0 {
	// 	return InvalidUserError("updated_at", u.Id)
	// }

	// if !IsValidUsername(u.Username) {
	// 	return InvalidUserError("username", u.Id)
	// }

	// if len(u.Email) > USER_EMAIL_MAX_LENGTH || len(u.Email) == 0 || !IsValidEmail(u.Email) {
	// 	return InvalidUserError("email", u.Id)
	// }

	// if utf8.RuneCountInString(u.FirstName) > USER_FIRST_NAME_MAX_RUNES {
	// 	return InvalidUserError("first_name", u.Id)
	// }

	// if utf8.RuneCountInString(u.LastName) > USER_LAST_NAME_MAX_RUNES {
	// 	return InvalidUserError("last_name", u.Id)
	// }

	// if u.AuthData != nil && len(*u.AuthData) > USER_AUTH_DATA_MAX_LENGTH {
	// 	return InvalidUserError("auth_data", u.Id)
	// }

	// if u.AuthData != nil && len(*u.AuthData) > 0 && len(u.AuthService) == 0 {
	// 	return InvalidUserError("auth_data_type", u.Id)
	// }

	// if len(u.Password) > 0 && u.AuthData != nil && len(*u.AuthData) > 0 {
	// 	return InvalidUserError("auth_data_pwd", u.Id)
	// }

	// if len(u.Password) > USER_PASSWORD_MAX_LENGTH {
	// 	return InvalidUserError("password_limit", u.Id)
	// }

	// return nil
}

func (u *User) Sanitize() {}

func (u *User) PreSave() {}

func (u *User) PreUpdate() {}

// InvalidUserError creates user specific AppError
// func InvalidUserError(fieldName string, userID string) *AppError {
// 	id := fmt.Sprintf("model.user.is_valid.%s.app_error", fieldName)
// 	details := ""
// 	if userID != "" {
// 		details = "user_id=" + userID
// 	}
// 	return NewAppError("User.IsValid", id, "ErrInvalidUser", details, nil, http.StatusBadRequest)

// 	// model.NewAppError("GetUsers", "api.user.get_users", "ErrGetUsers", "could not get users", map[string]interface{}{"some": "stuff"}, http.StatusInternalServerError))
// }

// NormalizeUsername trims space and returns lowercase username
func NormalizeUsername(username string) string {
	return strings.ToLower(strings.TrimSpace(username))
}

// NormalizeEmail trims space and returns lowercase email
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// HashPassword generates a hash using bcrypt
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), BCRYPT_COST)
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
