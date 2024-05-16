package check

import (
	"errors"
	"regexp"
)

// func ValidatePhone(phone string) error {
// 	phoneRegex := regexp.MustCompile(`^\+?998\d{9}$`)
// 	if phoneRegex.MatchString(phone) {
// 		return errors.New("phone is not valid")
// 	}
// 	return nil
// }

// "+998123456789",// Kорректный номер
// "998123456789",// Kорректный номер
// "+997123456789", // Некорректный номер
// "99812345678", // Некорректный номер

func ValidatePhone(phone string) error {
	phoneRegex := regexp.MustCompile(`^\+?998\d{9}$`)
	if !phoneRegex.MatchString(phone) {
		return errors.New("phone is not valid")
	}
	return nil
}

func ValidatePassword(pasword string) error {
	paswordRegex := regexp.MustCompile(`^(.*[A-Za-z])(.*\d).{8,}$`)
	if paswordRegex.MatchString(pasword) {
		return errors.New("password is not valid")
	}
	return nil
}
