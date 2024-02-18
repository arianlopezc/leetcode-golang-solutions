package MaximumNumberVowelsSubstringGivenLength

import "testing"

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", want: 3, args: args{k: 3, s: "abciiidef"}},
		{name: "test", want: 2, args: args{k: 2, s: "aeiou"}},
		{name: "test", want: 3, args: args{k: 3, s: "aei"}},
		{name: "test", want: 2, args: args{k: 3, s: "leetcode"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
