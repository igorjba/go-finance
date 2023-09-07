package validator

import (
	"regexp"
)

func IsEmailValid(email string) bool {
	regex := `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	return regexp.MustCompile(regex).MatchString(email)
}
