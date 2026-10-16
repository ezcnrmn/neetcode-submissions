func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, word := range strs {
		var counter [26]int
		for _, char := range word {
			counter[char-'a']++
		}
		groups[counter] = append(groups[counter], word)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
