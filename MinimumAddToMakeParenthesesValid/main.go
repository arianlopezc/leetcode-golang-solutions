package MinimumAddToMakeParenthesesValid

func minAddToMakeValid(s string) int {
	var stacked []rune
	for _, letter := range s {
		if letter == '(' {
			stacked = append(stacked, letter)
		} else if letter == ')' {
			if len(stacked) == 0 || stacked[len(stacked)-1] == ')' {
				stacked = append(stacked, letter)
			} else {
				stacked = stacked[:len(stacked)-1]
			}
		}
	}
	return len(stacked)
}
