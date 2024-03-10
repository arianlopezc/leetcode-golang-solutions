package VerifyingAnAlienDictionary

import (
	"sort"
)

func isAlienSorted(words []string, order string) bool {
	var alphabet = make(map[rune]int)
	for index := 0; index < len(order); index++ {
		alphabet[rune(order[index])] = index
	}
	var res bool
	res = sort.SliceIsSorted(words, func(a, b int) bool {
		var length = min(len(words[a]), len(words[b]))
		for index := 0; index < length; index++ {
			if words[a][index] != words[b][index] {
				return alphabet[rune(words[a][index])] < alphabet[rune(words[b][index])]
			}
		}
		return len(words[a]) < len(words[b])
	})
	return res
}
