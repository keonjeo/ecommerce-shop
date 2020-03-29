package pwutil

import (
	"errors"
	"strings"

	"github.com/dankobgd/ecommerce-shop/config"
)

const (
	NUMBERS           = "0123456789"
	SYMBOLS           = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"
	LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// IsValidPasswordCriteria checks if password fulfills the criteria
func IsValidPasswordCriteria(password string, settings *config.PasswordSettings) error {
	var err error

	if len(password) < settings.MinLength || len(password) > settings.MaxLength {
		err = errors.New("invalid password length")
	}

	if settings.Lowercase {
		if !strings.ContainsAny(password, LOWERCASE_LETTERS) {
			err = errors.New("must have a lowercase letter")
		}
	}

	if settings.Uppercase {
		if !strings.ContainsAny(password, UPPERCASE_LETTERS) {
			err = errors.New("must have an uppercase letter")
		}
	}

	if settings.Number {
		if !strings.ContainsAny(password, NUMBERS) {
			err = errors.New("must have uppercase letter")
		}
	}

	if settings.Symbol {
		if !strings.ContainsAny(password, SYMBOLS) {
			err = errors.New("must have a symbol")
		}
	}

	return err
}
