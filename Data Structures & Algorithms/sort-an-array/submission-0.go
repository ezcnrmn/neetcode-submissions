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
	p := hoar(n)
	quickSort(n[:p])
	quickSort(n[p:])
}

func hoar(n []int) int {
	pivot := n[len(n)/2]
	l, r := 0, len(n)-1
	for {
		for l < r {
			if n[l] > pivot {
				break
			}
			l++
		}
		for r > l {
			if n[r] <= pivot {
				break
			}
			r--
		}
		if l >= r {
			break
		}
		n[l], n[r] = n[r], n[l]
	}
	return r
}