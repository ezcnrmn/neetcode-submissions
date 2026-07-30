
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for !unicode.IsLetter(rune(s[i])) { i++ }
		for !unicode.IsLetter(rune(s[j])) { j-- }
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
