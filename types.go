package const_errs

type Error string
type Warning string

// Error returns error description.
func (err Error) Error() string {
	return string(err)
}

// Error returns error description.
func (err Warning) Error() string {
	return string(err)
}
