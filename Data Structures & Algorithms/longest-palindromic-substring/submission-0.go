func longestPalindrome(s string) string {
	n := len(s)
	var result string

	for i := 0; i < n*2-1; i++ {
		var c1, c2 int
		c1 = i/2
		if i % 2 == 0 {
			c2 = c1
		} else {
			c2 = c1+1
		}

		for gap := range n {
			if c2 + gap >= n || c1 - gap < 0 || s[c2+gap] != s[c1-gap] {
				break
			}
			if len(s[c1-gap:c2+gap+1]) > len(result) {
				result = s[c1-gap:c2+gap+1]
			}
		}
	}

	return result
}
