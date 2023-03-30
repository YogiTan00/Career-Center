package utils

import (
	"regexp"
)

func IsNumber(s string) bool {
	re := regexp.MustCompile("^[0-9]+$")
	return re.MatchString(s)
}
