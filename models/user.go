package models

import "gorm.io/gorm"

type Role int

const (
	RoleAdmin  Role = 1
	RoleMember Role = 0
)

type User struct {
	gorm.Model

	Login string `json:"login"`

	PasswordHash string `json:"password_hash"`
	PasswordSalt string `json:"password_salt"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Email string `json:"email"`

	Role Role `json:"role"`

	Checkouts []Checkout `json:"checkouts"`
}
