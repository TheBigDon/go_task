package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// User ...
type User struct {
	ID	   int    `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u, validation.Field(&u.Mobile, validation.Required, is.E164),
validation.Field(&u.Name, validation.Required, is.UTFLetter)) 
}

func (u *User) Sanitize() {
	u.Mobile = ""
}