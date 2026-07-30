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
				left++
				right--
			} else if nums[i] + nums[left] + nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return
}