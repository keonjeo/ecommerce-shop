package model

import (
	"encoding/json"
	"strings"
)

// AppError represents the app error
type AppError struct {
	ID            string `json:"id"`                       // the error id (in form of -> package.file_name.error_info.app_error)
	Where         string `json:"-"`                        // where the error happened (in form of -> struct.func)
	Name          string `json:"name"`                     // the error name
	Message       string `json:"message"`                  // the message to display to the end user
	RequestID     string `json:"request_id,omitempty"`     // requestID that's set in the headers
	DetailedError string `json:"detailed_error,omitempty"` // error details to help the developer
	StatusCode    int    `json:"status_code"`              // the http status code

	Params map[string]interface{} `json:"-"`
}

// NewAppError creates the app error
func NewAppError(where string, id string, params map[string]interface{}, details string, status int) *AppError {
	e := &AppError{}
	e.ID = id
	e.Where = where
	e.Name = "Err" + strings.Title(where)
	e.Message = id
	e.DetailedError = details
	e.StatusCode = status
	e.Params = params
	return e
}

func (e *AppError) Error() string {
	return e.Where + ": " + e.Message + ", " + e.DetailedError
}

// ToJSON converts AppError to json string
func (e *AppError) ToJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}
