func isAnagram(s string, t string) bool {
	counter := make([]int, 26)
	for _, r := range s {
		counter[r-97]++
	}
	for _, r := range t {
		counter[r-97]--
	}
	for _, c := range counter {
		if c != 0 {
			return false
		}
	}
	return true
}
