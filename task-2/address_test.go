package address

import "testing"

// TestAddressIsValid tests the address validation function IsValid()
// The Go standard library includes a testing framework in the testing package.
//
// This test is quite incomplete, can you improve it?
func TestAddressIsValid(t *testing.T) {

	// this is a struct literal
	addr := Address{
		Country: "austria",
		City:    "salzburg",
		ZipCode: "5050",
		Street:  "",
		Name:    "",
	}

	// This is a sub-test
	t.Run("isValidCountry", func(t *testing.T) {
		// add test here
		err := isValidCountry(addr.Country)
		if err != nil {
			// in case of error we print it and fail the test
			t.Error(err)
		}
	})

	t.Run("isValidZipCode", func(t *testing.T) {
		// now that I have your attention with this failing test, can you improve
		// this package tests?
		t.Error("this test needs to improve")
	})

	// TODO more tests
}
