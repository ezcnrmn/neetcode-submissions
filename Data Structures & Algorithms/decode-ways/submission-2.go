func numDecodings(s string) int {
	var result int
	stack := []string{s}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(cur) == 0 {
			result++
			continue
		}
		if cur[0] == '0' {
			continue
		}

		stack = append(stack, cur[1:])
		if len(cur) > 1 {
			parsed, _ := strconv.Atoi(cur[:2])
			if parsed <= 26 {
				stack = append(stack, cur[2:])
			}
		}
	}
	return result
}
