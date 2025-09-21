package main

func getMessageCosts(messages []string) []float64 {
	msgLen := len(messages)

	costSlice := make([]float64, msgLen)

	for i := 0; i < msgLen; i++ {
		costOfMsg := float64(len(messages[i])) * 0.01
		costSlice[i] = costOfMsg
	}

	return costSlice
}
