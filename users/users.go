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
	// variables can be defined with tea var keyword
	var newUser = User{
		Name:  name,
		Email: email,
	}
	// mmm this seems too easy, maybe we should validate the input before
	// returning the new user. There should be a function somewhere...
	return newUser, nil
}

// validateEmail check if the email is valid.
// This function is very simple, do you think you can improve it?
func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("email must contain '@' symbol")
	}
	return nil
}

// String implements the Stringer interface for the User type.
func (u User) String() string {
	return fmt.Sprintf("Name: %s, Email: %s", u.Name, u.Email)
}
