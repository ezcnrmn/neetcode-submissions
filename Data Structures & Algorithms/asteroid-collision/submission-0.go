func asteroidCollision(asteroids []int) []int {
	result := make([]int, 0, len(asteroids))

	i := 0
	for i < len(asteroids) {
		if asteroids[i] > 0 || len(result) == 0 {
			result = append(result, asteroids[i])
			i++
			continue
		}

		collisional := result[len(result)-1]
		if collisional + asteroids[i] < 0 {
			result = result[:len(result)-1]
		} else if collisional + asteroids[i] == 0 {
			result = result[:len(result)-1]
			i++
		} else {
			i++
		}
	}

	return result
}