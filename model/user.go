package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/language"

	"github.com/dankobgd/ecommerce-shop/config"
	"github.com/dankobgd/ecommerce-shop/utils/is"
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
	USER_USERNAME_MAX_LENGTH  = 64
	USER_USERNAME_MIN_LENGTH  = 1
	USER_LOCALE_MAX_LENGTH    = 5
	USER_DEFAULT_LOCALE       = "en"

	NUMBERS           = "0123456789"
	SYMBOLS           = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"
	LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var reservedNames = []string{"app", "api", "admin", "signup", "login", "oauth", "error", "help"}
var restrictedUsernames = []string{"app", "api", "admin", "system"}
var validUsernameChars = regexp.MustCompile(`^[a-z0-9\.\-_]+$`)

// User represents the shop user model
type User struct {
	ID             int       `json:"id" db:"id"`
	FirstName      string    `json:"first_name" db:"first_name"`
	LastName       string    `json:"last_name" db:"last_name"`
	Username       string    `json:"username" db:"username"`
	Email          string    `json:"email" db:"email"`
	Password       string    `json:"password" db:"password"`
	Gender         string    `json:"gender" db:"gender"`
	Locale         string    `json:"locale" db:"locale"`
	AvatarURL      string    `json:"avatar_url" db:"avatar_url"`
	Active         bool      `json:"active" db:"active"`
	EmailVerified  bool      `json:"email_verified" db:"email_verified"`
	FailedAttempts int       `json:"failed_attempts" db:"failed_attempts"`
	LastLoginAt    time.Time `json:"last_login_at" db:"last_login_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at" db:"deleted_at"`
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

// NewInvalidUserError builds the invalid user error
func NewInvalidUserError(fieldName string, userID int) *AppError {
	id := fmt.Sprintf("model.user.validate.%s.app_error", fieldName)
	details := ""
	if userID != 0 {
		details = fmt.Sprintf("userID: %d", userID)
	}
	return NewAppError("User.Validate", id, nil, details, http.StatusBadRequest)
}

// IsValidPasswordCriteria checks if password fulfills the criteria
func IsValidPasswordCriteria(password string, settings *config.PasswordSettings) *AppError {
	id := "model.user.validate.password"
	isError := false

	if len(password) < settings.MinLength || len(password) > settings.MaxLength {
		isError = true
	}
	if settings.Lowercase {
		if !strings.ContainsAny(password, LOWERCASE_LETTERS) {
			isError = true
		}
		id = id + "_lowercase"
	}
	if settings.Uppercase {
		if !strings.ContainsAny(password, UPPERCASE_LETTERS) {
			isError = true
		}
		id = id + "_uppercase"
	}
	if settings.Number {
		if !strings.ContainsAny(password, NUMBERS) {
			isError = true
		}
		id = id + "_number"
	}
	if settings.Symbol {
		if !strings.ContainsAny(password, SYMBOLS) {
			isError = true
		}
		id = id + "_symbol"
	}
	if isError {
		return NewAppError("User.Validate", id+".app_error", map[string]interface{}{"Min": settings.MinLength}, "", http.StatusBadRequest)
	}
	return nil
}

// IsValidUsername validates whether username matches the criteria
func IsValidUsername(username string) bool {
	if len(username) < USER_USERNAME_MIN_LENGTH || len(username) > USER_USERNAME_MAX_LENGTH {
		return false
	}
	if !validUsernameChars.MatchString(username) {
		return false
	}
	for _, run := range restrictedUsernames {
		if username == run {
			return false
		}
	}
	return true
}

// IsValidLocale checks if locale is valid
func IsValidLocale(locale string) bool {
	if locale != "" {
		if len(locale) > USER_LOCALE_MAX_LENGTH {
			return false
		} else if _, err := language.Parse(locale); err != nil {
			return false
		}
	}
	return true
}

// Validate validates the user and returns an error if it doesn't pass criteria
func (u *User) Validate() *AppError {
	if u.ID < 0 {
		return NewInvalidUserError("id", u.ID)
	}
	if u.CreatedAt.IsZero() {
		return NewInvalidUserError("created_at", u.ID)
	}
	if u.UpdatedAt.IsZero() {
		return NewInvalidUserError("updated_at", u.ID)
	}
	if !IsValidUsername(u.Username) {
		return NewInvalidUserError("username", u.ID)
	}
	if len(u.Email) > USER_EMAIL_MAX_LENGTH || len(u.Email) == 0 || !is.ValidEmail(u.Email) {
		return NewInvalidUserError("email", u.ID)
	}
	if utf8.RuneCountInString(u.Username) > USER_USERNAME_MAX_RUNES {
		return NewInvalidUserError("nickname", u.ID)
	}
	if utf8.RuneCountInString(u.FirstName) > USER_FIRST_NAME_MAX_RUNES {
		return NewInvalidUserError("first_name", u.ID)
	}
	if utf8.RuneCountInString(u.LastName) > USER_LAST_NAME_MAX_RUNES {
		return NewInvalidUserError("last_name", u.ID)
	}
	if len(u.Password) > USER_PASSWORD_MAX_LENGTH {
		return NewInvalidUserError("password_limit", u.ID)
	}
	if !IsValidLocale(u.Locale) {
		return NewInvalidUserError("locale", u.ID)
	}
	return nil
}

// Sanitize removes any private data from the user object
func (u *User) Sanitize(options map[string]bool) {
	u.Password = ""
	if len(options) != 0 && !options["email"] {
		u.Email = ""
	}
}

// PreSave will set missing defaults and fill CreatedAt and UpdatedAt times
// It will also hash the password and it should be called before saving the user to the db
func (u *User) PreSave() {
	u.Username = NormalizeUsername(u.Username)
	u.Email = NormalizeEmail(u.Email)
	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt

	if u.Locale == "" {
		u.Locale = USER_DEFAULT_LOCALE
	}
	if len(u.Password) > 0 {
		u.Password = HashPassword(u.Password)
	}
}

// PreUpdate should be called before updating the user in the db
func (u *User) PreUpdate() {
	u.Username = NormalizeUsername(u.Username)
	u.Email = NormalizeEmail(u.Email)
	u.UpdatedAt = time.Now()
}

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
