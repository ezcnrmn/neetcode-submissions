func mergeAlternately(word1 string, word2 string) string {
	m := min(len(word1), len(word2))
	res := ""
	i := 0
	for ; i < m; i++ {
		res += string(word1[i]) + string(word2[i])
	}
	res += word1[i:]
	res += word2[i:]
	return res
}