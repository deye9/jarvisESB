package entity

import (
	"net/mail"
	"time"
)

// Email struct
type Email struct {
	From     mail.Address
	To       mail.Address
	Subject  string
	Body     string
	Template string
}

// ResponseMessage struct
type ResponseMessage struct {
	Success   string                 `json:"success,omitempty"`
	Message   string                 `json:"message,omitempty"`
	ErrorCode string                 `json:"error_code,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// User struct
type User struct {
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
