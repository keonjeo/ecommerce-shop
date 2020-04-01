package model

import "encoding/json"

// AppError represents the app error
type AppError struct {
	ID            string `json:"id"`
	Message       string `json:"message"`
	Where         string `json:"where"`
	RequestID     string `json:"request_id,omitempty"`
	InternalError string `json:"internal_error,omitempty"`
	DetailedError string `json:"detailed_error,omitempty"`
	StatusCode    int    `json:"status_code"`

	Params map[string]interface{} `json:"-"`
}

// NewAppError creates the app error
func NewAppError(where string, id string, message string, details string, params map[string]interface{}, status int) *AppError {
	e := &AppError{}
	e.ID = id
	e.Where = where
	e.Message = message
	e.DetailedError = details
	e.StatusCode = status
	e.Params = params
	return e
}

func (e *AppError) Error() string {
	return e.Where + ": " + e.ID + ", " + e.DetailedError
}

// ToJSON converts AppError to json string
func (e *AppError) ToJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}
