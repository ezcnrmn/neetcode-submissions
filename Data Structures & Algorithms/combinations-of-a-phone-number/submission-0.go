func letterCombinations(d string) []string {
	letter := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	n := len(d)
	if n == 0 {
		return []string{}
	}

	digits := make([]string, 0, len(d))
	for _, r := range d {
		digits = append(digits, string(r))
	}

	queue := []string{""}
	for len(queue[0]) < n {
		cur := queue[0]
		queue[0] = ""
		queue = queue[1:]

		letters := letter[digits[len(cur)]]

		for _, l := range letters {
			queue = append(queue, cur + l)
		}
	}

	return queue
}
