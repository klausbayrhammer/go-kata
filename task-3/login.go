package user

// Loginer is an interface with a login method.
// Yes the name sucks, but it is idiomatic to call interfaces with a -er suffix.
// Can yo find a better name?
type Loginer interface {
	// Login allows the type implementing this method to login by providing
	// an opaque token.
	// In case of success a nil error and a session token is returned.
	// In case of error the session in not to valid.
	Login(token string) (session string, err error)
}

// TODO how can we make our user implement this interface?
// And what if we have different login mechanisms?
