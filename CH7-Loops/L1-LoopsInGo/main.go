package main

func bulkSend(numMessages int) float64 {
	baseCost := 1.0
	addFee := 0.01
	totalCost := 0.0

	for i := 0; i < numMessages; i++ {
		cost := baseCost + (float64(i) * addFee)
		totalCost += cost

	}
	return totalCost
}
