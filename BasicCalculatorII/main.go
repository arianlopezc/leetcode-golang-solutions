package BasicCalculatorII

import (
	"unicode"
)

func calculate(s string) int {
	var stk []int
	curr := 0
	operator := '+'
	for i, ch := range s {
		if unicode.IsDigit(ch) {
			curr = (curr * 10) + int(ch-'0')

			if i != len(s)-1 {
				continue
			}
		}
		if ch == ' ' && i != len(s)-1 {
			continue
		}
		switch operator {
		case '+':
			stk = append(stk, curr)
		case '-':
			stk = append(stk, -curr)
		case '*':
			newNum := stk[len(stk)-1] * curr
			stk[len(stk)-1] = newNum
		case '/':
			newNum := stk[len(stk)-1] / curr
			stk[len(stk)-1] = newNum
		}
		operator = ch
		curr = 0
	}
	res := 0
	for _, el := range stk {
		res += el
	}
	return res
}
