func mergeAlternately(word1 string, word2 string) string {
	m := min(len(word1), len(word2))
	res := make([]byte, 0, len(word1)+len(word2))
	i := 0
	for ; i < m; i++ {
		res = append(res, word1[i], word2[i])
	}
	res = append(res, word1[i:]...)
	res = append(res, word2[i:]...)
	return string(res)
}