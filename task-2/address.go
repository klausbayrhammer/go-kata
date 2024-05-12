// The address package defines the address type and some validation functions
// In Go the symbols starting with a capital letter are exported, visible from
// outside the package scope. In other languages the equivalent is to use a
// keyword like `public` or `export`.

package address

import "fmt"

// Address is user defined struct, in Go there are no classes but any user
// defined type can have methods
type Address struct {
	Country string
	City    string
	ZipCode string
	Street  string
	Name    string
}

// Validate returns an error in case the address is not valid.
// This is an example of a method for a type in Go.
//
// There are not enough validations! Can you improve it?
func (addr Address) Validate() error {
	// TODO implement the validation, some is already started can you finish it?
	err := isValidCountry(addr.Country)
	if err != nil {
		return fmt.Errorf("invalid country %q: %w", addr.Country, err)
	}
	// err = isValidZipCode()

	return nil
}

// isValidCountry is not exported and not visible from outside the address package.
func isValidCountry(country string) error {
	// countries don't change very often, maybe a simple way to validate this
	// field is to have a list of countries and check if it is present?
	// let's have a for loop "range" over a slice of strings ([]string)
	// maybe there is some explanation here https://go.dev/doc/effective_go#for

	return nil
}

// countryList is a list of some country for this Kata also not exported
var countryList = []string{
	"Austria",
	"Italy",
	"Switzerland",
	"Germany",
	"Hungary",
	"Slovenia",
	"Slovakia",
	"Czech Republic",
	"Croatia",
}

// isValidZipCode return an error if the zipcode is not valid for the city
func isValidZipCode(city, zipcode string) error {
	// in this case a map (dictionary) can be very useful to see all the
	// possible zipcodes in a city
	return nil
}

var cityZipcode = map[string][]string{
	"Austria": {"123", "456"},
	"Italy":   {"123", "456"},
}

// 💡 all these variable are hardcoded, would it be better to load the validation
// data from a file? Maybe a JSON?
