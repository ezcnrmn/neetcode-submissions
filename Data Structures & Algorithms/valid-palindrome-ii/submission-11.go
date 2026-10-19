func validPalindrome(s string) bool {
	return recVal(s, 1)
}

func recVal(s string, changes int) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			if changes == 0 {
				return false
			}
			return recVal(s[l+1:r+1], changes-1) || recVal(s[l:r], changes-1)
		}
		l++
		r--
	}
	return true
}