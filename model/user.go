package model

import (
    "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	ID            string          `json:"id" bson:"_id"`
	First         string          `json:"first"`
	Last          string          `json:"last"`
	Email         string          `json:"email"`
	Notifications []*Notification `json:"notifications"`
}

func (createUser *User) CreateUserValidate() error {
    return validation.Validate(&createUser.Email,
        is.Email.Error("正しいメールアドレスの形で入力してください"),
    )
}