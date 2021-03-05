package const_errs

// Error type
type Error string

// Warning type
type Warning string

// Error returns error description.
func (err Error) Error() string {
	return string(err)
}

// Error returns error description.
func (err Warning) Error() string {
	return string(err)
}

// NewError creates a new error of Error type.
func NewError(text string) Error {
	return Error(text)
}

// NewWarning creates a new error of Warning type.
func NewWarning(text string) Warning {
	return Warning(text)
}
