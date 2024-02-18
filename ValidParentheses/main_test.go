package ValidParentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "valid", want: true, args: args{s: "()[]{}"}},
		{name: "valid", want: true, args: args{s: "({[]})"}},
		{name: "invalid", want: false, args: args{s: "(]"}},
		{name: "invalid", want: false, args: args{s: "({[)]})"}},
		{name: "invalid", want: false, args: args{s: "("}},
		{name: "invalid", want: false, args: args{s: ")"}},
		{name: "invalid", want: false, args: args{s: "())"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
