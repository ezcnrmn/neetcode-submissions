import "slices"

type car struct {
	speed float64
	pos   int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]car, len(speed))
	for i := range len(speed) {
		cars[i] = car{speed: float64(speed[i]), pos: position[i]}
	}

	slices.SortFunc(cars, func(a, b car) int {
		return a.pos - b.pos
	})

	var stack []float64
	for _, c := range cars {
		time := float64(target - c.pos) / c.speed

		for len(stack) > 0 && stack[len(stack)-1] <= time {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, time)
	}

	return len(stack)
}
