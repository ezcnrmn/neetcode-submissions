func calPoints(operations []string) int {
	var score []int
	for _, op := range operations {
		switch op {
		case "D":
			score = append(score, score[len(score)-1]*2)
		case "C":
			score = score[:len(score)-1]
		case "+":
			score = append(score, score[len(score)-1] + score[len(score)-2])
		default:
			s, _ := strconv.Atoi(op)
			score = append(score, s)
		}
	}
	var acc int
	for _, s := range score {
		acc += s
	}
	return acc
}