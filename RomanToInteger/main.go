package RomanToInteger

const (
	One         int = 1
	Five        int = 5
	Ten         int = 10
	Fifty       int = 50
	Hundred     int = 100
	FiveHundred int = 500
	Thousand    int = 1000

	I rune = 'I'
	V rune = 'V'
	X rune = 'X'
	L rune = 'L'
	C rune = 'C'
	D rune = 'D'
	M rune = 'M'
)

func wasPreviousBigger(curr rune, prev rune) bool {
	if curr == D && prev == M {
		return true
	} else if curr == C && (prev == D || prev == M) {
		return true
	} else if curr == L && (prev == C || prev == D || prev == M) {
		return true
	} else if curr == X && (prev == L || prev == C || prev == D || prev == M) {
		return true
	} else if curr == V && (prev == X || prev == L || prev == C || prev == D || prev == M) {
		return true
	} else if curr == I && (prev == V || prev == X || prev == L || prev == C || prev == D || prev == M) {
		return true
	}
	return false
}

func romanToInt(s string) int {
	var split = []rune(s)
	var index = len(s) - 1
	var total int
	var previousValue = split[len(split)-1]
	for index >= 0 {
		var valueToAdd int
		switch split[index] {
		case I:
			valueToAdd = One
		case V:
			valueToAdd = Five
		case X:
			valueToAdd = Ten
		case L:
			valueToAdd = Fifty
		case C:
			valueToAdd = Hundred
		case D:
			valueToAdd = FiveHundred
		case M:
			valueToAdd = Thousand
		}
		if wasPreviousBigger(split[index], previousValue) {
			valueToAdd = valueToAdd * -1
		}
		total += valueToAdd
		previousValue = split[index]
		index--
	}
	return total
}
