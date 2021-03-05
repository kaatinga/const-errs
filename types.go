package const_errs

import "fmt"

// Error type
type Error struct {
	error
	string
}

// Warning type
type Warning string

// Debug type
type Debug string

// Error returns error description.
func (err Error) Error() string {
	if err.error != nil {
		return fmt.Sprintf("%s: %v", err.string, err)
	}
	return err.string
}

// WrapByError wraps the current error by another error.
func (err Error) WrapByError(errWrapper Error) {
	errWrapper.error = err
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
	return Error{err, text}
}

// NewWarning creates a new error of Warning type.
func NewWarning(text string) Warning {
	return Warning(text)
}

// NewDebug creates a new error of Debug type.
func NewDebug(text string) Debug {
	return Debug(text)
}
