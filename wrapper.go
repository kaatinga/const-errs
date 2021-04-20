package const_errs

// AnnotatedError type.
type AnnotatedError struct {
	cause error
	string
}

// Error returns error description.
func (err AnnotatedError) Error() string {
	if err.cause != nil {
		return err.cause.Error() + ": " + err.string
	}
	return err.string
}

// Annotate adds an additional description to go with an error.
func (err AnnotatedError) Annotate(msg string) AnnotatedError {
	return AnnotatedError{
		cause:  err,
		string: msg,
	}
}

// Annotation returns the additional description that goes with error.
func (err AnnotatedError) Annotation() string {
	return err.string
}

// Is reports whether any error in err's chain matches target.
func (err AnnotatedError) Is(wrappedErr error) bool {
	if err.cause != nil {

		if err.cause == wrappedErr {
			return true
		}

		switch err.cause.(type) {
		case AnnotatedError:
			return err.cause.(AnnotatedError).Is(wrappedErr)
		default:
			return false
		}
	}

	return false
}

func (err AnnotatedError) Cause() error  { return err.cause }
func (err AnnotatedError) Unwrap() error { return err.cause }
