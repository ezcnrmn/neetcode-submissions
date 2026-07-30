import "slices"

func majorityElement(nums []int) int {
    slices.Sort(nums)
	return nums[len(nums)/2]
}
