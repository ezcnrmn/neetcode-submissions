func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(n []int) []int {
	if len(n) < 2 {
		return n
	}
	mid := len(n)/2
	l := mergeSort(n[:mid])
	r := mergeSort(n[mid:])
	return merge(l, r)
}

func merge(l, r []int) []int {
	sorted := make([]int, 0, len(l) + len(r))
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] > r[j] {
			sorted = append(sorted, r[j])
			j++
		} else {
			sorted = append(sorted, l[i])
			i++
		}
	}
	if i < len(l) {
		sorted = append(sorted, l[i:]...)
	}
	if j < len(r) {
		sorted = append(sorted, r[j:]...)
	}
	return sorted
}