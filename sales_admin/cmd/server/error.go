package main

// exitCode is a type alias to make things a little more explicit
type exitCode = int

const (
	// exitUnknown is the error code that is used if the error code can't be pulled
	// from the error. 200 is just an arbitrarily large number.
	exitUnknown exitCode = 200
	badFlags    exitCode = iota + 1
	failedSetup
	failedServerSetup
	serverError
	shutdownError
)

type cmdError struct {
	code exitCode
	err  error
}

// Error implements the error interface. It returns the underlying error's error
// string.
func (s cmdError) Error() string {
	return s.err.Error()
}
