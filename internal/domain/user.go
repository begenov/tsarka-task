package domain

import "errors"

var (
	NotFound = errors.New("not found")
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}
