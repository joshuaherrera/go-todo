package models

import (
	"errors"
)

// ErrNoRecords returns an error if no model found
var ErrNoRecords = errors.New("models: no matching records found")

// Todo describes the todo record
type Todo struct{
	TodoID int `json:"todo_id"`
	Description string `json:"description"`
}