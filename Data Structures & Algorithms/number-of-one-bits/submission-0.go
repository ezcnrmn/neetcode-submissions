func hammingWeight(n int) int {
	var counter int
	for n > 0 {
		fmt.Println(n)
		if n%2 != 0 {
			counter++
		}
		n /= 2
	}
	return counter
}
