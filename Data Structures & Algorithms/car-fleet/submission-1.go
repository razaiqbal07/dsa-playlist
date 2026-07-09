func carFleet(target int, position []int, speed []int) int {
	n := len(position)

	type Car struct {
		pos   int
		speed int
	}

	cars := make([]Car, n)
	for i := 0; i < n; i++ {
		cars[i] = Car{position[i], speed[i]}
	}

	// Sort by position descending (closest to target first)
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 0
	lastTime := 0.0

	for _, car := range cars {
		time := float64(target-car.pos) / float64(car.speed)

		if time > lastTime {
			fleets++
			lastTime = time
		}
	}

	return fleets
}