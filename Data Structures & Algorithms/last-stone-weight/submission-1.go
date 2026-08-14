import "slices"

func lastStoneWeight(s []int) int {
	stones := make([]int, len(s))
	copy(stones, s)
	slices.Sort(stones)
	for len(stones) > 1 {
		s1, s2 := stones[len(stones)-1], stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if s1 == s2 {
			continue
		}

		var newStone int
		if s1 > s2 {
			newStone = s1-s2
		} else if s2 > s1 {
			newStone = s2-s1
		}

		insertionIndex := binarySearch(stones, newStone)
		stones = append(stones, 0)
		for i := len(stones)-1; i >= insertionIndex+1; i-- {
			stones[i] = stones[i-1]
		}
		stones[insertionIndex] = newStone
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target || mid == len(nums)-1 {
			return mid
		}

		if nums[mid] > target {
			right = mid-1
		} else {
			left = mid+1
		}
	}

	return 0
}