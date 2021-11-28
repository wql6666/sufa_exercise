package easy

import "testing"

func Test_strIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strIsValid(tt.args.s); got != tt.want {
				t.Errorf("strIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
