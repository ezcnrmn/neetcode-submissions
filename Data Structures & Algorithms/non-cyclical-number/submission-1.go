func isHappy(n int) bool {
	mem := make(map[int]struct{})
	mem[n] = struct{}{}
	for {
		var temp int
		for n > 0 {
			temp += (n%10) * (n%10)
			n /= 10
		}
		n = temp
		if n == 1 {
			return true
		}
		if _, ok := mem[n]; ok {
			break
		}
		mem[n] = struct{}{}
	}
	return false
}
