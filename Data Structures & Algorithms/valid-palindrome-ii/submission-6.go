func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	err := false
	for i < j {
		if s[i] != s[j] {
			if err {
				return false
			}
			if j - i < 2 {
			} else if s[i+1] == s[j] && s[i] != s[j-1]  {
				i++
			} else if s[i] == s[j-1] && s[i+1] != s[j] {
				j--
			} else {
				return false
			}
			err = true
		}
		i++
		j--
	}
	return true
}