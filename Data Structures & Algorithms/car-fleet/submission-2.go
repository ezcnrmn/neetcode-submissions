import "slices"

func carFleet(target int, position []int, speed []int) int {
	cars := make([][2]int, len(position))
	for i := range position {
		cars[i] = [2]int{position[i], speed[i]}
	}
	slices.SortFunc(cars, func(a, b [2]int) int {
		return b[0]-a[0]
	})

	var fleets int
	var prev float64
	for _, car := range cars {
		hour := float64(target-car[0])/float64(car[1])
		if hour > prev {
			fleets++
			prev = hour
		}
	}

	return fleets
}
