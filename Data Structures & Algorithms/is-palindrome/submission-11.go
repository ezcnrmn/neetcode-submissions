func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for !isAlphNum(s[l]) {
			l++
		}
		for !isAlphNum(s[r]) {
			r--
		}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphNum(char byte) bool {
	return (char >= '0' && char <= '9') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= 'a' && char <= 'z')
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		diff := byte('a' - 'A')
		return char + diff
	}
	return char
}