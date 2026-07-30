func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !check(rune(s[i])) { i++ }
		for i < j && !check(rune(s[j])) { j-- }

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func check(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}