package bracket

// Handler ...
type Handler interface {
	IsValid(s string) bool
}

type bracket struct {
}

// IsValid ...
func (b *bracket) IsValid(s string) bool {
	var stack []rune
	var ok bool
	if len(s) == 0 {
		return true
	}
	for _, saab := range s {
		if openParenthesis(saab) == true {
			stack = append(stack, saab)
		} else {
			stack, ok = removeBrackets(stack, saab)
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

// openParenthesis check the bracket to make it open
func openParenthesis(s rune) bool {
	str := []rune{'(', '{', '['}
	for _, saab := range str {
		if saab == s {
			return true
		}
	}
	return false
}

// removeBrackets check the previous bracket to match the current one
func removeBrackets(stack []rune, symbol rune) ([]rune, bool) {
	i := len(stack)
	if i > 0 {
		if stack[i-1] == symbol-1 || stack[i-1] == symbol-2 { // the magic number is the difference between ASCII codes
			stack := append(stack[:i-1], stack[i:]...)
			return stack, true
		}
	}
	return stack, false
}

// NewHandler ...
func NewHandler() Handler{
	return &bracket{}
}
