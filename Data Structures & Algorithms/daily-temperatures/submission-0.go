type maxTemp struct {
	index int
	value int
}

func dailyTemperatures(temps []int) []int {
	maxes := make([]maxTemp, 0, len(temps))
	maxes = append(maxes, maxTemp{index: len(temps)-1, value: temps[len(temps)-1]})
	result := make([]int, len(temps))

	for i := len(temps)-2; i >= 0; i-- {
		if temps[i] > temps[i+1] {
			for len(maxes) > 0 {
				if maxes[len(maxes)-1].value > temps[i] {
					break
				}
				maxes = maxes[:len(maxes)-1]
			}

			if len(maxes) != 0 {
				result[i] = maxes[len(maxes)-1].index - i
			}

			maxes = append(maxes, maxTemp{value: temps[i], index: i})
		} else if temps[i] < temps[i+1] {
			result[i] = 1
			maxes = append(maxes, maxTemp{value: temps[i], index: i})
		} else if result[i+1] > 0 {
			result[i] = result[i+1] + 1
		}
	}

	return result
}