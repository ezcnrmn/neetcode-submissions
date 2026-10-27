import "slices"

func carFleet(target int, position []int, speed []int) int {
	cars := make([][2]int, len(position))
	for i := range position {
		cars[i] = [2]int{position[i], speed[i]}
	}
	slices.SortFunc(cars, func(a, b [2]int) int {
		return b[0]-a[0]
	})

	var fleets, prev int
	for _, car := range cars {
		hour := div(target-car[0], car[1])
		if hour > prev {
			fleets++
			prev = hour
		}
	}

	return fleets
}

func div(x, y int) int {
	if x%y != 0 {
		return x/y+1
	}
	return x/y
}