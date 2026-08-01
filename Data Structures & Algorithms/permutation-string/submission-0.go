func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1counter, s2counter [26]int
	for i := 0; i < len(s1); i++ {
		s1counter[s1[i]-'a']++
	}

	var left int
	for right := 0; right < len(s2); right++ {
		s2counter[s2[right]-'a']++
		if right-left+1 > len(s1) {
			s2counter[s2[left]-'a']--
			left++
		}
		if right-left+1 == len(s1) && s1counter == s2counter {
			return true
		}
	}
	return false
}
