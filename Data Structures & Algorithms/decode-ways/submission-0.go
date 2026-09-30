func numDecodings(s string) int {
	cache := make(map[string]int)

	var bt func(string) int
	bt = func (str string) int {
		if len(str) <= 1 {
			return 1
		}
		if str[0] == '0' {
			return 0
		}

		if cached, ok := cache[str]; ok {
			return cached
		}

		case1 := bt(str[1:])

		parsed, _ := strconv.Atoi(str[:2])
		if parsed > 26 {
			cache[str] = case1
			return case1
		}
		case2 := bt(str[2:])

		cache[str] = case1 + case2
		return cache[str]
	}

	return bt(s)
}
