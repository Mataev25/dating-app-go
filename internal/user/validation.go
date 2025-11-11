package user

import (
	"errors"
	"unicode/utf8"
)

var (
	ErrNameShort = errors.New(
		"name must be at least 2 characters long",
	)
	ErrNameLong = errors.New(
		"name must be no more than 50 characters long",
	)
	ErrInvalidAge = errors.New(
		"age must be between 18 and 100",
	)
	ErrInvalidGender = errors.New("invalid gender")
	ErrInvalidLookingFor = errors.New("invalid looking_for value")
	ErrDescriptionLong = errors.New(
		"description must be no more" +
		"500 characters long",
	)

	ErrEmailRequired = errors.New("email is required")
)

func (u *User) Validate() error {
	nameLength := utf8.RuneCountInString(u.Name)
	if nameLength < 2 {
		return ErrNameShort
	}
	if nameLength > 50 {
		return ErrNameLong
	}
	if u.Age < 18 || u.Age > 100 {
		return ErrInvalidAge
	}
	if u.Gender != GenderMale && u.Gender != GenderFemale {
		return ErrInvalidGender
	}
	switch u.LookingFor {
	case LookingMale, LookingFemale, LookingBoth:
	default:
		return ErrInvalidLookingFor
	}
	if utf8.RuneCountInString(u.Description) > 500 {
		return ErrDescriptionLong
	}
	if u.Email == "" {
		return ErrEmailRequired
	}

	return nil
}

