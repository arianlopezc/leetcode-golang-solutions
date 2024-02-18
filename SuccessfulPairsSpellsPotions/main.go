package SuccessfulPairsSpellsPotions

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	var result = make([]int, len(spells))
	var length = len(potions)
	for index, spell := range spells {
		var l, r = 0, length
		for l <= r {
			middle := (r + l) / 2
			var multi = int64(potions[middle] * spell)
			if leftLimit := middle-1 >= 0; leftLimit && multi >= success && int64(potions[middle-1]*spell) < success || (!leftLimit && multi >= success) {
				l, r = middle+1, middle
				result[index] = length - middle
			} else if middle == length-1 && int64(potions[middle]*spell) < success {
				l, r = middle+1, middle
				result[index] = 0
			} else if leftLimit && int64(potions[middle-1]*spell) >= success {
				r = middle - 1
			} else {
				l = middle + 1
			}
		}
	}
	return result
}
