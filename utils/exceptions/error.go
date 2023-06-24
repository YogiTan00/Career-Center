package exceptions

import (
	"errors"
)

var (
	Unauthorized              = errors.New("user role not admin")
	ErrorWrongEmailorPassword = errors.New("wrong email or password")
	ErrorStartDate            = errors.New("need start date")
	ErrorEndDate              = errors.New("end date cannot be smaller than start date")
)

func ErrCustomString(s string) error {
	return errors.New(s)
}
