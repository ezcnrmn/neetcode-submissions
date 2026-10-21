import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	counter := make(map[int]int, len(nums))
	for _, num := range nums {
		counter[num]++
	}

	var result [][]int
	for i := range nums {
		counter[nums[i]]--
		for j := i+1; j < len(nums); j++ {
			counter[nums[j]]--
			if counter[-(nums[i] + nums[j])] > 0 {
				result = append(result, []int{nums[i], nums[j], -(nums[i] + nums[j])})
				counter[-(nums[i] + nums[j])]--
			} else {
				counter[nums[j]]++
			}
		}
	}

	return result
}