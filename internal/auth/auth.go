package auth

import "fmt"

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterLoginPayload struct {
	LoginPayload
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type ConfirmEmailPayload struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

var (
	ErrEmailIsRequired                          = fmt.Errorf("email is required")
	ErrPasswordIsRequired                       = fmt.Errorf("password is required")
	ErrPasswordConfirmationIsRequired           = fmt.Errorf("password confirmation is required")
	ErrPasswordAndPasswordConfirmationMustMatch = fmt.Errorf("password and password confirmation must match")
)

func MustHaveEmail(input LoginPayload) error {
	if input.Email == "" {
		return ErrEmailIsRequired
	}

	return nil
}

func MustHaveEmailAndPassword(input LoginPayload) error {
	if input.Email == "" {
		return ErrEmailIsRequired
	}

	if input.Password == "" {
		return ErrPasswordIsRequired
	}

	return nil
}

func MustHaveEmailAndMatchingPasswords(input RegisterLoginPayload) error {
	if input.Email == "" {
		return ErrEmailIsRequired
	}

	if input.Password == "" {
		return ErrPasswordIsRequired
	}

	if input.PasswordConfirmation == "" {
		return ErrPasswordConfirmationIsRequired
	}

	if input.Password != input.PasswordConfirmation {
		return ErrPasswordAndPasswordConfirmationMustMatch
	}

	return nil
}
