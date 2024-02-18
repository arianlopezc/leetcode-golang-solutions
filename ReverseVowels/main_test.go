package ReverseVowels

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "standard", want: "leotcede", args: args{s: "leetcode"}},
		{name: "empty", want: "", args: args{s: ""}},
		{name: "only consonants", want: "bbb", args: args{s: "bbb"}},
		{name: "only vowels", want: "aei", args: args{s: "iea"}},
		{name: "only one", want: "a", args: args{s: "a"}},
		{name: "standard", want: "holle", args: args{s: "hello"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
