package exceptions

import (
	"errors"
)

var (
	ErrorWrongEmailorPassword = errors.New("wrong email or password")
)

func ErrCustomString(s string) error {
	return errors.New(s)
}
