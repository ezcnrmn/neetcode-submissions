func isValid(s string) bool {
	stack := make([]rune, 0, len(s)/2)
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
			continue
		}
		if len(stack) == 0 {
			return false
		} 
		b := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if (r == ')' && b != '(') || (r == ']' && b != '[') || (r == '}' && b != '{') {
			return false
		}
	}
	return len(stack) == 0
}
