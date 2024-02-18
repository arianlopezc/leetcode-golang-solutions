package ReverseWords

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "standard", want: "blue is sky the", args: args{s: "the sky is blue"}},
		{name: "standard", want: "world hello", args: args{s: "  hello world  "}},
		{name: "standard", want: "content tres suis Je", args: args{s: "Je suis tres content"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
