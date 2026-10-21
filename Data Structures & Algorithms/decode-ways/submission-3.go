func numDecodings(s string) int {
	cache := make(map[string]int)

	var bt func(string) int
	bt = func(cur string) int {
		if len(cur) == 0 {
			return 1
		}
		if cur[0] == '0' {
			return 0
		}

		if _, ok := cache[cur]; !ok {
			amount := bt(cur[1:])
			if len(cur) > 1 {
				parsed, err := strconv.Atoi(cur[:2])
				if err == nil && parsed <= 26 {
					amount += bt(cur[2:])
				}
			}
			cache[cur] = amount
		}

		return cache[cur]
	}

	return bt(s)
}
