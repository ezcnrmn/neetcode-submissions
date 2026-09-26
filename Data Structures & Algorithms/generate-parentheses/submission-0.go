func generateParenthesis(n int) []string {
	var result []string

	var bt func(string, int, int)
	bt = func(cur string, opened, closed int) {
		if opened == n && closed == n {
			result = append(result, cur)
			return
		}

		if opened < n {
			bt(cur + "(", opened + 1, closed)
		}
		if closed < n && opened > closed {
			bt(cur + ")", opened, closed + 1)
		}
	}

	bt("", 0, 0)
	return result
}
