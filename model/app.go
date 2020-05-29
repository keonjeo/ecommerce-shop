package model

// package model

// import (
// 	"encoding/json"
// 	"strings"

// 	"github.com/dankobgd/ecommerce-shop/utils/locale"
// 	"github.com/nicksnyder/go-i18n/v2/i18n"
// )

// // AppErr represents the app error
// type AppErr struct {
// 	ID         string `json:"id"`                // the error id (in form of -> package.file_name.error_info.app_error)
// 	Op         string `json:"-"`                 // Operation where the error happened (in form of -> struct.func)
// 	Name       string `json:"name"`              // the error name
// 	Message    string `json:"message"`           // the message to display to the end user
// 	Details    string `json:"details,omitempty"` // error details to help the developer
// 	StatusCode int    `json:"status_code"`       // the http status code

// 	Msg    *i18n.Message `json:"-"`
// 	params map[string]interface{}
// }

// // NewAppErr creates the app error
// func NewAppErr(op string, l *i18n.Localizer, msg *i18n.Message, details string, status int) *AppErr {
// 	e := &AppErr{}
// 	e.ID = msg.ID
// 	e.Op = op
// 	e.Name = "Err" + strings.Title(op)
// 	e.Details = details
// 	e.StatusCode = status

// 	e.Msg = msg
// 	e.Localize(l)
// 	return e
// }

// func (e *AppErr) Error() string {
// 	return e.Op + ": " + e.Message + ", " + e.Details
// }

// // Localize translates the error message
// func (e *AppErr) Localize(l *i18n.Localizer) {
// 	msg := locale.LocalizeDefaultMessage(l, e.Msg)
// 	e.Message = msg
// }

// // ToJSON converts AppErr to json string
// func (e *AppErr) ToJSON() string {
// 	b, _ := json.Marshal(e)
// 	return string(b)
// }
