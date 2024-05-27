package models

import "github.com/go-playground/validator/v10"

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age" `
	Gmail    string `json:"gmail" validate:"required, containsany=@"`
	Address  string `json:"address"`
	Password string `json:"password" validate:"required, containsany=!@#?$"`
}

func (a *User) Validate() error {
	var validate = validator.New()
	return validate.Struct(a)
}

type LoginReply struct {
	RefreshToken string
	AccessToken  string
}
