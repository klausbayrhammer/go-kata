// packages in Go are namespaces for all the symbols defined inside the package
package main

// importing a package
import "fmt"

// function main() in the main package is the entry point of a Go executable
func main() {
	// exported symbols start with an uppercase letter, in most language the
	// same effect is achieved with a keyword, e.g. public, export, etc..
	fmt.Println("Hello, World!")
}
