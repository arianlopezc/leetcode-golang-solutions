package LongestSubstringWithoutRepeatingCharacters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", want: 3, args: args{s: "abcabcbb"}},
		{name: "test", want: 1, args: args{s: "bbbbb"}},
		{name: "test", want: 3, args: args{s: "pwwkew"}},
		{name: "test", want: 14, args: args{s: "abcdefghijklmn"}},
		{name: "test", want: 5, args: args{s: "abcdeeee"}},
		{name: "test", want: 2, args: args{s: "aaccee"}},
		{name: "test", want: 4, args: args{s: "aacrtcee"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
