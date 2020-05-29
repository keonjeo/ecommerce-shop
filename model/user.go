package model

import (
	"encoding/json"
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
	"github.com/dankobgd/ecommerce-shop/utils/locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const (
	bcryptCost = 14

	numbers          = "0123456789"
	symbols          = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	userEmailMaxLength    = 128
	userPasswordMaxLength = 72
	userUsernameMaxRunes  = 64
	userFirstnameMaxRunes = 64
	userLastnameMaxRunes  = 64
	userUsernameMaxLength = 64
	userUsernameMinLength = 1
	userLocaleMaxLength   = 5
	userDefaultLocale     = "en"
)

var reservedNames = []string{"app", "api", "admin", "signup", "login", "oauth", "error", "help"}
var restrictedUsernames = []string{"app", "api", "admin", "system"}
var validUsernameChars = regexp.MustCompile(`^[a-z0-9\.\-_]+$`)

var (
	msgValidateID                = &i18n.Message{ID: "model.user.validate.id.app_error", Other: "uppercase letter required"}
	msgValidateCreatedAt         = &i18n.Message{ID: "model.user.validate.created_at.app_error", Other: "invalid created_at timestamp"}
	msgValidateUpdatedAt         = &i18n.Message{ID: "model.user.validate.updated_at.app_error", Other: "invalid updated_at timestamp"}
	msgValidateUsername          = &i18n.Message{ID: "model.user.validate.username.app_error", Other: "invalid username"}
	msgValidateEmail             = &i18n.Message{ID: "model.user.validate.email.app_error", Other: "invalid email"}
	msgValidateFirstName         = &i18n.Message{ID: "model.user.validate.first_name.app_error", Other: "invalid first name"}
	msgValidateLastName          = &i18n.Message{ID: "model.user.validate.last_name.app_error", Other: "invalid last name"}
	msgValidatePassword          = &i18n.Message{ID: "model.user.validate.password.app_error", Other: "invalid password length"}
	msgValidateLocale            = &i18n.Message{ID: "model.user.validate.locale.app_error", Other: "invalid locale"}
	msgValidatePasswordUppercase = &i18n.Message{ID: "model.user.validate.password_uppercase", Other: "uppercase letter required"}
	msgValidatePasswordLowercase = &i18n.Message{ID: "model.user.validate.password_lowercase", Other: "lowercase letter required"}
	msgValidatePasswordNumbers   = &i18n.Message{ID: "model.user.validate.password_numbers", Other: "number required"}
	msgValidatePasswordSymbols   = &i18n.Message{ID: "model.user.validate.password_symbols", Other: "symbol required"}
)

// User represents the shop user model
type User struct {
	ID             int       `json:"id" db:"id"`
	FirstName      string    `json:"first_name" db:"first_name"`
	LastName       string    `json:"last_name" db:"last_name"`
	Username       string    `json:"username" db:"username"`
	Email          string    `json:"email" db:"email"`
	Password       string    `json:"password,omitempty" db:"password"`
	Gender         *string   `json:"gender" db:"gender"`
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

// IsValidUsername validates whether username matches the criteria
func IsValidUsername(username string) bool {
	if len(username) < userUsernameMinLength || len(username) > userUsernameMaxLength {
		return false
	}
	if !validUsernameChars.MatchString(username) {
		return false
	}
	for _, ru := range restrictedUsernames {
		if username == ru {
			return false
		}
	}
	return true
}

// IsValidLocale checks if locale is valid
func IsValidLocale(locale string) bool {
	if locale != "" {
		if len(locale) > userLocaleMaxLength {
			return false
		} else if _, err := language.Parse(locale); err != nil {
			return false
		}
	}
	return true
}

// IsValidPasswordCriteria checks if password fulfills the criteria
func IsValidPasswordCriteria(password string, settings *config.PasswordSettings) *AppErr {
	if len(password) < settings.MinLength || len(password) > settings.MaxLength {
		return NewInvalidUserError(msgValidatePasswordLowercase, 0)
	}
	if settings.Lowercase {
		if !strings.ContainsAny(password, lowercaseLetters) {
			return NewInvalidUserError(msgValidatePassword, 0)
		}
	}
	if settings.Uppercase {
		if !strings.ContainsAny(password, uppercaseLetters) {
			return NewInvalidUserError(msgValidatePasswordUppercase, 0)
		}
	}
	if settings.Number {
		if !strings.ContainsAny(password, numbers) {
			return NewInvalidUserError(msgValidatePasswordNumbers, 0)
		}
	}
	if settings.Symbol {
		if !strings.ContainsAny(password, symbols) {
			return NewInvalidUserError(msgValidatePasswordSymbols, 0)
		}
	}
	return nil
}

// NewInvalidUserError builds the invalid user error
func NewInvalidUserError(msg *i18n.Message, userID int) *AppErr {
	details := map[string]interface{}{}
	if userID != 0 {
		details["userID"] = "userID"
	}
	return NewAppErr("User.Validate", ErrInvalid, locale.GetUserLocalizer("en"), msg, http.StatusUnprocessableEntity, details)
}

// Validate validates the user and returns an error if it doesn't pass criteria
func (u *User) Validate() *AppErr {
	if u.ID < 0 {
		return NewInvalidUserError(msgValidateID, u.ID)
	}
	if u.CreatedAt.IsZero() {
		return NewInvalidUserError(msgValidateCreatedAt, u.ID)
	}
	if u.UpdatedAt.IsZero() {
		return NewInvalidUserError(msgValidateUpdatedAt, u.ID)
	}
	if !IsValidUsername(u.Username) {
		return NewInvalidUserError(msgValidateUsername, u.ID)
	}
	if len(u.Email) > userEmailMaxLength || len(u.Email) == 0 || !is.ValidEmail(u.Email) {
		return NewInvalidUserError(msgValidateEmail, u.ID)
	}
	if utf8.RuneCountInString(u.Username) > userUsernameMaxRunes {
		return NewInvalidUserError(msgValidateUsername, u.ID)
	}
	if utf8.RuneCountInString(u.FirstName) > userFirstnameMaxRunes {
		return NewInvalidUserError(msgValidateFirstName, u.ID)
	}
	if utf8.RuneCountInString(u.LastName) > userLastnameMaxRunes {
		return NewInvalidUserError(msgValidateLastName, u.ID)
	}
	if len(u.Password) > userPasswordMaxLength {
		return NewInvalidUserError(msgValidatePassword, u.ID)
	}
	if !IsValidLocale(u.Locale) {
		return NewInvalidUserError(msgValidateLocale, u.ID)
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
		u.Locale = userDefaultLocale
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
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
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
