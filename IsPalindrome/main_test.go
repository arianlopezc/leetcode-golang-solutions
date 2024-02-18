package IsPalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "valid", want: true, args: args{s: "A man, a plan, a canal: Panama"}},
		{name: "invalid", want: false, args: args{s: "race a car"}},
		{name: "invalid", want: false, args: args{s: "0D"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
