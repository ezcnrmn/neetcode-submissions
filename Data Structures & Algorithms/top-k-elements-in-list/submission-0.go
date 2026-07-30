func topKFrequent(nums []int, k int) []int {
	memo := make(map[int]int, len(nums))
	for _, number := range nums {
		memo[number]++
	}

	buckets := make([][]int, len(nums)+1)
	for number, amount := range memo {
		buckets[amount] = append(buckets[amount], number)
	}

	result := make([]int, 0, k)
	outer:
	for i := len(nums); i >= 0; i-- {
		for _, number := range buckets[i] {
			result = append(result, number)
			if len(result) == k {
				break outer
			}
		}
	}
	return result
}
