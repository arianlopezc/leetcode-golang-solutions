package LongestCommonPrefix

import "sort"

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	first, last := strs[0], strs[len(strs)-1]
	length := len(first)
	if len(last) < length {
		length = len(last)
	}
	result := make([]rune, 0)
	var index int
	for index < length {
		if first[index] != last[index] {
			return string(result)
		}
		result = append(result, rune(first[index]))
		index++
	}
	return string(result)
}
