package VerifyingAnAlienDictionary

import "testing"

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "standard", want: true, args: args{
				words: []string{"hello", "leetcode"},
				order: "hlabcdefgijkmnopqrstuvwxyz",
			},
		}, {
			name: "standard", want: false, args: args{
				words: []string{"word", "world", "row"},
				order: "worldabcefghijkmnpqstuvxyz",
			},
		}, {
			name: "standard", want: false, args: args{
				words: []string{"apple", "app"},
				order: "abcdefghijklmnopqrstuvwxyz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
