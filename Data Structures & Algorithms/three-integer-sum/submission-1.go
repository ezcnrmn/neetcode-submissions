import "slices"

func threeSum(nums []int) (res [][]int) {
	slices.Sort(nums)
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i] + nums[left] + nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left = moveIndex(nums, left, 1)
				right = moveIndex(nums, right, -1)
			} else if nums[i] + nums[left] + nums[right] > 0 {
				right = moveIndex(nums, right, -1)
			} else {
				left = moveIndex(nums, left, 1)
			}
		}
	}
	return
}

func moveIndex(nums []int, index, dir int) int {
	cur := index + dir
	for cur >= 0 && cur < len(nums) {
		if nums[cur] != nums[index] {
			break
		}
		cur += dir
	}
	return cur
}