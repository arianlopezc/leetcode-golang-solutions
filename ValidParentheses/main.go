package ValidParentheses

func isValid(s string) bool {
	var letters []rune = []rune(s)
	const (
		openKey          rune = '{'
		closeKey         rune = '}'
		openParentheses  rune = '('
		closeParentheses rune = ')'
		openBrackets     rune = '['
		closeBrackets    rune = ']'
	)
	var stack []rune = make([]rune, len(s))
	var index int = 0
	for _, letter := range letters {
		if letter == openKey || letter == openParentheses || letter == openBrackets {
			stack[index] = letter
			index++
		} else {
			index--
			if index < 0 || letter == closeKey && stack[index] != openKey || letter == closeParentheses && stack[index] != openParentheses || letter == closeBrackets && stack[index] != openBrackets {
				return false
			}
		}
	}
	return index == 0
}
