package model

import (
	"encoding/json"
	"strings"

	"github.com/dankobgd/ecommerce-shop/utils/locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// NewBool returns pointer to bool
func NewBool(b bool) *bool { return &b }

// NewInt returns pointer to int
func NewInt(n int) *int { return &n }

// NewInt64 returns pointer to int64
func NewInt64(n int64) *int64 { return &n }

// NewString returns pointer to string
func NewString(s string) *string { return &s }

// AppError represents the app error
type AppError struct {
	ID         string `json:"id"`                // the error id (in form of -> package.file_name.error_info.app_error)
	Op         string `json:"-"`                 // Operation where the error happened (in form of -> struct.func)
	Name       string `json:"name"`              // the error name
	Message    string `json:"message"`           // the message to display to the end user
	Details    string `json:"details,omitempty"` // error details to help the developer
	StatusCode int    `json:"status_code"`       // the http status code

	Msg    *i18n.Message `json:"-"`
	params map[string]interface{}
}

// NewAppError creates the app error
func NewAppError(op string, l *i18n.Localizer, msg *i18n.Message, details string, status int) *AppError {
	e := &AppError{}
	e.ID = msg.ID
	e.Op = op
	e.Name = "Err" + strings.Title(op)
	e.Details = details
	e.StatusCode = status

	e.Msg = msg
	e.Localize(l)
	return e
}

func (e *AppError) Error() string {
	return e.Op + ": " + e.Message + ", " + e.Details
}

// Localize translates the error message
func (e *AppError) Localize(l *i18n.Localizer) {
	msg := locale.LocalizeDefaultMessage(l, e.Msg)
	e.Message = msg
}

// ToJSON converts AppError to json string
func (e *AppError) ToJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}
