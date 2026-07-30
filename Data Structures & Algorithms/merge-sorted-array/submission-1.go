func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2, j := m-1, n-1, len(nums1)-1
	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[j] = nums1[i1]
			i1--
		} else {
			nums1[j] = nums2[i2]
			i2--
		}
		j--
	}
	copy(nums1[:j+1], nums2[:i2+1])
}
