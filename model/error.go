package model

import (
	"encoding/json"
	"fmt"

	"github.com/dankobgd/ecommerce-shop/utils/locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Application error codes
const (
	ErrConflict        = "Conflict"         // action cannot be performed
	ErrInternal        = "Internal"         // internal error
	ErrInvalid         = "Invalid"          // validation failed
	ErrNotFound        = "Not Found"        // entity does not exist
	ErrPermision       = "Permision Denied" // permission denied
	ErrUnauthenticated = "Unauthenticated"  // unauthenticated access
)

// AppErr is the main app error
type AppErr struct {
	ID         string `json:"id"`            // unique string which is the same as translation id
	Op         string `json:"op"`            // operation where it failed (Struct.Func)
	Code       string `json:"code"`          // machine readable error code
	StatusCode int    `json:"status_code"`   // http status code
	Message    string `json:"message"`       // meaningful end user message
	Err        error  `json:"err,omitempty"` // embeded error

	Msg     *i18n.Message          `json:"-"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// NewAppErr creates the new app error
func NewAppErr(op string, code string, l *i18n.Localizer, msg *i18n.Message, statusCode int, details map[string]interface{}) *AppErr {
	e := &AppErr{
		ID:         msg.ID,
		Op:         op,
		Code:       code,
		StatusCode: statusCode,
		Details:    details,
		Msg:        msg,
	}

	e.Localize(l)
	return e
}

func (e *AppErr) Error() string {
	return fmt.Sprintf("%v, %v: %v\n", e.Code, e.Op, e.Message)
}

// Localize translates the error message
func (e *AppErr) Localize(l *i18n.Localizer) {
	msg := locale.LocalizeDefaultMessage(l, e.Msg)
	e.Message = msg
}

// ToJSON converts AppErr to json string
func (e *AppErr) ToJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// ErrorCode returns the code of the root error, if available. Otherwise returns ErrInternal
func ErrorCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*AppErr); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return ErrInternal
}

// ErrorMessage returns the human-readable message of the error, if available
// Otherwise returns a generic error message
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*AppErr); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "an internal error has occurred, please contact support"
}
