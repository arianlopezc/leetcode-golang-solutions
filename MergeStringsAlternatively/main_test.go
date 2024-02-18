package MergeStringsAlternatively

import "testing"

func Test_mergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "normal", want: "apbqcr", args: struct {
			word1 string
			word2 string
		}{word1: "abc", word2: "pqr"}},
		{name: "normal", want: "apbqrs", args: struct {
			word1 string
			word2 string
		}{word1: "ab", word2: "pqrs"}},
		{name: "normal", want: "apbqcd", args: struct {
			word1 string
			word2 string
		}{word1: "abcd", word2: "pq"}},
		{name: "normal", want: "abcd", args: struct {
			word1 string
			word2 string
		}{word1: "abcd", word2: ""}},
		{name: "normal", want: "abcd", args: struct {
			word1 string
			word2 string
		}{word1: "", word2: "abcd"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
