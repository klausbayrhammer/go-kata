package user

import (
	"fmt"
	address "go-kata/task-2"
	"testing"
)

// TestUserLogin is a function that tests the user logins.
func TestUserLogin(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		// variables can be declared and defined with the := operator
		// the type is inferred
		name, email := "Bob", "bob's email.com"
		// we don't need the new user for this test so we can assign it to
		// the blank identifier https://go.dev/doc/effective_go#blank
		_, err := New(name, email, address.Address{})
		if err != nil {
			t.Logf("user should be valid with name: %s and email: %s", name, email)
			t.Fail()
		}
	})
}

// testLogin implements the Loginer interface  for tests.
// 💡can you create other implementations?
type testLogin struct {
	hash string
}

func (tl testLogin) Login(token string) (string, error) {
	if token == tl.hash {
		return "session", nil
	}
	return "", fmt.Errorf("login error, check credentials")
}
