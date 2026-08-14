import "slices"

func lastStoneWeight(s []int) int {
	stones := make([]int, len(s))
	copy(stones, s)
	slices.Sort(stones)
	for len(stones) > 1 {
		s1, s2 := stones[len(stones)-1], stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if s1 > s2 {
			stones = append(stones, s1-s2)
		} else if s2 > s1 {
			stones = append(stones, s2-s1)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
