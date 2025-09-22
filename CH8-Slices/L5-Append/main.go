package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	cst := []float64{}

	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			cst = append(cst, costs[i].value)
		}
	}

	return cst
}
