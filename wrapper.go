package const_errs

import "fmt"

// ErrorWrapper type
type ErrorWrapper struct {
	error
	string
}

// Error returns error description.
func (err ErrorWrapper) Error() string {
	if err.error != nil {
		return fmt.Sprintf("%s: %v", err.string, err.error)
	}
	return err.string
}

// Wrap wraps the current error by another error.
func (err ErrorWrapper) Wrap(msg string) ErrorWrapper {
	return ErrorWrapper{
		error:  err,
		string: msg,
	}
}