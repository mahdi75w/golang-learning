package main

import (
	"errors"
	"fmt"
)

func ErrorSampleMain() {
	SampleRaiseError()

	err := ValidateUser()
	fmt.Println(err.Error())
	validationErr := new(UserValidation)
	if errors.As(err, validationErr) {
		fmt.Println(validationErr.Code)
		fmt.Println(validationErr.Email)
		fmt.Println(validationErr.Message)
	} else {
		fmt.Println("Its not a userValidation Error")
	}
	fmt.Println("----------------------------------")
	recoverPanic()
	fmt.Println("print after panic")
	fmt.Println("----------------------------------")
	regErr := WrapError()
	fmt.Println(regErr)

	valErr := errors.Unwrap(regErr)
	fmt.Println(valErr)

	fmt.Println("----------------------------------")
}

func checkPassword(password string) error {
	if password != "somethings" {
		return errors.New("password mismatch")
	}
	return nil
}

func SampleRaiseError() {
	err := checkPassword("password")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("password ok")
	}
	fmt.Println("----------------------------------")
}

type UserValidation struct { // return custom type as error
	Code    int
	Email   string
	Message string
}

func (r UserValidation) Error() string { // need attach this method to custom error
	return r.Message
}

func ValidateUser() error {
	return UserValidation{
		Code:    1,
		Email:   "a@b.com",
		Message: "Email is valid",
	}

}

func recoverPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in recoverPanic:", r)
		}
	}()
	panic("This should never happen")
}

func ValidateUserInner() error {
	return errors.New("invalid user")
}

func WrapError() error {
	err := ValidateUserInner()
	if err != nil {
		return fmt.Errorf("inner text in wrap: %w", err)
	}
	return nil
}
