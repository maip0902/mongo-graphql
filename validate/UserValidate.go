package validate

import (
    "fmt"

    "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
    "github.com/maip0902/mongo-graphql/model"
)

func (createUser *model.NewUser) CreateUserValidate() err {
    return validation.ValidateStruct(&createUser,
    validation.Field(&createUser.Email, is.Email)
    )
}


