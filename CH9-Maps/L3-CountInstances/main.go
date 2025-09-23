package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if count, ok := validUsers[user]; ok {
			validUsers[user] = count + 1
		}
	}
}
