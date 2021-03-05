package const_errs

import "testing"

const err1 Error = "text1"
const err2 Error = "text2"

func TestErrorWrapper_Is(t *testing.T) {

	tests := []struct {
		name   string
		fields ErrorWrapper
		err    error
		want   bool
	}{
		{"ok", ErrorWrapper{
			error:  err1,
			string: "text",
		}, err1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ErrorWrapper{
				error:  tt.fields.error,
				string: tt.fields.string,
			}
			if got := err.Is(tt.err); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
