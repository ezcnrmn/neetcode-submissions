func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens))
	for _, t := range tokens {
		parsed, err := strconv.Atoi(t)
		if err == nil {
			nums = append(nums, parsed)
		} else {
			n1, n2 := nums[len(nums)-1], nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			switch t {
			case "+":
				nums = append(nums, n2+n1)
			case "-":
				nums = append(nums, n2-n1)
			case "*":
				nums = append(nums, n2*n1)
			case "/":
				nums = append(nums, n2/n1)
			}
		}
	}
	return nums[len(nums)-1]
}
