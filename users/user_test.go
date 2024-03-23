package users

import (
	"testing"
)

// TestUsers is a function that tests the users package.
func TestUsers(t *testing.T) {
	// this is a subtest, it allows us to group tests together
	t.Run("New", func(t *testing.T) {
		// variables can be defined also with the := operator
		name, email := "Bob", "bob's email.com"
		// we don't need the new user for this test so we can assign it to
		// the blank identifier https://go.dev/doc/effective_go#blank
		_, err := New(name, email)
		if err != nil {
			t.Errorf("user should be valid with name: %s and email: %s", name, email)
		}
	})
	t.Run("validateEmail", func(t *testing.T) {
		err := validateEmail("bob's email.com")
		if err != nil {
			t.Error(err)
		}
	})
	// maybe we should have some test for the not so happy path
}
