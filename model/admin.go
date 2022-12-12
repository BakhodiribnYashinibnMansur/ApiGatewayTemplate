package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

// SignUpModel ...
type UserSignIn struct {
	Username string `json:"username" validate:"required" default:"username"`
	Password string `json:"password" validate:"required" default:"password"`
}

// Validate Register Model
func (um *UserSignIn) Validate() error {
	return validation.ValidateStruct(
		um,
		validation.Field(&um.Password, validation.Required, validation.Length(8, 30), validation.Match(regexp.MustCompile("[a-z]|[A-Z][0-9]"))),
		validation.Field(&um.Username, validation.Required, validation.Length(5, 30)),
	)
}
