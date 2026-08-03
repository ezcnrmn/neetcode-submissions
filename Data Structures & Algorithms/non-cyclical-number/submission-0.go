func isHappy(n int) bool {
	mem := make(map[int]struct{})
	for {
		var temp int
		for n > 0 {
			temp += n%10 * n%10
			n /= 10
		}
		if temp == 1 {
			return true
		}
		if _, ok := mem[temp]; ok {
			break
		}
		mem[temp] = struct{}{}
	}
	return false
}
