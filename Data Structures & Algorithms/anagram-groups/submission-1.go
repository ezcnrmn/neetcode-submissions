type Letters [26]int

func groupAnagrams(strs []string) [][]string {
	mem := make(map[Letters][]string)
	for _, s := range strs {
		l := wordToLetters(s)
		mem[l] = append(mem[l], s)
	}
	var res [][]string
	for _, group := range mem {
		res = append(res, group)
	}
	return res
}

func wordToLetters(w string) Letters {
	var l Letters
	for _, r := range w {
		l[r-97]++
	}
	return l
}