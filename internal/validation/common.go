package validation

import (
	"fmt"
	"unicode/utf8"
)

func ValidateMaxLength(s string, fieldName string, max int) (bool, error) {
	if l := utf8.RuneCountInString(s); l > max {
		return false, fmt.Errorf("%s must not exeed %d characters", fieldName, max)
	}
	return true, nil
}

func ValidateEmpty(s string, fieldName string) (bool, error) {
	if s == "" {
		return false, fmt.Errorf("%s cannot be empty", fieldName)
	}
	return true, nil
}
