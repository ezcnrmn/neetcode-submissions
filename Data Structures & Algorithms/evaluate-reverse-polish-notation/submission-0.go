func evalRPN(tokens []string) int {
	calc := make([]int, 0, len(tokens))
	for _, t := range tokens {
		switch t {
		case "+":
			prev1, prev2 := calc[len(calc)-1], calc[len(calc)-2]
			calc = append(calc, prev1+prev2)
		case "-":
			prev1, prev2 := calc[len(calc)-1], calc[len(calc)-2]
			calc = append(calc, prev1-prev2)
		case "*":
			prev1, prev2 := calc[len(calc)-1], calc[len(calc)-2]
			calc = append(calc, prev1*prev2)
		case "/":
			prev1, prev2 := calc[len(calc)-1], calc[len(calc)-2]
			calc = append(calc, prev1/prev2)
		default:
			parsed, _ := strconv.Atoi(t)
			calc = append(calc, parsed)
		}
	}
	return calc[len(calc)-1]
}
