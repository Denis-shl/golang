package bracket

func openParenthesis(s rune) bool {
	str := []rune{'(', '{', '['}
	for _, symb := range str {
		if symb == s {
			return true
		}
	}
	return false
}

func removeBrackets(stack []rune, symbol rune) ([]rune, bool) {
	i := len(stack)
	if i > 0 {
		if stack[i-1] == symbol-1 || stack[i-1] == symbol-2 {
			stack := append(stack[:i-1], stack[i:]...)
			return stack, true
		}
	}
	return stack, false

}

func isValid(s string) bool {
	var stack []rune
	var ok bool

	if len(s) == 0 {
		return true
	}
	for _, symb := range s {
		if openParenthesis(symb) == true {
			stack = append(stack, symb)
		} else {
			stack, ok = removeBrackets(stack, symb)
			if ok == false {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
