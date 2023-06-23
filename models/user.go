package models

import (
	"time"
)

type User struct {
	Id        int        `json:"id"`
	Name      *string    `json:"name,omitempty" form:"name"`
	Email     *string    `json:"email,omitempty" form:"email"`
	Age       *int64     `json:"age,omitempty" form:"age"`
	Height    *float32   `json:"height,omitempty" form:"height"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}