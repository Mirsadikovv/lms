package check

import (
	"errors"
	"time"
)

func ValidateYear(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}

	return nil
}
