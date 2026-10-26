func canCompleteCircuit(gas []int, cost []int) int {
	var maxIdx int
	maxVal := gas[0]-cost[0]
	for i := 1; i < len(gas); i++ {
		if gas[i]-cost[i] > maxVal {
			maxVal = gas[i]-cost[i]
			maxIdx = i
		}
	}

	var tank int
	for i := 0; i < len(gas); i++ {
		cur := i + maxIdx
		if cur >= len(gas) {
			cur -= len(gas)
		}
		tank += gas[cur]-cost[cur]
		if tank < 0 {
			return -1
		}
	}

	return maxIdx
}
