package main

func countConnections(groupSize int) int {
	totalConnection := 0
	for i := 1; i <= groupSize; i++ {
		totalConnection += (i - 1)
	}
	return totalConnection
}
