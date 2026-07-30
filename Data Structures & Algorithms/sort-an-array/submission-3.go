func sortArray(nums []int) []int {
	result := make([]int, len(nums))
	copy(result, nums)
	radixSort(result)
	return result
}

func radixSort(n []int) {
	maxEl := n[0]
	for i := 1; i < len(n); i++ {
		if maxEl < n[i] {
			maxEl = n[i]
		}
	}
	maxRadix := getRadix(maxEl)

	buf := make([]int, len(n))

	for r := 1; r <= maxRadix; r++ {
		var counter [10]int
		for _, num := range n {
			counter[getDigitForRadix(num, r)]++
		}

		for i := 1; i < len(counter); i++ {
			counter[i] += counter[i-1]
		}

		for i := len(n)-1; i >= 0; i-- {
			rad := getDigitForRadix(n[i], r)
			buf[counter[rad]-1] = n[i]
			counter[rad]--	
		}

		copy(n, buf)
	}
}

func getDigitForRadix(i, radix int) int {
	return (i/powOfTen(radix-1))%10
}

func getRadix(i int) (radix int) {
	for i > 0 {
		i /= 10
		radix++
	}
	return radix
}

func powOfTen(p int) int {
	switch p {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	case 4:
		return 10000
	case 5:
		return 100000
	}
	panic(":(")
}