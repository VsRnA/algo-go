package validparentheses

func IsValid(s string) bool {
	if len(s) % 2 != 0 {
			return false
	}

	stack := []rune{}
	pairs := map[rune]rune{')': '(',  ']': '[', '}' : '{'}

	for _, symbol := range s {
		if symbol == '(' || symbol == '[' || symbol == '{' {
			stack = append(stack, symbol)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[symbol] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
