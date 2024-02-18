package ValidPalindromeII

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "standard", want: true, args: args{s: "abca"},
		}, {
			name: "aba", want: true, args: args{s: "aba"},
		}, {
			name: "abc", want: false, args: args{s: "abc"},
		}, {
			name: "cbbcc", want: true, args: args{s: "cbbcc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
