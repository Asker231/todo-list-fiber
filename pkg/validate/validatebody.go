package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)


func ValidateBody[T any](data T)error{
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil{
		fmt.Println(err.Error())
		return err
	}
	return err
}