func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		n := min(len(prefix), len(strs[i]))
		j := 0
		for ; j < n; j++ {
			if prefix[j] != strs[i][j] {
				break
			}
		}
		prefix = prefix[:j]
	}

	return prefix
}