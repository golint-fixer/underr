// Package underr provides access to underlying errors when wrapping errors.
//
// Often errors are "wrapped" in Go to augment them with additional
// information. This package provides for some generic utilities to work with
// such errors assuming they satisfy the interface defined here, which is that
// they provide a method called `Underlying` which returns the error being
// wrapped.
package underr

// Error is an error that has been wrapped, and Underlying returns the error it
// wraps.
type Error interface {
	error

	Underlying() error
}

// All returns all the underlying errors by iteratively checking if the
// error has an Underlying error. If e is nil, the returned slice will be nil.
func All(e error) []error {
	if e == nil {
		return nil
	}

	var errs []error
	for {
		if e == nil {
			return errs
		}
		errs = append(errs, e)

		if eh, ok := e.(Error); ok {
			e = eh.Underlying()
		} else {
			e = nil
		}
	}
}

// Matcher defines the function to be used as a predicate for checking for a
// kind of error.
type Matcher func(error) bool

// Equal returns a Matcher to check if an error equals the given error.
func Equal(e error) Matcher {
	return func(o error) bool { return e == o }
}

// Has returns true if any of the underlying errors satisfy the given Matcher.
func Has(e error, f Matcher) bool {
	for _, o := range All(e) {
		if f(o) {
			return true
		}
	}
	return false
}
