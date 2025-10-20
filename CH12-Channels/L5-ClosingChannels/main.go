package main

func countReports(numSentCh chan int) int {
	num := 0

	for {
		v, ok := <-numSentCh

		if !ok {
			break
		}
		num += v
	}
	return num
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
