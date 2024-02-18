package RemovingStarsFromAString

func removeStars(s string) string {
	var navInx int
	var stack []rune
	for navInx < len(s) {
		if s[navInx] == '*' {
			if newLength := len(stack) - 1; newLength >= 0 {
				stack = stack[:newLength]
			}
		} else {
			stack = append(stack, rune(s[navInx]))
		}
		navInx++
	}
	return string(stack)
}
