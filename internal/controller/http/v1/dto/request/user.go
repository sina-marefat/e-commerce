package request

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidPhoneNumber = errors.New("Invalid phone number.")
)

type CreateUser struct {
	Email       string `json:"email" validate:"email,required"`
	Password    string `json:"password" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	CityID      int    `json:"city_id"`
}

func (cu CreateUser) validate() error {
	if match, _ := regexp.MatchString("09(1[0-9]|3[1-9]|2[1-9])-?[0-9]{3}-?[0-9]{4}", cu.PhoneNumber); !match {
		return ErrInvalidPhoneNumber
	}
	return nil
}

type UpdateUser struct {
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	CityID      int    `json:"city_id"`
}

func (cu UpdateUser) validate() error {
	if match, _ := regexp.MatchString("09(1[0-9]|3[1-9]|2[1-9])-?[0-9]{3}-?[0-9]{4}", cu.PhoneNumber); !match {
		return ErrInvalidPhoneNumber
	}
	return nil
}

type Login struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}
