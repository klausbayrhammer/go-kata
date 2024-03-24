package users

import (
	"fmt"
	"strings"
)

type User struct {
	Name  string
	Email string
}

// New returns a new User.
func New(name, email string) (User, error) {
	err := validateEmail(email)
	if err != nil {
		return User{}, fmt.Errorf("email validation failed: %w", err)
	}
	newUser := User{
		Name:  name,
		Email: email,
	}
	return newUser, nil
}

// String implements the Stringer interface for the User type.
// This is how a Go type implements methods
func (u User) String() string {
	return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}

// validateEmail checks if the email is valid.
// We all know that the true way to validate an email address is to send a
// confirmation email, but this will do for now.
func validateEmail(email string) error {
	if strings.Contains(email, "@") == false {
		return fmt.Errorf("email must contain '@' symbol")
	}
	return nil
}
