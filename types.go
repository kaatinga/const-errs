package const_errs

import "fmt"

// ErrorWrapper type
type ErrorWrapper struct {
	error
	string
}

// Error type
type Error string

// Warning type
type Warning string

// Debug type
type Debug string

// Error returns error description.
func (err ErrorWrapper) Error() string {
	if err.error != nil {
		return fmt.Sprintf("%s: %v", err.string, err.error)
	}
	return err.string
}

// Wrap wraps the current error by another error.
func (err Error) Wrap(msg string) ErrorWrapper {
	return ErrorWrapper{
		error:  err,
		string: msg,
	}
}

// Error returns error description.
func (err Error) Error() string {
	return string(err)
}

// Error returns error description.
func (err Warning) Error() string {
	return string(err)
}

// Error returns error description.
func (err Debug) Error() string {
	return string(err)
}

// NewError creates a new error of Error type that wraps an error optionally.
func NewError(text string, err error) Error {
	return Error(text)
}

// NewWarning creates a new error of Warning type.
func NewWarning(text string) Warning {
	return Warning(text)
}

// NewDebug creates a new error of Debug type.
func NewDebug(text string) Debug {
	return Debug(text)
}
