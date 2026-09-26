func climbStairs(n int) int {
	a, b := 1, 1
	for range n-1 {
		a, b = b, a+b
	}
	return b
}
