package main

func maxMessages(thresh int) int {
	baseCost := 100
	total := 0
	tMsg := 0

	for i := 0; ; i++ {
		cost := baseCost + (i * 1)
		total += cost
		if total > thresh {
			tMsg = i
			break
		}
	}
	return tMsg
}
