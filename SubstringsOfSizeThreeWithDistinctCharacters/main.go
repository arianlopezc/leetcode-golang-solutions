package SubstringsOfSizeThreeWithDistinctCharacters

func countGoodSubstrings(s string) int {
	var res int
	front := 2
	for front < len(s) {
		if s[front] != s[front-1] && s[front] != s[front-2] && s[front-1] != s[front-2] {
			res++
		}
		front++
	}
	return res
}
