package model

import (
	"github.com/go-playground/validator"
)

type InputData struct {
	Created_at string `json:"-"`
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"email,required"`
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" vlidate:"required"`
}

func (inputData *InputData) Validate() error {
	validate := validator.New()
	return validate.Struct(inputData)
}
