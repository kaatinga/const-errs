package const_errs

import "testing"

const err1 Error = "text1"

func TestErrorWrapper_Is(t *testing.T) {

	tests := []struct {
		name           string
		annotatedError AnnotatedError
		err            error
		want           bool
	}{
		{"ok", AnnotatedError{
			cause: AnnotatedError{
				cause:  err1,
				string: "text 1",
			},
			string: "text 2",
		}, err1, true},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.annotatedError
			if got := err.Is(tt.err); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
