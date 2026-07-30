func sortArray(nums []int) []int {
	result := make([]int, len(nums))
	copy(result, nums)
	quickSort(result)
	return result
}

func quickSort(n []int) {
	if len(n) < 2 {
		return
	}
	p := lomuto(n)
	quickSort(n[:p])
	quickSort(n[p+1:])
}

func lomuto(n []int) int {
	pivot := n[len(n)-1]
	i := 0
	for j := range len(n)-1 {
		if n[j] <= pivot {
			n[i], n[j] = n[j], n[i] 
			i++
		}
	}
	n[i], n[len(n)-1] = n[len(n)-1], n[i]
	return i
}