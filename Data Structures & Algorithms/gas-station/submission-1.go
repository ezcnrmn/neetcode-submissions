func canCompleteCircuit(gas []int, cost []int) int {
	var prev, cur, idx int
	for i := range gas {
		dif := gas[i]-cost[i]
		if cur + dif < 0 {
			prev += cur + dif
			cur = 0
		} else {
			cur += dif
			idx = i
		}
	}

	if cur + prev >= 0 {
		return idx
	}
	return -1
}
